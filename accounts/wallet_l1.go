package accounts

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/bridgehub"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/l1bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l1messenger"
	"github.com/zksync-sdk/zksync2-go/contracts/l1sharedbridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/l2sharedbridge"
	"github.com/zksync-sdk/zksync2-go/contracts/zksynchyperchain"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"strings"
)

// WalletL1 is associated with an account and provides common operations on the
// L1 network for the associated account.
type WalletL1 struct {
	clientL1        *ethclient.Client
	clientL2        *clients.Client
	auth            *bind.TransactOpts
	l1ChainId       *big.Int
	l2ChainId       *big.Int
	baseToken       common.Address
	isEthBasedChain bool

	mainContractAddress common.Address
	mainContract        *zksynchyperchain.IZkSyncHyperchain

	bridgehubAddress common.Address
	bridgehub        *bridgehub.IBridgehub

	sharedL1BridgeAddress common.Address
	sharedL1Bridge        *l1sharedbridge.IL1SharedBridge

	defaultL1BridgeAddress common.Address
	defaultL1Bridge        *l1bridge.IL1Bridge
}

// NewWalletL1 creates an instance of WalletL1 associated with the account provided by the raw private key.
func NewWalletL1(rawPrivateKey []byte, clientL1 *ethclient.Client, clientL2 *clients.Client) (*WalletL1, error) {
	chainID, err := clientL1.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	signer, err := NewECDSASignerFromRawPrivateKey(rawPrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	return NewWalletL1FromSigner(signer, clientL1, clientL2)
}

// NewWalletL1FromSigner creates an instance of WalletL1 associated with the account provided by the signer.
func NewWalletL1FromSigner(signer *ECDSASigner, clientL1 *ethclient.Client, clientL2 *clients.Client) (*WalletL1, error) {
	if signer == nil {
		return nil, errors.New("signer is not provided")
	} else if clientL1 == nil {
		return nil, errors.New("clientL1 is not provided")
	} else if clientL2 == nil {
		return nil, errors.New("clientL2 is not provided")
	}

	mainContractAddress, err := clientL2.MainContractAddress(context.Background())
	if err != nil {
		return nil, err
	}
	iZkSync, err := zksynchyperchain.NewIZkSyncHyperchain(mainContractAddress, clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IZkSync: %w", err)
	}

	bridgehubContractAddress, err := clientL2.BridgehubContractAddress(context.Background())
	if err != nil {
		return nil, err
	}
	iBridgehub, err := bridgehub.NewIBridgehub(bridgehubContractAddress, clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IBridgehub: %w", err)
	}

	bridgeContracts, err := clientL2.BridgeContracts(context.Background())
	if err != nil {
		return nil, err
	}

	sharedL1Bridge, err := l1sharedbridge.NewIL1SharedBridge(bridgeContracts.L1SharedBridge, clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IL1SharedBridge: %w", err)
	}
	defaultL1Bridge, err := l1bridge.NewIL1Bridge(bridgeContracts.L1Erc20Bridge, clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IL1Bridge: %w", err)
	}

	l1ChainId, err := clientL1.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	l2ChainId, err := clientL2.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := newTransactorWithSigner(signer, l1ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to init TransactOpts: %w", err)
	}

	baseToken, err := clientL2.BaseTokenContractAddress(context.Background())
	if err != nil {
		return nil, err
	}
	isEthBasedChain, err := clientL2.IsEthBasedChain(context.Background())
	if err != nil {
		return nil, err
	}

	return &WalletL1{
		clientL1:               clientL1,
		clientL2:               clientL2,
		auth:                   auth,
		l1ChainId:              l1ChainId,
		l2ChainId:              l2ChainId,
		baseToken:              baseToken,
		isEthBasedChain:        isEthBasedChain,
		mainContractAddress:    mainContractAddress,
		mainContract:           iZkSync,
		bridgehubAddress:       bridgehubContractAddress,
		bridgehub:              iBridgehub,
		sharedL1BridgeAddress:  bridgeContracts.L1SharedBridge,
		sharedL1Bridge:         sharedL1Bridge,
		defaultL1BridgeAddress: bridgeContracts.L1Erc20Bridge,
		defaultL1Bridge:        defaultL1Bridge,
	}, nil
}

// MainContract returns the ZKsync L1 smart contract.
func (w *WalletL1) MainContract(_ context.Context) (*zksynchyperchain.IZkSyncHyperchain, error) {
	return w.mainContract, nil
}

// BridgehubContract returns the Bridgehub L1 smart contract.
func (w *WalletL1) BridgehubContract(_ context.Context) (*bridgehub.IBridgehub, error) {
	return w.bridgehub, nil
}

// L1BridgeContracts returns L1 bridge contracts.
func (w *WalletL1) L1BridgeContracts(_ context.Context) (*types.L1BridgeContracts, error) {
	return &types.L1BridgeContracts{Erc20: w.defaultL1Bridge, Shared: w.sharedL1Bridge}, nil
}

// BaseToken returns the address of the base token on L1.
func (w *WalletL1) BaseToken(opts *CallOpts) (common.Address, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(w.auth.From)
	return w.bridgehub.BaseToken(callOpts, w.l2ChainId)
}

// IsEthBasedChain returns whether the chain is ETH-based.
func (w *WalletL1) IsEthBasedChain(ctx context.Context) (bool, error) {
	return w.clientL2.IsEthBasedChain(ctx)
}

// BalanceL1 returns the balance of the specified token on L1 that can be
// either base token or any ERC20 token.
func (w *WalletL1) BalanceL1(opts *CallOpts, token common.Address) (*big.Int, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(w.auth.From)
	if token == utils.LegacyEthAddress || token == utils.EthAddressInContracts {
		return w.clientL1.BalanceAt(callOpts.Context, w.auth.From, callOpts.BlockNumber)
	} else {
		erc20Contract, err := erc20.NewIERC20(token, w.clientL1)
		if err != nil {
			return nil, fmt.Errorf("failed to load IERC20: %w", err)
		}
		return erc20Contract.BalanceOf(callOpts, w.auth.From)
	}
}

// AllowanceL1 returns the amount of approved tokens for a specific L1 bridge.
func (w *WalletL1) AllowanceL1(opts *CallOpts, token common.Address, bridgeAddress common.Address) (*big.Int, error) {
	if token == (common.Address{}) {
		token = utils.LegacyEthAddress
	}
	if bridgeAddress == (common.Address{}) {
		bridgeAddress = w.sharedL1BridgeAddress
	}

	erc20Contract, err := erc20.NewIERC20(token, w.clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IERC20: %w", err)
	}
	return erc20Contract.Allowance(ensureCallOpts(opts).ToCallOpts(w.auth.From), w.auth.From, bridgeAddress)
}

// L2TokenAddress returns the corresponding address on the L2 network for the token on the L1 network.
func (w *WalletL1) L2TokenAddress(ctx context.Context, token common.Address) (common.Address, error) {
	return w.clientL2.L2TokenAddress(ensureContext(ctx), token)
}

// ApproveERC20 approves the specified amount of tokens for the specified L1 bridge.
func (w *WalletL1) ApproveERC20(auth *TransactOpts, token common.Address, amount *big.Int,
	bridgeAddress common.Address) (*ethTypes.Transaction, error) {
	if token == (common.Address{}) {
		return nil, errors.New("token L1 address must be provided")
	}
	if token == utils.LegacyEthAddress {
		return nil, errors.New("ETH token can't be approved. The address of the token does not exist on L1")
	}

	opts := ensureTransactOpts(auth).ToTransactOpts(w.auth.From, w.auth.Signer)
	if bridgeAddress == (common.Address{}) {
		if !w.isEthBasedChain && token == w.baseToken {
			sharedBridge, err := w.bridgehub.SharedBridge(&bind.CallOpts{
				From:    opts.From,
				Context: opts.Context,
			})
			if err != nil {
				return nil, err
			}
			bridgeAddress = sharedBridge
		} else {
			bridgeAddress = w.sharedL1BridgeAddress
		}
	}

	erc20Contract, err := erc20.NewIERC20(token, w.clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to load IERC20: %w", err)
	}
	return erc20Contract.Approve(opts, bridgeAddress, amount)
}

// BaseCost returns base cost for L2 transaction.
func (w *WalletL1) BaseCost(opts *CallOpts, gasLimit, gasPerPubdataByte, gasPrice *big.Int) (*big.Int, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(w.auth.From)
	if gasPrice == nil {
		var err error
		if gasPrice, err = w.clientL1.SuggestGasPrice(callOpts.Context); err != nil {
			return nil, err
		}
	}
	return w.bridgehub.L2TransactionBaseCost(callOpts,
		w.l2ChainId,
		gasPrice,
		gasLimit,
		gasPerPubdataByte,
	)
}

// DepositAllowanceParams returns the parameters for the approval token transaction based on the deposit token and amount.
// Some deposit transactions require multiple approvals. Existing allowance for the bridge is not checked;
// allowance is calculated solely based on the specified amount.
func (w *WalletL1) DepositAllowanceParams(opts *CallOpts, msg DepositCallMsg) ([]struct {
	Token     common.Address
	Allowance *big.Int
}, error) {
	opts = ensureCallOpts(opts)
	if msg.Token == utils.LegacyEthAddress {
		msg.Token = utils.EthAddressInContracts
	}
	if w.isEthBasedChain && msg.Token == utils.EthAddressInContracts {
		return nil, errors.New("token ETH can't be approved. The address of the token does not exist on L1")
	}

	transactOpts := msg.ToTransactOpts()
	tx := msg.ToDepositTransaction()
	mintValue, err := w.calculateMintValue(&transactOpts, &tx)
	if err != nil {
		return nil, err
	}
	if w.isEthBasedChain || msg.Token == utils.EthAddressInContracts || msg.Token == w.baseToken {
		return []struct {
			Token     common.Address
			Allowance *big.Int
		}{{w.baseToken, mintValue}}, nil
	} else {
		return []struct {
			Token     common.Address
			Allowance *big.Int
		}{{w.baseToken, mintValue}, {msg.Token, msg.Amount}}, nil
	}
}

// Deposit transfers the specified token from the associated account on the L1 network
// to the target account on the L2 network. The token can be either ETH or any ERC20 token.
// For ERC20 tokens, enough approved tokens must be associated with the specified L1 bridge
// (default one or the one defined in DepositTransaction.BridgeAddress).
// In this case, depending on is the chain ETH-based or not DepositTransaction.ApproveERC20 or
// DepositTransaction.ApproveBaseERC20 can be enabled to perform token approval.
// If there are already enough approved tokens for the L1 bridge, token approval will be skipped.
// To check the amount of approved tokens for a specific bridge, use the AdapterL1.AllowanceL1 method.
func (w *WalletL1) Deposit(auth *TransactOpts, tx DepositTransaction) (*ethTypes.Transaction, error) {
	opts := ensureTransactOpts(auth)
	if tx.Token == utils.LegacyEthAddress {
		tx.Token = utils.EthAddressInContracts
	}
	if w.isEthBasedChain && tx.Token == utils.EthAddressInContracts {
		return w.depositEthToEthBasedChain(opts, &tx)
	} else if w.isEthBasedChain {
		return w.depositTokenToEthBasedChain(opts, &tx)
	} else if tx.Token == utils.EthAddressInContracts {
		return w.depositEthToNonEthBasedChain(opts, &tx)
	} else if tx.Token == w.baseToken {
		return w.depositBaseTokenToNonEthBasedChain(opts, &tx)
	} else {
		return w.depositNonBaseTokenToNonEthBasedChain(opts, &tx)
	}
}

// EstimateGasDeposit estimates the amount of gas required for a deposit transaction on L1 network.
// Gas of approving ERC20 token is not included in estimation.
func (w *WalletL1) EstimateGasDeposit(ctx context.Context, msg DepositCallMsg) (uint64, error) {
	if msg.Token == utils.LegacyEthAddress {
		msg.Token = utils.EthAddressInContracts
	}

	opts := msg.ToTransactOpts()
	opts.Context = ctx
	depositTx := msg.ToDepositTransaction()
	if err := w.populateDepositTxWithDefaults(&opts, &depositTx); err != nil {
		return 0, err
	}
	mintValue, err := w.calculateMintValue(&opts, &depositTx)
	if err != nil {
		return 0, err
	}
	if w.isEthBasedChain && msg.Token == utils.EthAddressInContracts {
		tx, prepareErr := w.prepareDepositEthToEthBasedChain(&opts, &depositTx, mintValue)
		if prepareErr != nil {
			return 0, prepareErr
		}
		return w.EstimateGasRequestExecute(ensureContext(ctx), tx.ToRequestExecuteCallMsg(&opts))
	} else if w.isEthBasedChain {
		tx, prepareErr := w.prepareDepositTokenToEthBasedChain(&opts, &depositTx, mintValue)
		if prepareErr != nil {
			return 0, prepareErr
		}
		return tx.Gas(), nil
	} else if msg.Token == utils.EthAddressInContracts {
		tx, prepareErr := w.prepareDepositEthToNonEthBasedChain(&opts, &depositTx, mintValue)
		if prepareErr != nil {
			return 0, prepareErr
		}
		return tx.Gas(), nil
	} else if msg.Token == w.baseToken {
		tx, prepareErr := w.prepareDepositBaseTokenToNonEthBasedChain(&opts, &depositTx, mintValue)
		if prepareErr != nil {
			return 0, prepareErr
		}
		return w.EstimateGasRequestExecute(ensureContext(ctx), tx.ToRequestExecuteCallMsg(&opts))
	} else {
		tx, prepareErr := w.prepareDepositNonBasedTokenToNonEthBasedChain(&opts, &depositTx, mintValue)
		if prepareErr != nil {
			return 0, prepareErr
		}
		return tx.Gas(), nil
	}
}

// FullRequiredDepositFee retrieves the full needed ETH fee for the deposit on both L1 and L2 networks.
func (w *WalletL1) FullRequiredDepositFee(ctx context.Context, msg DepositCallMsg) (*FullDepositFee, error) {
	if msg.Token == utils.LegacyEthAddress {
		msg.Token = utils.EthAddressInContracts
	}

	// It is assumed that the L2 fee for the transaction does not depend on its value.
	dummyAmount := big.NewInt(1)
	opts := msg.ToTransactOpts()
	tx := msg.ToDepositTransaction()
	tx.Amount = dummyAmount

	if err := w.populateDepositTxWithDefaults(&opts, &tx); err != nil {
		return nil, err
	}
	baseCost, err := w.calculateBaseCost(&opts, &tx)
	if err != nil {
		return nil, err
	}
	mintValue, err := w.calculateMintValueFromBaseCost(baseCost, &tx)
	if err != nil {
		return nil, err
	}

	gasPriceForEstimation := opts.GasPrice
	if opts.GasFeeCap != nil {
		gasPriceForEstimation = opts.GasFeeCap
	}

	if w.isEthBasedChain {
		selfBalanceETH, balanceErr := w.BalanceL1(nil, utils.EthAddressInContracts)
		if balanceErr != nil {
			return nil, balanceErr
		}
		// To ensure that L1 gas estimation succeeds when using estimateGasDeposit,
		// the account needs to have w sufficient ETH balance.
		if baseCost.Cmp(new(big.Int).Add(selfBalanceETH, dummyAmount)) >= 0 {
			recommendedETHBalance := utils.L1RecommendedMinEthDepositGasLimit
			if msg.Token == utils.EthAddressInContracts {
				recommendedETHBalance = utils.L1RecommendedMinErc20DepositGasLimit
			}
			recommendedETHBalance.Mul(recommendedETHBalance, gasPriceForEstimation)
			recommendedETHBalance.Add(recommendedETHBalance, baseCost)
			return nil, fmt.Errorf("not enough balance for deposit. Under the provided gas price, "+
				"the recommended balance to perform w deposit is %v ETH", recommendedETHBalance)
		}

		if msg.Token != utils.EthAddressInContracts {
			bridgeAddress := w.sharedL1BridgeAddress
			if msg.BridgeAddress != nil {
				bridgeAddress = *msg.BridgeAddress
			}
			allowance, allowanceErr := w.AllowanceL1(&CallOpts{Context: ensureContext(ctx)}, msg.Token, bridgeAddress)
			if allowanceErr != nil {
				return nil, allowanceErr
			}
			if allowance.Cmp(dummyAmount) < 0 {
				return nil, errors.New("not enough allowance to cover the deposit")
			}
		}
	} else {
		allowance, allowanceErr := w.AllowanceL1(&CallOpts{Context: ensureContext(ctx)}, w.baseToken, w.sharedL1BridgeAddress)
		if allowanceErr != nil {
			return nil, allowanceErr
		}
		if allowance.Cmp(mintValue) < 0 {
			return nil, errors.New("not enough base token allowance to cover the deposit")
		}

		if msg.Token == utils.EthAddressInContracts || msg.Token == w.baseToken {
			if msg.Value == nil {
				msg.Value = msg.Amount
			}
		} else {
			if msg.Value == nil {
				msg.Value = big.NewInt(0)
			}
			allowance, allowanceErr = w.AllowanceL1(&CallOpts{Context: ensureContext(ctx)}, msg.Token, w.sharedL1BridgeAddress)
			if allowanceErr != nil {
				return nil, allowanceErr
			}
			if allowance.Cmp(dummyAmount) < 0 {
				return nil, errors.New("not enough token allowance to cover the deposit")
			}
		}
	}

	// Deleting the explicit gas limits in the fee estimation
	// in order to prevent the situation where the transaction
	// fails because the user does not have enough balance.
	// In that matter, new DepositCalMsg is created with empty
	// GasPrice, GasFeeCap and GasTipCap field.
	l1GasLimit, err := w.EstimateGasDeposit(ensureContext(ctx), DepositCallMsg{
		To:                tx.To,
		Token:             tx.Token,
		Amount:            dummyAmount,
		OperatorTip:       tx.OperatorTip,
		BridgeAddress:     tx.BridgeAddress,
		L2GasLimit:        tx.L2GasLimit,
		GasPerPubdataByte: tx.GasPerPubdataByte,
		RefundRecipient:   tx.RefundRecipient,
		CustomBridgeData:  tx.CustomBridgeData,
		Value:             opts.Value,
	})
	if err != nil {
		return nil, err
	}

	fullConst := &FullDepositFee{
		BaseCost:   baseCost,
		L1GasLimit: new(big.Int).SetUint64(l1GasLimit),
		L2GasLimit: tx.L2GasLimit,
	}

	if msg.GasPrice != nil {
		fullConst.GasPrice = opts.GasPrice
	} else {
		fullConst.MaxPriorityFeePerGas = opts.GasTipCap
		fullConst.MaxFeePerGas = opts.GasFeeCap
	}
	return fullConst, nil
}

// FinalizeWithdraw proves the inclusion of the L2 -> L1 withdrawal message.
func (w *WalletL1) FinalizeWithdraw(auth *TransactOpts, withdrawalHash common.Hash, index int) (*ethTypes.Transaction, error) {
	if w.clientL1 == nil {
		return nil, errors.New("ethereum provider is not initialized")
	}
	var opts *bind.TransactOpts
	if auth == nil {
		opts = &bind.TransactOpts{
			From:    w.auth.From,
			Signer:  w.auth.Signer,
			Context: context.Background(),
		}
	} else {
		opts = auth.ToTransactOpts(w.auth.From, w.auth.Signer)
	}

	log, l1BatchTxId, err := w.getWithdrawalLog(opts.Context, withdrawalHash, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get WithdrawalLog: %w", err)
	}
	if l1BatchTxId == nil {
		return nil, errors.New("empty l1BatchTxIndex")
	}
	l2ToL1LogIndex, _, err := w.getWithdrawalL2ToL1Log(opts.Context, withdrawalHash, index)
	if err != nil {
		return nil, fmt.Errorf("failed to get WithdrawalL2ToL1Log: %w", err)
	}
	if len(log.Topics) < 2 {
		return nil, errors.New("not enough Topics count")
	}

	proof, err := w.clientL2.LogProof(opts.Context, withdrawalHash, l2ToL1LogIndex)
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
	sender := common.BytesToAddress(log.Topics[1].Bytes()[12:])

	if sender == utils.L2BaseTokenAddress {
		return w.sharedL1Bridge.FinalizeWithdrawal(
			opts,
			w.l2ChainId,
			log.L1BatchNumber.ToInt(),
			big.NewInt(int64(proof.Id)),
			uint16(l1BatchTxId.Uint64()),
			message,
			proof32)
	}

	var l1Bridge *l1sharedbridge.IL1SharedBridge
	isLegacyBridge, err := w.clientL2.IsL2BridgeLegacy(opts.Context, sender)
	if err != nil {
		return nil, fmt.Errorf("failed to check if L2 bridge is legacy: %w", err)
	}

	if isLegacyBridge {
		l2Bridge, errL2Bridge := l2bridge.NewIL2Bridge(sender, w.clientL2)
		if errL2Bridge != nil {
			return nil, fmt.Errorf("failed to connect to legacy L2 bridge: %w", errL2Bridge)
		}
		l1BridgeAddress, errL1Bridge := l2Bridge.L1Bridge(&bind.CallOpts{Context: opts.Context})
		if errL1Bridge != nil {
			return nil, errL1Bridge
		}
		l1Bridge, errL1Bridge = l1sharedbridge.NewIL1SharedBridge(l1BridgeAddress, w.clientL1)
	} else {
		l2Bridge, errL2Bridge := l2sharedbridge.NewIL2SharedBridge(sender, w.clientL2)
		if errL2Bridge != nil {
			return nil, fmt.Errorf("failed to connect to legacy L2 bridge: %w", errL2Bridge)
		}
		l1BridgeAddress, errL1Bridge := l2Bridge.L1SharedBridge(&bind.CallOpts{Context: opts.Context})
		if errL1Bridge != nil {
			return nil, errL1Bridge
		}
		l1Bridge, errL1Bridge = l1sharedbridge.NewIL1SharedBridge(l1BridgeAddress, w.clientL1)
	}

	return l1Bridge.FinalizeWithdrawal(
		opts,
		w.l2ChainId,
		log.L1BatchNumber.ToInt(),
		big.NewInt(int64(proof.Id)),
		uint16(l1BatchTxId.Uint64()),
		message,
		proof32)
}

// IsWithdrawFinalized checks if the withdrawal finalized on L1 network.
func (w *WalletL1) IsWithdrawFinalized(opts *CallOpts, withdrawalHash common.Hash, index int) (bool, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(w.auth.From)
	if w.clientL1 == nil {
		return false, errors.New("ethereum provider is not initialized")
	}
	log, _, err := w.getWithdrawalLog(callOpts.Context, withdrawalHash, index)
	if err != nil {
		return false, fmt.Errorf("failed to get WithdrawalLog: %w", err)
	}
	l2ToL1LogIndex, _, err := w.getWithdrawalL2ToL1Log(callOpts.Context, withdrawalHash, index)
	if err != nil {
		return false, fmt.Errorf("failed to get WithdrawalL2ToL1Log: %w", err)
	}
	if len(log.Topics) < 2 {
		return false, errors.New("not enough Topics count")
	}
	sender := common.BytesToAddress(log.Topics[1].Bytes()[12:])
	proof, err := w.clientL2.LogProof(callOpts.Context, withdrawalHash, l2ToL1LogIndex)
	if err != nil {
		return false, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}

	l1BridgeAddress := w.sharedL1BridgeAddress
	isBaseToken, err := w.clientL2.IsBaseToken(callOpts.Context, sender)
	if err != nil {
		return false, err
	}
	if !isBaseToken {
		l2Bridge, bridgeErr := l2sharedbridge.NewIL2SharedBridge(sender, w.clientL2)
		if bridgeErr != nil {
			return false, fmt.Errorf("failed to init l2Bridge: %w", err)
		}
		l1BridgeAddress, bridgeErr = l2Bridge.L1SharedBridge(callOpts)
		if bridgeErr != nil {
			return false, fmt.Errorf("failed to get l1BridgeAddress: %w", err)
		}
	}

	l1Bridge, err := l1sharedbridge.NewIL1SharedBridge(l1BridgeAddress, w.clientL1)
	if err != nil {
		return false, fmt.Errorf("failed to init l1Bridge: %w", err)
	}
	return l1Bridge.IsWithdrawalFinalized(callOpts, w.l2ChainId, log.L1BatchNumber.ToInt(), big.NewInt(int64(proof.Id)))
}

// ClaimFailedDeposit withdraws funds from the initiated deposit, which failed when finalizing on L2.
// If the deposit L2 transaction has failed, it sends an L1 transaction calling ClaimFailedDeposit method
// of the L1 bridge, which results in returning L1 tokens back to the depositor, otherwise throws the error.
func (w *WalletL1) ClaimFailedDeposit(auth *TransactOpts, depositHash common.Hash) (*ethTypes.Transaction, error) {
	opts := ensureTransactOpts(auth)
	receipt, err := w.clientL2.TransactionReceipt(opts.Context, depositHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return nil, errors.New("transaction receipt not found")
	}
	var successL2ToL1Log *types.L2ToL1Log
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

	tx, _, err := w.clientL2.TransactionByHash(opts.Context, depositHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get Transaction: %w", err)
	}
	if tx == nil {
		return nil, errors.New("transaction not found")
	}

	// Undo the aliasing, since the Mailbox contract set it as for contract address.
	l1BridgeAddress := utils.UndoL1ToL2Alias(receipt.From)
	l1Bridge, err := l1sharedbridge.NewIL1SharedBridge(l1BridgeAddress, w.clientL1)
	if err != nil {
		return nil, fmt.Errorf("failed to init l1Bridge: %w", err)
	}
	l2Bridge, err := abi.JSON(strings.NewReader(l2bridge.IL2BridgeMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load l2Bridge ABI: %w", err)
	}
	// skip first 4 bytes is the function selector, the input values are the rest
	finalizeDepositMethod, err := l2Bridge.MethodById(tx.Data[:4])
	if err != nil {
		return nil, err
	}
	inputValues, err := finalizeDepositMethod.Inputs.Unpack(tx.Data[4:])
	if err != nil {
		return nil, err
	}
	if len(inputValues) < 4 {
		return nil, errors.New("unpacked calldata is empty")
	}
	l1Sender := inputValues[0].(common.Address)
	l1Token := inputValues[2].(common.Address)
	amount := inputValues[3].(*big.Int)

	proof, err := w.clientL2.LogProof(opts.Context, depositHash, successL2ToL1LogIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2ToL1LogProof: %w", err)
	}

	proof32 := make([][32]byte, len(proof.Proof))
	for i, pr := range proof.Proof {
		proof32[i] = pr
	}

	return l1Bridge.ClaimFailedDeposit(
		opts.ToTransactOpts(w.auth.From, w.auth.Signer),
		w.l2ChainId,
		l1Sender,
		l1Token,
		amount,
		depositHash,
		receipt.L1BatchNumber.ToInt(),
		big.NewInt(int64(proof.Id)),
		uint16(receipt.L1BatchTxIndex.ToInt().Uint64()),
		proof32,
	)
}

// RequestExecute request execution of L2 transaction from L1.
func (w *WalletL1) RequestExecute(auth *TransactOpts, tx RequestExecuteTransaction) (*ethTypes.Transaction, error) {
	opts := ensureTransactOpts(auth)
	err := w.prepareRequestExecuteTx(opts, &tx)
	if err != nil {
		return nil, err
	}
	return w.bridgehub.RequestL2TransactionDirect(
		opts.ToTransactOpts(w.auth.From, w.auth.Signer),
		bridgehub.L2TransactionRequestDirect{
			ChainId:                  w.l2ChainId,
			MintValue:                tx.MintValue,
			L2Contract:               tx.ContractAddress,
			L2Value:                  tx.L2Value,
			L2Calldata:               tx.Calldata,
			L2GasLimit:               tx.L2GasLimit,
			L2GasPerPubdataByteLimit: tx.GasPerPubdataByte,
			FactoryDeps:              tx.FactoryDeps,
			RefundRecipient:          tx.RefundRecipient,
		},
	)
}

// EstimateGasRequestExecute estimates the amount of gas required for a request execute transaction.
func (w *WalletL1) EstimateGasRequestExecute(ctx context.Context, msg RequestExecuteCallMsg) (uint64, error) {
	opts := msg.ToTransactOpts()
	tx := msg.ToRequestExecuteTransaction()
	err := w.prepareRequestExecuteTx(&opts, &tx)
	if err != nil {
		return 0, err
	}

	preparedRequestExecuteCallMsg := tx.ToRequestExecuteCallMsg(&opts)
	callMsg, err := preparedRequestExecuteCallMsg.ToCallMsgWithChainID(w.auth.From, w.l2ChainId)
	if err != nil {
		return 0, err
	}

	return w.clientL1.EstimateGas(ensureContext(ctx), callMsg)
}

// RequestExecuteAllowanceParams returns the parameters for the approval token transaction based on the request execute transaction.
// Existing allowance for the bridge is not checked; allowance is calculated solely based on the specified transaction.
func (w *WalletL1) RequestExecuteAllowanceParams(opts *CallOpts, msg RequestExecuteCallMsg) (AllowanceParams, error) {
	if w.isEthBasedChain {
		return AllowanceParams{}, errors.New("token ETH can't be approved. The address of the token does not exist on L1")
	}

	opts = ensureCallOpts(opts)
	transactOpts := msg.ToTransactOpts()
	tx := msg.ToRequestExecuteTransaction()
	err := w.prepareRequestExecuteTx(&transactOpts, &tx)
	if err != nil {
		return AllowanceParams{}, err
	}
	return AllowanceParams{Token: w.baseToken, Allowance: tx.MintValue}, nil

}

// PriorityOpConfirmation returns the transaction confirmation data that is part of `L2->L1` message.
// The txHash is the  hash of the L2 transaction where the message was initiated.
// The index is used in case there were multiple transactions in one message, you may pass an index of the
// transaction which confirmation data should be fetched.
func (w *WalletL1) PriorityOpConfirmation(ctx context.Context, txHash common.Hash, index int) (*types.PriorityOpConfirmation, error) {
	return w.clientL2.PriorityOpConfirmation(ctx, txHash, index)
}

// EstimateDepositL2GasFromCustomBridge used by EstimateDepositL2GasFromDefaultBridge to estimate L2 gas required for token
// bridging via a custom ERC20 bridge.
func (w *WalletL1) EstimateDepositL2GasFromCustomBridge(ctx context.Context, l1BridgeAddress, l2BridgeAddress, token common.Address,
	amount *big.Int, to common.Address, bridgeData []byte, from common.Address, gasPerPubdataByte, l2Value *big.Int) (uint64, error) {
	calldata, err := utils.Erc20BridgeCalldata(token, from, to, amount, bridgeData)
	if err != nil {
		return 0, err
	}

	return w.clientL2.EstimateL1ToL2Execute(ensureContext(ctx), types.CallMsg{
		From:          utils.ApplyL1ToL2Alias(l1BridgeAddress),
		To:            &l2BridgeAddress,
		Data:          calldata,
		Value:         l2Value,
		GasPerPubdata: gasPerPubdataByte,
	})
}

// EstimateDepositL2GasFromDefaultBridge returns an estimation of L2 gas required for token bridging via the default ERC20
// bridge.
func (w *WalletL1) EstimateDepositL2GasFromDefaultBridge(ctx context.Context, token common.Address, amount *big.Int,
	to, from common.Address, gasPerPubdataByte *big.Int) (uint64, error) {

	// If the `from` address is not provided, we use w random address, because
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

	if token == w.baseToken {
		return w.clientL2.EstimateL1ToL2Execute(ensureContext(ctx), types.CallMsg{
			From:          from,
			To:            &to,
			Value:         amount,
			GasPerPubdata: gasPerPubdataByte,
		})
	} else {
		bridgeContracts, err := w.clientL2.BridgeContracts(ensureContext(ctx))
		if err != nil {
			return 0, err
		}
		calldata, err := utils.Erc20DefaultBridgeData(token, w.clientL1)
		if err != nil {
			return 0, err
		}
		if token == utils.LegacyEthAddress {
			token = utils.EthAddressInContracts
		}
		return w.EstimateDepositL2GasFromCustomBridge(
			ensureContext(ctx),
			bridgeContracts.L1SharedBridge,
			bridgeContracts.L2SharedBridge,
			token,
			amount,
			to,
			calldata,
			from,
			gasPerPubdataByte,
			common.Big0)
	}
}

func (w *WalletL1) depositEthToEthBasedChain(opts *TransactOpts, tx *DepositTransaction) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	mintValue, err := w.calculateMintValue(opts, tx)
	if err != nil {
		return nil, err
	}

	preparedTx, err := w.prepareDepositEthToEthBasedChain(opts, tx, mintValue)
	if err != nil {
		return nil, err
	}

	return w.RequestExecute(opts, *preparedTx)
}

func (w *WalletL1) prepareDepositEthToEthBasedChain(opts *TransactOpts, tx *DepositTransaction, mintValue *big.Int) (*RequestExecuteTransaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	if opts.Value == nil {
		opts.Value = mintValue
	}
	return &RequestExecuteTransaction{
		ContractAddress:   tx.To,
		MintValue:         opts.Value,
		L2Value:           tx.Amount,
		L2GasLimit:        tx.L2GasLimit,
		GasPerPubdataByte: tx.GasPerPubdataByte,
		OperatorTip:       tx.OperatorTip,
		RefundRecipient:   tx.RefundRecipient,
	}, nil
}

func (w *WalletL1) depositTokenToEthBasedChain(opts *TransactOpts, tx *DepositTransaction) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	baseCost, err := w.calculateBaseCost(opts, tx)
	if err != nil {
		return nil, err
	}
	mintValue, err := w.calculateMintValueFromBaseCost(baseCost, tx)
	if err != nil {
		return nil, err
	}
	checkErr := utils.CheckBaseCost(baseCost, mintValue)
	if checkErr != nil {
		return nil, checkErr
	}

	if tx.ApproveERC20 {
		bridge := w.sharedL1BridgeAddress
		if tx.BridgeAddress != nil {
			bridge = *tx.BridgeAddress
		}
		approveErr := w.approveERC20(tx.ApproveAuth, tx.Token, tx.Amount, bridge)
		if approveErr != nil {
			return nil, approveErr
		}
	}

	preparedTx, err := w.prepareDepositTokenToEthBasedChain(opts, tx, mintValue)
	if err != nil {
		return nil, err
	}
	if sendErr := w.clientL1.SendTransaction(opts.Context, preparedTx); sendErr != nil {
		return nil, sendErr
	}
	return preparedTx, nil
}

