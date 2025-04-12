// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownermanager

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

// OwnerManagerMetaData contains all meta data concerning the OwnerManager contract.
var OwnerManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notSelf\",\"type\":\"address\"}],\"name\":\"NOT_FROM_SELF\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OWNER_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OWNER_NOT_FOUND\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isK1Owner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listK1Owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"k1OwnerList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnerManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnerManagerMetaData.ABI instead.
var OwnerManagerABI = OwnerManagerMetaData.ABI

// OwnerManager is an auto generated Go binding around an Ethereum contract.
type OwnerManager struct {
	OwnerManagerCaller     // Read-only binding to the contract
	OwnerManagerTransactor // Write-only binding to the contract
	OwnerManagerFilterer   // Log filterer for contract events
}

// OwnerManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerManagerSession struct {
	Contract     *OwnerManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerManagerCallerSession struct {
	Contract *OwnerManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OwnerManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerManagerTransactorSession struct {
	Contract     *OwnerManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OwnerManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerManagerRaw struct {
	Contract *OwnerManager // Generic contract binding to access the raw methods on
}

// OwnerManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerManagerCallerRaw struct {
	Contract *OwnerManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerManagerTransactorRaw struct {
	Contract *OwnerManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnerManager creates a new instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManager(address common.Address, backend bind.ContractBackend) (*OwnerManager, error) {
	contract, err := bindOwnerManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnerManager{OwnerManagerCaller: OwnerManagerCaller{contract: contract}, OwnerManagerTransactor: OwnerManagerTransactor{contract: contract}, OwnerManagerFilterer: OwnerManagerFilterer{contract: contract}}, nil
}

// NewOwnerManagerCaller creates a new read-only instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManagerCaller(address common.Address, caller bind.ContractCaller) (*OwnerManagerCaller, error) {
	contract, err := bindOwnerManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerCaller{contract: contract}, nil
}

// NewOwnerManagerTransactor creates a new write-only instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerManagerTransactor, error) {
	contract, err := bindOwnerManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerTransactor{contract: contract}, nil
}

// NewOwnerManagerFilterer creates a new log filterer instance of OwnerManager, bound to a specific deployed contract.
func NewOwnerManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerManagerFilterer, error) {
	contract, err := bindOwnerManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerFilterer{contract: contract}, nil
}

