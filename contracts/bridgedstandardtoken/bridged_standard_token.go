// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgedstandardtoken

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

// IBridgedStandardTokenMetaData contains all meta data concerning the IBridgedStandardToken contract.
var IBridgedStandardTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"BridgeInitialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeMint\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assetId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bridgeBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bridgeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBridgedStandardTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgedStandardTokenMetaData.ABI instead.
var IBridgedStandardTokenABI = IBridgedStandardTokenMetaData.ABI

// IBridgedStandardToken is an auto generated Go binding around an Ethereum contract.
type IBridgedStandardToken struct {
	IBridgedStandardTokenCaller     // Read-only binding to the contract
	IBridgedStandardTokenTransactor // Write-only binding to the contract
	IBridgedStandardTokenFilterer   // Log filterer for contract events
}

// IBridgedStandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgedStandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgedStandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgedStandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgedStandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgedStandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgedStandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgedStandardTokenSession struct {
	Contract     *IBridgedStandardToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IBridgedStandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgedStandardTokenCallerSession struct {
	Contract *IBridgedStandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IBridgedStandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgedStandardTokenTransactorSession struct {
	Contract     *IBridgedStandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IBridgedStandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgedStandardTokenRaw struct {
	Contract *IBridgedStandardToken // Generic contract binding to access the raw methods on
}

// IBridgedStandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgedStandardTokenCallerRaw struct {
	Contract *IBridgedStandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgedStandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgedStandardTokenTransactorRaw struct {
	Contract *IBridgedStandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgedStandardToken creates a new instance of IBridgedStandardToken, bound to a specific deployed contract.
func NewIBridgedStandardToken(address common.Address, backend bind.ContractBackend) (*IBridgedStandardToken, error) {
	contract, err := bindIBridgedStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgedStandardToken{IBridgedStandardTokenCaller: IBridgedStandardTokenCaller{contract: contract}, IBridgedStandardTokenTransactor: IBridgedStandardTokenTransactor{contract: contract}, IBridgedStandardTokenFilterer: IBridgedStandardTokenFilterer{contract: contract}}, nil
}

// NewIBridgedStandardTokenCaller creates a new read-only instance of IBridgedStandardToken, bound to a specific deployed contract.
func NewIBridgedStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*IBridgedStandardTokenCaller, error) {
	contract, err := bindIBridgedStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgedStandardTokenCaller{contract: contract}, nil
}

// NewIBridgedStandardTokenTransactor creates a new write-only instance of IBridgedStandardToken, bound to a specific deployed contract.
func NewIBridgedStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgedStandardTokenTransactor, error) {
	contract, err := bindIBridgedStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgedStandardTokenTransactor{contract: contract}, nil
}

// NewIBridgedStandardTokenFilterer creates a new log filterer instance of IBridgedStandardToken, bound to a specific deployed contract.
func NewIBridgedStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgedStandardTokenFilterer, error) {
	contract, err := bindIBridgedStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgedStandardTokenFilterer{contract: contract}, nil
}

// bindIBridgedStandardToken binds a generic wrapper to an already deployed contract.
func bindIBridgedStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBridgedStandardTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgedStandardToken *IBridgedStandardTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgedStandardToken.Contract.IBridgedStandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgedStandardToken *IBridgedStandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.IBridgedStandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgedStandardToken *IBridgedStandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.IBridgedStandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgedStandardToken *IBridgedStandardTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgedStandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgedStandardToken *IBridgedStandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgedStandardToken *IBridgedStandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x44de240a.
//
// Solidity: function assetId() view returns(bytes32)
func (_IBridgedStandardToken *IBridgedStandardTokenCaller) AssetId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IBridgedStandardToken.contract.Call(opts, &out, "assetId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x44de240a.
//
// Solidity: function assetId() view returns(bytes32)
func (_IBridgedStandardToken *IBridgedStandardTokenSession) AssetId() ([32]byte, error) {
	return _IBridgedStandardToken.Contract.AssetId(&_IBridgedStandardToken.CallOpts)
}

// AssetId is a free data retrieval call binding the contract method 0x44de240a.
//
// Solidity: function assetId() view returns(bytes32)
func (_IBridgedStandardToken *IBridgedStandardTokenCallerSession) AssetId() ([32]byte, error) {
	return _IBridgedStandardToken.Contract.AssetId(&_IBridgedStandardToken.CallOpts)
}

// L1Address is a free data retrieval call binding the contract method 0xc2eeeebd.
//
// Solidity: function l1Address() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCaller) L1Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridgedStandardToken.contract.Call(opts, &out, "l1Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Address is a free data retrieval call binding the contract method 0xc2eeeebd.
//
// Solidity: function l1Address() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenSession) L1Address() (common.Address, error) {
	return _IBridgedStandardToken.Contract.L1Address(&_IBridgedStandardToken.CallOpts)
}

