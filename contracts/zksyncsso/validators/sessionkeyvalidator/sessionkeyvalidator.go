// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sessionkeyvalidator

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

// SessionLibCallSpec is an auto generated low-level Go binding around an user-defined struct.
type SessionLibCallSpec struct {
	Target         common.Address
	Selector       [4]byte
	MaxValuePerUse *big.Int
	ValueLimit     SessionLibUsageLimit
	Constraints    []SessionLibConstraint
}

// SessionLibConstraint is an auto generated low-level Go binding around an user-defined struct.
type SessionLibConstraint struct {
	Condition uint8
	Index     uint64
	RefValue  [32]byte
	Limit     SessionLibUsageLimit
}

// SessionLibLimitState is an auto generated low-level Go binding around an user-defined struct.
type SessionLibLimitState struct {
	Remaining *big.Int
	Target    common.Address
	Selector  [4]byte
	Index     *big.Int
}

// SessionLibSessionSpec is an auto generated low-level Go binding around an user-defined struct.
type SessionLibSessionSpec struct {
	Signer           common.Address
	ExpiresAt        *big.Int
	FeeLimit         SessionLibUsageLimit
	CallPolicies     []SessionLibCallSpec
	TransferPolicies []SessionLibTransferSpec
}

// SessionLibSessionState is an auto generated low-level Go binding around an user-defined struct.
type SessionLibSessionState struct {
	Status        uint8
	FeesRemaining *big.Int
	TransferValue []SessionLibLimitState
	CallValue     []SessionLibLimitState
	CallParams    []SessionLibLimitState
}

// SessionLibTransferSpec is an auto generated low-level Go binding around an user-defined struct.
type SessionLibTransferSpec struct {
	Target         common.Address
	MaxValuePerUse *big.Int
	ValueLimit     SessionLibUsageLimit
}

