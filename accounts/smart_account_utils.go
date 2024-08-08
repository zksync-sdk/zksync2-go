package accounts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
)

// SignPayloadWithECDSA signs the payload using an ECDSA private key.
var SignPayloadWithECDSA PayloadSigner = func(ctx context.Context, payload []byte, secret interface{}, client *clients.Client) ([]byte, error) {
	privateKeyHex, ok := (secret).(string)
	if !ok {
		return nil, errors.New("secret should be hex of ECDSA private key")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	signature, err := crypto.Sign(payload, privateKey)
	if err != nil {
		return nil, err
	}
	// crypto.Sign uses the traditional implementation where v is either 0 or 1,
	// while Ethereum uses newer implementation where v is either 27 or 28.
	if signature[64] < 27 {
		signature[64] += 27
	}
	return signature, nil
}

// SignPayloadWithMultipleECDSA signs the payload using multiple ECDSA private keys.
// The signature is generated by concatenating signatures created by signing with each key individually.
// The length of the resulting signature is len(secret) * 65.
var SignPayloadWithMultipleECDSA PayloadSigner = func(ctx context.Context, payload []byte, secret interface{}, client *clients.Client) ([]byte, error) {
	privateKeys, ok := (secret).([]string)
	if !ok {
		return nil, errors.New("secret should be a slice of ECDSA private keys")
	}
	if len(privateKeys) < 2 {
		return nil, errors.New("secret should be a slice of ECDSA private keys")
	}
	multiSig := make([]byte, len(privateKeys)*65)
	for i, privateKeyHex := range privateKeys {
		privateKey, err := crypto.HexToECDSA(privateKeyHex)
		if err != nil {
			return nil, err
		}
		signature, err := crypto.Sign(payload, privateKey)
		if err != nil {
			return nil, err
		}
		// crypto.Sign uses the traditional implementation where v is either 0 or 1,
		// while Ethereum uses newer implementation where v is either 27 or 28.
		if signature[64] < 27 {
			signature[64] += 27
		}
		copy(multiSig[i*65:(i+1)*65], signature)
	}
	return multiSig, nil
}

// PopulateTransactionECDSA Populates missing properties meant for signing using an ECDSA private key:
//
//   - Populates tx.From using the address derived from the ECDSA private key.
//   - Populates tx.Nonce via client.NonceAt().
//   - Populates tx.Gas via client.EstimateFee().GasLimit. If tx.From is not EOA, the estimation is done with address
//     derived from the ECDSA private key.
//   - Populates tx.GasFeeCap via client.EstimateFee().MaxFeePerGas.
//   - Populates tx.GasTipCap with 0 if is not set.
//   - Populates tx.ChainID via client.ChainID().
//   - Populates tx.Data with "0x".
//   - Populates tx.Meta.GasPerPubdata with utils.DefaultGasPerPubdataLimit.
//
// Expects the secret to be ECDSA private in hex format.
var PopulateTransactionECDSA TransactionBuilder = func(ctx context.Context, tx *types.Transaction712, secret interface{}, client *clients.Client) error {
	var err error
	if client == nil {
		return errors.New("client is required but is not provided")
	}

	if tx.ChainID == nil {
		if tx.ChainID, err = client.ChainID(ensureContext(ctx)); err != nil {
			return err
		}
	}
	if tx.Nonce == nil {
		nonce, err := client.NonceAt(ensureContext(ctx), *tx.From, nil)
		if err != nil {
			return fmt.Errorf("failed to get nonce: %w", err)
		}
		tx.Nonce = new(big.Int).SetUint64(nonce)
	}
	if tx.GasTipCap == nil {
		tx.GasTipCap = common.Big0
	}
	if tx.Meta == nil {
		tx.Meta = &types.Eip712Meta{GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64())}
	} else if tx.Meta.GasPerPubdata == nil {
		tx.Meta.GasPerPubdata = utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64())
	}
	if (tx.Gas == nil || tx.Gas.Uint64() == 0) || (tx.GasFeeCap == nil) {
		from := *tx.From
		// Gas estimation does not work when initiator is contract account (works only with EOA).
		// In order to estimation gas, the transaction's from value is replaced with signer's address.
		if bytecode, err := client.CodeAt(ensureContext(ctx), *tx.From, nil); len(bytecode) > 0 {
			privateKeyHex, ok := (secret).(string)
			if !ok {
				return errors.New("secret should be hex of ECDSA private key")
			}
			privateKey, err := crypto.HexToECDSA(privateKeyHex)
			if err != nil {
				return err
			}
			publicKey := privateKey.Public().(*ecdsa.PublicKey)
			from = crypto.PubkeyToAddress(*publicKey)
		} else if err != nil {
			return err
		}

		fee, err := client.EstimateFee(ensureContext(ctx), types.CallMsg{
			From:      from,
			To:        tx.To,
			GasFeeCap: tx.GasFeeCap,
			GasTipCap: tx.GasTipCap,
			Value:     tx.Value,
			Data:      tx.Data,
			Meta:      tx.Meta,
		})
		if err != nil {
			return fmt.Errorf("failed to EstimateFee: %w", err)
		}

		if tx.Gas == nil || tx.Gas.Uint64() == 0 {
			tx.Gas = fee.GasLimit.ToInt()
		}
		if tx.GasFeeCap == nil || tx.GasFeeCap.Uint64() == 0 {
			tx.GasFeeCap = fee.MaxFeePerGas.ToInt()
		}
	}
	if tx.Data == nil {
		tx.Data = hexutil.Bytes{}
	}
	return nil
}

// PopulateTransactionMultipleECDSA populates missing properties meant for signing using multiple ECDSA private keys.
// It uses PopulateTransactionECDSA, where the address of the first ECDSA key is set as the secret argument.
// Expects the secret to be a slice of ECDSA private in hex format.
var PopulateTransactionMultipleECDSA TransactionBuilder = func(ctx context.Context, tx *types.Transaction712, secret interface{}, client *clients.Client) error {
	privateKeys, ok := (secret).([]string)
	if !ok {
		return errors.New("secret should be a slice of ECDSA private keys")
	}
	if len(privateKeys) < 2 {
		return errors.New("secret should be a slice of ECDSA private keys")
	}
	// estimates gas accepts only one address, so the first signer is chosen.
	return PopulateTransactionECDSA(ctx, tx, privateKeys[0], client)
}
