package test

import "github.com/ethereum/go-ethereum/common"

const EthereumProvider = "http://localhost:8545"
const ZkSyncEraProvider = "http://localhost:3050"
const PrivateKey = "7726827caac94a7f9e1b160f7ea819f172f7b6f9d2a97f992c38edeab82d4110"

var Address = common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
var Receiver = common.HexToAddress("0xa61464658AfeAf65CccaaFD3a512b69A83B77618")

var L1Tokens []TokenData
var L2Dai common.Address
var L1Dai common.Address
var L2DepositTx common.Hash
var L1DepositTx common.Hash

type TokenData struct {
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	Decimals int            `json:"decimals"`
	Address  common.Address `json:"address"`
}
