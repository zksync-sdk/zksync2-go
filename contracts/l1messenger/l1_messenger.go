// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1messenger

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

// L2ToL1Log is an auto generated low-level Go binding around an user-defined struct.
type L2ToL1Log struct {
	L2ShardId       uint8
	IsService       bool
	TxNumberInBlock uint16
	Sender          common.Address
	Key             [32]byte
	Value           [32]byte
}

// IL1MessengerMetaData contains all meta data concerning the IL1Messenger contract.
var IL1MessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"BytecodeL1PublicationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"L1MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"l2ShardId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isService\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"txNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structL2ToL1Log\",\"name\":\"_l2log\",\"type\":\"tuple\"}],\"name\":\"L2ToL1LogSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"requestBytecodeL1Publication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isService\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"sendL2ToL1Log\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"logIdInMerkleTree\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendToL1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL1MessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1MessengerMetaData.ABI instead.
var IL1MessengerABI = IL1MessengerMetaData.ABI

// IL1Messenger is an auto generated Go binding around an Ethereum contract.
type IL1Messenger struct {
	IL1MessengerCaller     // Read-only binding to the contract
	IL1MessengerTransactor // Write-only binding to the contract
	IL1MessengerFilterer   // Log filterer for contract events
}

// IL1MessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1MessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1MessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1MessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1MessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1MessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1MessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1MessengerSession struct {
	Contract     *IL1Messenger     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1MessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1MessengerCallerSession struct {
	Contract *IL1MessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IL1MessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1MessengerTransactorSession struct {
	Contract     *IL1MessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IL1MessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1MessengerRaw struct {
	Contract *IL1Messenger // Generic contract binding to access the raw methods on
}

// IL1MessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1MessengerCallerRaw struct {
	Contract *IL1MessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IL1MessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1MessengerTransactorRaw struct {
	Contract *IL1MessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1Messenger creates a new instance of IL1Messenger, bound to a specific deployed contract.
func NewIL1Messenger(address common.Address, backend bind.ContractBackend) (*IL1Messenger, error) {
	contract, err := bindIL1Messenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1Messenger{IL1MessengerCaller: IL1MessengerCaller{contract: contract}, IL1MessengerTransactor: IL1MessengerTransactor{contract: contract}, IL1MessengerFilterer: IL1MessengerFilterer{contract: contract}}, nil
}

// NewIL1MessengerCaller creates a new read-only instance of IL1Messenger, bound to a specific deployed contract.
func NewIL1MessengerCaller(address common.Address, caller bind.ContractCaller) (*IL1MessengerCaller, error) {
	contract, err := bindIL1Messenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1MessengerCaller{contract: contract}, nil
}

// NewIL1MessengerTransactor creates a new write-only instance of IL1Messenger, bound to a specific deployed contract.
func NewIL1MessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1MessengerTransactor, error) {
	contract, err := bindIL1Messenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1MessengerTransactor{contract: contract}, nil
}

// NewIL1MessengerFilterer creates a new log filterer instance of IL1Messenger, bound to a specific deployed contract.
func NewIL1MessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1MessengerFilterer, error) {
	contract, err := bindIL1Messenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1MessengerFilterer{contract: contract}, nil
}

// bindIL1Messenger binds a generic wrapper to an already deployed contract.
func bindIL1Messenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1MessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Messenger *IL1MessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Messenger.Contract.IL1MessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Messenger *IL1MessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Messenger.Contract.IL1MessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Messenger *IL1MessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Messenger.Contract.IL1MessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Messenger *IL1MessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Messenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Messenger *IL1MessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Messenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Messenger *IL1MessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Messenger.Contract.contract.Transact(opts, method, params...)
}

// RequestBytecodeL1Publication is a paid mutator transaction binding the contract method 0x39b34c6e.
//
// Solidity: function requestBytecodeL1Publication(bytes32 _bytecodeHash) returns()
func (_IL1Messenger *IL1MessengerTransactor) RequestBytecodeL1Publication(opts *bind.TransactOpts, _bytecodeHash [32]byte) (*types.Transaction, error) {
	return _IL1Messenger.contract.Transact(opts, "requestBytecodeL1Publication", _bytecodeHash)
}

