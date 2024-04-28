// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1bridge

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// L2TransactionRequestTwoBridgesInner is an auto generated low-level Go binding around an user-defined struct.
type L2TransactionRequestTwoBridgesInner struct {
	MagicValue  [32]byte
	L2Contract  common.Address
	L2Calldata  []byte
	FactoryDeps [][]byte
	TxDataHash  [32]byte
}

// IL1BridgeMetaData contains all meta data concerning the IL1Bridge contract.
var IL1BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"}],\"name\":\"BridgehubDepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositInitiatedSharedBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDepositSharedBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiatedSharedBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalizedSharedBridge\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgehub\",\"outputs\":[{\"internalType\":\"contractIBridgehub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"bridgehubConfirmL2Transaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prevMsgSender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"bridgehubDeposit\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"magicValue\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Contract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"l2Calldata\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structL2TransactionRequestTwoBridgesInner\",\"name\":\"request\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prevMsgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bridgehubDepositBaseToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"}],\"name\":\"depositHappened\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalizedShared\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"l2BridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IL1BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1BridgeMetaData.ABI instead.
var IL1BridgeABI = IL1BridgeMetaData.ABI

// IL1Bridge is an auto generated Go binding around an Ethereum contract.
type IL1Bridge struct {
	IL1BridgeCaller     // Read-only binding to the contract
	IL1BridgeTransactor // Write-only binding to the contract
	IL1BridgeFilterer   // Log filterer for contract events
}

// IL1BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1BridgeSession struct {
	Contract     *IL1Bridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1BridgeCallerSession struct {
	Contract *IL1BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IL1BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1BridgeTransactorSession struct {
	Contract     *IL1BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IL1BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1BridgeRaw struct {
	Contract *IL1Bridge // Generic contract binding to access the raw methods on
}

// IL1BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1BridgeCallerRaw struct {
	Contract *IL1BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IL1BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1BridgeTransactorRaw struct {
	Contract *IL1BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1Bridge creates a new instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1Bridge(address common.Address, backend bind.ContractBackend) (*IL1Bridge, error) {
	contract, err := bindIL1Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1Bridge{IL1BridgeCaller: IL1BridgeCaller{contract: contract}, IL1BridgeTransactor: IL1BridgeTransactor{contract: contract}, IL1BridgeFilterer: IL1BridgeFilterer{contract: contract}}, nil
}

// NewIL1BridgeCaller creates a new read-only instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1BridgeCaller(address common.Address, caller bind.ContractCaller) (*IL1BridgeCaller, error) {
	contract, err := bindIL1Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeCaller{contract: contract}, nil
}

// NewIL1BridgeTransactor creates a new write-only instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1BridgeTransactor, error) {
	contract, err := bindIL1Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeTransactor{contract: contract}, nil
}

// NewIL1BridgeFilterer creates a new log filterer instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1BridgeFilterer, error) {
	contract, err := bindIL1Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeFilterer{contract: contract}, nil
}

// bindIL1Bridge binds a generic wrapper to an already deployed contract.
func bindIL1Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Bridge *IL1BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Bridge.Contract.IL1BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Bridge *IL1BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Bridge.Contract.IL1BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Bridge *IL1BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Bridge.Contract.IL1BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Bridge *IL1BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Bridge *IL1BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Bridge *IL1BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Bridge.Contract.contract.Transact(opts, method, params...)
}