func (w *WalletL1) prepareDepositTokenToEthBasedChain(opts *TransactOpts, tx *DepositTransaction, mintValue *big.Int) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	if opts.Value == nil {
		opts.Value = mintValue
	}

	secondBridgeAddress := w.sharedL1BridgeAddress
	if tx.BridgeAddress != nil {
		secondBridgeAddress = *tx.BridgeAddress
	}

	secondBridgeCalldata, err := w.secondBridgeCalldata(tx.Token, tx.To, tx.Amount)
	if err != nil {
		return nil, err
	}

	noSendOpts := opts.ToTransactOpts(w.auth.From, w.auth.Signer)
	noSendOpts.NoSend = true
	depositTx, err := w.bridgehub.RequestL2TransactionTwoBridges(noSendOpts, bridgehub.L2TransactionRequestTwoBridgesOuter{
		ChainId:                  w.l2ChainId,
		MintValue:                mintValue,
		L2Value:                  big.NewInt(0),
		L2GasLimit:               tx.L2GasLimit,
		L2GasPerPubdataByteLimit: tx.GasPerPubdataByte,
		RefundRecipient:          tx.RefundRecipient,
		SecondBridgeAddress:      secondBridgeAddress,
		SecondBridgeValue:        big.NewInt(0),
		SecondBridgeCalldata:     secondBridgeCalldata,
	})
	if err != nil {
		return nil, err
	}
	return depositTx, nil
}

