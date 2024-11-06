package accounts

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/ethtoken"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/nonceholder"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
)

// WalletL2 associated with an account and provides common operations on the
// L2 network for the associated account.
type WalletL2 struct {
	client *clients.Client
	signer *ECDSASigner
	auth   *bind.TransactOpts
	cache  *Cache
}

// NewWalletL2 creates an instance of WalletL2 associated with the account provided by the raw private key.
func NewWalletL2(rawPrivateKey []byte, client *clients.Client) (*WalletL2, error) {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	signer, err := NewECDSASignerFromRawPrivateKey(rawPrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	return NewWalletL2FromSigner(signer, client)
}

// NewWalletL2FromSigner creates an instance of WalletL2.
// The client can be optional and if it is not provided, only WalletL2.SignTransaction,
// WalletL2.Address, WalletL2.Signer can be performed, as the rest of the
// functionalities require communication to the network.
func NewWalletL2FromSigner(signer *ECDSASigner, client *clients.Client) (*WalletL2, error) {
	if client == nil {
		return &WalletL2{signer: signer}, nil
	}
	chainId, err := client.ChainID(context.Background())
	auth, err := newTransactorWithSigner(signer, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to init TransactOptsL1: %w", err)
	}

	return &WalletL2{
		client: client,
		signer: signer,
		auth:   auth,
		cache:  NewCache(client, nil),
	}, nil
}

// NewWalletL2FromSignerAndCache creates an instance of WalletL2 with cache.
// The cache is optional and if it is not provided, new empty cache is used.
// The clientL2 can be optional and if it is not provided, only WalletL2.SignTransaction,
// WalletL2.Address, WalletL2.Signer can be performed, as the rest of the
// functionalities require communication to the network.
func NewWalletL2FromSignerAndCache(signer *ECDSASigner, client *clients.Client, cache *Cache) (*WalletL2, error) {
	if client == nil {
		return &WalletL2{signer: signer}, nil
	}
	chainId, err := client.ChainID(context.Background())
	auth, err := newTransactorWithSigner(signer, chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to init TransactOptsL1: %w", err)
	}

	c := cache
	if c == nil {
		c = NewCache(client, nil)
	}

	return &WalletL2{
		client: client,
		signer: signer,
		auth:   auth,
		cache:  c,
	}, nil
}

// Address returns the address of the associated account.
func (w *WalletL2) Address() common.Address {
	return w.auth.From
}

// Signer returns the signer of the associated account.
func (w *WalletL2) Signer() *ECDSASigner {
	return w.signer
}

// Balance returns the balance of the specified token that can be either base token or any ERC20 token.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (w *WalletL2) Balance(opts *CallOpts, token common.Address) (*big.Int, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(w.auth.From)
	if token == utils.LegacyEthAddress || token == utils.EthAddressInContracts {
		var err error
		token, err = w.client.L2TokenAddress(callOpts.Context, token)
		if err != nil {
			return nil, err
		}
	}

	if token == utils.L2BaseTokenAddress {
		return w.client.BalanceAt(callOpts.Context, w.Address(), callOpts.BlockNumber)
	}
	erc20Token, err := erc20.NewIERC20(token, w.client)
	if err != nil {
		return nil, err
	}
	return erc20Token.BalanceOf(callOpts, w.Address())
}

// AllBalances returns all balances for confirmed tokens given by an associated account.
func (w *WalletL2) AllBalances(ctx context.Context) (map[common.Address]*big.Int, error) {
	return w.client.AllAccountBalances(ensureContext(ctx), w.Address())
}

// L2BridgeContracts returns L2 bridge contracts.
func (w *WalletL2) L2BridgeContracts(_ context.Context) (*types.L2BridgeContracts, error) {
	return w.cache.L2BridgeContracts()
}

