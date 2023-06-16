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

type ethereumClient interface {
	//ChainID(ctx context.Context) (*big.Int, error)
	//BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	//BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	//BlockNumber(ctx context.Context) (uint64, error)
	//PeerCount(ctx context.Context)

	GetClient() *ethclient.Client
	GetBalance(address common.Address, blockNumber zkTypes.BlockNumber) (*big.Int, error)
	GetBlockByNumber(blockNumber zkTypes.BlockNumber) (*zkTypes.Block, error)
	GetBlockByHash(blockHash common.Hash) (*zkTypes.Block, error)
	GetTransactionCount(address common.Address, blockNumber zkTypes.BlockNumber) (*big.Int, error)
	GetTransactionReceipt(txHash common.Hash) (*zkTypes.TransactionReceipt, error)
	GetTransaction(txHash common.Hash) (*zkTypes.TransactionResponse, error)
	WaitMined(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionReceipt, error)
	WaitFinalized(ctx context.Context, txHash common.Hash) (*zkTypes.TransactionReceipt, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error)
	GetGasPrice() (*big.Int, error)
	SendRawTransaction(tx []byte) (common.Hash, error)
	GetLogs(q zkTypes.FilterQuery) ([]zkTypes.Log, error)
}

type zksyncClient interface {
	ZksGetMainContract() (common.Address, error)
	ZksL1ChainId() (*big.Int, error)
	ZksL1BatchNumber() (*big.Int, error)
	ZksGetConfirmedTokens(from uint32, limit uint8) ([]*zkTypes.Token, error)
	ZksIsTokenLiquid(address common.Address) (bool, error)
	ZksGetTokenPrice(address common.Address) (*big.Float, error)
	L2TokenAddress(token common.Address) (common.Address, error)
	L1TokenAddress(token common.Address) (common.Address, error)
	ZksGetL2ToL1LogProof(txHash common.Hash, logIndex int) (*zkTypes.L2ToL1MessageProof, error)
	ZksGetL2ToL1MsgProof(block uint32, sender common.Address, msg common.Hash) (*zkTypes.L2ToL1MessageProof, error)
	ZksGetAllAccountBalances(address common.Address) (map[common.Address]*big.Int, error)
	ZksGetBridgeContracts() (*zkTypes.BridgeContracts, error)
	ZksEstimateFee(tx *zkTypes.Transaction) (*zkTypes.Fee, error)
	ZksGetTestnetPaymaster() (common.Address, error)
	ZksGetBlockDetails(block uint32) (*zkTypes.BlockDetails, error)
	EstimateGas712(tx *zkTypes.Transaction) (*big.Int, error)
	EstimateGasL1(tx *zkTypes.Transaction) (*big.Int, error)
	EstimateL1ToL2Execute(tx *zkTypes.Transaction) (*big.Int, error)
	GetL1BatchBlockRange(l1BatchNumber *big.Int) (*zkTypes.BlockRange, error)
	GetL1BatchDetails(l1BatchNumber *big.Int) (*zkTypes.BatchDetails, error)
	GetTransactionDetails(txHash common.Hash) (*zkTypes.TransactionDetails, error)
}

type Provider interface {
	ethereumClient
	zksyncClient
}

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

type DefaultProvider struct {
	c *rpc.Client
	// also inherit default Ethereum clients
	*ethclient.Client
}

func (p *DefaultProvider) GetClient() *ethclient.Client {
	return p.Client
}

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
	ethBlock, err := p.Client.BlockByNumber(context.Background(), resp.Number.ToInt())
	if err != nil {
		return nil, err
	}
	return &zkTypes.Block{
		Block:            *ethBlock,
		L1BatchNumber:    resp.L1BatchNumber,
		L1BatchTimestamp: resp.L1BatchTimestamp,
	}, nil
}

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
	ethBlock, err := p.Client.BlockByHash(context.Background(), blockHash)
	if err != nil {
		return nil, err
	}
	return &zkTypes.Block{
		Block:            *ethBlock,
		L1BatchNumber:    resp.L1BatchNumber,
		L1BatchTimestamp: resp.L1BatchTimestamp,
	}, nil
}

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

// EstimateGas does same as EstimateGas712 method and should be used instead.
// This method is implemented to be compatible with bind.ContractBackend
func (p *DefaultProvider) EstimateGas(_ context.Context, call ethereum.CallMsg) (uint64, error) {
	gas, err := p.EstimateGas712(zkTypes.NewTransactionFromCallMsg(call))
	if err != nil {
		return 0, err
	}
	return gas.Uint64(), nil
}

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

