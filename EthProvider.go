package zksync2

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/contracts/ERC20"
	"github.com/zksync-sdk/zksync2-go/contracts/IL1Bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/IZkSync"
	"math/big"
)

var (
	MaxApproveAmount             = big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	DefaultThreshold             = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(255), nil)
	RecommendedDepositL2GasLimit = big.NewInt(10000000)
	DepositGasPerPubdataLimit    = big.NewInt(800)
)

type EthProvider interface {
	GetClient() *ethclient.Client
	ApproveDeposit(token *Token, limit *big.Int, options *GasOptions) (*types.Transaction, error)
	IsDepositApproved(token *Token, to common.Address, threshold *big.Int) (bool, error)
	Deposit(token *Token, amount *big.Int, address common.Address, options *GasOptions) (*types.Transaction, error)
	RequestExecute(contractL2 common.Address, l2Value *big.Int, calldata []byte,
		l2GasLimit *big.Int, l2GasPerPubdataByteLimit *big.Int,
		factoryDeps [][]byte, refundRecipient common.Address, auth *bind.TransactOpts) (*types.Transaction, error)
	FinalizeEthWithdrawal(l2BlockNumber *big.Int, l2MessageIndex *big.Int, l2TxNumberInBlock *big.Int,
		message []byte, proof []common.Hash, options *GasOptions) (*types.Transaction, error)
	FinalizeWithdrawal(l1BridgeAddress common.Address, l2BlockNumber *big.Int, l2MessageIndex *big.Int,
		l2TxNumberInBlock *big.Int, message []byte, proof []common.Hash, options *GasOptions) (*types.Transaction, error)
	IsEthWithdrawalFinalized(l2BlockNumber *big.Int, l2MessageIndex *big.Int) (bool, error)
	IsWithdrawalFinalized(l1BridgeAddress common.Address, l2BlockNumber *big.Int, l2MessageIndex *big.Int) (bool, error)
	ClaimFailedDeposit(l1BridgeAddress common.Address, depositSender common.Address,
		l1Token common.Address, l2TxHash common.Hash, l2BlockNumber *big.Int, l2MessageIndex *big.Int,
		l2TxNumberInBlock *big.Int, proof []common.Hash, options *GasOptions) (*types.Transaction, error)
	GetL2HashFromPriorityOp(l1Receipt *types.Receipt) (common.Hash, error)
	GetBaseCost(l2GasLimit *big.Int, l2GasPerPubdataByteLimit *big.Int, gasPrice *big.Int) (*big.Int, error)
}

func NewDefaultEthProvider(rpcClient *rpc.Client, auth *bind.TransactOpts,
	mainContractAddress, l1ERC20BridgeAddress common.Address) (*DefaultEthProvider, error) {
	ec := ethclient.NewClient(rpcClient)
	iZkSync, err := IZkSync.NewIZkSync(mainContractAddress, ec)
	if err != nil {
		return nil, fmt.Errorf("failed to load IZkSync: %w", err)
	}
	l1ERC20Bridge, err := IL1Bridge.NewIL1Bridge(l1ERC20BridgeAddress, ec)
	if err != nil {
		return nil, fmt.Errorf("failed to load L1ERC20Bridge: %w", err)
	}
	return &DefaultEthProvider{
		rc:                   rpcClient,
		ec:                   ec,
		auth:                 auth,
		mainContractAddress:  mainContractAddress,
		l1ERC20BridgeAddress: l1ERC20BridgeAddress,
		l1ERC20Bridge:        l1ERC20Bridge,
		iZkSync:              iZkSync,
	}, nil
}

type DefaultEthProvider struct {
	rc   *rpc.Client
	ec   *ethclient.Client
	auth *bind.TransactOpts

	mainContractAddress  common.Address
	l1ERC20BridgeAddress common.Address
	l1ERC20Bridge        *IL1Bridge.IL1Bridge
	iZkSync              *IZkSync.IZkSync
}

func (p *DefaultEthProvider) GetClient() *ethclient.Client {
	return p.ec
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
	baseCost, err := p.GetBaseCost(RecommendedDepositL2GasLimit, DepositGasPerPubdataLimit, auth.GasPrice)
	if err != nil {
		return nil, fmt.Errorf("failed to GetBaseCost: %w", err)
	}
	if token.IsETH() {
		auth.Value = big.NewInt(0).Add(baseCost, amount)
		return p.RequestExecute(address, amount, nil,
			RecommendedDepositL2GasLimit, DepositGasPerPubdataLimit,
			nil, address, auth)
	} else {
		auth.Value = baseCost
		return p.l1ERC20Bridge.Deposit(auth,
			address, token.L1Address, amount,
			RecommendedDepositL2GasLimit, DepositGasPerPubdataLimit)
	}
}

