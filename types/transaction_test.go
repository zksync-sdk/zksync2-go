package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/contracts/paymasterflow"
	"math/big"
	"strings"
	"testing"
)

var Address1 = common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
var Address2 = common.HexToAddress("0xa61464658AfeAf65CccaaFD3a512b69A83B77618")
var Paymaster = common.HexToAddress("0xa222f0c183AFA73a8Bc1AFb48D34C88c9Bf7A174")
var ApprovalToken = common.HexToAddress("0x841c43Fa5d8fFfdB9efE3358906f7578d8700Dd4")

func TestTransaction712_DecodeUnsignedTx(t *testing.T) {
	serializedTx := "0x71f8418080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c080c0"

	tx := Transaction{
		Nonce:           big.NewInt(0),
		GasTipCap:       big.NewInt(0),
		GasFeeCap:       big.NewInt(0),
		Gas:             big.NewInt(0),
		Value:           big.NewInt(1_000_000),
		Data:            hexutil.Bytes{},
		ChainID:         big.NewInt(270),
		From:            &Address1,
		To:              &Address2,
		GasPerPubdata:   big.NewInt(50_000),
		CustomSignature: hexutil.Bytes{},
		FactoryDeps:     []hexutil.Bytes{},
	}

	decodedTx := new(Transaction)
	err := decodedTx.Decode(common.FromHex(serializedTx))
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, *decodedTx, tx, "Transaction should be same")
}

func TestTransaction712_DecodeUnsignedTxWithPaymaster(t *testing.T) {
	serializedTx := "0x71f8dd8080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c080f89b94a222f0c183afa73a8bc1afb48d34c88c9bf7a174b884949431dc000000000000000000000000841c43fa5d8fffdb9efe3358906f7578d8700dd4000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"

	paymasterFlowAbi, err := abi.JSON(strings.NewReader(paymasterflow.IPaymasterFlowMetaData.ABI))
	paymasterInput, err := paymasterFlowAbi.Pack("approvalBased",
		ApprovalToken,
		big.NewInt(1),
		[]byte{},
	)
	assert.NoError(t, err, "Pack should not return an error")

	tx := Transaction{
		Nonce:           big.NewInt(0),
		GasTipCap:       big.NewInt(0),
		GasFeeCap:       big.NewInt(0),
		Gas:             big.NewInt(0),
		Value:           big.NewInt(1_000_000),
		Data:            hexutil.Bytes{},
		ChainID:         big.NewInt(270),
		From:            &Address1,
		To:              &Address2,
		GasPerPubdata:   big.NewInt(50_000),
		CustomSignature: hexutil.Bytes{},
		FactoryDeps:     []hexutil.Bytes{},
		PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
	}

	decodedTx := new(Transaction)
	err = decodedTx.Decode(common.FromHex(serializedTx))
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, *decodedTx, tx, "Transaction should be same")
}

func TestTransaction712_DecodeSignedTx(t *testing.T) {
	serializedTx := "0x71f8c68080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c0b884307861666533646639333965383139373437613036386664356134346633316335313765643965313933336166333730633062636532623232303139623666373139326161336536316364363165373039356264353665343636316639336130303231396133393663313130366364383935623361663338623534633236373032313163c0"
	signature := "0xafe3df939e819747a068fd5a44f31c517ed9e1933af370c0bce2b22019b6f7192aa3e61cd61e7095bd56e4661f93a00219a396c1106cd895b3af38b54c2670211c"

	tx := Transaction{
		Nonce:           big.NewInt(0),
		GasTipCap:       big.NewInt(0),
		GasFeeCap:       big.NewInt(0),
		Gas:             big.NewInt(0),
		Value:           big.NewInt(1_000_000),
		Data:            hexutil.Bytes{},
		ChainID:         big.NewInt(270),
		From:            &Address1,
		To:              &Address2,
		GasPerPubdata:   big.NewInt(50_000),
		CustomSignature: hexutil.Bytes(signature),
		FactoryDeps:     []hexutil.Bytes{},
	}

	decodedTx := new(Transaction)
	err := decodedTx.Decode(common.FromHex(serializedTx))
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, *decodedTx, tx, "Transaction should be same")
}

