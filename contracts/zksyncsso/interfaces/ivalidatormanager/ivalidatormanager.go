// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivalidatormanager

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

// IValidatorManagerMetaData contains all meta data concerning the IValidatorManager contract.
var IValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isModuleValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listModuleValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IValidatorManagerMetaData.ABI instead.
var IValidatorManagerABI = IValidatorManagerMetaData.ABI

// IValidatorManager is an auto generated Go binding around an Ethereum contract.
type IValidatorManager struct {
	IValidatorManagerCaller     // Read-only binding to the contract
	IValidatorManagerTransactor // Write-only binding to the contract
	IValidatorManagerFilterer   // Log filterer for contract events
}

// IValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IValidatorManagerSession struct {
	Contract     *IValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IValidatorManagerCallerSession struct {
	Contract *IValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IValidatorManagerTransactorSession struct {
	Contract     *IValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IValidatorManagerRaw struct {
	Contract *IValidatorManager // Generic contract binding to access the raw methods on
}

// IValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IValidatorManagerCallerRaw struct {
	Contract *IValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IValidatorManagerTransactorRaw struct {
	Contract *IValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIValidatorManager creates a new instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManager(address common.Address, backend bind.ContractBackend) (*IValidatorManager, error) {
	contract, err := bindIValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IValidatorManager{IValidatorManagerCaller: IValidatorManagerCaller{contract: contract}, IValidatorManagerTransactor: IValidatorManagerTransactor{contract: contract}, IValidatorManagerFilterer: IValidatorManagerFilterer{contract: contract}}, nil
}

// NewIValidatorManagerCaller creates a new read-only instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*IValidatorManagerCaller, error) {
	contract, err := bindIValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerCaller{contract: contract}, nil
}

// NewIValidatorManagerTransactor creates a new write-only instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IValidatorManagerTransactor, error) {
	contract, err := bindIValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerTransactor{contract: contract}, nil
}

// NewIValidatorManagerFilterer creates a new log filterer instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IValidatorManagerFilterer, error) {
	contract, err := bindIValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerFilterer{contract: contract}, nil
}

// bindIValidatorManager binds a generic wrapper to an already deployed contract.
func bindIValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidatorManager *IValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IValidatorManager.Contract.IValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidatorManager *IValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidatorManager.Contract.IValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidatorManager *IValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidatorManager.Contract.IValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidatorManager *IValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidatorManager *IValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidatorManager *IValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_IValidatorManager *IValidatorManagerCaller) IsModuleValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _IValidatorManager.contract.Call(opts, &out, "isModuleValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_IValidatorManager *IValidatorManagerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _IValidatorManager.Contract.IsModuleValidator(&_IValidatorManager.CallOpts, validator)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_IValidatorManager *IValidatorManagerCallerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _IValidatorManager.Contract.IsModuleValidator(&_IValidatorManager.CallOpts, validator)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_IValidatorManager *IValidatorManagerCaller) ListModuleValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IValidatorManager.contract.Call(opts, &out, "listModuleValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_IValidatorManager *IValidatorManagerSession) ListModuleValidators() ([]common.Address, error) {
	return _IValidatorManager.Contract.ListModuleValidators(&_IValidatorManager.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_IValidatorManager *IValidatorManagerCallerSession) ListModuleValidators() ([]common.Address, error) {
	return _IValidatorManager.Contract.ListModuleValidators(&_IValidatorManager.CallOpts)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_IValidatorManager *IValidatorManagerTransactor) AddModuleValidator(opts *bind.TransactOpts, validator common.Address, initData []byte) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "addModuleValidator", validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_IValidatorManager *IValidatorManagerSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.AddModuleValidator(&_IValidatorManager.TransactOpts, validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.AddModuleValidator(&_IValidatorManager.TransactOpts, validator, initData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_IValidatorManager *IValidatorManagerTransactor) RemoveModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "removeModuleValidator", validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_IValidatorManager *IValidatorManagerSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.RemoveModuleValidator(&_IValidatorManager.TransactOpts, validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.RemoveModuleValidator(&_IValidatorManager.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_IValidatorManager *IValidatorManagerTransactor) UnlinkModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "unlinkModuleValidator", validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_IValidatorManager *IValidatorManagerSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.UnlinkModuleValidator(&_IValidatorManager.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.UnlinkModuleValidator(&_IValidatorManager.TransactOpts, validator, deinitData)
}

// IValidatorManagerValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the IValidatorManager contract.
type IValidatorManagerValidatorAddedIterator struct {
	Event *IValidatorManagerValidatorAdded // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerValidatorAdded)
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
		it.Event = new(IValidatorManagerValidatorAdded)
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
func (it *IValidatorManagerValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerValidatorAdded represents a ValidatorAdded event raised by the IValidatorManager contract.
type IValidatorManagerValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_IValidatorManager *IValidatorManagerFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*IValidatorManagerValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerValidatorAddedIterator{contract: _IValidatorManager.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_IValidatorManager *IValidatorManagerFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *IValidatorManagerValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerValidatorAdded)
				if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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
func (_IValidatorManager *IValidatorManagerFilterer) ParseValidatorAdded(log types.Log) (*IValidatorManagerValidatorAdded, error) {
	event := new(IValidatorManagerValidatorAdded)
	if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IValidatorManagerValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the IValidatorManager contract.
type IValidatorManagerValidatorRemovedIterator struct {
	Event *IValidatorManagerValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerValidatorRemoved)
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
		it.Event = new(IValidatorManagerValidatorRemoved)
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
func (it *IValidatorManagerValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerValidatorRemoved represents a ValidatorRemoved event raised by the IValidatorManager contract.
type IValidatorManagerValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_IValidatorManager *IValidatorManagerFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*IValidatorManagerValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerValidatorRemovedIterator{contract: _IValidatorManager.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_IValidatorManager *IValidatorManagerFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *IValidatorManagerValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerValidatorRemoved)
				if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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
func (_IValidatorManager *IValidatorManagerFilterer) ParseValidatorRemoved(log types.Log) (*IValidatorManagerValidatorRemoved, error) {
	event := new(IValidatorManagerValidatorRemoved)
	if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
