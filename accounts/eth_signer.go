package accounts

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/eip712"
)

type EthSigner interface {
	GetAddress() common.Address
	GetDomain() *eip712.Domain
	SignHash(msg []byte) ([]byte, error)
	SignTypedData(d *eip712.Domain, data eip712.TypedData) ([]byte, error)
}

type DefaultEthSigner struct {
	pk      *ecdsa.PrivateKey
	address common.Address
	domain  *eip712.Domain
}

func NewEthSignerFromMnemonic(mnemonic string, chainId int64) (*DefaultEthSigner, error) {
	return NewEthSignerFromMnemonicAndAccountId(mnemonic, 0, chainId)
}

func NewEthSignerFromMnemonicAndAccountId(mnemonic string, accountId uint32, chainId int64) (*DefaultEthSigner, error) {
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
	return &DefaultEthSigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
		domain:  eip712.DefaultEip712Domain(chainId),
	}, nil
}

func NewEthSignerFromRawPrivateKey(rawPk []byte, chainId int64) (*DefaultEthSigner, error) {
	pk, err := crypto.ToECDSA(rawPk)
	if err != nil {
		return nil, errors.Wrap(err, "invalid raw private key")
	}
	pub := pk.Public().(*ecdsa.PublicKey)
	return &DefaultEthSigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
		domain:  eip712.DefaultEip712Domain(chainId),
	}, nil
}

func (s *DefaultEthSigner) GetAddress() common.Address {
	return s.address
}

func (s *DefaultEthSigner) GetDomain() *eip712.Domain {
	return s.domain
}

func (s *DefaultEthSigner) SignTypedData(domain *eip712.Domain, data eip712.TypedData) ([]byte, error) {
	// compile TypedData structure
	eip712Msg, err := data.GetEIP712Message()
	if err != nil {
		return nil, err
	}
	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			data.GetEIP712Type():   data.GetEIP712Types(),
			domain.GetEIP712Type(): domain.GetEIP712Types(),
		},
		PrimaryType: data.GetEIP712Type(),
		Domain:      domain.GetEIP712Domain(),
		Message:     eip712Msg,
	}
	hash, err := s.HashTypedData(typedData)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of typed data: %w", err)
	}
	sig, err := crypto.Sign(hash, s.pk)
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash of typed data: %w", err)
	}
	if sig[64] < 27 {
		sig[64] += 27
	}
	return sig, nil
}

func (s *DefaultEthSigner) HashTypedData(data apitypes.TypedData) ([]byte, error) {
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

func (s *DefaultEthSigner) SignHash(msg []byte) ([]byte, error) {
	sig, err := crypto.Sign(msg, s.pk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign hash")
	}
	return sig, nil
}