func TestTransaction712_DecodeSignedTxWithPaymaster(t *testing.T) {
	serializedTx := "0x71f901628080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c0b884307837373262396162343735386435636630386637643732303161646332653534383933616532376263666562323162396337643666643430393766346464653063303166376630353332323866346636643838653662663334333436343931343135363761633930363632306661653832633239333339393062353563613336363162f89b94a222f0c183afa73a8bc1afb48d34c88c9bf7a174b884949431dc000000000000000000000000841c43fa5d8fffdb9efe3358906f7578d8700dd4000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"
	signature := "0x772b9ab4758d5cf08f7d7201adc2e54893ae27bcfeb21b9c7d6fd4097f4dde0c01f7f053228f4f6d88e6bf3434649141567ac906620fae82c2933990b55ca3661b"

	paymasterFlowAbi, err := abi.JSON(strings.NewReader(paymasterflow.IPaymasterFlowMetaData.ABI))
	paymasterInput, err := paymasterFlowAbi.Pack("approvalBased",
		ApprovalToken,
		big.NewInt(1),
		[]byte{},
	)
	assert.NoError(t, err, "Pack should not return an error")

	tx := Transaction{
		Nonce:           big.NewInt(0),
		GasTipCap:       big.NewInt(0),
		GasFeeCap:       big.NewInt(0),
		Gas:             big.NewInt(0),
		Value:           big.NewInt(1_000_000),
		Data:            hexutil.Bytes{},
		ChainID:         big.NewInt(270),
		From:            &Address1,
		To:              &Address2,
		GasPerPubdata:   big.NewInt(50_000),
		CustomSignature: hexutil.Bytes(signature),
		FactoryDeps:     []hexutil.Bytes{},
		PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
	}

	decodedTx := new(Transaction)
	err = decodedTx.Decode(common.FromHex(serializedTx))
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, *decodedTx, tx, "Transaction should be same")
}

func TestTransaction712_RLPValues(t *testing.T) {
	serializedTx := "0x71f8c68080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c0b884307861666533646639333965383139373437613036386664356134346633316335313765643965313933336166333730633062636532623232303139623666373139326161336536316364363165373039356264353665343636316639336130303231396133393663313130366364383935623361663338623534633236373032313163c0"
	signature := "0xafe3df939e819747a068fd5a44f31c517ed9e1933af370c0bce2b22019b6f7192aa3e61cd61e7095bd56e4661f93a00219a396c1106cd895b3af38b54c2670211c"

	tx := Transaction{
		Nonce:           big.NewInt(0),
		GasTipCap:       big.NewInt(0),
		GasFeeCap:       big.NewInt(0),
		Gas:             big.NewInt(0),
		Value:           big.NewInt(1_000_000),
		Data:            hexutil.Bytes{},
		ChainID:         big.NewInt(270),
		From:            &Address1,
		To:              &Address2,
		GasPerPubdata:   big.NewInt(50_000),
		CustomSignature: hexutil.Bytes(signature),
		FactoryDeps:     []hexutil.Bytes{},
	}

	result, err := tx.Encode(nil)
	assert.NoError(t, err, "Encode should not return an error")
	assert.Equal(t, common.FromHex(serializedTx), result, "Serialized transactions should be same")
}

func TestTransaction712_RLPValuesPaymasterWithoutSignature(t *testing.T) {
	serializedTx := "0x71f8dd8080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c080f89b94a222f0c183afa73a8bc1afb48d34c88c9bf7a174b884949431dc000000000000000000000000841c43fa5d8fffdb9efe3358906f7578d8700dd4000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"

	paymasterFlowAbi, err := abi.JSON(strings.NewReader(paymasterflow.IPaymasterFlowMetaData.ABI))
	paymasterInput, err := paymasterFlowAbi.Pack("approvalBased",
		ApprovalToken,
		big.NewInt(1),
		[]byte{},
	)
	assert.NoError(t, err, "Pack should not return an error")

	tx := Transaction{
		Nonce:           big.NewInt(0),
		GasTipCap:       big.NewInt(0),
		GasFeeCap:       big.NewInt(0),
		Gas:             big.NewInt(0),
		Value:           big.NewInt(1_000_000),
		Data:            hexutil.Bytes{},
		ChainID:         big.NewInt(270),
		From:            &Address1,
		To:              &Address2,
		GasPerPubdata:   big.NewInt(50_000),
		CustomSignature: hexutil.Bytes{},
		FactoryDeps:     []hexutil.Bytes{},
		PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
	}

	result, err := tx.Encode(nil)
	assert.NoError(t, err, "Encode should not return an error")
	assert.Equal(t, common.FromHex(serializedTx), result, "Serialized transactions should be same")
}

