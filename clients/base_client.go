package clients

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/contracts/contractdeployer"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
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

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) Client {
	return &BaseClient{
		rpcClient: c,
		ethClient: ethclient.NewClient(c),
	}
}

func (c *BaseClient) Client() *rpc.Client {
	return c.rpcClient
}

func (c *BaseClient) Close() {
	c.ethClient.Close()
}

func (c *BaseClient) ChainID(ctx context.Context) (*big.Int, error) {
	return c.ethClient.ChainID(ctx)
}

func (c *BaseClient) BlockByHash(ctx context.Context, hash common.Hash) (*zkTypes.Block, error) {
	return c.getBlock(ctx, "eth_getBlockByHash", hash, true)
}

func (c *BaseClient) BlockByNumber(ctx context.Context, number *big.Int) (*zkTypes.Block, error) {
	return c.getBlock(ctx, "eth_getBlockByNumber", toBlockNumArg(number), true)
}

func (c *BaseClient) BlockNumber(ctx context.Context) (uint64, error) {
	return c.ethClient.BlockNumber(ctx)
}

func (c *BaseClient) PeerCount(ctx context.Context) (uint64, error) {
	return c.ethClient.PeerCount(ctx)
}

func (c *BaseClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return c.ethClient.HeaderByHash(ctx, hash)
}

func (c *BaseClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.ethClient.HeaderByNumber(ctx, number)
}

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

func (c *BaseClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return c.ethClient.TransactionCount(ctx, blockHash)
}

func (c *BaseClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*zkTypes.TransactionResponse, error) {
	var tx *zkTypes.TransactionResponse
	err := c.rpcClient.CallContext(ctx, &tx, "eth_getTransactionByBlockHashAndIndex", blockHash, hexutil.Uint64(index))
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, ethereum.NotFound
	} else if tx.R == nil {
		return nil, errors.New("server returned transaction without signature")
	}
	return tx, err
}

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

func (c *BaseClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	return c.ethClient.SyncProgress(ctx)
}

func (c *BaseClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return c.ethClient.SubscribeNewHead(ctx, ch)
}

func (c *BaseClient) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.ethClient.NetworkID(ctx)
}

func (c *BaseClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return c.ethClient.BalanceAt(ctx, account, blockNumber)
}

func (c *BaseClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	return c.ethClient.StorageAt(ctx, account, key, blockNumber)
}

func (c *BaseClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return c.ethClient.CodeAt(ctx, account, blockNumber)
}

func (c *BaseClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return c.ethClient.NonceAt(ctx, account, blockNumber)
}

func (c *BaseClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return c.ethClient.FilterLogs(ctx, query)
}

func (c *BaseClient) FilterLogsL2(ctx context.Context, q ethereum.FilterQuery) ([]zkTypes.Log, error) {
	var result []zkTypes.Log
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = c.rpcClient.CallContext(ctx, &result, "eth_getLogs", arg)
	return result, err
}

func (c *BaseClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.ethClient.SubscribeFilterLogs(ctx, query, ch)
}

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

func (c *BaseClient) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	return c.ethClient.PendingBalanceAt(ctx, account)
}
func (c *BaseClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	return c.ethClient.PendingStorageAt(ctx, account, key)
}
func (c *BaseClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return c.ethClient.PendingCodeAt(ctx, account)
}
func (c *BaseClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return c.ethClient.PendingNonceAt(ctx, account)
}

func (c *BaseClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	return c.ethClient.PendingTransactionCount(ctx)
}

func (c *BaseClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.ethClient.CallContract(ctx, msg, blockNumber)
}

func (c *BaseClient) CallContractL2(ctx context.Context, msg zkTypes.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", msg, toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func (c *BaseClient) CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error) {
	return c.ethClient.CallContractAtHash(ctx, msg, blockHash)
}

func (c *BaseClient) CallContractAtHashL2(ctx context.Context, msg zkTypes.CallMsg, blockHash common.Hash) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", msg, rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func (c *BaseClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	return c.ethClient.PendingCallContract(ctx, msg)
}

func (c *BaseClient) PendingCallContractL2(ctx context.Context, msg zkTypes.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpcClient.CallContext(ctx, &hex, "eth_call", msg, "pending")
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func (c *BaseClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.ethClient.SuggestGasPrice(ctx)
}

func (c *BaseClient) SuggestGasTipCap(_ context.Context) (*big.Int, error) {
	return utils.MaxPriorityFeePerGas, nil
}

func (c *BaseClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return c.ethClient.EstimateGas(ctx, call)
}

func (c *BaseClient) EstimateGasL2(ctx context.Context, msg zkTypes.CallMsg) (uint64, error) {
	var hex hexutil.Uint64
	err := c.rpcClient.CallContext(ctx, &hex, "eth_estimateGas", msg)
	if err != nil {
		return 0, fmt.Errorf("failed to query eth_estimateGas: %w", err)
	}
	return uint64(hex), nil
}

func (c *BaseClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.ethClient.SendTransaction(ctx, tx)
}

func (c *BaseClient) SendRawTransaction(ctx context.Context, tx []byte) (common.Hash, error) {
	var res string
	err := c.rpcClient.CallContext(ctx, &res, "eth_sendRawTransaction", hexutil.Encode(tx))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to call eth_sendRawTransaction: %w", err)
	}
	return common.HexToHash(res), nil
}

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

