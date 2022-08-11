package zksync2

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
)

type EthSigner interface {
	GetAddress() common.Address
	GetDomain() *Eip712Domain
	SignHash(msg []byte) ([]byte, error)
	//SignTypedData(d *Eip712Domain,
}

type DefaultEthSigner struct {
	pk      *ecdsa.PrivateKey
	address common.Address
	domain  *Eip712Domain
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
		domain:  DefaultEip712Domain(chainId),
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
		domain:  DefaultEip712Domain(chainId),
	}, nil
}

func (s *DefaultEthSigner) GetAddress() common.Address {
	return s.address
}

func (s *DefaultEthSigner) GetDomain() *Eip712Domain {
	return s.domain
}

func (s *DefaultEthSigner) SignHash(msg []byte) ([]byte, error) {
	sig, err := crypto.Sign(msg, s.pk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign hash")
	}
	return sig, nil
}
