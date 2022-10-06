package zksync2

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/contracts/ERC20"
	"github.com/zksync-sdk/zksync2-go/contracts/IL1Bridge"
	"math/big"
)

var (
	MaxApproveAmount = big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	DefaultThreshold = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(255), nil)
)

type EthProvider interface {
	ApproveDeposit(token *Token, limit *big.Int, options *GasOptions) (*types.Transaction, error)
	IsDepositApproved(token *Token, to common.Address, threshold *big.Int) (bool, error)
	Deposit(token *Token, amount *big.Int, address common.Address, options *GasOptions) (*types.Transaction, error)
}

func NewDefaultEthProvider(rpcClient *rpc.Client, auth *bind.TransactOpts,
	l1EthBridgeAddress, l1ERC20BridgeAddress common.Address) (*DefaultEthProvider, error) {
	ec := ethclient.NewClient(rpcClient)
	l1EthBridge, err := IL1Bridge.NewIL1Bridge(l1EthBridgeAddress, ec)
	if err != nil {
		return nil, fmt.Errorf("failed to load L1EthBridge: %w", err)
	}
	l1ERC20Bridge, err := IL1Bridge.NewIL1Bridge(l1ERC20BridgeAddress, ec)
	if err != nil {
		return nil, fmt.Errorf("failed to load L1ERC20Bridge: %w", err)
	}
	return &DefaultEthProvider{
		rc:                   rpcClient,
		ec:                   ec,
		auth:                 auth,
		l1EthBridgeAddress:   l1EthBridgeAddress,
		l1ERC20BridgeAddress: l1ERC20BridgeAddress,
		l1EthBridge:          l1EthBridge,
		l1ERC20Bridge:        l1ERC20Bridge,
	}, nil
}

type DefaultEthProvider struct {
	rc   *rpc.Client
	ec   *ethclient.Client
	auth *bind.TransactOpts

	l1EthBridgeAddress   common.Address
	l1ERC20BridgeAddress common.Address
	l1EthBridge          *IL1Bridge.IL1Bridge
	l1ERC20Bridge        *IL1Bridge.IL1Bridge
}

type GasOptions struct {
	GasPrice *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit uint64   // Gas limit to set for the transaction execution (0 = estimate)
}

func (p *DefaultEthProvider) ApproveDeposit(token *Token, limit *big.Int, options *GasOptions) (*types.Transaction, error) {
	if token == nil {
		token = CreateETH()
	}
	auth := p.getAuth(options)
	erc20, err := ERC20.NewERC20(token.L1Address, p.ec)
	if err != nil {
		return nil, fmt.Errorf("failed to load ERC20: %w", err)
	}
	if limit == nil || len(limit.Bits()) == 0 {
		limit = MaxApproveAmount
	}
	return erc20.Approve(auth, p.l1ERC20BridgeAddress, limit)
}

func (p *DefaultEthProvider) IsDepositApproved(token *Token, to common.Address, threshold *big.Int) (bool, error) {
	if token == nil {
		token = CreateETH()
	}
	erc20, err := ERC20.NewERC20(token.L1Address, p.ec)
	if err != nil {
		return false, fmt.Errorf("failed to load ERC20: %w", err)
	}
	if threshold == nil || len(threshold.Bits()) == 0 {
		threshold = DefaultThreshold
	}
	res, err := erc20.Allowance(&bind.CallOpts{}, to, p.l1ERC20BridgeAddress)
	if err != nil {
		return false, fmt.Errorf("failed to query Allowance: %w", err)
	}
	return res.Cmp(threshold) >= 0, nil
}

func (p *DefaultEthProvider) Deposit(token *Token, amount *big.Int, address common.Address, options *GasOptions) (*types.Transaction, error) {
	if token == nil {
		token = CreateETH()
	}
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
