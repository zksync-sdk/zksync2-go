package accounts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/stephenlacy/go-ethereum-hdwallet"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
)

// Signer provides support for signing various types of payloads using some kind of secret.
type Signer interface {
	// SignMessage sings an arbitrary message.
	SignMessage(ctx context.Context, message []byte) ([]byte, error)
	// SignTransaction signs the given transaction.
	SignTransaction(ctx context.Context, tx *types.Transaction) ([]byte, error)
	// SignTypedData signs the given EIP712 typed data.
	SignTypedData(ctx context.Context, typedData *apitypes.TypedData) ([]byte, error)
}

// ECDSASigner represents basis implementation of Signer interface.
type ECDSASigner struct {
	pk      *ecdsa.PrivateKey
	address common.Address
	chainId *big.Int
}

// NewECDSASignerFromMnemonic creates a new instance of ECDSASigner based on the provided mnemonic phrase.
func NewECDSASignerFromMnemonic(mnemonic string, chainId *big.Int) (*ECDSASigner, error) {
	return NewECDSASignerFromMnemonicAndAccountId(mnemonic, 0, chainId)
}

// NewECDSASignerFromMnemonicAndAccountId creates a new instance of ECDSASigner based on the provided mnemonic phrase and
// account ID.
func NewECDSASignerFromMnemonicAndAccountId(mnemonic string, accountId uint32, chainId *big.Int) (*ECDSASigner, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HD wallet from mnemonic")
	}
	path, err := accounts.ParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", accountId))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse derivation path")
	}
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to derive account from HD wallet")
	}
	pk, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account's private key from HD wallet")
	}
	pub := pk.Public().(*ecdsa.PublicKey)
	return &ECDSASigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
		chainId: chainId,
	}, nil
}

// NewECDSASignerFromRawPrivateKey creates a new instance of ECDSASigner based on the provided raw private key.
func NewECDSASignerFromRawPrivateKey(rawPk []byte, chainId *big.Int) (*ECDSASigner, error) {
	pk, err := crypto.ToECDSA(rawPk)
	if err != nil {
		return nil, errors.Wrap(err, "invalid raw private key")
	}
	pub := pk.Public().(*ecdsa.PublicKey)
	return &ECDSASigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
		chainId: chainId,
	}, nil
}

// NewRandomBaseSigner creates an instance of Signer with a randomly generated private key.
func NewRandomBaseSigner(chainId *big.Int) (*ECDSASigner, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate radnom private key: %w", err)
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to convert public key to ECDSA")
	}
	return &ECDSASigner{
		pk:      privateKey,
		address: crypto.PubkeyToAddress(*publicKey),
		chainId: chainId,
	}, nil
}

// Address returns the address of the associated account.
func (s *ECDSASigner) Address() common.Address {
	return s.address
}

// ChainID returns the chain ID of the associated account.
func (s *ECDSASigner) ChainID() *big.Int {
	return s.chainId
}

// PrivateKey returns the private key of the associated account.
func (s *ECDSASigner) PrivateKey() *ecdsa.PrivateKey {
	return s.pk
}

// SignMessage signs the message using the private key.
func (s *ECDSASigner) SignMessage(_ context.Context, msg []byte) ([]byte, error) {
	sig, err := crypto.Sign(msg, s.pk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign hash")
	}
	return sig, nil
}

// SignTransaction signs the ZKsync transaction.
func (s *ECDSASigner) SignTransaction(ctx context.Context, tx *types.Transaction) ([]byte, error) {
	typedData, err := tx.TypedData()
	if err != nil {
		return nil, err
	}
	return s.SignTypedData(ctx, typedData)
}

// SignTypedData signs the typed data.
func (s *ECDSASigner) SignTypedData(_ context.Context, typedData *apitypes.TypedData) ([]byte, error) {
	hash, err := s.hashTypedData(typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of typed data: %w", err)
	}
	sig, err := crypto.Sign(hash, s.pk)
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash of typed data: %w", err)
	}
	// crypto.Sign uses the traditional implementation where v is either 0 or 1,
	// while Ethereum uses newer implementation where v is either 27 or 28.
	if sig[64] < 27 {
		sig[64] += 27
	}
	return sig, nil
}

func (s *ECDSASigner) hashTypedData(data *apitypes.TypedData) ([]byte, error) {
	domain, err := data.HashStruct("EIP712Domain", data.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of typed data domain: %w", err)
	}
	dataHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of typed message: %w", err)
	}
	prefixedData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domain), string(dataHash)))
	prefixedDataHash := crypto.Keccak256(prefixedData)
	return prefixedDataHash, nil
}
