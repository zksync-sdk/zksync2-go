package clients

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/zksync-sdk/zksync2-go/contracts/contractdeployer"
	"github.com/zksync-sdk/zksync2-go/contracts/erc1271"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2sharedbridge"
	"github.com/zksync-sdk/zksync2-go/contracts/zksync"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"time"
)

type BaseClient struct {
	rpcClient *rpc.Client
	ethClient *ethclient.Client

	mainContractAddress common.Address
	mainContract        *zksync.IZkSync

	testnetPaymasterAddress common.Address
	bridgehubAddress        common.Address
	baseTokenAddress        common.Address

	bridgeContracts *zkTypes.BridgeContracts
}

// Dial connects a client to the given URL.
func Dial(rawUrl string) (Client, error) {
	return DialContext(context.Background(), rawUrl)
}

// DialContext connects a client to the given URL with context.
func DialContext(ctx context.Context, rawUrl string) (Client, error) {
	c, err := rpc.DialContext(ctx, rawUrl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// DialBase connects a client to the given URL.
func DialBase(rawUrl string) (*BaseClient, error) {
	c, err := rpc.DialContext(context.Background(), rawUrl)
	if err != nil {
		return nil, err
	}
	return &BaseClient{
		rpcClient: c,
		ethClient: ethclient.NewClient(c),
	}, nil
}

// DialContextBase connects a client to the given URL with context.
func DialContextBase(ctx context.Context, rawUrl string) (*BaseClient, error) {
	c, err := rpc.DialContext(ctx, rawUrl)
	if err != nil {
		return nil, err
	}
	return &BaseClient{
		rpcClient: c,
		ethClient: ethclient.NewClient(c),
	}, nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) Client {
	return &BaseClient{
		rpcClient: c,
		ethClient: ethclient.NewClient(c),
	}
}

// Client gets the underlying RPC client.
func (c *BaseClient) Client() *rpc.Client {
	return c.rpcClient
}

// Close closes the underlying RPC connection.
func (c *BaseClient) Close() {
	c.ethClient.Close()
}

// ChainID retrieves the current chain ID for transaction replay protection.
func (c *BaseClient) ChainID(ctx context.Context) (*big.Int, error) {
	return c.ethClient.ChainID(ctx)
}

// BlockByHash returns the given full block.
//
// Note that loading full blocks requires two requests. Use HeaderByHash
// if you don't need all transactions or uncle headers.
func (c *BaseClient) BlockByHash(ctx context.Context, hash common.Hash) (*zkTypes.Block, error) {
	return c.getBlock(ctx, "eth_getBlockByHash", hash, true)
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
//
// Note that loading full blocks requires two requests. Use HeaderByNumber
// if you don't need all transactions or uncle headers.
func (c *BaseClient) BlockByNumber(ctx context.Context, number *big.Int) (*zkTypes.Block, error) {
	return c.getBlock(ctx, "eth_getBlockByNumber", toBlockNumArg(number), true)
}

// BlockNumber returns the most recent block number
func (c *BaseClient) BlockNumber(ctx context.Context) (uint64, error) {
	return c.ethClient.BlockNumber(ctx)
}

// PeerCount returns the number of p2p peers as reported by the net_peerCount method
func (c *BaseClient) PeerCount(ctx context.Context) (uint64, error) {
	return c.ethClient.PeerCount(ctx)
}

// HeaderByHash returns the block header with the given hash.
func (c *BaseClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return c.ethClient.HeaderByHash(ctx, hash)
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (c *BaseClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.ethClient.HeaderByNumber(ctx, number)
}

// TransactionByHash returns the transaction with the given hash.
func (c *BaseClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *zkTypes.TransactionResponse, isPending bool, err error) {
	var resp *zkTypes.TransactionResponse
	err = c.rpcClient.CallContext(ctx, &resp, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, false, fmt.Errorf("failed to query eth_getTransactionByHash: %w", err)
	} else if resp == nil {
		return nil, false, ethereum.NotFound
	} // else if resp.S == nil || resp.R == nil || resp.V == nil {
	//	return nil, false, errors.New("server returned transaction without signature")
	//}
	return resp, resp.BlockHash == nil, nil
}

// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
func (c *BaseClient) TransactionSender(ctx context.Context, tx *zkTypes.TransactionResponse, block common.Hash, index uint) (common.Address, error) {
	var meta struct {
		Hash *common.Hash
		From common.Address
	}
	if err := c.rpcClient.CallContext(ctx, &meta, "eth_getTransactionByBlockHashAndIndex", block, hexutil.Uint64(index)); err != nil {
		return common.Address{}, err
	}
	if meta.Hash == nil || *meta.Hash != tx.Hash {
		return common.Address{}, errors.New("wrong inclusion block/index")
	}
	return meta.From, nil
}

// TransactionCount returns the total number of transactions in the given block.
func (c *BaseClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return c.ethClient.TransactionCount(ctx, blockHash)
}

// TransactionInBlock returns a single transaction at index in the given block.
func (c *BaseClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*zkTypes.TransactionResponse, error) {
	var tx *zkTypes.TransactionResponse
	err := c.rpcClient.CallContext(ctx, &tx, "eth_getTransactionByBlockHashAndIndex", blockHash, hexutil.Uint64(index))
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, ethereum.NotFound
	} // else if tx.R == nil {
	//	return nil, errors.New("server returned transaction without signature")
	//}
	return tx, err
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (c *BaseClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*zkTypes.Receipt, error) {
	var resp *zkTypes.Receipt
	err := c.rpcClient.CallContext(ctx, &resp, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_getTransactionReceipt: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (c *BaseClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	return c.ethClient.SyncProgress(ctx)
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (c *BaseClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return c.ethClient.SubscribeNewHead(ctx, ch)
}

// NetworkID returns the network ID for this client.
func (c *BaseClient) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.ethClient.NetworkID(ctx)
}

// BalanceAt returns the base token balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (c *BaseClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return c.ethClient.BalanceAt(ctx, account, blockNumber)
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (c *BaseClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	return c.ethClient.StorageAt(ctx, account, key, blockNumber)
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (c *BaseClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return c.ethClient.CodeAt(ctx, account, blockNumber)
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (c *BaseClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return c.ethClient.NonceAt(ctx, account, blockNumber)
}

// FilterLogs performs the same function as FilterLogsL2, and that method should be used instead.
// This method is designed to be compatible with bind.ContractBackend.
func (c *BaseClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return c.ethClient.FilterLogs(ctx, query)
}

// FilterLogsL2 executes a log filter operation, blocking during execution and
// returning all the results in one batch.
func (c *BaseClient) FilterLogsL2(ctx context.Context, q ethereum.FilterQuery) ([]zkTypes.Log, error) {
	var result []zkTypes.Log
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = c.rpcClient.CallContext(ctx, &result, "eth_getLogs", arg)
	return result, err
}

// SubscribeFilterLogs performs the same function as SubscribeFilterLogsL2, and that method should be used instead.
// This method is designed to be compatible with bind.ContractBackend.
func (c *BaseClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.ethClient.SubscribeFilterLogs(ctx, query, ch)
}

// SubscribeFilterLogsL2 creates a background log filtering operation, returning
// a subscription immediately, which can be used to stream the found events.
func (c *BaseClient) SubscribeFilterLogsL2(ctx context.Context, query ethereum.FilterQuery, ch chan<- zkTypes.Log) (ethereum.Subscription, error) {
	arg, err := toFilterArg(query)
	if err != nil {
		return nil, err
	}
	sub, err := c.rpcClient.EthSubscribe(ctx, ch, "logs", arg)
	if err != nil {
		// Defensively prefer returning nil interface explicitly on error-path, instead
		// of letting default golang behavior wrap it with non-nil interface that stores
		// nil concrete type value.
		return nil, err
	}
	return sub, nil
}

// PendingBalanceAt returns the wei balance of the given account in the pending state.
func (c *BaseClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	return c.ethClient.PendingBalanceAt(ctx, account)
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (c *BaseClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	return c.ethClient.PendingStorageAt(ctx, account, key)
}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (c *BaseClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return c.ethClient.PendingCodeAt(ctx, account)
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (c *BaseClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return c.ethClient.PendingNonceAt(ctx, account)
}

// PendingTransactionCount returns the total number of transactions in the pending state.
func (c *BaseClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	return c.ethClient.PendingTransactionCount(ctx)
}

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (c *BaseClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// CallContractL2 is almost the same as CallContract except that it executes a message call
// for EIP-712 transaction.
func (c *BaseClient) CallContractL2(ctx context.Context, msg zkTypes.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", msg, toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// CallContractAtHash is almost the same as CallContract except that it selects
// the block by block hash instead of block height.
func (c *BaseClient) CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", toCallArg(msg), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// CallContractAtHashL2 is almost the same as CallContractL2 except that it selects
// the block by block hash instead of block height.
func (c *BaseClient) CallContractAtHashL2(ctx context.Context, msg zkTypes.CallMsg, blockHash common.Hash) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", msg, rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (c *BaseClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", toCallArg(msg), "pending")
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// PendingCallContractL2 executes a message call for EIP-712 transaction using the EVM.
// The state seen by the contract call is the pending state.
func (c *BaseClient) PendingCallContractL2(ctx context.Context, msg zkTypes.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", msg, "pending")
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (c *BaseClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.ethClient.SuggestGasPrice(ctx)
}

// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
// allow a timely execution of a transaction.
func (c *BaseClient) SuggestGasTipCap(_ context.Context) (*big.Int, error) {
	return utils.MaxPriorityFeePerGas, nil
}

// EstimateGas tries to estimate the gas needed to execute a transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (c *BaseClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	var hex hexutil.Uint64
	err := c.rpcClient.CallContext(ctx, &hex, "eth_estimateGas", toCallArg(call))
	if err != nil {
		return 0, err
	}
	return uint64(hex), nil
}

// EstimateGasL2 is almost the same as EstimateGas except that it executes an EIP-712 transaction.
func (c *BaseClient) EstimateGasL2(ctx context.Context, msg zkTypes.CallMsg) (uint64, error) {
	var hex hexutil.Uint64
	err := c.rpcClient.CallContext(ctx, &hex, "eth_estimateGas", msg)
	if err != nil {
		return 0, fmt.Errorf("failed to query eth_estimateGas: %w", err)
	}
	return uint64(hex), nil
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (c *BaseClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.ethClient.SendTransaction(ctx, tx)
}

// SendRawTransaction injects a signed raw transaction into the pending pool for execution.
func (c *BaseClient) SendRawTransaction(ctx context.Context, tx []byte) (common.Hash, error) {
	var res string
	err := c.rpcClient.CallContext(ctx, &res, "eth_sendRawTransaction", hexutil.Encode(tx))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to call eth_sendRawTransaction: %w", err)
	}
	return common.HexToHash(res), nil
}

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func (c *BaseClient) WaitMined(ctx context.Context, txHash common.Hash) (*zkTypes.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	for {
		receipt, err := c.TransactionReceipt(ctx, txHash)
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

// WaitFinalized waits for tx to be finalized on the blockchain.
// It stops waiting when the context is canceled.
func (c *BaseClient) WaitFinalized(ctx context.Context, txHash common.Hash) (*zkTypes.Receipt, error) {
	receipt, err := c.WaitMined(ctx, txHash)
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
		err = c.rpcClient.CallContext(ctx, &blockHead, "eth_getBlockByNumber", zkTypes.BlockNumberFinalized, false)
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

// BridgehubContractAddress returns the Bridgehub smart contract address.
func (c *BaseClient) BridgehubContractAddress(ctx context.Context) (common.Address, error) {
	if c.bridgehubAddress == (common.Address{}) {
		var res string
		err := c.rpcClient.CallContext(ctx, &res, "zks_getBridgehubContract")
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to query zks_getBridgehubContract: %w", err)
		}
		c.bridgehubAddress = common.HexToAddress(res)
	}
	return c.bridgehubAddress, nil
}

// MainContractAddress returns the address of the ZKsync Era contract.
func (c *BaseClient) MainContractAddress(ctx context.Context) (common.Address, error) {
	if c.mainContractAddress == (common.Address{}) {
		var res string
		err := c.rpcClient.CallContext(ctx, &res, "zks_getMainContract")
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to query zks_getMainContract: %w", err)
		}
		c.mainContractAddress = common.HexToAddress(res)
	}
	return c.mainContractAddress, nil

}

// TestnetPaymaster returns the testnet paymaster address if available, or nil.
func (c *BaseClient) TestnetPaymaster(ctx context.Context) (common.Address, error) {
	if c.testnetPaymasterAddress == (common.Address{}) {
		var res string
		err := c.rpcClient.CallContext(ctx, &res, "zks_getTestnetPaymaster")
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to query zks_getTestnetPaymaster: %w", err)
		}
		c.testnetPaymasterAddress = common.HexToAddress(res)
	}
	return c.testnetPaymasterAddress, nil
}

// BridgeContracts returns the addresses of the default ZKsync Era bridge
// contracts on both L1 and L2.
func (c *BaseClient) BridgeContracts(ctx context.Context) (*zkTypes.BridgeContracts, error) {
	if c.bridgeContracts == nil {
		res := zkTypes.BridgeContracts{}
		err := c.rpcClient.CallContext(ctx, &res, "zks_getBridgeContracts")
		if err != nil {
			return nil, fmt.Errorf("failed to query zks_getBridgeContracts: %w", err)
		}
		c.bridgeContracts = &res
	}
	return c.bridgeContracts, nil
}

// BaseTokenContractAddress returns the L1 base token address.
func (c *BaseClient) BaseTokenContractAddress(ctx context.Context) (common.Address, error) {
	if c.baseTokenAddress == (common.Address{}) {
		var res string
		err := c.rpcClient.CallContext(ctx, &res, "zks_getBaseTokenL1Address")
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to query zks_getBaseTokenL1Address: %w", err)
		}
		c.baseTokenAddress = common.HexToAddress(res)
	}
	return c.baseTokenAddress, nil
}

// IsEthBasedChain returns whether the chain is ETH-based.
func (c *BaseClient) IsEthBasedChain(ctx context.Context) (bool, error) {
	baseToken, err := c.BaseTokenContractAddress(ctx)
	if err != nil {
		return false, err
	}
	return baseToken == utils.EthAddressInContracts, nil
}

// IsBaseToken returns whether the token is the base token.
func (c *BaseClient) IsBaseToken(ctx context.Context, token common.Address) (bool, error) {
	if token == utils.LegacyEthAddress {
		token = utils.EthAddressInContracts
	}
	baseToken, err := c.BaseTokenContractAddress(ctx)
	if err != nil {
		return false, err
	}
	return token == baseToken || token == utils.L2BaseTokenAddress, nil
}

// ContractAccountInfo returns the version of the supported account abstraction
// and nonce ordering from a given contract address.
func (c *BaseClient) ContractAccountInfo(ctx context.Context, address common.Address) (*zkTypes.ContractAccountInfo, error) {
	contractDeployer, err := contractdeployer.NewIContractDeployerCaller(utils.ContractDeployerAddress, c)
	if err != nil {
		return nil, err
	}
	accountInfo, err := contractDeployer.GetAccountInfo(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}
	return &zkTypes.ContractAccountInfo{
		SupportedAAVersion: zkTypes.AccountAbstractionVersion(accountInfo.SupportedAAVersion),
		NonceOrdering:      zkTypes.AccountNonceOrdering(accountInfo.NonceOrdering),
	}, nil
}

// L1ChainID returns the chain id of the underlying L1.
func (c *BaseClient) L1ChainID(ctx context.Context) (*big.Int, error) {
	var res string
	err := c.rpcClient.CallContext(ctx, &res, "zks_L1ChainId")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_L1ChainId: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// L1BatchNumber returns the latest L1 batch number.
func (c *BaseClient) L1BatchNumber(ctx context.Context) (*big.Int, error) {
	var res string
	err := c.rpcClient.CallContext(ctx, &res, "zks_L1BatchNumber")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_L1BatchNumber: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

// L1BatchBlockRange returns the range of blocks contained within a batch given by batch number.
func (c *BaseClient) L1BatchBlockRange(ctx context.Context, l1BatchNumber *big.Int) (*BlockRange, error) {
	var resp *BlockRange
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getL1BatchBlockRange", l1BatchNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL1BatchBlockRange: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// L1BatchDetails returns data pertaining to a given batch.
func (c *BaseClient) L1BatchDetails(ctx context.Context, l1BatchNumber *big.Int) (*zkTypes.BatchDetails, error) {
	var resp *zkTypes.BatchDetails
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getL1BatchDetails", l1BatchNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL1BatchDetails: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// BlockDetails returns additional ZKsync Era-specific information about the L2 block.
func (c *BaseClient) BlockDetails(ctx context.Context, block uint32) (*zkTypes.BlockDetails, error) {
	var resp *zkTypes.BlockDetails
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getBlockDetails", block)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getBlockDetails: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// TransactionDetails returns data from a specific transaction given by the transaction hash.
func (c *BaseClient) TransactionDetails(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionDetails, error) {
	var resp *zkTypes.TransactionDetails
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getTransactionDetails", txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getTransactionDetails: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// BytecodeByHash returns bytecode of a contract given by its hash.
func (c *BaseClient) BytecodeByHash(ctx context.Context, bytecodeHash common.Hash) ([]byte, error) {
	var resp []byte
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getBytecodeByHash", bytecodeHash)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getBytecodeByHash: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// RawBlockTransactions returns data of transactions in a block.
func (c *BaseClient) RawBlockTransactions(ctx context.Context, number uint64) ([]zkTypes.RawBlockTransaction, error) {
	var resp []zkTypes.RawBlockTransaction
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getRawBlockTransactions", number)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getRawBlockTransactions: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// LogProof returns the proof for a transaction's L2 to L1 log sent via the L1Messenger system contract.
func (c *BaseClient) LogProof(ctx context.Context, txHash common.Hash, logIndex int) (*zkTypes.LogProof, error) {
	var resp *zkTypes.LogProof
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getL2ToL1LogProof", txHash, logIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL2ToL1LogProof: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Endpoint will be deprecated in favor of LogProof
func (c *BaseClient) MsgProof(ctx context.Context, block uint32, sender common.Address, msg common.Hash) (*zkTypes.MessageProof, error) {
	var resp *zkTypes.MessageProof
	err := c.rpcClient.CallContext(ctx, &resp, "zks_getL2ToL1MsgProof", block, sender, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL2ToL1MsgProof: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Proof returns Merkle proofs for one or more storage values at the specified account along with a Merkle proof of their authenticity.
func (c *BaseClient) Proof(ctx context.Context, address common.Address, keys []common.Hash, l1BatchNumber *big.Int) (*zkTypes.StorageProof, error) {
	var res zkTypes.StorageProof
	err := c.rpcClient.CallContext(ctx, &res, "zks_getProof", address, keys, l1BatchNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getProof: %w", err)
	}
	return &res, nil
}

// SendRawTransactionWithDetailedOutput executes a transaction and returns its hash, storage logs, and events that would
// have been generated if the transaction had already been included in the block. The API has a similar behaviour to
// `eth_sendRawTransaction` but with some extra data returned from it.
//
// With this API Consumer apps can apply "optimistic" events in their applications instantly without having to
// wait for ZKsync block confirmation time.
//
// It’s expected that the optimistic logs of two uncommitted transactions that modify the same state will not
// have causal relationships between each other.
func (c *BaseClient) SendRawTransactionWithDetailedOutput(ctx context.Context, tx []byte) (*zkTypes.TransactionWithDetailedOutput, error) {
	var res zkTypes.TransactionWithDetailedOutput
	err := c.rpcClient.CallContext(ctx, &res, "zks_sendRawTransactionWithDetailedOutput", hexutil.Encode(tx))
	if err != nil {
		return nil, fmt.Errorf("failed to call zks_sendRawTransactionWithDetailedOutput: %w", err)
	}
	return &res, nil
}

// L2TransactionFromPriorityOp returns transaction on L2 network from transaction
// receipt on L1 network.
func (c *BaseClient) L2TransactionFromPriorityOp(ctx context.Context, l1TxReceipt *types.Receipt) (*zkTypes.TransactionResponse, error) {
	if c.mainContract == nil {
		if err := c.cacheMainContract(ctx); err != nil {
			return nil, err
		}
	}

	for _, l := range l1TxReceipt.Logs {
		if l.Address == c.mainContractAddress {
			req, err := c.mainContract.ParseNewPriorityRequest(*l)
			if err != nil {
				return nil, fmt.Errorf("failed to ParseNewPriorityRequest: %w", err)
			}
			_, err = c.WaitMined(ctx, req.TxHash)
			if err != nil {
				return nil, err
			}
			tx, _, err := c.TransactionByHash(ctx, req.TxHash)
			return tx, err
		}
	}
	return nil, errors.New("wrong tx")
}

// PriorityOpConfirmation returns the transaction confirmation data that is part of `L2->L1` message.
// The txHash is the  hash of the L2 transaction where the message was initiated.
// The index is used in case there were multiple transactions in one message, you may pass an index of the
// transaction which confirmation data should be fetched.
func (c *BaseClient) PriorityOpConfirmation(ctx context.Context, txHash common.Hash, index int) (*zkTypes.PriorityOpConfirmation, error) {
	receipt, err := c.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return nil, errors.New("transaction receipt not found")
	}
	fLogs := make([]struct {
		i int
		l *zkTypes.L2ToL1Log
	}, 0)
	for i, l := range receipt.L2ToL1Logs {
		if l.Sender == utils.BootloaderFormalAddress {
			fLogs = append(fLogs, struct {
				i int
				l *zkTypes.L2ToL1Log
			}{i, l})
		}
	}

	if len(fLogs) < index+1 {
		return nil, errors.New("message not found")
	}

	l2ToL1LogIndex := fLogs[index].i
	l2ToL1Log := fLogs[index].l
	proof, err := c.LogProof(ctx, txHash, l2ToL1LogIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}

	return &zkTypes.PriorityOpConfirmation{
		L1BatchNumber:     l2ToL1Log.L1BatchNumber.ToInt(),
		L2MessageIndex:    proof.Id,
		L2TxNumberInBlock: receipt.L1BatchNumber.ToInt(),
		Proof:             proof.Proof,
	}, nil
}

// ConfirmedTokens returns confirmed tokens. Confirmed token is any token bridged to ZKsync Era via the official bridge.
func (c *BaseClient) ConfirmedTokens(ctx context.Context, from uint32, limit uint8) ([]*zkTypes.Token, error) {
	res := make([]*zkTypes.Token, 0)
	err := c.rpcClient.CallContext(ctx, &res, "zks_getConfirmedTokens", from, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getConfirmedTokens: %w", err)
	}
	return res, nil
}

// Deprecated: Method is deprecated and will be removed in the near future.
func (c *BaseClient) TokenPrice(ctx context.Context, address common.Address) (*big.Float, error) {
	var res string
	err := c.rpcClient.CallContext(ctx, &res, "zks_getTokenPrice", address)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getTokenPrice: %w", err)
	}
	resp, ok := big.NewFloat(0).SetString(res)
	if !ok {
		return nil, errors.New("failed to decode response as big.Float")
	}
	return resp, nil
}

// L2TokenAddress returns the L2 token address equivalent for a L1 token address
// as they are not equal. ETH address is set to zero address.
func (c *BaseClient) L2TokenAddress(ctx context.Context, token common.Address) (common.Address, error) {
	if token == utils.LegacyEthAddress {
		token = utils.EthAddressInContracts
	}

	if baseToken, err := c.BaseTokenContractAddress(ctx); err == nil && token == baseToken {
		return utils.L2BaseTokenAddress, nil
	} else if err != nil {
		return common.Address{}, err
	}

	bridgeContracts, err := c.BridgeContracts(ctx)
	if err != nil {
		return common.Address{}, err
	}
	bridge, err := l2bridge.NewIL2Bridge(bridgeContracts.L2SharedBridge, c)
	if err != nil {
		return common.Address{}, err
	}
	tokenAddress, err := bridge.L2TokenAddress(&bind.CallOpts{Context: ctx}, token)
	if err != nil {
		return common.Address{}, err
	}
	return tokenAddress, nil
}

// L2TokenAddressFromCustomBridge returns the L2 token address equivalent for a L1 token address
// using custom bridge. ETH address is set to zero address.
func (c *BaseClient) L2TokenAddressFromCustomBridge(ctx context.Context, token, bridge common.Address) (common.Address, error) {
	if token == utils.LegacyEthAddress {
		token = utils.EthAddressInContracts
	}

	if baseToken, err := c.BaseTokenContractAddress(ctx); err == nil && token == baseToken {
		return utils.L2BaseTokenAddress, nil
	} else if err != nil {
		return common.Address{}, err
	}

	bridgeContract, err := l2bridge.NewIL2Bridge(bridge, c)
	if err != nil {
		return common.Address{}, err
	}
	tokenAddress, err := bridgeContract.L2TokenAddress(&bind.CallOpts{Context: ctx}, token)
	if err != nil {
		return common.Address{}, err
	}
	return tokenAddress, nil
}

// L1TokenAddress returns the L1 token address equivalent for a L2 token address
// as they are not equal. ETH address is set to zero address.
func (c *BaseClient) L1TokenAddress(ctx context.Context, token common.Address) (common.Address, error) {
	if token == utils.LegacyEthAddress {
		return utils.LegacyEthAddress, nil
	} else {
		bridgeContracts, err := c.BridgeContracts(ctx)
		if err != nil {
			return common.Address{}, err
		}
		bridge, err := l2bridge.NewIL2Bridge(bridgeContracts.L2SharedBridge, c)
		if err != nil {
			return common.Address{}, err
		}
		tokenAddress, err := bridge.L1TokenAddress(&bind.CallOpts{Context: ctx}, token)
		if err != nil {
			return common.Address{}, err
		}
		return tokenAddress, nil
	}
}

// ProtocolVersion return the protocol version
func (c *BaseClient) ProtocolVersion(ctx context.Context) (*zkTypes.ProtocolVersion, error) {
	var res zkTypes.ProtocolVersion
	err := c.rpcClient.CallContext(ctx, &res, "zks_getProtocolVersion")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getProtocolVersion: %w", err)
	}
	return &res, nil
}

// AllAccountBalances returns all balances for confirmed tokens given by an account address.
func (c *BaseClient) AllAccountBalances(ctx context.Context, address common.Address) (map[common.Address]*big.Int, error) {
	res := make(map[common.Address]string)
	err := c.rpcClient.CallContext(ctx, &res, "zks_getAllAccountBalances", address)
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

// EstimateFee returns the fee for the transaction.
func (c *BaseClient) EstimateFee(ctx context.Context, msg zkTypes.CallMsg) (*zkTypes.Fee, error) {
	var res zkTypes.Fee
	err := c.rpcClient.CallContext(ctx, &res, "zks_estimateFee", msg)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_estimateFee: %w", err)
	}
	return &res, nil
}

// FeeParams returns the current fee parameters.
func (c *BaseClient) FeeParams(ctx context.Context) (*zkTypes.FeeParams, error) {
	var res zkTypes.FeeParams
	err := c.rpcClient.CallContext(ctx, &res, "zks_getFeeParams")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getFeeParams: %w", err)
	}
	return &res, nil
}

// EstimateGasL1 estimates the amount of gas required to submit a transaction from L1 to L2.
func (c *BaseClient) EstimateGasL1(ctx context.Context, msg zkTypes.CallMsg) (uint64, error) {
	var res hexutil.Uint64
	err := c.rpcClient.CallContext(ctx, &res, "zks_estimateGasL1ToL2", msg)
	if err != nil {
		return 0, fmt.Errorf("failed to query zks_estimateGasL1ToL2: %w", err)
	}
	return uint64(res), nil
}

// EstimateGasTransfer estimates the amount of gas required for a transfer transaction.
func (c *BaseClient) EstimateGasTransfer(ctx context.Context, msg TransferCallMsg) (uint64, error) {
	callMsg, err := msg.ToZkCallMsg()
	if err != nil {
		return 0, err
	}
	return c.EstimateGasL2(ctx, *callMsg)
}

// EstimateGasWithdraw estimates the amount of gas required for a withdrawal transaction.
func (c *BaseClient) EstimateGasWithdraw(ctx context.Context, msg WithdrawalCallMsg) (uint64, error) {
	var (
		callMsg *zkTypes.CallMsg
		err     error
	)

	if msg.Token == utils.LegacyEthAddress || msg.Token == utils.EthAddressInContracts {
		msg.Token, err = c.L2TokenAddress(ctx, msg.Token)
		if err != nil {
			return 0, err
		}
	}

	if msg.BridgeAddress == nil {
		contracts, errBridge := c.BridgeContracts(ctx)
		if errBridge != nil {
			return 0, fmt.Errorf("failed to getBridgeContracts: %w", errBridge)
		}
		callMsg, err = msg.ToZkCallMsg(&contracts.L2SharedBridge)
	} else {
		callMsg, err = msg.ToZkCallMsg(nil)
		if err != nil {
			return 0, err
		}
	}
	return c.EstimateGasL2(ctx, *callMsg)
}

// EstimateL1ToL2Execute estimates the amount of gas required for an L1 to L2 execute operation.
func (c *BaseClient) EstimateL1ToL2Execute(ctx context.Context, msg zkTypes.CallMsg) (uint64, error) {
	if msg.Meta == nil || msg.Meta.GasPerPubdata == nil {
		msg.Meta = &zkTypes.Eip712Meta{GasPerPubdata: utils.NewBig(utils.RequiredL1ToL2GasPerPubdataLimit.Int64())}
	}

	// If the `from` address is not provided, we use a random address, because
	// due to storage slot aggregation, the gas estimation will depend on the address
	// and so estimation for the zero address may be smaller than for the sender.
	if msg.From == (common.Address{}) {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return 0, fmt.Errorf("failed to generate radnom private key: %w", err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return 0, fmt.Errorf("failed to convert public key to ECDSA")
		}
		msg.From = crypto.PubkeyToAddress(*publicKeyECDSA)
	}

	return c.EstimateGasL1(ctx, msg)
}

// IsL2BridgeLegacy returns true if passed bridge address is legacy and false if its shared bridge.
func (c *BaseClient) IsL2BridgeLegacy(ctx context.Context, address common.Address) (bool, error) {
	bridge, err := l2sharedbridge.NewIL2SharedBridge(address, c)
	if err != nil {
		return false, err
	}
	_, err = bridge.L1SharedBridge(&bind.CallOpts{Context: ctx})
	if err != nil {
		return true, nil
	}
	return false, nil
}

// IsMessageSignatureCorrect checks whether the account abstraction message signature is correct.
// Signature can be created using EIP1271 or ECDSA.
func (c *BaseClient) IsMessageSignatureCorrect(ctx context.Context, address common.Address, msg, sig []byte) (bool, error) {
	code, err := c.CodeAt(ctx, address, nil)
	if err != nil {
		return false, err
	}
	if len(code) > 0 {
		return c.isEIP1271SignatureCorrect(ctx, address, common.Hash(accounts.TextHash(msg)), sig)
	}
	return utils.IsMessageSignatureCorrect(address, msg, sig)
}

// IsTypedDataSignatureCorrect checks whether the account abstraction typed data signature is correct.
// Signature can be created using EIP1271 or ECDSA.
func (c *BaseClient) IsTypedDataSignatureCorrect(ctx context.Context, address common.Address, typedData apitypes.TypedData, sig []byte) (bool, error) {
	code, err := c.CodeAt(ctx, address, nil)
	if err != nil {
		return false, err
	}
	if len(code) > 0 {
		hash, _, errHash := apitypes.TypedDataAndHash(typedData)
		if errHash != nil {
			return false, errHash
		}
		return c.isEIP1271SignatureCorrect(ctx, address, common.Hash(hash), sig)
	}
	return utils.IsTypedDataSignatureCorrect(address, typedData, sig)
}

func (c *BaseClient) cacheMainContract(ctx context.Context) error {
	mainContractAddress, err := c.MainContractAddress(ctx)
	if err != nil {
		return err
	}
	// parsing events does not require backend to be set
	mainContract, err := zksync.NewIZkSync(mainContractAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to load IZkSync: %w", err)
	}
	c.mainContract = mainContract
	c.mainContractAddress = mainContractAddress
	return nil
}

func (c *BaseClient) isEIP1271SignatureCorrect(ctx context.Context, address common.Address, hash common.Hash, sig []byte) (bool, error) {
	account, err := erc1271.NewIERC1271(address, c)
	if err != nil {
		return false, err
	}
	value, err := account.IsValidSignature(&bind.CallOpts{Context: ctx}, hash, sig)
	if err != nil {
		return false, err
	}
	return value == utils.Eip1271MagicValue, nil
}

func (c *BaseClient) getBlock(ctx context.Context, method string, args ...interface{}) (*zkTypes.Block, error) {
	var raw json.RawMessage
	err := c.rpcClient.CallContext(ctx, &raw, method, args...)
	if err != nil {
		return nil, err
	}

	var block *blockMarshaling
	if err = json.Unmarshal(raw, &block); err != nil {
		return nil, err
	}
	if block == nil {
		return nil, ethereum.NotFound
	}

	// Quick-verify transaction and uncle lists. This mostly helps with debugging the server.
	if block.UncleHash == types.EmptyUncleHash && len(block.Uncles) > 0 {
		return nil, errors.New("server returned non-empty uncle list but block header indicates no uncles")
	}
	if block.UncleHash != types.EmptyUncleHash && len(block.Uncles) == 0 {
		return nil, errors.New("server returned empty uncle list but block header indicates uncles")
	}
	if block.TxHash == types.EmptyTxsHash && len(block.Transactions) > 0 {
		return nil, errors.New("server returned non-empty transaction list but block header indicates no transactions")
	}
	// TODO use this validation when transaction root is fixed on zksync node
	//if head.TxHash != types.EmptyTxsHash && len(body.Transactions) == 0 {
	//	return nil, errors.New("server returned empty transaction list but block header indicates transactions")
	//}
	// Load uncles because they are not included in the block response.
	var uncles []*types.Header
	if len(block.Uncles) > 0 {
		uncles = make([]*types.Header, len(block.Uncles))
		reqs := make([]rpc.BatchElem, len(block.Uncles))
		for i := range reqs {
			reqs[i] = rpc.BatchElem{
				Method: "eth_getUncleByBlockHashAndIndex",
				Args:   []interface{}{block.Hash, hexutil.EncodeUint64(uint64(i))},
				Result: &uncles[i],
			}
		}
		if err := c.rpcClient.BatchCallContext(ctx, reqs); err != nil {
			return nil, err
		}
		for i := range reqs {
			if reqs[i].Error != nil {
				return nil, reqs[i].Error
			}
			if uncles[i] == nil {
				return nil, fmt.Errorf("got null header for uncle %d of block %x", i, block.Hash[:])
			}
		}
	}

	return &zkTypes.Block{
		Header: &types.Header{
			ParentHash:       block.ParentHash,
			UncleHash:        block.UncleHash,
			Coinbase:         block.Coinbase,
			Root:             block.Root,
			TxHash:           block.TxHash,
			ReceiptHash:      block.ReceiptHash,
			Bloom:            block.Bloom,
			Difficulty:       block.Difficulty.ToInt(),
			Number:           block.Number.ToInt(),
			GasLimit:         uint64(block.GasLimit),
			GasUsed:          uint64(block.GasUsed),
			Time:             uint64(block.Time),
			Extra:            block.Extra,
			MixDigest:        block.MixDigest,
			Nonce:            block.Nonce,
			BaseFee:          block.BaseFee.ToInt(),
			WithdrawalsHash:  nil,
			BlobGasUsed:      nil,
			ExcessBlobGas:    nil,
			ParentBeaconRoot: nil,
		},
		Uncles:           uncles,
		Transactions:     block.Transactions,
		Hash:             *block.Hash,
		Size:             block.Size.ToInt(),
		TotalDifficulty:  block.TotalDifficulty.ToInt(),
		SealFields:       block.SealFields,
		L1BatchNumber:    block.L1BatchNumber.ToInt(),
		L1BatchTimestamp: block.L1BatchTimestamp.ToInt(),
	}, nil
}
