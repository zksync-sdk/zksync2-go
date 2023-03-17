package zksync2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/contracts/ERC20"
	"github.com/zksync-sdk/zksync2-go/contracts/IEthToken"
	"github.com/zksync-sdk/zksync2-go/contracts/IL1Messenger"
	"github.com/zksync-sdk/zksync2-go/contracts/IL2Bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/IL2Messenger"
	"math/big"
	"strings"
)

type Wallet struct {
	es EthSigner
	zp Provider
	ep EthProvider

	bcs          *BridgeContracts
	mainContract common.Address
	erc20abi     abi.ABI
	ethTokenAbi  abi.ABI
	l2BridgeAbi  abi.ABI
	l1Messenger  abi.ABI
	l2Messenger  abi.ABI
}

func NewWallet(es EthSigner, zp Provider) (*Wallet, error) {
	erc20abi, err := abi.JSON(strings.NewReader(ERC20.ERC20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load erc20abi: %w", err)
	}
	ethTokenAbi, err := abi.JSON(strings.NewReader(IEthToken.IEthTokenMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load ethTokenAbi: %w", err)
	}
	l2BridgeAbi, err := abi.JSON(strings.NewReader(IL2Bridge.IL2BridgeMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l2BridgeAbi: %w", err)
	}
	l1Messenger, err := abi.JSON(strings.NewReader(IL1Messenger.IL1MessengerMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l1Messenger: %w", err)
	}
	l2Messenger, err := abi.JSON(strings.NewReader(IL2Messenger.IL2MessengerMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l2Messenger: %w", err)
	}
	return &Wallet{
		es:          es,
		zp:          zp,
		erc20abi:    erc20abi,
		ethTokenAbi: ethTokenAbi,
		l2BridgeAbi: l2BridgeAbi,
		l1Messenger: l1Messenger,
		l2Messenger: l2Messenger,
	}, nil
}

func (w *Wallet) GetEthSigner() EthSigner {
	return w.es
}

func (w *Wallet) GetProvider() Provider {
	return w.zp
}

func (w *Wallet) CreateEthereumProvider(rpcClient *rpc.Client) (EthProvider, error) {
	ethClient := ethclient.NewClient(rpcClient)
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain Id: %w", err)
	}
	auth, err := w.newTransactorWithSigner(w.es, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to init TransactOpts: %w", err)
	}
	mc, err := w.getMainContract()
	if err != nil {
		return nil, fmt.Errorf("failed to getMainContract: %w", err)
	}
	bcs, err := w.getBridgeContracts()
	if err != nil {
		return nil, fmt.Errorf("failed to getBridgeContracts: %w", err)
	}
	w.ep, err = NewDefaultEthProvider(rpcClient, auth, mc, bcs.L1Erc20DefaultBridge)
	if err != nil {
		return nil, fmt.Errorf("failed to init NewDefaultEthProvider: %w", err)
	}
	return w.ep, nil
}

func (w *Wallet) SetEthereumProvider(ep EthProvider) {
	w.ep = ep
}

func (w *Wallet) GetEthereumProvider() EthProvider {
	return w.ep
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

func (w *Wallet) Transfer(to common.Address, amount *big.Int, token *Token, nonce *big.Int) (common.Hash, error) {
	var err error
	if token == nil {
		token = CreateETH()
	}
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	var data hexutil.Bytes
	if !token.IsETH() {
		data, err = w.erc20abi.Pack("transfer", to, amount)
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to pack transfer function: %w", err)
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
		nil, nil,
	)
	return w.EstimateAndSend(tx, nonce)
}

func (w *Wallet) Withdraw(to common.Address, amount *big.Int, token *Token, nonce *big.Int) (common.Hash, error) {
	var err error
	if token == nil {
		token = CreateETH()
	}
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	var data hexutil.Bytes
	if token.IsETH() {
		data, err = w.ethTokenAbi.Pack("withdraw", to)
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to pack withdraw function: %w", err)
		}
		tx := CreateFunctionCallTransaction(
			w.es.GetAddress(),
			L2EthTokenAddress,
			big.NewInt(0),
			big.NewInt(0),
			amount,
			data,
			nil, nil,
		)
		return w.EstimateAndSend(tx, nonce)
	} else {
		data, err = w.l2BridgeAbi.Pack("withdraw", to, token.L2Address, amount)
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to pack withdraw function: %w", err)
		}
		bcs, err := w.getBridgeContracts()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to getBridgeContracts: %w", err)
		}
		tx := CreateFunctionCallTransaction(
			w.es.GetAddress(),
			bcs.L2Erc20DefaultBridge,
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			data,
			nil, nil,
		)
		return w.EstimateAndSend(tx, nonce)
	}
}

func (w *Wallet) FinalizeWithdraw(withdrawalHash common.Hash, index int) (common.Hash, error) {
	if w.ep == nil {
		return common.Hash{}, errors.New("ethereum provider is not initialized")
	}
	log, l1BatchTxId, err := w.getWithdrawalLog(withdrawalHash, index)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get WithdrawalLog: %w", err)
	}
	l2ToL1LogIndex, _, err := w.getWithdrawalL2ToL1Log(withdrawalHash, index)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get WithdrawalL2ToL1Log: %w", err)
	}
	if len(log.Topics) < 2 {
		return common.Hash{}, errors.New("not enough Topics count")
	}
	sender := common.BytesToAddress(log.Topics[1].Bytes()[12:])
	proof, err := w.zp.ZksGetL2ToL1LogProof(withdrawalHash, l2ToL1LogIndex)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}
	ev, err := w.l1Messenger.EventByID(log.Topics[0])
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get EventByID: %w", err)
	}
	dl, err := w.l1Messenger.Unpack(ev.Name, log.Data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to Unpack log data: %w", err)
	}
	message := dl[0].([]byte)

	// ETH token
	if sender == L2EthTokenAddress {
		tx, err := w.ep.FinalizeEthWithdrawal(
			log.L1BatchNumber.ToInt(),
			big.NewInt(int64(proof.Id)),
			l1BatchTxId,
			message,
			proof.Proof,
			nil,
		)
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to FinalizeEthWithdrawal: %w", err)
		}
		return tx.Hash(), nil
	}
	// other tokens
	l2Bridge, err := IL2Bridge.NewIL2Bridge(sender, w.zp.GetClient())
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to init l2Bridge: %w", err)
	}
	l1BridgeAddress, err := l2Bridge.L1Bridge(&bind.CallOpts{})
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get l1BridgeAddress: %w", err)
	}
	tx, err := w.ep.FinalizeWithdrawal(
		l1BridgeAddress,
		log.L1BatchNumber.ToInt(),
		big.NewInt(int64(proof.Id)),
		l1BatchTxId,
		message,
		proof.Proof,
		nil,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to FinalizeWithdrawal: %w", err)
	}
	return tx.Hash(), nil
}

