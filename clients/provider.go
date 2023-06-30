package clients

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"time"
)

// Deprecated: Deprecated in favor of EthereumClient.
type EthereumProvider interface {
	// Close closes the underlying RPC connection.
	Close()

	// ChainID retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockByHash returns the given full block.
	//
	// Note that loading full blocks requires two requests. Use HeaderByHash
	// if you don't need all transactions or uncle headers.
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	// BlockByNumber returns a block from the current canonical chain. If number is nil, the
	// latest known block is returned.
	//
	// Note that loading full blocks requires two requests. Use HeaderByNumber
	// if you don't need all transactions or uncle headers.
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
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
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	// TransactionSender returns the sender address of the given transaction. The transaction
	// must be known to the remote node and included in the blockchain at the given block and
	// index. The sender is the one derived by the protocol at the time of inclusion.
	//
	// There is a fast-path for transactions retrieved by TransactionByHash and
	// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
	TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error)
	// TransactionCount returns the total number of transactions in the given block.
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	// TransactionInBlock returns a single transaction at index in the given block.
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)
	// TransactionReceipt returns the receipt of a transaction by transaction hash.
	// Note that the receipt is not available for pending transactions.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
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

	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	// SubscribeFilterLogs creates a background log filtering operation, returning
	// a subscription immediately, which can be used to stream the found events.
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)

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
	// CallContractAtHash is almost the same as CallContract except that it selects
	// the block by block hash instead of block height.
	CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error)
	// PendingCallContract executes a message call transaction using the EVM.
	// The state seen by the contract call is the pending state.
	PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
	// allow a timely execution of a transaction.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// FeeHistory retrieves the fee market history.
	FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error)
	// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
	// the current pending state of the backend blockchain. There is no guarantee that this is
	// the true gas limit requirement as other transactions may be added or removed by miners,
	// but it should provide a basis for setting a reasonable default.
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	// SendTransaction injects a signed transaction into the pending pool for execution.
	//
	// If the transaction was a contract creation use the TransactionReceipt method to get the
	// contract address after the transaction has been mined.
	SendTransaction(ctx context.Context, tx *types.Transaction) error

	GetClient() *ethclient.Client
	GetBalance(address common.Address, blockNumber zkTypes.BlockNumber) (*big.Int, error)
	GetBlockByNumber(blockNumber zkTypes.BlockNumber) (*zkTypes.Block, error)
	GetBlockByHash(blockHash common.Hash) (*zkTypes.Block, error)
	GetTransactionCount(address common.Address, blockNumber zkTypes.BlockNumber) (*big.Int, error)
	GetTransactionReceipt(txHash common.Hash) (*zkTypes.TransactionReceipt, error)
	GetTransaction(txHash common.Hash) (*zkTypes.TransactionResponse, error)
	WaitMined(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionReceipt, error)
	WaitFinalized(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionReceipt, error)
	GetGasPrice() (*big.Int, error)
	SendRawTransaction(tx []byte) (common.Hash, error)
	GetLogs(q zkTypes.FilterQuery) ([]zkTypes.Log, error)
}

