// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itimestampasserter

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

// ITimestampAsserterMetaData contains all meta data concerning the ITimestampAsserter contract.
var ITimestampAsserterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"assertTimestampInRange\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ITimestampAsserterABI is the input ABI used to generate the binding from.
// Deprecated: Use ITimestampAsserterMetaData.ABI instead.
var ITimestampAsserterABI = ITimestampAsserterMetaData.ABI

// ITimestampAsserter is an auto generated Go binding around an Ethereum contract.
type ITimestampAsserter struct {
	ITimestampAsserterCaller     // Read-only binding to the contract
	ITimestampAsserterTransactor // Write-only binding to the contract
	ITimestampAsserterFilterer   // Log filterer for contract events
}

// ITimestampAsserterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITimestampAsserterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimestampAsserterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITimestampAsserterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimestampAsserterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITimestampAsserterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimestampAsserterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITimestampAsserterSession struct {
	Contract     *ITimestampAsserter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ITimestampAsserterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITimestampAsserterCallerSession struct {
	Contract *ITimestampAsserterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ITimestampAsserterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITimestampAsserterTransactorSession struct {
	Contract     *ITimestampAsserterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ITimestampAsserterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITimestampAsserterRaw struct {
	Contract *ITimestampAsserter // Generic contract binding to access the raw methods on
}

// ITimestampAsserterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITimestampAsserterCallerRaw struct {
	Contract *ITimestampAsserterCaller // Generic read-only contract binding to access the raw methods on
}

// ITimestampAsserterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITimestampAsserterTransactorRaw struct {
	Contract *ITimestampAsserterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITimestampAsserter creates a new instance of ITimestampAsserter, bound to a specific deployed contract.
func NewITimestampAsserter(address common.Address, backend bind.ContractBackend) (*ITimestampAsserter, error) {
	contract, err := bindITimestampAsserter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITimestampAsserter{ITimestampAsserterCaller: ITimestampAsserterCaller{contract: contract}, ITimestampAsserterTransactor: ITimestampAsserterTransactor{contract: contract}, ITimestampAsserterFilterer: ITimestampAsserterFilterer{contract: contract}}, nil
}

// NewITimestampAsserterCaller creates a new read-only instance of ITimestampAsserter, bound to a specific deployed contract.
func NewITimestampAsserterCaller(address common.Address, caller bind.ContractCaller) (*ITimestampAsserterCaller, error) {
	contract, err := bindITimestampAsserter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITimestampAsserterCaller{contract: contract}, nil
}

// NewITimestampAsserterTransactor creates a new write-only instance of ITimestampAsserter, bound to a specific deployed contract.
func NewITimestampAsserterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITimestampAsserterTransactor, error) {
	contract, err := bindITimestampAsserter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITimestampAsserterTransactor{contract: contract}, nil
}

// NewITimestampAsserterFilterer creates a new log filterer instance of ITimestampAsserter, bound to a specific deployed contract.
func NewITimestampAsserterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITimestampAsserterFilterer, error) {
	contract, err := bindITimestampAsserter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITimestampAsserterFilterer{contract: contract}, nil
}

// bindITimestampAsserter binds a generic wrapper to an already deployed contract.
func bindITimestampAsserter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITimestampAsserterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimestampAsserter *ITimestampAsserterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimestampAsserter.Contract.ITimestampAsserterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimestampAsserter *ITimestampAsserterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimestampAsserter.Contract.ITimestampAsserterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimestampAsserter *ITimestampAsserterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimestampAsserter.Contract.ITimestampAsserterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimestampAsserter *ITimestampAsserterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimestampAsserter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimestampAsserter *ITimestampAsserterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimestampAsserter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimestampAsserter *ITimestampAsserterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimestampAsserter.Contract.contract.Transact(opts, method, params...)
}

// AssertTimestampInRange is a free data retrieval call binding the contract method 0x5b1a0c91.
//
// Solidity: function assertTimestampInRange(uint256 start, uint256 end) view returns()
func (_ITimestampAsserter *ITimestampAsserterCaller) AssertTimestampInRange(opts *bind.CallOpts, start *big.Int, end *big.Int) error {
	var out []interface{}
	err := _ITimestampAsserter.contract.Call(opts, &out, "assertTimestampInRange", start, end)

	if err != nil {
		return err
	}

	return err

}

// AssertTimestampInRange is a free data retrieval call binding the contract method 0x5b1a0c91.
//
// Solidity: function assertTimestampInRange(uint256 start, uint256 end) view returns()
func (_ITimestampAsserter *ITimestampAsserterSession) AssertTimestampInRange(start *big.Int, end *big.Int) error {
	return _ITimestampAsserter.Contract.AssertTimestampInRange(&_ITimestampAsserter.CallOpts, start, end)
}

// AssertTimestampInRange is a free data retrieval call binding the contract method 0x5b1a0c91.
//
// Solidity: function assertTimestampInRange(uint256 start, uint256 end) view returns()
func (_ITimestampAsserter *ITimestampAsserterCallerSession) AssertTimestampInRange(start *big.Int, end *big.Int) error {
	return _ITimestampAsserter.Contract.AssertTimestampInRange(&_ITimestampAsserter.CallOpts, start, end)
}
