package accounts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/l1bridge"
	"github.com/zksync-sdk/zksync2-go/contracts/zksync"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
)

// CallOpts is the collection of options to fine tune a contract call request from
// a specific account associated with AdapterL1.
// Its primary purpose is to be transformed into bind.CallOpts, wherein the 'From'
// field represents the associated account.
type CallOpts struct {
	Pending     bool            // Whether to operate on the pending state or the last known one
	BlockNumber *big.Int        // Optional the block number on which the call should be performed
	Context     context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

func (o *CallOpts) ToCallOpts(from common.Address) *bind.CallOpts {
	return &bind.CallOpts{
		Pending:     o.Pending,
		From:        from,
		BlockNumber: o.BlockNumber,
		Context:     o.Context,
	}
}

// CallMsg contains parameters for contract call from a specific account associated
// with AdapterL1, AdapterL2 or Deployer.
// Its primary purpose is to be transformed into types.CallMsg, wherein the 'From'
// field represents the associated account.
type CallMsg struct {
	To         *common.Address     // The address of the recipient.
	Gas        uint64              // If 0, the call executes with near-infinite gas.
	GasPrice   *big.Int            // Wei <-> gas exchange ratio.
	GasFeeCap  *big.Int            // EIP-1559 fee cap per gas.
	GasTipCap  *big.Int            // EIP-1559 tip per gas.
	Value      *big.Int            // Amount of wei sent along with the call.
	Data       []byte              // Input data, usually an ABI-encoded contract method invocation
	AccessList types.AccessList    // EIP-2930 access list.
	Meta       *zkTypes.Eip712Meta // EIP-712 metadata.
}

func (m *CallMsg) ToCallMsg(from common.Address) zkTypes.CallMsg {
	return zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:       from,
			To:         m.To,
			Gas:        m.Gas,
			GasPrice:   m.GasPrice,
			GasFeeCap:  m.GasFeeCap,
			GasTipCap:  m.GasTipCap,
			Value:      m.Value,
			Data:       m.Data,
			AccessList: m.AccessList,
		},
		Meta: m.Meta,
	}
}

// WithdrawalCallMsg contains the common data required to execute a withdrawal call on L1 from L2.
// This execution is initiated by the account associated with AdapterL2.
type WithdrawalCallMsg struct {
	To            common.Address  // The address of the recipient on L1.
	Amount        *big.Int        // The amount of the token to transfer.
	Token         common.Address  // The address of the token. ETH by default.
	BridgeAddress *common.Address // The address of the bridge contract to be used.

	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	AccessList types.AccessList // EIP-2930 access list.
}

func (m *WithdrawalCallMsg) ToWithdrawalCallMsg(from common.Address) clients.WithdrawalCallMsg {
	return clients.WithdrawalCallMsg{
		To:            m.To,
		Amount:        m.Amount,
		Token:         m.Token,
		BridgeAddress: m.BridgeAddress,
		From:          from,
		Gas:           m.Gas,
		GasPrice:      m.GasPrice,
		GasFeeCap:     m.GasFeeCap,
		GasTipCap:     m.GasTipCap,
		AccessList:    m.AccessList,
	}
}

// TransferCallMsg contains the common data required to execute a transfer call on L2.
// This execution is initiated by the account associated with AdapterL2.
type TransferCallMsg struct {
	To     common.Address // The address of the recipient.
	Amount *big.Int       // The amount of the token to transfer.
	Token  common.Address // The address of the token. ETH by default.

	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	AccessList types.AccessList // EIP-2930 access list.
}

func (m *TransferCallMsg) ToTransferCallMsg(from common.Address) clients.TransferCallMsg {
	return clients.TransferCallMsg{
		To:         m.To,
		Amount:     m.Amount,
		Token:      m.Token,
		From:       from,
		Gas:        m.Gas,
		GasPrice:   m.GasPrice,
		GasFeeCap:  m.GasFeeCap,
		GasTipCap:  m.GasTipCap,
		AccessList: m.AccessList,
	}
}