// L1Address is a free data retrieval call binding the contract method 0xc2eeeebd.
//
// Solidity: function l1Address() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCallerSession) L1Address() (common.Address, error) {
	return _IBridgedStandardToken.Contract.L1Address(&_IBridgedStandardToken.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridgedStandardToken.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenSession) L2Bridge() (common.Address, error) {
	return _IBridgedStandardToken.Contract.L2Bridge(&_IBridgedStandardToken.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCallerSession) L2Bridge() (common.Address, error) {
	return _IBridgedStandardToken.Contract.L2Bridge(&_IBridgedStandardToken.CallOpts)
}

// NativeTokenVault is a free data retrieval call binding the contract method 0x64e130cf.
//
// Solidity: function nativeTokenVault() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCaller) NativeTokenVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridgedStandardToken.contract.Call(opts, &out, "nativeTokenVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenVault is a free data retrieval call binding the contract method 0x64e130cf.
//
// Solidity: function nativeTokenVault() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenSession) NativeTokenVault() (common.Address, error) {
	return _IBridgedStandardToken.Contract.NativeTokenVault(&_IBridgedStandardToken.CallOpts)
}

// NativeTokenVault is a free data retrieval call binding the contract method 0x64e130cf.
//
// Solidity: function nativeTokenVault() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCallerSession) NativeTokenVault() (common.Address, error) {
	return _IBridgedStandardToken.Contract.NativeTokenVault(&_IBridgedStandardToken.CallOpts)
}

// OriginToken is a free data retrieval call binding the contract method 0x13096a41.
//
// Solidity: function originToken() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCaller) OriginToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridgedStandardToken.contract.Call(opts, &out, "originToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OriginToken is a free data retrieval call binding the contract method 0x13096a41.
//
// Solidity: function originToken() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenSession) OriginToken() (common.Address, error) {
	return _IBridgedStandardToken.Contract.OriginToken(&_IBridgedStandardToken.CallOpts)
}

// OriginToken is a free data retrieval call binding the contract method 0x13096a41.
//
// Solidity: function originToken() view returns(address)
func (_IBridgedStandardToken *IBridgedStandardTokenCallerSession) OriginToken() (common.Address, error) {
	return _IBridgedStandardToken.Contract.OriginToken(&_IBridgedStandardToken.CallOpts)
}

// BridgeBurn is a paid mutator transaction binding the contract method 0x74f4f547.
//
// Solidity: function bridgeBurn(address _account, uint256 _amount) returns()
func (_IBridgedStandardToken *IBridgedStandardTokenTransactor) BridgeBurn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBridgedStandardToken.contract.Transact(opts, "bridgeBurn", _account, _amount)
}

// BridgeBurn is a paid mutator transaction binding the contract method 0x74f4f547.
//
// Solidity: function bridgeBurn(address _account, uint256 _amount) returns()
func (_IBridgedStandardToken *IBridgedStandardTokenSession) BridgeBurn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.BridgeBurn(&_IBridgedStandardToken.TransactOpts, _account, _amount)
}

// BridgeBurn is a paid mutator transaction binding the contract method 0x74f4f547.
//
// Solidity: function bridgeBurn(address _account, uint256 _amount) returns()
func (_IBridgedStandardToken *IBridgedStandardTokenTransactorSession) BridgeBurn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.BridgeBurn(&_IBridgedStandardToken.TransactOpts, _account, _amount)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x8c2a993e.
//
// Solidity: function bridgeMint(address _account, uint256 _amount) returns()
func (_IBridgedStandardToken *IBridgedStandardTokenTransactor) BridgeMint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBridgedStandardToken.contract.Transact(opts, "bridgeMint", _account, _amount)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x8c2a993e.
//
// Solidity: function bridgeMint(address _account, uint256 _amount) returns()
func (_IBridgedStandardToken *IBridgedStandardTokenSession) BridgeMint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.BridgeMint(&_IBridgedStandardToken.TransactOpts, _account, _amount)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x8c2a993e.
//
// Solidity: function bridgeMint(address _account, uint256 _amount) returns()
func (_IBridgedStandardToken *IBridgedStandardTokenTransactorSession) BridgeMint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBridgedStandardToken.Contract.BridgeMint(&_IBridgedStandardToken.TransactOpts, _account, _amount)
}