func (w *WalletL1) depositEthToNonEthBasedChain(opts *TransactOpts, tx *DepositTransaction) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	baseCost, err := w.calculateBaseCost(opts, tx)
	if err != nil {
		return nil, err
	}
	mintValue, err := w.calculateMintValueFromBaseCost(baseCost, tx)
	if err != nil {
		return nil, err
	}
	checkErr := utils.CheckBaseCost(baseCost, mintValue)
	if checkErr != nil {
		return nil, checkErr
	}

	if tx.ApproveBaseERC20 {
		approveErr := w.approveERC20(ensureTransactOpts(tx.ApproveBaseAuth), w.baseToken, mintValue, w.sharedL1BridgeAddress)
		if approveErr != nil {
			return nil, approveErr
		}
	}

	preparedTx, err := w.prepareDepositEthToNonEthBasedChain(opts, tx, mintValue)
	if err != nil {
		return nil, err
	}

	if sendErr := w.clientL1.SendTransaction(opts.Context, preparedTx); sendErr != nil {
		return nil, sendErr
	}
	return preparedTx, nil
}

func (w *WalletL1) prepareDepositEthToNonEthBasedChain(opts *TransactOpts, tx *DepositTransaction, mintValue *big.Int) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	if opts.Value == nil {
		opts.Value = tx.Amount
	}
	secondBridgeCalldata, err := w.secondBridgeCalldata(utils.EthAddressInContracts, tx.To, big.NewInt(0))
	if err != nil {
		return nil, err
	}

	secondBridgeAddress := w.sharedL1BridgeAddress
	if tx.BridgeAddress != nil {
		secondBridgeAddress = *tx.BridgeAddress
	}

	noSendOpts := opts.ToTransactOpts(w.auth.From, w.auth.Signer)
	noSendOpts.NoSend = true
	depositTx, err := w.bridgehub.RequestL2TransactionTwoBridges(noSendOpts, bridgehub.L2TransactionRequestTwoBridgesOuter{
		ChainId:                  w.l2ChainId,
		MintValue:                mintValue,
		L2Value:                  big.NewInt(0),
		L2GasLimit:               tx.L2GasLimit,
		L2GasPerPubdataByteLimit: tx.GasPerPubdataByte,
		RefundRecipient:          tx.RefundRecipient,
		SecondBridgeAddress:      secondBridgeAddress,
		SecondBridgeValue:        tx.Amount,
		SecondBridgeCalldata:     secondBridgeCalldata,
	})
	if err != nil {
		return nil, err
	}
	return depositTx, nil
}

