package accounts

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/stephenlacy/go-ethereum-hdwallet"
	"github.com/zksync-sdk/zksync2-go/eip712"
)

// Signer provides support for signing EIP-712 transactions as well as other types of transactions supported by
// types.Signer.
type Signer interface {
	// Address returns the  address associated with the signer.
	Address() common.Address
	// Domain returns the EIP-712 domain used for signing.
	Domain() *eip712.Domain
	// PrivateKey returns the private key associated with the signer.
	PrivateKey() *ecdsa.PrivateKey
	// SignHash signs the given hash using the signer's private key and returns the signature.
	// The hash should be the 32-byte hash of the data to be signed.
	SignHash(msg []byte) ([]byte, error)
	// SignTypedData signs the given EIP-712 typed data using the signer's private key and returns the signature.
	// The domain parameter is the EIP-712 domain separator, and the data parameter is the EIP-712 typed data.
	SignTypedData(typedData apitypes.TypedData) ([]byte, error)
}

// BaseSigner represents basis implementation of Signer interface.
type BaseSigner struct {
	pk      *ecdsa.PrivateKey
	address common.Address
	domain  *eip712.Domain
}

// NewBaseSignerFromMnemonic creates a new instance of BaseSigner based on the provided mnemonic phrase.
func NewBaseSignerFromMnemonic(mnemonic string, chainId int64) (*BaseSigner, error) {
	return NewBaseSignerFromMnemonicAndAccountId(mnemonic, 0, chainId)
}

// NewBaseSignerFromMnemonicAndAccountId creates a new instance of BaseSigner based on the provided mnemonic phrase and
// account ID.
func NewBaseSignerFromMnemonicAndAccountId(mnemonic string, accountId uint32, chainId int64) (*BaseSigner, error) {
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
	return &BaseSigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
		domain:  eip712.ZkSyncEraEIP712Domain(chainId),
	}, nil
}

// NewBaseSignerFromRawPrivateKey creates a new instance of BaseSigner based on the provided raw private key.
func NewBaseSignerFromRawPrivateKey(rawPk []byte, chainId int64) (*BaseSigner, error) {
	pk, err := crypto.ToECDSA(rawPk)
	if err != nil {
		return nil, errors.Wrap(err, "invalid raw private key")
	}
	pub := pk.Public().(*ecdsa.PublicKey)
	return &BaseSigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
		domain:  eip712.ZkSyncEraEIP712Domain(chainId),
	}, nil
}

// NewRandomBaseSigner creates an instance of Signer with a randomly generated private key.
func NewRandomBaseSigner(chainId int64) (*BaseSigner, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate radnom private key: %w", err)
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to convert public key to ECDSA")
	}
	return &BaseSigner{
		pk:      privateKey,
		address: crypto.PubkeyToAddress(*publicKey),
		domain:  eip712.ZkSyncEraEIP712Domain(chainId),
	}, nil
}

func (s *BaseSigner) Address() common.Address {
	return s.address
}

func (s *BaseSigner) Domain() *eip712.Domain {
	return s.domain
}

func (s *BaseSigner) PrivateKey() *ecdsa.PrivateKey {
	return s.pk
}

func (s *BaseSigner) SignTypedData(typedData apitypes.TypedData) ([]byte, error) {
	hash, err := s.HashTypedData(typedData)
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

func (s *BaseSigner) HashTypedData(data apitypes.TypedData) ([]byte, error) {
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

func (s *BaseSigner) SignHash(msg []byte) ([]byte, error) {
	sig, err := crypto.Sign(msg, s.pk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign hash")
	}
	return sig, nil
}
