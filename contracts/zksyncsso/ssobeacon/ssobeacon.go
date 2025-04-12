// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ssobeacon

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

// SsoBeaconMetaData contains all meta data concerning the SsoBeacon contract.
var SsoBeaconMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SsoBeaconABI is the input ABI used to generate the binding from.
// Deprecated: Use SsoBeaconMetaData.ABI instead.
var SsoBeaconABI = SsoBeaconMetaData.ABI

// SsoBeacon is an auto generated Go binding around an Ethereum contract.
type SsoBeacon struct {
	SsoBeaconCaller     // Read-only binding to the contract
	SsoBeaconTransactor // Write-only binding to the contract
	SsoBeaconFilterer   // Log filterer for contract events
}

// SsoBeaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type SsoBeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoBeaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SsoBeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoBeaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SsoBeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoBeaconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SsoBeaconSession struct {
	Contract     *SsoBeacon        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsoBeaconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SsoBeaconCallerSession struct {
	Contract *SsoBeaconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SsoBeaconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SsoBeaconTransactorSession struct {
	Contract     *SsoBeaconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SsoBeaconRaw is an auto generated low-level Go binding around an Ethereum contract.
type SsoBeaconRaw struct {
	Contract *SsoBeacon // Generic contract binding to access the raw methods on
}

// SsoBeaconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SsoBeaconCallerRaw struct {
	Contract *SsoBeaconCaller // Generic read-only contract binding to access the raw methods on
}

// SsoBeaconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SsoBeaconTransactorRaw struct {
	Contract *SsoBeaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSsoBeacon creates a new instance of SsoBeacon, bound to a specific deployed contract.
func NewSsoBeacon(address common.Address, backend bind.ContractBackend) (*SsoBeacon, error) {
	contract, err := bindSsoBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SsoBeacon{SsoBeaconCaller: SsoBeaconCaller{contract: contract}, SsoBeaconTransactor: SsoBeaconTransactor{contract: contract}, SsoBeaconFilterer: SsoBeaconFilterer{contract: contract}}, nil
}

// NewSsoBeaconCaller creates a new read-only instance of SsoBeacon, bound to a specific deployed contract.
func NewSsoBeaconCaller(address common.Address, caller bind.ContractCaller) (*SsoBeaconCaller, error) {
	contract, err := bindSsoBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SsoBeaconCaller{contract: contract}, nil
}

// NewSsoBeaconTransactor creates a new write-only instance of SsoBeacon, bound to a specific deployed contract.
func NewSsoBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*SsoBeaconTransactor, error) {
	contract, err := bindSsoBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SsoBeaconTransactor{contract: contract}, nil
}

// NewSsoBeaconFilterer creates a new log filterer instance of SsoBeacon, bound to a specific deployed contract.
func NewSsoBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*SsoBeaconFilterer, error) {
	contract, err := bindSsoBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SsoBeaconFilterer{contract: contract}, nil
}

// bindSsoBeacon binds a generic wrapper to an already deployed contract.
func bindSsoBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SsoBeaconMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsoBeacon *SsoBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsoBeacon.Contract.SsoBeaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsoBeacon *SsoBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoBeacon.Contract.SsoBeaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsoBeacon *SsoBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsoBeacon.Contract.SsoBeaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsoBeacon *SsoBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsoBeacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsoBeacon *SsoBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoBeacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsoBeacon *SsoBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsoBeacon.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SsoBeacon *SsoBeaconCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SsoBeacon.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SsoBeacon *SsoBeaconSession) Implementation() (common.Address, error) {
	return _SsoBeacon.Contract.Implementation(&_SsoBeacon.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SsoBeacon *SsoBeaconCallerSession) Implementation() (common.Address, error) {
	return _SsoBeacon.Contract.Implementation(&_SsoBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsoBeacon *SsoBeaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SsoBeacon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsoBeacon *SsoBeaconSession) Owner() (common.Address, error) {
	return _SsoBeacon.Contract.Owner(&_SsoBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SsoBeacon *SsoBeaconCallerSession) Owner() (common.Address, error) {
	return _SsoBeacon.Contract.Owner(&_SsoBeacon.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsoBeacon *SsoBeaconTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoBeacon.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsoBeacon *SsoBeaconSession) RenounceOwnership() (*types.Transaction, error) {
	return _SsoBeacon.Contract.RenounceOwnership(&_SsoBeacon.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SsoBeacon *SsoBeaconTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SsoBeacon.Contract.RenounceOwnership(&_SsoBeacon.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsoBeacon *SsoBeaconTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SsoBeacon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsoBeacon *SsoBeaconSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SsoBeacon.Contract.TransferOwnership(&_SsoBeacon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SsoBeacon *SsoBeaconTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SsoBeacon.Contract.TransferOwnership(&_SsoBeacon.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsoBeacon *SsoBeaconTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SsoBeacon.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsoBeacon *SsoBeaconSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SsoBeacon.Contract.UpgradeTo(&_SsoBeacon.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SsoBeacon *SsoBeaconTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SsoBeacon.Contract.UpgradeTo(&_SsoBeacon.TransactOpts, newImplementation)
}

// SsoBeaconOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SsoBeacon contract.
type SsoBeaconOwnershipTransferredIterator struct {
	Event *SsoBeaconOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SsoBeaconOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoBeaconOwnershipTransferred)
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
		it.Event = new(SsoBeaconOwnershipTransferred)
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
func (it *SsoBeaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoBeaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoBeaconOwnershipTransferred represents a OwnershipTransferred event raised by the SsoBeacon contract.
type SsoBeaconOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsoBeacon *SsoBeaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsoBeaconOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsoBeacon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsoBeaconOwnershipTransferredIterator{contract: _SsoBeacon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsoBeacon *SsoBeaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SsoBeaconOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SsoBeacon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoBeaconOwnershipTransferred)
				if err := _SsoBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SsoBeacon *SsoBeaconFilterer) ParseOwnershipTransferred(log types.Log) (*SsoBeaconOwnershipTransferred, error) {
	event := new(SsoBeaconOwnershipTransferred)
	if err := _SsoBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoBeaconUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SsoBeacon contract.
type SsoBeaconUpgradedIterator struct {
	Event *SsoBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SsoBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoBeaconUpgraded)
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
		it.Event = new(SsoBeaconUpgraded)
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
func (it *SsoBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoBeaconUpgraded represents a Upgraded event raised by the SsoBeacon contract.
type SsoBeaconUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsoBeacon *SsoBeaconFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SsoBeaconUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SsoBeacon.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SsoBeaconUpgradedIterator{contract: _SsoBeacon.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsoBeacon *SsoBeaconFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SsoBeaconUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SsoBeacon.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoBeaconUpgraded)
				if err := _SsoBeacon.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SsoBeacon *SsoBeaconFilterer) ParseUpgraded(log types.Log) (*SsoBeaconUpgraded, error) {
	event := new(SsoBeaconUpgraded)
	if err := _SsoBeacon.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
