// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bootloaderauth

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

// BootloaderAuthMetaData contains all meta data concerning the BootloaderAuth contract.
var BootloaderAuthMetaData = &bind.MetaData{
	ABI: "[]",
}

// BootloaderAuthABI is the input ABI used to generate the binding from.
// Deprecated: Use BootloaderAuthMetaData.ABI instead.
var BootloaderAuthABI = BootloaderAuthMetaData.ABI

// BootloaderAuth is an auto generated Go binding around an Ethereum contract.
type BootloaderAuth struct {
	BootloaderAuthCaller     // Read-only binding to the contract
	BootloaderAuthTransactor // Write-only binding to the contract
	BootloaderAuthFilterer   // Log filterer for contract events
}

// BootloaderAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type BootloaderAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootloaderAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BootloaderAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootloaderAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BootloaderAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootloaderAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BootloaderAuthSession struct {
	Contract     *BootloaderAuth   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BootloaderAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BootloaderAuthCallerSession struct {
	Contract *BootloaderAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BootloaderAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BootloaderAuthTransactorSession struct {
	Contract     *BootloaderAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BootloaderAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type BootloaderAuthRaw struct {
	Contract *BootloaderAuth // Generic contract binding to access the raw methods on
}

// BootloaderAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BootloaderAuthCallerRaw struct {
	Contract *BootloaderAuthCaller // Generic read-only contract binding to access the raw methods on
}

// BootloaderAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BootloaderAuthTransactorRaw struct {
	Contract *BootloaderAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBootloaderAuth creates a new instance of BootloaderAuth, bound to a specific deployed contract.
func NewBootloaderAuth(address common.Address, backend bind.ContractBackend) (*BootloaderAuth, error) {
	contract, err := bindBootloaderAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BootloaderAuth{BootloaderAuthCaller: BootloaderAuthCaller{contract: contract}, BootloaderAuthTransactor: BootloaderAuthTransactor{contract: contract}, BootloaderAuthFilterer: BootloaderAuthFilterer{contract: contract}}, nil
}

// NewBootloaderAuthCaller creates a new read-only instance of BootloaderAuth, bound to a specific deployed contract.
func NewBootloaderAuthCaller(address common.Address, caller bind.ContractCaller) (*BootloaderAuthCaller, error) {
	contract, err := bindBootloaderAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BootloaderAuthCaller{contract: contract}, nil
}

// NewBootloaderAuthTransactor creates a new write-only instance of BootloaderAuth, bound to a specific deployed contract.
func NewBootloaderAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*BootloaderAuthTransactor, error) {
	contract, err := bindBootloaderAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BootloaderAuthTransactor{contract: contract}, nil
}

// NewBootloaderAuthFilterer creates a new log filterer instance of BootloaderAuth, bound to a specific deployed contract.
func NewBootloaderAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*BootloaderAuthFilterer, error) {
	contract, err := bindBootloaderAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BootloaderAuthFilterer{contract: contract}, nil
}

// bindBootloaderAuth binds a generic wrapper to an already deployed contract.
func bindBootloaderAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BootloaderAuthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BootloaderAuth *BootloaderAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BootloaderAuth.Contract.BootloaderAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BootloaderAuth *BootloaderAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BootloaderAuth.Contract.BootloaderAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BootloaderAuth *BootloaderAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BootloaderAuth.Contract.BootloaderAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BootloaderAuth *BootloaderAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BootloaderAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BootloaderAuth *BootloaderAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BootloaderAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BootloaderAuth *BootloaderAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BootloaderAuth.Contract.contract.Transact(opts, method, params...)
}
