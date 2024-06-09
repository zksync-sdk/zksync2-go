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

	tx := Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       big.NewInt(0),
		Value:     big.NewInt(1_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		To:        &Address2,
		Meta: &Eip712Meta{
			GasPerPubdata:   (*hexutil.Big)(big.NewInt(50_000)),
			CustomSignature: hexutil.Bytes{},
			FactoryDeps:     []hexutil.Bytes{},
		},
	}

	decodedTx := new(Transaction712)
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

	tx := Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       big.NewInt(0),
		Value:     big.NewInt(1_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		To:        &Address2,
		Meta: &Eip712Meta{
			GasPerPubdata:   (*hexutil.Big)(big.NewInt(50_000)),
			CustomSignature: hexutil.Bytes{},
			FactoryDeps:     []hexutil.Bytes{},
			PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
		},
	}

	decodedTx := new(Transaction712)
	err = decodedTx.Decode(common.FromHex(serializedTx))
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, *decodedTx, tx, "Transaction should be same")
}

func TestTransaction712_DecodeSignedTx(t *testing.T) {
	serializedTx := "0x71f8c68080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c0b884307861666533646639333965383139373437613036386664356134346633316335313765643965313933336166333730633062636532623232303139623666373139326161336536316364363165373039356264353665343636316639336130303231396133393663313130366364383935623361663338623534633236373032313163c0"
	signature := "0xafe3df939e819747a068fd5a44f31c517ed9e1933af370c0bce2b22019b6f7192aa3e61cd61e7095bd56e4661f93a00219a396c1106cd895b3af38b54c2670211c"

	tx := Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       big.NewInt(0),
		Value:     big.NewInt(1_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		To:        &Address2,
		Meta: &Eip712Meta{
			GasPerPubdata:   (*hexutil.Big)(big.NewInt(50_000)),
			CustomSignature: hexutil.Bytes(signature),
			FactoryDeps:     []hexutil.Bytes{},
		},
	}

	decodedTx := new(Transaction712)
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

	tx := Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       big.NewInt(0),
		Value:     big.NewInt(1_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		To:        &Address2,
		Meta: &Eip712Meta{
			GasPerPubdata:   (*hexutil.Big)(big.NewInt(50_000)),
			CustomSignature: hexutil.Bytes(signature),
			FactoryDeps:     []hexutil.Bytes{},
			PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
		},
	}

	decodedTx := new(Transaction712)
	err = decodedTx.Decode(common.FromHex(serializedTx))
	assert.NoError(t, err, "Decode should not return an error")
	assert.Equal(t, *decodedTx, tx, "Transaction should be same")
}

func TestTransaction712_RLPValues(t *testing.T) {
	serializedTx := "0x71f8c68080808094a61464658afeaf65cccaafd3a512b69a83b77618830f42408082010e808082010e9436615cf349d7f6344891b1e7ca7c72883f5dc04982c350c0b884307861666533646639333965383139373437613036386664356134346633316335313765643965313933336166333730633062636532623232303139623666373139326161336536316364363165373039356264353665343636316639336130303231396133393663313130366364383935623361663338623534633236373032313163c0"
	signature := "0xafe3df939e819747a068fd5a44f31c517ed9e1933af370c0bce2b22019b6f7192aa3e61cd61e7095bd56e4661f93a00219a396c1106cd895b3af38b54c2670211c"

	tx := Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       big.NewInt(0),
		Value:     big.NewInt(1_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		To:        &Address2,
		Meta: &Eip712Meta{
			GasPerPubdata:   (*hexutil.Big)(big.NewInt(50_000)),
			CustomSignature: hexutil.Bytes(signature),
			FactoryDeps:     []hexutil.Bytes{},
		},
	}

	result, err := tx.RLPValues(nil)
	assert.NoError(t, err, "RLPValues should not return an error")
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

	tx := Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       big.NewInt(0),
		Value:     big.NewInt(1_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		To:        &Address2,
		Meta: &Eip712Meta{
			GasPerPubdata:   (*hexutil.Big)(big.NewInt(50_000)),
			CustomSignature: hexutil.Bytes{},
			FactoryDeps:     []hexutil.Bytes{},
			PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
		},
	}

	result, err := tx.RLPValues(nil)
	assert.NoError(t, err, "RLPValues should not return an error")
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

	tx := Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(0),
		Gas:       big.NewInt(0),
		Value:     big.NewInt(1_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		To:        &Address2,
		Meta: &Eip712Meta{
			GasPerPubdata:   (*hexutil.Big)(big.NewInt(50_000)),
			CustomSignature: hexutil.Bytes(signature),
			FactoryDeps:     []hexutil.Bytes{},
			PaymasterParams: &PaymasterParams{Paymaster, paymasterInput},
		},
	}
	result, err := tx.RLPValues(nil)
	assert.NoError(t, err, "RLPValues should not return an error")
	assert.Equal(t, common.FromHex(serializedTx), result, "Serialized transactions should be same")
}