// Deprecated: Deprecated in favor of ZkSyncEraClient.
type ZksyncProvider interface {
	ZksGetMainContract() (common.Address, error)
	ZksL1ChainId() (*big.Int, error)
	ZksL1BatchNumber() (*big.Int, error)
	ZksGetConfirmedTokens(from uint32, limit uint8) ([]*zkTypes.Token, error)
	ZksIsTokenLiquid(address common.Address) (bool, error)
	ZksGetTokenPrice(address common.Address) (*big.Float, error)
	L2TokenAddress(token common.Address) (common.Address, error)
	L1TokenAddress(token common.Address) (common.Address, error)
	ZksGetL2ToL1LogProof(txHash common.Hash, logIndex int) (*zkTypes.MessageProof, error)
	ZksGetL2ToL1MsgProof(block uint32, sender common.Address, msg common.Hash) (*zkTypes.MessageProof, error)
	ZksGetAllAccountBalances(address common.Address) (map[common.Address]*big.Int, error)
	ZksGetBridgeContracts() (*zkTypes.BridgeContracts, error)
	ZksEstimateFee(tx *zkTypes.Transaction) (*zkTypes.Fee, error)
	ZksGetTestnetPaymaster() (common.Address, error)
	ZksGetBlockDetails(block uint32) (*zkTypes.BlockDetails, error)
	EstimateGas712(tx *zkTypes.Transaction) (*big.Int, error)
	EstimateGasL1(tx *zkTypes.Transaction) (*big.Int, error)
	EstimateL1ToL2Execute(tx *zkTypes.Transaction) (*big.Int, error)
	GetL1BatchBlockRange(l1BatchNumber *big.Int) (*BlockRange, error)
	GetL1BatchDetails(l1BatchNumber *big.Int) (*zkTypes.BatchDetails, error)
	GetTransactionDetails(txHash common.Hash) (*zkTypes.TransactionDetails, error)
}

// Deprecated: Interface is deprecated in favor of Client.
type Provider interface {
	EthereumProvider
	ZksyncProvider
}

// Deprecated: Interface is deprecated in favor of NewClient.
func NewDefaultProvider(rawUrl string) (*DefaultProvider, error) {
	rpcClient, err := rpc.Dial(rawUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to rpc.Dial(): %w", err)
	}
	return &DefaultProvider{
		c:      rpcClient,
		Client: ethclient.NewClient(rpcClient),
	}, nil
}

// Deprecated: Interface is deprecated in favor of BaseClient.
type DefaultProvider struct {
	c *rpc.Client
	// also inherit default Ethereum clients
	*ethclient.Client
}

// Deprecated: Will be removed in the future releases.
func (p *DefaultProvider) GetClient() *ethclient.Client {
	return p.Client
}