func (c *BaseClient) MainContractAddress(ctx context.Context) (common.Address, error) {
	var res string
	err := c.rpcClient.CallContext(ctx, &res, "zks_getMainContract")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to query zks_getMainContract: %w", err)
	}
	return common.HexToAddress(res), nil
}

func (c *BaseClient) TestnetPaymaster(ctx context.Context) (common.Address, error) {
	var res string
	err := c.rpcClient.CallContext(ctx, &res, "zks_getTestnetPaymaster")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to query zks_getTestnetPaymaster: %w", err)
	}
	return common.HexToAddress(res), nil
}

func (c *BaseClient) BridgeContracts(ctx context.Context) (*zkTypes.BridgeContracts, error) {
	res := zkTypes.BridgeContracts{}
	err := c.rpcClient.CallContext(ctx, &res, "zks_getBridgeContracts")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getBridgeContracts: %w", err)
	}
	return &res, nil
}

func (c *BaseClient) ContractAccountInfo(ctx context.Context, address common.Address) (*zkTypes.ContractAccountInfo, error) {
	contractDeployer, err := contractdeployer.NewContractDeployerCaller(utils.ContractDeployerAddress, c)
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

func (c *BaseClient) LogProof(ctx context.Context, txHash common.Hash, logIndex int) (*zkTypes.MessageProof, error) {
	var resp *zkTypes.MessageProof
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

func (c *BaseClient) ConfirmedTokens(ctx context.Context, from uint32, limit uint8) ([]*zkTypes.Token, error) {
	res := make([]*zkTypes.Token, 0)
	err := c.rpcClient.CallContext(ctx, &res, "zks_getConfirmedTokens", from, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getConfirmedTokens: %w", err)
	}
	return res, nil
}

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

func (c *BaseClient) L2TokenAddress(ctx context.Context, token common.Address) (common.Address, error) {
	if token == utils.EthAddress {
		return utils.EthAddress, nil
	} else {
		bridgeContracts, err := c.BridgeContracts(ctx)
		if err != nil {
			return common.Address{}, err
		}
		bridge, err := l2bridge.NewIL2Bridge(bridgeContracts.L2Erc20DefaultBridge, c)
		if err != nil {
			return common.Address{}, err
		}
		tokenAddress, err := bridge.L2TokenAddress(&bind.CallOpts{Context: ctx}, token)
		if err != nil {
			return common.Address{}, err
		}
		return tokenAddress, nil
	}
}

func (c *BaseClient) L1TokenAddress(ctx context.Context, token common.Address) (common.Address, error) {
	if token == utils.EthAddress {
		return utils.EthAddress, nil
	} else {
		bridgeContracts, err := c.BridgeContracts(ctx)
		if err != nil {
			return common.Address{}, err
		}
		bridge, err := l2bridge.NewIL2Bridge(bridgeContracts.L2Erc20DefaultBridge, c)
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

func (c *BaseClient) EstimateFee(ctx context.Context, msg zkTypes.CallMsg) (*zkTypes.Fee, error) {
	var res zkTypes.Fee
	err := c.rpcClient.CallContext(ctx, &res, "zks_estimateFee", msg)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_estimateFee: %w", err)
	}
	return &res, nil
}

func (c *BaseClient) EstimateGasL1(ctx context.Context, msg zkTypes.CallMsg) (uint64, error) {
	var res hexutil.Uint64
	err := c.rpcClient.CallContext(ctx, &res, "zks_estimateGasL1ToL2", msg)
	if err != nil {
		return 0, fmt.Errorf("failed to query zks_estimateGasL1ToL2: %w", err)
	}
	return uint64(res), nil
}

func (c *BaseClient) EstimateGasTransfer(ctx context.Context, msg TransferCallMsg) (uint64, error) {
	callMsg, err := msg.ToCallMsg()
	if err != nil {
		return 0, err
	}
	return c.EstimateGas(ctx, *callMsg)
}

func (c *BaseClient) EstimateGasWithdraw(ctx context.Context, msg WithdrawalCallMsg) (uint64, error) {
	var (
		callMsg *ethereum.CallMsg
		err     error
	)

	if msg.BridgeAddress == nil {
		contracts, errBridge := c.BridgeContracts(ctx)
		if errBridge != nil {
			return 0, fmt.Errorf("failed to getBridgeContracts: %w", errBridge)
		}
		callMsg, err = msg.ToCallMsg(&contracts.L2Erc20DefaultBridge)
	} else {
		callMsg, err = msg.ToCallMsg(nil)
		if err != nil {
			return 0, err
		}
	}
	return c.EstimateGas(ctx, *callMsg)
}

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
			ParentHash:      block.ParentHash,
			UncleHash:       block.UncleHash,
			Coinbase:        block.Coinbase,
			Root:            block.Root,
			TxHash:          block.TxHash,
			ReceiptHash:     block.ReceiptHash,
			Bloom:           block.Bloom,
			Difficulty:      block.Difficulty.ToInt(),
			Number:          block.Number.ToInt(),
			GasLimit:        uint64(block.GasLimit),
			GasUsed:         uint64(block.GasUsed),
			Time:            uint64(block.Time),
			Extra:           block.Extra,
			MixDigest:       block.MixDigest,
			Nonce:           block.Nonce,
			BaseFee:         block.BaseFee.ToInt(),
			WithdrawalsHash: nil,
			ExcessDataGas:   block.ExcessDataGas.ToInt(),
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
