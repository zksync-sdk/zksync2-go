// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hookmanager

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

// HookManagerMetaData contains all meta data concerning the HookManager contract.
var HookManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hookAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_ERC165_FAIL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notSelf\",\"type\":\"address\"}],\"name\":\"NOT_FROM_SELF\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_ALREADY_EXISTS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHook\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"listHooks\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"hookList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HookManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use HookManagerMetaData.ABI instead.
var HookManagerABI = HookManagerMetaData.ABI

// HookManager is an auto generated Go binding around an Ethereum contract.
type HookManager struct {
	HookManagerCaller     // Read-only binding to the contract
	HookManagerTransactor // Write-only binding to the contract
	HookManagerFilterer   // Log filterer for contract events
}

// HookManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type HookManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HookManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HookManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HookManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HookManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HookManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HookManagerSession struct {
	Contract     *HookManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HookManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HookManagerCallerSession struct {
	Contract *HookManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HookManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HookManagerTransactorSession struct {
	Contract     *HookManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HookManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type HookManagerRaw struct {
	Contract *HookManager // Generic contract binding to access the raw methods on
}

// HookManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HookManagerCallerRaw struct {
	Contract *HookManagerCaller // Generic read-only contract binding to access the raw methods on
}

// HookManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HookManagerTransactorRaw struct {
	Contract *HookManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHookManager creates a new instance of HookManager, bound to a specific deployed contract.
func NewHookManager(address common.Address, backend bind.ContractBackend) (*HookManager, error) {
	contract, err := bindHookManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HookManager{HookManagerCaller: HookManagerCaller{contract: contract}, HookManagerTransactor: HookManagerTransactor{contract: contract}, HookManagerFilterer: HookManagerFilterer{contract: contract}}, nil
}

// NewHookManagerCaller creates a new read-only instance of HookManager, bound to a specific deployed contract.
func NewHookManagerCaller(address common.Address, caller bind.ContractCaller) (*HookManagerCaller, error) {
	contract, err := bindHookManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HookManagerCaller{contract: contract}, nil
}

// NewHookManagerTransactor creates a new write-only instance of HookManager, bound to a specific deployed contract.
func NewHookManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*HookManagerTransactor, error) {
	contract, err := bindHookManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HookManagerTransactor{contract: contract}, nil
}

// NewHookManagerFilterer creates a new log filterer instance of HookManager, bound to a specific deployed contract.
func NewHookManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*HookManagerFilterer, error) {
	contract, err := bindHookManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HookManagerFilterer{contract: contract}, nil
}

