// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batchcaller

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

// Call is an auto generated low-level Go binding around an user-defined struct.
type Call struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// BatchCallerMetaData contains all meta data concerning the BatchCaller contract.
var BatchCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedValue\",\"type\":\"uint256\"}],\"name\":\"BATCH_MSG_VALUE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notSelf\",\"type\":\"address\"}],\"name\":\"NOT_FROM_SELF\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"BatchCallFailure\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCall[]\",\"name\":\"_calls\",\"type\":\"tuple[]\"}],\"name\":\"batchCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BatchCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchCallerMetaData.ABI instead.
var BatchCallerABI = BatchCallerMetaData.ABI

// BatchCaller is an auto generated Go binding around an Ethereum contract.
type BatchCaller struct {
	BatchCallerCaller     // Read-only binding to the contract
	BatchCallerTransactor // Write-only binding to the contract
	BatchCallerFilterer   // Log filterer for contract events
}

// BatchCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchCallerSession struct {
	Contract     *BatchCaller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchCallerCallerSession struct {
	Contract *BatchCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BatchCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchCallerTransactorSession struct {
	Contract     *BatchCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BatchCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchCallerRaw struct {
	Contract *BatchCaller // Generic contract binding to access the raw methods on
}

// BatchCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchCallerCallerRaw struct {
	Contract *BatchCallerCaller // Generic read-only contract binding to access the raw methods on
}

// BatchCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchCallerTransactorRaw struct {
	Contract *BatchCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchCaller creates a new instance of BatchCaller, bound to a specific deployed contract.
func NewBatchCaller(address common.Address, backend bind.ContractBackend) (*BatchCaller, error) {
	contract, err := bindBatchCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchCaller{BatchCallerCaller: BatchCallerCaller{contract: contract}, BatchCallerTransactor: BatchCallerTransactor{contract: contract}, BatchCallerFilterer: BatchCallerFilterer{contract: contract}}, nil
}

// NewBatchCallerCaller creates a new read-only instance of BatchCaller, bound to a specific deployed contract.
func NewBatchCallerCaller(address common.Address, caller bind.ContractCaller) (*BatchCallerCaller, error) {
	contract, err := bindBatchCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchCallerCaller{contract: contract}, nil
}

// NewBatchCallerTransactor creates a new write-only instance of BatchCaller, bound to a specific deployed contract.
func NewBatchCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchCallerTransactor, error) {
	contract, err := bindBatchCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchCallerTransactor{contract: contract}, nil
}

// NewBatchCallerFilterer creates a new log filterer instance of BatchCaller, bound to a specific deployed contract.
func NewBatchCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchCallerFilterer, error) {
	contract, err := bindBatchCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchCallerFilterer{contract: contract}, nil
}

// bindBatchCaller binds a generic wrapper to an already deployed contract.
func bindBatchCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchCaller *BatchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchCaller.Contract.BatchCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchCaller *BatchCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchCaller.Contract.BatchCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchCaller *BatchCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchCaller.Contract.BatchCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchCaller *BatchCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchCaller *BatchCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchCaller *BatchCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchCaller.Contract.contract.Transact(opts, method, params...)
}

// BatchCall is a paid mutator transaction binding the contract method 0x8f0273a9.
//
// Solidity: function batchCall((address,bool,uint256,bytes)[] _calls) payable returns()
func (_BatchCaller *BatchCallerTransactor) BatchCall(opts *bind.TransactOpts, _calls []Call) (*types.Transaction, error) {
	return _BatchCaller.contract.Transact(opts, "batchCall", _calls)
}

// BatchCall is a paid mutator transaction binding the contract method 0x8f0273a9.
//
// Solidity: function batchCall((address,bool,uint256,bytes)[] _calls) payable returns()
func (_BatchCaller *BatchCallerSession) BatchCall(_calls []Call) (*types.Transaction, error) {
	return _BatchCaller.Contract.BatchCall(&_BatchCaller.TransactOpts, _calls)
}

// BatchCall is a paid mutator transaction binding the contract method 0x8f0273a9.
//
// Solidity: function batchCall((address,bool,uint256,bytes)[] _calls) payable returns()
func (_BatchCaller *BatchCallerTransactorSession) BatchCall(_calls []Call) (*types.Transaction, error) {
	return _BatchCaller.Contract.BatchCall(&_BatchCaller.TransactOpts, _calls)
}

// BatchCallerBatchCallFailureIterator is returned from FilterBatchCallFailure and is used to iterate over the raw logs and unpacked data for BatchCallFailure events raised by the BatchCaller contract.
type BatchCallerBatchCallFailureIterator struct {
	Event *BatchCallerBatchCallFailure // Event containing the contract specifics and raw log

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
func (it *BatchCallerBatchCallFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchCallerBatchCallFailure)
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
		it.Event = new(BatchCallerBatchCallFailure)
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
func (it *BatchCallerBatchCallFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchCallerBatchCallFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchCallerBatchCallFailure represents a BatchCallFailure event raised by the BatchCaller contract.
type BatchCallerBatchCallFailure struct {
	Index      *big.Int
	RevertData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchCallFailure is a free log retrieval operation binding the contract event 0x860d60a3c3ed451a144846eba67081eb134233fed0aff20997e5110b942b3d0e.
//
// Solidity: event BatchCallFailure(uint256 indexed index, bytes revertData)
func (_BatchCaller *BatchCallerFilterer) FilterBatchCallFailure(opts *bind.FilterOpts, index []*big.Int) (*BatchCallerBatchCallFailureIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BatchCaller.contract.FilterLogs(opts, "BatchCallFailure", indexRule)
	if err != nil {
		return nil, err
	}
	return &BatchCallerBatchCallFailureIterator{contract: _BatchCaller.contract, event: "BatchCallFailure", logs: logs, sub: sub}, nil
}

// WatchBatchCallFailure is a free log subscription operation binding the contract event 0x860d60a3c3ed451a144846eba67081eb134233fed0aff20997e5110b942b3d0e.
//
// Solidity: event BatchCallFailure(uint256 indexed index, bytes revertData)
func (_BatchCaller *BatchCallerFilterer) WatchBatchCallFailure(opts *bind.WatchOpts, sink chan<- *BatchCallerBatchCallFailure, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BatchCaller.contract.WatchLogs(opts, "BatchCallFailure", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchCallerBatchCallFailure)
				if err := _BatchCaller.contract.UnpackLog(event, "BatchCallFailure", log); err != nil {
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

// ParseBatchCallFailure is a log parse operation binding the contract event 0x860d60a3c3ed451a144846eba67081eb134233fed0aff20997e5110b942b3d0e.
//
// Solidity: event BatchCallFailure(uint256 indexed index, bytes revertData)
func (_BatchCaller *BatchCallerFilterer) ParseBatchCallFailure(log types.Log) (*BatchCallerBatchCallFailure, error) {
	event := new(BatchCallerBatchCallFailure)
	if err := _BatchCaller.contract.UnpackLog(event, "BatchCallFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
