// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testvalidationhook

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

// TestValidationHookMetaData contains all meta data concerning the TestValidationHook contract.
var TestValidationHookMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidationHookInstalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"ValidationHookTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidationHookUninstalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validationHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TestValidationHookABI is the input ABI used to generate the binding from.
// Deprecated: Use TestValidationHookMetaData.ABI instead.
var TestValidationHookABI = TestValidationHookMetaData.ABI

// TestValidationHook is an auto generated Go binding around an Ethereum contract.
type TestValidationHook struct {
	TestValidationHookCaller     // Read-only binding to the contract
	TestValidationHookTransactor // Write-only binding to the contract
	TestValidationHookFilterer   // Log filterer for contract events
}

// TestValidationHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestValidationHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestValidationHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestValidationHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestValidationHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestValidationHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestValidationHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestValidationHookSession struct {
	Contract     *TestValidationHook // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestValidationHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestValidationHookCallerSession struct {
	Contract *TestValidationHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TestValidationHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestValidationHookTransactorSession struct {
	Contract     *TestValidationHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TestValidationHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestValidationHookRaw struct {
	Contract *TestValidationHook // Generic contract binding to access the raw methods on
}

// TestValidationHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestValidationHookCallerRaw struct {
	Contract *TestValidationHookCaller // Generic read-only contract binding to access the raw methods on
}

// TestValidationHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestValidationHookTransactorRaw struct {
	Contract *TestValidationHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestValidationHook creates a new instance of TestValidationHook, bound to a specific deployed contract.
func NewTestValidationHook(address common.Address, backend bind.ContractBackend) (*TestValidationHook, error) {
	contract, err := bindTestValidationHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestValidationHook{TestValidationHookCaller: TestValidationHookCaller{contract: contract}, TestValidationHookTransactor: TestValidationHookTransactor{contract: contract}, TestValidationHookFilterer: TestValidationHookFilterer{contract: contract}}, nil
}

// NewTestValidationHookCaller creates a new read-only instance of TestValidationHook, bound to a specific deployed contract.
func NewTestValidationHookCaller(address common.Address, caller bind.ContractCaller) (*TestValidationHookCaller, error) {
	contract, err := bindTestValidationHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestValidationHookCaller{contract: contract}, nil
}

// NewTestValidationHookTransactor creates a new write-only instance of TestValidationHook, bound to a specific deployed contract.
func NewTestValidationHookTransactor(address common.Address, transactor bind.ContractTransactor) (*TestValidationHookTransactor, error) {
	contract, err := bindTestValidationHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestValidationHookTransactor{contract: contract}, nil
}

// NewTestValidationHookFilterer creates a new log filterer instance of TestValidationHook, bound to a specific deployed contract.
func NewTestValidationHookFilterer(address common.Address, filterer bind.ContractFilterer) (*TestValidationHookFilterer, error) {
	contract, err := bindTestValidationHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestValidationHookFilterer{contract: contract}, nil
}

// bindTestValidationHook binds a generic wrapper to an already deployed contract.
func bindTestValidationHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestValidationHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestValidationHook *TestValidationHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestValidationHook.Contract.TestValidationHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestValidationHook *TestValidationHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestValidationHook.Contract.TestValidationHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestValidationHook *TestValidationHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestValidationHook.Contract.TestValidationHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestValidationHook *TestValidationHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestValidationHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestValidationHook *TestValidationHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestValidationHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestValidationHook *TestValidationHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestValidationHook.Contract.contract.Transact(opts, method, params...)
}

