package accounts

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/l1bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l1messenger"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/zksync"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"strings"
)

// WalletL1 implements the AdapterL1 interface.
type WalletL1 struct {
	clientL1 *ethclient.Client
	clientL2 *clients.Client
	auth     *bind.TransactOpts

	mainContractAddress common.Address
	mainContract        *zksync.IZkSync

	defaultL1BridgeAddress common.Address
	defaultL1Bridge        *l1bridge.IL1Bridge
}

// NewWalletL1 creates an instance of WalletL1 associated with the account provided by the raw private key.
func NewWalletL1(rawPrivateKey []byte, clientL1 *ethclient.Client, clientL2 *clients.Client) (*WalletL1, error) {
	chainID, err := (*clientL2).ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	signer, err := NewBaseSignerFromRawPrivateKey(rawPrivateKey, chainID.Int64())
	if err != nil {
		return nil, err
	}
	s := Signer(signer)
	return NewWalletL1FromSigner(&s, clientL1, clientL2)
}

// NewWalletL1FromSigner creates an instance of WalletL1 associated with the account provided by the signer.
func NewWalletL1FromSigner(signer *Signer, clientL1 *ethclient.Client, clientL2 *clients.Client) (*WalletL1, error) {
	if signer == nil {
		return nil, errors.New("signer is not provided")
	} else if clientL1 == nil {
		return nil, errors.New("clientL1 is not provided")
	} else if clientL2 == nil {
		return nil, errors.New("clientL2 is not provided")
	}

	mainContractAddress, err := (*clientL2).MainContractAddress(context.Background())
	if err != nil {
		return nil, err
	}
	iZkSync, err := zksync.NewIZkSync(mainContractAddress, clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IZkSync: %w", err)
	}
	bridgeContracts, err := (*clientL2).BridgeContracts(context.Background())
	if err != nil {
		return nil, err
	}
	iL1Bridge, err := l1bridge.NewIL1Bridge(bridgeContracts.L1Erc20DefaultBridge, clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IL1Bridge: %w", err)
	}
	chainId, err := clientL1.ChainID(context.Background())
	auth, err := newTransactorWithSigner(signer, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to init TransactOpts: %w", err)
	}

	return &WalletL1{
		clientL1:               clientL1,
		clientL2:               clientL2,
		auth:                   auth,
		mainContractAddress:    mainContractAddress,
		defaultL1BridgeAddress: bridgeContracts.L1Erc20DefaultBridge,
		defaultL1Bridge:        iL1Bridge,
		mainContract:           iZkSync,
	}, nil
}

func (a *WalletL1) MainContract(_ context.Context) (*zksync.IZkSync, error) {
	return a.mainContract, nil
}

func (a *WalletL1) L1BridgeContracts(_ context.Context) (*zkTypes.L1BridgeContracts, error) {
	return &zkTypes.L1BridgeContracts{Erc20: a.defaultL1Bridge}, nil
}

func (a *WalletL1) BalanceL1(opts *CallOpts, token common.Address) (*big.Int, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(a.auth.From)
	if token == utils.EthAddress {
		return a.clientL1.BalanceAt(callOpts.Context, a.auth.From, callOpts.BlockNumber)
	} else {
		erc20Contract, err := erc20.NewIERC20(token, a.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load IERC20: %w", err)
		}
		return erc20Contract.BalanceOf(callOpts, a.auth.From)
	}
}

func (a *WalletL1) AllowanceL1(opts *CallOpts, token common.Address, bridgeAddress common.Address) (*big.Int, error) {
	if token == (common.Address{}) {
		token = utils.EthAddress
	}
	erc20Contract, err := erc20.NewIERC20(token, a.clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IERC20: %w", err)
	}
	return erc20Contract.Allowance(ensureCallOpts(opts).ToCallOpts(a.auth.From), a.auth.From, bridgeAddress)
}

func (a *WalletL1) L2TokenAddress(ctx context.Context, token common.Address) (common.Address, error) {
	return (*a.clientL2).L2TokenAddress(ensureContext(ctx), token)
}

func (a *WalletL1) ApproveERC20(auth *TransactOpts, token common.Address, amount *big.Int,
	bridgeAddress common.Address) (*types.Transaction, error) {
	if token == (common.Address{}) {
		return nil, errors.New("token L1 address must be provided")
	}
	if token == utils.EthAddress {
		return nil, errors.New("ETH token can't be approved. The address of the token does not exist on L1")
	}
	if bridgeAddress == (common.Address{}) {
		bridgeAddress = a.defaultL1BridgeAddress
	}

	erc20Contract, err := erc20.NewIERC20(token, a.clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IERC20: %w", err)
	}
	return erc20Contract.Approve(ensureTransactOpts(auth).ToTransactOpts(a.auth.From, a.auth.Signer), bridgeAddress, amount)
}

