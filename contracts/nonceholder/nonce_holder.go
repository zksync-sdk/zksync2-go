// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nonceholder

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

// INonceHolderMetaData contains all meta data concerning the INonceHolder contract.
var INonceHolderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValueSetUnderNonce\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getDeploymentNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getMinNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getRawNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"}],\"name\":\"getValueUnderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"increaseMinNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"incrementDeploymentNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expectedNonce\",\"type\":\"uint256\"}],\"name\":\"incrementMinNonceIfEquals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"isNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setValueUnderNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_key\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_shouldBeUsed\",\"type\":\"bool\"}],\"name\":\"validateNonceUsage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// INonceHolderABI is the input ABI used to generate the binding from.
// Deprecated: Use INonceHolderMetaData.ABI instead.
var INonceHolderABI = INonceHolderMetaData.ABI

// INonceHolder is an auto generated Go binding around an Ethereum contract.
type INonceHolder struct {
	INonceHolderCaller     // Read-only binding to the contract
	INonceHolderTransactor // Write-only binding to the contract
	INonceHolderFilterer   // Log filterer for contract events
}

// INonceHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type INonceHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonceHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INonceHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonceHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INonceHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonceHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INonceHolderSession struct {
	Contract     *INonceHolder     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INonceHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INonceHolderCallerSession struct {
	Contract *INonceHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// INonceHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INonceHolderTransactorSession struct {
	Contract     *INonceHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// INonceHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type INonceHolderRaw struct {
	Contract *INonceHolder // Generic contract binding to access the raw methods on
}

// INonceHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INonceHolderCallerRaw struct {
	Contract *INonceHolderCaller // Generic read-only contract binding to access the raw methods on
}

// INonceHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INonceHolderTransactorRaw struct {
	Contract *INonceHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINonceHolder creates a new instance of INonceHolder, bound to a specific deployed contract.
func NewINonceHolder(address common.Address, backend bind.ContractBackend) (*INonceHolder, error) {
	contract, err := bindINonceHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INonceHolder{INonceHolderCaller: INonceHolderCaller{contract: contract}, INonceHolderTransactor: INonceHolderTransactor{contract: contract}, INonceHolderFilterer: INonceHolderFilterer{contract: contract}}, nil
}

// NewINonceHolderCaller creates a new read-only instance of INonceHolder, bound to a specific deployed contract.
func NewINonceHolderCaller(address common.Address, caller bind.ContractCaller) (*INonceHolderCaller, error) {
	contract, err := bindINonceHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INonceHolderCaller{contract: contract}, nil
}

// NewINonceHolderTransactor creates a new write-only instance of INonceHolder, bound to a specific deployed contract.
func NewINonceHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*INonceHolderTransactor, error) {
	contract, err := bindINonceHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INonceHolderTransactor{contract: contract}, nil
}

// NewINonceHolderFilterer creates a new log filterer instance of INonceHolder, bound to a specific deployed contract.
func NewINonceHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*INonceHolderFilterer, error) {
	contract, err := bindINonceHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INonceHolderFilterer{contract: contract}, nil
}

// bindINonceHolder binds a generic wrapper to an already deployed contract.
func bindINonceHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INonceHolderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonceHolder *INonceHolderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonceHolder.Contract.INonceHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonceHolder *INonceHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonceHolder.Contract.INonceHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonceHolder *INonceHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonceHolder.Contract.INonceHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonceHolder *INonceHolderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonceHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonceHolder *INonceHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonceHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonceHolder *INonceHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonceHolder.Contract.contract.Transact(opts, method, params...)
}

// GetDeploymentNonce is a free data retrieval call binding the contract method 0xfb1a9a57.
//
// Solidity: function getDeploymentNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderCaller) GetDeploymentNonce(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _INonceHolder.contract.Call(opts, &out, "getDeploymentNonce", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeploymentNonce is a free data retrieval call binding the contract method 0xfb1a9a57.
//
// Solidity: function getDeploymentNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderSession) GetDeploymentNonce(_address common.Address) (*big.Int, error) {
	return _INonceHolder.Contract.GetDeploymentNonce(&_INonceHolder.CallOpts, _address)
}

// GetDeploymentNonce is a free data retrieval call binding the contract method 0xfb1a9a57.
//
// Solidity: function getDeploymentNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderCallerSession) GetDeploymentNonce(_address common.Address) (*big.Int, error) {
	return _INonceHolder.Contract.GetDeploymentNonce(&_INonceHolder.CallOpts, _address)
}

