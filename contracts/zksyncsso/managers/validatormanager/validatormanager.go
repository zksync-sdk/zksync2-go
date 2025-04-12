// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatormanager

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

// ValidatorManagerMetaData contains all meta data concerning the ValidatorManager contract.
var ValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notSelf\",\"type\":\"address\"}],\"name\":\"NOT_FROM_SELF\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_ERC165_FAIL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_NOT_FOUND\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isModuleValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listModuleValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorManagerMetaData.ABI instead.
var ValidatorManagerABI = ValidatorManagerMetaData.ABI

// ValidatorManager is an auto generated Go binding around an Ethereum contract.
type ValidatorManager struct {
	ValidatorManagerCaller     // Read-only binding to the contract
	ValidatorManagerTransactor // Write-only binding to the contract
	ValidatorManagerFilterer   // Log filterer for contract events
}

// ValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorManagerSession struct {
	Contract     *ValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorManagerCallerSession struct {
	Contract *ValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorManagerTransactorSession struct {
	Contract     *ValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorManagerRaw struct {
	Contract *ValidatorManager // Generic contract binding to access the raw methods on
}

// ValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorManagerCallerRaw struct {
	Contract *ValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactorRaw struct {
	Contract *ValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorManager creates a new instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManager(address common.Address, backend bind.ContractBackend) (*ValidatorManager, error) {
	contract, err := bindValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorManager{ValidatorManagerCaller: ValidatorManagerCaller{contract: contract}, ValidatorManagerTransactor: ValidatorManagerTransactor{contract: contract}, ValidatorManagerFilterer: ValidatorManagerFilterer{contract: contract}}, nil
}

// NewValidatorManagerCaller creates a new read-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*ValidatorManagerCaller, error) {
	contract, err := bindValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerCaller{contract: contract}, nil
}

// NewValidatorManagerTransactor creates a new write-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorManagerTransactor, error) {
	contract, err := bindValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerTransactor{contract: contract}, nil
}

// NewValidatorManagerFilterer creates a new log filterer instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorManagerFilterer, error) {
	contract, err := bindValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerFilterer{contract: contract}, nil
}

// bindValidatorManager binds a generic wrapper to an already deployed contract.
func bindValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.ValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ValidatorManager *ValidatorManagerCaller) IsModuleValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "isModuleValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ValidatorManager *ValidatorManagerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsModuleValidator(&_ValidatorManager.CallOpts, validator)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ValidatorManager *ValidatorManagerCallerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _ValidatorManager.Contract.IsModuleValidator(&_ValidatorManager.CallOpts, validator)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ValidatorManager *ValidatorManagerCaller) ListModuleValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "listModuleValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ValidatorManager *ValidatorManagerSession) ListModuleValidators() ([]common.Address, error) {
	return _ValidatorManager.Contract.ListModuleValidators(&_ValidatorManager.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ValidatorManager *ValidatorManagerCallerSession) ListModuleValidators() ([]common.Address, error) {
	return _ValidatorManager.Contract.ListModuleValidators(&_ValidatorManager.CallOpts)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ValidatorManager *ValidatorManagerTransactor) AddModuleValidator(opts *bind.TransactOpts, validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "addModuleValidator", validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ValidatorManager *ValidatorManagerSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.AddModuleValidator(&_ValidatorManager.TransactOpts, validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.AddModuleValidator(&_ValidatorManager.TransactOpts, validator, initData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ValidatorManager *ValidatorManagerTransactor) RemoveModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "removeModuleValidator", validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ValidatorManager *ValidatorManagerSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.RemoveModuleValidator(&_ValidatorManager.TransactOpts, validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.RemoveModuleValidator(&_ValidatorManager.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ValidatorManager *ValidatorManagerTransactor) UnlinkModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "unlinkModuleValidator", validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ValidatorManager *ValidatorManagerSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.UnlinkModuleValidator(&_ValidatorManager.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.UnlinkModuleValidator(&_ValidatorManager.TransactOpts, validator, deinitData)
}

// ValidatorManagerValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ValidatorManager contract.
type ValidatorManagerValidatorAddedIterator struct {
	Event *ValidatorManagerValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorAdded)
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
		it.Event = new(ValidatorManagerValidatorAdded)
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
func (it *ValidatorManagerValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorAdded represents a ValidatorAdded event raised by the ValidatorManager contract.
type ValidatorManagerValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*ValidatorManagerValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorAddedIterator{contract: _ValidatorManager.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorAdded)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorAdded(log types.Log) (*ValidatorManagerValidatorAdded, error) {
	event := new(ValidatorManagerValidatorAdded)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the ValidatorManager contract.
type ValidatorManagerValidatorRemovedIterator struct {
	Event *ValidatorManagerValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorRemoved)
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
		it.Event = new(ValidatorManagerValidatorRemoved)
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
func (it *ValidatorManagerValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorRemoved represents a ValidatorRemoved event raised by the ValidatorManager contract.
type ValidatorManagerValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*ValidatorManagerValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorRemovedIterator{contract: _ValidatorManager.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorRemoved)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorRemoved(log types.Log) (*ValidatorManagerValidatorRemoved, error) {
	event := new(ValidatorManagerValidatorRemoved)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