func (a *WalletL1) BaseCost(opts *CallOpts, gasLimit, gasPerPubdataByte, gasPrice *big.Int) (*big.Int, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(a.auth.From)
	if gasPrice == nil {
		var err error
		if gasPrice, err = a.clientL1.SuggestGasPrice(callOpts.Context); err != nil {
			return nil, err
		}
	}
	return a.mainContract.L2TransactionBaseCost(callOpts,
		gasPrice,
		gasLimit,
		gasPerPubdataByte,
	)
}

func (a *WalletL1) Deposit(auth *TransactOpts, tx DepositTransaction) (*types.Transaction, error) {
	opts, depositTx, err := a.prepareDepositTx(*ensureTransactOpts(auth), tx)
	if err != nil {
		return nil, err
	}

	if depositTx.Token == utils.EthAddress {
		return a.depositETH(opts, depositTx)
	} else {
		if depositTx.ApproveERC20 {
			errApprove := a.approveERC20(opts, depositTx)
			if errApprove != nil {
				return nil, errApprove
			}
		}
		return a.depositERC20(opts, depositTx)
	}
}

func (a *WalletL1) EstimateGasDeposit(ctx context.Context, msg DepositCallMsg) (uint64, error) {
	auth, prepareDepositTx, err := a.prepareDepositTx(msg.ToTransactOpts(), msg.ToDepositTransaction())
	if err != nil {
		return 0, err
	}

	preparedDepositCallMsg := prepareDepositTx.ToDepositCallMsg(auth)
	if preparedDepositCallMsg.Token == utils.EthAddress {
		return a.EstimateGasRequestExecute(ensureContext(ctx), preparedDepositCallMsg.ToRequestExecuteCallMsg())
	} else {
		return a.estimateDepositERC20(ctx, preparedDepositCallMsg)
	}
}

