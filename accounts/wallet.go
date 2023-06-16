package accounts

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
	zkClient "github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/ethtoken"
	"github.com/zksync-sdk/zksync2-go/contracts/l1messenger"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/paymasterflow"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"strings"
)

type Wallet struct {
	es EthSigner
	zp zkClient.Provider
	ep zkClient.EthProvider

	bcs              *zkTypes.BridgeContracts
	mainContract     common.Address
	erc20abi         abi.ABI
	ethTokenAbi      abi.ABI
	l2BridgeAbi      abi.ABI
	l1MessengerAbi   abi.ABI
	paymasterFlowAbi abi.ABI
}

func NewWallet(es EthSigner, zp zkClient.Provider) (*Wallet, error) {
	erc20abi, err := abi.JSON(strings.NewReader(erc20.IERC20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load erc20abi: %w", err)
	}
	ethTokenAbi, err := abi.JSON(strings.NewReader(ethtoken.IEthTokenMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load ethTokenAbi: %w", err)
	}
	l2BridgeAbi, err := abi.JSON(strings.NewReader(l2bridge.IL2BridgeMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l2BridgeAbi: %w", err)
	}
	l1MessengerAbi, err := abi.JSON(strings.NewReader(l1messenger.IL1MessengerMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l1MessengerAbi: %w", err)
	}
	paymasterFlowAbi, err := abi.JSON(strings.NewReader(paymasterflow.IPaymasterFlowMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load paymasterFlowAbi: %w", err)
	}
	return &Wallet{
		es:               es,
		zp:               zp,
		erc20abi:         erc20abi,
		ethTokenAbi:      ethTokenAbi,
		l2BridgeAbi:      l2BridgeAbi,
		l1MessengerAbi:   l1MessengerAbi,
		paymasterFlowAbi: paymasterFlowAbi,
	}, nil
}

func (w *Wallet) GetEthSigner() EthSigner {
	return w.es
}

func (w *Wallet) GetAddress() common.Address {
	return w.es.GetAddress()
}

func (w *Wallet) GetProvider() zkClient.Provider {
	return w.zp
}

func (w *Wallet) CreateEthereumProvider(rpcClient *rpc.Client) (zkClient.EthProvider, error) {
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
	w.ep, err = zkClient.NewDefaultEthProvider(rpcClient, auth, mc, bcs.L1Erc20DefaultBridge)
	if err != nil {
		return nil, fmt.Errorf("failed to init NewDefaultEthProvider: %w", err)
	}
	return w.ep, nil
}

func (w *Wallet) SetEthereumProvider(ep zkClient.EthProvider) {
	w.ep = ep
}

func (w *Wallet) GetEthereumProvider() zkClient.EthProvider {
	return w.ep
}

func (w *Wallet) GetBalance() (*big.Int, error) {
	return w.zp.GetBalance(w.es.GetAddress(), zkTypes.BlockNumberCommitted)
}

func (w *Wallet) GetTokenBalance(token *zkTypes.Token) (*big.Int, error) {
	if token.IsETH() {
		return w.zp.GetBalance(w.es.GetAddress(), zkTypes.BlockNumberCommitted)
	}
	erc20Contract, err := erc20.NewIERC20(token.L2Address, w.zp.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to load IERC20: %w", err)
	}
	return erc20Contract.BalanceOf(nil, w.es.GetAddress())
}

func (w *Wallet) GetBalanceOf(token *zkTypes.Token, at zkTypes.BlockNumber) (*big.Int, error) {
	if token == nil || token.IsETH() {
		return w.zp.GetBalance(w.es.GetAddress(), at)
	}
	erc20Contract, err := erc20.NewIERC20(token.L2Address, w.zp.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to load IERC20: %w", err)
	}
	opts := &bind.CallOpts{}
	if at == zkTypes.BlockNumberPending {
		opts.Pending = true
	} else if bn, err := hexutil.DecodeBig(string(at)); err == nil && bn != nil {
		opts.BlockNumber = bn
	}
	return erc20Contract.BalanceOf(opts, w.es.GetAddress())
}

func (w *Wallet) GetNonce() (*big.Int, error) {
	return w.GetNonceAt(zkTypes.BlockNumberCommitted)
}

func (w *Wallet) GetNonceAt(at zkTypes.BlockNumber) (*big.Int, error) {
	return w.zp.GetTransactionCount(w.es.GetAddress(), at)
}

func (w *Wallet) Transfer(to common.Address, amount *big.Int, token *zkTypes.Token, nonce *big.Int) (common.Hash, error) {
	var err error
	if token == nil {
		token = utils.CreateETH()
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
	tx := utils.CreateFunctionCallTransaction(
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

func (w *Wallet) Withdraw(to common.Address, amount *big.Int, token *zkTypes.Token, nonce *big.Int) (common.Hash, error) {
	var err error
	if token == nil {
		token = utils.CreateETH()
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
		tx := utils.CreateFunctionCallTransaction(
			w.es.GetAddress(),
			utils.L2EthTokenAddress,
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
		tx := utils.CreateFunctionCallTransaction(
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
	if l1BatchTxId == nil {
		return common.Hash{}, errors.New("empty l1BatchTxIndex")
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
	ev, err := w.l1MessengerAbi.EventByID(log.Topics[0])
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get EventByID: %w", err)
	}
	dl, err := w.l1MessengerAbi.Unpack(ev.Name, log.Data)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to Unpack log data: %w", err)
	}
	message := dl[0].([]byte)

	// ETH token
	if sender == utils.L2EthTokenAddress {
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
	l2Bridge, err := l2bridge.NewIL2Bridge(sender, w.zp.GetClient())
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
	if sender == utils.L2EthTokenAddress {
		return w.ep.IsEthWithdrawalFinalized(
			log.L1BatchNumber.ToInt(),
			big.NewInt(int64(proof.Id)),
		)
	}
	// other tokens
	l2Bridge, err := l2bridge.NewIL2Bridge(sender, w.zp.GetClient())
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

func (w *Wallet) ClaimFailedDeposit(depositHash common.Hash, ep zkClient.EthProvider) (common.Hash, error) {
	receipt, err := w.zp.GetTransactionReceipt(depositHash)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return common.Hash{}, errors.New("transaction receipt not found")
	}
	var successL2ToL1Log *zkTypes.L2ToL1Log
	var successL2ToL1LogIndex int
	for i, l := range receipt.L2ToL1Logs {
		if l.Sender == utils.BootloaderFormalAddress && l.Key == depositHash.String() {
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
	if tx == nil {
		return common.Hash{}, errors.New("transaction not found")
	}

	// Undo the aliasing, since the Mailbox contract set it as for contract address.
	l1BridgeAddress := utils.UndoL1ToL2Alias(receipt.From)
	l2Bridge, err := abi.JSON(strings.NewReader(l2bridge.IL2BridgeMetaData.ABI))
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
	data, err := utils.EncodeCreate2(bytecode, calldata, salt)
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
	tx := utils.Create2ContractTransaction(
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

func (w *Wallet) DeployWithCreate(bytecode []byte, calldata []byte, deps [][]byte, nonce *big.Int) (common.Hash, error) {
	var err error
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	data, err := utils.EncodeCreate(bytecode, calldata)
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
	tx := utils.Create2ContractTransaction(
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

func (w *Wallet) DeployAccount(bytecode []byte, calldata []byte, salt []byte, nonce *big.Int) (common.Hash, error) {
	var err error
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	data, err := utils.EncodeCreate2Account(bytecode, calldata, salt, zkTypes.Version1)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to encode create2Account call: %w", err)
	}

	factoryDeps := []hexutil.Bytes{bytecode}
	tx := utils.Create2ContractTransaction(
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

func (w *Wallet) DeployAccountWithCreate(bytecode []byte, calldata []byte, nonce *big.Int) (common.Hash, error) {
	var err error
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	data, err := utils.EncodeCreateAccount(bytecode, calldata, zkTypes.Version1)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to encode create2Account call: %w", err)
	}

	factoryDeps := []hexutil.Bytes{bytecode}
	tx := utils.Create2ContractTransaction(
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

func (w *Wallet) Execute(contract common.Address, calldata []byte, value *big.Int, nonce *big.Int) (common.Hash, error) {
	var err error
	if value == nil {
		value = big.NewInt(0)
	}
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to get nonce: %w", err)
		}
	}
	tx := utils.CreateFunctionCallTransaction(
		w.es.GetAddress(),
		contract,
		big.NewInt(0),
		big.NewInt(0),
		value,
		calldata,
		nil, nil,
	)
	return w.EstimateAndSend(tx, nonce)
}

func (w *Wallet) EstimateAndSend(tx *zkTypes.Transaction, nonce *big.Int) (common.Hash, error) {
	gas, err := w.zp.EstimateGas712(tx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to EstimateGas712: %w", err)
	}
	chainId := w.es.GetDomain().ChainId
	gasPrice, err := w.zp.GetGasPrice()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to GetGasPrice: %w", err)
	}
	prepared := &zkTypes.Transaction712{
		ChainID:   chainId,
		Nonce:     nonce,
		Gas:       gas,
		To:        &tx.To,
		Value:     tx.Value.ToInt(),
		Data:      tx.Data,
		GasTipCap: big.NewInt(100_000_000), // TODO: Estimate correct one
		GasFeeCap: gasPrice,
		From:      &tx.From,
		Meta:      tx.Eip712Meta,
	}
	signature, err := w.es.SignTypedData(w.es.GetDomain(), prepared)
	if err != nil {
		return common.Hash{}, err
	}
	rawTx, err := prepared.RLPValues(signature)
	if err != nil {
		return common.Hash{}, err
	}
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

func (w *Wallet) getBridgeContracts() (*zkTypes.BridgeContracts, error) {
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

func (w *Wallet) getWithdrawalLog(withdrawalHash common.Hash, index int) (*zkTypes.Log, *big.Int, error) {
	receipt, err := w.zp.GetTransactionReceipt(withdrawalHash)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return nil, nil, errors.New("transaction receipt not found")
	}
	fLogs := make([]*zkTypes.Log, 0)
	for _, l := range receipt.Logs {
		if l.Address == utils.MessengerAddress && len(l.Topics) > 0 &&
			bytes.Equal(l.Topics[0].Bytes(), crypto.Keccak256([]byte("L1MessageSent(address,bytes32,bytes)"))) {
			fLogs = append(fLogs, l)
		}
	}
	if len(fLogs) < index+1 {
		return nil, nil, errors.New("withdrawal log not found")
	}
	return fLogs[index], receipt.L1BatchTxIndex.ToInt(), nil
}

func (w *Wallet) getWithdrawalL2ToL1Log(withdrawalHash common.Hash, index int) (int, *zkTypes.L2ToL1Log, error) {
	receipt, err := w.zp.GetTransactionReceipt(withdrawalHash)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return 0, nil, errors.New("transaction receipt not found")
	}
	fLogs := make([]struct {
		i int
		l *zkTypes.L2ToL1Log
	}, 0)
	for i, l := range receipt.L2ToL1Logs {
		if l.Sender == utils.MessengerAddress {
			fLogs = append(fLogs, struct {
				i int
				l *zkTypes.L2ToL1Log
			}{i, l})
		}
	}
	if len(fLogs) < index+1 {
		return 0, nil, errors.New("withdrawal L2ToL1 log not found")
	}
	return fLogs[index].i, fLogs[index].l, nil
}
