// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iguardianrecoveryvalidator

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

// IGuardianRecoveryValidatorGuardian is an auto generated low-level Go binding around an user-defined struct.
type IGuardianRecoveryValidatorGuardian struct {
	Addr    common.Address
	IsReady bool
	AddedAt uint64
}

// IGuardianRecoveryValidatorRecoveryRequest is an auto generated low-level Go binding around an user-defined struct.
type IGuardianRecoveryValidatorRecoveryRequest struct {
	HashedCredentialId [32]byte
	RawPublicKey       [2][32]byte
	Timestamp          *big.Int
}

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

// IGuardianRecoveryValidatorMetaData contains all meta data concerning the IGuardianRecoveryValidator contract.
var IGuardianRecoveryValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"AccountAlreadyGuardedByGuardian\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"AccountNotGuardedByAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountRecoveryInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GuardianCannotBeSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianNotProposed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccountToGuardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccountToRecoverAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGuardianAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWebAuthValidatorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonFunctionCallTransaction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"UnknownHashedOriginDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WebAuthValidatorNotEnabled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"HashedOriginDomainDisabledForAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"HashedOriginDomainEnabledForAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"}],\"name\":\"RecoveryDiscarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"}],\"name\":\"RecoveryFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"RecoveryInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"accountToGuard\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"discardRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPendingRecoveryData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"rawPublicKey\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIGuardianRecoveryValidator.RecoveryRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardianOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"guardiansFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isReady\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"addedAt\",\"type\":\"uint64\"}],\"internalType\":\"structIGuardianRecoveryValidator.Guardian[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"rawPublicKey\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"initRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"proposeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"guardianToRemove\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGuardianRecoveryValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IGuardianRecoveryValidatorMetaData.ABI instead.
var IGuardianRecoveryValidatorABI = IGuardianRecoveryValidatorMetaData.ABI

// IGuardianRecoveryValidator is an auto generated Go binding around an Ethereum contract.
type IGuardianRecoveryValidator struct {
	IGuardianRecoveryValidatorCaller     // Read-only binding to the contract
	IGuardianRecoveryValidatorTransactor // Write-only binding to the contract
	IGuardianRecoveryValidatorFilterer   // Log filterer for contract events
}

// IGuardianRecoveryValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGuardianRecoveryValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGuardianRecoveryValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGuardianRecoveryValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGuardianRecoveryValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGuardianRecoveryValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGuardianRecoveryValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGuardianRecoveryValidatorSession struct {
	Contract     *IGuardianRecoveryValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IGuardianRecoveryValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGuardianRecoveryValidatorCallerSession struct {
	Contract *IGuardianRecoveryValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IGuardianRecoveryValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGuardianRecoveryValidatorTransactorSession struct {
	Contract     *IGuardianRecoveryValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IGuardianRecoveryValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGuardianRecoveryValidatorRaw struct {
	Contract *IGuardianRecoveryValidator // Generic contract binding to access the raw methods on
}

// IGuardianRecoveryValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGuardianRecoveryValidatorCallerRaw struct {
	Contract *IGuardianRecoveryValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// IGuardianRecoveryValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGuardianRecoveryValidatorTransactorRaw struct {
	Contract *IGuardianRecoveryValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGuardianRecoveryValidator creates a new instance of IGuardianRecoveryValidator, bound to a specific deployed contract.
func NewIGuardianRecoveryValidator(address common.Address, backend bind.ContractBackend) (*IGuardianRecoveryValidator, error) {
	contract, err := bindIGuardianRecoveryValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidator{IGuardianRecoveryValidatorCaller: IGuardianRecoveryValidatorCaller{contract: contract}, IGuardianRecoveryValidatorTransactor: IGuardianRecoveryValidatorTransactor{contract: contract}, IGuardianRecoveryValidatorFilterer: IGuardianRecoveryValidatorFilterer{contract: contract}}, nil
}

// NewIGuardianRecoveryValidatorCaller creates a new read-only instance of IGuardianRecoveryValidator, bound to a specific deployed contract.
func NewIGuardianRecoveryValidatorCaller(address common.Address, caller bind.ContractCaller) (*IGuardianRecoveryValidatorCaller, error) {
	contract, err := bindIGuardianRecoveryValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorCaller{contract: contract}, nil
}

// NewIGuardianRecoveryValidatorTransactor creates a new write-only instance of IGuardianRecoveryValidator, bound to a specific deployed contract.
func NewIGuardianRecoveryValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IGuardianRecoveryValidatorTransactor, error) {
	contract, err := bindIGuardianRecoveryValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorTransactor{contract: contract}, nil
}

// NewIGuardianRecoveryValidatorFilterer creates a new log filterer instance of IGuardianRecoveryValidator, bound to a specific deployed contract.
func NewIGuardianRecoveryValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IGuardianRecoveryValidatorFilterer, error) {
	contract, err := bindIGuardianRecoveryValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorFilterer{contract: contract}, nil
}