// DeploymentNonce returns the deployment nonce of the account.
func (w *WalletL2) DeploymentNonce(opts *CallOpts) (*big.Int, error) {
	callOpts := ensureCallOpts(opts).ToCallOpts(w.auth.From)
	nonceHolder, err := nonceholder.NewINonceHolder(utils.NonceHolderAddress, w.client)
	if err != nil {
		return nil, err
	}
	return nonceHolder.GetDeploymentNonce(callOpts, w.Address())
}

// IsBaseToken returns whether the token is the base token.
func (w *WalletL2) IsBaseToken(ctx context.Context, token common.Address) (bool, error) {
	return w.client.IsBaseToken(ensureContext(ctx), token)
}

// Withdraw initiates the withdrawal process which withdraws ETH or any ERC20
// token from the associated account on L2 network to the target account on L1
// network.
func (w *WalletL2) Withdraw(auth *TransactOpts, tx WithdrawalTransaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	insertGasPrice(opts)

	if tx.Token == utils.LegacyEthAddress || tx.Token == utils.EthAddressInContracts {
		var err error
		tx.Token, err = w.client.L2TokenAddress(opts.Context, tx.Token)
		if err != nil {
			return common.Hash{}, err
		}
	}

	if tx.Token == utils.L2BaseTokenAddress {
		if opts.Value != nil && opts.Value != tx.Amount {
			return common.Hash{}, errors.New("the tx.value is not equal to the value withdrawn")
		} else {
			opts.Value = tx.Amount
		}

		abi, abiErr := ethtoken.IEthTokenMetaData.GetAbi()
		if abiErr != nil {
			return common.Hash{}, abiErr
		}

		data, packErr := abi.Pack("withdraw", tx.To)
		if packErr != nil {
			return common.Hash{}, packErr
		}

		return w.SendTransaction(opts.Context, &Transaction{
			Nonce:           opts.Nonce,
			GasTipCap:       opts.GasTipCap,
			GasFeeCap:       opts.GasFeeCap,
			Gas:             opts.GasLimit,
			To:              &utils.L2BaseTokenAddress,
			Value:           opts.Value,
			Data:            data,
			ChainID:         w.signer.ChainID(),
			GasPerPubdata:   utils.DefaultGasPerPubdataLimit,
			PaymasterParams: opts.PaymasterParams,
		})
	}

	if tx.BridgeAddress == nil {
		sharedL2BridgeAddress, err := w.cache.L2SharedBridgeAddress()
		if err != nil {
			return common.Hash{}, err
		}
		tx.BridgeAddress = &sharedL2BridgeAddress
	}

	abi, abiErr := l2bridge.IL2BridgeMetaData.GetAbi()
	if abiErr != nil {
		return common.Hash{}, abiErr
	}

	data, abiPack := abi.Pack("withdraw", tx.To, tx.Token, tx.Amount)
	if abiPack != nil {
		return common.Hash{}, abiPack
	}

	return w.SendTransaction(opts.Context, &Transaction{
		Nonce:           opts.Nonce,
		GasTipCap:       opts.GasTipCap,
		GasFeeCap:       opts.GasFeeCap,
		Gas:             opts.GasLimit,
		To:              tx.BridgeAddress,
		Value:           opts.Value,
		Data:            data,
		ChainID:         w.signer.ChainID(),
		GasPerPubdata:   utils.DefaultGasPerPubdataLimit,
		PaymasterParams: opts.PaymasterParams,
	})
}

// EstimateGasWithdraw estimates the amount of gas required for a withdrawal transaction.
func (w *WalletL2) EstimateGasWithdraw(ctx context.Context, msg WithdrawalCallMsg) (uint64, error) {
	return w.client.EstimateGasWithdraw(ensureContext(ctx), msg.ToWithdrawalCallMsg(w.Address()))
}

