// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package timestampasserterlocator

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

// TimestampAsserterLocatorMetaData contains all meta data concerning the TimestampAsserterLocator contract.
var TimestampAsserterLocatorMetaData = &bind.MetaData{
	ABI: "[]",
}

// TimestampAsserterLocatorABI is the input ABI used to generate the binding from.
// Deprecated: Use TimestampAsserterLocatorMetaData.ABI instead.
var TimestampAsserterLocatorABI = TimestampAsserterLocatorMetaData.ABI

// TimestampAsserterLocator is an auto generated Go binding around an Ethereum contract.
type TimestampAsserterLocator struct {
	TimestampAsserterLocatorCaller     // Read-only binding to the contract
	TimestampAsserterLocatorTransactor // Write-only binding to the contract
	TimestampAsserterLocatorFilterer   // Log filterer for contract events
}

// TimestampAsserterLocatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimestampAsserterLocatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimestampAsserterLocatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimestampAsserterLocatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimestampAsserterLocatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimestampAsserterLocatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimestampAsserterLocatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimestampAsserterLocatorSession struct {
	Contract     *TimestampAsserterLocator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TimestampAsserterLocatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimestampAsserterLocatorCallerSession struct {
	Contract *TimestampAsserterLocatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// TimestampAsserterLocatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimestampAsserterLocatorTransactorSession struct {
	Contract     *TimestampAsserterLocatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TimestampAsserterLocatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimestampAsserterLocatorRaw struct {
	Contract *TimestampAsserterLocator // Generic contract binding to access the raw methods on
}

// TimestampAsserterLocatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimestampAsserterLocatorCallerRaw struct {
	Contract *TimestampAsserterLocatorCaller // Generic read-only contract binding to access the raw methods on
}

// TimestampAsserterLocatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimestampAsserterLocatorTransactorRaw struct {
	Contract *TimestampAsserterLocatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimestampAsserterLocator creates a new instance of TimestampAsserterLocator, bound to a specific deployed contract.
func NewTimestampAsserterLocator(address common.Address, backend bind.ContractBackend) (*TimestampAsserterLocator, error) {
	contract, err := bindTimestampAsserterLocator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimestampAsserterLocator{TimestampAsserterLocatorCaller: TimestampAsserterLocatorCaller{contract: contract}, TimestampAsserterLocatorTransactor: TimestampAsserterLocatorTransactor{contract: contract}, TimestampAsserterLocatorFilterer: TimestampAsserterLocatorFilterer{contract: contract}}, nil
}

// NewTimestampAsserterLocatorCaller creates a new read-only instance of TimestampAsserterLocator, bound to a specific deployed contract.
func NewTimestampAsserterLocatorCaller(address common.Address, caller bind.ContractCaller) (*TimestampAsserterLocatorCaller, error) {
	contract, err := bindTimestampAsserterLocator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimestampAsserterLocatorCaller{contract: contract}, nil
}

// NewTimestampAsserterLocatorTransactor creates a new write-only instance of TimestampAsserterLocator, bound to a specific deployed contract.
func NewTimestampAsserterLocatorTransactor(address common.Address, transactor bind.ContractTransactor) (*TimestampAsserterLocatorTransactor, error) {
	contract, err := bindTimestampAsserterLocator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimestampAsserterLocatorTransactor{contract: contract}, nil
}

// NewTimestampAsserterLocatorFilterer creates a new log filterer instance of TimestampAsserterLocator, bound to a specific deployed contract.
func NewTimestampAsserterLocatorFilterer(address common.Address, filterer bind.ContractFilterer) (*TimestampAsserterLocatorFilterer, error) {
	contract, err := bindTimestampAsserterLocator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimestampAsserterLocatorFilterer{contract: contract}, nil
}

// bindTimestampAsserterLocator binds a generic wrapper to an already deployed contract.
func bindTimestampAsserterLocator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimestampAsserterLocatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimestampAsserterLocator *TimestampAsserterLocatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimestampAsserterLocator.Contract.TimestampAsserterLocatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimestampAsserterLocator *TimestampAsserterLocatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimestampAsserterLocator.Contract.TimestampAsserterLocatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimestampAsserterLocator *TimestampAsserterLocatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimestampAsserterLocator.Contract.TimestampAsserterLocatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimestampAsserterLocator *TimestampAsserterLocatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimestampAsserterLocator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimestampAsserterLocator *TimestampAsserterLocatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimestampAsserterLocator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimestampAsserterLocator *TimestampAsserterLocatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimestampAsserterLocator.Contract.contract.Transact(opts, method, params...)
}