// GetMinNonce is a free data retrieval call binding the contract method 0x896909dc.
//
// Solidity: function getMinNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderCaller) GetMinNonce(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _INonceHolder.contract.Call(opts, &out, "getMinNonce", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNonce is a free data retrieval call binding the contract method 0x896909dc.
//
// Solidity: function getMinNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderSession) GetMinNonce(_address common.Address) (*big.Int, error) {
	return _INonceHolder.Contract.GetMinNonce(&_INonceHolder.CallOpts, _address)
}

// GetMinNonce is a free data retrieval call binding the contract method 0x896909dc.
//
// Solidity: function getMinNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderCallerSession) GetMinNonce(_address common.Address) (*big.Int, error) {
	return _INonceHolder.Contract.GetMinNonce(&_INonceHolder.CallOpts, _address)
}

// GetRawNonce is a free data retrieval call binding the contract method 0x5aa9b6b5.
//
// Solidity: function getRawNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderCaller) GetRawNonce(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _INonceHolder.contract.Call(opts, &out, "getRawNonce", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRawNonce is a free data retrieval call binding the contract method 0x5aa9b6b5.
//
// Solidity: function getRawNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderSession) GetRawNonce(_address common.Address) (*big.Int, error) {
	return _INonceHolder.Contract.GetRawNonce(&_INonceHolder.CallOpts, _address)
}

// GetRawNonce is a free data retrieval call binding the contract method 0x5aa9b6b5.
//
// Solidity: function getRawNonce(address _address) view returns(uint256)
func (_INonceHolder *INonceHolderCallerSession) GetRawNonce(_address common.Address) (*big.Int, error) {
	return _INonceHolder.Contract.GetRawNonce(&_INonceHolder.CallOpts, _address)
}

