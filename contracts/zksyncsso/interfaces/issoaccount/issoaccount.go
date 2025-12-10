// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package issoaccount

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

// ISsoAccountMetaData contains all meta data concerning the ISsoAccount contract.
var ISsoAccountMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_suggestedSignedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"executeTransactionFromOutside\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"initialValidators\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"initialK1Owners\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHook\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isK1Owner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isModuleValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"listHooks\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"hookList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listK1Owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"k1OwnerList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listModuleValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_suggestedSignedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"payForTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_possibleSignedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"prepareForPaymaster\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_suggestedSignedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magic\",\"type\":\"bytes4\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ISsoAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use ISsoAccountMetaData.ABI instead.
var ISsoAccountABI = ISsoAccountMetaData.ABI

// ISsoAccount is an auto generated Go binding around an Ethereum contract.
type ISsoAccount struct {
	ISsoAccountCaller     // Read-only binding to the contract
	ISsoAccountTransactor // Write-only binding to the contract
	ISsoAccountFilterer   // Log filterer for contract events
}

// ISsoAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISsoAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISsoAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISsoAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISsoAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISsoAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISsoAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISsoAccountSession struct {
	Contract     *ISsoAccount      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISsoAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISsoAccountCallerSession struct {
	Contract *ISsoAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ISsoAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISsoAccountTransactorSession struct {
	Contract     *ISsoAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISsoAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISsoAccountRaw struct {
	Contract *ISsoAccount // Generic contract binding to access the raw methods on
}

// ISsoAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISsoAccountCallerRaw struct {
	Contract *ISsoAccountCaller // Generic read-only contract binding to access the raw methods on
}

// ISsoAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISsoAccountTransactorRaw struct {
	Contract *ISsoAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISsoAccount creates a new instance of ISsoAccount, bound to a specific deployed contract.
func NewISsoAccount(address common.Address, backend bind.ContractBackend) (*ISsoAccount, error) {
	contract, err := bindISsoAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISsoAccount{ISsoAccountCaller: ISsoAccountCaller{contract: contract}, ISsoAccountTransactor: ISsoAccountTransactor{contract: contract}, ISsoAccountFilterer: ISsoAccountFilterer{contract: contract}}, nil
}

// NewISsoAccountCaller creates a new read-only instance of ISsoAccount, bound to a specific deployed contract.
func NewISsoAccountCaller(address common.Address, caller bind.ContractCaller) (*ISsoAccountCaller, error) {
	contract, err := bindISsoAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountCaller{contract: contract}, nil
}

// NewISsoAccountTransactor creates a new write-only instance of ISsoAccount, bound to a specific deployed contract.
func NewISsoAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*ISsoAccountTransactor, error) {
	contract, err := bindISsoAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountTransactor{contract: contract}, nil
}

// NewISsoAccountFilterer creates a new log filterer instance of ISsoAccount, bound to a specific deployed contract.
func NewISsoAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*ISsoAccountFilterer, error) {
	contract, err := bindISsoAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountFilterer{contract: contract}, nil
}