func (w *WalletL1) depositBaseTokenToNonEthBasedChain(opts *TransactOpts, tx *DepositTransaction) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	mintValue, err := w.calculateMintValue(opts, tx)
	if err != nil {
		return nil, err
	}

	if tx.ApproveBaseERC20 || tx.ApproveERC20 {
		approveOpts := tx.ApproveBaseAuth
		if tx.ApproveBaseERC20 {
			approveOpts = tx.ApproveAuth
		}
		approveErr := w.approveERC20(ensureTransactOpts(approveOpts), w.baseToken, mintValue, w.sharedL1BridgeAddress)
		if approveErr != nil {
			return nil, approveErr
		}
	}

	preparedTx, err := w.prepareDepositBaseTokenToNonEthBasedChain(opts, tx, mintValue)
	if err != nil {
		return nil, err
	}

	return w.RequestExecute(opts, *preparedTx)
}

func (w *WalletL1) prepareDepositBaseTokenToNonEthBasedChain(opts *TransactOpts, tx *DepositTransaction, mintValue *big.Int) (*RequestExecuteTransaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	opts.Value = big.NewInt(0)

	return &RequestExecuteTransaction{
		ContractAddress:   tx.To,
		L2GasLimit:        tx.L2GasLimit,
		MintValue:         mintValue,
		L2Value:           tx.Amount,
		OperatorTip:       tx.OperatorTip,
		GasPerPubdataByte: tx.GasPerPubdataByte,
		RefundRecipient:   tx.RefundRecipient,
	}, nil
}

