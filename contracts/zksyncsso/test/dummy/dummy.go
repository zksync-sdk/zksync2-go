// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dummy

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

// DummyMetaData contains all meta data concerning the Dummy contract.
var DummyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"dummy\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DummyABI is the input ABI used to generate the binding from.
// Deprecated: Use DummyMetaData.ABI instead.
var DummyABI = DummyMetaData.ABI

// Dummy is an auto generated Go binding around an Ethereum contract.
type Dummy struct {
	DummyCaller     // Read-only binding to the contract
	DummyTransactor // Write-only binding to the contract
	DummyFilterer   // Log filterer for contract events
}

// DummyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummySession struct {
	Contract     *Dummy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyCallerSession struct {
	Contract *DummyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DummyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyTransactorSession struct {
	Contract     *DummyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyRaw struct {
	Contract *Dummy // Generic contract binding to access the raw methods on
}

// DummyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyCallerRaw struct {
	Contract *DummyCaller // Generic read-only contract binding to access the raw methods on
}

// DummyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyTransactorRaw struct {
	Contract *DummyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummy creates a new instance of Dummy, bound to a specific deployed contract.
func NewDummy(address common.Address, backend bind.ContractBackend) (*Dummy, error) {
	contract, err := bindDummy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dummy{DummyCaller: DummyCaller{contract: contract}, DummyTransactor: DummyTransactor{contract: contract}, DummyFilterer: DummyFilterer{contract: contract}}, nil
}

// NewDummyCaller creates a new read-only instance of Dummy, bound to a specific deployed contract.
func NewDummyCaller(address common.Address, caller bind.ContractCaller) (*DummyCaller, error) {
	contract, err := bindDummy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyCaller{contract: contract}, nil
}

// NewDummyTransactor creates a new write-only instance of Dummy, bound to a specific deployed contract.
func NewDummyTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyTransactor, error) {
	contract, err := bindDummy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyTransactor{contract: contract}, nil
}

// NewDummyFilterer creates a new log filterer instance of Dummy, bound to a specific deployed contract.
func NewDummyFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyFilterer, error) {
	contract, err := bindDummy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyFilterer{contract: contract}, nil
}

// bindDummy binds a generic wrapper to an already deployed contract.
func bindDummy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DummyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dummy *DummyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dummy.Contract.DummyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dummy *DummyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dummy.Contract.DummyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dummy *DummyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dummy.Contract.DummyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dummy *DummyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dummy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dummy *DummyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dummy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dummy *DummyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dummy.Contract.contract.Transact(opts, method, params...)
}

// Dummy is a free data retrieval call binding the contract method 0x32e43a11.
//
// Solidity: function dummy() view returns(string)
func (_Dummy *DummyCaller) Dummy(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dummy.contract.Call(opts, &out, "dummy")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Dummy is a free data retrieval call binding the contract method 0x32e43a11.
//
// Solidity: function dummy() view returns(string)
func (_Dummy *DummySession) Dummy() (string, error) {
	return _Dummy.Contract.Dummy(&_Dummy.CallOpts)
}

// Dummy is a free data retrieval call binding the contract method 0x32e43a11.
//
// Solidity: function dummy() view returns(string)
func (_Dummy *DummyCallerSession) Dummy() (string, error) {
	return _Dummy.Contract.Dummy(&_Dummy.CallOpts)
}