func (w *Wallet) IsWithdrawFinalized(withdrawalHash common.Hash, index int) (bool, error) {
	if w.ep == nil {
		return false, errors.New("ethereum provider is not initialized")
	}
	log, _, err := w.getWithdrawalLog(withdrawalHash, index)
	if err != nil {
		return false, fmt.Errorf("failed to get WithdrawalLog: %w", err)
	}
	l2ToL1LogIndex, _, err := w.getWithdrawalL2ToL1Log(withdrawalHash, index)
	if err != nil {
		return false, fmt.Errorf("failed to get WithdrawalL2ToL1Log: %w", err)
	}
	if len(log.Topics) < 2 {
		return false, errors.New("not enough Topics count")
	}
	sender := common.BytesToAddress(log.Topics[1].Bytes()[12:])
	proof, err := w.zp.ZksGetL2ToL1LogProof(withdrawalHash, l2ToL1LogIndex)
	if err != nil {
		return false, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}
	// ETH token
	if sender == L2EthTokenAddress {
		return w.ep.IsEthWithdrawalFinalized(
			log.L1BatchNumber.ToInt(),
			big.NewInt(int64(proof.Id)),
		)
	}
	// other tokens
	l2Bridge, err := IL2Bridge.NewIL2Bridge(sender, w.zp.GetClient())
	if err != nil {
		return false, fmt.Errorf("failed to init l2Bridge: %w", err)
	}
	l1BridgeAddress, err := l2Bridge.L1Bridge(&bind.CallOpts{})
	if err != nil {
		return false, fmt.Errorf("failed to get l1BridgeAddress: %w", err)
	}
	return w.ep.IsWithdrawalFinalized(
		l1BridgeAddress,
		log.L1BatchNumber.ToInt(),
		big.NewInt(int64(proof.Id)),
	)
}

func (w *Wallet) ClaimFailedDeposit(depositHash common.Hash, ep EthProvider) (common.Hash, error) {
	receipt, err := w.zp.GetTransactionReceipt(depositHash)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	var successL2ToL1Log *L2ToL1Log
	var successL2ToL1LogIndex int
	for i, l := range receipt.L2ToL1Logs {
		if l.Sender == BootloaderFormalAddress && l.Key == depositHash.String() {
			successL2ToL1LogIndex = i
			successL2ToL1Log = l
		}
	}
	if successL2ToL1Log.Value != (common.Hash{}).String() {
		return common.Hash{}, errors.New("can't claim successful deposit")
	}

	tx, err := w.zp.GetTransaction(depositHash)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get Transaction: %w", err)
	}

	// Undo the aliasing, since the Mailbox contract set it as for contract address.
	l1BridgeAddress := UndoL1ToL2Alias(receipt.From)
	l2Bridge, err := abi.JSON(strings.NewReader(IL2Bridge.IL2BridgeMetaData.ABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to load l2Bridge ABI: %w", err)
	}
	calldata, err := l2Bridge.Unpack("finalizeDeposit", tx.Data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to Unpack finalizeDeposit data: %w", err)
	}
	if len(calldata) < 3 {
		return common.Hash{}, errors.New("unpacked calldata is empty")
	}
	b20, ok := calldata[0].([20]byte)
	if !ok {
		return common.Hash{}, errors.New("failed to parse l1Sender from unpacked calldata")
	}
	l1Sender := common.BytesToAddress(b20[:])
	b20, ok = calldata[2].([20]byte)
	if !ok {
		return common.Hash{}, errors.New("failed to parse l1Token from unpacked calldata")
	}
	l1Token := common.BytesToAddress(b20[:])

	proof, err := w.zp.ZksGetL2ToL1LogProof(depositHash, successL2ToL1LogIndex)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}

	res, err := ep.ClaimFailedDeposit(
		l1BridgeAddress,
		l1Sender,
		l1Token,
		depositHash,
		receipt.L1BatchNumber.ToInt(),
		big.NewInt(int64(proof.Id)),
		receipt.L1BatchTxIndex.ToInt(),
		proof.Proof,
		nil,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to call ClaimFailedDeposit: %w", err)
	}
	return res.Hash(), nil
}

