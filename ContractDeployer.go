package zksync2

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zksync-sdk/zksync2-go/contracts/ContractDeployer"
	"math/big"
	"strings"
)

var contractDeployerABI *abi.ABI

func getContractDeployerABI() (*abi.ABI, error) {
	if contractDeployerABI == nil {
		cda, err := abi.JSON(strings.NewReader(ContractDeployer.ContractDeployerMetaData.ABI))
		if err != nil {
			return nil, fmt.Errorf("failed to load ContractDeployer ABI: %w", err)
		}
		contractDeployerABI = &cda
	}
	return contractDeployerABI, nil
}

func ComputeL2Create2Address(sender common.Address, bytecode, constructor, salt []byte) (common.Address, error) {
	if len(salt) == 0 {
		salt = make([]byte, 32)
	} else if len(salt) != 32 {
		return common.Address{}, errors.New("salt must be 32 bytes")
	}
	if constructor == nil {
		constructor = []byte{}
	}
	bytecodeHash, err := HashBytecode(bytecode)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	bytes := make([]byte, 5*32)                                  // concatenate five 32-bytes slices
	copy(bytes[0:32], crypto.Keccak256([]byte("zksyncCreate2"))) // 1 - prefix
	copy(bytes[44:64], sender[:])                                // 2 - sender (20 bytes right padded to 32)
	copy(bytes[64:96], salt)                                     // 3 - salt
	copy(bytes[96:128], bytecodeHash)                            // 4 - bytecode hash
	copy(bytes[128:160], crypto.Keccak256(constructor))          // 5 - constructor hash
	result := crypto.Keccak256(bytes)
	return common.BytesToAddress(result[12:]), nil
}

func ComputeL2CreateAddress(sender common.Address, nonce *big.Int) (common.Address, error) {
	nonceBytes := nonce.Bytes()
	bytes := make([]byte, 3*32)                                 // concatenate three 32-bytes slices
	copy(bytes[0:32], crypto.Keccak256([]byte("zksyncCreate"))) // 1 - prefix
	copy(bytes[44:64], sender[:])                               // 2 - sender (20 bytes right padded to 32)
	copy(bytes[64+(32-len(nonceBytes)):96], nonceBytes)         // 3 - nonce (some bytes right padded to 32)
	result := crypto.Keccak256(bytes)
	return common.BytesToAddress(result[12:]), nil
}

func EncodeCreate2(bytecode, calldata, salt []byte) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	// prepare
	if len(salt) == 0 {
		salt = make([]byte, 32)
	} else if len(salt) != 32 {
		return nil, errors.New("salt must be 32 bytes")
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{}
	copy(salt32[:], salt)
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("create2", salt32, hash32, calldata)
	if err != nil {
		return nil, fmt.Errorf("failed to pack create2 function: %w", err)
	}
	return res, nil
}

func EncodeCreate(bytecode, calldata []byte) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{} // will be empty
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("create", salt32, hash32, calldata)
	if err != nil {
		return nil, fmt.Errorf("failed to pack create function: %w", err)
	}
	return res, nil
}

func EncodeCreate2Account(bytecode, calldata, salt []byte, version AccountAbstractionVersion) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	// prepare
	if len(salt) == 0 {
		salt = make([]byte, 32)
	} else if len(salt) != 32 {
		return nil, errors.New("salt must be 32 bytes")
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{}
	copy(salt32[:], salt)
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("create2Account", salt32, hash32, calldata, version)
	if err != nil {
		return nil, fmt.Errorf("failed to pack create2Account function: %w", err)
	}
	return res, nil
}

func EncodeCreateAccount(bytecode, calldata []byte, version AccountAbstractionVersion) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{} // will be empty
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("createAccount", salt32, hash32, calldata, version)
	if err != nil {
		return nil, fmt.Errorf("failed to pack createAccount function: %w", err)
	}
	return res, nil
}

func HashBytecode(bytecode []byte) ([]byte, error) {
	if len(bytecode)%32 != 0 {
		return nil, errors.New("bytecode length in bytes must be divisible by 32")
	}
	bytecodeHash := sha256.Sum256(bytecode)
	// get real length of bytecode, which is presented as 32-byte words
	length := big.NewInt(int64(len(bytecode) / 32))
	if length.BitLen() > 16 {
		return nil, errors.New("bytecode length must be less than 2^16 bytes")
	}
	// replace first 2 bytes of hash with version
	version := []byte{1, 0}
	copy(bytecodeHash[0:2], version)
	// replace second 2 bytes of hash with bytecode length
	length2b := make([]byte, 2)
	length2b = length.FillBytes(length2b) // 0-padded in 2 bytes
	copy(bytecodeHash[2:4], length2b)
	return bytecodeHash[:], nil
}