func (a *WalletL1) FullRequiredDepositFee(ctx context.Context, msg DepositCallMsg) (*FullDepositFee, error) {
	// It is assumed that the L2 fee for the transaction does not depend on its value.
	dummyAmount := big.NewInt(1)
	msg.PopulateEmptyFields(a.auth.From)

	if msg.BridgeAddress != nil {
		if msg.Token == utils.EthAddress {
			return nil, errors.New("ETH token can not be deposited with custom bridge")
		}
		bridge, err := l1bridge.NewIL1Bridge(*msg.BridgeAddress, a.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load custom bridge: %w", err)
		}
		if msg.CustomBridgeData == nil {
			msg.CustomBridgeData, err = utils.Erc20DefaultBridgeData(msg.Token, a.clientL1)
			if err != nil {
				return nil, err
			}
		}

		if msg.L2GasLimit == nil {
			l2Address, errBridge := bridge.L2Bridge(nil)
			if errBridge != nil {
				return nil, errBridge
			}

			gas, errGas := a.EstimateCustomBridgeDepositL2Gas(ensureContext(ctx), *msg.BridgeAddress, l2Address,
				msg.Token, dummyAmount, msg.To, msg.CustomBridgeData, a.auth.From, msg.GasPerPubdataByte)
			if errGas != nil {
				return nil, errGas
			}
			msg.L2GasLimit = new(big.Int).SetUint64(gas)
		}

	} else {
		gas, err := a.EstimateDefaultBridgeDepositL2Gas(ensureContext(ctx), msg.Token, dummyAmount, msg.To,
			a.auth.From, msg.GasPerPubdataByte)
		if err != nil {
			return nil, err
		}
		msg.L2GasLimit = new(big.Int).SetUint64(gas)
	}

	if err := a.insertGasPriceInDepositMsg(ensureContext(ctx), &msg); err != nil {
		return nil, err
	}

	var gasPriceForEstimation *big.Int
	if msg.GasFeeCap != nil {
		gasPriceForEstimation = msg.GasFeeCap
	} else {
		gasPriceForEstimation = msg.GasPrice
	}

	baseCost, err := a.BaseCost(&CallOpts{
		Context: ensureContext(ctx),
	}, msg.L2GasLimit, msg.GasPerPubdataByte, gasPriceForEstimation)
	if err != nil {
		return nil, err
	}

	selfBalanceETH, err := a.BalanceL1(nil, utils.EthAddress)
	if err != nil {
		return nil, err
	}
	// We could use 0, because the final fee will anyway be bigger than
	if baseCost.Cmp(new(big.Int).Add(selfBalanceETH, dummyAmount)) >= 0 {
		recommendedETHBalance := utils.L1RecommendedMinEthDepositGasLimit
		if msg.Token != utils.EthAddress {
			recommendedETHBalance = utils.L1RecommendedMinErc20DepositGasLimit
		}
		recommendedETHBalance.Mul(recommendedETHBalance, gasPriceForEstimation)
		recommendedETHBalance.Add(recommendedETHBalance, baseCost)
		return nil, fmt.Errorf("not enough balance for deposit. Under the provided gas price, "+
			"the recommended balance to perform a deposit is %w ETH", recommendedETHBalance)
	}

	// For ETH token the value that the user passes to the estimation is the one which has the
	// value for the L2 commission subtracted.
	amountForEstimation := big.NewInt(dummyAmount.Int64())
	if msg.Token != utils.EthAddress {
		allowance, errAllowance := a.AllowanceL1(nil, msg.Token, *msg.BridgeAddress)
		if errAllowance != nil {
			return nil, errAllowance
		}
		if allowance.Cmp(amountForEstimation) < 0 {
			return nil, errors.New("not enough allowance to cover the deposit")
		}
	}

	// Deleting the explicit gas limits in the fee estimation
	// in order to prevent the situation where the transaction
	// fails because the user does not have enough balance.
	// In that matter, new DepositCalMsg is created with empty
	// GasPrice, GasFeeCap and GasTipCap field.
	l1GasLimit, err := a.EstimateGasDeposit(ensureContext(ctx), DepositCallMsg{
		To:                msg.To,
		Token:             msg.Token,
		Amount:            amountForEstimation,
		OperatorTip:       msg.OperatorTip,
		BridgeAddress:     msg.BridgeAddress,
		L2GasLimit:        msg.L2GasLimit,
		GasPerPubdataByte: msg.GasPerPubdataByte,
		RefundRecipient:   msg.RefundRecipient,
		CustomBridgeData:  msg.CustomBridgeData,
		Value:             msg.Value,
	})
	if err != nil {
		return nil, err
	}

	fullConst := &FullDepositFee{
		BaseCost:   baseCost,
		L1GasLimit: new(big.Int).SetUint64(l1GasLimit),
		L2GasLimit: msg.L2GasLimit,
	}

	if msg.GasPrice != nil {
		fullConst.GasPrice = msg.GasPrice
	} else {
		fullConst.MaxPriorityFeePerGas = msg.GasTipCap
		fullConst.MaxFeePerGas = msg.GasFeeCap
	}
	return fullConst, nil
}