func (w *WalletL1) depositNonBaseTokenToNonEthBasedChain(opts *TransactOpts, tx *DepositTransaction) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	baseCost, err := w.calculateBaseCost(opts, tx)
	if err != nil {
		return nil, err
	}
	mintValue, err := w.calculateMintValueFromBaseCost(baseCost, tx)
	if err != nil {
		return nil, err
	}
	checkErr := utils.CheckBaseCost(baseCost, mintValue)
	if checkErr != nil {
		return nil, checkErr
	}

	if tx.ApproveBaseERC20 {
		approveErr := w.approveERC20(ensureTransactOpts(tx.ApproveBaseAuth), w.baseToken, mintValue, w.sharedL1BridgeAddress)
		if approveErr != nil {
			return nil, approveErr
		}
	}

	if tx.ApproveERC20 {
		bridge := w.sharedL1BridgeAddress
		if tx.BridgeAddress != nil {
			bridge = *tx.BridgeAddress
		}
		approveErr := w.approveERC20(ensureTransactOpts(tx.ApproveAuth), tx.Token, tx.Amount, bridge)
		if approveErr != nil {
			return nil, approveErr
		}
	}

	preparedTx, err := w.prepareDepositNonBasedTokenToNonEthBasedChain(opts, tx, mintValue)
	if err != nil {
		return nil, err
	}

	if sendErr := w.clientL1.SendTransaction(opts.Context, preparedTx); sendErr != nil {
		return nil, sendErr
	}
	return preparedTx, nil
}