func TestTransaction712_RLPValuesPaymasterWithSignature(t *testing.T) {
	serializedTx := "0x71f901628080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c0b884307837373262396162343735386435636630386637643732303161646332653534383933616532376263666562323162396337643666643430393766346464653063303166376630353332323866346636643838653662663334333436343931343135363761633930363632306661653832633239333339393062353563613336363162f89b94a222f0c183afa73a8bc1afb48d34c88c9bf7a174b884949431dc000000000000000000000000841c43fa5d8fffdb9efe3358906f7578d8700dd4000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"
	signature := "0x772b9ab4758d5cf08f7d7201adc2e54893ae27bcfeb21b9c7d6fd4097f4dde0c01f7f053228f4f6d88e6bf3434649141567ac906620fae82c2933990b55ca3661b"

	paymasterFlowAbi, err := abi.JSON(strings.NewReader(paymasterflow.IPaymasterFlowMetaData.ABI))
	paymasterInput, err := paymasterFlowAbi.Pack("approvalBased",
		ApprovalToken,
		big.NewInt(1),
		[]byte{},
	)
	assert.NoError(t, err, "Pack should not return an error")

	tx := Transaction{
		Nonce:           big.NewInt(0),
		GasTipCap:       big.NewInt(0),
		GasFeeCap:       big.NewInt(0),
		Gas:             big.NewInt(0),
		Value:           big.NewInt(1_000_000),
		Data:            hexutil.Bytes{},
		ChainID:         big.NewInt(270),
		From:            &Address1,
		To:              &Address2,
		GasPerPubdata:   big.NewInt(50_000),
		CustomSignature: hexutil.Bytes(signature),
		FactoryDeps:     []hexutil.Bytes{},
		PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
	}
	result, err := tx.Encode(nil)
	assert.NoError(t, err, "Encode should not return an error")
	assert.Equal(t, common.FromHex(serializedTx), result, "Serialized transactions should be same")
}

