package zksync2

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/contracts/IL1Bridge"
	"math/big"
)

type EthProvider interface {
	Deposit(token *Token, amount *big.Int, address common.Address, options *GasOptions) (*types.Transaction, error)
}

func NewDefaultEthProvider(rpcClient *rpc.Client, auth *bind.TransactOpts,
	l1EthBridge *IL1Bridge.IL1Bridge, l1ERC20Bridge *IL1Bridge.IL1Bridge) *DefaultEthProvider {
	return &DefaultEthProvider{
		rc:            rpcClient,
		ec:            ethclient.NewClient(rpcClient),
		auth:          auth,
		l1EthBridge:   l1EthBridge,
		l1ERC20Bridge: l1ERC20Bridge,
	}
}

type DefaultEthProvider struct {
	rc   *rpc.Client
	ec   *ethclient.Client
	auth *bind.TransactOpts

	l1EthBridge   *IL1Bridge.IL1Bridge
	l1ERC20Bridge *IL1Bridge.IL1Bridge
}

type GasOptions struct {
	GasPrice *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit uint64   // Gas limit to set for the transaction execution (0 = estimate)
}

func (p *DefaultEthProvider) Deposit(token *Token, amount *big.Int, address common.Address, options *GasOptions) (*types.Transaction, error) {
	auth := p.getAuth(options)
	if token.IsETH() {
		auth.Value = amount
		return p.l1EthBridge.Deposit(auth, address, EthAddress, amount)
	} else {
		auth.Value = nil
		return p.l1ERC20Bridge.Deposit(auth, address, token.L1Address, amount)
	}
}

// getAuth make a new copy of origin TransactOpts to be used safely for each call
func (p *DefaultEthProvider) getAuth(options *GasOptions) *bind.TransactOpts {
	newAuth := &bind.TransactOpts{
		From:   p.auth.From,
		Signer: p.auth.Signer,
	}
	if options != nil {
		newAuth.GasPrice = options.GasPrice
		newAuth.GasLimit = options.GasLimit
	}
	return newAuth
}