func (w *WalletL1) prepareDepositNonBasedTokenToNonEthBasedChain(opts *TransactOpts, tx *DepositTransaction, mintValue *big.Int) (*ethTypes.Transaction, error) {
	if err := w.populateDepositTxWithDefaults(opts, tx); err != nil {
		return nil, err
	}
	if opts.Value == nil {
		opts.Value = big.NewInt(0)
	}
	secondBridgeCalldata, err := w.secondBridgeCalldata(tx.Token, tx.To, tx.Amount)
	if err != nil {
		return nil, err
	}

	secondBridgeAddress := w.sharedL1BridgeAddress
	if tx.BridgeAddress != nil {
		secondBridgeAddress = *tx.BridgeAddress
	}

	noSendOpts := opts.ToTransactOpts(w.auth.From, w.auth.Signer)
	noSendOpts.NoSend = true
	depositTx, err := w.bridgehub.RequestL2TransactionTwoBridges(noSendOpts, bridgehub.L2TransactionRequestTwoBridgesOuter{
		ChainId:                  w.l2ChainId,
		MintValue:                mintValue,
		L2Value:                  big.NewInt(0),
		L2GasLimit:               tx.L2GasLimit,
		L2GasPerPubdataByteLimit: tx.GasPerPubdataByte,
		RefundRecipient:          tx.RefundRecipient,
		SecondBridgeAddress:      secondBridgeAddress,
		SecondBridgeValue:        big.NewInt(0),
		SecondBridgeCalldata:     secondBridgeCalldata,
	})
	if err != nil {
		return nil, err
	}
	return depositTx, nil
}