// bindIGuardianRecoveryValidator binds a generic wrapper to an already deployed contract.
func bindIGuardianRecoveryValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGuardianRecoveryValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGuardianRecoveryValidator.Contract.IGuardianRecoveryValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.IGuardianRecoveryValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.IGuardianRecoveryValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGuardianRecoveryValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.contract.Transact(opts, method, params...)
}

// GetPendingRecoveryData is a free data retrieval call binding the contract method 0xab9d5fff.
//
// Solidity: function getPendingRecoveryData(bytes32 hashedOriginDomain, address account) view returns((bytes32,bytes32[2],uint256))
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCaller) GetPendingRecoveryData(opts *bind.CallOpts, hashedOriginDomain [32]byte, account common.Address) (IGuardianRecoveryValidatorRecoveryRequest, error) {
	var out []interface{}
	err := _IGuardianRecoveryValidator.contract.Call(opts, &out, "getPendingRecoveryData", hashedOriginDomain, account)

	if err != nil {
		return *new(IGuardianRecoveryValidatorRecoveryRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IGuardianRecoveryValidatorRecoveryRequest)).(*IGuardianRecoveryValidatorRecoveryRequest)

	return out0, err

}

// GetPendingRecoveryData is a free data retrieval call binding the contract method 0xab9d5fff.
//
// Solidity: function getPendingRecoveryData(bytes32 hashedOriginDomain, address account) view returns((bytes32,bytes32[2],uint256))
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) GetPendingRecoveryData(hashedOriginDomain [32]byte, account common.Address) (IGuardianRecoveryValidatorRecoveryRequest, error) {
	return _IGuardianRecoveryValidator.Contract.GetPendingRecoveryData(&_IGuardianRecoveryValidator.CallOpts, hashedOriginDomain, account)
}

// GetPendingRecoveryData is a free data retrieval call binding the contract method 0xab9d5fff.
//
// Solidity: function getPendingRecoveryData(bytes32 hashedOriginDomain, address account) view returns((bytes32,bytes32[2],uint256))
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCallerSession) GetPendingRecoveryData(hashedOriginDomain [32]byte, account common.Address) (IGuardianRecoveryValidatorRecoveryRequest, error) {
	return _IGuardianRecoveryValidator.Contract.GetPendingRecoveryData(&_IGuardianRecoveryValidator.CallOpts, hashedOriginDomain, account)
}

// GuardianOf is a free data retrieval call binding the contract method 0xcb339850.
//
// Solidity: function guardianOf(bytes32 hashedOriginDomain, address guardian) view returns(address[])
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCaller) GuardianOf(opts *bind.CallOpts, hashedOriginDomain [32]byte, guardian common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _IGuardianRecoveryValidator.contract.Call(opts, &out, "guardianOf", hashedOriginDomain, guardian)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GuardianOf is a free data retrieval call binding the contract method 0xcb339850.
//
// Solidity: function guardianOf(bytes32 hashedOriginDomain, address guardian) view returns(address[])
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) GuardianOf(hashedOriginDomain [32]byte, guardian common.Address) ([]common.Address, error) {
	return _IGuardianRecoveryValidator.Contract.GuardianOf(&_IGuardianRecoveryValidator.CallOpts, hashedOriginDomain, guardian)
}