// SessionLibUsageLimit is an auto generated low-level Go binding around an user-defined struct.
type SessionLibUsageLimit struct {
	LimitType uint8
	Limit     *big.Int
	Period    *big.Int
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

// SessionKeyValidatorMetaData contains all meta data concerning the SessionKeyValidator contract.
var SessionKeyValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ADDRESS_CAST_OVERFLOW\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"INVALID_PAYMASTER_INPUT\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notInitialized\",\"type\":\"address\"}],\"name\":\"NOT_FROM_INITIALIZED_ACCOUNT\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"NO_TIMESTAMP_ASSERTER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"period\",\"type\":\"uint64\"}],\"name\":\"SESSION_ALLOWANCE_EXCEEDED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sessionHash\",\"type\":\"bytes32\"}],\"name\":\"SESSION_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"SESSION_CALL_POLICY_VIOLATED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"param\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"refValue\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"condition\",\"type\":\"uint8\"}],\"name\":\"SESSION_CONDITION_FAILED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"name\":\"SESSION_EXPIRES_TOO_SOON\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"}],\"name\":\"SESSION_INVALID_DATA_LENGTH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recovered\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"SESSION_INVALID_SIGNER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lifetimeUsage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxUsage\",\"type\":\"uint256\"}],\"name\":\"SESSION_LIFETIME_USAGE_EXCEEDED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usedValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValuePerUse\",\"type\":\"uint256\"}],\"name\":\"SESSION_MAX_VALUE_EXCEEDED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SESSION_NOT_ACTIVE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"SESSION_TRANSFER_POLICY_VIOLATED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SESSION_UNLIMITED_FEES\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SESSION_ZERO_SIGNER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"openSessions\",\"type\":\"uint256\"}],\"name\":\"UNINSTALL_WITH_OPEN_SESSIONS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sessionHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"feeLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"maxValuePerUse\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"valueLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumSessionLib.Condition\",\"name\":\"condition\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refValue\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"limit\",\"type\":\"tuple\"}],\"internalType\":\"structSessionLib.Constraint[]\",\"name\":\"constraints\",\"type\":\"tuple[]\"}],\"internalType\":\"structSessionLib.CallSpec[]\",\"name\":\"callPolicies\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxValuePerUse\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"valueLimit\",\"type\":\"tuple\"}],\"internalType\":\"structSessionLib.TransferSpec[]\",\"name\":\"transferPolicies\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structSessionLib.SessionSpec\",\"name\":\"sessionSpec\",\"type\":\"tuple\"}],\"name\":\"SessionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sessionHash\",\"type\":\"bytes32\"}],\"name\":\"SessionRevoked\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"feeLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"maxValuePerUse\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"valueLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumSessionLib.Condition\",\"name\":\"condition\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refValue\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"limit\",\"type\":\"tuple\"}],\"internalType\":\"structSessionLib.Constraint[]\",\"name\":\"constraints\",\"type\":\"tuple[]\"}],\"internalType\":\"structSessionLib.CallSpec[]\",\"name\":\"callPolicies\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxValuePerUse\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"valueLimit\",\"type\":\"tuple\"}],\"internalType\":\"structSessionLib.TransferSpec[]\",\"name\":\"transferPolicies\",\"type\":\"tuple[]\"}],\"internalType\":\"structSessionLib.SessionSpec\",\"name\":\"sessionSpec\",\"type\":\"tuple\"}],\"name\":\"createSession\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"smartAccount\",\"type\":\"address\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sessionHash\",\"type\":\"bytes32\"}],\"name\":\"revokeKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"sessionHashes\",\"type\":\"bytes32[]\"}],\"name\":\"revokeKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"feeLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"maxValuePerUse\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"valueLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumSessionLib.Condition\",\"name\":\"condition\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refValue\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"limit\",\"type\":\"tuple\"}],\"internalType\":\"structSessionLib.Constraint[]\",\"name\":\"constraints\",\"type\":\"tuple[]\"}],\"internalType\":\"structSessionLib.CallSpec[]\",\"name\":\"callPolicies\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxValuePerUse\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSessionLib.LimitType\",\"name\":\"limitType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.UsageLimit\",\"name\":\"valueLimit\",\"type\":\"tuple\"}],\"internalType\":\"structSessionLib.TransferSpec[]\",\"name\":\"transferPolicies\",\"type\":\"tuple[]\"}],\"internalType\":\"structSessionLib.SessionSpec\",\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"sessionState\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSessionLib.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"feesRemaining\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.LimitState[]\",\"name\":\"transferValue\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.LimitState[]\",\"name\":\"callValue\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structSessionLib.LimitState[]\",\"name\":\"callParams\",\"type\":\"tuple[]\"}],\"internalType\":\"structSessionLib.SessionState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sessionHash\",\"type\":\"bytes32\"}],\"name\":\"sessionStatus\",\"outputs\":[{\"internalType\":\"enumSessionLib.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SessionKeyValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SessionKeyValidatorMetaData.ABI instead.
var SessionKeyValidatorABI = SessionKeyValidatorMetaData.ABI

// SessionKeyValidator is an auto generated Go binding around an Ethereum contract.
type SessionKeyValidator struct {
	SessionKeyValidatorCaller     // Read-only binding to the contract
	SessionKeyValidatorTransactor // Write-only binding to the contract
	SessionKeyValidatorFilterer   // Log filterer for contract events
}

// SessionKeyValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SessionKeyValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SessionKeyValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SessionKeyValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionKeyValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SessionKeyValidatorSession struct {
	Contract     *SessionKeyValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SessionKeyValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SessionKeyValidatorCallerSession struct {
	Contract *SessionKeyValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SessionKeyValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SessionKeyValidatorTransactorSession struct {
	Contract     *SessionKeyValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SessionKeyValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SessionKeyValidatorRaw struct {
	Contract *SessionKeyValidator // Generic contract binding to access the raw methods on
}

// SessionKeyValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SessionKeyValidatorCallerRaw struct {
	Contract *SessionKeyValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// SessionKeyValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SessionKeyValidatorTransactorRaw struct {
	Contract *SessionKeyValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSessionKeyValidator creates a new instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidator(address common.Address, backend bind.ContractBackend) (*SessionKeyValidator, error) {
	contract, err := bindSessionKeyValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidator{SessionKeyValidatorCaller: SessionKeyValidatorCaller{contract: contract}, SessionKeyValidatorTransactor: SessionKeyValidatorTransactor{contract: contract}, SessionKeyValidatorFilterer: SessionKeyValidatorFilterer{contract: contract}}, nil
}

// NewSessionKeyValidatorCaller creates a new read-only instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidatorCaller(address common.Address, caller bind.ContractCaller) (*SessionKeyValidatorCaller, error) {
	contract, err := bindSessionKeyValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorCaller{contract: contract}, nil
}

// NewSessionKeyValidatorTransactor creates a new write-only instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SessionKeyValidatorTransactor, error) {
	contract, err := bindSessionKeyValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorTransactor{contract: contract}, nil
}

