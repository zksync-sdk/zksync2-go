// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ssostorage

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

// SsoStorageMetaData contains all meta data concerning the SsoStorage contract.
var SsoStorageMetaData = &bind.MetaData{
	ABI: "[]",
}

// SsoStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SsoStorageMetaData.ABI instead.
var SsoStorageABI = SsoStorageMetaData.ABI

// SsoStorage is an auto generated Go binding around an Ethereum contract.
type SsoStorage struct {
	SsoStorageCaller     // Read-only binding to the contract
	SsoStorageTransactor // Write-only binding to the contract
	SsoStorageFilterer   // Log filterer for contract events
}

// SsoStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SsoStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SsoStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SsoStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SsoStorageSession struct {
	Contract     *SsoStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsoStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SsoStorageCallerSession struct {
	Contract *SsoStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SsoStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SsoStorageTransactorSession struct {
	Contract     *SsoStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SsoStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SsoStorageRaw struct {
	Contract *SsoStorage // Generic contract binding to access the raw methods on
}

// SsoStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SsoStorageCallerRaw struct {
	Contract *SsoStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SsoStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SsoStorageTransactorRaw struct {
	Contract *SsoStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSsoStorage creates a new instance of SsoStorage, bound to a specific deployed contract.
func NewSsoStorage(address common.Address, backend bind.ContractBackend) (*SsoStorage, error) {
	contract, err := bindSsoStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SsoStorage{SsoStorageCaller: SsoStorageCaller{contract: contract}, SsoStorageTransactor: SsoStorageTransactor{contract: contract}, SsoStorageFilterer: SsoStorageFilterer{contract: contract}}, nil
}

// NewSsoStorageCaller creates a new read-only instance of SsoStorage, bound to a specific deployed contract.
func NewSsoStorageCaller(address common.Address, caller bind.ContractCaller) (*SsoStorageCaller, error) {
	contract, err := bindSsoStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SsoStorageCaller{contract: contract}, nil
}

// NewSsoStorageTransactor creates a new write-only instance of SsoStorage, bound to a specific deployed contract.
func NewSsoStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SsoStorageTransactor, error) {
	contract, err := bindSsoStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SsoStorageTransactor{contract: contract}, nil
}

// NewSsoStorageFilterer creates a new log filterer instance of SsoStorage, bound to a specific deployed contract.
func NewSsoStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SsoStorageFilterer, error) {
	contract, err := bindSsoStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SsoStorageFilterer{contract: contract}, nil
}

// bindSsoStorage binds a generic wrapper to an already deployed contract.
func bindSsoStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SsoStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsoStorage *SsoStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsoStorage.Contract.SsoStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsoStorage *SsoStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoStorage.Contract.SsoStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsoStorage *SsoStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsoStorage.Contract.SsoStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsoStorage *SsoStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsoStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsoStorage *SsoStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsoStorage *SsoStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsoStorage.Contract.contract.Transact(opts, method, params...)
}
