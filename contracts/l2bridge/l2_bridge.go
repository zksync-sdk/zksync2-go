// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2bridge

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

// IL2BridgeMetaData contains all meta data concerning the IL2Bridge contract.
var IL2BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"l1TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL2BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2BridgeMetaData.ABI instead.
var IL2BridgeABI = IL2BridgeMetaData.ABI

// IL2Bridge is an auto generated Go binding around an Ethereum contract.
type IL2Bridge struct {
	IL2BridgeCaller     // Read-only binding to the contract
	IL2BridgeTransactor // Write-only binding to the contract
	IL2BridgeFilterer   // Log filterer for contract events
}

// IL2BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2BridgeSession struct {
	Contract     *IL2Bridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL2BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2BridgeCallerSession struct {
	Contract *IL2BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IL2BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2BridgeTransactorSession struct {
	Contract     *IL2BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IL2BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2BridgeRaw struct {
	Contract *IL2Bridge // Generic contract binding to access the raw methods on
}

// IL2BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2BridgeCallerRaw struct {
	Contract *IL2BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IL2BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2BridgeTransactorRaw struct {
	Contract *IL2BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2Bridge creates a new instance of IL2Bridge, bound to a specific deployed contract.
func NewIL2Bridge(address common.Address, backend bind.ContractBackend) (*IL2Bridge, error) {
	contract, err := bindIL2Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2Bridge{IL2BridgeCaller: IL2BridgeCaller{contract: contract}, IL2BridgeTransactor: IL2BridgeTransactor{contract: contract}, IL2BridgeFilterer: IL2BridgeFilterer{contract: contract}}, nil
}

// NewIL2BridgeCaller creates a new read-only instance of IL2Bridge, bound to a specific deployed contract.
func NewIL2BridgeCaller(address common.Address, caller bind.ContractCaller) (*IL2BridgeCaller, error) {
	contract, err := bindIL2Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2BridgeCaller{contract: contract}, nil
}

// NewIL2BridgeTransactor creates a new write-only instance of IL2Bridge, bound to a specific deployed contract.
func NewIL2BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2BridgeTransactor, error) {
	contract, err := bindIL2Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2BridgeTransactor{contract: contract}, nil
}

// NewIL2BridgeFilterer creates a new log filterer instance of IL2Bridge, bound to a specific deployed contract.
func NewIL2BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2BridgeFilterer, error) {
	contract, err := bindIL2Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2BridgeFilterer{contract: contract}, nil
}

// bindIL2Bridge binds a generic wrapper to an already deployed contract.
func bindIL2Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL2BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2Bridge *IL2BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2Bridge.Contract.IL2BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2Bridge *IL2BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2Bridge.Contract.IL2BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2Bridge *IL2BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2Bridge.Contract.IL2BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2Bridge *IL2BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2Bridge *IL2BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2Bridge *IL2BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2Bridge.Contract.contract.Transact(opts, method, params...)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_IL2Bridge *IL2BridgeCaller) L1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL2Bridge.contract.Call(opts, &out, "l1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_IL2Bridge *IL2BridgeSession) L1Bridge() (common.Address, error) {
	return _IL2Bridge.Contract.L1Bridge(&_IL2Bridge.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_IL2Bridge *IL2BridgeCallerSession) L1Bridge() (common.Address, error) {
	return _IL2Bridge.Contract.L1Bridge(&_IL2Bridge.CallOpts)
}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address _l2Token) view returns(address)
func (_IL2Bridge *IL2BridgeCaller) L1TokenAddress(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL2Bridge.contract.Call(opts, &out, "l1TokenAddress", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address _l2Token) view returns(address)
func (_IL2Bridge *IL2BridgeSession) L1TokenAddress(_l2Token common.Address) (common.Address, error) {
	return _IL2Bridge.Contract.L1TokenAddress(&_IL2Bridge.CallOpts, _l2Token)
}

// L1TokenAddress is a free data retrieval call binding the contract method 0xf54266a2.
//
// Solidity: function l1TokenAddress(address _l2Token) view returns(address)
func (_IL2Bridge *IL2BridgeCallerSession) L1TokenAddress(_l2Token common.Address) (common.Address, error) {
	return _IL2Bridge.Contract.L1TokenAddress(&_IL2Bridge.CallOpts, _l2Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL2Bridge *IL2BridgeCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL2Bridge.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL2Bridge *IL2BridgeSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL2Bridge.Contract.L2TokenAddress(&_IL2Bridge.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL2Bridge *IL2BridgeCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL2Bridge.Contract.L2TokenAddress(&_IL2Bridge.CallOpts, _l1Token)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_IL2Bridge *IL2BridgeTransactor) FinalizeDeposit(opts *bind.TransactOpts, _l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2Bridge.contract.Transact(opts, "finalizeDeposit", _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_IL2Bridge *IL2BridgeSession) FinalizeDeposit(_l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2Bridge.Contract.FinalizeDeposit(&_IL2Bridge.TransactOpts, _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0xcfe7af7c.
//
// Solidity: function finalizeDeposit(address _l1Sender, address _l2Receiver, address _l1Token, uint256 _amount, bytes _data) returns()
func (_IL2Bridge *IL2BridgeTransactorSession) FinalizeDeposit(_l1Sender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2Bridge.Contract.FinalizeDeposit(&_IL2Bridge.TransactOpts, _l1Sender, _l2Receiver, _l1Token, _amount, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_IL2Bridge *IL2BridgeTransactor) Withdraw(opts *bind.TransactOpts, _l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL2Bridge.contract.Transact(opts, "withdraw", _l1Receiver, _l2Token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_IL2Bridge *IL2BridgeSession) Withdraw(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL2Bridge.Contract.Withdraw(&_IL2Bridge.TransactOpts, _l1Receiver, _l2Token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _l1Receiver, address _l2Token, uint256 _amount) returns()
func (_IL2Bridge *IL2BridgeTransactorSession) Withdraw(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL2Bridge.Contract.Withdraw(&_IL2Bridge.TransactOpts, _l1Receiver, _l2Token, _amount)
}
