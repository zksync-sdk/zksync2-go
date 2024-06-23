package clients

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"math/big"
)

// EthereumClient provides Ethereum RPC methods on  ZKsync Era node, ones that has `eth_` prefix.
// Interface contains same methods as ethclient.Client except for FeeHistory method.
// Additionally, it has extra methods capable of working with EIP-712 transactions.
// It is designed to be compatible with bind.ContractBackend interface, enabling support for
// smart contracts generated using the abigen tool.
// Some types of function parameters and return values are different from original geth library
// because are not compatible with ZKsync Era node and are adjusted accordingly.
type EthereumClient interface {
	// Client gets the underlying RPC client.
	Client() *rpc.Client
	// Close closes the underlying RPC connection.
	Close()

	// ChainID retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockByHash returns the given full block.
	//
	// Note that loading full blocks requires two requests. Use HeaderByHash
	// if you don't need all transactions or uncle headers.
	BlockByHash(ctx context.Context, hash common.Hash) (*zkTypes.Block, error)
	// BlockByNumber returns a block from the current canonical chain. If number is nil, the
	// latest known block is returned.
	//
	// Note that loading full blocks requires two requests. Use HeaderByNumber
	// if you don't need all transactions or uncle headers.
	BlockByNumber(ctx context.Context, number *big.Int) (*zkTypes.Block, error)
	// BlockNumber returns the most recent block number
	BlockNumber(ctx context.Context) (uint64, error)
	// PeerCount returns the number of p2p peers as reported by the net_peerCount method
	PeerCount(ctx context.Context) (uint64, error)
	// HeaderByHash returns the block header with the given hash.
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	// HeaderByNumber returns a block header from the current canonical chain. If number is
	// nil, the latest known header is returned.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// TransactionByHash returns the transaction with the given hash.
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *zkTypes.TransactionResponse, isPending bool, err error)
	// TransactionSender returns the sender address of the given transaction. The transaction
	// must be known to the remote node and included in the blockchain at the given block and
	// index. The sender is the one derived by the protocol at the time of inclusion.
	TransactionSender(ctx context.Context, tx *zkTypes.TransactionResponse, block common.Hash, index uint) (common.Address, error)
	// TransactionCount returns the total number of transactions in the given block.
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	// TransactionInBlock returns a single transaction at index in the given block.
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*zkTypes.TransactionResponse, error)
	// TransactionReceipt returns the receipt of a transaction by transaction hash.
	// Note that the receipt is not available for pending transactions.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*zkTypes.Receipt, error)
	// SyncProgress retrieves the current progress of the sync algorithm. If there's
	// no sync currently running, it returns nil.
	SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error)
	// SubscribeNewHead subscribes to notifications about the current blockchain head
	// on the given channel.
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)

	// NetworkID returns the network ID for this client.
	NetworkID(ctx context.Context) (*big.Int, error)
	// BalanceAt returns the wei balance of the given account.
	// The block number can be nil, in which case the balance is taken from the latest known block.
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	// StorageAt returns the value of key in the contract storage of the given account.
	// The block number can be nil, in which case the value is taken from the latest known block.
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	// CodeAt returns the contract code of the given account.
	// The block number can be nil, in which case the code is taken from the latest known block.
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	// NonceAt returns the account nonce of the given account.
	// The block number can be nil, in which case the nonce is taken from the latest known block.
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)

	// FilterLogs performs the same function as FilterLogsL2, and that method should be used instead.
	// This method is designed to be compatible with bind.ContractBackend.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	// FilterLogsL2 executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	FilterLogsL2(ctx context.Context, query ethereum.FilterQuery) ([]zkTypes.Log, error)
	// SubscribeFilterLogs performs the same function as SubscribeFilterLogsL2, and that method should be used instead.
	// This method is designed to be compatible with bind.ContractBackend.
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	// SubscribeFilterLogsL2 creates a background log filtering operation, returning
	// a subscription immediately, which can be used to stream the found events.
	SubscribeFilterLogsL2(ctx context.Context, query ethereum.FilterQuery, ch chan<- zkTypes.Log) (ethereum.Subscription, error)

	// PendingBalanceAt returns the wei balance of the given account in the pending state.
	PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)
	// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
	PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error)
	// PendingCodeAt returns the contract code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// PendingTransactionCount returns the total number of transactions in the pending state.
	PendingTransactionCount(ctx context.Context) (uint, error)

	// CallContract executes a message call transaction, which is directly executed in the VM
	// of the node, but never mined into the blockchain.
	//
	// blockNumber selects the block height at which the call runs. It can be nil, in which
	// case the code is taken from the latest known block. Note that state from very old
	// blocks might not be available.
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	// CallContractL2 is almost the same as CallContract except that it executes a message call
	// for EIP-712 transaction.
	CallContractL2(ctx context.Context, msg zkTypes.CallMsg, blockNumber *big.Int) ([]byte, error)
	// CallContractAtHash is almost the same as CallContract except that it selects
	// the block by block hash instead of block height.
	CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error)
	// CallContractAtHashL2 is almost the same as CallContractL2 except that it selects
	// the block by block hash instead of block height.
	CallContractAtHashL2(ctx context.Context, msg zkTypes.CallMsg, blockHash common.Hash) ([]byte, error)
	// PendingCallContract executes a message call transaction using the EVM.
	// The state seen by the contract call is the pending state.
	PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error)
	// PendingCallContractL2 executes a message call for EIP-712 transaction using the EVM.
	// The state seen by the contract call is the pending state.
	PendingCallContractL2(ctx context.Context, msg zkTypes.CallMsg) ([]byte, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
	// allow a timely execution of a transaction.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// EstimateGas tries to estimate the gas needed to execute a transaction based on
	// the current pending state of the backend blockchain. There is no guarantee that this is
	// the true gas limit requirement as other transactions may be added or removed by miners,
	// but it should provide a basis for setting a reasonable default.
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	// EstimateGasL2 is almost the same as EstimateGas except that it executes an EIP-712 transaction.
	EstimateGasL2(ctx context.Context, msg zkTypes.CallMsg) (uint64, error)
	// SendTransaction injects a signed transaction into the pending pool for execution.
	//
	// If the transaction was a contract creation use the TransactionReceipt method to get the
	// contract address after the transaction has been mined.
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	// SendRawTransaction injects a signed raw transaction into the pending pool for execution.
	SendRawTransaction(ctx context.Context, tx []byte) (common.Hash, error)

	// WaitMined waits for tx to be mined on the blockchain.
	// It stops waiting when the context is canceled.
	WaitMined(ctx context.Context, txHash common.Hash) (*zkTypes.Receipt, error)
	// WaitFinalized waits for tx to be finalized on the blockchain.
	// It stops waiting when the context is canceled.
	WaitFinalized(ctx context.Context, txHash common.Hash) (*zkTypes.Receipt, error)
}