func (a *WalletL1) FinalizeWithdraw(auth *TransactOpts, withdrawalHash common.Hash, index int) (*types.Transaction, error) {
	if a.clientL1 == nil {
		return nil, errors.New("ethereum provider is not initialized")
	}
	var opts *bind.TransactOpts
	if auth == nil {
		opts = &bind.TransactOpts{
			From:    a.auth.From,
			Signer:  a.auth.Signer,
			Context: context.Background(),
		}
	} else {
		opts = auth.ToTransactOpts(a.auth.From, a.auth.Signer)
	}

	log, l1BatchTxId, err := a.getWithdrawalLog(opts.Context, withdrawalHash, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get WithdrawalLog: %w", err)
	}
	if l1BatchTxId == nil {
		return nil, errors.New("empty l1BatchTxIndex")
	}
	l2ToL1LogIndex, _, err := a.getWithdrawalL2ToL1Log(opts.Context, withdrawalHash, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get WithdrawalL2ToL1Log: %w", err)
	}
	if len(log.Topics) < 2 {
		return nil, errors.New("not enough Topics count")
	}
	sender := common.BytesToAddress(log.Topics[1].Bytes()[12:])
	proof, err := (*a.clientL2).LogProof(opts.Context, withdrawalHash, l2ToL1LogIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}

	l1MessengerAbi, err := abi.JSON(strings.NewReader(l1messenger.IL1MessengerMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l1MessengerAbi: %w", err)
	}
	ev, err := l1MessengerAbi.EventByID(log.Topics[0])
	if err != nil {
		return nil, fmt.Errorf("failed to get EventByID: %w", err)
	}
	dl, err := l1MessengerAbi.Unpack(ev.Name, log.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to Unpack log data: %w", err)
	}
	message := dl[0].([]byte)

	proof32 := make([][32]byte, len(proof.Proof))
	for i, pr := range proof.Proof {
		proof32[i] = pr
	}

	// ETH token
	if sender == utils.L2EthTokenAddress {
		return a.mainContract.FinalizeEthWithdrawal(opts,
			log.L1BatchNumber.ToInt(),
			big.NewInt(int64(proof.Id)),
			uint16(l1BatchTxId.Uint64()),
			message,
			proof32,
		)
	}
	// other tokens
	l2Bridge, err := l2bridge.NewIL2Bridge(sender, *a.clientL2)
	if err != nil {
		return nil, fmt.Errorf("failed to init l2Bridge: %w", err)
	}
	l1BridgeAddress, err := l2Bridge.L1Bridge(&bind.CallOpts{Context: opts.Context})
	if err != nil {
		return nil, fmt.Errorf("failed to get l1BridgeAddress: %w", err)
	}
	l1Bridge, err := l1bridge.NewIL1Bridge(l1BridgeAddress, a.clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to init l1Bridge: %w", err)
	}

	return l1Bridge.FinalizeWithdrawal(opts,
		log.L1BatchNumber.ToInt(),
		big.NewInt(int64(proof.Id)),
		uint16(l1BatchTxId.Uint64()),
		message,
		proof32,
	)
}

func (a *WalletL1) IsWithdrawFinalized(opts *CallOpts, withdrawalHash common.Hash, index int) (bool, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(a.auth.From)
	if a.clientL1 == nil {
		return false, errors.New("ethereum provider is not initialized")
	}
	log, _, err := a.getWithdrawalLog(callOpts.Context, withdrawalHash, index)
	if err != nil {
		return false, fmt.Errorf("failed to get WithdrawalLog: %w", err)
	}
	l2ToL1LogIndex, _, err := a.getWithdrawalL2ToL1Log(callOpts.Context, withdrawalHash, index)
	if err != nil {
		return false, fmt.Errorf("failed to get WithdrawalL2ToL1Log: %w", err)
	}
	if len(log.Topics) < 2 {
		return false, errors.New("not enough Topics count")
	}
	sender := common.BytesToAddress(log.Topics[1].Bytes()[12:])
	proof, err := (*a.clientL2).LogProof(callOpts.Context, withdrawalHash, l2ToL1LogIndex)
	if err != nil {
		return false, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}
	// ETH token
	if sender == utils.L2EthTokenAddress {
		return a.mainContract.IsEthWithdrawalFinalized(callOpts, log.L1BatchNumber.ToInt(), big.NewInt(int64(proof.Id)))
	}
	// other tokens
	l2Bridge, err := l2bridge.NewIL2Bridge(sender, *a.clientL2)
	if err != nil {
		return false, fmt.Errorf("failed to init l2Bridge: %w", err)
	}
	l1BridgeAddress, err := l2Bridge.L1Bridge(callOpts)
	if err != nil {
		return false, fmt.Errorf("failed to get l1BridgeAddress: %w", err)
	}
	l1Bridge, err := l1bridge.NewIL1Bridge(l1BridgeAddress, a.clientL1)
	if err != nil {
		return false, fmt.Errorf("failed to init l1Bridge: %w", err)
	}
	return l1Bridge.IsWithdrawalFinalized(callOpts, log.L1BatchNumber.ToInt(), big.NewInt(int64(proof.Id)))
}

func (a *WalletL1) ClaimFailedDeposit(auth *TransactOpts, depositHash common.Hash) (*types.Transaction, error) {
	opts := ensureTransactOpts(auth)
	receipt, err := (*a.clientL2).TransactionReceipt(opts.Context, depositHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return nil, errors.New("transaction receipt not found")
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
		return nil, errors.New("can't claim successful deposit")
	}

	tx, _, err := (*a.clientL2).TransactionByHash(opts.Context, depositHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get Transaction: %w", err)
	}
	if tx == nil {
		return nil, errors.New("transaction not found")
	}

	// Undo the aliasing, since the Mailbox contract set it as for contract address.
	l1BridgeAddress := utils.UndoL1ToL2Alias(receipt.From)
	l1Bridge, err := l1bridge.NewIL1Bridge(l1BridgeAddress, a.clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to init l1Bridge: %w", err)
	}
	l2Bridge, err := abi.JSON(strings.NewReader(l2bridge.IL2BridgeMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l2Bridge ABI: %w", err)
	}
	calldata, err := l2Bridge.Unpack("finalizeDeposit", tx.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to Unpack finalizeDeposit data: %w", err)
	}
	if len(calldata) < 3 {
		return nil, errors.New("unpacked calldata is empty")
	}
	b20, ok := calldata[0].([20]byte)
	if !ok {
		return nil, errors.New("failed to parse l1Sender from unpacked calldata")
	}
	l1Sender := common.BytesToAddress(b20[:])
	b20, ok = calldata[2].([20]byte)
	if !ok {
		return nil, errors.New("failed to parse l1Token from unpacked calldata")
	}
	l1Token := common.BytesToAddress(b20[:])

	proof, err := (*a.clientL2).LogProof(opts.Context, depositHash, successL2ToL1LogIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}

	proof32 := make([][32]byte, len(proof.Proof))
	for i, pr := range proof.Proof {
		proof32[i] = pr
	}

	return l1Bridge.ClaimFailedDeposit(
		opts.ToTransactOpts(a.auth.From, a.auth.Signer),
		l1Sender,
		l1Token,
		depositHash,
		receipt.L1BatchNumber.ToInt(),
		big.NewInt(int64(proof.Id)),
		uint16(receipt.L1BatchTxIndex.ToInt().Uint64()),
		proof32,
	)
}

func (a *WalletL1) RequestExecute(auth *TransactOpts, tx RequestExecuteTransaction) (*types.Transaction, error) {
	opts, requestExecuteTx, err := a.prepareRequestExecuteTx(*ensureTransactOpts(auth), tx)
	if err != nil {
		return nil, err
	}
	return a.mainContract.RequestL2Transaction(
		opts.ToTransactOpts(a.auth.From, a.auth.Signer),
		requestExecuteTx.ContractAddress,
		requestExecuteTx.L2Value,
		requestExecuteTx.Calldata,
		requestExecuteTx.L2GasLimit,
		requestExecuteTx.GasPerPubdataByte,
		requestExecuteTx.FactoryDeps,
		requestExecuteTx.RefundRecipient,
	)
}

func (a *WalletL1) EstimateGasRequestExecute(ctx context.Context, msg RequestExecuteCallMsg) (uint64, error) {
	opts, prepareRequestExecuteTx, err := a.prepareRequestExecuteTx(msg.ToTransactOpts(), msg.ToRequestExecuteTransaction())
	if err != nil {
		return 0, err
	}

	preparedRequestExecuteCallMsg := prepareRequestExecuteTx.ToRequestExecuteCallMsg(opts)
	callMsg, err := preparedRequestExecuteCallMsg.ToCallMsg(a.auth.From)
	if err != nil {
		return 0, err
	}
	return a.clientL1.EstimateGas(ensureContext(ctx), callMsg)
}

// EstimateCustomBridgeDepositL2Gas used by EstimateDefaultBridgeDepositL2Gas to estimate L2 gas required for token
// bridging via a custom ERC20 bridge.
func (a *WalletL1) EstimateCustomBridgeDepositL2Gas(ctx context.Context, l1BridgeAddress, l2BridgeAddress, token common.Address,
	amount *big.Int, to common.Address, bridgeData []byte, from common.Address, gasPerPubdataByte *big.Int) (uint64, error) {
	calldata, err := utils.Erc20BridgeCalldata(token, from, to, amount, bridgeData)
	if err != nil {
		return 0, err
	}

	return (*a.clientL2).EstimateL1ToL2Execute(ensureContext(ctx), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From: utils.ApplyL1ToL2Alias(l1BridgeAddress),
			To:   &l2BridgeAddress,
			Data: calldata,
		},
		Meta: &zkTypes.Eip712Meta{
			GasPerPubdata: utils.NewBig(gasPerPubdataByte.Int64()),
		},
	})
}