// Transfer moves the base token or any ERC20 token from the associated account to the target account.
func (w *WalletL2) Transfer(auth *TransactOpts, tx TransferTransaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	insertGasPrice(opts)

	if tx.Token == utils.LegacyEthAddress || tx.Token == utils.EthAddressInContracts {
		var err error
		tx.Token, err = w.client.L2TokenAddress(opts.Context, tx.Token)
		if err != nil {
			return common.Hash{}, err
		}
	}

	if tx.Token == utils.L2BaseTokenAddress {
		return w.SendTransaction(opts.Context, &Transaction{
			Nonce:           opts.Nonce,
			GasTipCap:       opts.GasTipCap,
			GasFeeCap:       opts.GasFeeCap,
			Gas:             opts.GasLimit,
			To:              &tx.To,
			Value:           tx.Amount,
			ChainID:         w.signer.ChainID(),
			GasPerPubdata:   utils.DefaultGasPerPubdataLimit,
			PaymasterParams: opts.PaymasterParams,
		})
	}

	abi, err := erc20.IERC20MetaData.GetAbi()
	if err != nil {
		return common.Hash{}, err
	}

	data, err := abi.Pack("transfer", tx.To, tx.Amount)
	if err != nil {
		return common.Hash{}, err
	}

	return w.SendTransaction(opts.Context, &Transaction{
		Nonce:           opts.Nonce,
		GasTipCap:       opts.GasTipCap,
		GasFeeCap:       opts.GasFeeCap,
		Gas:             opts.GasLimit,
		To:              &tx.Token,
		Value:           big.NewInt(0),
		Data:            data,
		ChainID:         w.signer.ChainID(),
		GasPerPubdata:   utils.DefaultGasPerPubdataLimit,
		PaymasterParams: opts.PaymasterParams,
	})
}

// EstimateGasTransfer estimates the amount of gas required for a transfer transaction.
func (w *WalletL2) EstimateGasTransfer(ctx context.Context, msg TransferCallMsg) (uint64, error) {
	return w.client.EstimateGasTransfer(ensureContext(ctx), msg.ToTransferCallMsg(w.Address()))
}

// CallContract executes a message call for EIP-712 transaction, which is
// directly executed in the VM of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in
// which case the code is taken from the latest known block. Note that state from
// very old blocks might not be available.
func (w *WalletL2) CallContract(ctx context.Context, msg CallMsg, blockNumber *big.Int) ([]byte, error) {
	return w.client.CallContractL2(ensureContext(ctx), msg.ToCallMsg(w.Address()), blockNumber)
}

// PopulateTransaction is designed for users who prefer a simplified approach by
// providing only the necessary data to create a valid transaction. The only
// required fields are Transaction.To and either Transaction.Data or
// Transaction.Value (or both, if the method is payable). Any other fields that
// are not set will be prepared by this method.
func (w *WalletL2) PopulateTransaction(ctx context.Context, tx Transaction) (*types.Transaction, error) {
	if tx.ChainID == nil {
		tx.ChainID = w.signer.ChainID()
	}
	if tx.Nonce == nil {
		nonce, err := w.client.NonceAt(ensureContext(ctx), w.Address(), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce: %w", err)
		}
		tx.Nonce = new(big.Int).SetUint64(nonce)
	}
	if tx.GasFeeCap == nil {
		gasFeeCap, err := w.client.SuggestGasPrice(ensureContext(ctx))
		if err != nil {
			return nil, fmt.Errorf("failed to SuggestGasPrice: %w", err)
		}
		tx.GasFeeCap = gasFeeCap
	}
	if tx.GasTipCap == nil {
		tx.GasTipCap = big.NewInt(0)
	}
	if tx.GasPerPubdata == nil {
		tx.GasPerPubdata = utils.DefaultGasPerPubdataLimit
	}
	if tx.Gas == 0 || tx.GasFeeCap == nil {
		fee, err := w.client.EstimateFee(ensureContext(ctx), tx.ToCallMsg(w.Address()))
		if err != nil {
			return nil, fmt.Errorf("failed to EstimateFee: %w", err)
		}
		if tx.Gas == 0 {
			tx.Gas = fee.GasLimit.ToInt().Uint64()
		}
		if tx.GasFeeCap == nil || tx.GasFeeCap.Uint64() == 0 {
			tx.GasFeeCap = fee.MaxFeePerGas.ToInt()
		}

	}
	if tx.Data == nil {
		tx.Data = hexutil.Bytes{}
	}
	return tx.ToTransaction(w.auth.From), nil
}