// bindISsoAccount binds a generic wrapper to an already deployed contract.
func bindISsoAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISsoAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISsoAccount *ISsoAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISsoAccount.Contract.ISsoAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISsoAccount *ISsoAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ISsoAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISsoAccount *ISsoAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ISsoAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISsoAccount *ISsoAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISsoAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISsoAccount *ISsoAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISsoAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISsoAccount *ISsoAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISsoAccount.Contract.contract.Transact(opts, method, params...)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_ISsoAccount *ISsoAccountCaller) IsHook(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "isHook", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_ISsoAccount *ISsoAccountSession) IsHook(addr common.Address) (bool, error) {
	return _ISsoAccount.Contract.IsHook(&_ISsoAccount.CallOpts, addr)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_ISsoAccount *ISsoAccountCallerSession) IsHook(addr common.Address) (bool, error) {
	return _ISsoAccount.Contract.IsHook(&_ISsoAccount.CallOpts, addr)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_ISsoAccount *ISsoAccountCaller) IsK1Owner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "isK1Owner", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_ISsoAccount *ISsoAccountSession) IsK1Owner(addr common.Address) (bool, error) {
	return _ISsoAccount.Contract.IsK1Owner(&_ISsoAccount.CallOpts, addr)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_ISsoAccount *ISsoAccountCallerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _ISsoAccount.Contract.IsK1Owner(&_ISsoAccount.CallOpts, addr)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ISsoAccount *ISsoAccountCaller) IsModuleValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "isModuleValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ISsoAccount *ISsoAccountSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _ISsoAccount.Contract.IsModuleValidator(&_ISsoAccount.CallOpts, validator)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_ISsoAccount *ISsoAccountCallerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _ISsoAccount.Contract.IsModuleValidator(&_ISsoAccount.CallOpts, validator)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ISsoAccount *ISsoAccountCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ISsoAccount *ISsoAccountSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _ISsoAccount.Contract.IsValidSignature(&_ISsoAccount.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_ISsoAccount *ISsoAccountCallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _ISsoAccount.Contract.IsValidSignature(&_ISsoAccount.CallOpts, hash, signature)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_ISsoAccount *ISsoAccountCaller) ListHooks(opts *bind.CallOpts, isValidation bool) ([]common.Address, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "listHooks", isValidation)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_ISsoAccount *ISsoAccountSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _ISsoAccount.Contract.ListHooks(&_ISsoAccount.CallOpts, isValidation)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_ISsoAccount *ISsoAccountCallerSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _ISsoAccount.Contract.ListHooks(&_ISsoAccount.CallOpts, isValidation)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_ISsoAccount *ISsoAccountCaller) ListK1Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "listK1Owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_ISsoAccount *ISsoAccountSession) ListK1Owners() ([]common.Address, error) {
	return _ISsoAccount.Contract.ListK1Owners(&_ISsoAccount.CallOpts)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_ISsoAccount *ISsoAccountCallerSession) ListK1Owners() ([]common.Address, error) {
	return _ISsoAccount.Contract.ListK1Owners(&_ISsoAccount.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ISsoAccount *ISsoAccountCaller) ListModuleValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "listModuleValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ISsoAccount *ISsoAccountSession) ListModuleValidators() ([]common.Address, error) {
	return _ISsoAccount.Contract.ListModuleValidators(&_ISsoAccount.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_ISsoAccount *ISsoAccountCallerSession) ListModuleValidators() ([]common.Address, error) {
	return _ISsoAccount.Contract.ListModuleValidators(&_ISsoAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ISsoAccount *ISsoAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ISsoAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ISsoAccount *ISsoAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ISsoAccount.Contract.SupportsInterface(&_ISsoAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ISsoAccount *ISsoAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ISsoAccount.Contract.SupportsInterface(&_ISsoAccount.CallOpts, interfaceId)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_ISsoAccount *ISsoAccountTransactor) AddHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "addHook", hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_ISsoAccount *ISsoAccountSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.AddHook(&_ISsoAccount.TransactOpts, hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.AddHook(&_ISsoAccount.TransactOpts, hook, isValidation, initData)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_ISsoAccount *ISsoAccountTransactor) AddK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "addK1Owner", addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_ISsoAccount *ISsoAccountSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ISsoAccount.Contract.AddK1Owner(&_ISsoAccount.TransactOpts, addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ISsoAccount.Contract.AddK1Owner(&_ISsoAccount.TransactOpts, addr)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ISsoAccount *ISsoAccountTransactor) AddModuleValidator(opts *bind.TransactOpts, validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "addModuleValidator", validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ISsoAccount *ISsoAccountSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.AddModuleValidator(&_ISsoAccount.TransactOpts, validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.AddModuleValidator(&_ISsoAccount.TransactOpts, validator, initData)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xdf9c1589.
//
// Solidity: function executeTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactor) ExecuteTransaction(opts *bind.TransactOpts, _txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "executeTransaction", _txHash, _suggestedSignedHash, _transaction)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xdf9c1589.
//
// Solidity: function executeTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountSession) ExecuteTransaction(_txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ExecuteTransaction(&_ISsoAccount.TransactOpts, _txHash, _suggestedSignedHash, _transaction)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xdf9c1589.
//
// Solidity: function executeTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactorSession) ExecuteTransaction(_txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ExecuteTransaction(&_ISsoAccount.TransactOpts, _txHash, _suggestedSignedHash, _transaction)
}

// ExecuteTransactionFromOutside is a paid mutator transaction binding the contract method 0xeeb8cb09.
//
// Solidity: function executeTransactionFromOutside((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactor) ExecuteTransactionFromOutside(opts *bind.TransactOpts, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "executeTransactionFromOutside", _transaction)
}

// ExecuteTransactionFromOutside is a paid mutator transaction binding the contract method 0xeeb8cb09.
//
// Solidity: function executeTransactionFromOutside((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountSession) ExecuteTransactionFromOutside(_transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ExecuteTransactionFromOutside(&_ISsoAccount.TransactOpts, _transaction)
}

// ExecuteTransactionFromOutside is a paid mutator transaction binding the contract method 0xeeb8cb09.
//
// Solidity: function executeTransactionFromOutside((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactorSession) ExecuteTransactionFromOutside(_transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ExecuteTransactionFromOutside(&_ISsoAccount.TransactOpts, _transaction)
}

// Initialize is a paid mutator transaction binding the contract method 0xb9094997.
//
// Solidity: function initialize(bytes[] initialValidators, address[] initialK1Owners) returns()
func (_ISsoAccount *ISsoAccountTransactor) Initialize(opts *bind.TransactOpts, initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "initialize", initialValidators, initialK1Owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xb9094997.
//
// Solidity: function initialize(bytes[] initialValidators, address[] initialK1Owners) returns()
func (_ISsoAccount *ISsoAccountSession) Initialize(initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _ISsoAccount.Contract.Initialize(&_ISsoAccount.TransactOpts, initialValidators, initialK1Owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xb9094997.
//
// Solidity: function initialize(bytes[] initialValidators, address[] initialK1Owners) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) Initialize(initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _ISsoAccount.Contract.Initialize(&_ISsoAccount.TransactOpts, initialValidators, initialK1Owners)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "onERC1155BatchReceived", operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.OnERC1155BatchReceived(&_ISsoAccount.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.OnERC1155BatchReceived(&_ISsoAccount.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "onERC1155Received", operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.OnERC1155Received(&_ISsoAccount.TransactOpts, operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountTransactorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.OnERC1155Received(&_ISsoAccount.TransactOpts, operator, from, id, value, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.OnERC721Received(&_ISsoAccount.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ISsoAccount *ISsoAccountTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.OnERC721Received(&_ISsoAccount.TransactOpts, operator, from, tokenId, data)
}

// PayForTransaction is a paid mutator transaction binding the contract method 0xe2f318e3.
//
// Solidity: function payForTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactor) PayForTransaction(opts *bind.TransactOpts, _txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "payForTransaction", _txHash, _suggestedSignedHash, _transaction)
}

// PayForTransaction is a paid mutator transaction binding the contract method 0xe2f318e3.
//
// Solidity: function payForTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountSession) PayForTransaction(_txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.PayForTransaction(&_ISsoAccount.TransactOpts, _txHash, _suggestedSignedHash, _transaction)
}

// PayForTransaction is a paid mutator transaction binding the contract method 0xe2f318e3.
//
// Solidity: function payForTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactorSession) PayForTransaction(_txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.PayForTransaction(&_ISsoAccount.TransactOpts, _txHash, _suggestedSignedHash, _transaction)
}