// ZkSyncEraClient provides the API to ZKsync Era features and
// specific RPC methods, ones that that has `zks_` prefix.
type ZkSyncEraClient interface {
	// MainContractAddress returns the address of the ZKsync Era contract.
	MainContractAddress(ctx context.Context) (common.Address, error)
	// TestnetPaymaster returns the testnet paymaster address if available, or nil.
	TestnetPaymaster(ctx context.Context) (common.Address, error)
	// BridgeContracts returns the addresses of the default ZKsync Era bridge
	// contracts on both L1 and L2.
	BridgeContracts(ctx context.Context) (*zkTypes.BridgeContracts, error)
	// ContractAccountInfo returns the version of the supported account abstraction
	// and nonce ordering from a given contract address.
	ContractAccountInfo(ctx context.Context, address common.Address) (*zkTypes.ContractAccountInfo, error)

	// L1ChainID returns the chain id of the underlying L1.
	L1ChainID(ctx context.Context) (*big.Int, error)
	// L1BatchNumber returns the latest L1 batch number.
	L1BatchNumber(ctx context.Context) (*big.Int, error)
	// L1BatchBlockRange returns the range of blocks contained within a batch given
	// by batch number.
	L1BatchBlockRange(ctx context.Context, l1BatchNumber *big.Int) (*BlockRange, error)
	// L1BatchDetails returns data pertaining to a given batch.
	L1BatchDetails(ctx context.Context, l1BatchNumber *big.Int) (*zkTypes.BatchDetails, error)
	// BlockDetails returns additional ZKsync Era-specific information about the L2
	// block.
	BlockDetails(ctx context.Context, block uint32) (*zkTypes.BlockDetails, error)
	// TransactionDetails returns data from a specific transaction given by the
	// transaction hash.
	TransactionDetails(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionDetails, error)
	// LogProof returns the proof for a transaction's L2 to L1 log sent via the
	// L1Messenger system contract.
	LogProof(ctx context.Context, txHash common.Hash, logIndex int) (*zkTypes.MessageProof, error)
	// Deprecated: Deprecated in favor of LogProof.
	MsgProof(ctx context.Context, block uint32, sender common.Address, msg common.Hash) (*zkTypes.MessageProof, error)
	// L2TransactionFromPriorityOp returns transaction on L2 network from transaction
	// receipt on L1 network.
	L2TransactionFromPriorityOp(ctx context.Context, l1TxReceipt *types.Receipt) (*zkTypes.TransactionResponse, error)

	// Deprecated: Method is deprecated and will be removed in the near future.
	ConfirmedTokens(ctx context.Context, from uint32, limit uint8) ([]*zkTypes.Token, error)
	// Deprecated: Method is deprecated and will be removed in the near future.
	TokenPrice(ctx context.Context, address common.Address) (*big.Float, error)
	// L2TokenAddress returns the L2 token address equivalent for a L1 token address
	// as they are not equal. ETH address is set to zero address.
	L2TokenAddress(ctx context.Context, token common.Address) (common.Address, error)
	// L1TokenAddress returns the L1 token address equivalent for a L2 token address
	// as they are not equal. ETH address is set to zero address.
	L1TokenAddress(ctx context.Context, token common.Address) (common.Address, error)
	// AllAccountBalances returns all balances for confirmed tokens given by an
	// account address.
	AllAccountBalances(ctx context.Context, address common.Address) (map[common.Address]*big.Int, error)

	// EstimateFee Returns the fee for the transaction.
	EstimateFee(ctx context.Context, tx zkTypes.CallMsg) (*zkTypes.Fee, error)
	// EstimateGasL1 estimates the amount of gas required to submit a transaction
	// from L1 to L2.
	EstimateGasL1(ctx context.Context, tx zkTypes.CallMsg) (uint64, error)
	// EstimateGasTransfer estimates the amount of gas required for a transfer
	// transaction.
	EstimateGasTransfer(ctx context.Context, msg TransferCallMsg) (uint64, error)
	// EstimateGasWithdraw estimates the amount of gas required for a withdrawal
	// transaction.
	EstimateGasWithdraw(ctx context.Context, msg WithdrawalCallMsg) (uint64, error)
	// EstimateL1ToL2Execute estimates the amount of gas required for an L1 to L2
	// execute operation.
	EstimateL1ToL2Execute(ctx context.Context, msg zkTypes.CallMsg) (uint64, error)
}

// Client defines both ethereum and ZKsync Era RPC methods and common features.
type Client interface {
	EthereumClient
	ZkSyncEraClient
}
