package zksync2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/contracts/ERC20"
	"github.com/zksync-sdk/zksync2-go/contracts/IL1Bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/IL2Bridge"
	"math/big"
	"strings"
)

type Wallet struct {
	es EthSigner
	zp Provider

	bcs         *BridgeContracts
	erc20abi    abi.ABI
	l2BridgeAbi abi.ABI
}

func NewWallet(es EthSigner, zp Provider) (*Wallet, error) {
	erc20abi, err := abi.JSON(strings.NewReader(ERC20.ERC20MetaData.ABI))
	l2BridgeAbi, err := abi.JSON(strings.NewReader(IL2Bridge.IL2BridgeMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load erc20abi: %w", err)
	}
	return &Wallet{
		es:          es,
		zp:          zp,
		erc20abi:    erc20abi,
		l2BridgeAbi: l2BridgeAbi,
	}, nil
}

func (w *Wallet) CreateEthereumProvider(rpcClient *rpc.Client) (*DefaultEthProvider, error) {
	ethClient := ethclient.NewClient(rpcClient)
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain Id: %w", err)
	}
	auth, err := w.newTransactorWithSigner(w.es, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to init TransactOpts: %w", err)
	}
	bcs, err := w.getBridgeContracts()
	if err != nil {
		return nil, fmt.Errorf("failed to getBridgeContracts: %w", err)
	}
	l1EthBridge, err := IL1Bridge.NewIL1Bridge(bcs.L1EthDefaultBridge, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to load L1EthBridge: %w", err)
	}
	l1ERC20Bridge, err := IL1Bridge.NewIL1Bridge(bcs.L1Erc20DefaultBridge, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to load L1ERC20Bridge: %w", err)
	}
	return NewDefaultEthProvider(rpcClient, auth, l1EthBridge, l1ERC20Bridge), nil
}

func (w *Wallet) GetBalance() (*big.Int, error) {
	return w.zp.GetBalance(w.es.GetAddress(), BlockNumberCommitted)
}

func (w *Wallet) GetBalanceOf(address common.Address, token *Token, at BlockNumber) (*big.Int, error) {
	if token.IsETH() {
		return w.zp.GetBalance(address, at)
	}
	erc20, err := ERC20.NewERC20(token.L2Address, w.zp.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to load ERC20: %w", err)
	}
	opts := &bind.CallOpts{}
	if at == BlockNumberPending {
		opts.Pending = true
	} else if bn, err := hexutil.DecodeBig(string(at)); err == nil && bn != nil {
		opts.BlockNumber = bn
	}
	return erc20.BalanceOf(opts, address)
}

func (w *Wallet) GetNonce() (*big.Int, error) {
	return w.GetNonceAt(BlockNumberCommitted)
}

func (w *Wallet) GetNonceAt(at BlockNumber) (*big.Int, error) {
	return w.zp.GetTransactionCount(w.es.GetAddress(), at)
}

func (w *Wallet) Transfer(to common.Address, amount *big.Int, token *Token, nonce *big.Int) (string, error) {
	var err error
	if token == nil {
		token = CreateETH()
	}
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return "", fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	var data hexutil.Bytes
	if !token.IsETH() {
		data, err = w.erc20abi.Pack("transfer", to, amount)
		if err != nil {
			return "", fmt.Errorf("failed to pack transfer function: %w", err)
		}
		to = token.L2Address
		amount = big.NewInt(0)
	}
	tx := CreateFunctionCallTransaction(
		w.es.GetAddress(),
		to,
		big.NewInt(0),
		big.NewInt(0),
		amount,
		data,
	)
	return w.estimateAndSend(tx, nonce)
}

func (w *Wallet) Withdraw(to common.Address, amount *big.Int, token *Token, nonce *big.Int) (string, error) {
	var err error
	if token == nil {
		token = CreateETH()
	}
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return "", fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	var data hexutil.Bytes
	data, err = w.l2BridgeAbi.Pack("withdraw", to, token.L2Address, amount)
	if err != nil {
		return "", fmt.Errorf("failed to pack transfer function: %w", err)
	}
	bcs, err := w.getBridgeContracts()
	if err != nil {
		return "", fmt.Errorf("failed to getBridgeContracts: %w", err)
	}
	l2BridgeAddress := bcs.L2EthDefaultBridge
	if !token.IsETH() {
		l2BridgeAddress = bcs.L2Erc20DefaultBridge
	}
	tx := CreateFunctionCallTransaction(
		w.es.GetAddress(),
		l2BridgeAddress,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		data,
	)
	return w.estimateAndSend(tx, nonce)
}

func (w *Wallet) Deploy(bytecode []byte, nonce *big.Int) (string, error) {
	var err error
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return "", fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	calldata, err := EncodeCreate2(bytecode, nil)
	if err != nil {
		return "", fmt.Errorf("failed to encode create2 call: %w", err)
	}
	tx := Create2ContractTransaction(
		w.es.GetAddress(),
		big.NewInt(0),
		big.NewInt(0),
		bytecode,
		calldata,
	)
	return w.estimateAndSend(tx, nonce)
}

func (w *Wallet) estimateAndSend(tx *Transaction, nonce *big.Int) (string, error) {
	fmt.Println("nonce", nonce)
	txj, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Println("Tx:", string(txj))

	gas, err := w.zp.EstimateGas(tx)
	if err != nil {
		return "", fmt.Errorf("failed to EstimateGas: %w", err)
	}
	fmt.Println("EstimateGas", gas)
	chainId := w.es.GetDomain().ChainId
	fmt.Println("chainId", chainId)

	gasPrice, err := w.zp.GetGasPrice()
	if err != nil {
		return "", fmt.Errorf("failed to GetGasPrice: %w", err)
	}
	fmt.Println("gasPrice", gasPrice)

	prepared := NewTransaction712(
		chainId,
		nonce,
		gas,
		tx.To,
		tx.Value.ToInt(),
		tx.Data,
		big.NewInt(100000000), // TODO: Estimate correct one
		gasPrice,
		tx.From,
		tx.Eip712Meta,
	)
	signature, err := w.es.SignTypedData(w.es.GetDomain(), prepared)
	fmt.Println("signature", hexutil.Encode(signature), err)

	rawTx, err := prepared.RLPValues(signature)
	fmt.Println("rawTx", hexutil.Encode(rawTx), err, len(rawTx))

	txHash, err := w.zp.SendRawTransaction(rawTx)
	fmt.Println("txHash", txHash, err)
	return txHash, err
}

func (w *Wallet) getBridgeContracts() (*BridgeContracts, error) {
	if w.bcs != nil {
		return w.bcs, nil
	}
	var err error
	w.bcs, err = w.zp.ZksGetBridgeContracts()
	if err != nil {
		return nil, err
	}
	return w.bcs, nil
}

func (w *Wallet) newTransactorWithSigner(ethSigner EthSigner, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, bind.ErrNoChainID
	}
	keyAddr := ethSigner.GetAddress()
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}
			signature, err := ethSigner.SignHash(signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}, nil
}