// DepositCallMsg contains the common data required to execute a deposit call on L2 from L1.
// This execution is initiated by the account associated with AdapterL1.
type DepositCallMsg struct {
	To     common.Address // The address that will receive the deposited tokens on L2.
	Token  common.Address // The address of the token to deposit.
	Amount *big.Int       // The amount of the token to be deposited.

	// If the ETH value passed with the transaction is not explicitly stated Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	// The address of the bridge contract to be used. Defaults to the default zkSync bridge
	// (either L1EthBridge or L1Erc20Bridge).
	BridgeAddress *common.Address

	L2GasLimit *big.Int // Maximum amount of L2 gas that transaction can consume during execution on L2.

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address

	CustomBridgeData []byte // Additional data that can be sent to a bridge.

	Value     *big.Int // The amount of wei sent along with the call.
	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	AccessList types.AccessList // EIP-2930 access list.
}

func (m *DepositCallMsg) ToDepositTransaction() DepositTransaction {
	return DepositTransaction{
		Token:             m.Token,
		Amount:            m.Amount,
		To:                m.To,
		OperatorTip:       m.OperatorTip,
		BridgeAddress:     m.BridgeAddress,
		L2GasLimit:        m.L2GasLimit,
		GasPerPubdataByte: m.GasPerPubdataByte,
		RefundRecipient:   m.RefundRecipient,
		CustomBridgeData:  m.CustomBridgeData,
	}
}

func (m *DepositCallMsg) ToRequestExecuteCallMsg() RequestExecuteCallMsg {
	return RequestExecuteCallMsg{
		ContractAddress:   m.To,
		L2GasLimit:        m.L2GasLimit,
		L2Value:           m.Amount,
		OperatorTip:       m.OperatorTip,
		GasPerPubdataByte: m.GasPerPubdataByte,
		RefundRecipient:   m.RefundRecipient,
		Value:             m.Value,
		Gas:               m.Gas,
		GasPrice:          m.GasPrice,
		GasFeeCap:         m.GasFeeCap,
		GasTipCap:         m.GasTipCap,
		AccessList:        m.AccessList,
	}
}

func (m *DepositCallMsg) ToCallMsg(from, l1Bridge common.Address) (ethereum.CallMsg, error) {
	l1BridgeAbi, err := l1bridge.IL1BridgeMetaData.GetAbi()
	if err != nil {
		return ethereum.CallMsg{}, fmt.Errorf("failed to load IL1Bridge ABI: %w", err)
	}

	calldata, err := l1BridgeAbi.Pack("deposit", m.To, m.Token, m.Amount, m.L2GasLimit, m.GasPerPubdataByte, m.RefundRecipient)
	if err != nil {
		return ethereum.CallMsg{}, err
	}

	var bridge common.Address
	if m.BridgeAddress != nil {
		bridge = *m.BridgeAddress
	} else {
		bridge = l1Bridge
	}

	return ethereum.CallMsg{
		From:      from,
		To:        &bridge,
		Gas:       m.Gas,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		Value:     m.Value,
		Data:      calldata,
	}, nil
}

func (m *DepositCallMsg) ToTransactOpts() TransactOpts {
	return TransactOpts{
		Value:     m.Value,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		GasLimit:  m.Gas,
	}
}

func (m *DepositCallMsg) PopulateEmptyFields(from common.Address) {
	if m.To == (common.Address{}) {
		m.To = from
	}
	if m.OperatorTip == nil {
		m.OperatorTip = big.NewInt(0)
	}
	if m.GasPerPubdataByte == nil {
		m.GasPerPubdataByte = utils.RequiredL1ToL2GasPerPubdataLimit
	}
	if m.Token == (common.Address{}) {
		m.Token = utils.EthAddress
	}
}

// RequestExecuteCallMsg contains the common data required to execute a call for a request execution of an L2
// transaction from L1. This execution is initiated by the account associated with AdapterL1.
type RequestExecuteCallMsg struct {
	ContractAddress common.Address // The L2 receiver address.
	Calldata        []byte         // The input of the L2 transaction.
	L2GasLimit      *big.Int       // Maximum amount of L2 gas that transaction can consume during execution on L2.
	L2Value         *big.Int       // `msg.value` of L2 transaction.
	FactoryDeps     [][]byte       // An array of L2 bytecodes that will be marked as known on L2.

	// If the ETH value passed with the transaction is not explicitly stated Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address

	Value     *big.Int // The amount of wei sent along with the call.
	Gas       uint64   // If 0, the call executes with near-infinite gas.
	GasPrice  *big.Int // Wei <-> gas exchange ratio.
	GasFeeCap *big.Int // EIP-1559 fee cap per gas.
	GasTipCap *big.Int // EIP-1559 tip per gas.

	AccessList types.AccessList // EIP-2930 access list.
}

