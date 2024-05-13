package test

import (
	"github.com/ethereum/go-ethereum/common"
	"reflect"
)

const EthereumProvider = "http://localhost:8545"
const ZkSyncEraProvider = "http://localhost:3050"
const PrivateKey1 = "7726827caac94a7f9e1b160f7ea819f172f7b6f9d2a97f992c38edeab82d4110"
const PrivateKey2 = "ac1e735be8536c6534bb4f17f06f6afc73b2b5ba84ac2cfb12f7461b20c0bbe3"

var Address1 = common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
var Address2 = common.HexToAddress("0xa61464658AfeAf65CccaaFD3a512b69A83B77618")

var Salt = common.Hex2Bytes("0x293328ad84b118194c65a0dc0defdb6483740d3163fd99b260907e15f2e2f642")

// ApprovalToken is deployed using create2 and Salt
var ApprovalToken = common.HexToAddress("0x834FF28392Ab0460f13286c389fEF4E3980e28F6")

// Paymaster is an approval based paymaster for approval token deployed using create2 and Salt
var Paymaster = common.HexToAddress("0xe1438081bF20c4C910266aa1229930473446b283")

// MultisigAccount is an account that signs transaction using PrivateKey1 and PrivateKey2
// and is deployed using create2 and Salt
var MultisigAccount = common.HexToAddress("0x60222D60b22D3e2A6F459Dc7264aEbf9f8735363")

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

func DeepEqualIgnore(s1, s2 interface{}, ignoreFields ...string) bool {
	// contains checks if a string slice contains a specific string.
	contains := func(slice []string, str string) bool {
		for _, s := range slice {
			if s == str {
				return true
			}
		}
		return false
	}

	v1 := reflect.ValueOf(s1)
	v2 := reflect.ValueOf(s2)

	for i := 0; i < v1.NumField(); i++ {
		field := v1.Type().Field(i)
		fieldName := field.Name

		if contains(ignoreFields, fieldName) {
			continue
		}

		val1 := v1.Field(i)
		val2 := v2.Field(i)

		if !reflect.DeepEqual(val1.Interface(), val2.Interface()) {
			return false
		}
	}

	return true
}