func TestTransactionCopy(t *testing.T) {
	original := &Transaction{
		Nonce:           big.NewInt(1),
		GasTipCap:       big.NewInt(2000000000),
		GasFeeCap:       big.NewInt(3000000000),
		Gas:             big.NewInt(21000),
		To:              &Address1,
		Value:           big.NewInt(1000000000000000000),
		Data:            hexutil.Bytes{1, 2, 3, 4},
		ChainID:         big.NewInt(1),
		From:            &Address1,
		GasPerPubdata:   big.NewInt(50000),
		CustomSignature: hexutil.Bytes{5, 6, 7, 8},
		FactoryDeps:     []hexutil.Bytes{{9, 10}, {11, 12}},
		PaymasterParams: &PaymasterParams{
			Paymaster:      Address1,
			PaymasterInput: hexutil.Bytes{13, 14, 15, 16},
		},
	}

	deepCopy := original.Copy()

	// Test equality of values
	assert.Equal(t, original.Nonce, deepCopy.Nonce, "Nonce should be equal")
	assert.Equal(t, original.GasTipCap, deepCopy.GasTipCap, "GasTipCap should be equal")
	assert.Equal(t, original.GasFeeCap, deepCopy.GasFeeCap, "GasFeeCap should be equal")
	assert.Equal(t, original.Gas, deepCopy.Gas, "Gas should be equal")
	assert.Equal(t, original.To, deepCopy.To, "To should be equal")
	assert.Equal(t, original.Value, deepCopy.Value, "Value should be equal")
	assert.Equal(t, original.Data, deepCopy.Data, "Data should be equal")
	assert.Equal(t, original.ChainID, deepCopy.ChainID, "ChainID should be equal")
	assert.Equal(t, original.From, deepCopy.From, "From should be equal")
	assert.Equal(t, original.GasPerPubdata, deepCopy.GasPerPubdata, "GasPerPubdata should be equal")
	assert.Equal(t, original.CustomSignature, deepCopy.CustomSignature, "CustomSignature should be equal")
	assert.Equal(t, original.FactoryDeps, deepCopy.FactoryDeps, "FactoryDeps should be equal")
	assert.Equal(t, original.PaymasterParams, deepCopy.PaymasterParams, "PaymasterParams should be equal")

	// Test different references for pointer and slice fields
	assert.NotSame(t, &original.Nonce, &deepCopy.Nonce, "Nonce should have different memory addresses")
	assert.NotSame(t, &original.GasTipCap, &deepCopy.GasTipCap, "GasTipCap should have different memory addresses")
	assert.NotSame(t, &original.GasFeeCap, &deepCopy.GasFeeCap, "GasFeeCap should have different memory addresses")
	assert.NotSame(t, &original.Gas, &deepCopy.Gas, "Gas should have different memory addresses")
	assert.NotSame(t, original.To, deepCopy.To, "To should have different memory addresses")
	assert.NotSame(t, &original.Value, &deepCopy.Value, "Value should have different memory addresses")
	assert.NotSame(t, &original.Data, &deepCopy.Data, "Data should have different memory addresses")
	assert.NotSame(t, &original.ChainID, &deepCopy.ChainID, "ChainID should have different memory addresses")
	assert.NotSame(t, original.From, deepCopy.From, "From should have different memory addresses")
	assert.NotSame(t, &original.GasPerPubdata, &deepCopy.GasPerPubdata, "GasPerPubdata should have different memory addresses")
	assert.NotSame(t, &original.CustomSignature, &deepCopy.CustomSignature, "CustomSignature should have different memory addresses")
	assert.NotSame(t, &original.FactoryDeps, &deepCopy.FactoryDeps, "FactoryDeps should have different memory addresses")
	assert.NotSame(t, original.PaymasterParams, deepCopy.PaymasterParams, "PaymasterParams should have different memory addresses")

	// Test deep Copy method by modifying the deepCopy
	deepCopy.Nonce.Add(deepCopy.Nonce, big.NewInt(2))
	assert.NotEqual(t, original.Nonce, deepCopy.Nonce, "Modifying deepCopy's Nonce should not affect original")

	newAddress := Address2
	deepCopy.To = &newAddress
	assert.NotEqual(t, original.To, deepCopy.To, "Modifying deepCopy's To should not affect original")

	deepCopy.Data[0] = 99
	assert.NotEqual(t, original.Data, deepCopy.Data, "Modifying deepCopy's Data should not affect original")

	deepCopy.FactoryDeps[0][0] = 99
	assert.NotEqual(t, original.FactoryDeps, deepCopy.FactoryDeps, "Modifying deepCopy's FactoryDeps should not affect original")

	deepCopy.PaymasterParams.PaymasterInput[0] = 99
	assert.NotEqual(t, original.PaymasterParams.PaymasterInput, deepCopy.PaymasterParams.PaymasterInput, "Modifying deepCopy's PaymasterParams should not affect original")

	// Test nil fields
	nilTx := &Transaction{}
	nilCopy := nilTx.Copy()
	assert.Equal(t, &Transaction{
		Nonce:           new(big.Int),
		GasTipCap:       new(big.Int),
		GasFeeCap:       new(big.Int),
		Gas:             new(big.Int),
		To:              nil,
		Value:           new(big.Int),
		Data:            hexutil.Bytes(nil),
		ChainID:         new(big.Int),
		From:            nil,
		GasPerPubdata:   new(big.Int),
		CustomSignature: hexutil.Bytes(nil),
		FactoryDeps:     []hexutil.Bytes{},
		PaymasterParams: nil,
	}, nilCopy, "Copying a Transaction with nil fields should return zero values where appropriate")

	// Test nil Transaction
	var nullTx *Transaction
	nullCopy := nullTx.Copy()
	assert.Nil(t, nullCopy, "Copying a nil Transaction should return nil")
}