// EstimateDefaultBridgeDepositL2Gas returns an estimation of L2 gas required for token bridging via the default ERC20
// bridge.
func (a *WalletL1) EstimateDefaultBridgeDepositL2Gas(ctx context.Context, token common.Address, amount *big.Int,
	to, from common.Address, gasPerPubdataByte *big.Int) (uint64, error) {

	// If the `from` address is not provided, we use a random address, because
	// due to storage slot aggregation, the gas estimation will depend on the address
	// and so estimation for the zero address may be smaller than for the sender.
	if from == (common.Address{}) {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return 0, fmt.Errorf("failed to generate radnom private key: %w", err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return 0, fmt.Errorf("failed to convert public key to ECDSA")
		}
		from = crypto.PubkeyToAddress(*publicKeyECDSA)
	}

	if token == utils.EthAddress {
		return (*a.clientL2).EstimateL1ToL2Execute(ensureContext(ctx), zkTypes.CallMsg{
			CallMsg: ethereum.CallMsg{
				From:  from,
				To:    &to,
				Value: amount,
			},
			Meta: &zkTypes.Eip712Meta{
				GasPerPubdata: utils.NewBig(gasPerPubdataByte.Int64()),
			},
		})
	} else {
		bridgeContracts, err := (*a.clientL2).BridgeContracts(ensureContext(ctx))
		if err != nil {
			return 0, err
		}
		calldata, err := utils.Erc20DefaultBridgeData(token, a.clientL1)
		if err != nil {
			return 0, err
		}

		return a.EstimateCustomBridgeDepositL2Gas(
			ensureContext(ctx),
			bridgeContracts.L1Erc20DefaultBridge,
			bridgeContracts.L2Erc20DefaultBridge,
			token,
			amount,
			to,
			calldata,
			from,
			gasPerPubdataByte)
	}
}

