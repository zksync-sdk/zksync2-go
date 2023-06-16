package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zksync-sdk/zksync2-go/types"
	"log"
	"os"
	"strings"
)

// ReadStandardJson reads standard-json file generated as output from zksolc.
// Returns standard json configuration and extracted contracts abi and bytecode from config file.
func ReadStandardJson(path string) (*types.StandardConfiguration, abi.ABI, []byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, abi.ABI{}, nil, fmt.Errorf("error reading standard json file: %w", err)
	}

	var config types.StandardConfiguration
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, abi.ABI{}, nil, fmt.Errorf("error decoding standard json file: %w", err)
	}

	bytecode, err := hexutil.Decode(config.Bytecode)
	if err != nil {
		return nil, abi.ABI{}, nil, fmt.Errorf("error decoding bytcode from standard json file: %w", err)
	}

	abiJson, err := json.Marshal(config.Abi)
	if err != nil {
		log.Panic(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		log.Panic(err)
	}

	return &config, contractAbi, bytecode, nil
}
