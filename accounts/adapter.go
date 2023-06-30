package accounts

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zksync-sdk/zksync2-go/contracts/zksync"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"math/big"
)

// AdapterL1 is associated with an account and provides common operations on the
// L1 network for the associated account.
type AdapterL1 interface {
	// MainContract returns the zkSync L1 smart contract.
	MainContract(ctx context.Context) (*zksync.IZkSync, error)
	// L1BridgeContracts returns L1 bridge contracts.
	L1BridgeContracts(ctx context.Context) (*zkTypes.L1BridgeContracts, error)
	// BalanceL1 returns the balance of the specified token on L1 that can be
	// either ETH or any ERC20 token.
	BalanceL1(opts *CallOpts, token common.Address) (*big.Int, error)
	// AllowanceL1 returns the amount of approved tokens for a specific L1 bridge.
	AllowanceL1(opts *CallOpts, token common.Address, bridgeAddress common.Address) (*big.Int, error)
	// L2TokenAddress returns the corresponding address on the L2 network for the token on the L1 network.
	L2TokenAddress(ctx context.Context, token common.Address) (common.Address, error)
	// ApproveERC20 approves the specified amount of tokens for the specified L1 bridge.
	ApproveERC20(auth *TransactOpts, token common.Address, amount *big.Int, bridgeAddress common.Address) (*types.Transaction, error)
	// BaseCost returns base cost for L2 transaction.
	BaseCost(opts *CallOpts, gasLimit, gasPerPubdataByte, gasPrice *big.Int) (*big.Int, error)
	// Deposit transfers the specified token from the associated account on the L1 network
	// to the target account on the L2 network. The token can be either ETH or any ERC20 token.
	// For ERC20 tokens, enough approved tokens must be associated with the specified L1 bridge
	// (default one or the one defined in DepositTransaction.BridgeAddress). In this case,
	// DepositTransaction.ApproveERC20 can be enabled to perform token approval.
	// If there are already enough approved tokens for the L1 bridge, token approval will be skipped.
	// To check the amount of approved tokens for a specific bridge, use the AdapterL1.AllowanceL1 method.
	Deposit(auth *TransactOpts, tx DepositTransaction) (*types.Transaction, error)
	// EstimateGasDeposit estimates the amount of gas required for a deposit transaction on L1 network.
	// Gas of approving ERC20 token is not included in estimation.
	EstimateGasDeposit(ctx context.Context, msg DepositCallMsg) (uint64, error)
	// FullRequiredDepositFee retrieves the full needed ETH fee for the deposit on both L1 and L2 networks.
	FullRequiredDepositFee(ctx context.Context, msg DepositCallMsg) (*FullDepositFee, error)
	// FinalizeWithdraw proves the inclusion of the L2 -> L1 withdrawal message.
	FinalizeWithdraw(auth *TransactOpts, withdrawalHash common.Hash, index int) (*types.Transaction, error)
	// IsWithdrawFinalized checks if the withdrawal finalized on L1 network.
	IsWithdrawFinalized(opts *CallOpts, withdrawalHash common.Hash, index int) (bool, error)
	// ClaimFailedDeposit withdraws funds from the initiated deposit, which failed when finalizing on L2.
	// If the deposit L2 transaction has failed, it sends an L1 transaction calling ClaimFailedDeposit method
	// of the L1 bridge, which results in returning L1 tokens back to the depositor, otherwise throws the error.
	ClaimFailedDeposit(auth *TransactOpts, depositHash common.Hash) (*types.Transaction, error)
	// RequestExecute request execution of L2 transaction from L1.
	RequestExecute(auth *TransactOpts, tx RequestExecuteTransaction) (*types.Transaction, error)
	// EstimateGasRequestExecute estimates the amount of gas required for a request execute transaction.
	EstimateGasRequestExecute(ctx context.Context, msg RequestExecuteCallMsg) (uint64, error)
}