// LastTarget is a free data retrieval call binding the contract method 0x99d48762.
//
// Solidity: function lastTarget(address ) view returns(address)
func (_TestValidationHook *TestValidationHookCaller) LastTarget(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TestValidationHook.contract.Call(opts, &out, "lastTarget", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastTarget is a free data retrieval call binding the contract method 0x99d48762.
//
// Solidity: function lastTarget(address ) view returns(address)
func (_TestValidationHook *TestValidationHookSession) LastTarget(arg0 common.Address) (common.Address, error) {
	return _TestValidationHook.Contract.LastTarget(&_TestValidationHook.CallOpts, arg0)
}

// LastTarget is a free data retrieval call binding the contract method 0x99d48762.
//
// Solidity: function lastTarget(address ) view returns(address)
func (_TestValidationHook *TestValidationHookCallerSession) LastTarget(arg0 common.Address) (common.Address, error) {
	return _TestValidationHook.Contract.LastTarget(&_TestValidationHook.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestValidationHook *TestValidationHookCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TestValidationHook.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestValidationHook *TestValidationHookSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestValidationHook.Contract.SupportsInterface(&_TestValidationHook.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_TestValidationHook *TestValidationHookCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TestValidationHook.Contract.SupportsInterface(&_TestValidationHook.CallOpts, interfaceId)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_TestValidationHook *TestValidationHookTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _TestValidationHook.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_TestValidationHook *TestValidationHookSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _TestValidationHook.Contract.OnInstall(&_TestValidationHook.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_TestValidationHook *TestValidationHookTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _TestValidationHook.Contract.OnInstall(&_TestValidationHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_TestValidationHook *TestValidationHookTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _TestValidationHook.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_TestValidationHook *TestValidationHookSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _TestValidationHook.Contract.OnUninstall(&_TestValidationHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_TestValidationHook *TestValidationHookTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _TestValidationHook.Contract.OnUninstall(&_TestValidationHook.TransactOpts, data)
}

// ValidationHook is a paid mutator transaction binding the contract method 0x37d5f03a.
//
// Solidity: function validationHook(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns()
func (_TestValidationHook *TestValidationHookTransactor) ValidationHook(opts *bind.TransactOpts, signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _TestValidationHook.contract.Transact(opts, "validationHook", signedHash, transaction)
}

// ValidationHook is a paid mutator transaction binding the contract method 0x37d5f03a.
//
// Solidity: function validationHook(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns()
func (_TestValidationHook *TestValidationHookSession) ValidationHook(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _TestValidationHook.Contract.ValidationHook(&_TestValidationHook.TransactOpts, signedHash, transaction)
}

// ValidationHook is a paid mutator transaction binding the contract method 0x37d5f03a.
//
// Solidity: function validationHook(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns()
func (_TestValidationHook *TestValidationHookTransactorSession) ValidationHook(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _TestValidationHook.Contract.ValidationHook(&_TestValidationHook.TransactOpts, signedHash, transaction)
}

// TestValidationHookValidationHookInstalledIterator is returned from FilterValidationHookInstalled and is used to iterate over the raw logs and unpacked data for ValidationHookInstalled events raised by the TestValidationHook contract.
type TestValidationHookValidationHookInstalledIterator struct {
	Event *TestValidationHookValidationHookInstalled // Event containing the contract specifics and raw log

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
func (it *TestValidationHookValidationHookInstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestValidationHookValidationHookInstalled)
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
		it.Event = new(TestValidationHookValidationHookInstalled)
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
func (it *TestValidationHookValidationHookInstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestValidationHookValidationHookInstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestValidationHookValidationHookInstalled represents a ValidationHookInstalled event raised by the TestValidationHook contract.
type TestValidationHookValidationHookInstalled struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidationHookInstalled is a free log retrieval operation binding the contract event 0xd76f8f3aa13d67cda2f8afeae0e727faf7d8a6ceac6431058522a06bbe0c2e3a.
//
// Solidity: event ValidationHookInstalled(address indexed account)
func (_TestValidationHook *TestValidationHookFilterer) FilterValidationHookInstalled(opts *bind.FilterOpts, account []common.Address) (*TestValidationHookValidationHookInstalledIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestValidationHook.contract.FilterLogs(opts, "ValidationHookInstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return &TestValidationHookValidationHookInstalledIterator{contract: _TestValidationHook.contract, event: "ValidationHookInstalled", logs: logs, sub: sub}, nil
}

// WatchValidationHookInstalled is a free log subscription operation binding the contract event 0xd76f8f3aa13d67cda2f8afeae0e727faf7d8a6ceac6431058522a06bbe0c2e3a.
//
// Solidity: event ValidationHookInstalled(address indexed account)
func (_TestValidationHook *TestValidationHookFilterer) WatchValidationHookInstalled(opts *bind.WatchOpts, sink chan<- *TestValidationHookValidationHookInstalled, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestValidationHook.contract.WatchLogs(opts, "ValidationHookInstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestValidationHookValidationHookInstalled)
				if err := _TestValidationHook.contract.UnpackLog(event, "ValidationHookInstalled", log); err != nil {
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

// ParseValidationHookInstalled is a log parse operation binding the contract event 0xd76f8f3aa13d67cda2f8afeae0e727faf7d8a6ceac6431058522a06bbe0c2e3a.
//
// Solidity: event ValidationHookInstalled(address indexed account)
func (_TestValidationHook *TestValidationHookFilterer) ParseValidationHookInstalled(log types.Log) (*TestValidationHookValidationHookInstalled, error) {
	event := new(TestValidationHookValidationHookInstalled)
	if err := _TestValidationHook.contract.UnpackLog(event, "ValidationHookInstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestValidationHookValidationHookTriggeredIterator is returned from FilterValidationHookTriggered and is used to iterate over the raw logs and unpacked data for ValidationHookTriggered events raised by the TestValidationHook contract.
type TestValidationHookValidationHookTriggeredIterator struct {
	Event *TestValidationHookValidationHookTriggered // Event containing the contract specifics and raw log

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
func (it *TestValidationHookValidationHookTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestValidationHookValidationHookTriggered)
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
		it.Event = new(TestValidationHookValidationHookTriggered)
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
func (it *TestValidationHookValidationHookTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestValidationHookValidationHookTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestValidationHookValidationHookTriggered represents a ValidationHookTriggered event raised by the TestValidationHook contract.
type TestValidationHookValidationHookTriggered struct {
	Account common.Address
	Target  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidationHookTriggered is a free log retrieval operation binding the contract event 0x940ef28342a837d716d690d57f37540f3cc25b12a8cba8ba05f292becd2607b0.
//
// Solidity: event ValidationHookTriggered(address indexed account, address indexed target)
func (_TestValidationHook *TestValidationHookFilterer) FilterValidationHookTriggered(opts *bind.FilterOpts, account []common.Address, target []common.Address) (*TestValidationHookValidationHookTriggeredIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TestValidationHook.contract.FilterLogs(opts, "ValidationHookTriggered", accountRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &TestValidationHookValidationHookTriggeredIterator{contract: _TestValidationHook.contract, event: "ValidationHookTriggered", logs: logs, sub: sub}, nil
}

// WatchValidationHookTriggered is a free log subscription operation binding the contract event 0x940ef28342a837d716d690d57f37540f3cc25b12a8cba8ba05f292becd2607b0.
//
// Solidity: event ValidationHookTriggered(address indexed account, address indexed target)
func (_TestValidationHook *TestValidationHookFilterer) WatchValidationHookTriggered(opts *bind.WatchOpts, sink chan<- *TestValidationHookValidationHookTriggered, account []common.Address, target []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TestValidationHook.contract.WatchLogs(opts, "ValidationHookTriggered", accountRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestValidationHookValidationHookTriggered)
				if err := _TestValidationHook.contract.UnpackLog(event, "ValidationHookTriggered", log); err != nil {
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

// ParseValidationHookTriggered is a log parse operation binding the contract event 0x940ef28342a837d716d690d57f37540f3cc25b12a8cba8ba05f292becd2607b0.
//
// Solidity: event ValidationHookTriggered(address indexed account, address indexed target)
func (_TestValidationHook *TestValidationHookFilterer) ParseValidationHookTriggered(log types.Log) (*TestValidationHookValidationHookTriggered, error) {
	event := new(TestValidationHookValidationHookTriggered)
	if err := _TestValidationHook.contract.UnpackLog(event, "ValidationHookTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestValidationHookValidationHookUninstalledIterator is returned from FilterValidationHookUninstalled and is used to iterate over the raw logs and unpacked data for ValidationHookUninstalled events raised by the TestValidationHook contract.
type TestValidationHookValidationHookUninstalledIterator struct {
	Event *TestValidationHookValidationHookUninstalled // Event containing the contract specifics and raw log

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
func (it *TestValidationHookValidationHookUninstalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestValidationHookValidationHookUninstalled)
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
		it.Event = new(TestValidationHookValidationHookUninstalled)
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
func (it *TestValidationHookValidationHookUninstalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestValidationHookValidationHookUninstalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestValidationHookValidationHookUninstalled represents a ValidationHookUninstalled event raised by the TestValidationHook contract.
type TestValidationHookValidationHookUninstalled struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidationHookUninstalled is a free log retrieval operation binding the contract event 0x1055ac92b9f330f212b97da321a321a1f6c5990238245855b75b9d7f6eaa6dce.
//
// Solidity: event ValidationHookUninstalled(address indexed account)
func (_TestValidationHook *TestValidationHookFilterer) FilterValidationHookUninstalled(opts *bind.FilterOpts, account []common.Address) (*TestValidationHookValidationHookUninstalledIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestValidationHook.contract.FilterLogs(opts, "ValidationHookUninstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return &TestValidationHookValidationHookUninstalledIterator{contract: _TestValidationHook.contract, event: "ValidationHookUninstalled", logs: logs, sub: sub}, nil
}

// WatchValidationHookUninstalled is a free log subscription operation binding the contract event 0x1055ac92b9f330f212b97da321a321a1f6c5990238245855b75b9d7f6eaa6dce.
//
// Solidity: event ValidationHookUninstalled(address indexed account)
func (_TestValidationHook *TestValidationHookFilterer) WatchValidationHookUninstalled(opts *bind.WatchOpts, sink chan<- *TestValidationHookValidationHookUninstalled, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestValidationHook.contract.WatchLogs(opts, "ValidationHookUninstalled", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestValidationHookValidationHookUninstalled)
				if err := _TestValidationHook.contract.UnpackLog(event, "ValidationHookUninstalled", log); err != nil {
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

// ParseValidationHookUninstalled is a log parse operation binding the contract event 0x1055ac92b9f330f212b97da321a321a1f6c5990238245855b75b9d7f6eaa6dce.
//
// Solidity: event ValidationHookUninstalled(address indexed account)
func (_TestValidationHook *TestValidationHookFilterer) ParseValidationHookUninstalled(log types.Log) (*TestValidationHookValidationHookUninstalled, error) {
	event := new(TestValidationHookValidationHookUninstalled)
	if err := _TestValidationHook.contract.UnpackLog(event, "ValidationHookUninstalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