// Deprecated: Deprecated in favor of BaseClient.BalanceAt.
func (p *DefaultProvider) GetBalance(address common.Address, blockNumber zkTypes.BlockNumber) (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "eth_getBalance", address, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_getBalance: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.BlockByNumber.
func (p *DefaultProvider) GetBlockByNumber(blockNumber zkTypes.BlockNumber) (*zkTypes.Block, error) {
	type TmpBlock struct {
		Number           hexutil.Big  `json:"number"`
		L1BatchNumber    *hexutil.Big `json:"l1BatchNumber"`
		L1BatchTimestamp *hexutil.Big `json:"l1BatchTimestamp"`
	}
	var resp *TmpBlock
	err := p.c.Call(&resp, "eth_getBlockByNumber", blockNumber, false)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_getBlockByNumber: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	//return resp, nil
	//ethBlock, err := p.Client.BlockByNumber(context.Background(), resp.Number.ToInt())
	if err != nil {
		return nil, err
	}
	return &zkTypes.Block{
		//L1BatchNumber:    resp.L1BatchNumber,
		//L1BatchTimestamp: resp.L1BatchTimestamp,
	}, nil
}

// Deprecated: Deprecated in favor of BaseClient.BlockByHash.
func (p *DefaultProvider) GetBlockByHash(blockHash common.Hash) (*zkTypes.Block, error) {
	type TmpBlock struct {
		L1BatchNumber    *hexutil.Big `json:"l1BatchNumber"`
		L1BatchTimestamp *hexutil.Big `json:"l1BatchTimestamp"`
	}
	var resp *TmpBlock
	err := p.c.Call(&resp, "eth_getBlockByHash", blockHash, false)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_getBlockByHash: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	//ethBlock, err := p.Client.BlockByHash(context.Background(), blockHash)
	if err != nil {
		return nil, err
	}
	return &zkTypes.Block{
		//Block:            *ethBlock,
		//L1BatchNumber:    resp.L1BatchNumber,
		//L1BatchTimestamp: resp.L1BatchTimestamp,
	}, nil
}

// Deprecated: Deprecated in favor of BaseClient.TransactionCount.
func (p *DefaultProvider) GetTransactionCount(address common.Address, blockNumber zkTypes.BlockNumber) (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "eth_getTransactionCount", address, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_getTransactionCount: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.TransactionReceipt.
func (p *DefaultProvider) GetTransactionReceipt(txHash common.Hash) (*zkTypes.TransactionReceipt, error) {
	var resp *zkTypes.TransactionReceipt
	err := p.c.Call(&resp, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_getTransactionReceipt: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.TransactionByHash.
func (p *DefaultProvider) GetTransaction(txHash common.Hash) (*zkTypes.TransactionResponse, error) {
	var resp *zkTypes.TransactionResponse
	err := p.c.Call(&resp, "eth_getTransactionByHash", txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_getTransactionByHash: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.EstimateGas.
func (p *DefaultProvider) EstimateGas(_ context.Context, call ethereum.CallMsg) (uint64, error) {
	gas, err := p.EstimateGas712(&zkTypes.Transaction{
		From:     call.From,
		To:       *call.To,
		Gas:      utils.NewBig(int64(call.Gas)),
		GasPrice: (*hexutil.Big)(call.GasPrice),
		Value:    (*hexutil.Big)(call.Value),
		Data:     call.Data,
		Eip712Meta: &zkTypes.Eip712Meta{
			GasPerPubdata: utils.NewBig(160_000),
		},
	})
	if err != nil {
		return 0, err
	}
	return gas.Uint64(), nil
}

// Deprecated: Deprecated in favor of BaseClient.EstimateGasL2.
func (p *DefaultProvider) EstimateGas712(tx *zkTypes.Transaction) (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "eth_estimateGas", tx, zkTypes.BlockNumberLatest)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_estimateGas: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.SuggestGasPrice.
func (p *DefaultProvider) GetGasPrice() (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "eth_gasPrice")
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_gasPrice: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.SendRawTransaction.
func (p *DefaultProvider) SendRawTransaction(tx []byte) (common.Hash, error) {
	var res string
	err := p.c.Call(&res, "eth_sendRawTransaction", hexutil.Encode(tx))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to call eth_sendRawTransaction: %w", err)
	}
	return common.HexToHash(res), nil
}

// Deprecated: Deprecated in favor of BaseClient.MainContractAddress.
func (p *DefaultProvider) ZksGetMainContract() (common.Address, error) {
	var res string
	err := p.c.Call(&res, "zks_getMainContract")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to query zks_getMainContract: %w", err)
	}
	return common.HexToAddress(res), nil
}

// Deprecated: Deprecated in favor of BaseClient.L1ChainID.
func (p *DefaultProvider) ZksL1ChainId() (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "zks_L1ChainId")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_L1ChainId: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.L1BatchNumber.
func (p *DefaultProvider) ZksL1BatchNumber() (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "zks_L1BatchNumber")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_L1BatchNumber: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.ConfirmedTokens.
func (p *DefaultProvider) ZksGetConfirmedTokens(from uint32, limit uint8) ([]*zkTypes.Token, error) {
	res := make([]*zkTypes.Token, 0)
	err := p.c.Call(&res, "zks_getConfirmedTokens", from, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getConfirmedTokens: %w", err)
	}
	return res, nil
}

// Deprecated: Deprecated in favor of BaseClient.IsTokenLiquid.
func (p *DefaultProvider) ZksIsTokenLiquid(address common.Address) (bool, error) {
	var res bool
	err := p.c.Call(&res, "zks_isTokenLiquid", address)
	if err != nil {
		return false, fmt.Errorf("failed to query zks_isTokenLiquid: %w", err)
	}
	return res, nil
}

// Deprecated: Deprecated in favor of BaseClient.TokenPrice.
func (p *DefaultProvider) ZksGetTokenPrice(address common.Address) (*big.Float, error) {
	var res string
	err := p.c.Call(&res, "zks_getTokenPrice", address)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getTokenPrice: %w", err)
	}
	resp, ok := big.NewFloat(0).SetString(res)
	if !ok {
		return nil, errors.New("failed to decode response as big.Float")
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.LogProof.
func (p *DefaultProvider) ZksGetL2ToL1LogProof(txHash common.Hash, logIndex int) (*zkTypes.MessageProof, error) {
	var resp *zkTypes.MessageProof
	err := p.c.Call(&resp, "zks_getL2ToL1LogProof", txHash, logIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL2ToL1LogProof: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of ZksGetL2ToL1LogProof.
func (p *DefaultProvider) ZksGetL2ToL1MsgProof(block uint32, sender common.Address, msg common.Hash) (*zkTypes.MessageProof, error) {
	var resp *zkTypes.MessageProof
	err := p.c.Call(&resp, "zks_getL2ToL1MsgProof", block, sender, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL2ToL1MsgProof: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.AllAccountBalances.
func (p *DefaultProvider) ZksGetAllAccountBalances(address common.Address) (map[common.Address]*big.Int, error) {
	res := make(map[common.Address]string)
	err := p.c.Call(&res, "zks_getAllAccountBalances", address)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getAllAccountBalances: %w", err)
	}
	resp := make(map[common.Address]*big.Int, len(res))
	for t, b := range res {
		resp[t], err = hexutil.DecodeBig(b)
		if err != nil {
			return nil, fmt.Errorf("failed to decode one of balances as big.Int: %w", err)
		}
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.BridgeContracts.
func (p *DefaultProvider) ZksGetBridgeContracts() (*zkTypes.BridgeContracts, error) {
	res := zkTypes.BridgeContracts{}
	err := p.c.Call(&res, "zks_getBridgeContracts")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getBridgeContracts: %w", err)
	}
	return &res, nil
}

// Deprecated: Deprecated in favor of BaseClient.EstimateFee.
func (p *DefaultProvider) ZksEstimateFee(tx *zkTypes.Transaction) (*zkTypes.Fee, error) {
	var res zkTypes.Fee
	err := p.c.Call(&res, "zks_estimateFee", tx)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_estimateFee: %w", err)
	}
	return &res, nil
}

// Deprecated: Deprecated in favor of BaseClient.TestnetPaymaster.
func (p *DefaultProvider) ZksGetTestnetPaymaster() (common.Address, error) {
	var res string
	err := p.c.Call(&res, "zks_getTestnetPaymaster")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to query zks_estimateFee: %w", err)
	}
	return common.HexToAddress(res), nil
}

// Deprecated: Deprecated in favor of BaseClient.BlockDetails.
func (p *DefaultProvider) ZksGetBlockDetails(block uint32) (*zkTypes.BlockDetails, error) {
	var resp *zkTypes.BlockDetails
	err := p.c.Call(&resp, "zks_getBlockDetails", block)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getBlockDetails: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.WaitMined.
func (p *DefaultProvider) WaitMined(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionReceipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	for {
		receipt, err := p.GetTransactionReceipt(txHash)
		if err == nil && receipt.BlockNumber != nil {
			return receipt, nil
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

// Deprecated: Deprecated in favor of BaseClient.WaitFinalized.
func (p *DefaultProvider) WaitFinalized(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionReceipt, error) {
	receipt, err := p.WaitMined(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx is mined: %w", err)
	}
	if receipt.BlockNumber == nil {
		return nil, errors.New("empty tx block number")
	}
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	var blockHead *types.Header
	for {
		err = p.c.CallContext(ctx, &blockHead, "eth_getBlockByNumber", zkTypes.BlockNumberFinalized, false)
		if err == nil && blockHead == nil {
			err = ethereum.NotFound
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get finalized block: %w", err)
		}
		if blockHead.Number.Cmp(receipt.BlockNumber) >= 0 {
			return receipt, nil
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

// Deprecated: Deprecated in favor of BaseClient.FilterLogsL2.
func (p *DefaultProvider) GetLogs(q zkTypes.FilterQuery) ([]zkTypes.Log, error) {
	var result []zkTypes.Log
	arg, err := utils.ToFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = p.c.Call(&result, "eth_getLogs", arg)
	return result, err
}

// Deprecated: Deprecated in favor of BaseClient.L2TokenAddress.
func (p *DefaultProvider) L2TokenAddress(token common.Address) (common.Address, error) {
	if token == utils.EthAddress {
		return utils.EthAddress, nil
	} else {
		bridgeContracts, err := p.ZksGetBridgeContracts()
		if err != nil {
			return common.Address{}, err
		}
		bridge, err := l2bridge.NewIL2Bridge(bridgeContracts.L2Erc20DefaultBridge, p.Client)
		if err != nil {
			return common.Address{}, err
		}
		tokenAddress, err := bridge.L2TokenAddress(nil, token)
		if err != nil {
			return common.Address{}, err
		}
		return tokenAddress, nil
	}
}

// Deprecated: Deprecated in favor of BaseClient.L1TokenAddress.
func (p *DefaultProvider) L1TokenAddress(token common.Address) (common.Address, error) {
	if token == utils.EthAddress {
		return utils.EthAddress, nil
	} else {
		bridgeContracts, err := p.ZksGetBridgeContracts()
		if err != nil {
			return common.Address{}, err
		}
		bridge, err := l2bridge.NewIL2Bridge(bridgeContracts.L2Erc20DefaultBridge, p.Client)
		if err != nil {
			return common.Address{}, err
		}
		tokenAddress, err := bridge.L1TokenAddress(nil, token)
		if err != nil {
			return common.Address{}, err
		}
		return tokenAddress, nil
	}
}

// Deprecated: Deprecated in favor of BaseClient.EstimateGasL1.
func (p *DefaultProvider) EstimateGasL1(tx *zkTypes.Transaction) (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "zks_estimateGasL1ToL2", tx)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_estimateGasL1ToL2: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.EstimateL1ToL2Execute.
func (p *DefaultProvider) EstimateL1ToL2Execute(tx *zkTypes.Transaction) (*big.Int, error) {
	if tx.Eip712Meta.GasPerPubdata == nil {
		tx.Eip712Meta.GasPerPubdata = utils.NewBig(utils.RequiredL1ToL2GasPerPubdataLimit.Int64())
	}

	// If the `from` address is not provided, we use a random address, because
	// due to storage slot aggregation, the gas estimation will depend on the address
	// and so estimation for the zero address may be smaller than for the sender.
	if tx.From == (common.Address{}) {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, fmt.Errorf("failed to generate radnom private key: %w", err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("failed to convert public key to ECDSA")
		}
		tx.From = crypto.PubkeyToAddress(*publicKeyECDSA)
	}

	gas, err := p.EstimateGasL1(tx)
	if err != nil {
		return nil, err
	}

	return gas, nil
}

// Deprecated: Deprecated in favor of BaseClient.L1BatchBlockRange.
func (p *DefaultProvider) GetL1BatchBlockRange(l1BatchNumber *big.Int) (*BlockRange, error) {
	var resp *BlockRange
	err := p.c.Call(&resp, "zks_getL1BatchBlockRange", l1BatchNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL1BatchBlockRange: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.L1BatchDetails.
func (p *DefaultProvider) GetL1BatchDetails(l1BatchNumber *big.Int) (*zkTypes.BatchDetails, error) {
	var resp *zkTypes.BatchDetails
	err := p.c.Call(&resp, "zks_getL1BatchDetails", l1BatchNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL1BatchDetails: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.TransactionDetails.
func (p *DefaultProvider) GetTransactionDetails(txHash common.Hash) (*zkTypes.TransactionDetails, error) {
	var resp *zkTypes.TransactionDetails
	err := p.c.Call(&resp, "zks_getTransactionDetails", txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getTransactionDetails: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Deprecated in favor of BaseClient.SuggestGasTipCap.
func (p *DefaultProvider) SuggestGasTipCap(_ context.Context) (*big.Int, error) {
	return utils.MaxPriorityFeePerGas, nil
}