// NewSessionKeyValidatorFilterer creates a new log filterer instance of SessionKeyValidator, bound to a specific deployed contract.
func NewSessionKeyValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SessionKeyValidatorFilterer, error) {
	contract, err := bindSessionKeyValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorFilterer{contract: contract}, nil
}

// bindSessionKeyValidator binds a generic wrapper to an already deployed contract.
func bindSessionKeyValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SessionKeyValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionKeyValidator *SessionKeyValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionKeyValidator.Contract.SessionKeyValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionKeyValidator *SessionKeyValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.SessionKeyValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionKeyValidator *SessionKeyValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.SessionKeyValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionKeyValidator *SessionKeyValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionKeyValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionKeyValidator *SessionKeyValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionKeyValidator *SessionKeyValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.contract.Transact(opts, method, params...)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address smartAccount) view returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCaller) IsInitialized(opts *bind.CallOpts, smartAccount common.Address) (bool, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "isInitialized", smartAccount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address smartAccount) view returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorSession) IsInitialized(smartAccount common.Address) (bool, error) {
	return _SessionKeyValidator.Contract.IsInitialized(&_SessionKeyValidator.CallOpts, smartAccount)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address smartAccount) view returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) IsInitialized(smartAccount common.Address) (bool, error) {
	return _SessionKeyValidator.Contract.IsInitialized(&_SessionKeyValidator.CallOpts, smartAccount)
}

// SessionState is a free data retrieval call binding the contract method 0xb6d75019.
//
// Solidity: function sessionState(address account, (address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) spec) view returns((uint8,uint256,(uint256,address,bytes4,uint256)[],(uint256,address,bytes4,uint256)[],(uint256,address,bytes4,uint256)[]))
func (_SessionKeyValidator *SessionKeyValidatorCaller) SessionState(opts *bind.CallOpts, account common.Address, spec SessionLibSessionSpec) (SessionLibSessionState, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "sessionState", account, spec)

	if err != nil {
		return *new(SessionLibSessionState), err
	}

	out0 := *abi.ConvertType(out[0], new(SessionLibSessionState)).(*SessionLibSessionState)

	return out0, err

}

// SessionState is a free data retrieval call binding the contract method 0xb6d75019.
//
// Solidity: function sessionState(address account, (address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) spec) view returns((uint8,uint256,(uint256,address,bytes4,uint256)[],(uint256,address,bytes4,uint256)[],(uint256,address,bytes4,uint256)[]))
func (_SessionKeyValidator *SessionKeyValidatorSession) SessionState(account common.Address, spec SessionLibSessionSpec) (SessionLibSessionState, error) {
	return _SessionKeyValidator.Contract.SessionState(&_SessionKeyValidator.CallOpts, account, spec)
}

// SessionState is a free data retrieval call binding the contract method 0xb6d75019.
//
// Solidity: function sessionState(address account, (address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) spec) view returns((uint8,uint256,(uint256,address,bytes4,uint256)[],(uint256,address,bytes4,uint256)[],(uint256,address,bytes4,uint256)[]))
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) SessionState(account common.Address, spec SessionLibSessionSpec) (SessionLibSessionState, error) {
	return _SessionKeyValidator.Contract.SessionState(&_SessionKeyValidator.CallOpts, account, spec)
}