// IBridgedStandardTokenBridgeBurnIterator is returned from FilterBridgeBurn and is used to iterate over the raw logs and unpacked data for BridgeBurn events raised by the IBridgedStandardToken contract.
type IBridgedStandardTokenBridgeBurnIterator struct {
	Event *IBridgedStandardTokenBridgeBurn // Event containing the contract specifics and raw log

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
func (it *IBridgedStandardTokenBridgeBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgedStandardTokenBridgeBurn)
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
		it.Event = new(IBridgedStandardTokenBridgeBurn)
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
func (it *IBridgedStandardTokenBridgeBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgedStandardTokenBridgeBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgedStandardTokenBridgeBurn represents a BridgeBurn event raised by the IBridgedStandardToken contract.
type IBridgedStandardTokenBridgeBurn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeBurn is a free log retrieval operation binding the contract event 0x9b5b9a05e4726d8bb959f1440e05c6b8109443f2083bc4e386237d7654526553.
//
// Solidity: event BridgeBurn(address indexed account, uint256 amount)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) FilterBridgeBurn(opts *bind.FilterOpts, account []common.Address) (*IBridgedStandardTokenBridgeBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IBridgedStandardToken.contract.FilterLogs(opts, "BridgeBurn", accountRule)
	if err != nil {
		return nil, err
	}
	return &IBridgedStandardTokenBridgeBurnIterator{contract: _IBridgedStandardToken.contract, event: "BridgeBurn", logs: logs, sub: sub}, nil
}

// WatchBridgeBurn is a free log subscription operation binding the contract event 0x9b5b9a05e4726d8bb959f1440e05c6b8109443f2083bc4e386237d7654526553.
//
// Solidity: event BridgeBurn(address indexed account, uint256 amount)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) WatchBridgeBurn(opts *bind.WatchOpts, sink chan<- *IBridgedStandardTokenBridgeBurn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IBridgedStandardToken.contract.WatchLogs(opts, "BridgeBurn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgedStandardTokenBridgeBurn)
				if err := _IBridgedStandardToken.contract.UnpackLog(event, "BridgeBurn", log); err != nil {
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

// ParseBridgeBurn is a log parse operation binding the contract event 0x9b5b9a05e4726d8bb959f1440e05c6b8109443f2083bc4e386237d7654526553.
//
// Solidity: event BridgeBurn(address indexed account, uint256 amount)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) ParseBridgeBurn(log types.Log) (*IBridgedStandardTokenBridgeBurn, error) {
	event := new(IBridgedStandardTokenBridgeBurn)
	if err := _IBridgedStandardToken.contract.UnpackLog(event, "BridgeBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgedStandardTokenBridgeInitializeIterator is returned from FilterBridgeInitialize and is used to iterate over the raw logs and unpacked data for BridgeInitialize events raised by the IBridgedStandardToken contract.
type IBridgedStandardTokenBridgeInitializeIterator struct {
	Event *IBridgedStandardTokenBridgeInitialize // Event containing the contract specifics and raw log

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
func (it *IBridgedStandardTokenBridgeInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgedStandardTokenBridgeInitialize)
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
		it.Event = new(IBridgedStandardTokenBridgeInitialize)
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
func (it *IBridgedStandardTokenBridgeInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgedStandardTokenBridgeInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgedStandardTokenBridgeInitialize represents a BridgeInitialize event raised by the IBridgedStandardToken contract.
type IBridgedStandardTokenBridgeInitialize struct {
	L1Token  common.Address
	Name     string
	Symbol   string
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitialize is a free log retrieval operation binding the contract event 0x81e8e92e5873539605a102eddae7ed06d19bea042099a437cbc3644415eb7404.
//
// Solidity: event BridgeInitialize(address indexed l1Token, string name, string symbol, uint8 decimals)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) FilterBridgeInitialize(opts *bind.FilterOpts, l1Token []common.Address) (*IBridgedStandardTokenBridgeInitializeIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IBridgedStandardToken.contract.FilterLogs(opts, "BridgeInitialize", l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IBridgedStandardTokenBridgeInitializeIterator{contract: _IBridgedStandardToken.contract, event: "BridgeInitialize", logs: logs, sub: sub}, nil
}

// WatchBridgeInitialize is a free log subscription operation binding the contract event 0x81e8e92e5873539605a102eddae7ed06d19bea042099a437cbc3644415eb7404.
//
// Solidity: event BridgeInitialize(address indexed l1Token, string name, string symbol, uint8 decimals)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) WatchBridgeInitialize(opts *bind.WatchOpts, sink chan<- *IBridgedStandardTokenBridgeInitialize, l1Token []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IBridgedStandardToken.contract.WatchLogs(opts, "BridgeInitialize", l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgedStandardTokenBridgeInitialize)
				if err := _IBridgedStandardToken.contract.UnpackLog(event, "BridgeInitialize", log); err != nil {
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

// ParseBridgeInitialize is a log parse operation binding the contract event 0x81e8e92e5873539605a102eddae7ed06d19bea042099a437cbc3644415eb7404.
//
// Solidity: event BridgeInitialize(address indexed l1Token, string name, string symbol, uint8 decimals)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) ParseBridgeInitialize(log types.Log) (*IBridgedStandardTokenBridgeInitialize, error) {
	event := new(IBridgedStandardTokenBridgeInitialize)
	if err := _IBridgedStandardToken.contract.UnpackLog(event, "BridgeInitialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgedStandardTokenBridgeMintIterator is returned from FilterBridgeMint and is used to iterate over the raw logs and unpacked data for BridgeMint events raised by the IBridgedStandardToken contract.
type IBridgedStandardTokenBridgeMintIterator struct {
	Event *IBridgedStandardTokenBridgeMint // Event containing the contract specifics and raw log

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
func (it *IBridgedStandardTokenBridgeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgedStandardTokenBridgeMint)
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
		it.Event = new(IBridgedStandardTokenBridgeMint)
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
func (it *IBridgedStandardTokenBridgeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgedStandardTokenBridgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgedStandardTokenBridgeMint represents a BridgeMint event raised by the IBridgedStandardToken contract.
type IBridgedStandardTokenBridgeMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeMint is a free log retrieval operation binding the contract event 0x397b33b307fc137878ebfc75b295289ec0ee25a31bb5bf034f33256fe8ea2aa6.
//
// Solidity: event BridgeMint(address indexed account, uint256 amount)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) FilterBridgeMint(opts *bind.FilterOpts, account []common.Address) (*IBridgedStandardTokenBridgeMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IBridgedStandardToken.contract.FilterLogs(opts, "BridgeMint", accountRule)
	if err != nil {
		return nil, err
	}
	return &IBridgedStandardTokenBridgeMintIterator{contract: _IBridgedStandardToken.contract, event: "BridgeMint", logs: logs, sub: sub}, nil
}

// WatchBridgeMint is a free log subscription operation binding the contract event 0x397b33b307fc137878ebfc75b295289ec0ee25a31bb5bf034f33256fe8ea2aa6.
//
// Solidity: event BridgeMint(address indexed account, uint256 amount)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) WatchBridgeMint(opts *bind.WatchOpts, sink chan<- *IBridgedStandardTokenBridgeMint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IBridgedStandardToken.contract.WatchLogs(opts, "BridgeMint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgedStandardTokenBridgeMint)
				if err := _IBridgedStandardToken.contract.UnpackLog(event, "BridgeMint", log); err != nil {
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

// ParseBridgeMint is a log parse operation binding the contract event 0x397b33b307fc137878ebfc75b295289ec0ee25a31bb5bf034f33256fe8ea2aa6.
//
// Solidity: event BridgeMint(address indexed account, uint256 amount)
func (_IBridgedStandardToken *IBridgedStandardTokenFilterer) ParseBridgeMint(log types.Log) (*IBridgedStandardTokenBridgeMint, error) {
	event := new(IBridgedStandardTokenBridgeMint)
	if err := _IBridgedStandardToken.contract.UnpackLog(event, "BridgeMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
