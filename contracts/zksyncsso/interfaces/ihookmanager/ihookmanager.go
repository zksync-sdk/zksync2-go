// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ihookmanager

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

// IHookManagerMetaData contains all meta data concerning the IHookManager contract.
var IHookManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHook\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"listHooks\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"hookList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IHookManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IHookManagerMetaData.ABI instead.
var IHookManagerABI = IHookManagerMetaData.ABI

// IHookManager is an auto generated Go binding around an Ethereum contract.
type IHookManager struct {
	IHookManagerCaller     // Read-only binding to the contract
	IHookManagerTransactor // Write-only binding to the contract
	IHookManagerFilterer   // Log filterer for contract events
}

// IHookManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHookManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHookManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHookManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHookManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHookManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHookManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHookManagerSession struct {
	Contract     *IHookManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHookManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHookManagerCallerSession struct {
	Contract *IHookManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IHookManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHookManagerTransactorSession struct {
	Contract     *IHookManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IHookManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHookManagerRaw struct {
	Contract *IHookManager // Generic contract binding to access the raw methods on
}

// IHookManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHookManagerCallerRaw struct {
	Contract *IHookManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IHookManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHookManagerTransactorRaw struct {
	Contract *IHookManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHookManager creates a new instance of IHookManager, bound to a specific deployed contract.
func NewIHookManager(address common.Address, backend bind.ContractBackend) (*IHookManager, error) {
	contract, err := bindIHookManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHookManager{IHookManagerCaller: IHookManagerCaller{contract: contract}, IHookManagerTransactor: IHookManagerTransactor{contract: contract}, IHookManagerFilterer: IHookManagerFilterer{contract: contract}}, nil
}

// NewIHookManagerCaller creates a new read-only instance of IHookManager, bound to a specific deployed contract.
func NewIHookManagerCaller(address common.Address, caller bind.ContractCaller) (*IHookManagerCaller, error) {
	contract, err := bindIHookManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHookManagerCaller{contract: contract}, nil
}

// NewIHookManagerTransactor creates a new write-only instance of IHookManager, bound to a specific deployed contract.
func NewIHookManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IHookManagerTransactor, error) {
	contract, err := bindIHookManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHookManagerTransactor{contract: contract}, nil
}

// NewIHookManagerFilterer creates a new log filterer instance of IHookManager, bound to a specific deployed contract.
func NewIHookManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IHookManagerFilterer, error) {
	contract, err := bindIHookManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHookManagerFilterer{contract: contract}, nil
}

// bindIHookManager binds a generic wrapper to an already deployed contract.
func bindIHookManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHookManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHookManager *IHookManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHookManager.Contract.IHookManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHookManager *IHookManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHookManager.Contract.IHookManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHookManager *IHookManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHookManager.Contract.IHookManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHookManager *IHookManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHookManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHookManager *IHookManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHookManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHookManager *IHookManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHookManager.Contract.contract.Transact(opts, method, params...)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_IHookManager *IHookManagerCaller) IsHook(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _IHookManager.contract.Call(opts, &out, "isHook", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_IHookManager *IHookManagerSession) IsHook(addr common.Address) (bool, error) {
	return _IHookManager.Contract.IsHook(&_IHookManager.CallOpts, addr)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_IHookManager *IHookManagerCallerSession) IsHook(addr common.Address) (bool, error) {
	return _IHookManager.Contract.IsHook(&_IHookManager.CallOpts, addr)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_IHookManager *IHookManagerCaller) ListHooks(opts *bind.CallOpts, isValidation bool) ([]common.Address, error) {
	var out []interface{}
	err := _IHookManager.contract.Call(opts, &out, "listHooks", isValidation)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_IHookManager *IHookManagerSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _IHookManager.Contract.ListHooks(&_IHookManager.CallOpts, isValidation)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_IHookManager *IHookManagerCallerSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _IHookManager.Contract.ListHooks(&_IHookManager.CallOpts, isValidation)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_IHookManager *IHookManagerTransactor) AddHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _IHookManager.contract.Transact(opts, "addHook", hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_IHookManager *IHookManagerSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _IHookManager.Contract.AddHook(&_IHookManager.TransactOpts, hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_IHookManager *IHookManagerTransactorSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _IHookManager.Contract.AddHook(&_IHookManager.TransactOpts, hook, isValidation, initData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_IHookManager *IHookManagerTransactor) RemoveHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _IHookManager.contract.Transact(opts, "removeHook", hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_IHookManager *IHookManagerSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _IHookManager.Contract.RemoveHook(&_IHookManager.TransactOpts, hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_IHookManager *IHookManagerTransactorSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _IHookManager.Contract.RemoveHook(&_IHookManager.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_IHookManager *IHookManagerTransactor) UnlinkHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _IHookManager.contract.Transact(opts, "unlinkHook", hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_IHookManager *IHookManagerSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _IHookManager.Contract.UnlinkHook(&_IHookManager.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_IHookManager *IHookManagerTransactorSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _IHookManager.Contract.UnlinkHook(&_IHookManager.TransactOpts, hook, isValidation, deinitData)
}

// IHookManagerHookAddedIterator is returned from FilterHookAdded and is used to iterate over the raw logs and unpacked data for HookAdded events raised by the IHookManager contract.
type IHookManagerHookAddedIterator struct {
	Event *IHookManagerHookAdded // Event containing the contract specifics and raw log

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
func (it *IHookManagerHookAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHookManagerHookAdded)
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
		it.Event = new(IHookManagerHookAdded)
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
func (it *IHookManagerHookAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHookManagerHookAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHookManagerHookAdded represents a HookAdded event raised by the IHookManager contract.
type IHookManagerHookAdded struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookAdded is a free log retrieval operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_IHookManager *IHookManagerFilterer) FilterHookAdded(opts *bind.FilterOpts, hook []common.Address) (*IHookManagerHookAddedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _IHookManager.contract.FilterLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return &IHookManagerHookAddedIterator{contract: _IHookManager.contract, event: "HookAdded", logs: logs, sub: sub}, nil
}

// WatchHookAdded is a free log subscription operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_IHookManager *IHookManagerFilterer) WatchHookAdded(opts *bind.WatchOpts, sink chan<- *IHookManagerHookAdded, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _IHookManager.contract.WatchLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHookManagerHookAdded)
				if err := _IHookManager.contract.UnpackLog(event, "HookAdded", log); err != nil {
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

// ParseHookAdded is a log parse operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_IHookManager *IHookManagerFilterer) ParseHookAdded(log types.Log) (*IHookManagerHookAdded, error) {
	event := new(IHookManagerHookAdded)
	if err := _IHookManager.contract.UnpackLog(event, "HookAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHookManagerHookRemovedIterator is returned from FilterHookRemoved and is used to iterate over the raw logs and unpacked data for HookRemoved events raised by the IHookManager contract.
type IHookManagerHookRemovedIterator struct {
	Event *IHookManagerHookRemoved // Event containing the contract specifics and raw log

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
func (it *IHookManagerHookRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHookManagerHookRemoved)
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
		it.Event = new(IHookManagerHookRemoved)
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
func (it *IHookManagerHookRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHookManagerHookRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHookManagerHookRemoved represents a HookRemoved event raised by the IHookManager contract.
type IHookManagerHookRemoved struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookRemoved is a free log retrieval operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_IHookManager *IHookManagerFilterer) FilterHookRemoved(opts *bind.FilterOpts, hook []common.Address) (*IHookManagerHookRemovedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _IHookManager.contract.FilterLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return &IHookManagerHookRemovedIterator{contract: _IHookManager.contract, event: "HookRemoved", logs: logs, sub: sub}, nil
}

// WatchHookRemoved is a free log subscription operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_IHookManager *IHookManagerFilterer) WatchHookRemoved(opts *bind.WatchOpts, sink chan<- *IHookManagerHookRemoved, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _IHookManager.contract.WatchLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHookManagerHookRemoved)
				if err := _IHookManager.contract.UnpackLog(event, "HookRemoved", log); err != nil {
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

// ParseHookRemoved is a log parse operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_IHookManager *IHookManagerFilterer) ParseHookRemoved(log types.Log) (*IHookManagerHookRemoved, error) {
	event := new(IHookManagerHookRemoved)
	if err := _IHookManager.contract.UnpackLog(event, "HookRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
