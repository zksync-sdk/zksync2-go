package accounts

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/clients"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"math/big"
)

// Wallet wraps all operations that interact with an associated account.
// An account typically contains a private key, allowing it to sign various types of payloads.
// Wallet implements the Adapter interface.
type Wallet struct {
	WalletL1
	WalletL2
	Deployer

	clientL1 *ethclient.Client
	clientL2 *clients.Client
}

// NewWallet creates an instance of Wallet associated with the account provided by the rawPrivateKey.
// The clientL1 parameters is optional; if not provided, only method form AdapterL2 and Deployer can be used,
// as the rest of the functionalities require communication with the L1 network.
// A Wallet can be configured to communicate with L1 networks by using and Wallet.ConnectL1 method.
func NewWallet(rawPrivateKey []byte, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	chainID, err := (*clientL2).ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	signer, err := NewBaseSignerFromRawPrivateKey(rawPrivateKey, chainID.Int64())
	if err != nil {
		return nil, err
	}
	s := Signer(signer)
	return NewWalletFromSigner(&s, clientL2, clientL1)
}

// NewWalletFromSigner creates an instance of Wallet associated with the account provided by the signer.
// The clientL2 and clientL1 parameters are optional; if not provided, only AdapterL2.SignTransaction,
// AdapterL2.Address, AdapterL2.Signer methods can be used, as the rest of the functionalities
// require communication with the network.
// A wallet that contains only a signer can be configured to communicate with L2 and L1 networks by
// using Wallet.Connect and Wallet.ConnectL1, respectively.
func NewWalletFromSigner(signer *Signer, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	if signer == nil {
		return nil, errors.New("signer must be provided")
	}
	var (
		walletL1 *WalletL1
		walletL2 *WalletL2
		err      error
	)

	walletL1 = new(WalletL1)
	if clientL1 != nil {
		if walletL1, err = NewWalletL1FromSigner(signer, clientL1, clientL2); err != nil {
			return nil, err
		}
	}
	if walletL2, err = NewWalletL2FromSigner(signer, clientL2); err != nil {
		return nil, err
	}
	a := AdapterL2(walletL2)
	return &Wallet{
		WalletL1: *walletL1,
		WalletL2: *walletL2,
		Deployer: NewBaseDeployer(&a),
		clientL1: clientL1,
		clientL2: clientL2,
	}, nil
}

// NewWalletFromMnemonic creates a new instance of Wallet based on the provided mnemonic phrase.
// The clientL2 and clientL1 parameters are optional, and can be configured with Wallet.Connect
// and Wallet.ConnectL1, respectively.
func NewWalletFromMnemonic(mnemonic string, chainId int64, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	signer, err := NewBaseSignerFromMnemonic(mnemonic, chainId)
	if err != nil {
		return nil, err
	}
	s := Signer(signer)
	return NewWalletFromSigner(&s, clientL2, clientL1)
}

// NewWalletFromRawPrivateKey creates a new instance of Wallet based on the provided private key
// of the account and chain ID.
// The clientL2 and clientL1 parameters are optional, and can be configured with Wallet.Connect
// and Wallet.ConnectL1, respectively.
func NewWalletFromRawPrivateKey(rawPk []byte, chainId int64, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	signer, err := NewBaseSignerFromRawPrivateKey(rawPk, chainId)
	if err != nil {
		return nil, err
	}
	s := Signer(signer)
	return NewWalletFromSigner(&s, clientL2, clientL1)
}

// NewRandomWallet creates an instance of Wallet with a randomly generated account.
// The clientL2 and clientL1 parameters are optional, and can be configured with Wallet.Connect
// and Wallet.ConnectL1, respectively.
func NewRandomWallet(chainId int64, clientL2 *clients.Client, clientL1 *ethclient.Client) (*Wallet, error) {
	signer, err := NewRandomBaseSigner(chainId)
	if err != nil {
		return nil, err
	}
	s := Signer(signer)
	return NewWalletFromSigner(&s, clientL2, clientL1)
}

// Nonce returns the account nonce of the associated account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (w *Wallet) Nonce(ctx context.Context, blockNumber *big.Int) (uint64, error) {
	return (*w.clientL2).NonceAt(ctx, w.Address(), blockNumber)
}

// PendingNonce returns the account nonce of the associated account in the pending state.
// This is the nonce that should be used for the next transaction.
func (w *Wallet) PendingNonce(ctx context.Context) (uint64, error) {
	return (*w.clientL2).PendingNonceAt(ctx, w.Address())
}

// Connect returns a new instance of Wallet with the provided client for the L2 network.
func (w *Wallet) Connect(client *clients.Client) (*Wallet, error) {
	s := w.Signer()
	return NewWalletFromSigner(&s, client, w.clientL1)
}

// ConnectL1 returns a new instance of Wallet with the provided client for the L1 network.
func (w *Wallet) ConnectL1(client *ethclient.Client) (*Wallet, error) {
	s := w.Signer()
	return NewWalletFromSigner(&s, w.clientL2, client)
}

// Deprecated: Deprecated in favor of Wallet.Signer.
func (w *Wallet) GetEthSigner() Signer {
	return w.Signer()
}

// Deprecated: Deprecated in favor of Wallet.Address.
func (w *Wallet) GetAddress() common.Address {
	return w.Address()
}

// Deprecated: Will be removed in the future releases.
func (w *Wallet) GetProvider() clients.Client {
	return *w.clientL2
}

// Deprecated: Will be removed in future releases.
func (w *Wallet) CreateEthereumProvider(rpcClient *rpc.Client) (clients.EthProvider, error) {
	mainContractAddress, err := (*w.clientL2).MainContractAddress(context.Background())
	if err != nil {
		return nil, err
	}
	clientL1 := ethclient.NewClient(rpcClient)
	chainId, err := clientL1.ChainID(context.Background())
	signer := w.Signer()
	auth, err := newTransactorWithSigner(&signer, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to init TransactOpts: %w", err)
	}
	bridgeContracts, err := (*w.clientL2).BridgeContracts(context.Background())
	if err != nil {
		return nil, err
	}
	return clients.NewDefaultEthProvider(rpcClient, auth, mainContractAddress, bridgeContracts.L1Erc20DefaultBridge)
}

// Deprecated: Deprecated in favor of Wallet.Balance.
func (w *Wallet) GetBalance() (*big.Int, error) {
	return w.Balance(context.Background(), w.Address(), nil)
}

// Deprecated: Deprecated in favor of Wallet.Balance.
func (w *Wallet) GetTokenBalance(token *zkTypes.Token) (*big.Int, error) {
	return w.Balance(context.Background(), token.L2Address, nil)
}

// Deprecated: Deprecated in favor of Wallet.Balance.
func (w *Wallet) GetBalanceOf(token *zkTypes.Token, blockNumber *big.Int) (*big.Int, error) {
	return w.Balance(context.Background(), token.L2Address, blockNumber)
}

// Deprecated: Deprecated in favor of Wallet.Nonce.
func (w *Wallet) GetNonce() (*big.Int, error) {
	nonce, err := w.Nonce(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetUint64(nonce), nil
}

// Deprecated: Deprecated in favor of Wallet.Nonce.
func (w *Wallet) GetNonceAt(blockNumber *big.Int) (*big.Int, error) {
	nonce, err := w.Nonce(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetUint64(nonce), nil
}
