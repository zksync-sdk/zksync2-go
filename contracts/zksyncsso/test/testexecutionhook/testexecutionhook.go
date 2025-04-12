// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testexecutionhook

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

// Transaction is an auto generated low-level Go binding around an user-defined struct.
type Transaction struct {
	TxType                 *big.Int
	From                   *big.Int
	To                     *big.Int
	GasLimit               *big.Int
	GasPerPubdataByteLimit *big.Int
	MaxFeePerGas           *big.Int
	MaxPriorityFeePerGas   *big.Int
	Paymaster              *big.Int
	Nonce                  *big.Int
	Value                  *big.Int
	Reserved               [4]*big.Int
	Data                   []byte
	Signature              []byte
	FactoryDeps            [][32]byte
	PaymasterInput         []byte
	ReservedDynamic        []byte
}

// TestExecutionHookMetaData contains all meta data concerning the TestExecutionHook contract.
var TestExecutionHookMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ExecutionHookInstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ExecutionHookUninstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PostExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"PreExecution\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"postExecutionHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"preExecutionHook\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// TestExecutionHookABI is the input ABI used to generate the binding from.
// Deprecated: Use TestExecutionHookMetaData.ABI instead.
var TestExecutionHookABI = TestExecutionHookMetaData.ABI

// TestExecutionHook is an auto generated Go binding around an Ethereum contract.
type TestExecutionHook struct {
	TestExecutionHookCaller     // Read-only binding to the contract
	TestExecutionHookTransactor // Write-only binding to the contract
	TestExecutionHookFilterer   // Log filterer for contract events
}

// TestExecutionHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestExecutionHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecutionHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestExecutionHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecutionHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestExecutionHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestExecutionHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestExecutionHookSession struct {
	Contract     *TestExecutionHook // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestExecutionHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestExecutionHookCallerSession struct {
	Contract *TestExecutionHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TestExecutionHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestExecutionHookTransactorSession struct {
	Contract     *TestExecutionHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TestExecutionHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestExecutionHookRaw struct {
	Contract *TestExecutionHook // Generic contract binding to access the raw methods on
}

// TestExecutionHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestExecutionHookCallerRaw struct {
	Contract *TestExecutionHookCaller // Generic read-only contract binding to access the raw methods on
}

// TestExecutionHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestExecutionHookTransactorRaw struct {
	Contract *TestExecutionHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestExecutionHook creates a new instance of TestExecutionHook, bound to a specific deployed contract.
func NewTestExecutionHook(address common.Address, backend bind.ContractBackend) (*TestExecutionHook, error) {
	contract, err := bindTestExecutionHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHook{TestExecutionHookCaller: TestExecutionHookCaller{contract: contract}, TestExecutionHookTransactor: TestExecutionHookTransactor{contract: contract}, TestExecutionHookFilterer: TestExecutionHookFilterer{contract: contract}}, nil
}

// NewTestExecutionHookCaller creates a new read-only instance of TestExecutionHook, bound to a specific deployed contract.
func NewTestExecutionHookCaller(address common.Address, caller bind.ContractCaller) (*TestExecutionHookCaller, error) {
	contract, err := bindTestExecutionHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHookCaller{contract: contract}, nil
}

// NewTestExecutionHookTransactor creates a new write-only instance of TestExecutionHook, bound to a specific deployed contract.
func NewTestExecutionHookTransactor(address common.Address, transactor bind.ContractTransactor) (*TestExecutionHookTransactor, error) {
	contract, err := bindTestExecutionHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHookTransactor{contract: contract}, nil
}

// NewTestExecutionHookFilterer creates a new log filterer instance of TestExecutionHook, bound to a specific deployed contract.
func NewTestExecutionHookFilterer(address common.Address, filterer bind.ContractFilterer) (*TestExecutionHookFilterer, error) {
	contract, err := bindTestExecutionHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHookFilterer{contract: contract}, nil
}

