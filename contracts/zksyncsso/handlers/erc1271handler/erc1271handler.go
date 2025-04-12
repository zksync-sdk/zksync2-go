// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1271handler

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

// ERC1271HandlerMetaData contains all meta data concerning the ERC1271Handler contract.
var ERC1271HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notSelf\",\"type\":\"address\"}],\"name\":\"NOT_FROM_SELF\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OWNER_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OWNER_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_ERC165_FAIL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_NOT_FOUND\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isK1Owner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isModuleValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listK1Owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"k1OwnerList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listModuleValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC1271HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1271HandlerMetaData.ABI instead.
var ERC1271HandlerABI = ERC1271HandlerMetaData.ABI

// ERC1271Handler is an auto generated Go binding around an Ethereum contract.
type ERC1271Handler struct {
	ERC1271HandlerCaller     // Read-only binding to the contract
	ERC1271HandlerTransactor // Write-only binding to the contract
	ERC1271HandlerFilterer   // Log filterer for contract events
}

// ERC1271HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1271HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1271HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1271HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1271HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1271HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1271HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1271HandlerSession struct {
	Contract     *ERC1271Handler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1271HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1271HandlerCallerSession struct {
	Contract *ERC1271HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC1271HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1271HandlerTransactorSession struct {
	Contract     *ERC1271HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC1271HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1271HandlerRaw struct {
	Contract *ERC1271Handler // Generic contract binding to access the raw methods on
}

// ERC1271HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1271HandlerCallerRaw struct {
	Contract *ERC1271HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1271HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1271HandlerTransactorRaw struct {
	Contract *ERC1271HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1271Handler creates a new instance of ERC1271Handler, bound to a specific deployed contract.
func NewERC1271Handler(address common.Address, backend bind.ContractBackend) (*ERC1271Handler, error) {
	contract, err := bindERC1271Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1271Handler{ERC1271HandlerCaller: ERC1271HandlerCaller{contract: contract}, ERC1271HandlerTransactor: ERC1271HandlerTransactor{contract: contract}, ERC1271HandlerFilterer: ERC1271HandlerFilterer{contract: contract}}, nil
}

// NewERC1271HandlerCaller creates a new read-only instance of ERC1271Handler, bound to a specific deployed contract.
func NewERC1271HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC1271HandlerCaller, error) {
	contract, err := bindERC1271Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1271HandlerCaller{contract: contract}, nil
}

// NewERC1271HandlerTransactor creates a new write-only instance of ERC1271Handler, bound to a specific deployed contract.
func NewERC1271HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1271HandlerTransactor, error) {
	contract, err := bindERC1271Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1271HandlerTransactor{contract: contract}, nil
}

// NewERC1271HandlerFilterer creates a new log filterer instance of ERC1271Handler, bound to a specific deployed contract.
func NewERC1271HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1271HandlerFilterer, error) {
	contract, err := bindERC1271Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1271HandlerFilterer{contract: contract}, nil
}