// RequestBytecodeL1Publication is a paid mutator transaction binding the contract method 0x39b34c6e.
//
// Solidity: function requestBytecodeL1Publication(bytes32 _bytecodeHash) returns()
func (_IL1Messenger *IL1MessengerSession) RequestBytecodeL1Publication(_bytecodeHash [32]byte) (*types.Transaction, error) {
	return _IL1Messenger.Contract.RequestBytecodeL1Publication(&_IL1Messenger.TransactOpts, _bytecodeHash)
}

// RequestBytecodeL1Publication is a paid mutator transaction binding the contract method 0x39b34c6e.
//
// Solidity: function requestBytecodeL1Publication(bytes32 _bytecodeHash) returns()
func (_IL1Messenger *IL1MessengerTransactorSession) RequestBytecodeL1Publication(_bytecodeHash [32]byte) (*types.Transaction, error) {
	return _IL1Messenger.Contract.RequestBytecodeL1Publication(&_IL1Messenger.TransactOpts, _bytecodeHash)
}

// SendL2ToL1Log is a paid mutator transaction binding the contract method 0x56079ac8.
//
// Solidity: function sendL2ToL1Log(bool _isService, bytes32 _key, bytes32 _value) returns(uint256 logIdInMerkleTree)
func (_IL1Messenger *IL1MessengerTransactor) SendL2ToL1Log(opts *bind.TransactOpts, _isService bool, _key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _IL1Messenger.contract.Transact(opts, "sendL2ToL1Log", _isService, _key, _value)
}

// SendL2ToL1Log is a paid mutator transaction binding the contract method 0x56079ac8.
//
// Solidity: function sendL2ToL1Log(bool _isService, bytes32 _key, bytes32 _value) returns(uint256 logIdInMerkleTree)
func (_IL1Messenger *IL1MessengerSession) SendL2ToL1Log(_isService bool, _key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _IL1Messenger.Contract.SendL2ToL1Log(&_IL1Messenger.TransactOpts, _isService, _key, _value)
}