func (a *WalletL1) prepareDepositTx(auth TransactOpts, tx DepositTransaction) (*TransactOpts, *DepositTransaction, error) {
	opts := ensureTransactOpts(&auth)
	tx.PopulateEmptyFields(a.auth.From)

	if tx.BridgeAddress != nil {
		if tx.Token == utils.EthAddress {
			return nil, nil, errors.New("ETH token can not be deposited with custom bridge")
		}
		bridge, err := l1bridge.NewIL1Bridge(*tx.BridgeAddress, a.clientL1)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to load custom bridge: %w", err)
		}
		if tx.CustomBridgeData == nil {
			tx.CustomBridgeData, err = utils.Erc20DefaultBridgeData(tx.Token, a.clientL1)
			if err != nil {
				return nil, nil, err
			}
		}

		if tx.L2GasLimit == nil {
			l2Address, errBridge := bridge.L2Bridge(nil)
			if errBridge != nil {
				return nil, nil, errBridge
			}

			gas, errGas := a.EstimateCustomBridgeDepositL2Gas(opts.Context, *tx.BridgeAddress, l2Address, tx.Token,
				tx.Amount, tx.To, tx.CustomBridgeData, a.auth.From, tx.GasPerPubdataByte)
			if errGas != nil {
				return nil, nil, errGas
			}
			tx.L2GasLimit = new(big.Int).SetUint64(gas)
		}

	} else {
		gas, err := a.EstimateDefaultBridgeDepositL2Gas(opts.Context, tx.Token, tx.Amount, tx.To,
			a.auth.From, tx.GasPerPubdataByte)
		if err != nil {
			return nil, nil, err
		}
		tx.L2GasLimit = new(big.Int).SetUint64(gas)
	}

	if err := a.insertGasPriceInTransactOpts(&opts); err != nil {
		return nil, nil, err
	}

	var gasPriceForEstimation *big.Int
	if opts.GasFeeCap != nil {
		gasPriceForEstimation = opts.GasFeeCap
	} else {
		gasPriceForEstimation = opts.GasPrice
	}

	baseCost, err := a.BaseCost(&CallOpts{
		Context: opts.Context,
	}, tx.L2GasLimit, tx.GasPerPubdataByte, gasPriceForEstimation)
	if err != nil {
		return nil, nil, err
	}

	if tx.Token == utils.EthAddress {
		if opts.Value == nil {
			opts.Value = new(big.Int).Add(baseCost, new(big.Int).Add(tx.OperatorTip, tx.Amount))
		}
		if tx.RefundRecipient == (common.Address{}) {
			tx.RefundRecipient = tx.To
		}
		return opts, &tx, nil
	} else {
		opts.Value = new(big.Int).Add(baseCost, tx.OperatorTip)
		checkErr := utils.CheckBaseCost(baseCost, opts.Value)
		if checkErr != nil {
			return opts, nil, checkErr
		}
		return opts, &tx, nil
	}
}

func (a *WalletL1) depositETH(auth *TransactOpts, tx *DepositTransaction) (*types.Transaction, error) {
	return a.RequestExecute(auth, RequestExecuteTransaction{
		ContractAddress:   tx.To,
		L2Value:           tx.Amount,
		L2GasLimit:        tx.L2GasLimit,
		GasPerPubdataByte: tx.GasPerPubdataByte,
		OperatorTip:       tx.OperatorTip,
		RefundRecipient:   tx.RefundRecipient,
	})
}