// bindERC1271Handler binds a generic wrapper to an already deployed contract.
func bindERC1271Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1271HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1271Handler *ERC1271HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1271Handler.Contract.ERC1271HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1271Handler *ERC1271HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.ERC1271HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1271Handler *ERC1271HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.ERC1271HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1271Handler *ERC1271HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1271Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1271Handler *ERC1271HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1271Handler *ERC1271HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.contract.Transact(opts, method, params...)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_ERC1271Handler *ERC1271HandlerCaller) IsK1Owner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1271Handler.contract.Call(opts, &out, "isK1Owner", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_ERC1271Handler *ERC1271HandlerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _ERC1271Handler.Contract.IsK1Owner(&_ERC1271Handler.CallOpts, addr)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_ERC1271Handler *ERC1271HandlerCallerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _ERC1271Handler.Contract.IsK1Owner(&_ERC1271Handler.CallOpts, addr)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ERC1271Handler *ERC1271HandlerCaller) IsModuleValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1271Handler.contract.Call(opts, &out, "isModuleValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ERC1271Handler *ERC1271HandlerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _ERC1271Handler.Contract.IsModuleValidator(&_ERC1271Handler.CallOpts, validator)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ERC1271Handler *ERC1271HandlerCallerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _ERC1271Handler.Contract.IsModuleValidator(&_ERC1271Handler.CallOpts, validator)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ERC1271Handler *ERC1271HandlerCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ERC1271Handler.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ERC1271Handler *ERC1271HandlerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _ERC1271Handler.Contract.IsValidSignature(&_ERC1271Handler.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ERC1271Handler *ERC1271HandlerCallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _ERC1271Handler.Contract.IsValidSignature(&_ERC1271Handler.CallOpts, hash, signature)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_ERC1271Handler *ERC1271HandlerCaller) ListK1Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC1271Handler.contract.Call(opts, &out, "listK1Owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_ERC1271Handler *ERC1271HandlerSession) ListK1Owners() ([]common.Address, error) {
	return _ERC1271Handler.Contract.ListK1Owners(&_ERC1271Handler.CallOpts)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_ERC1271Handler *ERC1271HandlerCallerSession) ListK1Owners() ([]common.Address, error) {
	return _ERC1271Handler.Contract.ListK1Owners(&_ERC1271Handler.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ERC1271Handler *ERC1271HandlerCaller) ListModuleValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC1271Handler.contract.Call(opts, &out, "listModuleValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ERC1271Handler *ERC1271HandlerSession) ListModuleValidators() ([]common.Address, error) {
	return _ERC1271Handler.Contract.ListModuleValidators(&_ERC1271Handler.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ERC1271Handler *ERC1271HandlerCallerSession) ListModuleValidators() ([]common.Address, error) {
	return _ERC1271Handler.Contract.ListModuleValidators(&_ERC1271Handler.CallOpts)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_ERC1271Handler *ERC1271HandlerTransactor) AddK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ERC1271Handler.contract.Transact(opts, "addK1Owner", addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_ERC1271Handler *ERC1271HandlerSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.AddK1Owner(&_ERC1271Handler.TransactOpts, addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_ERC1271Handler *ERC1271HandlerTransactorSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.AddK1Owner(&_ERC1271Handler.TransactOpts, addr)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ERC1271Handler *ERC1271HandlerTransactor) AddModuleValidator(opts *bind.TransactOpts, validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.contract.Transact(opts, "addModuleValidator", validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ERC1271Handler *ERC1271HandlerSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.AddModuleValidator(&_ERC1271Handler.TransactOpts, validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ERC1271Handler *ERC1271HandlerTransactorSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.AddModuleValidator(&_ERC1271Handler.TransactOpts, validator, initData)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_ERC1271Handler *ERC1271HandlerTransactor) RemoveK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ERC1271Handler.contract.Transact(opts, "removeK1Owner", addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_ERC1271Handler *ERC1271HandlerSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.RemoveK1Owner(&_ERC1271Handler.TransactOpts, addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_ERC1271Handler *ERC1271HandlerTransactorSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.RemoveK1Owner(&_ERC1271Handler.TransactOpts, addr)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ERC1271Handler *ERC1271HandlerTransactor) RemoveModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.contract.Transact(opts, "removeModuleValidator", validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ERC1271Handler *ERC1271HandlerSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.RemoveModuleValidator(&_ERC1271Handler.TransactOpts, validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ERC1271Handler *ERC1271HandlerTransactorSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.RemoveModuleValidator(&_ERC1271Handler.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ERC1271Handler *ERC1271HandlerTransactor) UnlinkModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.contract.Transact(opts, "unlinkModuleValidator", validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ERC1271Handler *ERC1271HandlerSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.UnlinkModuleValidator(&_ERC1271Handler.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ERC1271Handler *ERC1271HandlerTransactorSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ERC1271Handler.Contract.UnlinkModuleValidator(&_ERC1271Handler.TransactOpts, validator, deinitData)
}

// ERC1271HandlerK1OwnerAddedIterator is returned from FilterK1OwnerAdded and is used to iterate over the raw logs and unpacked data for K1OwnerAdded events raised by the ERC1271Handler contract.
type ERC1271HandlerK1OwnerAddedIterator struct {
	Event *ERC1271HandlerK1OwnerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1271HandlerK1OwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1271HandlerK1OwnerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1271HandlerK1OwnerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1271HandlerK1OwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1271HandlerK1OwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1271HandlerK1OwnerAdded represents a K1OwnerAdded event raised by the ERC1271Handler contract.
type ERC1271HandlerK1OwnerAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerAdded is a free log retrieval operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_ERC1271Handler *ERC1271HandlerFilterer) FilterK1OwnerAdded(opts *bind.FilterOpts, addr []common.Address) (*ERC1271HandlerK1OwnerAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ERC1271Handler.contract.FilterLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &ERC1271HandlerK1OwnerAddedIterator{contract: _ERC1271Handler.contract, event: "K1OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchK1OwnerAdded is a free log subscription operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_ERC1271Handler *ERC1271HandlerFilterer) WatchK1OwnerAdded(opts *bind.WatchOpts, sink chan<- *ERC1271HandlerK1OwnerAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ERC1271Handler.contract.WatchLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1271HandlerK1OwnerAdded)
				if err := _ERC1271Handler.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseK1OwnerAdded is a log parse operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_ERC1271Handler *ERC1271HandlerFilterer) ParseK1OwnerAdded(log types.Log) (*ERC1271HandlerK1OwnerAdded, error) {
	event := new(ERC1271HandlerK1OwnerAdded)
	if err := _ERC1271Handler.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1271HandlerK1OwnerRemovedIterator is returned from FilterK1OwnerRemoved and is used to iterate over the raw logs and unpacked data for K1OwnerRemoved events raised by the ERC1271Handler contract.
type ERC1271HandlerK1OwnerRemovedIterator struct {
	Event *ERC1271HandlerK1OwnerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1271HandlerK1OwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1271HandlerK1OwnerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1271HandlerK1OwnerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1271HandlerK1OwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1271HandlerK1OwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1271HandlerK1OwnerRemoved represents a K1OwnerRemoved event raised by the ERC1271Handler contract.
type ERC1271HandlerK1OwnerRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerRemoved is a free log retrieval operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_ERC1271Handler *ERC1271HandlerFilterer) FilterK1OwnerRemoved(opts *bind.FilterOpts, addr []common.Address) (*ERC1271HandlerK1OwnerRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ERC1271Handler.contract.FilterLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &ERC1271HandlerK1OwnerRemovedIterator{contract: _ERC1271Handler.contract, event: "K1OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchK1OwnerRemoved is a free log subscription operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_ERC1271Handler *ERC1271HandlerFilterer) WatchK1OwnerRemoved(opts *bind.WatchOpts, sink chan<- *ERC1271HandlerK1OwnerRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ERC1271Handler.contract.WatchLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1271HandlerK1OwnerRemoved)
				if err := _ERC1271Handler.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseK1OwnerRemoved is a log parse operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_ERC1271Handler *ERC1271HandlerFilterer) ParseK1OwnerRemoved(log types.Log) (*ERC1271HandlerK1OwnerRemoved, error) {
	event := new(ERC1271HandlerK1OwnerRemoved)
	if err := _ERC1271Handler.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1271HandlerValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ERC1271Handler contract.
type ERC1271HandlerValidatorAddedIterator struct {
	Event *ERC1271HandlerValidatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1271HandlerValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1271HandlerValidatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1271HandlerValidatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1271HandlerValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1271HandlerValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1271HandlerValidatorAdded represents a ValidatorAdded event raised by the ERC1271Handler contract.
type ERC1271HandlerValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ERC1271Handler *ERC1271HandlerFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*ERC1271HandlerValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ERC1271Handler.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1271HandlerValidatorAddedIterator{contract: _ERC1271Handler.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ERC1271Handler *ERC1271HandlerFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ERC1271HandlerValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ERC1271Handler.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1271HandlerValidatorAdded)
				if err := _ERC1271Handler.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ERC1271Handler *ERC1271HandlerFilterer) ParseValidatorAdded(log types.Log) (*ERC1271HandlerValidatorAdded, error) {
	event := new(ERC1271HandlerValidatorAdded)
	if err := _ERC1271Handler.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1271HandlerValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the ERC1271Handler contract.
type ERC1271HandlerValidatorRemovedIterator struct {
	Event *ERC1271HandlerValidatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1271HandlerValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1271HandlerValidatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1271HandlerValidatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1271HandlerValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1271HandlerValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1271HandlerValidatorRemoved represents a ValidatorRemoved event raised by the ERC1271Handler contract.
type ERC1271HandlerValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ERC1271Handler *ERC1271HandlerFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*ERC1271HandlerValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ERC1271Handler.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1271HandlerValidatorRemovedIterator{contract: _ERC1271Handler.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ERC1271Handler *ERC1271HandlerFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ERC1271HandlerValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ERC1271Handler.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1271HandlerValidatorRemoved)
				if err := _ERC1271Handler.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ERC1271Handler *ERC1271HandlerFilterer) ParseValidatorRemoved(log types.Log) (*ERC1271HandlerValidatorRemoved, error) {
	event := new(ERC1271HandlerValidatorRemoved)
	if err := _ERC1271Handler.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