// SessionStatus is a free data retrieval call binding the contract method 0x59130708.
//
// Solidity: function sessionStatus(address account, bytes32 sessionHash) view returns(uint8)
func (_SessionKeyValidator *SessionKeyValidatorCaller) SessionStatus(opts *bind.CallOpts, account common.Address, sessionHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "sessionStatus", account, sessionHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SessionStatus is a free data retrieval call binding the contract method 0x59130708.
//
// Solidity: function sessionStatus(address account, bytes32 sessionHash) view returns(uint8)
func (_SessionKeyValidator *SessionKeyValidatorSession) SessionStatus(account common.Address, sessionHash [32]byte) (uint8, error) {
	return _SessionKeyValidator.Contract.SessionStatus(&_SessionKeyValidator.CallOpts, account, sessionHash)
}

// SessionStatus is a free data retrieval call binding the contract method 0x59130708.
//
// Solidity: function sessionStatus(address account, bytes32 sessionHash) view returns(uint8)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) SessionStatus(account common.Address, sessionHash [32]byte) (uint8, error) {
	return _SessionKeyValidator.Contract.SessionStatus(&_SessionKeyValidator.CallOpts, account, sessionHash)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SessionKeyValidator.Contract.SupportsInterface(&_SessionKeyValidator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SessionKeyValidator.Contract.SupportsInterface(&_SessionKeyValidator.CallOpts, interfaceId)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCaller) ValidateSignature(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _SessionKeyValidator.contract.Call(opts, &out, "validateSignature", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorSession) ValidateSignature(arg0 [32]byte, arg1 []byte) (bool, error) {
	return _SessionKeyValidator.Contract.ValidateSignature(&_SessionKeyValidator.CallOpts, arg0, arg1)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorCallerSession) ValidateSignature(arg0 [32]byte, arg1 []byte) (bool, error) {
	return _SessionKeyValidator.Contract.ValidateSignature(&_SessionKeyValidator.CallOpts, arg0, arg1)
}

// CreateSession is a paid mutator transaction binding the contract method 0x5a0694d2.
//
// Solidity: function createSession((address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) sessionSpec) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) CreateSession(opts *bind.TransactOpts, sessionSpec SessionLibSessionSpec) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "createSession", sessionSpec)
}

// CreateSession is a paid mutator transaction binding the contract method 0x5a0694d2.
//
// Solidity: function createSession((address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) sessionSpec) returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) CreateSession(sessionSpec SessionLibSessionSpec) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.CreateSession(&_SessionKeyValidator.TransactOpts, sessionSpec)
}

// CreateSession is a paid mutator transaction binding the contract method 0x5a0694d2.
//
// Solidity: function createSession((address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) sessionSpec) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) CreateSession(sessionSpec SessionLibSessionSpec) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.CreateSession(&_SessionKeyValidator.TransactOpts, sessionSpec)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.OnInstall(&_SessionKeyValidator.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.OnInstall(&_SessionKeyValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.OnUninstall(&_SessionKeyValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.OnUninstall(&_SessionKeyValidator.TransactOpts, data)
}

// RevokeKey is a paid mutator transaction binding the contract method 0x572f2210.
//
// Solidity: function revokeKey(bytes32 sessionHash) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) RevokeKey(opts *bind.TransactOpts, sessionHash [32]byte) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "revokeKey", sessionHash)
}

// RevokeKey is a paid mutator transaction binding the contract method 0x572f2210.
//
// Solidity: function revokeKey(bytes32 sessionHash) returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) RevokeKey(sessionHash [32]byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.RevokeKey(&_SessionKeyValidator.TransactOpts, sessionHash)
}

// RevokeKey is a paid mutator transaction binding the contract method 0x572f2210.
//
// Solidity: function revokeKey(bytes32 sessionHash) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) RevokeKey(sessionHash [32]byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.RevokeKey(&_SessionKeyValidator.TransactOpts, sessionHash)
}

// RevokeKeys is a paid mutator transaction binding the contract method 0x3ec59279.
//
// Solidity: function revokeKeys(bytes32[] sessionHashes) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactor) RevokeKeys(opts *bind.TransactOpts, sessionHashes [][32]byte) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "revokeKeys", sessionHashes)
}