func (w *WalletL1) populateDepositTxWithDefaults(opts *TransactOpts, tx *DepositTransaction) error {
	*opts = *ensureTransactOpts(opts)
	if tx.Token == utils.LegacyEthAddress {
		tx.Token = utils.EthAddressInContracts
	}
	tx.PopulateEmptyFields(w.auth.From)
	if tx.L2GasLimit == nil {
		if err := w.estimateL2GasLimit(opts, tx); err != nil {
			return err
		}
	}

	if err := w.insertGasPriceInTransactOpts(&opts); err != nil {
		return err
	}
	return nil
}

func (w *WalletL1) estimateL2GasLimit(auth *TransactOpts, tx *DepositTransaction) error {
	if tx.BridgeAddress != nil {
		err := w.estimateL2GasLimitFromCustomBridge(auth, tx)
		if err != nil {
			return err
		}
	} else {
		gas, err := w.EstimateDepositL2GasFromDefaultBridge(
			auth.Context,
			tx.Token,
			tx.Amount,
			tx.To, w.auth.From, tx.GasPerPubdataByte)
		if err != nil {
			return err
		}
		tx.L2GasLimit = new(big.Int).SetUint64(gas)
	}
	return nil
}

func (w *WalletL1) estimateL2GasLimitFromCustomBridge(auth *TransactOpts, tx *DepositTransaction) error {
	bridge, err := l1bridge.NewIL1Bridge(*tx.BridgeAddress, w.clientL1)
	if err != nil {
		return fmt.Errorf("failed to load custom bridge: %w", err)
	}
	if tx.CustomBridgeData == nil {
		tx.CustomBridgeData, err = utils.Erc20DefaultBridgeData(tx.Token, w.clientL1)
		if err != nil {
			return err
		}
	}

	l2Address, errBridge := bridge.L2BridgeAddress(&bind.CallOpts{Context: auth.Context, From: w.auth.From}, w.l2ChainId)
	if errBridge != nil {
		return errBridge
	}

	gas, errGas := w.EstimateDepositL2GasFromCustomBridge(auth.Context, *tx.BridgeAddress, l2Address, tx.Token,
		tx.Amount, tx.To, tx.CustomBridgeData, w.auth.From, tx.GasPerPubdataByte, common.Big0)
	if errGas != nil {
		return errGas
	}
	tx.L2GasLimit = new(big.Int).SetUint64(gas)
	return nil
}

func (w *WalletL1) calculateBaseCost(opts *TransactOpts, tx *DepositTransaction) (*big.Int, error) {
	err := w.populateDepositTxWithDefaults(opts, tx)
	if err != nil {
		return nil, err
	}
	gasPriceForEstimation := opts.GasPrice
	if opts.GasFeeCap != nil {
		gasPriceForEstimation = opts.GasFeeCap
	}

	return w.BaseCost(&CallOpts{
		Context: opts.Context,
	}, tx.L2GasLimit, tx.GasPerPubdataByte, gasPriceForEstimation)
}

func (w *WalletL1) calculateMintValue(opts *TransactOpts, tx *DepositTransaction) (*big.Int, error) {
	baseCost, err := w.calculateBaseCost(opts, tx)
	if err != nil {
		return nil, err
	}
	return w.calculateMintValueFromBaseCost(baseCost, tx)

}

