// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package efficientproxy

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

// EfficientProxyMetaData contains all meta data concerning the EfficientProxy contract.
var EfficientProxyMetaData = &bind.MetaData{
	ABI: "[]",
}

// EfficientProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use EfficientProxyMetaData.ABI instead.
var EfficientProxyABI = EfficientProxyMetaData.ABI

// EfficientProxy is an auto generated Go binding around an Ethereum contract.
type EfficientProxy struct {
	EfficientProxyCaller     // Read-only binding to the contract
	EfficientProxyTransactor // Write-only binding to the contract
	EfficientProxyFilterer   // Log filterer for contract events
}

// EfficientProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type EfficientProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EfficientProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EfficientProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EfficientProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EfficientProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EfficientProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EfficientProxySession struct {
	Contract     *EfficientProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EfficientProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EfficientProxyCallerSession struct {
	Contract *EfficientProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EfficientProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EfficientProxyTransactorSession struct {
	Contract     *EfficientProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EfficientProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type EfficientProxyRaw struct {
	Contract *EfficientProxy // Generic contract binding to access the raw methods on
}

// EfficientProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EfficientProxyCallerRaw struct {
	Contract *EfficientProxyCaller // Generic read-only contract binding to access the raw methods on
}

// EfficientProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EfficientProxyTransactorRaw struct {
	Contract *EfficientProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEfficientProxy creates a new instance of EfficientProxy, bound to a specific deployed contract.
func NewEfficientProxy(address common.Address, backend bind.ContractBackend) (*EfficientProxy, error) {
	contract, err := bindEfficientProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EfficientProxy{EfficientProxyCaller: EfficientProxyCaller{contract: contract}, EfficientProxyTransactor: EfficientProxyTransactor{contract: contract}, EfficientProxyFilterer: EfficientProxyFilterer{contract: contract}}, nil
}

// NewEfficientProxyCaller creates a new read-only instance of EfficientProxy, bound to a specific deployed contract.
func NewEfficientProxyCaller(address common.Address, caller bind.ContractCaller) (*EfficientProxyCaller, error) {
	contract, err := bindEfficientProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EfficientProxyCaller{contract: contract}, nil
}

// NewEfficientProxyTransactor creates a new write-only instance of EfficientProxy, bound to a specific deployed contract.
func NewEfficientProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*EfficientProxyTransactor, error) {
	contract, err := bindEfficientProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EfficientProxyTransactor{contract: contract}, nil
}

// NewEfficientProxyFilterer creates a new log filterer instance of EfficientProxy, bound to a specific deployed contract.
func NewEfficientProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*EfficientProxyFilterer, error) {
	contract, err := bindEfficientProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EfficientProxyFilterer{contract: contract}, nil
}

// bindEfficientProxy binds a generic wrapper to an already deployed contract.
func bindEfficientProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EfficientProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EfficientProxy *EfficientProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EfficientProxy.Contract.EfficientProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EfficientProxy *EfficientProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EfficientProxy.Contract.EfficientProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EfficientProxy *EfficientProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EfficientProxy.Contract.EfficientProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EfficientProxy *EfficientProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EfficientProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EfficientProxy *EfficientProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EfficientProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EfficientProxy *EfficientProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EfficientProxy.Contract.contract.Transact(opts, method, params...)
}