// RevokeKeys is a paid mutator transaction binding the contract method 0x3ec59279.
//
// Solidity: function revokeKeys(bytes32[] sessionHashes) returns()
func (_SessionKeyValidator *SessionKeyValidatorSession) RevokeKeys(sessionHashes [][32]byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.RevokeKeys(&_SessionKeyValidator.TransactOpts, sessionHashes)
}

// RevokeKeys is a paid mutator transaction binding the contract method 0x3ec59279.
//
// Solidity: function revokeKeys(bytes32[] sessionHashes) returns()
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) RevokeKeys(sessionHashes [][32]byte) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.RevokeKeys(&_SessionKeyValidator.TransactOpts, sessionHashes)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorTransactor) ValidateTransaction(opts *bind.TransactOpts, signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _SessionKeyValidator.contract.Transact(opts, "validateTransaction", signedHash, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.ValidateTransaction(&_SessionKeyValidator.TransactOpts, signedHash, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_SessionKeyValidator *SessionKeyValidatorTransactorSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _SessionKeyValidator.Contract.ValidateTransaction(&_SessionKeyValidator.TransactOpts, signedHash, transaction)
}

// SessionKeyValidatorSessionCreatedIterator is returned from FilterSessionCreated and is used to iterate over the raw logs and unpacked data for SessionCreated events raised by the SessionKeyValidator contract.
type SessionKeyValidatorSessionCreatedIterator struct {
	Event *SessionKeyValidatorSessionCreated // Event containing the contract specifics and raw log

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
func (it *SessionKeyValidatorSessionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyValidatorSessionCreated)
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
		it.Event = new(SessionKeyValidatorSessionCreated)
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
func (it *SessionKeyValidatorSessionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyValidatorSessionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyValidatorSessionCreated represents a SessionCreated event raised by the SessionKeyValidator contract.
type SessionKeyValidatorSessionCreated struct {
	Account     common.Address
	SessionHash [32]byte
	SessionSpec SessionLibSessionSpec
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSessionCreated is a free log retrieval operation binding the contract event 0x225b7c57a9711566a3a2dabe5ba5d2c5965272461cd8af89883376ecb2f69ac4.
//
// Solidity: event SessionCreated(address indexed account, bytes32 indexed sessionHash, (address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) sessionSpec)
func (_SessionKeyValidator *SessionKeyValidatorFilterer) FilterSessionCreated(opts *bind.FilterOpts, account []common.Address, sessionHash [][32]byte) (*SessionKeyValidatorSessionCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var sessionHashRule []interface{}
	for _, sessionHashItem := range sessionHash {
		sessionHashRule = append(sessionHashRule, sessionHashItem)
	}

	logs, sub, err := _SessionKeyValidator.contract.FilterLogs(opts, "SessionCreated", accountRule, sessionHashRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorSessionCreatedIterator{contract: _SessionKeyValidator.contract, event: "SessionCreated", logs: logs, sub: sub}, nil
}

// WatchSessionCreated is a free log subscription operation binding the contract event 0x225b7c57a9711566a3a2dabe5ba5d2c5965272461cd8af89883376ecb2f69ac4.
//
// Solidity: event SessionCreated(address indexed account, bytes32 indexed sessionHash, (address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) sessionSpec)
func (_SessionKeyValidator *SessionKeyValidatorFilterer) WatchSessionCreated(opts *bind.WatchOpts, sink chan<- *SessionKeyValidatorSessionCreated, account []common.Address, sessionHash [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var sessionHashRule []interface{}
	for _, sessionHashItem := range sessionHash {
		sessionHashRule = append(sessionHashRule, sessionHashItem)
	}

	logs, sub, err := _SessionKeyValidator.contract.WatchLogs(opts, "SessionCreated", accountRule, sessionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyValidatorSessionCreated)
				if err := _SessionKeyValidator.contract.UnpackLog(event, "SessionCreated", log); err != nil {
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

// ParseSessionCreated is a log parse operation binding the contract event 0x225b7c57a9711566a3a2dabe5ba5d2c5965272461cd8af89883376ecb2f69ac4.
//
// Solidity: event SessionCreated(address indexed account, bytes32 indexed sessionHash, (address,uint256,(uint8,uint256,uint256),(address,bytes4,uint256,(uint8,uint256,uint256),(uint8,uint64,bytes32,(uint8,uint256,uint256))[])[],(address,uint256,(uint8,uint256,uint256))[]) sessionSpec)
func (_SessionKeyValidator *SessionKeyValidatorFilterer) ParseSessionCreated(log types.Log) (*SessionKeyValidatorSessionCreated, error) {
	event := new(SessionKeyValidatorSessionCreated)
	if err := _SessionKeyValidator.contract.UnpackLog(event, "SessionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionKeyValidatorSessionRevokedIterator is returned from FilterSessionRevoked and is used to iterate over the raw logs and unpacked data for SessionRevoked events raised by the SessionKeyValidator contract.
type SessionKeyValidatorSessionRevokedIterator struct {
	Event *SessionKeyValidatorSessionRevoked // Event containing the contract specifics and raw log

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
func (it *SessionKeyValidatorSessionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionKeyValidatorSessionRevoked)
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
		it.Event = new(SessionKeyValidatorSessionRevoked)
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
func (it *SessionKeyValidatorSessionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionKeyValidatorSessionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionKeyValidatorSessionRevoked represents a SessionRevoked event raised by the SessionKeyValidator contract.
type SessionKeyValidatorSessionRevoked struct {
	Account     common.Address
	SessionHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSessionRevoked is a free log retrieval operation binding the contract event 0x006be5f629d8df9f032986474cfc3bfc82c9b1cd74dcb2fcb5b3d48f98a50024.
//
// Solidity: event SessionRevoked(address indexed account, bytes32 indexed sessionHash)
func (_SessionKeyValidator *SessionKeyValidatorFilterer) FilterSessionRevoked(opts *bind.FilterOpts, account []common.Address, sessionHash [][32]byte) (*SessionKeyValidatorSessionRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var sessionHashRule []interface{}
	for _, sessionHashItem := range sessionHash {
		sessionHashRule = append(sessionHashRule, sessionHashItem)
	}

	logs, sub, err := _SessionKeyValidator.contract.FilterLogs(opts, "SessionRevoked", accountRule, sessionHashRule)
	if err != nil {
		return nil, err
	}
	return &SessionKeyValidatorSessionRevokedIterator{contract: _SessionKeyValidator.contract, event: "SessionRevoked", logs: logs, sub: sub}, nil
}

// WatchSessionRevoked is a free log subscription operation binding the contract event 0x006be5f629d8df9f032986474cfc3bfc82c9b1cd74dcb2fcb5b3d48f98a50024.
//
// Solidity: event SessionRevoked(address indexed account, bytes32 indexed sessionHash)
func (_SessionKeyValidator *SessionKeyValidatorFilterer) WatchSessionRevoked(opts *bind.WatchOpts, sink chan<- *SessionKeyValidatorSessionRevoked, account []common.Address, sessionHash [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var sessionHashRule []interface{}
	for _, sessionHashItem := range sessionHash {
		sessionHashRule = append(sessionHashRule, sessionHashItem)
	}

	logs, sub, err := _SessionKeyValidator.contract.WatchLogs(opts, "SessionRevoked", accountRule, sessionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionKeyValidatorSessionRevoked)
				if err := _SessionKeyValidator.contract.UnpackLog(event, "SessionRevoked", log); err != nil {
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

// ParseSessionRevoked is a log parse operation binding the contract event 0x006be5f629d8df9f032986474cfc3bfc82c9b1cd74dcb2fcb5b3d48f98a50024.
//
// Solidity: event SessionRevoked(address indexed account, bytes32 indexed sessionHash)
func (_SessionKeyValidator *SessionKeyValidatorFilterer) ParseSessionRevoked(log types.Log) (*SessionKeyValidatorSessionRevoked, error) {
	event := new(SessionKeyValidatorSessionRevoked)
	if err := _SessionKeyValidator.contract.UnpackLog(event, "SessionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