// PrepareForPaymaster is a paid mutator transaction binding the contract method 0xa28c1aee.
//
// Solidity: function prepareForPaymaster(bytes32 _txHash, bytes32 _possibleSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactor) PrepareForPaymaster(opts *bind.TransactOpts, _txHash [32]byte, _possibleSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "prepareForPaymaster", _txHash, _possibleSignedHash, _transaction)
}

// PrepareForPaymaster is a paid mutator transaction binding the contract method 0xa28c1aee.
//
// Solidity: function prepareForPaymaster(bytes32 _txHash, bytes32 _possibleSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountSession) PrepareForPaymaster(_txHash [32]byte, _possibleSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.PrepareForPaymaster(&_ISsoAccount.TransactOpts, _txHash, _possibleSignedHash, _transaction)
}

// PrepareForPaymaster is a paid mutator transaction binding the contract method 0xa28c1aee.
//
// Solidity: function prepareForPaymaster(bytes32 _txHash, bytes32 _possibleSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_ISsoAccount *ISsoAccountTransactorSession) PrepareForPaymaster(_txHash [32]byte, _possibleSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.PrepareForPaymaster(&_ISsoAccount.TransactOpts, _txHash, _possibleSignedHash, _transaction)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactor) RemoveHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "removeHook", hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.RemoveHook(&_ISsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.RemoveHook(&_ISsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_ISsoAccount *ISsoAccountTransactor) RemoveK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "removeK1Owner", addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_ISsoAccount *ISsoAccountSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ISsoAccount.Contract.RemoveK1Owner(&_ISsoAccount.TransactOpts, addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _ISsoAccount.Contract.RemoveK1Owner(&_ISsoAccount.TransactOpts, addr)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactor) RemoveModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "removeModuleValidator", validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.RemoveModuleValidator(&_ISsoAccount.TransactOpts, validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.RemoveModuleValidator(&_ISsoAccount.TransactOpts, validator, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactor) UnlinkHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "unlinkHook", hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.UnlinkHook(&_ISsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.UnlinkHook(&_ISsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactor) UnlinkModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "unlinkModuleValidator", validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.UnlinkModuleValidator(&_ISsoAccount.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_ISsoAccount *ISsoAccountTransactorSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _ISsoAccount.Contract.UnlinkModuleValidator(&_ISsoAccount.TransactOpts, validator, deinitData)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x202bcce7.
//
// Solidity: function validateTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic)
func (_ISsoAccount *ISsoAccountTransactor) ValidateTransaction(opts *bind.TransactOpts, _txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.contract.Transact(opts, "validateTransaction", _txHash, _suggestedSignedHash, _transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x202bcce7.
//
// Solidity: function validateTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic)
func (_ISsoAccount *ISsoAccountSession) ValidateTransaction(_txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ValidateTransaction(&_ISsoAccount.TransactOpts, _txHash, _suggestedSignedHash, _transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x202bcce7.
//
// Solidity: function validateTransaction(bytes32 _txHash, bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic)
func (_ISsoAccount *ISsoAccountTransactorSession) ValidateTransaction(_txHash [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ISsoAccount.Contract.ValidateTransaction(&_ISsoAccount.TransactOpts, _txHash, _suggestedSignedHash, _transaction)
}

// ISsoAccountHookAddedIterator is returned from FilterHookAdded and is used to iterate over the raw logs and unpacked data for HookAdded events raised by the ISsoAccount contract.
type ISsoAccountHookAddedIterator struct {
	Event *ISsoAccountHookAdded // Event containing the contract specifics and raw log

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
func (it *ISsoAccountHookAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISsoAccountHookAdded)
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
		it.Event = new(ISsoAccountHookAdded)
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
func (it *ISsoAccountHookAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISsoAccountHookAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISsoAccountHookAdded represents a HookAdded event raised by the ISsoAccount contract.
type ISsoAccountHookAdded struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookAdded is a free log retrieval operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_ISsoAccount *ISsoAccountFilterer) FilterHookAdded(opts *bind.FilterOpts, hook []common.Address) (*ISsoAccountHookAddedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _ISsoAccount.contract.FilterLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountHookAddedIterator{contract: _ISsoAccount.contract, event: "HookAdded", logs: logs, sub: sub}, nil
}

// WatchHookAdded is a free log subscription operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_ISsoAccount *ISsoAccountFilterer) WatchHookAdded(opts *bind.WatchOpts, sink chan<- *ISsoAccountHookAdded, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _ISsoAccount.contract.WatchLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISsoAccountHookAdded)
				if err := _ISsoAccount.contract.UnpackLog(event, "HookAdded", log); err != nil {
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

// ParseHookAdded is a log parse operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_ISsoAccount *ISsoAccountFilterer) ParseHookAdded(log types.Log) (*ISsoAccountHookAdded, error) {
	event := new(ISsoAccountHookAdded)
	if err := _ISsoAccount.contract.UnpackLog(event, "HookAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISsoAccountHookRemovedIterator is returned from FilterHookRemoved and is used to iterate over the raw logs and unpacked data for HookRemoved events raised by the ISsoAccount contract.
type ISsoAccountHookRemovedIterator struct {
	Event *ISsoAccountHookRemoved // Event containing the contract specifics and raw log

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
func (it *ISsoAccountHookRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISsoAccountHookRemoved)
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
		it.Event = new(ISsoAccountHookRemoved)
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
func (it *ISsoAccountHookRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISsoAccountHookRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISsoAccountHookRemoved represents a HookRemoved event raised by the ISsoAccount contract.
type ISsoAccountHookRemoved struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookRemoved is a free log retrieval operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_ISsoAccount *ISsoAccountFilterer) FilterHookRemoved(opts *bind.FilterOpts, hook []common.Address) (*ISsoAccountHookRemovedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _ISsoAccount.contract.FilterLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountHookRemovedIterator{contract: _ISsoAccount.contract, event: "HookRemoved", logs: logs, sub: sub}, nil
}

// WatchHookRemoved is a free log subscription operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_ISsoAccount *ISsoAccountFilterer) WatchHookRemoved(opts *bind.WatchOpts, sink chan<- *ISsoAccountHookRemoved, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _ISsoAccount.contract.WatchLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISsoAccountHookRemoved)
				if err := _ISsoAccount.contract.UnpackLog(event, "HookRemoved", log); err != nil {
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

// ParseHookRemoved is a log parse operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_ISsoAccount *ISsoAccountFilterer) ParseHookRemoved(log types.Log) (*ISsoAccountHookRemoved, error) {
	event := new(ISsoAccountHookRemoved)
	if err := _ISsoAccount.contract.UnpackLog(event, "HookRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISsoAccountK1OwnerAddedIterator is returned from FilterK1OwnerAdded and is used to iterate over the raw logs and unpacked data for K1OwnerAdded events raised by the ISsoAccount contract.
type ISsoAccountK1OwnerAddedIterator struct {
	Event *ISsoAccountK1OwnerAdded // Event containing the contract specifics and raw log

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
func (it *ISsoAccountK1OwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISsoAccountK1OwnerAdded)
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
		it.Event = new(ISsoAccountK1OwnerAdded)
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
func (it *ISsoAccountK1OwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISsoAccountK1OwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISsoAccountK1OwnerAdded represents a K1OwnerAdded event raised by the ISsoAccount contract.
type ISsoAccountK1OwnerAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerAdded is a free log retrieval operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_ISsoAccount *ISsoAccountFilterer) FilterK1OwnerAdded(opts *bind.FilterOpts, addr []common.Address) (*ISsoAccountK1OwnerAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ISsoAccount.contract.FilterLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountK1OwnerAddedIterator{contract: _ISsoAccount.contract, event: "K1OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchK1OwnerAdded is a free log subscription operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_ISsoAccount *ISsoAccountFilterer) WatchK1OwnerAdded(opts *bind.WatchOpts, sink chan<- *ISsoAccountK1OwnerAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ISsoAccount.contract.WatchLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISsoAccountK1OwnerAdded)
				if err := _ISsoAccount.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
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
func (_ISsoAccount *ISsoAccountFilterer) ParseK1OwnerAdded(log types.Log) (*ISsoAccountK1OwnerAdded, error) {
	event := new(ISsoAccountK1OwnerAdded)
	if err := _ISsoAccount.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISsoAccountK1OwnerRemovedIterator is returned from FilterK1OwnerRemoved and is used to iterate over the raw logs and unpacked data for K1OwnerRemoved events raised by the ISsoAccount contract.
type ISsoAccountK1OwnerRemovedIterator struct {
	Event *ISsoAccountK1OwnerRemoved // Event containing the contract specifics and raw log

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
func (it *ISsoAccountK1OwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISsoAccountK1OwnerRemoved)
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
		it.Event = new(ISsoAccountK1OwnerRemoved)
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
func (it *ISsoAccountK1OwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISsoAccountK1OwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISsoAccountK1OwnerRemoved represents a K1OwnerRemoved event raised by the ISsoAccount contract.
type ISsoAccountK1OwnerRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerRemoved is a free log retrieval operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_ISsoAccount *ISsoAccountFilterer) FilterK1OwnerRemoved(opts *bind.FilterOpts, addr []common.Address) (*ISsoAccountK1OwnerRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ISsoAccount.contract.FilterLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountK1OwnerRemovedIterator{contract: _ISsoAccount.contract, event: "K1OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchK1OwnerRemoved is a free log subscription operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_ISsoAccount *ISsoAccountFilterer) WatchK1OwnerRemoved(opts *bind.WatchOpts, sink chan<- *ISsoAccountK1OwnerRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ISsoAccount.contract.WatchLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISsoAccountK1OwnerRemoved)
				if err := _ISsoAccount.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
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
func (_ISsoAccount *ISsoAccountFilterer) ParseK1OwnerRemoved(log types.Log) (*ISsoAccountK1OwnerRemoved, error) {
	event := new(ISsoAccountK1OwnerRemoved)
	if err := _ISsoAccount.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISsoAccountValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ISsoAccount contract.
type ISsoAccountValidatorAddedIterator struct {
	Event *ISsoAccountValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ISsoAccountValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISsoAccountValidatorAdded)
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
		it.Event = new(ISsoAccountValidatorAdded)
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
func (it *ISsoAccountValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISsoAccountValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISsoAccountValidatorAdded represents a ValidatorAdded event raised by the ISsoAccount contract.
type ISsoAccountValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ISsoAccount *ISsoAccountFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*ISsoAccountValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ISsoAccount.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountValidatorAddedIterator{contract: _ISsoAccount.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ISsoAccount *ISsoAccountFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ISsoAccountValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ISsoAccount.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISsoAccountValidatorAdded)
				if err := _ISsoAccount.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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
func (_ISsoAccount *ISsoAccountFilterer) ParseValidatorAdded(log types.Log) (*ISsoAccountValidatorAdded, error) {
	event := new(ISsoAccountValidatorAdded)
	if err := _ISsoAccount.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISsoAccountValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the ISsoAccount contract.
type ISsoAccountValidatorRemovedIterator struct {
	Event *ISsoAccountValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ISsoAccountValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISsoAccountValidatorRemoved)
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
		it.Event = new(ISsoAccountValidatorRemoved)
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
func (it *ISsoAccountValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISsoAccountValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISsoAccountValidatorRemoved represents a ValidatorRemoved event raised by the ISsoAccount contract.
type ISsoAccountValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ISsoAccount *ISsoAccountFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*ISsoAccountValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ISsoAccount.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ISsoAccountValidatorRemovedIterator{contract: _ISsoAccount.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ISsoAccount *ISsoAccountFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ISsoAccountValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ISsoAccount.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISsoAccountValidatorRemoved)
				if err := _ISsoAccount.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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
func (_ISsoAccount *ISsoAccountFilterer) ParseValidatorRemoved(log types.Log) (*ISsoAccountValidatorRemoved, error) {
	event := new(ISsoAccountValidatorRemoved)
	if err := _ISsoAccount.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
