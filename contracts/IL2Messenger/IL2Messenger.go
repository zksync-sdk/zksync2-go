// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IL2Messenger

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
)

// IL2MessengerMetaData contains all meta data concerning the IL2Messenger contract.
var IL2MessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendToL1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL2MessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2MessengerMetaData.ABI instead.
var IL2MessengerABI = IL2MessengerMetaData.ABI

// IL2Messenger is an auto generated Go binding around an Ethereum contract.
type IL2Messenger struct {
	IL2MessengerCaller     // Read-only binding to the contract
	IL2MessengerTransactor // Write-only binding to the contract
	IL2MessengerFilterer   // Log filterer for contract events
}

// IL2MessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2MessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2MessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2MessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2MessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2MessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2MessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2MessengerSession struct {
	Contract     *IL2Messenger     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL2MessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2MessengerCallerSession struct {
	Contract *IL2MessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IL2MessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2MessengerTransactorSession struct {
	Contract     *IL2MessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IL2MessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2MessengerRaw struct {
	Contract *IL2Messenger // Generic contract binding to access the raw methods on
}

// IL2MessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2MessengerCallerRaw struct {
	Contract *IL2MessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IL2MessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2MessengerTransactorRaw struct {
	Contract *IL2MessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2Messenger creates a new instance of IL2Messenger, bound to a specific deployed contract.
func NewIL2Messenger(address common.Address, backend bind.ContractBackend) (*IL2Messenger, error) {
	contract, err := bindIL2Messenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2Messenger{IL2MessengerCaller: IL2MessengerCaller{contract: contract}, IL2MessengerTransactor: IL2MessengerTransactor{contract: contract}, IL2MessengerFilterer: IL2MessengerFilterer{contract: contract}}, nil
}

// NewIL2MessengerCaller creates a new read-only instance of IL2Messenger, bound to a specific deployed contract.
func NewIL2MessengerCaller(address common.Address, caller bind.ContractCaller) (*IL2MessengerCaller, error) {
	contract, err := bindIL2Messenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2MessengerCaller{contract: contract}, nil
}

// NewIL2MessengerTransactor creates a new write-only instance of IL2Messenger, bound to a specific deployed contract.
func NewIL2MessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2MessengerTransactor, error) {
	contract, err := bindIL2Messenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2MessengerTransactor{contract: contract}, nil
}

// NewIL2MessengerFilterer creates a new log filterer instance of IL2Messenger, bound to a specific deployed contract.
func NewIL2MessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2MessengerFilterer, error) {
	contract, err := bindIL2Messenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2MessengerFilterer{contract: contract}, nil
}

// bindIL2Messenger binds a generic wrapper to an already deployed contract.
func bindIL2Messenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL2MessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2Messenger *IL2MessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2Messenger.Contract.IL2MessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2Messenger *IL2MessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2Messenger.Contract.IL2MessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2Messenger *IL2MessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2Messenger.Contract.IL2MessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2Messenger *IL2MessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2Messenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2Messenger *IL2MessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2Messenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2Messenger *IL2MessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2Messenger.Contract.contract.Transact(opts, method, params...)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_IL2Messenger *IL2MessengerTransactor) SendToL1(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _IL2Messenger.contract.Transact(opts, "sendToL1", _message)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_IL2Messenger *IL2MessengerSession) SendToL1(_message []byte) (*types.Transaction, error) {
	return _IL2Messenger.Contract.SendToL1(&_IL2Messenger.TransactOpts, _message)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_IL2Messenger *IL2MessengerTransactorSession) SendToL1(_message []byte) (*types.Transaction, error) {
	return _IL2Messenger.Contract.SendToL1(&_IL2Messenger.TransactOpts, _message)
}
