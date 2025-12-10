// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iownermanager

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

// IOwnerManagerMetaData contains all meta data concerning the IOwnerManager contract.
var IOwnerManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isK1Owner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listK1Owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"k1OwnerList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IOwnerManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IOwnerManagerMetaData.ABI instead.
var IOwnerManagerABI = IOwnerManagerMetaData.ABI

// IOwnerManager is an auto generated Go binding around an Ethereum contract.
type IOwnerManager struct {
	IOwnerManagerCaller     // Read-only binding to the contract
	IOwnerManagerTransactor // Write-only binding to the contract
	IOwnerManagerFilterer   // Log filterer for contract events
}

// IOwnerManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnerManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnerManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnerManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnerManagerSession struct {
	Contract     *IOwnerManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnerManagerCallerSession struct {
	Contract *IOwnerManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IOwnerManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnerManagerTransactorSession struct {
	Contract     *IOwnerManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IOwnerManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnerManagerRaw struct {
	Contract *IOwnerManager // Generic contract binding to access the raw methods on
}

// IOwnerManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnerManagerCallerRaw struct {
	Contract *IOwnerManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IOwnerManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnerManagerTransactorRaw struct {
	Contract *IOwnerManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwnerManager creates a new instance of IOwnerManager, bound to a specific deployed contract.
func NewIOwnerManager(address common.Address, backend bind.ContractBackend) (*IOwnerManager, error) {
	contract, err := bindIOwnerManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwnerManager{IOwnerManagerCaller: IOwnerManagerCaller{contract: contract}, IOwnerManagerTransactor: IOwnerManagerTransactor{contract: contract}, IOwnerManagerFilterer: IOwnerManagerFilterer{contract: contract}}, nil
}

// NewIOwnerManagerCaller creates a new read-only instance of IOwnerManager, bound to a specific deployed contract.
func NewIOwnerManagerCaller(address common.Address, caller bind.ContractCaller) (*IOwnerManagerCaller, error) {
	contract, err := bindIOwnerManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerManagerCaller{contract: contract}, nil
}

// NewIOwnerManagerTransactor creates a new write-only instance of IOwnerManager, bound to a specific deployed contract.
func NewIOwnerManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IOwnerManagerTransactor, error) {
	contract, err := bindIOwnerManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerManagerTransactor{contract: contract}, nil
}

// NewIOwnerManagerFilterer creates a new log filterer instance of IOwnerManager, bound to a specific deployed contract.
func NewIOwnerManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IOwnerManagerFilterer, error) {
	contract, err := bindIOwnerManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnerManagerFilterer{contract: contract}, nil
}

// bindIOwnerManager binds a generic wrapper to an already deployed contract.
func bindIOwnerManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOwnerManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnerManager *IOwnerManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwnerManager.Contract.IOwnerManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnerManager *IOwnerManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnerManager.Contract.IOwnerManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnerManager *IOwnerManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnerManager.Contract.IOwnerManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnerManager *IOwnerManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwnerManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnerManager *IOwnerManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnerManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnerManager *IOwnerManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnerManager.Contract.contract.Transact(opts, method, params...)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_IOwnerManager *IOwnerManagerCaller) IsK1Owner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _IOwnerManager.contract.Call(opts, &out, "isK1Owner", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_IOwnerManager *IOwnerManagerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _IOwnerManager.Contract.IsK1Owner(&_IOwnerManager.CallOpts, addr)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_IOwnerManager *IOwnerManagerCallerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _IOwnerManager.Contract.IsK1Owner(&_IOwnerManager.CallOpts, addr)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_IOwnerManager *IOwnerManagerCaller) ListK1Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IOwnerManager.contract.Call(opts, &out, "listK1Owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_IOwnerManager *IOwnerManagerSession) ListK1Owners() ([]common.Address, error) {
	return _IOwnerManager.Contract.ListK1Owners(&_IOwnerManager.CallOpts)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_IOwnerManager *IOwnerManagerCallerSession) ListK1Owners() ([]common.Address, error) {
	return _IOwnerManager.Contract.ListK1Owners(&_IOwnerManager.CallOpts)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_IOwnerManager *IOwnerManagerTransactor) AddK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _IOwnerManager.contract.Transact(opts, "addK1Owner", addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_IOwnerManager *IOwnerManagerSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _IOwnerManager.Contract.AddK1Owner(&_IOwnerManager.TransactOpts, addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_IOwnerManager *IOwnerManagerTransactorSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _IOwnerManager.Contract.AddK1Owner(&_IOwnerManager.TransactOpts, addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_IOwnerManager *IOwnerManagerTransactor) RemoveK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _IOwnerManager.contract.Transact(opts, "removeK1Owner", addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_IOwnerManager *IOwnerManagerSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _IOwnerManager.Contract.RemoveK1Owner(&_IOwnerManager.TransactOpts, addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_IOwnerManager *IOwnerManagerTransactorSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _IOwnerManager.Contract.RemoveK1Owner(&_IOwnerManager.TransactOpts, addr)
}

// IOwnerManagerK1OwnerAddedIterator is returned from FilterK1OwnerAdded and is used to iterate over the raw logs and unpacked data for K1OwnerAdded events raised by the IOwnerManager contract.
type IOwnerManagerK1OwnerAddedIterator struct {
	Event *IOwnerManagerK1OwnerAdded // Event containing the contract specifics and raw log

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
func (it *IOwnerManagerK1OwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOwnerManagerK1OwnerAdded)
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
		it.Event = new(IOwnerManagerK1OwnerAdded)
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
func (it *IOwnerManagerK1OwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOwnerManagerK1OwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOwnerManagerK1OwnerAdded represents a K1OwnerAdded event raised by the IOwnerManager contract.
type IOwnerManagerK1OwnerAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerAdded is a free log retrieval operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_IOwnerManager *IOwnerManagerFilterer) FilterK1OwnerAdded(opts *bind.FilterOpts, addr []common.Address) (*IOwnerManagerK1OwnerAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IOwnerManager.contract.FilterLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &IOwnerManagerK1OwnerAddedIterator{contract: _IOwnerManager.contract, event: "K1OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchK1OwnerAdded is a free log subscription operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_IOwnerManager *IOwnerManagerFilterer) WatchK1OwnerAdded(opts *bind.WatchOpts, sink chan<- *IOwnerManagerK1OwnerAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IOwnerManager.contract.WatchLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOwnerManagerK1OwnerAdded)
				if err := _IOwnerManager.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
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
func (_IOwnerManager *IOwnerManagerFilterer) ParseK1OwnerAdded(log types.Log) (*IOwnerManagerK1OwnerAdded, error) {
	event := new(IOwnerManagerK1OwnerAdded)
	if err := _IOwnerManager.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOwnerManagerK1OwnerRemovedIterator is returned from FilterK1OwnerRemoved and is used to iterate over the raw logs and unpacked data for K1OwnerRemoved events raised by the IOwnerManager contract.
type IOwnerManagerK1OwnerRemovedIterator struct {
	Event *IOwnerManagerK1OwnerRemoved // Event containing the contract specifics and raw log

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
func (it *IOwnerManagerK1OwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOwnerManagerK1OwnerRemoved)
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
		it.Event = new(IOwnerManagerK1OwnerRemoved)
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
func (it *IOwnerManagerK1OwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOwnerManagerK1OwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOwnerManagerK1OwnerRemoved represents a K1OwnerRemoved event raised by the IOwnerManager contract.
type IOwnerManagerK1OwnerRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerRemoved is a free log retrieval operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_IOwnerManager *IOwnerManagerFilterer) FilterK1OwnerRemoved(opts *bind.FilterOpts, addr []common.Address) (*IOwnerManagerK1OwnerRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IOwnerManager.contract.FilterLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &IOwnerManagerK1OwnerRemovedIterator{contract: _IOwnerManager.contract, event: "K1OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchK1OwnerRemoved is a free log subscription operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_IOwnerManager *IOwnerManagerFilterer) WatchK1OwnerRemoved(opts *bind.WatchOpts, sink chan<- *IOwnerManagerK1OwnerRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IOwnerManager.contract.WatchLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOwnerManagerK1OwnerRemoved)
				if err := _IOwnerManager.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
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
func (_IOwnerManager *IOwnerManagerFilterer) ParseK1OwnerRemoved(log types.Log) (*IOwnerManagerK1OwnerRemoved, error) {
	event := new(IOwnerManagerK1OwnerRemoved)
	if err := _IOwnerManager.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