func (a *WalletL1) depositERC20(auth *TransactOpts, tx *DepositTransaction) (*types.Transaction, error) {
	if tx.BridgeAddress != nil {
		l1Bridge, err := l1bridge.NewIL1Bridge(*tx.BridgeAddress, a.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load IL1Bridge: %w", err)
		}
		return l1Bridge.Deposit(
			auth.ToTransactOpts(a.auth.From, a.auth.Signer),
			tx.To,
			tx.Token,
			tx.Amount,
			tx.L2GasLimit,
			tx.GasPerPubdataByte,
			tx.RefundRecipient,
		)
	} else {
		return a.defaultL1Bridge.Deposit(
			auth.ToTransactOpts(a.auth.From, a.auth.Signer),
			tx.To,
			tx.Token,
			tx.Amount,
			tx.L2GasLimit,
			tx.GasPerPubdataByte,
			tx.RefundRecipient,
		)
	}
}

func (a *WalletL1) approveERC20(auth *TransactOpts, tx *DepositTransaction) error {
	// We only request the allowance if the current one is not enough.
	bridge := a.defaultL1BridgeAddress
	if tx.BridgeAddress != nil {
		bridge = *tx.BridgeAddress
	}
	allowance, err := a.AllowanceL1(&CallOpts{
		Context: auth.Context,
	}, tx.Token, bridge)
	if err != nil {
		return err
	}
	if allowance.Cmp(tx.Amount) < 0 {
		approveTx, errApprove := a.ApproveERC20(
			tx.ApproveAuth, tx.Token, tx.Amount, bridge,
		)
		if errApprove != nil {
			return errApprove
		}
		_, err = bind.WaitMined(ensureContext(tx.ApproveAuth.Context), a.clientL1, approveTx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *WalletL1) estimateDepositERC20(ctx context.Context, msg DepositCallMsg) (uint64, error) {
	var bridge common.Address
	if msg.BridgeAddress == nil {
		bridge = a.defaultL1BridgeAddress
	} else {
		bridge = *msg.BridgeAddress
	}
	callMsg, err := msg.ToCallMsg(a.auth.From, bridge)
	if err != nil {
		return 0, err
	}
	return a.clientL1.EstimateGas(ensureContext(ctx), callMsg)
}

func (a *WalletL1) prepareRequestExecuteTx(auth TransactOpts, tx RequestExecuteTransaction) (*TransactOpts, *RequestExecuteTransaction, error) {
	opts := ensureTransactOpts(&auth)
	if tx.L2Value == nil {
		tx.L2Value = big.NewInt(0)
	}
	if tx.OperatorTip == nil {
		tx.OperatorTip = big.NewInt(0)
	}
	if tx.GasPerPubdataByte == nil {
		tx.GasPerPubdataByte = utils.RequiredL1ToL2GasPerPubdataLimit
	}
	if tx.RefundRecipient == (common.Address{}) {
		tx.RefundRecipient = a.auth.From
	}
	if tx.L2GasLimit == nil {
		gas, err := (*a.clientL2).EstimateL1ToL2Execute(opts.Context, tx.ToCallMsg(a.auth.From, opts))
		if err != nil {
			return nil, nil, err
		}
		tx.L2GasLimit = new(big.Int).SetUint64(gas)
	}

	if err := a.insertGasPriceInTransactOpts(&opts); err != nil {
		return nil, nil, err
	}

	var gasPriceForEstimation *big.Int
	if opts.GasFeeCap != nil {
		gasPriceForEstimation = opts.GasFeeCap
	} else {
		gasPriceForEstimation = opts.GasPrice
	}

	baseCost, err := a.BaseCost(&CallOpts{
		Context: opts.Context,
	}, tx.L2GasLimit, tx.GasPerPubdataByte, gasPriceForEstimation)
	if err != nil {
		return nil, nil, err
	}

	if opts.Value == nil {
		opts.Value = new(big.Int).Add(baseCost, new(big.Int).Add(tx.OperatorTip, tx.L2Value))
	}

	err = utils.CheckBaseCost(baseCost, opts.Value)
	if err != nil {
		return nil, nil, err
	}
	return opts, &tx, nil
}

// Checks if the options contain a GasPrice or GasFeeCap, if not it will try to insert
// the GasFeeCap (maxFeePerGas) and GasTipCap (maxPriorityFee) if chain supports EIP1559,
// otherwise it sets GasPrice
func (a *WalletL1) insertGasPriceInTransactOpts(opts **TransactOpts) error {
	if *opts == nil {
		*opts = ensureTransactOpts(*opts)
	}
	if (*opts).GasFeeCap == nil && (*opts).GasPrice == nil {
		if isReady, head, err := a.checkIfL1ChainIsLondonReady(ensureContext((*opts).Context)); err != nil {
			return err
		} else if isReady {
			if (*opts).GasTipCap == nil {
				gasTipCap, errGasTipCap := a.clientL1.SuggestGasTipCap(ensureContext((*opts).Context))
				if errGasTipCap != nil {
					return errGasTipCap
				}
				(*opts).GasTipCap = gasTipCap
			}
			// geth by default uses multiplication by 2 (abi.BoundContract.createDynamicTx -> basefeeWiggleMultiplier),
			// but since the price for the L2 part will depend on the L1 part, doubling base fee is typically too much.
			// BaseFee * 3 / 2 + GasTipCap
			(*opts).GasFeeCap = new(big.Int).Add(
				(*opts).GasTipCap,
				new(big.Int).Div(new(big.Int).Mul(head.BaseFee, big.NewInt(3)), big.NewInt(2)),
			)
		} else {
			// Chain is not London ready -> use legacy transaction
			gasPrice, errGasPrice := a.clientL1.SuggestGasPrice(ensureContext((*opts).Context))
			if errGasPrice != nil {
				return errGasPrice
			}
			(*opts).GasPrice = gasPrice
		}
	}
	return nil
}

// Checks if the options contain a GasPrice or GasFeeCap, if not it will try to insert
// the GasFeeCap (maxFeePerGas) and GasTipCap (maxPriorityFee) if chain supports EIP1559,
// otherwise it sets GasPrice
func (a *WalletL1) insertGasPriceInDepositMsg(ctx context.Context, msg *DepositCallMsg) error {
	if msg == nil {
		msg = new(DepositCallMsg)
	}
	if msg.GasFeeCap == nil && msg.GasPrice == nil {
		if isReady, head, err := a.checkIfL1ChainIsLondonReady(ensureContext(ctx)); err != nil {
			return err
		} else if isReady {
			if msg.GasTipCap == nil {
				gasTipCap, errGasTipCap := a.clientL1.SuggestGasTipCap(ensureContext(ctx))
				if errGasTipCap != nil {
					return errGasTipCap
				}
				msg.GasTipCap = gasTipCap
			}
			// geth by default uses multiplication by 2 (abi.BoundContract.createDynamicTx -> basefeeWiggleMultiplier),
			// but since the price for the L2 part will depend on the L1 part, doubling base fee is typically too much.
			// BaseFee * 3 / 2 + GasTipCap
			msg.GasFeeCap = new(big.Int).Add(
				msg.GasTipCap,
				new(big.Int).Div(new(big.Int).Mul(head.BaseFee, big.NewInt(3)), big.NewInt(2)),
			)
		} else {
			// Chain is not London ready -> use legacy transaction
			gasPrice, errGasPrice := a.clientL1.SuggestGasPrice(ensureContext(ctx))
			if errGasPrice != nil {
				return errGasPrice
			}
			msg.GasPrice = gasPrice
		}
	}
	return nil
}

func (a *WalletL1) getWithdrawalLog(ctx context.Context, withdrawalHash common.Hash, index int) (*zkTypes.Log, *big.Int, error) {
	receipt, err := (*a.clientL2).TransactionReceipt(ctx, withdrawalHash)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return nil, nil, errors.New("transaction receipt not found")
	}
	fLogs := make([]*zkTypes.Log, 0)
	for _, l := range receipt.Logs {
		if l.Address == utils.L1MessengerAddress && len(l.Topics) > 0 &&
			bytes.Equal(l.Topics[0].Bytes(), crypto.Keccak256([]byte("L1MessageSent(address,bytes32,bytes)"))) {
			fLogs = append(fLogs, l)
		}
	}
	if len(fLogs) < index+1 {
		return nil, nil, errors.New("withdrawal log not found")
	}
	return fLogs[index], receipt.L1BatchTxIndex.ToInt(), nil
}

func (a *WalletL1) getWithdrawalL2ToL1Log(ctx context.Context, withdrawalHash common.Hash, index int) (int, *zkTypes.L2ToL1Log, error) {
	receipt, err := (*a.clientL2).TransactionReceipt(ctx, withdrawalHash)
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
		if l.Sender == utils.L1MessengerAddress {
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

func (a *WalletL1) checkIfL1ChainIsLondonReady(ctx context.Context) (bool, *types.Header, error) {
	// Only query for block header not whole block with transactions
	if head, err := a.clientL1.HeaderByNumber(ensureContext(ctx), nil); err != nil {
		return false, nil, err
	} else if head.BaseFee != nil {
		return true, head, nil
	} else {
		return false, head, nil
	}
}
