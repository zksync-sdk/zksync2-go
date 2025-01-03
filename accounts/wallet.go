package accounts

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/clients"
	"math/big"
)

// Wallet wraps all operations that interact with an associated account.
// An account typically contains a private key, allowing it to sign various types of payloads.
type Wallet struct {
	WalletL1
	WalletL2
	Deployer

	clientL1 *ethclient.Client
	clientL2 *clients.Client
}

// NewWallet creates an instance of Wallet associated with the account provided by the rawPrivateKey.
// The clientL1 parameters is optional; if not provided, only method form WalletL2 and Deployer can be used,
// as the rest of the functionalities require communication with the L1 network.
// A Wallet can be configured to communicate with L1 networks by using and Wallet.ConnectL1 method.
func NewWallet(rawPrivateKey []byte, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	chainID, err := clientL2.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	signer, err := NewECDSASignerFromRawPrivateKey(rawPrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	return NewWalletFromSigner(signer, clientL2, clientL1)
}

// NewWalletFromSigner creates an instance of Wallet associated with the account provided by the signer.
// The clientL2 and clientL1 parameters are optional; if not provided, only WalletL2.SignTransaction,
// WalletL2.Address, WalletL2.Signer methods can be used, as the rest of the functionalities
// require communication with the network.
// A runner that contains only a signer can be configured to communicate with L2 and L1 networks by
// using Wallet.Connect and Wallet.ConnectL1, respectively.
func NewWalletFromSigner(signer *ECDSASigner, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	if signer == nil {
		return nil, errors.New("signer must be provided")
	}
	var (
		walletL1 *WalletL1
		walletL2 *WalletL2
		err      error
	)

	cache := NewCache(clientL2, clientL1)

	walletL1 = new(WalletL1)
	if clientL1 != nil {
		if walletL1, err = NewWalletL1FromSignerAndCache(signer, clientL1, clientL2, cache); err != nil {
			return nil, err
		}
	}
	if walletL2, err = NewWalletL2FromSignerAndCache(signer, clientL2, cache); err != nil {
		return nil, err
	}
	return &Wallet{
		WalletL1: *walletL1,
		WalletL2: *walletL2,
		Deployer: *NewDeployer(walletL2),
		clientL1: clientL1,
		clientL2: clientL2,
	}, nil
}

// NewWalletFromMnemonic creates a new instance of Wallet based on the provided mnemonic phrase.
// The clientL2 and clientL1 parameters are optional, and can be configured with Wallet.Connect
// and Wallet.ConnectL1, respectively.
func NewWalletFromMnemonic(mnemonic string, chainId *big.Int, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	signer, err := NewECDSASignerFromMnemonic(mnemonic, chainId)
	if err != nil {
		return nil, err
	}
	return NewWalletFromSigner(signer, clientL2, clientL1)
}

// NewWalletFromRawPrivateKey creates a new instance of Wallet based on the provided private key
// of the account and chain ID.
// The clientL2 and clientL1 parameters are optional, and can be configured with Wallet.Connect
// and Wallet.ConnectL1, respectively.
func NewWalletFromRawPrivateKey(rawPk []byte, chainId *big.Int, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	signer, err := NewECDSASignerFromRawPrivateKey(rawPk, chainId)
	if err != nil {
		return nil, err
	}
	return NewWalletFromSigner(signer, clientL2, clientL1)
}

// NewRandomWallet creates an instance of Wallet with a randomly generated account.
// The clientL2 and clientL1 parameters are optional, and can be configured with Wallet.Connect
// and Wallet.ConnectL1, respectively.
func NewRandomWallet(chainId *big.Int, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	signer, err := NewRandomBaseSigner(chainId)
	if err != nil {
		return nil, err
	}
	return NewWalletFromSigner(signer, clientL2, clientL1)
}

// Nonce returns the account nonce of the associated account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (w *Wallet) Nonce(ctx context.Context, blockNumber *big.Int) (uint64, error) {
	return w.clientL2.NonceAt(ctx, w.Address(), blockNumber)
}

// PendingNonce returns the account nonce of the associated account in the pending state.
// This is the nonce that should be used for the next transaction.
func (w *Wallet) PendingNonce(ctx context.Context) (uint64, error) {
	return w.clientL2.PendingNonceAt(ctx, w.Address())
}

// Connect returns a new instance of Wallet with the provided clientL2 for the L2 network.
func (w *Wallet) Connect(client *clients.Client) (*Wallet, error) {
	return NewWalletFromSigner(w.Signer(), client, w.clientL1)
}

// ConnectL1 returns a new instance of Wallet with the provided clientL2 for the L1 network.
func (w *Wallet) ConnectL1(client *ethclient.Client) (*Wallet, error) {
	return NewWalletFromSigner(w.Signer(), w.clientL2, client)
}