// GuardianOf is a free data retrieval call binding the contract method 0xcb339850.
//
// Solidity: function guardianOf(bytes32 hashedOriginDomain, address guardian) view returns(address[])
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCallerSession) GuardianOf(hashedOriginDomain [32]byte, guardian common.Address) ([]common.Address, error) {
	return _IGuardianRecoveryValidator.Contract.GuardianOf(&_IGuardianRecoveryValidator.CallOpts, hashedOriginDomain, guardian)
}

// GuardiansFor is a free data retrieval call binding the contract method 0x511f723d.
//
// Solidity: function guardiansFor(bytes32 hashedOriginDomain, address addr) view returns((address,bool,uint64)[])
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCaller) GuardiansFor(opts *bind.CallOpts, hashedOriginDomain [32]byte, addr common.Address) ([]IGuardianRecoveryValidatorGuardian, error) {
	var out []interface{}
	err := _IGuardianRecoveryValidator.contract.Call(opts, &out, "guardiansFor", hashedOriginDomain, addr)

	if err != nil {
		return *new([]IGuardianRecoveryValidatorGuardian), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGuardianRecoveryValidatorGuardian)).(*[]IGuardianRecoveryValidatorGuardian)

	return out0, err

}

// GuardiansFor is a free data retrieval call binding the contract method 0x511f723d.
//
// Solidity: function guardiansFor(bytes32 hashedOriginDomain, address addr) view returns((address,bool,uint64)[])
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) GuardiansFor(hashedOriginDomain [32]byte, addr common.Address) ([]IGuardianRecoveryValidatorGuardian, error) {
	return _IGuardianRecoveryValidator.Contract.GuardiansFor(&_IGuardianRecoveryValidator.CallOpts, hashedOriginDomain, addr)
}

// GuardiansFor is a free data retrieval call binding the contract method 0x511f723d.
//
// Solidity: function guardiansFor(bytes32 hashedOriginDomain, address addr) view returns((address,bool,uint64)[])
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCallerSession) GuardiansFor(hashedOriginDomain [32]byte, addr common.Address) ([]IGuardianRecoveryValidatorGuardian, error) {
	return _IGuardianRecoveryValidator.Contract.GuardiansFor(&_IGuardianRecoveryValidator.CallOpts, hashedOriginDomain, addr)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IGuardianRecoveryValidator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IGuardianRecoveryValidator.Contract.SupportsInterface(&_IGuardianRecoveryValidator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IGuardianRecoveryValidator.Contract.SupportsInterface(&_IGuardianRecoveryValidator.CallOpts, interfaceId)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCaller) ValidateSignature(opts *bind.CallOpts, signedHash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _IGuardianRecoveryValidator.contract.Call(opts, &out, "validateSignature", signedHash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) ValidateSignature(signedHash [32]byte, signature []byte) (bool, error) {
	return _IGuardianRecoveryValidator.Contract.ValidateSignature(&_IGuardianRecoveryValidator.CallOpts, signedHash, signature)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorCallerSession) ValidateSignature(signedHash [32]byte, signature []byte) (bool, error) {
	return _IGuardianRecoveryValidator.Contract.ValidateSignature(&_IGuardianRecoveryValidator.CallOpts, signedHash, signature)
}

// AddGuardian is a paid mutator transaction binding the contract method 0x85431596.
//
// Solidity: function addGuardian(bytes32 hashedOriginDomain, address accountToGuard) returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) AddGuardian(opts *bind.TransactOpts, hashedOriginDomain [32]byte, accountToGuard common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "addGuardian", hashedOriginDomain, accountToGuard)
}

// AddGuardian is a paid mutator transaction binding the contract method 0x85431596.
//
// Solidity: function addGuardian(bytes32 hashedOriginDomain, address accountToGuard) returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) AddGuardian(hashedOriginDomain [32]byte, accountToGuard common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.AddGuardian(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain, accountToGuard)
}

// AddGuardian is a paid mutator transaction binding the contract method 0x85431596.
//
// Solidity: function addGuardian(bytes32 hashedOriginDomain, address accountToGuard) returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) AddGuardian(hashedOriginDomain [32]byte, accountToGuard common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.AddGuardian(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain, accountToGuard)
}