// bindTestExecutionHook binds a generic wrapper to an already deployed contract.
func bindTestExecutionHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestExecutionHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExecutionHook *TestExecutionHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExecutionHook.Contract.TestExecutionHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExecutionHook *TestExecutionHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.TestExecutionHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExecutionHook *TestExecutionHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.TestExecutionHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestExecutionHook *TestExecutionHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestExecutionHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestExecutionHook *TestExecutionHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestExecutionHook *TestExecutionHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.contract.Transact(opts, method, params...)
}

// LastTarget is a free data retrieval call binding the contract method 0x99d48762.
//
// Solidity: function lastTarget(address ) view returns(address)
func (_TestExecutionHook *TestExecutionHookCaller) LastTarget(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TestExecutionHook.contract.Call(opts, &out, "lastTarget", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastTarget is a free data retrieval call binding the contract method 0x99d48762.
//
// Solidity: function lastTarget(address ) view returns(address)
func (_TestExecutionHook *TestExecutionHookSession) LastTarget(arg0 common.Address) (common.Address, error) {
	return _TestExecutionHook.Contract.LastTarget(&_TestExecutionHook.CallOpts, arg0)
}

// LastTarget is a free data retrieval call binding the contract method 0x99d48762.
//
// Solidity: function lastTarget(address ) view returns(address)
func (_TestExecutionHook *TestExecutionHookCallerSession) LastTarget(arg0 common.Address) (common.Address, error) {
	return _TestExecutionHook.Contract.LastTarget(&_TestExecutionHook.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestExecutionHook *TestExecutionHookCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestExecutionHook.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestExecutionHook *TestExecutionHookSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestExecutionHook.Contract.SupportsInterface(&_TestExecutionHook.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestExecutionHook *TestExecutionHookCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestExecutionHook.Contract.SupportsInterface(&_TestExecutionHook.CallOpts, interfaceId)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_TestExecutionHook *TestExecutionHookTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _TestExecutionHook.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_TestExecutionHook *TestExecutionHookSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.OnInstall(&_TestExecutionHook.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_TestExecutionHook *TestExecutionHookTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.OnInstall(&_TestExecutionHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_TestExecutionHook *TestExecutionHookTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _TestExecutionHook.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_TestExecutionHook *TestExecutionHookSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.OnUninstall(&_TestExecutionHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_TestExecutionHook *TestExecutionHookTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.OnUninstall(&_TestExecutionHook.TransactOpts, data)
}

// PostExecutionHook is a paid mutator transaction binding the contract method 0x3920e539.
//
// Solidity: function postExecutionHook(bytes context) returns()
func (_TestExecutionHook *TestExecutionHookTransactor) PostExecutionHook(opts *bind.TransactOpts, context []byte) (*types.Transaction, error) {
	return _TestExecutionHook.contract.Transact(opts, "postExecutionHook", context)
}

// PostExecutionHook is a paid mutator transaction binding the contract method 0x3920e539.
//
// Solidity: function postExecutionHook(bytes context) returns()
func (_TestExecutionHook *TestExecutionHookSession) PostExecutionHook(context []byte) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.PostExecutionHook(&_TestExecutionHook.TransactOpts, context)
}

// PostExecutionHook is a paid mutator transaction binding the contract method 0x3920e539.
//
// Solidity: function postExecutionHook(bytes context) returns()
func (_TestExecutionHook *TestExecutionHookTransactorSession) PostExecutionHook(context []byte) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.PostExecutionHook(&_TestExecutionHook.TransactOpts, context)
}

// PreExecutionHook is a paid mutator transaction binding the contract method 0x926a74a6.
//
// Solidity: function preExecutionHook((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bytes context)
func (_TestExecutionHook *TestExecutionHookTransactor) PreExecutionHook(opts *bind.TransactOpts, transaction Transaction) (*types.Transaction, error) {
	return _TestExecutionHook.contract.Transact(opts, "preExecutionHook", transaction)
}

// PreExecutionHook is a paid mutator transaction binding the contract method 0x926a74a6.
//
// Solidity: function preExecutionHook((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bytes context)
func (_TestExecutionHook *TestExecutionHookSession) PreExecutionHook(transaction Transaction) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.PreExecutionHook(&_TestExecutionHook.TransactOpts, transaction)
}

// PreExecutionHook is a paid mutator transaction binding the contract method 0x926a74a6.
//
// Solidity: function preExecutionHook((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bytes context)
func (_TestExecutionHook *TestExecutionHookTransactorSession) PreExecutionHook(transaction Transaction) (*types.Transaction, error) {
	return _TestExecutionHook.Contract.PreExecutionHook(&_TestExecutionHook.TransactOpts, transaction)
}

// TestExecutionHookExecutionHookInstalledIterator is returned from FilterExecutionHookInstalled and is used to iterate over the raw logs and unpacked data for ExecutionHookInstalled events raised by the TestExecutionHook contract.
type TestExecutionHookExecutionHookInstalledIterator struct {
	Event *TestExecutionHookExecutionHookInstalled // Event containing the contract specifics and raw log

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
func (it *TestExecutionHookExecutionHookInstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecutionHookExecutionHookInstalled)
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
		it.Event = new(TestExecutionHookExecutionHookInstalled)
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
func (it *TestExecutionHookExecutionHookInstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecutionHookExecutionHookInstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecutionHookExecutionHookInstalled represents a ExecutionHookInstalled event raised by the TestExecutionHook contract.
type TestExecutionHookExecutionHookInstalled struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionHookInstalled is a free log retrieval operation binding the contract event 0x1a4792d8240e108e601cf588feb48e55008ad1ce1e56e8d8f7e737496ad36d5a.
//
// Solidity: event ExecutionHookInstalled(address indexed account)
func (_TestExecutionHook *TestExecutionHookFilterer) FilterExecutionHookInstalled(opts *bind.FilterOpts, account []common.Address) (*TestExecutionHookExecutionHookInstalledIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestExecutionHook.contract.FilterLogs(opts, "ExecutionHookInstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHookExecutionHookInstalledIterator{contract: _TestExecutionHook.contract, event: "ExecutionHookInstalled", logs: logs, sub: sub}, nil
}

// WatchExecutionHookInstalled is a free log subscription operation binding the contract event 0x1a4792d8240e108e601cf588feb48e55008ad1ce1e56e8d8f7e737496ad36d5a.
//
// Solidity: event ExecutionHookInstalled(address indexed account)
func (_TestExecutionHook *TestExecutionHookFilterer) WatchExecutionHookInstalled(opts *bind.WatchOpts, sink chan<- *TestExecutionHookExecutionHookInstalled, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestExecutionHook.contract.WatchLogs(opts, "ExecutionHookInstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecutionHookExecutionHookInstalled)
				if err := _TestExecutionHook.contract.UnpackLog(event, "ExecutionHookInstalled", log); err != nil {
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

// ParseExecutionHookInstalled is a log parse operation binding the contract event 0x1a4792d8240e108e601cf588feb48e55008ad1ce1e56e8d8f7e737496ad36d5a.
//
// Solidity: event ExecutionHookInstalled(address indexed account)
func (_TestExecutionHook *TestExecutionHookFilterer) ParseExecutionHookInstalled(log types.Log) (*TestExecutionHookExecutionHookInstalled, error) {
	event := new(TestExecutionHookExecutionHookInstalled)
	if err := _TestExecutionHook.contract.UnpackLog(event, "ExecutionHookInstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExecutionHookExecutionHookUninstalledIterator is returned from FilterExecutionHookUninstalled and is used to iterate over the raw logs and unpacked data for ExecutionHookUninstalled events raised by the TestExecutionHook contract.
type TestExecutionHookExecutionHookUninstalledIterator struct {
	Event *TestExecutionHookExecutionHookUninstalled // Event containing the contract specifics and raw log

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
func (it *TestExecutionHookExecutionHookUninstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecutionHookExecutionHookUninstalled)
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
		it.Event = new(TestExecutionHookExecutionHookUninstalled)
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
func (it *TestExecutionHookExecutionHookUninstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecutionHookExecutionHookUninstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecutionHookExecutionHookUninstalled represents a ExecutionHookUninstalled event raised by the TestExecutionHook contract.
type TestExecutionHookExecutionHookUninstalled struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionHookUninstalled is a free log retrieval operation binding the contract event 0xf726cef83066d912247b1fe30d29f3cf8ad61df09f96937b3ccee866fbb787f1.
//
// Solidity: event ExecutionHookUninstalled(address indexed account)
func (_TestExecutionHook *TestExecutionHookFilterer) FilterExecutionHookUninstalled(opts *bind.FilterOpts, account []common.Address) (*TestExecutionHookExecutionHookUninstalledIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestExecutionHook.contract.FilterLogs(opts, "ExecutionHookUninstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHookExecutionHookUninstalledIterator{contract: _TestExecutionHook.contract, event: "ExecutionHookUninstalled", logs: logs, sub: sub}, nil
}

// WatchExecutionHookUninstalled is a free log subscription operation binding the contract event 0xf726cef83066d912247b1fe30d29f3cf8ad61df09f96937b3ccee866fbb787f1.
//
// Solidity: event ExecutionHookUninstalled(address indexed account)
func (_TestExecutionHook *TestExecutionHookFilterer) WatchExecutionHookUninstalled(opts *bind.WatchOpts, sink chan<- *TestExecutionHookExecutionHookUninstalled, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestExecutionHook.contract.WatchLogs(opts, "ExecutionHookUninstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecutionHookExecutionHookUninstalled)
				if err := _TestExecutionHook.contract.UnpackLog(event, "ExecutionHookUninstalled", log); err != nil {
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

// ParseExecutionHookUninstalled is a log parse operation binding the contract event 0xf726cef83066d912247b1fe30d29f3cf8ad61df09f96937b3ccee866fbb787f1.
//
// Solidity: event ExecutionHookUninstalled(address indexed account)
func (_TestExecutionHook *TestExecutionHookFilterer) ParseExecutionHookUninstalled(log types.Log) (*TestExecutionHookExecutionHookUninstalled, error) {
	event := new(TestExecutionHookExecutionHookUninstalled)
	if err := _TestExecutionHook.contract.UnpackLog(event, "ExecutionHookUninstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExecutionHookPostExecutionIterator is returned from FilterPostExecution and is used to iterate over the raw logs and unpacked data for PostExecution events raised by the TestExecutionHook contract.
type TestExecutionHookPostExecutionIterator struct {
	Event *TestExecutionHookPostExecution // Event containing the contract specifics and raw log

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
func (it *TestExecutionHookPostExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecutionHookPostExecution)
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
		it.Event = new(TestExecutionHookPostExecution)
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
func (it *TestExecutionHookPostExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecutionHookPostExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecutionHookPostExecution represents a PostExecution event raised by the TestExecutionHook contract.
type TestExecutionHookPostExecution struct {
	Account common.Address
	Target  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPostExecution is a free log retrieval operation binding the contract event 0x61f378cfbc325e0bff59b7fea67f609d10e0e3910e06bf03b8ec822082de5de6.
//
// Solidity: event PostExecution(address indexed account, address indexed target)
func (_TestExecutionHook *TestExecutionHookFilterer) FilterPostExecution(opts *bind.FilterOpts, account []common.Address, target []common.Address) (*TestExecutionHookPostExecutionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TestExecutionHook.contract.FilterLogs(opts, "PostExecution", accountRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHookPostExecutionIterator{contract: _TestExecutionHook.contract, event: "PostExecution", logs: logs, sub: sub}, nil
}

// WatchPostExecution is a free log subscription operation binding the contract event 0x61f378cfbc325e0bff59b7fea67f609d10e0e3910e06bf03b8ec822082de5de6.
//
// Solidity: event PostExecution(address indexed account, address indexed target)
func (_TestExecutionHook *TestExecutionHookFilterer) WatchPostExecution(opts *bind.WatchOpts, sink chan<- *TestExecutionHookPostExecution, account []common.Address, target []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TestExecutionHook.contract.WatchLogs(opts, "PostExecution", accountRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecutionHookPostExecution)
				if err := _TestExecutionHook.contract.UnpackLog(event, "PostExecution", log); err != nil {
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

// ParsePostExecution is a log parse operation binding the contract event 0x61f378cfbc325e0bff59b7fea67f609d10e0e3910e06bf03b8ec822082de5de6.
//
// Solidity: event PostExecution(address indexed account, address indexed target)
func (_TestExecutionHook *TestExecutionHookFilterer) ParsePostExecution(log types.Log) (*TestExecutionHookPostExecution, error) {
	event := new(TestExecutionHookPostExecution)
	if err := _TestExecutionHook.contract.UnpackLog(event, "PostExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestExecutionHookPreExecutionIterator is returned from FilterPreExecution and is used to iterate over the raw logs and unpacked data for PreExecution events raised by the TestExecutionHook contract.
type TestExecutionHookPreExecutionIterator struct {
	Event *TestExecutionHookPreExecution // Event containing the contract specifics and raw log

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
func (it *TestExecutionHookPreExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestExecutionHookPreExecution)
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
		it.Event = new(TestExecutionHookPreExecution)
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
func (it *TestExecutionHookPreExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestExecutionHookPreExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestExecutionHookPreExecution represents a PreExecution event raised by the TestExecutionHook contract.
type TestExecutionHookPreExecution struct {
	Account common.Address
	Target  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPreExecution is a free log retrieval operation binding the contract event 0x4a145db413bc0ef36bb6eff13d165b8dc2aed778d6a3be333378c25d3a4cc835.
//
// Solidity: event PreExecution(address indexed account, address indexed target)
func (_TestExecutionHook *TestExecutionHookFilterer) FilterPreExecution(opts *bind.FilterOpts, account []common.Address, target []common.Address) (*TestExecutionHookPreExecutionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TestExecutionHook.contract.FilterLogs(opts, "PreExecution", accountRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &TestExecutionHookPreExecutionIterator{contract: _TestExecutionHook.contract, event: "PreExecution", logs: logs, sub: sub}, nil
}

// WatchPreExecution is a free log subscription operation binding the contract event 0x4a145db413bc0ef36bb6eff13d165b8dc2aed778d6a3be333378c25d3a4cc835.
//
// Solidity: event PreExecution(address indexed account, address indexed target)
func (_TestExecutionHook *TestExecutionHookFilterer) WatchPreExecution(opts *bind.WatchOpts, sink chan<- *TestExecutionHookPreExecution, account []common.Address, target []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TestExecutionHook.contract.WatchLogs(opts, "PreExecution", accountRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestExecutionHookPreExecution)
				if err := _TestExecutionHook.contract.UnpackLog(event, "PreExecution", log); err != nil {
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

// ParsePreExecution is a log parse operation binding the contract event 0x4a145db413bc0ef36bb6eff13d165b8dc2aed778d6a3be333378c25d3a4cc835.
//
// Solidity: event PreExecution(address indexed account, address indexed target)
func (_TestExecutionHook *TestExecutionHookFilterer) ParsePreExecution(log types.Log) (*TestExecutionHookPreExecution, error) {
	event := new(TestExecutionHookPreExecution)
	if err := _TestExecutionHook.contract.UnpackLog(event, "PreExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