// AdapterL2 is associated with an account and provides common operations on the
// L2 network for the associated account.
type AdapterL2 interface {
	// Address returns the address of the associated account.
	Address() common.Address
	// Signer returns the signer of the associated account.
	Signer() Signer
	// Balance returns the balance of the specified token that can be either ETH or any ERC20 token.
	// The block number can be nil, in which case the balance is taken from the latest known block.
	Balance(ctx context.Context, token common.Address, at *big.Int) (*big.Int, error)
	// AllBalances returns all balances for confirmed tokens given by an associated
	// account.
	AllBalances(ctx context.Context) (map[common.Address]*big.Int, error)
	// L2BridgeContracts returns L2 bridge contracts.
	L2BridgeContracts(ctx context.Context) (*zkTypes.L2BridgeContracts, error)
	// Withdraw initiates the withdrawal process which withdraws ETH or any ERC20
	// token from the associated account on L2 network to the target account on L1
	// network.
	Withdraw(auth *TransactOpts, tx WithdrawalTransaction) (*types.Transaction, error)
	// EstimateGasWithdraw estimates the amount of gas required for a withdrawal
	// transaction.
	EstimateGasWithdraw(ctx context.Context, msg WithdrawalCallMsg) (uint64, error)
	// Transfer moves the ETH or any ERC20 token from the associated account to the
	// target account.
	Transfer(auth *TransactOpts, tx TransferTransaction) (*types.Transaction, error)
	// EstimateGasTransfer estimates the amount of gas required for a transfer
	// transaction.
	EstimateGasTransfer(ctx context.Context, msg TransferCallMsg) (uint64, error)
	// CallContract executes a message call for EIP-712 transaction, which is
	// directly executed in the VM of the node, but never mined into the blockchain.
	//
	// blockNumber selects the block height at which the call runs. It can be nil, in
	// which case the code is taken from the latest known block. Note that state from
	// very old blocks might not be available.
	CallContract(ctx context.Context, msg CallMsg, blockNumber *big.Int) ([]byte, error)
	// PopulateTransaction is designed for users who prefer a simplified approach by
	// providing only the necessary data to create a valid transaction. The only
	// required fields are Transaction.To and either Transaction.Data or
	// Transaction.Value (or both, if the method is payable). Any other fields that
	// are not set will be prepared by this method.
	PopulateTransaction(ctx context.Context, tx Transaction) (*zkTypes.Transaction712, error)
	// SignTransaction returns a signed transaction that is ready to be broadcast to
	// the network. The input transaction must be a valid transaction with all fields
	// having appropriate values. To obtain a valid transaction, you can use the
	// PopulateTransaction method.
	SignTransaction(tx *zkTypes.Transaction712) ([]byte, error)
	// SendTransaction injects a transaction into the pending pool for execution. Any
	// unset transaction fields are prepared using the PopulateTransaction method.
	SendTransaction(ctx context.Context, tx *Transaction) (common.Hash, error)
}

// Deployer is associated with an account and provides deployment of smart contracts
// and smart accounts on L2 network for the associated account.
type Deployer interface {
	// Deploy deploys smart contract using CREATE2 opcode.
	Deploy(auth *TransactOpts, tx Create2Transaction) (common.Hash, error)
	// DeployWithCreate deploys smart contract using CREATE opcode.
	DeployWithCreate(auth *TransactOpts, tx CreateTransaction) (common.Hash, error)
	// DeployAccount deploys smart account using CREATE2 opcode.
	DeployAccount(auth *TransactOpts, tx Create2Transaction) (common.Hash, error)
	// DeployAccountWithCreate deploys smart account using CREATE opcode.
	DeployAccountWithCreate(auth *TransactOpts, tx CreateTransaction) (common.Hash, error)
}

// Adapter is associated with an account and provides common functionalities on
// both L1 and L2 network as well as deployment of smart contracts and smart
// accounts for the associated account.
type Adapter interface {
	AdapterL1
	AdapterL2
	Deployer
}