// bindOwnerManager binds a generic wrapper to an already deployed contract.
func bindOwnerManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnerManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerManager *OwnerManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerManager.Contract.OwnerManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerManager *OwnerManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerManager.Contract.OwnerManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerManager *OwnerManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerManager.Contract.OwnerManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerManager *OwnerManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerManager *OwnerManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerManager *OwnerManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerManager.Contract.contract.Transact(opts, method, params...)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_OwnerManager *OwnerManagerCaller) IsK1Owner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _OwnerManager.contract.Call(opts, &out, "isK1Owner", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_OwnerManager *OwnerManagerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _OwnerManager.Contract.IsK1Owner(&_OwnerManager.CallOpts, addr)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_OwnerManager *OwnerManagerCallerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _OwnerManager.Contract.IsK1Owner(&_OwnerManager.CallOpts, addr)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_OwnerManager *OwnerManagerCaller) ListK1Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OwnerManager.contract.Call(opts, &out, "listK1Owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_OwnerManager *OwnerManagerSession) ListK1Owners() ([]common.Address, error) {
	return _OwnerManager.Contract.ListK1Owners(&_OwnerManager.CallOpts)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_OwnerManager *OwnerManagerCallerSession) ListK1Owners() ([]common.Address, error) {
	return _OwnerManager.Contract.ListK1Owners(&_OwnerManager.CallOpts)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_OwnerManager *OwnerManagerTransactor) AddK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _OwnerManager.contract.Transact(opts, "addK1Owner", addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_OwnerManager *OwnerManagerSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _OwnerManager.Contract.AddK1Owner(&_OwnerManager.TransactOpts, addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_OwnerManager *OwnerManagerTransactorSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _OwnerManager.Contract.AddK1Owner(&_OwnerManager.TransactOpts, addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_OwnerManager *OwnerManagerTransactor) RemoveK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _OwnerManager.contract.Transact(opts, "removeK1Owner", addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_OwnerManager *OwnerManagerSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _OwnerManager.Contract.RemoveK1Owner(&_OwnerManager.TransactOpts, addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_OwnerManager *OwnerManagerTransactorSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _OwnerManager.Contract.RemoveK1Owner(&_OwnerManager.TransactOpts, addr)
}

// OwnerManagerK1OwnerAddedIterator is returned from FilterK1OwnerAdded and is used to iterate over the raw logs and unpacked data for K1OwnerAdded events raised by the OwnerManager contract.
type OwnerManagerK1OwnerAddedIterator struct {
	Event *OwnerManagerK1OwnerAdded // Event containing the contract specifics and raw log

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
func (it *OwnerManagerK1OwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerManagerK1OwnerAdded)
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
		it.Event = new(OwnerManagerK1OwnerAdded)
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
func (it *OwnerManagerK1OwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerManagerK1OwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerManagerK1OwnerAdded represents a K1OwnerAdded event raised by the OwnerManager contract.
type OwnerManagerK1OwnerAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerAdded is a free log retrieval operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_OwnerManager *OwnerManagerFilterer) FilterK1OwnerAdded(opts *bind.FilterOpts, addr []common.Address) (*OwnerManagerK1OwnerAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OwnerManager.contract.FilterLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerK1OwnerAddedIterator{contract: _OwnerManager.contract, event: "K1OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchK1OwnerAdded is a free log subscription operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_OwnerManager *OwnerManagerFilterer) WatchK1OwnerAdded(opts *bind.WatchOpts, sink chan<- *OwnerManagerK1OwnerAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OwnerManager.contract.WatchLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerManagerK1OwnerAdded)
				if err := _OwnerManager.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
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
func (_OwnerManager *OwnerManagerFilterer) ParseK1OwnerAdded(log types.Log) (*OwnerManagerK1OwnerAdded, error) {
	event := new(OwnerManagerK1OwnerAdded)
	if err := _OwnerManager.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerManagerK1OwnerRemovedIterator is returned from FilterK1OwnerRemoved and is used to iterate over the raw logs and unpacked data for K1OwnerRemoved events raised by the OwnerManager contract.
type OwnerManagerK1OwnerRemovedIterator struct {
	Event *OwnerManagerK1OwnerRemoved // Event containing the contract specifics and raw log

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
func (it *OwnerManagerK1OwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerManagerK1OwnerRemoved)
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
		it.Event = new(OwnerManagerK1OwnerRemoved)
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
func (it *OwnerManagerK1OwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerManagerK1OwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerManagerK1OwnerRemoved represents a K1OwnerRemoved event raised by the OwnerManager contract.
type OwnerManagerK1OwnerRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerRemoved is a free log retrieval operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_OwnerManager *OwnerManagerFilterer) FilterK1OwnerRemoved(opts *bind.FilterOpts, addr []common.Address) (*OwnerManagerK1OwnerRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OwnerManager.contract.FilterLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &OwnerManagerK1OwnerRemovedIterator{contract: _OwnerManager.contract, event: "K1OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchK1OwnerRemoved is a free log subscription operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_OwnerManager *OwnerManagerFilterer) WatchK1OwnerRemoved(opts *bind.WatchOpts, sink chan<- *OwnerManagerK1OwnerRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OwnerManager.contract.WatchLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerManagerK1OwnerRemoved)
				if err := _OwnerManager.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
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
func (_OwnerManager *OwnerManagerFilterer) ParseK1OwnerRemoved(log types.Log) (*OwnerManagerK1OwnerRemoved, error) {
	event := new(OwnerManagerK1OwnerRemoved)
	if err := _OwnerManager.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