// DiscardRecovery is a paid mutator transaction binding the contract method 0x51bd1f69.
//
// Solidity: function discardRecovery(bytes32 hashedOriginDomain) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) DiscardRecovery(opts *bind.TransactOpts, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "discardRecovery", hashedOriginDomain)
}

// DiscardRecovery is a paid mutator transaction binding the contract method 0x51bd1f69.
//
// Solidity: function discardRecovery(bytes32 hashedOriginDomain) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) DiscardRecovery(hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.DiscardRecovery(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain)
}

// DiscardRecovery is a paid mutator transaction binding the contract method 0x51bd1f69.
//
// Solidity: function discardRecovery(bytes32 hashedOriginDomain) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) DiscardRecovery(hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.DiscardRecovery(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain)
}

// InitRecovery is a paid mutator transaction binding the contract method 0x236ad679.
//
// Solidity: function initRecovery(address accountToRecover, bytes32 hashedCredentialId, bytes32[2] rawPublicKey, bytes32 hashedOriginDomain) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) InitRecovery(opts *bind.TransactOpts, accountToRecover common.Address, hashedCredentialId [32]byte, rawPublicKey [2][32]byte, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "initRecovery", accountToRecover, hashedCredentialId, rawPublicKey, hashedOriginDomain)
}

// InitRecovery is a paid mutator transaction binding the contract method 0x236ad679.
//
// Solidity: function initRecovery(address accountToRecover, bytes32 hashedCredentialId, bytes32[2] rawPublicKey, bytes32 hashedOriginDomain) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) InitRecovery(accountToRecover common.Address, hashedCredentialId [32]byte, rawPublicKey [2][32]byte, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.InitRecovery(&_IGuardianRecoveryValidator.TransactOpts, accountToRecover, hashedCredentialId, rawPublicKey, hashedOriginDomain)
}

// InitRecovery is a paid mutator transaction binding the contract method 0x236ad679.
//
// Solidity: function initRecovery(address accountToRecover, bytes32 hashedCredentialId, bytes32[2] rawPublicKey, bytes32 hashedOriginDomain) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) InitRecovery(accountToRecover common.Address, hashedCredentialId [32]byte, rawPublicKey [2][32]byte, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.InitRecovery(&_IGuardianRecoveryValidator.TransactOpts, accountToRecover, hashedCredentialId, rawPublicKey, hashedOriginDomain)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.OnInstall(&_IGuardianRecoveryValidator.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.OnInstall(&_IGuardianRecoveryValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.OnUninstall(&_IGuardianRecoveryValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.OnUninstall(&_IGuardianRecoveryValidator.TransactOpts, data)
}

// ProposeGuardian is a paid mutator transaction binding the contract method 0x9f5d29c9.
//
// Solidity: function proposeGuardian(bytes32 hashedOriginDomain, address newGuardian) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) ProposeGuardian(opts *bind.TransactOpts, hashedOriginDomain [32]byte, newGuardian common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "proposeGuardian", hashedOriginDomain, newGuardian)
}

// ProposeGuardian is a paid mutator transaction binding the contract method 0x9f5d29c9.
//
// Solidity: function proposeGuardian(bytes32 hashedOriginDomain, address newGuardian) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) ProposeGuardian(hashedOriginDomain [32]byte, newGuardian common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.ProposeGuardian(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain, newGuardian)
}