func (p *DefaultProvider) SendRawTransaction(tx []byte) (common.Hash, error) {
	var res string
	err := p.c.Call(&res, "eth_sendRawTransaction", hexutil.Encode(tx))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to call eth_sendRawTransaction: %w", err)
	}
	return common.HexToHash(res), nil
}

func (p *DefaultProvider) ZksGetMainContract() (common.Address, error) {
	var res string
	err := p.c.Call(&res, "zks_getMainContract")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to query zks_getMainContract: %w", err)
	}
	return common.HexToAddress(res), nil
}

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

func (p *DefaultProvider) ZksGetConfirmedTokens(from uint32, limit uint8) ([]*zkTypes.Token, error) {
	res := make([]*zkTypes.Token, 0)
	err := p.c.Call(&res, "zks_getConfirmedTokens", from, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getConfirmedTokens: %w", err)
	}
	return res, nil
}

func (p *DefaultProvider) ZksIsTokenLiquid(address common.Address) (bool, error) {
	var res bool
	err := p.c.Call(&res, "zks_isTokenLiquid", address)
	if err != nil {
		return false, fmt.Errorf("failed to query zks_isTokenLiquid: %w", err)
	}
	return res, nil
}

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

func (p *DefaultProvider) ZksGetL2ToL1LogProof(txHash common.Hash, logIndex int) (*zkTypes.L2ToL1MessageProof, error) {
	var resp *zkTypes.L2ToL1MessageProof
	err := p.c.Call(&resp, "zks_getL2ToL1LogProof", txHash, logIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL2ToL1LogProof: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

// Deprecated: Endpoint will be deprecated in favor of ZksGetL2ToL1LogProof
func (p *DefaultProvider) ZksGetL2ToL1MsgProof(block uint32, sender common.Address, msg common.Hash) (*zkTypes.L2ToL1MessageProof, error) {
	var resp *zkTypes.L2ToL1MessageProof
	err := p.c.Call(&resp, "zks_getL2ToL1MsgProof", block, sender, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL2ToL1MsgProof: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

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

func (p *DefaultProvider) ZksGetBridgeContracts() (*zkTypes.BridgeContracts, error) {
	res := zkTypes.BridgeContracts{}
	err := p.c.Call(&res, "zks_getBridgeContracts")
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getBridgeContracts: %w", err)
	}
	return &res, nil
}

func (p *DefaultProvider) ZksEstimateFee(tx *zkTypes.Transaction) (*zkTypes.Fee, error) {
	var res zkTypes.Fee
	err := p.c.Call(&res, "zks_estimateFee", tx)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_estimateFee: %w", err)
	}
	return &res, nil
}

func (p *DefaultProvider) ZksGetTestnetPaymaster() (common.Address, error) {
	var res string
	err := p.c.Call(&res, "zks_getTestnetPaymaster")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to query zks_estimateFee: %w", err)
	}
	return common.HexToAddress(res), nil
}

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

func (p *DefaultProvider) GetLogs(q zkTypes.FilterQuery) ([]zkTypes.Log, error) {
	var result []zkTypes.Log
	arg, err := utils.ToFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = p.c.Call(&result, "eth_getLogs", arg)
	return result, err
}

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

func (p *DefaultProvider) EstimateGasL1(tx *zkTypes.Transaction) (*big.Int, error) {
	var res string
	err := p.c.Call(&res, "zks_estimateGasL1ToL2", tx)
	if err != nil {
		return nil, fmt.Errorf("failed to query eth_estimateGas: %w", err)
	}
	resp, err := hexutil.DecodeBig(res)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response as big.Int: %w", err)
	}
	return resp, nil
}

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

func (p *DefaultProvider) GetL1BatchBlockRange(l1BatchNumber *big.Int) (*zkTypes.BlockRange, error) {
	var resp *zkTypes.BlockRange
	err := p.c.Call(&resp, "zks_getL1BatchBlockRange", l1BatchNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to query zks_getL1BatchBlockRange: %w", err)
	} else if resp == nil {
		return nil, ethereum.NotFound
	}
	return resp, nil
}

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

func (p *DefaultProvider) SuggestGasTipCap(_ context.Context) (*big.Int, error) {
	return utils.MaxPriorityFeePerGas, nil
}