// Bridgehub is a free data retrieval call binding the contract method 0x5d38962d.
//
// Solidity: function bridgehub() view returns(address)
func (_IL1Bridge *IL1BridgeCaller) Bridgehub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Bridge.contract.Call(opts, &out, "bridgehub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridgehub is a free data retrieval call binding the contract method 0x5d38962d.
//
// Solidity: function bridgehub() view returns(address)
func (_IL1Bridge *IL1BridgeSession) Bridgehub() (common.Address, error) {
	return _IL1Bridge.Contract.Bridgehub(&_IL1Bridge.CallOpts)
}

// Bridgehub is a free data retrieval call binding the contract method 0x5d38962d.
//
// Solidity: function bridgehub() view returns(address)
func (_IL1Bridge *IL1BridgeCallerSession) Bridgehub() (common.Address, error) {
	return _IL1Bridge.Contract.Bridgehub(&_IL1Bridge.CallOpts)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 _chainId, bytes32 _l2TxHash) view returns(bytes32)
func (_IL1Bridge *IL1BridgeCaller) DepositHappened(opts *bind.CallOpts, _chainId *big.Int, _l2TxHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IL1Bridge.contract.Call(opts, &out, "depositHappened", _chainId, _l2TxHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 _chainId, bytes32 _l2TxHash) view returns(bytes32)
func (_IL1Bridge *IL1BridgeSession) DepositHappened(_chainId *big.Int, _l2TxHash [32]byte) ([32]byte, error) {
	return _IL1Bridge.Contract.DepositHappened(&_IL1Bridge.CallOpts, _chainId, _l2TxHash)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 _chainId, bytes32 _l2TxHash) view returns(bytes32)
func (_IL1Bridge *IL1BridgeCallerSession) DepositHappened(_chainId *big.Int, _l2TxHash [32]byte) ([32]byte, error) {
	return _IL1Bridge.Contract.DepositHappened(&_IL1Bridge.CallOpts, _chainId, _l2TxHash)
}

// IsWithdrawalFinalizedShared is a free data retrieval call binding the contract method 0x1f41f3f2.
//
// Solidity: function isWithdrawalFinalizedShared(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1Bridge *IL1BridgeCaller) IsWithdrawalFinalizedShared(opts *bind.CallOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1Bridge.contract.Call(opts, &out, "isWithdrawalFinalizedShared", _chainId, _l2BatchNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalizedShared is a free data retrieval call binding the contract method 0x1f41f3f2.
//
// Solidity: function isWithdrawalFinalizedShared(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1Bridge *IL1BridgeSession) IsWithdrawalFinalizedShared(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1Bridge.Contract.IsWithdrawalFinalizedShared(&_IL1Bridge.CallOpts, _chainId, _l2BatchNumber, _l2MessageIndex)
}

// IsWithdrawalFinalizedShared is a free data retrieval call binding the contract method 0x1f41f3f2.
//
// Solidity: function isWithdrawalFinalizedShared(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1Bridge *IL1BridgeCallerSession) IsWithdrawalFinalizedShared(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1Bridge.Contract.IsWithdrawalFinalizedShared(&_IL1Bridge.CallOpts, _chainId, _l2BatchNumber, _l2MessageIndex)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1Bridge *IL1BridgeCaller) L2BridgeAddress(opts *bind.CallOpts, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IL1Bridge.contract.Call(opts, &out, "l2BridgeAddress", _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1Bridge *IL1BridgeSession) L2BridgeAddress(_chainId *big.Int) (common.Address, error) {
	return _IL1Bridge.Contract.L2BridgeAddress(&_IL1Bridge.CallOpts, _chainId)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1Bridge *IL1BridgeCallerSession) L2BridgeAddress(_chainId *big.Int) (common.Address, error) {
	return _IL1Bridge.Contract.L2BridgeAddress(&_IL1Bridge.CallOpts, _chainId)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1Bridge *IL1BridgeTransactor) BridgehubConfirmL2Transaction(opts *bind.TransactOpts, _chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "bridgehubConfirmL2Transaction", _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1Bridge *IL1BridgeSession) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.BridgehubConfirmL2Transaction(&_IL1Bridge.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1Bridge *IL1BridgeTransactorSession) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.BridgehubConfirmL2Transaction(&_IL1Bridge.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0x7d7366b3.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1Bridge *IL1BridgeTransactor) BridgehubDeposit(opts *bind.TransactOpts, _chainId *big.Int, _prevMsgSender common.Address, _data []byte) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "bridgehubDeposit", _chainId, _prevMsgSender, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0x7d7366b3.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1Bridge *IL1BridgeSession) BridgehubDeposit(_chainId *big.Int, _prevMsgSender common.Address, _data []byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.BridgehubDeposit(&_IL1Bridge.TransactOpts, _chainId, _prevMsgSender, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0x7d7366b3.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1Bridge *IL1BridgeTransactorSession) BridgehubDeposit(_chainId *big.Int, _prevMsgSender common.Address, _data []byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.BridgehubDeposit(&_IL1Bridge.TransactOpts, _chainId, _prevMsgSender, _data)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_IL1Bridge *IL1BridgeTransactor) BridgehubDepositBaseToken(opts *bind.TransactOpts, _chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "bridgehubDepositBaseToken", _chainId, _prevMsgSender, _l1Token, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_IL1Bridge *IL1BridgeSession) BridgehubDepositBaseToken(_chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1Bridge.Contract.BridgehubDepositBaseToken(&_IL1Bridge.TransactOpts, _chainId, _prevMsgSender, _l1Token, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_IL1Bridge *IL1BridgeTransactorSession) BridgehubDepositBaseToken(_chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1Bridge.Contract.BridgehubDepositBaseToken(&_IL1Bridge.TransactOpts, _chainId, _prevMsgSender, _l1Token, _amount)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "claimFailedDeposit", _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.ClaimFailedDeposit(&_IL1Bridge.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactorSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.ClaimFailedDeposit(&_IL1Bridge.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0x55895e9c.
//
// Solidity: function deposit(uint256 _chainId, address _l2Receiver, address _l1Token, uint256 _mintValue, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1Bridge *IL1BridgeTransactor) Deposit(opts *bind.TransactOpts, _chainId *big.Int, _l2Receiver common.Address, _l1Token common.Address, _mintValue *big.Int, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "deposit", _chainId, _l2Receiver, _l1Token, _mintValue, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x55895e9c.
//
// Solidity: function deposit(uint256 _chainId, address _l2Receiver, address _l1Token, uint256 _mintValue, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1Bridge *IL1BridgeSession) Deposit(_chainId *big.Int, _l2Receiver common.Address, _l1Token common.Address, _mintValue *big.Int, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1Bridge.Contract.Deposit(&_IL1Bridge.TransactOpts, _chainId, _l2Receiver, _l1Token, _mintValue, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x55895e9c.
//
// Solidity: function deposit(uint256 _chainId, address _l2Receiver, address _l1Token, uint256 _mintValue, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1Bridge *IL1BridgeTransactorSession) Deposit(_chainId *big.Int, _l2Receiver common.Address, _l1Token common.Address, _mintValue *big.Int, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1Bridge.Contract.Deposit(&_IL1Bridge.TransactOpts, _chainId, _l2Receiver, _l1Token, _mintValue, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "finalizeWithdrawal", _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.FinalizeWithdrawal(&_IL1Bridge.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactorSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.FinalizeWithdrawal(&_IL1Bridge.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// IL1BridgeBridgehubDepositFinalizedIterator is returned from FilterBridgehubDepositFinalized and is used to iterate over the raw logs and unpacked data for BridgehubDepositFinalized events raised by the IL1Bridge contract.
type IL1BridgeBridgehubDepositFinalizedIterator struct {
	Event *IL1BridgeBridgehubDepositFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IL1BridgeBridgehubDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeBridgehubDepositFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IL1BridgeBridgehubDepositFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IL1BridgeBridgehubDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeBridgehubDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeBridgehubDepositFinalized represents a BridgehubDepositFinalized event raised by the IL1Bridge contract.
type IL1BridgeBridgehubDepositFinalized struct {
	ChainId         *big.Int
	TxDataHash      [32]byte
	L2DepositTxHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositFinalized is a free log retrieval operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1Bridge *IL1BridgeFilterer) FilterBridgehubDepositFinalized(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (*IL1BridgeBridgehubDepositFinalizedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeBridgehubDepositFinalizedIterator{contract: _IL1Bridge.contract, event: "BridgehubDepositFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositFinalized is a free log subscription operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1Bridge *IL1BridgeFilterer) WatchBridgehubDepositFinalized(opts *bind.WatchOpts, sink chan<- *IL1BridgeBridgehubDepositFinalized, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeBridgehubDepositFinalized)
				if err := _IL1Bridge.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgehubDepositFinalized is a log parse operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1Bridge *IL1BridgeFilterer) ParseBridgehubDepositFinalized(log types.Log) (*IL1BridgeBridgehubDepositFinalized, error) {
	event := new(IL1BridgeBridgehubDepositFinalized)
	if err := _IL1Bridge.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1BridgeBridgehubDepositInitiatedSharedBridgeIterator is returned from FilterBridgehubDepositInitiatedSharedBridge and is used to iterate over the raw logs and unpacked data for BridgehubDepositInitiatedSharedBridge events raised by the IL1Bridge contract.
type IL1BridgeBridgehubDepositInitiatedSharedBridgeIterator struct {
	Event *IL1BridgeBridgehubDepositInitiatedSharedBridge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IL1BridgeBridgehubDepositInitiatedSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeBridgehubDepositInitiatedSharedBridge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IL1BridgeBridgehubDepositInitiatedSharedBridge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IL1BridgeBridgehubDepositInitiatedSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeBridgehubDepositInitiatedSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeBridgehubDepositInitiatedSharedBridge represents a BridgehubDepositInitiatedSharedBridge event raised by the IL1Bridge contract.
type IL1BridgeBridgehubDepositInitiatedSharedBridge struct {
	ChainId    *big.Int
	TxDataHash [32]byte
	From       common.Address
	To         common.Address
	L1Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositInitiatedSharedBridge is a free log retrieval operation binding the contract event 0xb58916f15ee1c87cb4ddf5a07f5dddb605cef92c8d7c0b9041057df9fdc02e60.
//
// Solidity: event BridgehubDepositInitiatedSharedBridge(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) FilterBridgehubDepositInitiatedSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (*IL1BridgeBridgehubDepositInitiatedSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "BridgehubDepositInitiatedSharedBridge", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeBridgehubDepositInitiatedSharedBridgeIterator{contract: _IL1Bridge.contract, event: "BridgehubDepositInitiatedSharedBridge", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositInitiatedSharedBridge is a free log subscription operation binding the contract event 0xb58916f15ee1c87cb4ddf5a07f5dddb605cef92c8d7c0b9041057df9fdc02e60.
//
// Solidity: event BridgehubDepositInitiatedSharedBridge(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) WatchBridgehubDepositInitiatedSharedBridge(opts *bind.WatchOpts, sink chan<- *IL1BridgeBridgehubDepositInitiatedSharedBridge, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "BridgehubDepositInitiatedSharedBridge", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeBridgehubDepositInitiatedSharedBridge)
				if err := _IL1Bridge.contract.UnpackLog(event, "BridgehubDepositInitiatedSharedBridge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgehubDepositInitiatedSharedBridge is a log parse operation binding the contract event 0xb58916f15ee1c87cb4ddf5a07f5dddb605cef92c8d7c0b9041057df9fdc02e60.
//
// Solidity: event BridgehubDepositInitiatedSharedBridge(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) ParseBridgehubDepositInitiatedSharedBridge(log types.Log) (*IL1BridgeBridgehubDepositInitiatedSharedBridge, error) {
	event := new(IL1BridgeBridgehubDepositInitiatedSharedBridge)
	if err := _IL1Bridge.contract.UnpackLog(event, "BridgehubDepositInitiatedSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1BridgeClaimedFailedDepositSharedBridgeIterator is returned from FilterClaimedFailedDepositSharedBridge and is used to iterate over the raw logs and unpacked data for ClaimedFailedDepositSharedBridge events raised by the IL1Bridge contract.
type IL1BridgeClaimedFailedDepositSharedBridgeIterator struct {
	Event *IL1BridgeClaimedFailedDepositSharedBridge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IL1BridgeClaimedFailedDepositSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeClaimedFailedDepositSharedBridge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IL1BridgeClaimedFailedDepositSharedBridge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IL1BridgeClaimedFailedDepositSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeClaimedFailedDepositSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeClaimedFailedDepositSharedBridge represents a ClaimedFailedDepositSharedBridge event raised by the IL1Bridge contract.
type IL1BridgeClaimedFailedDepositSharedBridge struct {
	ChainId *big.Int
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDepositSharedBridge is a free log retrieval operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) FilterClaimedFailedDepositSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, to []common.Address, l1Token []common.Address) (*IL1BridgeClaimedFailedDepositSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "ClaimedFailedDepositSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeClaimedFailedDepositSharedBridgeIterator{contract: _IL1Bridge.contract, event: "ClaimedFailedDepositSharedBridge", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDepositSharedBridge is a free log subscription operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) WatchClaimedFailedDepositSharedBridge(opts *bind.WatchOpts, sink chan<- *IL1BridgeClaimedFailedDepositSharedBridge, chainId []*big.Int, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "ClaimedFailedDepositSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeClaimedFailedDepositSharedBridge)
				if err := _IL1Bridge.contract.UnpackLog(event, "ClaimedFailedDepositSharedBridge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedFailedDepositSharedBridge is a log parse operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) ParseClaimedFailedDepositSharedBridge(log types.Log) (*IL1BridgeClaimedFailedDepositSharedBridge, error) {
	event := new(IL1BridgeClaimedFailedDepositSharedBridge)
	if err := _IL1Bridge.contract.UnpackLog(event, "ClaimedFailedDepositSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1BridgeDepositInitiatedSharedBridgeIterator is returned from FilterDepositInitiatedSharedBridge and is used to iterate over the raw logs and unpacked data for DepositInitiatedSharedBridge events raised by the IL1Bridge contract.
type IL1BridgeDepositInitiatedSharedBridgeIterator struct {
	Event *IL1BridgeDepositInitiatedSharedBridge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IL1BridgeDepositInitiatedSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeDepositInitiatedSharedBridge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IL1BridgeDepositInitiatedSharedBridge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IL1BridgeDepositInitiatedSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeDepositInitiatedSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeDepositInitiatedSharedBridge represents a DepositInitiatedSharedBridge event raised by the IL1Bridge contract.
type IL1BridgeDepositInitiatedSharedBridge struct {
	ChainId         *big.Int
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Token         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositInitiatedSharedBridge is a free log retrieval operation binding the contract event 0x461e9f9a4610edbd60c7fc3c7b509fe60c53343c5a9f310b8cf9e18af00bef31.
//
// Solidity: event DepositInitiatedSharedBridge(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) FilterDepositInitiatedSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (*IL1BridgeDepositInitiatedSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "DepositInitiatedSharedBridge", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeDepositInitiatedSharedBridgeIterator{contract: _IL1Bridge.contract, event: "DepositInitiatedSharedBridge", logs: logs, sub: sub}, nil
}

// WatchDepositInitiatedSharedBridge is a free log subscription operation binding the contract event 0x461e9f9a4610edbd60c7fc3c7b509fe60c53343c5a9f310b8cf9e18af00bef31.
//
// Solidity: event DepositInitiatedSharedBridge(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) WatchDepositInitiatedSharedBridge(opts *bind.WatchOpts, sink chan<- *IL1BridgeDepositInitiatedSharedBridge, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "DepositInitiatedSharedBridge", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeDepositInitiatedSharedBridge)
				if err := _IL1Bridge.contract.UnpackLog(event, "DepositInitiatedSharedBridge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositInitiatedSharedBridge is a log parse operation binding the contract event 0x461e9f9a4610edbd60c7fc3c7b509fe60c53343c5a9f310b8cf9e18af00bef31.
//
// Solidity: event DepositInitiatedSharedBridge(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) ParseDepositInitiatedSharedBridge(log types.Log) (*IL1BridgeDepositInitiatedSharedBridge, error) {
	event := new(IL1BridgeDepositInitiatedSharedBridge)
	if err := _IL1Bridge.contract.UnpackLog(event, "DepositInitiatedSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1BridgeWithdrawalFinalizedSharedBridgeIterator is returned from FilterWithdrawalFinalizedSharedBridge and is used to iterate over the raw logs and unpacked data for WithdrawalFinalizedSharedBridge events raised by the IL1Bridge contract.
type IL1BridgeWithdrawalFinalizedSharedBridgeIterator struct {
	Event *IL1BridgeWithdrawalFinalizedSharedBridge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IL1BridgeWithdrawalFinalizedSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeWithdrawalFinalizedSharedBridge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IL1BridgeWithdrawalFinalizedSharedBridge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IL1BridgeWithdrawalFinalizedSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeWithdrawalFinalizedSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeWithdrawalFinalizedSharedBridge represents a WithdrawalFinalizedSharedBridge event raised by the IL1Bridge contract.
type IL1BridgeWithdrawalFinalizedSharedBridge struct {
	ChainId *big.Int
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalizedSharedBridge is a free log retrieval operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) FilterWithdrawalFinalizedSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, to []common.Address, l1Token []common.Address) (*IL1BridgeWithdrawalFinalizedSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "WithdrawalFinalizedSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeWithdrawalFinalizedSharedBridgeIterator{contract: _IL1Bridge.contract, event: "WithdrawalFinalizedSharedBridge", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalizedSharedBridge is a free log subscription operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) WatchWithdrawalFinalizedSharedBridge(opts *bind.WatchOpts, sink chan<- *IL1BridgeWithdrawalFinalizedSharedBridge, chainId []*big.Int, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "WithdrawalFinalizedSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeWithdrawalFinalizedSharedBridge)
				if err := _IL1Bridge.contract.UnpackLog(event, "WithdrawalFinalizedSharedBridge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalFinalizedSharedBridge is a log parse operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) ParseWithdrawalFinalizedSharedBridge(log types.Log) (*IL1BridgeWithdrawalFinalizedSharedBridge, error) {
	event := new(IL1BridgeWithdrawalFinalizedSharedBridge)
	if err := _IL1Bridge.contract.UnpackLog(event, "WithdrawalFinalizedSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
