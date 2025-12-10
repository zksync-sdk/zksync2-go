// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sessionlib

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

// SessionLibMetaData contains all meta data concerning the SessionLib contract.
var SessionLibMetaData = &bind.MetaData{
	ABI: "[]",
}

// SessionLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SessionLibMetaData.ABI instead.
var SessionLibABI = SessionLibMetaData.ABI

// SessionLib is an auto generated Go binding around an Ethereum contract.
type SessionLib struct {
	SessionLibCaller     // Read-only binding to the contract
	SessionLibTransactor // Write-only binding to the contract
	SessionLibFilterer   // Log filterer for contract events
}

// SessionLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SessionLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SessionLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SessionLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SessionLibSession struct {
	Contract     *SessionLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SessionLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SessionLibCallerSession struct {
	Contract *SessionLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SessionLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SessionLibTransactorSession struct {
	Contract     *SessionLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SessionLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SessionLibRaw struct {
	Contract *SessionLib // Generic contract binding to access the raw methods on
}

// SessionLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SessionLibCallerRaw struct {
	Contract *SessionLibCaller // Generic read-only contract binding to access the raw methods on
}

// SessionLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SessionLibTransactorRaw struct {
	Contract *SessionLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSessionLib creates a new instance of SessionLib, bound to a specific deployed contract.
func NewSessionLib(address common.Address, backend bind.ContractBackend) (*SessionLib, error) {
	contract, err := bindSessionLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SessionLib{SessionLibCaller: SessionLibCaller{contract: contract}, SessionLibTransactor: SessionLibTransactor{contract: contract}, SessionLibFilterer: SessionLibFilterer{contract: contract}}, nil
}

// NewSessionLibCaller creates a new read-only instance of SessionLib, bound to a specific deployed contract.
func NewSessionLibCaller(address common.Address, caller bind.ContractCaller) (*SessionLibCaller, error) {
	contract, err := bindSessionLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SessionLibCaller{contract: contract}, nil
}

// NewSessionLibTransactor creates a new write-only instance of SessionLib, bound to a specific deployed contract.
func NewSessionLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SessionLibTransactor, error) {
	contract, err := bindSessionLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SessionLibTransactor{contract: contract}, nil
}

// NewSessionLibFilterer creates a new log filterer instance of SessionLib, bound to a specific deployed contract.
func NewSessionLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SessionLibFilterer, error) {
	contract, err := bindSessionLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SessionLibFilterer{contract: contract}, nil
}

// bindSessionLib binds a generic wrapper to an already deployed contract.
func bindSessionLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SessionLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionLib *SessionLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionLib.Contract.SessionLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionLib *SessionLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionLib.Contract.SessionLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionLib *SessionLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionLib.Contract.SessionLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionLib *SessionLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionLib *SessionLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionLib *SessionLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionLib.Contract.contract.Transact(opts, method, params...)
}