// bindHookManager binds a generic wrapper to an already deployed contract.
func bindHookManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HookManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HookManager *HookManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HookManager.Contract.HookManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HookManager *HookManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HookManager.Contract.HookManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HookManager *HookManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HookManager.Contract.HookManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HookManager *HookManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HookManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HookManager *HookManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HookManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HookManager *HookManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HookManager.Contract.contract.Transact(opts, method, params...)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_HookManager *HookManagerCaller) IsHook(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _HookManager.contract.Call(opts, &out, "isHook", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_HookManager *HookManagerSession) IsHook(addr common.Address) (bool, error) {
	return _HookManager.Contract.IsHook(&_HookManager.CallOpts, addr)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_HookManager *HookManagerCallerSession) IsHook(addr common.Address) (bool, error) {
	return _HookManager.Contract.IsHook(&_HookManager.CallOpts, addr)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_HookManager *HookManagerCaller) ListHooks(opts *bind.CallOpts, isValidation bool) ([]common.Address, error) {
	var out []interface{}
	err := _HookManager.contract.Call(opts, &out, "listHooks", isValidation)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_HookManager *HookManagerSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _HookManager.Contract.ListHooks(&_HookManager.CallOpts, isValidation)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_HookManager *HookManagerCallerSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _HookManager.Contract.ListHooks(&_HookManager.CallOpts, isValidation)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_HookManager *HookManagerTransactor) AddHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _HookManager.contract.Transact(opts, "addHook", hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_HookManager *HookManagerSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _HookManager.Contract.AddHook(&_HookManager.TransactOpts, hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_HookManager *HookManagerTransactorSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _HookManager.Contract.AddHook(&_HookManager.TransactOpts, hook, isValidation, initData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_HookManager *HookManagerTransactor) RemoveHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _HookManager.contract.Transact(opts, "removeHook", hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_HookManager *HookManagerSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _HookManager.Contract.RemoveHook(&_HookManager.TransactOpts, hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_HookManager *HookManagerTransactorSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _HookManager.Contract.RemoveHook(&_HookManager.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_HookManager *HookManagerTransactor) UnlinkHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _HookManager.contract.Transact(opts, "unlinkHook", hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_HookManager *HookManagerSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _HookManager.Contract.UnlinkHook(&_HookManager.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_HookManager *HookManagerTransactorSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _HookManager.Contract.UnlinkHook(&_HookManager.TransactOpts, hook, isValidation, deinitData)
}

// HookManagerHookAddedIterator is returned from FilterHookAdded and is used to iterate over the raw logs and unpacked data for HookAdded events raised by the HookManager contract.
type HookManagerHookAddedIterator struct {
	Event *HookManagerHookAdded // Event containing the contract specifics and raw log

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
func (it *HookManagerHookAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HookManagerHookAdded)
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
		it.Event = new(HookManagerHookAdded)
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
func (it *HookManagerHookAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HookManagerHookAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HookManagerHookAdded represents a HookAdded event raised by the HookManager contract.
type HookManagerHookAdded struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookAdded is a free log retrieval operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_HookManager *HookManagerFilterer) FilterHookAdded(opts *bind.FilterOpts, hook []common.Address) (*HookManagerHookAddedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _HookManager.contract.FilterLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return &HookManagerHookAddedIterator{contract: _HookManager.contract, event: "HookAdded", logs: logs, sub: sub}, nil
}

// WatchHookAdded is a free log subscription operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_HookManager *HookManagerFilterer) WatchHookAdded(opts *bind.WatchOpts, sink chan<- *HookManagerHookAdded, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _HookManager.contract.WatchLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HookManagerHookAdded)
				if err := _HookManager.contract.UnpackLog(event, "HookAdded", log); err != nil {
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
func (_HookManager *HookManagerFilterer) ParseHookAdded(log types.Log) (*HookManagerHookAdded, error) {
	event := new(HookManagerHookAdded)
	if err := _HookManager.contract.UnpackLog(event, "HookAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HookManagerHookRemovedIterator is returned from FilterHookRemoved and is used to iterate over the raw logs and unpacked data for HookRemoved events raised by the HookManager contract.
type HookManagerHookRemovedIterator struct {
	Event *HookManagerHookRemoved // Event containing the contract specifics and raw log

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
func (it *HookManagerHookRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HookManagerHookRemoved)
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
		it.Event = new(HookManagerHookRemoved)
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
func (it *HookManagerHookRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HookManagerHookRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HookManagerHookRemoved represents a HookRemoved event raised by the HookManager contract.
type HookManagerHookRemoved struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookRemoved is a free log retrieval operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_HookManager *HookManagerFilterer) FilterHookRemoved(opts *bind.FilterOpts, hook []common.Address) (*HookManagerHookRemovedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _HookManager.contract.FilterLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return &HookManagerHookRemovedIterator{contract: _HookManager.contract, event: "HookRemoved", logs: logs, sub: sub}, nil
}

// WatchHookRemoved is a free log subscription operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_HookManager *HookManagerFilterer) WatchHookRemoved(opts *bind.WatchOpts, sink chan<- *HookManagerHookRemoved, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _HookManager.contract.WatchLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HookManagerHookRemoved)
				if err := _HookManager.contract.UnpackLog(event, "HookRemoved", log); err != nil {
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
func (_HookManager *HookManagerFilterer) ParseHookRemoved(log types.Log) (*HookManagerHookRemoved, error) {
	event := new(HookManagerHookRemoved)
	if err := _HookManager.contract.UnpackLog(event, "HookRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
