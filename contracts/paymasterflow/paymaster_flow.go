// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymasterflow

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

// IPaymasterFlowMetaData contains all meta data concerning the IPaymasterFlow contract.
var IPaymasterFlowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minAllowance\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_innerInput\",\"type\":\"bytes\"}],\"name\":\"approvalBased\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"general\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPaymasterFlowABI is the input ABI used to generate the binding from.
// Deprecated: Use IPaymasterFlowMetaData.ABI instead.
var IPaymasterFlowABI = IPaymasterFlowMetaData.ABI

// IPaymasterFlow is an auto generated Go binding around an Ethereum contract.
type IPaymasterFlow struct {
	IPaymasterFlowCaller     // Read-only binding to the contract
	IPaymasterFlowTransactor // Write-only binding to the contract
	IPaymasterFlowFilterer   // Log filterer for contract events
}

// IPaymasterFlowCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPaymasterFlowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterFlowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPaymasterFlowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterFlowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPaymasterFlowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymasterFlowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPaymasterFlowSession struct {
	Contract     *IPaymasterFlow   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPaymasterFlowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPaymasterFlowCallerSession struct {
	Contract *IPaymasterFlowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IPaymasterFlowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPaymasterFlowTransactorSession struct {
	Contract     *IPaymasterFlowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IPaymasterFlowRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPaymasterFlowRaw struct {
	Contract *IPaymasterFlow // Generic contract binding to access the raw methods on
}

// IPaymasterFlowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPaymasterFlowCallerRaw struct {
	Contract *IPaymasterFlowCaller // Generic read-only contract binding to access the raw methods on
}

// IPaymasterFlowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPaymasterFlowTransactorRaw struct {
	Contract *IPaymasterFlowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPaymasterFlow creates a new instance of IPaymasterFlow, bound to a specific deployed contract.
func NewIPaymasterFlow(address common.Address, backend bind.ContractBackend) (*IPaymasterFlow, error) {
	contract, err := bindIPaymasterFlow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPaymasterFlow{IPaymasterFlowCaller: IPaymasterFlowCaller{contract: contract}, IPaymasterFlowTransactor: IPaymasterFlowTransactor{contract: contract}, IPaymasterFlowFilterer: IPaymasterFlowFilterer{contract: contract}}, nil
}

// NewIPaymasterFlowCaller creates a new read-only instance of IPaymasterFlow, bound to a specific deployed contract.
func NewIPaymasterFlowCaller(address common.Address, caller bind.ContractCaller) (*IPaymasterFlowCaller, error) {
	contract, err := bindIPaymasterFlow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymasterFlowCaller{contract: contract}, nil
}

// NewIPaymasterFlowTransactor creates a new write-only instance of IPaymasterFlow, bound to a specific deployed contract.
func NewIPaymasterFlowTransactor(address common.Address, transactor bind.ContractTransactor) (*IPaymasterFlowTransactor, error) {
	contract, err := bindIPaymasterFlow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymasterFlowTransactor{contract: contract}, nil
}

// NewIPaymasterFlowFilterer creates a new log filterer instance of IPaymasterFlow, bound to a specific deployed contract.
func NewIPaymasterFlowFilterer(address common.Address, filterer bind.ContractFilterer) (*IPaymasterFlowFilterer, error) {
	contract, err := bindIPaymasterFlow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPaymasterFlowFilterer{contract: contract}, nil
}

// bindIPaymasterFlow binds a generic wrapper to an already deployed contract.
func bindIPaymasterFlow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPaymasterFlowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymasterFlow *IPaymasterFlowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymasterFlow.Contract.IPaymasterFlowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymasterFlow *IPaymasterFlowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.IPaymasterFlowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymasterFlow *IPaymasterFlowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.IPaymasterFlowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPaymasterFlow *IPaymasterFlowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPaymasterFlow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPaymasterFlow *IPaymasterFlowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPaymasterFlow *IPaymasterFlowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.contract.Transact(opts, method, params...)
}

// ApprovalBased is a paid mutator transaction binding the contract method 0x949431dc.
//
// Solidity: function approvalBased(address _token, uint256 _minAllowance, bytes _innerInput) returns()
func (_IPaymasterFlow *IPaymasterFlowTransactor) ApprovalBased(opts *bind.TransactOpts, _token common.Address, _minAllowance *big.Int, _innerInput []byte) (*types.Transaction, error) {
	return _IPaymasterFlow.contract.Transact(opts, "approvalBased", _token, _minAllowance, _innerInput)
}

// ApprovalBased is a paid mutator transaction binding the contract method 0x949431dc.
//
// Solidity: function approvalBased(address _token, uint256 _minAllowance, bytes _innerInput) returns()
func (_IPaymasterFlow *IPaymasterFlowSession) ApprovalBased(_token common.Address, _minAllowance *big.Int, _innerInput []byte) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.ApprovalBased(&_IPaymasterFlow.TransactOpts, _token, _minAllowance, _innerInput)
}

// ApprovalBased is a paid mutator transaction binding the contract method 0x949431dc.
//
// Solidity: function approvalBased(address _token, uint256 _minAllowance, bytes _innerInput) returns()
func (_IPaymasterFlow *IPaymasterFlowTransactorSession) ApprovalBased(_token common.Address, _minAllowance *big.Int, _innerInput []byte) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.ApprovalBased(&_IPaymasterFlow.TransactOpts, _token, _minAllowance, _innerInput)
}

// General is a paid mutator transaction binding the contract method 0x8c5a3445.
//
// Solidity: function general(bytes input) returns()
func (_IPaymasterFlow *IPaymasterFlowTransactor) General(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _IPaymasterFlow.contract.Transact(opts, "general", input)
}

// General is a paid mutator transaction binding the contract method 0x8c5a3445.
//
// Solidity: function general(bytes input) returns()
func (_IPaymasterFlow *IPaymasterFlowSession) General(input []byte) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.General(&_IPaymasterFlow.TransactOpts, input)
}

// General is a paid mutator transaction binding the contract method 0x8c5a3445.
//
// Solidity: function general(bytes input) returns()
func (_IPaymasterFlow *IPaymasterFlowTransactorSession) General(input []byte) (*types.Transaction, error) {
	return _IPaymasterFlow.Contract.General(&_IPaymasterFlow.TransactOpts, input)
}