// GetValueUnderNonce is a free data retrieval call binding the contract method 0x55d35d18.
//
// Solidity: function getValueUnderNonce(uint256 _key) view returns(uint256)
func (_INonceHolder *INonceHolderCaller) GetValueUnderNonce(opts *bind.CallOpts, _key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INonceHolder.contract.Call(opts, &out, "getValueUnderNonce", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValueUnderNonce is a free data retrieval call binding the contract method 0x55d35d18.
//
// Solidity: function getValueUnderNonce(uint256 _key) view returns(uint256)
func (_INonceHolder *INonceHolderSession) GetValueUnderNonce(_key *big.Int) (*big.Int, error) {
	return _INonceHolder.Contract.GetValueUnderNonce(&_INonceHolder.CallOpts, _key)
}

// GetValueUnderNonce is a free data retrieval call binding the contract method 0x55d35d18.
//
// Solidity: function getValueUnderNonce(uint256 _key) view returns(uint256)
func (_INonceHolder *INonceHolderCallerSession) GetValueUnderNonce(_key *big.Int) (*big.Int, error) {
	return _INonceHolder.Contract.GetValueUnderNonce(&_INonceHolder.CallOpts, _key)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address _address, uint256 _nonce) view returns(bool)
func (_INonceHolder *INonceHolderCaller) IsNonceUsed(opts *bind.CallOpts, _address common.Address, _nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _INonceHolder.contract.Call(opts, &out, "isNonceUsed", _address, _nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address _address, uint256 _nonce) view returns(bool)
func (_INonceHolder *INonceHolderSession) IsNonceUsed(_address common.Address, _nonce *big.Int) (bool, error) {
	return _INonceHolder.Contract.IsNonceUsed(&_INonceHolder.CallOpts, _address, _nonce)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0xcab7e8eb.
//
// Solidity: function isNonceUsed(address _address, uint256 _nonce) view returns(bool)
func (_INonceHolder *INonceHolderCallerSession) IsNonceUsed(_address common.Address, _nonce *big.Int) (bool, error) {
	return _INonceHolder.Contract.IsNonceUsed(&_INonceHolder.CallOpts, _address, _nonce)
}

// ValidateNonceUsage is a free data retrieval call binding the contract method 0x6ee1dc20.
//
// Solidity: function validateNonceUsage(address _address, uint256 _key, bool _shouldBeUsed) view returns()
func (_INonceHolder *INonceHolderCaller) ValidateNonceUsage(opts *bind.CallOpts, _address common.Address, _key *big.Int, _shouldBeUsed bool) error {
	var out []interface{}
	err := _INonceHolder.contract.Call(opts, &out, "validateNonceUsage", _address, _key, _shouldBeUsed)

	if err != nil {
		return err
	}

	return err

}

// ValidateNonceUsage is a free data retrieval call binding the contract method 0x6ee1dc20.
//
// Solidity: function validateNonceUsage(address _address, uint256 _key, bool _shouldBeUsed) view returns()
func (_INonceHolder *INonceHolderSession) ValidateNonceUsage(_address common.Address, _key *big.Int, _shouldBeUsed bool) error {
	return _INonceHolder.Contract.ValidateNonceUsage(&_INonceHolder.CallOpts, _address, _key, _shouldBeUsed)
}

// ValidateNonceUsage is a free data retrieval call binding the contract method 0x6ee1dc20.
//
// Solidity: function validateNonceUsage(address _address, uint256 _key, bool _shouldBeUsed) view returns()
func (_INonceHolder *INonceHolderCallerSession) ValidateNonceUsage(_address common.Address, _key *big.Int, _shouldBeUsed bool) error {
	return _INonceHolder.Contract.ValidateNonceUsage(&_INonceHolder.CallOpts, _address, _key, _shouldBeUsed)
}

// IncreaseMinNonce is a paid mutator transaction binding the contract method 0x38a78092.
//
// Solidity: function increaseMinNonce(uint256 _value) returns(uint256)
func (_INonceHolder *INonceHolderTransactor) IncreaseMinNonce(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _INonceHolder.contract.Transact(opts, "increaseMinNonce", _value)
}

// IncreaseMinNonce is a paid mutator transaction binding the contract method 0x38a78092.
//
// Solidity: function increaseMinNonce(uint256 _value) returns(uint256)
func (_INonceHolder *INonceHolderSession) IncreaseMinNonce(_value *big.Int) (*types.Transaction, error) {
	return _INonceHolder.Contract.IncreaseMinNonce(&_INonceHolder.TransactOpts, _value)
}

// IncreaseMinNonce is a paid mutator transaction binding the contract method 0x38a78092.
//
// Solidity: function increaseMinNonce(uint256 _value) returns(uint256)
func (_INonceHolder *INonceHolderTransactorSession) IncreaseMinNonce(_value *big.Int) (*types.Transaction, error) {
	return _INonceHolder.Contract.IncreaseMinNonce(&_INonceHolder.TransactOpts, _value)
}

// IncrementDeploymentNonce is a paid mutator transaction binding the contract method 0x306395c6.
//
// Solidity: function incrementDeploymentNonce(address _address) returns(uint256)
func (_INonceHolder *INonceHolderTransactor) IncrementDeploymentNonce(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _INonceHolder.contract.Transact(opts, "incrementDeploymentNonce", _address)
}

// IncrementDeploymentNonce is a paid mutator transaction binding the contract method 0x306395c6.
//
// Solidity: function incrementDeploymentNonce(address _address) returns(uint256)
func (_INonceHolder *INonceHolderSession) IncrementDeploymentNonce(_address common.Address) (*types.Transaction, error) {
	return _INonceHolder.Contract.IncrementDeploymentNonce(&_INonceHolder.TransactOpts, _address)
}

// IncrementDeploymentNonce is a paid mutator transaction binding the contract method 0x306395c6.
//
// Solidity: function incrementDeploymentNonce(address _address) returns(uint256)
func (_INonceHolder *INonceHolderTransactorSession) IncrementDeploymentNonce(_address common.Address) (*types.Transaction, error) {
	return _INonceHolder.Contract.IncrementDeploymentNonce(&_INonceHolder.TransactOpts, _address)
}

// IncrementMinNonceIfEquals is a paid mutator transaction binding the contract method 0xe1239cd8.
//
// Solidity: function incrementMinNonceIfEquals(uint256 _expectedNonce) returns()
func (_INonceHolder *INonceHolderTransactor) IncrementMinNonceIfEquals(opts *bind.TransactOpts, _expectedNonce *big.Int) (*types.Transaction, error) {
	return _INonceHolder.contract.Transact(opts, "incrementMinNonceIfEquals", _expectedNonce)
}

// IncrementMinNonceIfEquals is a paid mutator transaction binding the contract method 0xe1239cd8.
//
// Solidity: function incrementMinNonceIfEquals(uint256 _expectedNonce) returns()
func (_INonceHolder *INonceHolderSession) IncrementMinNonceIfEquals(_expectedNonce *big.Int) (*types.Transaction, error) {
	return _INonceHolder.Contract.IncrementMinNonceIfEquals(&_INonceHolder.TransactOpts, _expectedNonce)
}

// IncrementMinNonceIfEquals is a paid mutator transaction binding the contract method 0xe1239cd8.
//
// Solidity: function incrementMinNonceIfEquals(uint256 _expectedNonce) returns()
func (_INonceHolder *INonceHolderTransactorSession) IncrementMinNonceIfEquals(_expectedNonce *big.Int) (*types.Transaction, error) {
	return _INonceHolder.Contract.IncrementMinNonceIfEquals(&_INonceHolder.TransactOpts, _expectedNonce)
}

// SetValueUnderNonce is a paid mutator transaction binding the contract method 0x155fd27a.
//
// Solidity: function setValueUnderNonce(uint256 _key, uint256 _value) returns()
func (_INonceHolder *INonceHolderTransactor) SetValueUnderNonce(opts *bind.TransactOpts, _key *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _INonceHolder.contract.Transact(opts, "setValueUnderNonce", _key, _value)
}

// SetValueUnderNonce is a paid mutator transaction binding the contract method 0x155fd27a.
//
// Solidity: function setValueUnderNonce(uint256 _key, uint256 _value) returns()
func (_INonceHolder *INonceHolderSession) SetValueUnderNonce(_key *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _INonceHolder.Contract.SetValueUnderNonce(&_INonceHolder.TransactOpts, _key, _value)
}

// SetValueUnderNonce is a paid mutator transaction binding the contract method 0x155fd27a.
//
// Solidity: function setValueUnderNonce(uint256 _key, uint256 _value) returns()
func (_INonceHolder *INonceHolderTransactorSession) SetValueUnderNonce(_key *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _INonceHolder.Contract.SetValueUnderNonce(&_INonceHolder.TransactOpts, _key, _value)
}

// INonceHolderValueSetUnderNonceIterator is returned from FilterValueSetUnderNonce and is used to iterate over the raw logs and unpacked data for ValueSetUnderNonce events raised by the INonceHolder contract.
type INonceHolderValueSetUnderNonceIterator struct {
	Event *INonceHolderValueSetUnderNonce // Event containing the contract specifics and raw log

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
func (it *INonceHolderValueSetUnderNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INonceHolderValueSetUnderNonce)
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
		it.Event = new(INonceHolderValueSetUnderNonce)
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
func (it *INonceHolderValueSetUnderNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INonceHolderValueSetUnderNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INonceHolderValueSetUnderNonce represents a ValueSetUnderNonce event raised by the INonceHolder contract.
type INonceHolderValueSetUnderNonce struct {
	AccountAddress common.Address
	Key            *big.Int
	Value          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValueSetUnderNonce is a free log retrieval operation binding the contract event 0xda2b716e5a5d5f602b9a5842bcd89c215b125258dfea271a03e5e0e801d93a8c.
//
// Solidity: event ValueSetUnderNonce(address indexed accountAddress, uint256 indexed key, uint256 value)
func (_INonceHolder *INonceHolderFilterer) FilterValueSetUnderNonce(opts *bind.FilterOpts, accountAddress []common.Address, key []*big.Int) (*INonceHolderValueSetUnderNonceIterator, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _INonceHolder.contract.FilterLogs(opts, "ValueSetUnderNonce", accountAddressRule, keyRule)
	if err != nil {
		return nil, err
	}
	return &INonceHolderValueSetUnderNonceIterator{contract: _INonceHolder.contract, event: "ValueSetUnderNonce", logs: logs, sub: sub}, nil
}

// WatchValueSetUnderNonce is a free log subscription operation binding the contract event 0xda2b716e5a5d5f602b9a5842bcd89c215b125258dfea271a03e5e0e801d93a8c.
//
// Solidity: event ValueSetUnderNonce(address indexed accountAddress, uint256 indexed key, uint256 value)
func (_INonceHolder *INonceHolderFilterer) WatchValueSetUnderNonce(opts *bind.WatchOpts, sink chan<- *INonceHolderValueSetUnderNonce, accountAddress []common.Address, key []*big.Int) (event.Subscription, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}
	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _INonceHolder.contract.WatchLogs(opts, "ValueSetUnderNonce", accountAddressRule, keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INonceHolderValueSetUnderNonce)
				if err := _INonceHolder.contract.UnpackLog(event, "ValueSetUnderNonce", log); err != nil {
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

// ParseValueSetUnderNonce is a log parse operation binding the contract event 0xda2b716e5a5d5f602b9a5842bcd89c215b125258dfea271a03e5e0e801d93a8c.
//
// Solidity: event ValueSetUnderNonce(address indexed accountAddress, uint256 indexed key, uint256 value)
func (_INonceHolder *INonceHolderFilterer) ParseValueSetUnderNonce(log types.Log) (*INonceHolderValueSetUnderNonce, error) {
	event := new(INonceHolderValueSetUnderNonce)
	if err := _INonceHolder.contract.UnpackLog(event, "ValueSetUnderNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