func (m *RequestExecuteCallMsg) ToRequestExecuteTransaction() RequestExecuteTransaction {
	return RequestExecuteTransaction{
		ContractAddress:   m.ContractAddress,
		Calldata:          m.Calldata,
		L2GasLimit:        m.L2GasLimit,
		L2Value:           m.L2Value,
		FactoryDeps:       m.FactoryDeps,
		OperatorTip:       m.OperatorTip,
		GasPerPubdataByte: m.GasPerPubdataByte,
		RefundRecipient:   m.RefundRecipient,
	}
}

func (m *RequestExecuteCallMsg) ToCallMsg(from common.Address) (ethereum.CallMsg, error) {
	zksyncAbi, err := zksync.IZkSyncMetaData.GetAbi()
	if err != nil {
		return ethereum.CallMsg{}, fmt.Errorf("failed to load IZkSync ABI: %w", err)
	}
	requestExecuteCalldata, err := zksyncAbi.Pack("requestL2Transaction", m.ContractAddress, m.L2Value,
		m.Calldata, m.L2GasLimit, m.GasPerPubdataByte, m.FactoryDeps, m.RefundRecipient)
	if err != nil {
		return ethereum.CallMsg{}, err
	}

	return ethereum.CallMsg{
		From:      from,
		To:        &m.ContractAddress,
		Gas:       m.Gas,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		Value:     m.Value,
		Data:      requestExecuteCalldata,
	}, nil
}

func (m *RequestExecuteCallMsg) ToTransactOpts() TransactOpts {
	return TransactOpts{
		Value:     m.Value,
		GasPrice:  m.GasPrice,
		GasFeeCap: m.GasFeeCap,
		GasTipCap: m.GasTipCap,
		GasLimit:  m.Gas,
	}
}

// TransactOpts contains common data required to create a valid transaction on both L1 and L2
// using the account associated with AdapterL1, AdapterL2 or Deployer.
// Its primary purpose is to be transformed into bind.TransactOpts, wherein the 'From' and 'Signer'
// fields are linked to the associated account.
type TransactOpts struct {
	Nonce     *big.Int        // Nonce to use for the transaction execution (nil = use pending state).
	Value     *big.Int        // Funds to transfer along the transaction (nil = 0 = no funds).
	GasPrice  *big.Int        // Gas price to use for the transaction execution (nil = gas price oracle).
	GasFeeCap *big.Int        // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle).
	GasTipCap *big.Int        // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle).
	GasLimit  uint64          // Gas limit to set for the transaction execution (0 = estimate).
	Context   context.Context // Network context to support cancellation and timeouts (nil = no timeout).
}

func (t *TransactOpts) ToTransactOpts(from common.Address, signer bind.SignerFn) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:      from,
		Nonce:     t.Nonce,
		Signer:    signer,
		Value:     t.Value,
		GasPrice:  t.GasPrice,
		GasFeeCap: t.GasFeeCap,
		GasTipCap: t.GasTipCap,
		GasLimit:  t.GasLimit,
		Context:   t.Context,
	}
}

// Transaction is similar to types.Transaction712 but does not include the From field. This design is intended for use
// with AdapterL1, AdapterL2, or Deployer, which already have an associated account. The From field is bound to
// their associated account, and thus, it is not included in this type.
type Transaction struct {
	To        *common.Address // The address of the recipient.
	Data      hexutil.Bytes   // Input data, usually an ABI-encoded contract method invocation.
	Value     *big.Int        // Funds to transfer along the transaction (nil = 0 = no funds).
	Nonce     *big.Int        // Nonce to use for the transaction execution.
	GasTipCap *big.Int        // EIP-1559 tip per gas.
	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
	Gas       uint64          // Gas limit to set for the transaction execution.

	AccessList types.AccessList // EIP-2930 access list.

	ChainID *big.Int            // Chain ID of the network.
	Meta    *zkTypes.Eip712Meta // EIP-712 metadata.
}