func (p *DefaultEthProvider) RequestExecute(contractL2 common.Address, l2Value *big.Int, calldata []byte,
	l2GasLimit *big.Int, l2GasPerPubdataByteLimit *big.Int,
	factoryDeps [][]byte, refundRecipient common.Address, auth *bind.TransactOpts) (*types.Transaction, error) {
	return p.iZkSync.RequestL2Transaction(auth,
		contractL2,
		l2Value,
		calldata,
		l2GasLimit,
		l2GasPerPubdataByteLimit,
		factoryDeps,
		refundRecipient,
	)
}

func (p *DefaultEthProvider) FinalizeEthWithdrawal(l2BlockNumber *big.Int, l2MessageIndex *big.Int,
	l2TxNumberInBlock *big.Int, message []byte, proof []common.Hash, options *GasOptions) (*types.Transaction, error) {
	proof32 := make([][32]byte, len(proof))
	for i, pr := range proof {
		proof32[i] = pr
	}
	auth := p.getAuth(options)
	return p.iZkSync.FinalizeEthWithdrawal(auth,
		l2BlockNumber,
		l2MessageIndex,
		uint16(l2TxNumberInBlock.Uint64()),
		message,
		proof32,
	)
}

func (p *DefaultEthProvider) FinalizeWithdrawal(l1BridgeAddress common.Address, l2BlockNumber *big.Int, l2MessageIndex *big.Int,
	l2TxNumberInBlock *big.Int, message []byte, proof []common.Hash, options *GasOptions) (*types.Transaction, error) {
	l1Bridge, err := IL1Bridge.NewIL1Bridge(l1BridgeAddress, p.ec)
	if err != nil {
		return nil, fmt.Errorf("failed to init l1Bridge: %w", err)
	}
	proof32 := make([][32]byte, len(proof))
	for i, pr := range proof {
		proof32[i] = pr
	}
	auth := p.getAuth(options)
	return l1Bridge.FinalizeWithdrawal(auth,
		l2BlockNumber,
		l2MessageIndex,
		uint16(l2TxNumberInBlock.Uint64()),
		message,
		proof32,
	)
}

func (p *DefaultEthProvider) ClaimFailedDeposit(l1BridgeAddress common.Address, depositSender common.Address,
	l1Token common.Address, l2TxHash common.Hash, l2BlockNumber *big.Int, l2MessageIndex *big.Int,
	l2TxNumberInBlock *big.Int, proof []common.Hash, options *GasOptions) (*types.Transaction, error) {
	// _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16,
	//_merkleProof [][32]byte
	l1Bridge, err := IL1Bridge.NewIL1Bridge(l1BridgeAddress, p.ec)
	if err != nil {
		return nil, fmt.Errorf("failed to init l1Bridge: %w", err)
	}
	proof32 := make([][32]byte, len(proof))
	for i, pr := range proof {
		proof32[i] = pr
	}
	auth := p.getAuth(options)
	return l1Bridge.ClaimFailedDeposit(auth,
		depositSender,
		l1Token,
		l2TxHash,
		l2BlockNumber,
		l2MessageIndex,
		uint16(l2TxNumberInBlock.Uint64()),
		proof32,
	)
}

func (p *DefaultEthProvider) IsEthWithdrawalFinalized(l2BlockNumber *big.Int, l2MessageIndex *big.Int) (bool, error) {
	return p.iZkSync.IsEthWithdrawalFinalized(&bind.CallOpts{},
		l2BlockNumber,
		l2MessageIndex,
	)
}

func (p *DefaultEthProvider) IsWithdrawalFinalized(l1BridgeAddress common.Address, l2BlockNumber *big.Int,
	l2MessageIndex *big.Int) (bool, error) {
	l1Bridge, err := IL1Bridge.NewIL1Bridge(l1BridgeAddress, p.ec)
	if err != nil {
		return false, fmt.Errorf("failed to init l1Bridge: %w", err)
	}
	return l1Bridge.IsWithdrawalFinalized(&bind.CallOpts{},
		l2BlockNumber,
		l2MessageIndex,
	)
}

func (p *DefaultEthProvider) GetL2HashFromPriorityOp(l1Receipt *types.Receipt) (common.Hash, error) {
	for _, l := range l1Receipt.Logs {
		if l.Address == p.mainContractAddress {
			req, err := p.iZkSync.ParseNewPriorityRequest(*l)
			if err != nil {
				return common.Hash{}, fmt.Errorf("failed to ParseNewPriorityRequest: %w", err)
			}
			return req.TxHash, nil
		}
	}
	return common.Hash{}, errors.New("wrong tx")
}

func (p *DefaultEthProvider) GetBaseCost(l2GasLimit *big.Int, l2GasPerPubdataByteLimit *big.Int, gasPrice *big.Int) (*big.Int, error) {
	if gasPrice == nil {
		var err error
		if gasPrice, err = p.ec.SuggestGasPrice(context.Background()); err != nil {
			return nil, err
		}
	}
	return p.iZkSync.L2TransactionBaseCost(&bind.CallOpts{},
		gasPrice,
		l2GasLimit,
		l2GasPerPubdataByteLimit,
	)
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