// SignTransaction returns a signed transaction that is ready to be broadcast to
// the network. The input transaction must be a valid transaction with all fields
// having appropriate values. To obtain a valid transaction, you can use the
// PopulateTransaction method.
func (w *WalletL2) SignTransaction(tx *types.Transaction) ([]byte, error) {
	var gas uint64 = 0
	if tx.Gas != nil {
		gas = tx.Gas.Uint64()
	}
	preparedTx, err := w.PopulateTransaction(context.Background(), Transaction{
		To:              tx.To,
		Data:            tx.Data,
		Value:           tx.Value,
		Nonce:           tx.Nonce,
		GasTipCap:       tx.GasTipCap,
		GasFeeCap:       tx.GasFeeCap,
		Gas:             gas,
		ChainID:         tx.ChainID,
		GasPerPubdata:   tx.GasPerPubdata,
		CustomSignature: tx.CustomSignature,
		FactoryDeps:     tx.FactoryDeps,
		PaymasterParams: tx.PaymasterParams,
	})
	if err != nil {
		return nil, err
	}

	typedData, err := preparedTx.TypedData()
	if err != nil {
		return nil, err
	}
	signature, err := w.signer.SignTypedData(context.Background(), typedData)
	if err != nil {
		return nil, err
	}
	return preparedTx.Encode(signature)
}

// SendTransaction injects a transaction into the pending pool for execution. Any
// unset transaction fields are prepared using the PopulateTransaction method.
func (w *WalletL2) SendTransaction(ctx context.Context, tx *Transaction) (common.Hash, error) {
	preparedTx, err := w.PopulateTransaction(ensureContext(ctx), *tx)
	if err != nil {
		return common.Hash{}, err
	}
	rawTx, err := w.SignTransaction(preparedTx)
	if err != nil {
		return common.Hash{}, err
	}
	return w.client.SendRawTransaction(ensureContext(ctx), rawTx)
}

func (w *WalletL2) transferBaseToken(auth *TransactOpts, tx TransferTransaction) (*ethTypes.Transaction, error) {
	if auth.GasPrice != nil {
		if auth.Nonce == nil {
			nonce, err := w.client.NonceAt(auth.Context, w.Address(), nil)
			if err != nil {
				return nil, fmt.Errorf("failed to get nonce: %w", err)
			}
			auth.Nonce = new(big.Int).SetUint64(nonce)
		}

		transaction := ethTypes.NewTx(
			&ethTypes.LegacyTx{
				Nonce:    auth.Nonce.Uint64(),
				GasPrice: auth.GasPrice,
				Gas:      auth.GasLimit,
				To:       &tx.To,
				Value:    tx.Amount,
			})

		signedTx, _ := ethTypes.SignTx(transaction, ethTypes.NewLondonSigner(w.signer.ChainID()), w.signer.PrivateKey())
		err := w.client.SendTransaction(auth.Context, signedTx)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	} else {
		preparedTx, err := w.PopulateTransaction(auth.Context, *tx.ToTransaction(auth))
		if err != nil {
			return nil, err
		}
		transaction := ethTypes.NewTx(
			&ethTypes.DynamicFeeTx{
				ChainID:   w.signer.ChainID(),
				Nonce:     preparedTx.Nonce.Uint64(),
				GasTipCap: preparedTx.GasTipCap,
				GasFeeCap: preparedTx.GasFeeCap,
				Gas:       preparedTx.Gas.Uint64(),
				To:        preparedTx.To,
				Value:     preparedTx.Value,
			})
		signedTx, _ := ethTypes.SignTx(transaction, ethTypes.NewLondonSigner(w.signer.ChainID()), w.signer.PrivateKey())
		err = w.client.SendTransaction(auth.Context, signedTx)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	}
}