func (t *Transaction) ToTransaction712(from common.Address) *zkTypes.Transaction712 {
	return &zkTypes.Transaction712{
		Nonce:      t.Nonce,
		GasTipCap:  t.GasTipCap,
		GasFeeCap:  t.GasFeeCap,
		Gas:        new(big.Int).SetUint64(t.Gas),
		To:         t.To,
		Value:      t.Value,
		Data:       t.Data,
		AccessList: t.AccessList,
		ChainID:    t.ChainID,
		From:       &from,
		Meta:       t.Meta,
	}
}

func (t *Transaction) ToCallMsg(from common.Address) zkTypes.CallMsg {
	return zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:       from,
			To:         t.To,
			Gas:        t.Gas,
			GasFeeCap:  t.GasFeeCap,
			GasTipCap:  t.GasTipCap,
			Value:      t.Value,
			Data:       t.Data,
			AccessList: t.AccessList,
		},
		Meta: t.Meta,
	}
}

// TransferTransaction represents a transfer transaction on L2
// initiated by the account associated with AdapterL2.
type TransferTransaction struct {
	To     common.Address // The address of the recipient.
	Amount *big.Int       // The amount of the token to transfer.
	Token  common.Address // The address of the token. ETH by default.
}

func (t *TransferTransaction) ToTransaction(opts *TransactOpts) *Transaction {
	return &Transaction{
		To:        &t.To,
		Value:     t.Amount,
		Nonce:     opts.Nonce,
		GasFeeCap: opts.GasFeeCap,
		GasTipCap: opts.GasTipCap,
		Gas:       opts.GasLimit,
	}
}

func (t *TransferTransaction) ToTransferCallMsg(from common.Address, opts *TransactOpts) clients.TransferCallMsg {
	return clients.TransferCallMsg{
		To:        t.To,
		Amount:    t.Amount,
		Token:     t.Token,
		From:      from,
		Gas:       opts.GasLimit,
		GasPrice:  opts.GasPrice,
		GasFeeCap: opts.GasFeeCap,
		GasTipCap: opts.GasTipCap,
	}
}

// WithdrawalTransaction represents a withdrawal transaction on L1 from L2
// initiated by the account associated with AdapterL2.
type WithdrawalTransaction struct {
	To     common.Address // The address that will receive the withdrawn tokens on L1.
	Token  common.Address // The address of the token to withdraw.
	Amount *big.Int       // The amount of the token to withdraw.

	// The address of the bridge contract to be used. Defaults to the default zkSync bridge
	// (either L2EthBridge or L2Erc20Bridge).
	BridgeAddress *common.Address
}

func (t *WithdrawalTransaction) ToWithdrawalCallMsg(from common.Address, opts *TransactOpts) *clients.WithdrawalCallMsg {
	return &clients.WithdrawalCallMsg{
		To:            t.To,
		Amount:        t.Amount,
		Token:         t.Token,
		BridgeAddress: t.BridgeAddress,
		From:          from,
		Gas:           opts.GasLimit,
		GasPrice:      opts.GasPrice,
		GasFeeCap:     opts.GasFeeCap,
		GasTipCap:     opts.GasTipCap,
	}
}

// RequestExecuteTransaction represents a request execution of L2 transaction from L1
// initiated by the account associated with AdapterL1.
type RequestExecuteTransaction struct {
	ContractAddress common.Address // The L2 receiver address.
	Calldata        []byte         // The input of the L2 transaction.
	L2GasLimit      *big.Int       // Maximum amount of L2 gas that transaction can consume during execution on L2.
	L2Value         *big.Int       // `msg.value` of L2 transaction.
	FactoryDeps     [][]byte       // An array of L2 bytecodes that will be marked as known on L2.

	// If the ETH value passed with the transaction is not explicitly stated Auth.Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address
}

func (t *RequestExecuteTransaction) ToRequestExecuteCallMsg(opts *TransactOpts) RequestExecuteCallMsg {
	return RequestExecuteCallMsg{
		ContractAddress:   t.ContractAddress,
		Calldata:          t.Calldata,
		L2GasLimit:        t.L2GasLimit,
		L2Value:           t.L2Value,
		FactoryDeps:       t.FactoryDeps,
		OperatorTip:       t.OperatorTip,
		GasPerPubdataByte: t.GasPerPubdataByte,
		RefundRecipient:   t.RefundRecipient,
		Value:             opts.Value,
		GasPrice:          opts.GasPrice,
		GasFeeCap:         opts.GasFeeCap,
		GasTipCap:         opts.GasTipCap,
	}
}