func (w *Wallet) Deploy(bytecode []byte, calldata []byte, salt []byte, deps [][]byte, nonce *big.Int) (common.Hash, error) {
	var err error
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	data, err := EncodeCreate2(bytecode, calldata, salt)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to encode create2 call: %w", err)
	}
	var factoryDeps []hexutil.Bytes
	if len(deps) > 0 {
		factoryDeps = make([]hexutil.Bytes, len(deps))
		for i, d := range deps {
			factoryDeps[i] = d
		}
	}
	tx := Create2ContractTransaction(
		w.es.GetAddress(),
		big.NewInt(0),
		big.NewInt(0),
		bytecode,
		data,
		factoryDeps,
		nil, nil,
	)
	return w.EstimateAndSend(tx, nonce)
}

func (w *Wallet) Execute(contract common.Address, calldata []byte, nonce *big.Int) (common.Hash, error) {
	var err error
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	tx := CreateFunctionCallTransaction(
		w.es.GetAddress(),
		contract,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		calldata,
		nil, nil,
	)
	return w.EstimateAndSend(tx, nonce)
}

func (w *Wallet) SendMessageToL1(message []byte, nonce *big.Int) (common.Hash, error) {
	var err error
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	var data hexutil.Bytes
	data, err = w.l2Messenger.Pack("sendToL1", message)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack sendToL1 function: %w", err)
	}
	tx := CreateFunctionCallTransaction(
		w.es.GetAddress(),
		MessengerAddress,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		data,
		nil, nil,
	)
	return w.EstimateAndSend(tx, nonce)
}

func (w *Wallet) EstimateAndSend(tx *Transaction, nonce *big.Int) (common.Hash, error) {
	gas, err := w.zp.EstimateGas(tx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to EstimateGas: %w", err)
	}
	chainId := w.es.GetDomain().ChainId
	gasPrice, err := w.zp.GetGasPrice()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to GetGasPrice: %w", err)
	}
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
	rawTx, err := prepared.RLPValues(signature)
	return w.zp.SendRawTransaction(rawTx)
}

func (w *Wallet) getMainContract() (common.Address, error) {
	if !bytes.Equal(w.mainContract.Bytes(), common.Address{}.Bytes()) {
		return w.mainContract, nil
	}
	var err error
	w.mainContract, err = w.zp.ZksGetMainContract()
	if err != nil {
		return common.Address{}, err
	}
	return w.mainContract, nil
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

func (w *Wallet) getWithdrawalLog(withdrawalHash common.Hash, index int) (*Log, *big.Int, error) {
	receipt, err := w.zp.GetTransactionReceipt(withdrawalHash)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	fLogs := make([]*Log, 0)
	for _, l := range receipt.Logs {
		if l.Address == MessengerAddress && len(l.Topics) > 0 &&
			bytes.Equal(l.Topics[0].Bytes(), crypto.Keccak256([]byte("L1MessageSent(address,bytes32,bytes)"))) {
			fLogs = append(fLogs, l)
		}
	}
	if len(fLogs) < index+1 {
		return nil, nil, errors.New("withdrawal log not found")
	}
	return fLogs[index], receipt.L1BatchTxIndex.ToInt(), nil
}

func (w *Wallet) getWithdrawalL2ToL1Log(withdrawalHash common.Hash, index int) (int, *L2ToL1Log, error) {
	receipt, err := w.zp.GetTransactionReceipt(withdrawalHash)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	fLogs := make([]struct {
		i int
		l *L2ToL1Log
	}, 0)
	for i, l := range receipt.L2ToL1Logs {
		if l.Sender == MessengerAddress {
			fLogs = append(fLogs, struct {
				i int
				l *L2ToL1Log
			}{i, l})
		}
	}
	if len(fLogs) < index+1 {
		return 0, nil, errors.New("withdrawal L2ToL1 log not found")
	}
	return fLogs[index].i, fLogs[index].l, nil
}