// SendL2ToL1Log is a paid mutator transaction binding the contract method 0x56079ac8.
//
// Solidity: function sendL2ToL1Log(bool _isService, bytes32 _key, bytes32 _value) returns(uint256 logIdInMerkleTree)
func (_IL1Messenger *IL1MessengerTransactorSession) SendL2ToL1Log(_isService bool, _key [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _IL1Messenger.Contract.SendL2ToL1Log(&_IL1Messenger.TransactOpts, _isService, _key, _value)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_IL1Messenger *IL1MessengerTransactor) SendToL1(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _IL1Messenger.contract.Transact(opts, "sendToL1", _message)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_IL1Messenger *IL1MessengerSession) SendToL1(_message []byte) (*types.Transaction, error) {
	return _IL1Messenger.Contract.SendToL1(&_IL1Messenger.TransactOpts, _message)
}

// SendToL1 is a paid mutator transaction binding the contract method 0x62f84b24.
//
// Solidity: function sendToL1(bytes _message) returns(bytes32)
func (_IL1Messenger *IL1MessengerTransactorSession) SendToL1(_message []byte) (*types.Transaction, error) {
	return _IL1Messenger.Contract.SendToL1(&_IL1Messenger.TransactOpts, _message)
}

// IL1MessengerBytecodeL1PublicationRequestedIterator is returned from FilterBytecodeL1PublicationRequested and is used to iterate over the raw logs and unpacked data for BytecodeL1PublicationRequested events raised by the IL1Messenger contract.
type IL1MessengerBytecodeL1PublicationRequestedIterator struct {
	Event *IL1MessengerBytecodeL1PublicationRequested // Event containing the contract specifics and raw log

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
func (it *IL1MessengerBytecodeL1PublicationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessengerBytecodeL1PublicationRequested)
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
		it.Event = new(IL1MessengerBytecodeL1PublicationRequested)
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
func (it *IL1MessengerBytecodeL1PublicationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessengerBytecodeL1PublicationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessengerBytecodeL1PublicationRequested represents a BytecodeL1PublicationRequested event raised by the IL1Messenger contract.
type IL1MessengerBytecodeL1PublicationRequested struct {
	BytecodeHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBytecodeL1PublicationRequested is a free log retrieval operation binding the contract event 0x480d3c9f727b5e5c1203d4c61fb185d37f08e6b2dc5e9bbf98591b1a7addf57c.
//
// Solidity: event BytecodeL1PublicationRequested(bytes32 _bytecodeHash)
func (_IL1Messenger *IL1MessengerFilterer) FilterBytecodeL1PublicationRequested(opts *bind.FilterOpts) (*IL1MessengerBytecodeL1PublicationRequestedIterator, error) {

	logs, sub, err := _IL1Messenger.contract.FilterLogs(opts, "BytecodeL1PublicationRequested")
	if err != nil {
		return nil, err
	}
	return &IL1MessengerBytecodeL1PublicationRequestedIterator{contract: _IL1Messenger.contract, event: "BytecodeL1PublicationRequested", logs: logs, sub: sub}, nil
}

// WatchBytecodeL1PublicationRequested is a free log subscription operation binding the contract event 0x480d3c9f727b5e5c1203d4c61fb185d37f08e6b2dc5e9bbf98591b1a7addf57c.
//
// Solidity: event BytecodeL1PublicationRequested(bytes32 _bytecodeHash)
func (_IL1Messenger *IL1MessengerFilterer) WatchBytecodeL1PublicationRequested(opts *bind.WatchOpts, sink chan<- *IL1MessengerBytecodeL1PublicationRequested) (event.Subscription, error) {

	logs, sub, err := _IL1Messenger.contract.WatchLogs(opts, "BytecodeL1PublicationRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessengerBytecodeL1PublicationRequested)
				if err := _IL1Messenger.contract.UnpackLog(event, "BytecodeL1PublicationRequested", log); err != nil {
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

// ParseBytecodeL1PublicationRequested is a log parse operation binding the contract event 0x480d3c9f727b5e5c1203d4c61fb185d37f08e6b2dc5e9bbf98591b1a7addf57c.
//
// Solidity: event BytecodeL1PublicationRequested(bytes32 _bytecodeHash)
func (_IL1Messenger *IL1MessengerFilterer) ParseBytecodeL1PublicationRequested(log types.Log) (*IL1MessengerBytecodeL1PublicationRequested, error) {
	event := new(IL1MessengerBytecodeL1PublicationRequested)
	if err := _IL1Messenger.contract.UnpackLog(event, "BytecodeL1PublicationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1MessengerL1MessageSentIterator is returned from FilterL1MessageSent and is used to iterate over the raw logs and unpacked data for L1MessageSent events raised by the IL1Messenger contract.
type IL1MessengerL1MessageSentIterator struct {
	Event *IL1MessengerL1MessageSent // Event containing the contract specifics and raw log

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
func (it *IL1MessengerL1MessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessengerL1MessageSent)
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
		it.Event = new(IL1MessengerL1MessageSent)
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
func (it *IL1MessengerL1MessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessengerL1MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessengerL1MessageSent represents a L1MessageSent event raised by the IL1Messenger contract.
type IL1MessengerL1MessageSent struct {
	Sender  common.Address
	Hash    [32]byte
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterL1MessageSent is a free log retrieval operation binding the contract event 0x3a36e47291f4201faf137fab081d92295bce2d53be2c6ca68ba82c7faa9ce241.
//
// Solidity: event L1MessageSent(address indexed _sender, bytes32 indexed _hash, bytes _message)
func (_IL1Messenger *IL1MessengerFilterer) FilterL1MessageSent(opts *bind.FilterOpts, _sender []common.Address, _hash [][32]byte) (*IL1MessengerL1MessageSentIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _IL1Messenger.contract.FilterLogs(opts, "L1MessageSent", _senderRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &IL1MessengerL1MessageSentIterator{contract: _IL1Messenger.contract, event: "L1MessageSent", logs: logs, sub: sub}, nil
}

// WatchL1MessageSent is a free log subscription operation binding the contract event 0x3a36e47291f4201faf137fab081d92295bce2d53be2c6ca68ba82c7faa9ce241.
//
// Solidity: event L1MessageSent(address indexed _sender, bytes32 indexed _hash, bytes _message)
func (_IL1Messenger *IL1MessengerFilterer) WatchL1MessageSent(opts *bind.WatchOpts, sink chan<- *IL1MessengerL1MessageSent, _sender []common.Address, _hash [][32]byte) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _IL1Messenger.contract.WatchLogs(opts, "L1MessageSent", _senderRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessengerL1MessageSent)
				if err := _IL1Messenger.contract.UnpackLog(event, "L1MessageSent", log); err != nil {
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

// ParseL1MessageSent is a log parse operation binding the contract event 0x3a36e47291f4201faf137fab081d92295bce2d53be2c6ca68ba82c7faa9ce241.
//
// Solidity: event L1MessageSent(address indexed _sender, bytes32 indexed _hash, bytes _message)
func (_IL1Messenger *IL1MessengerFilterer) ParseL1MessageSent(log types.Log) (*IL1MessengerL1MessageSent, error) {
	event := new(IL1MessengerL1MessageSent)
	if err := _IL1Messenger.contract.UnpackLog(event, "L1MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1MessengerL2ToL1LogSentIterator is returned from FilterL2ToL1LogSent and is used to iterate over the raw logs and unpacked data for L2ToL1LogSent events raised by the IL1Messenger contract.
type IL1MessengerL2ToL1LogSentIterator struct {
	Event *IL1MessengerL2ToL1LogSent // Event containing the contract specifics and raw log

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
func (it *IL1MessengerL2ToL1LogSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessengerL2ToL1LogSent)
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
		it.Event = new(IL1MessengerL2ToL1LogSent)
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
func (it *IL1MessengerL2ToL1LogSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessengerL2ToL1LogSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessengerL2ToL1LogSent represents a L2ToL1LogSent event raised by the IL1Messenger contract.
type IL1MessengerL2ToL1LogSent struct {
	L2log L2ToL1Log
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterL2ToL1LogSent is a free log retrieval operation binding the contract event 0x27fe8c0b49f49507b9d4fe5968c9f49edfe5c9df277d433a07a0717ede97638d.
//
// Solidity: event L2ToL1LogSent((uint8,bool,uint16,address,bytes32,bytes32) _l2log)
func (_IL1Messenger *IL1MessengerFilterer) FilterL2ToL1LogSent(opts *bind.FilterOpts) (*IL1MessengerL2ToL1LogSentIterator, error) {

	logs, sub, err := _IL1Messenger.contract.FilterLogs(opts, "L2ToL1LogSent")
	if err != nil {
		return nil, err
	}
	return &IL1MessengerL2ToL1LogSentIterator{contract: _IL1Messenger.contract, event: "L2ToL1LogSent", logs: logs, sub: sub}, nil
}

// WatchL2ToL1LogSent is a free log subscription operation binding the contract event 0x27fe8c0b49f49507b9d4fe5968c9f49edfe5c9df277d433a07a0717ede97638d.
//
// Solidity: event L2ToL1LogSent((uint8,bool,uint16,address,bytes32,bytes32) _l2log)
func (_IL1Messenger *IL1MessengerFilterer) WatchL2ToL1LogSent(opts *bind.WatchOpts, sink chan<- *IL1MessengerL2ToL1LogSent) (event.Subscription, error) {

	logs, sub, err := _IL1Messenger.contract.WatchLogs(opts, "L2ToL1LogSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessengerL2ToL1LogSent)
				if err := _IL1Messenger.contract.UnpackLog(event, "L2ToL1LogSent", log); err != nil {
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

// ParseL2ToL1LogSent is a log parse operation binding the contract event 0x27fe8c0b49f49507b9d4fe5968c9f49edfe5c9df277d433a07a0717ede97638d.
//
// Solidity: event L2ToL1LogSent((uint8,bool,uint16,address,bytes32,bytes32) _l2log)
func (_IL1Messenger *IL1MessengerFilterer) ParseL2ToL1LogSent(log types.Log) (*IL1MessengerL2ToL1LogSent, error) {
	event := new(IL1MessengerL2ToL1LogSent)
	if err := _IL1Messenger.contract.UnpackLog(event, "L2ToL1LogSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