func (t *RequestExecuteTransaction) ToCallMsg(from common.Address, opts *TransactOpts) zkTypes.CallMsg {
	factoryDeps := make([]hexutil.Bytes, len(t.FactoryDeps))
	if len(t.FactoryDeps) > 0 {
		for i, d := range t.FactoryDeps {
			factoryDeps[i] = d
		}
	}
	return zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:     from,
			To:       &t.ContractAddress,
			Gas:      opts.GasLimit,
			GasPrice: opts.GasPrice,
			Value:    opts.Value,
			Data:     t.Calldata,
		},
		Meta: &zkTypes.Eip712Meta{
			GasPerPubdata: utils.NewBig(t.GasPerPubdataByte.Int64()),
			FactoryDeps:   factoryDeps,
		},
	}
}

// DepositTransaction represents a deposit transaction on L2 from L1
// initiated by the account associated with AdapterL1.
type DepositTransaction struct {
	To     common.Address // The address of the token to deposit.
	Token  common.Address // The address of the token to deposit.
	Amount *big.Int       // The amount of the token to be deposited.

	// If the ETH value passed with the transaction is not explicitly stated Auth.Value,
	// this field will be equal to the tip the operator will receive on top of the base cost
	// of the transaction.
	OperatorTip *big.Int

	// The address of the bridge contract to be used. Defaults to the default zkSync bridge
	// (either L1EthBridge or L1Erc20Bridge).
	BridgeAddress *common.Address

	// Whether should the token approval be performed under the hood. Set this flag to true if you
	// bridge an ERC20 token and didn't call the approveERC20 function beforehand.
	ApproveERC20 bool

	L2GasLimit *big.Int // Maximum amount of L2 gas that transaction can consume during execution on L2.

	// The maximum amount L2 gas that the operator may charge the user for single byte of pubdata.
	GasPerPubdataByte *big.Int

	// The address on L2 that will receive the refund for the transaction.
	// If the transaction fails, it will also be the address to receive L2Value.
	RefundRecipient common.Address

	CustomBridgeData []byte // Additional data that can be sent to a bridge.

	ApproveAuth *TransactOpts // Authorization data for the approval token transaction.
}

func (t *DepositTransaction) ToRequestExecuteTransaction() *RequestExecuteTransaction {
	return &RequestExecuteTransaction{
		ContractAddress:   t.To,
		L2Value:           t.Amount,
		L2GasLimit:        t.L2GasLimit,
		GasPerPubdataByte: t.GasPerPubdataByte,
		RefundRecipient:   t.RefundRecipient,
	}
}

func (t *DepositTransaction) ToDepositCallMsg(opts *TransactOpts) DepositCallMsg {
	return DepositCallMsg{
		To:                t.To,
		Token:             t.Token,
		Amount:            t.Amount,
		OperatorTip:       t.OperatorTip,
		BridgeAddress:     t.BridgeAddress,
		L2GasLimit:        t.L2GasLimit,
		GasPerPubdataByte: t.GasPerPubdataByte,
		RefundRecipient:   t.RefundRecipient,
		CustomBridgeData:  t.CustomBridgeData,
		Value:             opts.Value,
		GasPrice:          opts.GasPrice,
		GasFeeCap:         opts.GasFeeCap,
		GasTipCap:         opts.GasTipCap,
	}
}

func (t *DepositTransaction) PopulateEmptyFields(from common.Address) {
	if t.To == (common.Address{}) {
		t.To = from
	}
	if t.OperatorTip == nil {
		t.OperatorTip = big.NewInt(0)
	}
	if t.GasPerPubdataByte == nil {
		t.GasPerPubdataByte = utils.RequiredL1ToL2GasPerPubdataLimit
	}
	if t.Token == (common.Address{}) {
		t.Token = utils.EthAddress
	}
	if t.ApproveERC20 && t.ApproveAuth == nil {
		t.ApproveAuth = ensureTransactOpts(t.ApproveAuth)
	}
}

// DeploymentType represents an enumeration of deployment types.
type DeploymentType string

const (
	DeployContract DeploymentType = "CONTRACT"
	DeployAccount  DeploymentType = "ACCOUNT"
)