// ProposeGuardian is a paid mutator transaction binding the contract method 0x9f5d29c9.
//
// Solidity: function proposeGuardian(bytes32 hashedOriginDomain, address newGuardian) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) ProposeGuardian(hashedOriginDomain [32]byte, newGuardian common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.ProposeGuardian(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain, newGuardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x3cab947d.
//
// Solidity: function removeGuardian(bytes32 hashedOriginDomain, address guardianToRemove) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) RemoveGuardian(opts *bind.TransactOpts, hashedOriginDomain [32]byte, guardianToRemove common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "removeGuardian", hashedOriginDomain, guardianToRemove)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x3cab947d.
//
// Solidity: function removeGuardian(bytes32 hashedOriginDomain, address guardianToRemove) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) RemoveGuardian(hashedOriginDomain [32]byte, guardianToRemove common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.RemoveGuardian(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain, guardianToRemove)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x3cab947d.
//
// Solidity: function removeGuardian(bytes32 hashedOriginDomain, address guardianToRemove) returns()
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) RemoveGuardian(hashedOriginDomain [32]byte, guardianToRemove common.Address) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.RemoveGuardian(&_IGuardianRecoveryValidator.TransactOpts, hashedOriginDomain, guardianToRemove)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactor) ValidateTransaction(opts *bind.TransactOpts, signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.contract.Transact(opts, "validateTransaction", signedHash, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.ValidateTransaction(&_IGuardianRecoveryValidator.TransactOpts, signedHash, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorTransactorSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IGuardianRecoveryValidator.Contract.ValidateTransaction(&_IGuardianRecoveryValidator.TransactOpts, signedHash, transaction)
}

// IGuardianRecoveryValidatorGuardianAddedIterator is returned from FilterGuardianAdded and is used to iterate over the raw logs and unpacked data for GuardianAdded events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorGuardianAddedIterator struct {
	Event *IGuardianRecoveryValidatorGuardianAdded // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorGuardianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorGuardianAdded)
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
		it.Event = new(IGuardianRecoveryValidatorGuardianAdded)
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
func (it *IGuardianRecoveryValidatorGuardianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorGuardianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorGuardianAdded represents a GuardianAdded event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorGuardianAdded struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdded is a free log retrieval operation binding the contract event 0x6aba2dccf4edc3319e24e776ffc5d7f8fa89456e0f15421abf316d28d427ed33.
//
// Solidity: event GuardianAdded(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterGuardianAdded(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (*IGuardianRecoveryValidatorGuardianAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "GuardianAdded", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorGuardianAddedIterator{contract: _IGuardianRecoveryValidator.contract, event: "GuardianAdded", logs: logs, sub: sub}, nil
}

// WatchGuardianAdded is a free log subscription operation binding the contract event 0x6aba2dccf4edc3319e24e776ffc5d7f8fa89456e0f15421abf316d28d427ed33.
//
// Solidity: event GuardianAdded(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchGuardianAdded(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorGuardianAdded, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "GuardianAdded", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorGuardianAdded)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
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

// ParseGuardianAdded is a log parse operation binding the contract event 0x6aba2dccf4edc3319e24e776ffc5d7f8fa89456e0f15421abf316d28d427ed33.
//
// Solidity: event GuardianAdded(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseGuardianAdded(log types.Log) (*IGuardianRecoveryValidatorGuardianAdded, error) {
	event := new(IGuardianRecoveryValidatorGuardianAdded)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGuardianRecoveryValidatorGuardianProposedIterator is returned from FilterGuardianProposed and is used to iterate over the raw logs and unpacked data for GuardianProposed events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorGuardianProposedIterator struct {
	Event *IGuardianRecoveryValidatorGuardianProposed // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorGuardianProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorGuardianProposed)
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
		it.Event = new(IGuardianRecoveryValidatorGuardianProposed)
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
func (it *IGuardianRecoveryValidatorGuardianProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorGuardianProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorGuardianProposed represents a GuardianProposed event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorGuardianProposed struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGuardianProposed is a free log retrieval operation binding the contract event 0x3deaf0877005efc7ea2c571de4f96ff33bf60c0b32e188f1fa32bab3da8b870e.
//
// Solidity: event GuardianProposed(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterGuardianProposed(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (*IGuardianRecoveryValidatorGuardianProposedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "GuardianProposed", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorGuardianProposedIterator{contract: _IGuardianRecoveryValidator.contract, event: "GuardianProposed", logs: logs, sub: sub}, nil
}

// WatchGuardianProposed is a free log subscription operation binding the contract event 0x3deaf0877005efc7ea2c571de4f96ff33bf60c0b32e188f1fa32bab3da8b870e.
//
// Solidity: event GuardianProposed(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchGuardianProposed(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorGuardianProposed, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "GuardianProposed", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorGuardianProposed)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "GuardianProposed", log); err != nil {
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

// ParseGuardianProposed is a log parse operation binding the contract event 0x3deaf0877005efc7ea2c571de4f96ff33bf60c0b32e188f1fa32bab3da8b870e.
//
// Solidity: event GuardianProposed(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseGuardianProposed(log types.Log) (*IGuardianRecoveryValidatorGuardianProposed, error) {
	event := new(IGuardianRecoveryValidatorGuardianProposed)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "GuardianProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGuardianRecoveryValidatorGuardianRemovedIterator is returned from FilterGuardianRemoved and is used to iterate over the raw logs and unpacked data for GuardianRemoved events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorGuardianRemovedIterator struct {
	Event *IGuardianRecoveryValidatorGuardianRemoved // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorGuardianRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorGuardianRemoved)
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
		it.Event = new(IGuardianRecoveryValidatorGuardianRemoved)
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
func (it *IGuardianRecoveryValidatorGuardianRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorGuardianRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorGuardianRemoved represents a GuardianRemoved event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorGuardianRemoved struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGuardianRemoved is a free log retrieval operation binding the contract event 0xb9c57fa4d0e004062979bfe03380ddabcf23ec66f36e471ee5b5ab67349485b2.
//
// Solidity: event GuardianRemoved(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterGuardianRemoved(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (*IGuardianRecoveryValidatorGuardianRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "GuardianRemoved", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorGuardianRemovedIterator{contract: _IGuardianRecoveryValidator.contract, event: "GuardianRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardianRemoved is a free log subscription operation binding the contract event 0xb9c57fa4d0e004062979bfe03380ddabcf23ec66f36e471ee5b5ab67349485b2.
//
// Solidity: event GuardianRemoved(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchGuardianRemoved(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorGuardianRemoved, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "GuardianRemoved", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorGuardianRemoved)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
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

// ParseGuardianRemoved is a log parse operation binding the contract event 0xb9c57fa4d0e004062979bfe03380ddabcf23ec66f36e471ee5b5ab67349485b2.
//
// Solidity: event GuardianRemoved(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseGuardianRemoved(log types.Log) (*IGuardianRecoveryValidatorGuardianRemoved, error) {
	event := new(IGuardianRecoveryValidatorGuardianRemoved)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator is returned from FilterHashedOriginDomainDisabledForAccount and is used to iterate over the raw logs and unpacked data for HashedOriginDomainDisabledForAccount events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator struct {
	Event *IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
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
		it.Event = new(IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
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
func (it *IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount represents a HashedOriginDomainDisabledForAccount event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterHashedOriginDomainDisabledForAccount is a free log retrieval operation binding the contract event 0x1373f4151802c1c04091d06c3f13c1948e40afad22bd00382d4702808d0320eb.
//
// Solidity: event HashedOriginDomainDisabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterHashedOriginDomainDisabledForAccount(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte) (*IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "HashedOriginDomainDisabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator{contract: _IGuardianRecoveryValidator.contract, event: "HashedOriginDomainDisabledForAccount", logs: logs, sub: sub}, nil
}

// WatchHashedOriginDomainDisabledForAccount is a free log subscription operation binding the contract event 0x1373f4151802c1c04091d06c3f13c1948e40afad22bd00382d4702808d0320eb.
//
// Solidity: event HashedOriginDomainDisabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchHashedOriginDomainDisabledForAccount(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount, account []common.Address, hashedOriginDomain [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "HashedOriginDomainDisabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainDisabledForAccount", log); err != nil {
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

// ParseHashedOriginDomainDisabledForAccount is a log parse operation binding the contract event 0x1373f4151802c1c04091d06c3f13c1948e40afad22bd00382d4702808d0320eb.
//
// Solidity: event HashedOriginDomainDisabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseHashedOriginDomainDisabledForAccount(log types.Log) (*IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount, error) {
	event := new(IGuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainDisabledForAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator is returned from FilterHashedOriginDomainEnabledForAccount and is used to iterate over the raw logs and unpacked data for HashedOriginDomainEnabledForAccount events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator struct {
	Event *IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
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
		it.Event = new(IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
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
func (it *IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount represents a HashedOriginDomainEnabledForAccount event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterHashedOriginDomainEnabledForAccount is a free log retrieval operation binding the contract event 0x49e308297bf7e3373130f54d11e9944d094c586954301b5c7da3b918195c7580.
//
// Solidity: event HashedOriginDomainEnabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterHashedOriginDomainEnabledForAccount(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte) (*IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "HashedOriginDomainEnabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator{contract: _IGuardianRecoveryValidator.contract, event: "HashedOriginDomainEnabledForAccount", logs: logs, sub: sub}, nil
}

// WatchHashedOriginDomainEnabledForAccount is a free log subscription operation binding the contract event 0x49e308297bf7e3373130f54d11e9944d094c586954301b5c7da3b918195c7580.
//
// Solidity: event HashedOriginDomainEnabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchHashedOriginDomainEnabledForAccount(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount, account []common.Address, hashedOriginDomain [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "HashedOriginDomainEnabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainEnabledForAccount", log); err != nil {
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

// ParseHashedOriginDomainEnabledForAccount is a log parse operation binding the contract event 0x49e308297bf7e3373130f54d11e9944d094c586954301b5c7da3b918195c7580.
//
// Solidity: event HashedOriginDomainEnabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseHashedOriginDomainEnabledForAccount(log types.Log) (*IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount, error) {
	event := new(IGuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainEnabledForAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGuardianRecoveryValidatorRecoveryDiscardedIterator is returned from FilterRecoveryDiscarded and is used to iterate over the raw logs and unpacked data for RecoveryDiscarded events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorRecoveryDiscardedIterator struct {
	Event *IGuardianRecoveryValidatorRecoveryDiscarded // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorRecoveryDiscardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorRecoveryDiscarded)
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
		it.Event = new(IGuardianRecoveryValidatorRecoveryDiscarded)
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
func (it *IGuardianRecoveryValidatorRecoveryDiscardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorRecoveryDiscardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorRecoveryDiscarded represents a RecoveryDiscarded event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorRecoveryDiscarded struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	HashedCredentialId [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryDiscarded is a free log retrieval operation binding the contract event 0xa4e8705d9ef0c51657f9c64b028c7199ff3cd3ee5d057d6e1088e64580853d13.
//
// Solidity: event RecoveryDiscarded(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterRecoveryDiscarded(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (*IGuardianRecoveryValidatorRecoveryDiscardedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "RecoveryDiscarded", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorRecoveryDiscardedIterator{contract: _IGuardianRecoveryValidator.contract, event: "RecoveryDiscarded", logs: logs, sub: sub}, nil
}

// WatchRecoveryDiscarded is a free log subscription operation binding the contract event 0xa4e8705d9ef0c51657f9c64b028c7199ff3cd3ee5d057d6e1088e64580853d13.
//
// Solidity: event RecoveryDiscarded(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchRecoveryDiscarded(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorRecoveryDiscarded, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "RecoveryDiscarded", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorRecoveryDiscarded)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryDiscarded", log); err != nil {
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

// ParseRecoveryDiscarded is a log parse operation binding the contract event 0xa4e8705d9ef0c51657f9c64b028c7199ff3cd3ee5d057d6e1088e64580853d13.
//
// Solidity: event RecoveryDiscarded(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseRecoveryDiscarded(log types.Log) (*IGuardianRecoveryValidatorRecoveryDiscarded, error) {
	event := new(IGuardianRecoveryValidatorRecoveryDiscarded)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryDiscarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGuardianRecoveryValidatorRecoveryFinishedIterator is returned from FilterRecoveryFinished and is used to iterate over the raw logs and unpacked data for RecoveryFinished events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorRecoveryFinishedIterator struct {
	Event *IGuardianRecoveryValidatorRecoveryFinished // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorRecoveryFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorRecoveryFinished)
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
		it.Event = new(IGuardianRecoveryValidatorRecoveryFinished)
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
func (it *IGuardianRecoveryValidatorRecoveryFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorRecoveryFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorRecoveryFinished represents a RecoveryFinished event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorRecoveryFinished struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	HashedCredentialId [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFinished is a free log retrieval operation binding the contract event 0x122f699d3998d7d483d87a65abbdba44780a412ca03b283c5c6fc45b6d7d94d2.
//
// Solidity: event RecoveryFinished(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterRecoveryFinished(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (*IGuardianRecoveryValidatorRecoveryFinishedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "RecoveryFinished", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorRecoveryFinishedIterator{contract: _IGuardianRecoveryValidator.contract, event: "RecoveryFinished", logs: logs, sub: sub}, nil
}

// WatchRecoveryFinished is a free log subscription operation binding the contract event 0x122f699d3998d7d483d87a65abbdba44780a412ca03b283c5c6fc45b6d7d94d2.
//
// Solidity: event RecoveryFinished(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchRecoveryFinished(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorRecoveryFinished, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "RecoveryFinished", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorRecoveryFinished)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryFinished", log); err != nil {
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

// ParseRecoveryFinished is a log parse operation binding the contract event 0x122f699d3998d7d483d87a65abbdba44780a412ca03b283c5c6fc45b6d7d94d2.
//
// Solidity: event RecoveryFinished(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseRecoveryFinished(log types.Log) (*IGuardianRecoveryValidatorRecoveryFinished, error) {
	event := new(IGuardianRecoveryValidatorRecoveryFinished)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGuardianRecoveryValidatorRecoveryInitiatedIterator is returned from FilterRecoveryInitiated and is used to iterate over the raw logs and unpacked data for RecoveryInitiated events raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorRecoveryInitiatedIterator struct {
	Event *IGuardianRecoveryValidatorRecoveryInitiated // Event containing the contract specifics and raw log

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
func (it *IGuardianRecoveryValidatorRecoveryInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGuardianRecoveryValidatorRecoveryInitiated)
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
		it.Event = new(IGuardianRecoveryValidatorRecoveryInitiated)
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
func (it *IGuardianRecoveryValidatorRecoveryInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGuardianRecoveryValidatorRecoveryInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGuardianRecoveryValidatorRecoveryInitiated represents a RecoveryInitiated event raised by the IGuardianRecoveryValidator contract.
type IGuardianRecoveryValidatorRecoveryInitiated struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	HashedCredentialId [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryInitiated is a free log retrieval operation binding the contract event 0x750203e1c6151c9136eca3be6c61d060ea236395244e97c72ee3735c131c3802.
//
// Solidity: event RecoveryInitiated(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId, address guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) FilterRecoveryInitiated(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (*IGuardianRecoveryValidatorRecoveryInitiatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.FilterLogs(opts, "RecoveryInitiated", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return &IGuardianRecoveryValidatorRecoveryInitiatedIterator{contract: _IGuardianRecoveryValidator.contract, event: "RecoveryInitiated", logs: logs, sub: sub}, nil
}

// WatchRecoveryInitiated is a free log subscription operation binding the contract event 0x750203e1c6151c9136eca3be6c61d060ea236395244e97c72ee3735c131c3802.
//
// Solidity: event RecoveryInitiated(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId, address guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) WatchRecoveryInitiated(opts *bind.WatchOpts, sink chan<- *IGuardianRecoveryValidatorRecoveryInitiated, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _IGuardianRecoveryValidator.contract.WatchLogs(opts, "RecoveryInitiated", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGuardianRecoveryValidatorRecoveryInitiated)
				if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryInitiated", log); err != nil {
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

// ParseRecoveryInitiated is a log parse operation binding the contract event 0x750203e1c6151c9136eca3be6c61d060ea236395244e97c72ee3735c131c3802.
//
// Solidity: event RecoveryInitiated(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId, address guardian)
func (_IGuardianRecoveryValidator *IGuardianRecoveryValidatorFilterer) ParseRecoveryInitiated(log types.Log) (*IGuardianRecoveryValidatorRecoveryInitiated, error) {
	event := new(IGuardianRecoveryValidatorRecoveryInitiated)
	if err := _IGuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
