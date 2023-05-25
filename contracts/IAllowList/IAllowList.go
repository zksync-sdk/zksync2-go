// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAllowList

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

// IAllowListDeposit is an auto generated low-level Go binding around an user-defined struct.
type IAllowListDeposit struct {
	DepositLimitation bool
	DepositCap        *big.Int
}

// IAllowListMetaData contains all meta data concerning the IAllowList contract.
var IAllowListMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIAllowList.AccessMode\",\"name\":\"previousMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumIAllowList.AccessMode\",\"name\":\"newMode\",\"type\":\"uint8\"}],\"name\":\"UpdateAccessMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"functionSig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateCallPermission\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_functionSig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"getAccessMode\",\"outputs\":[{\"internalType\":\"enumIAllowList.AccessMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"getTokenDepositLimitData\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"depositLimitation\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"depositCap\",\"type\":\"uint256\"}],\"internalType\":\"structIAllowList.Deposit\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_functionSig\",\"type\":\"bytes4\"}],\"name\":\"hasSpecialAccessToCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"enumIAllowList.AccessMode\",\"name\":\"_accessMode\",\"type\":\"uint8\"}],\"name\":\"setAccessMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"enumIAllowList.AccessMode[]\",\"name\":\"_accessMode\",\"type\":\"uint8[]\"}],\"name\":\"setBatchAccessMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_callers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"_functionSigs\",\"type\":\"bytes4[]\"},{\"internalType\":\"bool[]\",\"name\":\"_enables\",\"type\":\"bool[]\"}],\"name\":\"setBatchPermissionToCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_depositLimitation\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_depositCap\",\"type\":\"uint256\"}],\"name\":\"setDepositLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_functionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setPermissionToCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAllowListABI is the input ABI used to generate the binding from.
// Deprecated: Use IAllowListMetaData.ABI instead.
var IAllowListABI = IAllowListMetaData.ABI

// IAllowList is an auto generated Go binding around an Ethereum contract.
type IAllowList struct {
	IAllowListCaller     // Read-only binding to the contract
	IAllowListTransactor // Write-only binding to the contract
	IAllowListFilterer   // Log filterer for contract events
}

// IAllowListCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAllowListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllowListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAllowListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllowListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAllowListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllowListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAllowListSession struct {
	Contract     *IAllowList       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAllowListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAllowListCallerSession struct {
	Contract *IAllowListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IAllowListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAllowListTransactorSession struct {
	Contract     *IAllowListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IAllowListRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAllowListRaw struct {
	Contract *IAllowList // Generic contract binding to access the raw methods on
}

// IAllowListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAllowListCallerRaw struct {
	Contract *IAllowListCaller // Generic read-only contract binding to access the raw methods on
}

// IAllowListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAllowListTransactorRaw struct {
	Contract *IAllowListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAllowList creates a new instance of IAllowList, bound to a specific deployed contract.
func NewIAllowList(address common.Address, backend bind.ContractBackend) (*IAllowList, error) {
	contract, err := bindIAllowList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAllowList{IAllowListCaller: IAllowListCaller{contract: contract}, IAllowListTransactor: IAllowListTransactor{contract: contract}, IAllowListFilterer: IAllowListFilterer{contract: contract}}, nil
}

// NewIAllowListCaller creates a new read-only instance of IAllowList, bound to a specific deployed contract.
func NewIAllowListCaller(address common.Address, caller bind.ContractCaller) (*IAllowListCaller, error) {
	contract, err := bindIAllowList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAllowListCaller{contract: contract}, nil
}

// NewIAllowListTransactor creates a new write-only instance of IAllowList, bound to a specific deployed contract.
func NewIAllowListTransactor(address common.Address, transactor bind.ContractTransactor) (*IAllowListTransactor, error) {
	contract, err := bindIAllowList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAllowListTransactor{contract: contract}, nil
}

// NewIAllowListFilterer creates a new log filterer instance of IAllowList, bound to a specific deployed contract.
func NewIAllowListFilterer(address common.Address, filterer bind.ContractFilterer) (*IAllowListFilterer, error) {
	contract, err := bindIAllowList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAllowListFilterer{contract: contract}, nil
}

// bindIAllowList binds a generic wrapper to an already deployed contract.
func bindIAllowList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAllowListMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllowList *IAllowListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllowList.Contract.IAllowListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllowList *IAllowListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllowList.Contract.IAllowListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllowList *IAllowListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllowList.Contract.IAllowListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllowList *IAllowListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllowList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllowList *IAllowListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllowList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllowList *IAllowListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllowList.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address _caller, address _target, bytes4 _functionSig) view returns(bool)
func (_IAllowList *IAllowListCaller) CanCall(opts *bind.CallOpts, _caller common.Address, _target common.Address, _functionSig [4]byte) (bool, error) {
	var out []interface{}
	err := _IAllowList.contract.Call(opts, &out, "canCall", _caller, _target, _functionSig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address _caller, address _target, bytes4 _functionSig) view returns(bool)
func (_IAllowList *IAllowListSession) CanCall(_caller common.Address, _target common.Address, _functionSig [4]byte) (bool, error) {
	return _IAllowList.Contract.CanCall(&_IAllowList.CallOpts, _caller, _target, _functionSig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address _caller, address _target, bytes4 _functionSig) view returns(bool)
func (_IAllowList *IAllowListCallerSession) CanCall(_caller common.Address, _target common.Address, _functionSig [4]byte) (bool, error) {
	return _IAllowList.Contract.CanCall(&_IAllowList.CallOpts, _caller, _target, _functionSig)
}

// GetAccessMode is a free data retrieval call binding the contract method 0x285712c8.
//
// Solidity: function getAccessMode(address _target) view returns(uint8)
func (_IAllowList *IAllowListCaller) GetAccessMode(opts *bind.CallOpts, _target common.Address) (uint8, error) {
	var out []interface{}
	err := _IAllowList.contract.Call(opts, &out, "getAccessMode", _target)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAccessMode is a free data retrieval call binding the contract method 0x285712c8.
//
// Solidity: function getAccessMode(address _target) view returns(uint8)
func (_IAllowList *IAllowListSession) GetAccessMode(_target common.Address) (uint8, error) {
	return _IAllowList.Contract.GetAccessMode(&_IAllowList.CallOpts, _target)
}

// GetAccessMode is a free data retrieval call binding the contract method 0x285712c8.
//
// Solidity: function getAccessMode(address _target) view returns(uint8)
func (_IAllowList *IAllowListCallerSession) GetAccessMode(_target common.Address) (uint8, error) {
	return _IAllowList.Contract.GetAccessMode(&_IAllowList.CallOpts, _target)
}

// GetTokenDepositLimitData is a free data retrieval call binding the contract method 0x7cf14701.
//
// Solidity: function getTokenDepositLimitData(address _l1Token) view returns((bool,uint256))
func (_IAllowList *IAllowListCaller) GetTokenDepositLimitData(opts *bind.CallOpts, _l1Token common.Address) (IAllowListDeposit, error) {
	var out []interface{}
	err := _IAllowList.contract.Call(opts, &out, "getTokenDepositLimitData", _l1Token)

	if err != nil {
		return *new(IAllowListDeposit), err
	}

	out0 := *abi.ConvertType(out[0], new(IAllowListDeposit)).(*IAllowListDeposit)

	return out0, err

}

// GetTokenDepositLimitData is a free data retrieval call binding the contract method 0x7cf14701.
//
// Solidity: function getTokenDepositLimitData(address _l1Token) view returns((bool,uint256))
func (_IAllowList *IAllowListSession) GetTokenDepositLimitData(_l1Token common.Address) (IAllowListDeposit, error) {
	return _IAllowList.Contract.GetTokenDepositLimitData(&_IAllowList.CallOpts, _l1Token)
}

// GetTokenDepositLimitData is a free data retrieval call binding the contract method 0x7cf14701.
//
// Solidity: function getTokenDepositLimitData(address _l1Token) view returns((bool,uint256))
func (_IAllowList *IAllowListCallerSession) GetTokenDepositLimitData(_l1Token common.Address) (IAllowListDeposit, error) {
	return _IAllowList.Contract.GetTokenDepositLimitData(&_IAllowList.CallOpts, _l1Token)
}

// HasSpecialAccessToCall is a free data retrieval call binding the contract method 0x5965cf8c.
//
// Solidity: function hasSpecialAccessToCall(address _caller, address _target, bytes4 _functionSig) view returns(bool)
func (_IAllowList *IAllowListCaller) HasSpecialAccessToCall(opts *bind.CallOpts, _caller common.Address, _target common.Address, _functionSig [4]byte) (bool, error) {
	var out []interface{}
	err := _IAllowList.contract.Call(opts, &out, "hasSpecialAccessToCall", _caller, _target, _functionSig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSpecialAccessToCall is a free data retrieval call binding the contract method 0x5965cf8c.
//
// Solidity: function hasSpecialAccessToCall(address _caller, address _target, bytes4 _functionSig) view returns(bool)
func (_IAllowList *IAllowListSession) HasSpecialAccessToCall(_caller common.Address, _target common.Address, _functionSig [4]byte) (bool, error) {
	return _IAllowList.Contract.HasSpecialAccessToCall(&_IAllowList.CallOpts, _caller, _target, _functionSig)
}

// HasSpecialAccessToCall is a free data retrieval call binding the contract method 0x5965cf8c.
//
// Solidity: function hasSpecialAccessToCall(address _caller, address _target, bytes4 _functionSig) view returns(bool)
func (_IAllowList *IAllowListCallerSession) HasSpecialAccessToCall(_caller common.Address, _target common.Address, _functionSig [4]byte) (bool, error) {
	return _IAllowList.Contract.HasSpecialAccessToCall(&_IAllowList.CallOpts, _caller, _target, _functionSig)
}

// SetAccessMode is a paid mutator transaction binding the contract method 0xffc7d0b4.
//
// Solidity: function setAccessMode(address _target, uint8 _accessMode) returns()
func (_IAllowList *IAllowListTransactor) SetAccessMode(opts *bind.TransactOpts, _target common.Address, _accessMode uint8) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setAccessMode", _target, _accessMode)
}

// SetAccessMode is a paid mutator transaction binding the contract method 0xffc7d0b4.
//
// Solidity: function setAccessMode(address _target, uint8 _accessMode) returns()
func (_IAllowList *IAllowListSession) SetAccessMode(_target common.Address, _accessMode uint8) (*types.Transaction, error) {
	return _IAllowList.Contract.SetAccessMode(&_IAllowList.TransactOpts, _target, _accessMode)
}

// SetAccessMode is a paid mutator transaction binding the contract method 0xffc7d0b4.
//
// Solidity: function setAccessMode(address _target, uint8 _accessMode) returns()
func (_IAllowList *IAllowListTransactorSession) SetAccessMode(_target common.Address, _accessMode uint8) (*types.Transaction, error) {
	return _IAllowList.Contract.SetAccessMode(&_IAllowList.TransactOpts, _target, _accessMode)
}

// SetBatchAccessMode is a paid mutator transaction binding the contract method 0xec5f109a.
//
// Solidity: function setBatchAccessMode(address[] _targets, uint8[] _accessMode) returns()
func (_IAllowList *IAllowListTransactor) SetBatchAccessMode(opts *bind.TransactOpts, _targets []common.Address, _accessMode []uint8) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setBatchAccessMode", _targets, _accessMode)
}

// SetBatchAccessMode is a paid mutator transaction binding the contract method 0xec5f109a.
//
// Solidity: function setBatchAccessMode(address[] _targets, uint8[] _accessMode) returns()
func (_IAllowList *IAllowListSession) SetBatchAccessMode(_targets []common.Address, _accessMode []uint8) (*types.Transaction, error) {
	return _IAllowList.Contract.SetBatchAccessMode(&_IAllowList.TransactOpts, _targets, _accessMode)
}

// SetBatchAccessMode is a paid mutator transaction binding the contract method 0xec5f109a.
//
// Solidity: function setBatchAccessMode(address[] _targets, uint8[] _accessMode) returns()
func (_IAllowList *IAllowListTransactorSession) SetBatchAccessMode(_targets []common.Address, _accessMode []uint8) (*types.Transaction, error) {
	return _IAllowList.Contract.SetBatchAccessMode(&_IAllowList.TransactOpts, _targets, _accessMode)
}

// SetBatchPermissionToCall is a paid mutator transaction binding the contract method 0x507d1bed.
//
// Solidity: function setBatchPermissionToCall(address[] _callers, address[] _targets, bytes4[] _functionSigs, bool[] _enables) returns()
func (_IAllowList *IAllowListTransactor) SetBatchPermissionToCall(opts *bind.TransactOpts, _callers []common.Address, _targets []common.Address, _functionSigs [][4]byte, _enables []bool) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setBatchPermissionToCall", _callers, _targets, _functionSigs, _enables)
}

// SetBatchPermissionToCall is a paid mutator transaction binding the contract method 0x507d1bed.
//
// Solidity: function setBatchPermissionToCall(address[] _callers, address[] _targets, bytes4[] _functionSigs, bool[] _enables) returns()
func (_IAllowList *IAllowListSession) SetBatchPermissionToCall(_callers []common.Address, _targets []common.Address, _functionSigs [][4]byte, _enables []bool) (*types.Transaction, error) {
	return _IAllowList.Contract.SetBatchPermissionToCall(&_IAllowList.TransactOpts, _callers, _targets, _functionSigs, _enables)
}

// SetBatchPermissionToCall is a paid mutator transaction binding the contract method 0x507d1bed.
//
// Solidity: function setBatchPermissionToCall(address[] _callers, address[] _targets, bytes4[] _functionSigs, bool[] _enables) returns()
func (_IAllowList *IAllowListTransactorSession) SetBatchPermissionToCall(_callers []common.Address, _targets []common.Address, _functionSigs [][4]byte, _enables []bool) (*types.Transaction, error) {
	return _IAllowList.Contract.SetBatchPermissionToCall(&_IAllowList.TransactOpts, _callers, _targets, _functionSigs, _enables)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xdff05449.
//
// Solidity: function setDepositLimit(address _l1Token, bool _depositLimitation, uint256 _depositCap) returns()
func (_IAllowList *IAllowListTransactor) SetDepositLimit(opts *bind.TransactOpts, _l1Token common.Address, _depositLimitation bool, _depositCap *big.Int) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setDepositLimit", _l1Token, _depositLimitation, _depositCap)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xdff05449.
//
// Solidity: function setDepositLimit(address _l1Token, bool _depositLimitation, uint256 _depositCap) returns()
func (_IAllowList *IAllowListSession) SetDepositLimit(_l1Token common.Address, _depositLimitation bool, _depositCap *big.Int) (*types.Transaction, error) {
	return _IAllowList.Contract.SetDepositLimit(&_IAllowList.TransactOpts, _l1Token, _depositLimitation, _depositCap)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xdff05449.
//
// Solidity: function setDepositLimit(address _l1Token, bool _depositLimitation, uint256 _depositCap) returns()
func (_IAllowList *IAllowListTransactorSession) SetDepositLimit(_l1Token common.Address, _depositLimitation bool, _depositCap *big.Int) (*types.Transaction, error) {
	return _IAllowList.Contract.SetDepositLimit(&_IAllowList.TransactOpts, _l1Token, _depositLimitation, _depositCap)
}

// SetPermissionToCall is a paid mutator transaction binding the contract method 0x73df5d8d.
//
// Solidity: function setPermissionToCall(address _caller, address _target, bytes4 _functionSig, bool _enable) returns()
func (_IAllowList *IAllowListTransactor) SetPermissionToCall(opts *bind.TransactOpts, _caller common.Address, _target common.Address, _functionSig [4]byte, _enable bool) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setPermissionToCall", _caller, _target, _functionSig, _enable)
}

// SetPermissionToCall is a paid mutator transaction binding the contract method 0x73df5d8d.
//
// Solidity: function setPermissionToCall(address _caller, address _target, bytes4 _functionSig, bool _enable) returns()
func (_IAllowList *IAllowListSession) SetPermissionToCall(_caller common.Address, _target common.Address, _functionSig [4]byte, _enable bool) (*types.Transaction, error) {
	return _IAllowList.Contract.SetPermissionToCall(&_IAllowList.TransactOpts, _caller, _target, _functionSig, _enable)
}

// SetPermissionToCall is a paid mutator transaction binding the contract method 0x73df5d8d.
//
// Solidity: function setPermissionToCall(address _caller, address _target, bytes4 _functionSig, bool _enable) returns()
func (_IAllowList *IAllowListTransactorSession) SetPermissionToCall(_caller common.Address, _target common.Address, _functionSig [4]byte, _enable bool) (*types.Transaction, error) {
	return _IAllowList.Contract.SetPermissionToCall(&_IAllowList.TransactOpts, _caller, _target, _functionSig, _enable)
}

// IAllowListUpdateAccessModeIterator is returned from FilterUpdateAccessMode and is used to iterate over the raw logs and unpacked data for UpdateAccessMode events raised by the IAllowList contract.
type IAllowListUpdateAccessModeIterator struct {
	Event *IAllowListUpdateAccessMode // Event containing the contract specifics and raw log

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
func (it *IAllowListUpdateAccessModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllowListUpdateAccessMode)
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
		it.Event = new(IAllowListUpdateAccessMode)
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
func (it *IAllowListUpdateAccessModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllowListUpdateAccessModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllowListUpdateAccessMode represents a UpdateAccessMode event raised by the IAllowList contract.
type IAllowListUpdateAccessMode struct {
	Target       common.Address
	PreviousMode uint8
	NewMode      uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAccessMode is a free log retrieval operation binding the contract event 0x3b3c3f982e4b12b1870d2ff77adfdb97d3838faab0fd8a6b255160e52e79a82f.
//
// Solidity: event UpdateAccessMode(address indexed target, uint8 previousMode, uint8 newMode)
func (_IAllowList *IAllowListFilterer) FilterUpdateAccessMode(opts *bind.FilterOpts, target []common.Address) (*IAllowListUpdateAccessModeIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IAllowList.contract.FilterLogs(opts, "UpdateAccessMode", targetRule)
	if err != nil {
		return nil, err
	}
	return &IAllowListUpdateAccessModeIterator{contract: _IAllowList.contract, event: "UpdateAccessMode", logs: logs, sub: sub}, nil
}

// WatchUpdateAccessMode is a free log subscription operation binding the contract event 0x3b3c3f982e4b12b1870d2ff77adfdb97d3838faab0fd8a6b255160e52e79a82f.
//
// Solidity: event UpdateAccessMode(address indexed target, uint8 previousMode, uint8 newMode)
func (_IAllowList *IAllowListFilterer) WatchUpdateAccessMode(opts *bind.WatchOpts, sink chan<- *IAllowListUpdateAccessMode, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IAllowList.contract.WatchLogs(opts, "UpdateAccessMode", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllowListUpdateAccessMode)
				if err := _IAllowList.contract.UnpackLog(event, "UpdateAccessMode", log); err != nil {
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

// ParseUpdateAccessMode is a log parse operation binding the contract event 0x3b3c3f982e4b12b1870d2ff77adfdb97d3838faab0fd8a6b255160e52e79a82f.
//
// Solidity: event UpdateAccessMode(address indexed target, uint8 previousMode, uint8 newMode)
func (_IAllowList *IAllowListFilterer) ParseUpdateAccessMode(log types.Log) (*IAllowListUpdateAccessMode, error) {
	event := new(IAllowListUpdateAccessMode)
	if err := _IAllowList.contract.UnpackLog(event, "UpdateAccessMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllowListUpdateCallPermissionIterator is returned from FilterUpdateCallPermission and is used to iterate over the raw logs and unpacked data for UpdateCallPermission events raised by the IAllowList contract.
type IAllowListUpdateCallPermissionIterator struct {
	Event *IAllowListUpdateCallPermission // Event containing the contract specifics and raw log

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
func (it *IAllowListUpdateCallPermissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllowListUpdateCallPermission)
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
		it.Event = new(IAllowListUpdateCallPermission)
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
func (it *IAllowListUpdateCallPermissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllowListUpdateCallPermissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllowListUpdateCallPermission represents a UpdateCallPermission event raised by the IAllowList contract.
type IAllowListUpdateCallPermission struct {
	Caller      common.Address
	Target      common.Address
	FunctionSig [4]byte
	Status      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateCallPermission is a free log retrieval operation binding the contract event 0x3336e7aa4c86fcb95fa993c8022c30690f1f696f67f138c845d81dc5484c9a32.
//
// Solidity: event UpdateCallPermission(address indexed caller, address indexed target, bytes4 indexed functionSig, bool status)
func (_IAllowList *IAllowListFilterer) FilterUpdateCallPermission(opts *bind.FilterOpts, caller []common.Address, target []common.Address, functionSig [][4]byte) (*IAllowListUpdateCallPermissionIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var functionSigRule []interface{}
	for _, functionSigItem := range functionSig {
		functionSigRule = append(functionSigRule, functionSigItem)
	}

	logs, sub, err := _IAllowList.contract.FilterLogs(opts, "UpdateCallPermission", callerRule, targetRule, functionSigRule)
	if err != nil {
		return nil, err
	}
	return &IAllowListUpdateCallPermissionIterator{contract: _IAllowList.contract, event: "UpdateCallPermission", logs: logs, sub: sub}, nil
}

// WatchUpdateCallPermission is a free log subscription operation binding the contract event 0x3336e7aa4c86fcb95fa993c8022c30690f1f696f67f138c845d81dc5484c9a32.
//
// Solidity: event UpdateCallPermission(address indexed caller, address indexed target, bytes4 indexed functionSig, bool status)
func (_IAllowList *IAllowListFilterer) WatchUpdateCallPermission(opts *bind.WatchOpts, sink chan<- *IAllowListUpdateCallPermission, caller []common.Address, target []common.Address, functionSig [][4]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var functionSigRule []interface{}
	for _, functionSigItem := range functionSig {
		functionSigRule = append(functionSigRule, functionSigItem)
	}

	logs, sub, err := _IAllowList.contract.WatchLogs(opts, "UpdateCallPermission", callerRule, targetRule, functionSigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllowListUpdateCallPermission)
				if err := _IAllowList.contract.UnpackLog(event, "UpdateCallPermission", log); err != nil {
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

// ParseUpdateCallPermission is a log parse operation binding the contract event 0x3336e7aa4c86fcb95fa993c8022c30690f1f696f67f138c845d81dc5484c9a32.
//
// Solidity: event UpdateCallPermission(address indexed caller, address indexed target, bytes4 indexed functionSig, bool status)
func (_IAllowList *IAllowListFilterer) ParseUpdateCallPermission(log types.Log) (*IAllowListUpdateCallPermission, error) {
	event := new(IAllowListUpdateCallPermission)
	if err := _IAllowList.contract.UnpackLog(event, "UpdateCallPermission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