// CreateTransaction represents the parameters for deploying a contract or account using the CREATE opcode.
// This transaction is initiated by the account associated with Deployer.
type CreateTransaction struct {
	Bytecode     []byte   // The bytecode of smart contract or smart account.
	Calldata     []byte   // The constructor calldata.
	Dependencies [][]byte // The bytecode of dependent smart contracts or smart accounts.
}

func (t *CreateTransaction) ToTransaction(deploymentType DeploymentType, opts *TransactOpts) (*Transaction, error) {
	var (
		data []byte
		err  error
	)
	if deploymentType == DeployContract {
		data, err = utils.EncodeCreate(t.Bytecode, t.Calldata)
		if err != nil {
			return nil, fmt.Errorf("failed to encode create call: %w", err)
		}
	} else {
		data, err = utils.EncodeCreateAccount(t.Bytecode, t.Calldata, zkTypes.Version1)
		if err != nil {
			return nil, fmt.Errorf("failed to encode createAccount call: %w", err)
		}

	}

	var factoryDeps []hexutil.Bytes
	if len(t.Dependencies) > 0 {
		factoryDeps = make([]hexutil.Bytes, len(t.Dependencies)+1)
		for i, d := range t.Dependencies {
			factoryDeps[i] = d
		}
		factoryDeps[len(t.Dependencies)] = t.Bytecode
	} else {
		factoryDeps = []hexutil.Bytes{t.Bytecode}
	}

	auth := opts
	if auth == nil {
		auth = &TransactOpts{Context: context.Background()}
	}

	return &Transaction{
		To:        &utils.ContractDeployerAddress,
		Data:      data,
		Nonce:     auth.Nonce,
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
		Gas:       auth.GasLimit,
		Meta: &zkTypes.Eip712Meta{
			GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
			FactoryDeps:   factoryDeps,
		},
	}, nil
}

// Create2Transaction represents the parameters for deploying a contract or account using the CREATE2 opcode.
// This transaction is initiated by the account associated with Deployer.
type Create2Transaction struct {
	Bytecode     []byte   // The bytecode of smart contract or smart account.
	Calldata     []byte   // The constructor calldata.
	Salt         []byte   // The create2 salt.
	Dependencies [][]byte // The bytecode of dependent smart contracts or smart accounts.
}

func (t *Create2Transaction) ToTransaction(deploymentType DeploymentType, opts *TransactOpts) (*Transaction, error) {
	var (
		data []byte
		err  error
	)

	if deploymentType == DeployContract {
		data, err = utils.EncodeCreate2(t.Bytecode, t.Calldata, t.Salt)
		if err != nil {
			return nil, fmt.Errorf("failed to encode create2 call: %w", err)
		}
	} else {
		data, err = utils.EncodeCreate2Account(t.Bytecode, t.Calldata, t.Salt, zkTypes.Version1)
		if err != nil {
			return nil, fmt.Errorf("failed to encode create2Account call: %w", err)
		}
	}

	var factoryDeps []hexutil.Bytes
	if len(t.Dependencies) > 0 {
		factoryDeps = make([]hexutil.Bytes, len(t.Dependencies)+1)
		for i, d := range t.Dependencies {
			factoryDeps[i] = d
		}
		factoryDeps[len(t.Dependencies)] = t.Bytecode
	} else {
		factoryDeps = []hexutil.Bytes{t.Bytecode}
	}

	auth := opts
	if auth == nil {
		auth = &TransactOpts{Context: context.Background()}
	}

	return &Transaction{
		To:        &utils.ContractDeployerAddress,
		Data:      data,
		Nonce:     auth.Nonce,
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
		Gas:       auth.GasLimit,
		Meta: &zkTypes.Eip712Meta{
			GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
			FactoryDeps:   factoryDeps,
		},
	}, nil
}

// FullDepositFee represents the total ETH fee required for performing the deposit on
// both L1 and L2 networks.
type FullDepositFee struct {
	MaxFeePerGas, // MaxFeePerGas of the L1 transaction if 1559 transaction is used.
	MaxPriorityFeePerGas, // MaxPriorityFeePerGas of the L1 transaction if 1559 transaction is used.
	GasPrice, // Gas price of the L1 transaction if legacy transaction is used.
	BaseCost, // Base cost of the L2 transaction.
	L1GasLimit, // Gas limit of the L1 transaction.
	L2GasLimit *big.Int // Gas limit of the L2 transaction.
}
