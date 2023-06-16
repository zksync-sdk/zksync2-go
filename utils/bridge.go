package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
	"strings"
)

const bridgeDataAbiDefinition = `[
	{
		"constant": true,
		"inputs": [
			{
				"name": "name",
				"type": "string"
			},
			{
				"name": "symbol",
				"type": "string"
			},
			{
				"name": "decimals",
				"type": "uint256"
			}
		],
		"name": "token",
		"outputs": [
				{
					"name": "",
					"type": "bool"
				}
			],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`

func CreateETH() *types.Token {
	return &types.Token{
		L1Address: common.Address{},
		L2Address: common.Address{},
		Name:      `ETH`,
		Symbol:    `ETH`,
		Decimals:  18,
	}
}

func GetERC20DefaultBridgeData(l1TokenAddress common.Address, backend bind.ContractBackend) ([]byte, error) {
	token, err := erc20.NewIERC20(l1TokenAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("failed to load IERC20: %w", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		return nil, err
	}
	symbol, err := token.Symbol(nil)
	if err != nil {
		return nil, err
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		return nil, err
	}

	bridgeDataAbi, err := abi.JSON(strings.NewReader(bridgeDataAbiDefinition))
	if err != nil {
		return nil, fmt.Errorf("failed to load bridge data ABI: %w", err)
	}
	// encoding constructor function returns encoded arguments which is required to be result of function
	// encoding different functions appends at beginning 4 bytes which is  hash of function signature
	return bridgeDataAbi.Pack("", name, symbol, decimals)
}

// GetERC20BridgeCalldata returns the calldata that will be sent by an L1 ERC20 bridge to its L2 counterpart
// during bridging of a token.
func GetERC20BridgeCalldata(l1TokenAddress, l1Sender, l2Receiver common.Address, amount *big.Int, bridgeData []byte) ([]byte, error) {
	l2BridgeAbi, err := l2bridge.IL2BridgeMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to load L2 bridge ABI: %w", err)
	}
	return l2BridgeAbi.Pack("finalizeDeposit", l1Sender, l2Receiver, l1TokenAddress, amount, bridgeData)
}