func (w *WalletL1) calculateMintValueFromBaseCost(baseCost *big.Int, tx *DepositTransaction) (*big.Int, error) {
	if tx.Token == w.baseToken || (w.isEthBasedChain && tx.Token == utils.EthAddressInContracts) {
		mintValue := new(big.Int).Add(baseCost, tx.OperatorTip)
		mintValue.Add(mintValue, tx.Amount)
		return mintValue, nil
	} else {
		return new(big.Int).Add(baseCost, tx.OperatorTip), nil
	}
}

func (w *WalletL1) secondBridgeCalldata(token, to common.Address, amount *big.Int) ([]byte, error) {
	addressAbiType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}
	uint256AbiType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	return abi.Arguments{
		abi.Argument{Name: "token", Type: addressAbiType},
		abi.Argument{Name: "amount", Type: uint256AbiType},
		abi.Argument{Name: "to", Type: addressAbiType},
	}.Pack(token, amount, to)
}

func (w *WalletL1) approveERC20(auth *TransactOpts, token common.Address, amount *big.Int, bridgeAddress common.Address) error {
	// We only request the allowance if the current one is not enough.
	auth = ensureTransactOpts(auth)
	allowance, err := w.AllowanceL1(&CallOpts{
		Context: auth.Context,
	}, token, bridgeAddress)
	if err != nil {
		return err
	}
	if allowance.Cmp(amount) < 0 {
		approveTx, errApprove := w.ApproveERC20(
			auth, token, amount, bridgeAddress,
		)
		if errApprove != nil {
			return errApprove
		}
		_, err = bind.WaitMined(ensureContext(auth.Context), w.clientL1, approveTx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *WalletL1) prepareRequestExecuteTx(auth *TransactOpts, tx *RequestExecuteTransaction) error {
	opts := ensureTransactOpts(auth)
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
		tx.RefundRecipient = w.auth.From
	}
	if tx.L2GasLimit == nil {
		gas, err := w.clientL2.EstimateL1ToL2Execute(opts.Context, tx.ToCallMsg(w.auth.From, opts))
		if err != nil {
			return err
		}
		tx.L2GasLimit = new(big.Int).SetUint64(gas)
	}
	if tx.MintValue == nil {
		tx.MintValue = big.NewInt(0)
	}

	if err := w.insertGasPriceInTransactOpts(&opts); err != nil {
		return err
	}

	var gasPriceForEstimation *big.Int
	if opts.GasFeeCap != nil {
		gasPriceForEstimation = opts.GasFeeCap
	} else {
		gasPriceForEstimation = opts.GasPrice
	}

	baseCost, err := w.BaseCost(&CallOpts{
		Context: opts.Context,
	}, tx.L2GasLimit, tx.GasPerPubdataByte, gasPriceForEstimation)
	if err != nil {
		return err
	}

	l2Costs := new(big.Int).Add(baseCost, new(big.Int).Add(tx.OperatorTip, tx.L2Value))
	providedValue := tx.MintValue
	if w.isEthBasedChain {
		providedValue = opts.Value
	}
	if providedValue == nil || providedValue.Cmp(big.NewInt(0)) == 0 {
		providedValue = l2Costs
		if w.isEthBasedChain {
			opts.Value = providedValue
		}
	}
	tx.MintValue = providedValue

	err = utils.CheckBaseCost(baseCost, providedValue)
	if err != nil {
		return err
	}
	return nil
}

// Checks if the options contain a GasPrice or GasFeeCap, if not it will try to insert
// the GasFeeCap (maxFeePerGas) and GasTipCap (maxPriorityFee) if chain supports EIP1559,
// otherwise it sets GasPrice
func (w *WalletL1) insertGasPriceInTransactOpts(opts **TransactOpts) error {
	if *opts == nil {
		*opts = ensureTransactOpts(*opts)
	}
	if (*opts).GasFeeCap == nil && (*opts).GasPrice == nil {
		if isReady, head, err := w.checkIfL1ChainIsLondonReady(ensureContext((*opts).Context)); err != nil {
			return err
		} else if isReady {
			if (*opts).GasTipCap == nil {
				gasTipCap, errGasTipCap := w.clientL1.SuggestGasTipCap(ensureContext((*opts).Context))
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
			gasPrice, errGasPrice := w.clientL1.SuggestGasPrice(ensureContext((*opts).Context))
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
func (w *WalletL1) insertGasPriceInDepositMsg(ctx context.Context, msg *DepositCallMsg) error {
	if msg == nil {
		msg = new(DepositCallMsg)
	}
	if msg.GasFeeCap == nil && msg.GasPrice == nil {
		if isReady, head, err := w.checkIfL1ChainIsLondonReady(ensureContext(ctx)); err != nil {
			return err
		} else if isReady {
			if msg.GasTipCap == nil {
				gasTipCap, errGasTipCap := w.clientL1.SuggestGasTipCap(ensureContext(ctx))
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
			gasPrice, errGasPrice := w.clientL1.SuggestGasPrice(ensureContext(ctx))
			if errGasPrice != nil {
				return errGasPrice
			}
			msg.GasPrice = gasPrice
		}
	}
	return nil
}

func (w *WalletL1) getWithdrawalLog(ctx context.Context, withdrawalHash common.Hash, index int) (*types.Log, *big.Int, error) {
	receipt, err := w.clientL2.TransactionReceipt(ctx, withdrawalHash)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return nil, nil, errors.New("transaction receipt not found")
	}
	fLogs := make([]*types.Log, 0)
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

func (w *WalletL1) getWithdrawalL2ToL1Log(ctx context.Context, withdrawalHash common.Hash, index int) (int, *types.L2ToL1Log, error) {
	receipt, err := w.clientL2.TransactionReceipt(ctx, withdrawalHash)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get TransactionReceipt: %w", err)
	}
	if receipt == nil {
		return 0, nil, errors.New("transaction receipt not found")
	}
	fLogs := make([]struct {
		i int
		l *types.L2ToL1Log
	}, 0)
	for i, l := range receipt.L2ToL1Logs {
		if l.Sender == utils.L1MessengerAddress {
			fLogs = append(fLogs, struct {
				i int
				l *types.L2ToL1Log
			}{i, l})
		}
	}
	if len(fLogs) < index+1 {
		return 0, nil, errors.New("withdrawal L2ToL1 log not found")
	}
	return fLogs[index].i, fLogs[index].l, nil
}

func (w *WalletL1) checkIfL1ChainIsLondonReady(ctx context.Context) (bool, *ethTypes.Header, error) {
	// Only query for block header not whole block with transactions
	if head, err := w.clientL1.HeaderByNumber(ensureContext(ctx), nil); err != nil {
		return false, nil, err
	} else if head.BaseFee != nil {
		return true, head, nil
	} else {
		return false, head, nil
	}
}
