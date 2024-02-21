// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

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

// IncrementerMetaData contains all meta data concerning the Incrementer contract.
var IncrementerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_incrementer\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IncrementerABI is the input ABI used to generate the binding from.
// Deprecated: Use IncrementerMetaData.ABI instead.
var IncrementerABI = IncrementerMetaData.ABI

// Incrementer is an auto generated Go binding around an Ethereum contract.
type Incrementer struct {
	IncrementerCaller     // Read-only binding to the contract
	IncrementerTransactor // Write-only binding to the contract
	IncrementerFilterer   // Log filterer for contract events
}

// IncrementerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncrementerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncrementerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncrementerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncrementerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncrementerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncrementerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncrementerSession struct {
	Contract     *Incrementer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncrementerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncrementerCallerSession struct {
	Contract *IncrementerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IncrementerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncrementerTransactorSession struct {
	Contract     *IncrementerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IncrementerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncrementerRaw struct {
	Contract *Incrementer // Generic contract binding to access the raw methods on
}

// IncrementerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncrementerCallerRaw struct {
	Contract *IncrementerCaller // Generic read-only contract binding to access the raw methods on
}

// IncrementerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncrementerTransactorRaw struct {
	Contract *IncrementerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncrementer creates a new instance of Incrementer, bound to a specific deployed contract.
func NewIncrementer(address common.Address, backend bind.ContractBackend) (*Incrementer, error) {
	contract, err := bindIncrementer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Incrementer{IncrementerCaller: IncrementerCaller{contract: contract}, IncrementerTransactor: IncrementerTransactor{contract: contract}, IncrementerFilterer: IncrementerFilterer{contract: contract}}, nil
}

// NewIncrementerCaller creates a new read-only instance of Incrementer, bound to a specific deployed contract.
func NewIncrementerCaller(address common.Address, caller bind.ContractCaller) (*IncrementerCaller, error) {
	contract, err := bindIncrementer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncrementerCaller{contract: contract}, nil
}

// NewIncrementerTransactor creates a new write-only instance of Incrementer, bound to a specific deployed contract.
func NewIncrementerTransactor(address common.Address, transactor bind.ContractTransactor) (*IncrementerTransactor, error) {
	contract, err := bindIncrementer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncrementerTransactor{contract: contract}, nil
}

// NewIncrementerFilterer creates a new log filterer instance of Incrementer, bound to a specific deployed contract.
func NewIncrementerFilterer(address common.Address, filterer bind.ContractFilterer) (*IncrementerFilterer, error) {
	contract, err := bindIncrementer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncrementerFilterer{contract: contract}, nil
}

// bindIncrementer binds a generic wrapper to an already deployed contract.
func bindIncrementer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IncrementerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incrementer *IncrementerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Incrementer.Contract.IncrementerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incrementer *IncrementerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incrementer.Contract.IncrementerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incrementer *IncrementerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incrementer.Contract.IncrementerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incrementer *IncrementerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Incrementer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incrementer *IncrementerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incrementer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incrementer *IncrementerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incrementer.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Incrementer *IncrementerCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Incrementer.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Incrementer *IncrementerSession) Get() (*big.Int, error) {
	return _Incrementer.Contract.Get(&_Incrementer.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Incrementer *IncrementerCallerSession) Get() (*big.Int, error) {
	return _Incrementer.Contract.Get(&_Incrementer.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Incrementer *IncrementerTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incrementer.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Incrementer *IncrementerSession) Increment() (*types.Transaction, error) {
	return _Incrementer.Contract.Increment(&_Incrementer.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Incrementer *IncrementerTransactorSession) Increment() (*types.Transaction, error) {
	return _Incrementer.Contract.Increment(&_Incrementer.TransactOpts)
}