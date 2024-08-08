package accounts

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/eip712"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"testing"
)

const PrivateKey1 = "7726827caac94a7f9e1b160f7ea819f172f7b6f9d2a97f992c38edeab82d4110"
const PrivateKey2 = "ac1e735be8536c6534bb4f17f06f6afc73b2b5ba84ac2cfb12f7461b20c0bbe3"

var Address1 = common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
var Address2 = common.HexToAddress("0xa61464658AfeAf65CccaaFD3a512b69A83B77618")

func TestSignPayloadWithECDSASignTransaction(t *testing.T) {
	signature := "0x475e207d1e5da85721e37118cea54b2a3ac8e5dcd79cd7935c59bdd5cbc71e9824d4ab9dbaa5f8542e51588f4187c406fc4311c2ce9a9aa2a269f14298e5777d1b"

	tx := types.Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(1_000_000_000),
		Gas:       big.NewInt(1_000_000_000),
		To:        &Address2,
		Value:     big.NewInt(7_000_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		Meta: &types.Eip712Meta{
			GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
		},
	}

	domain := eip712.ZkSyncEraEIP712Domain(270)

	eip712Msg, err := tx.EIP712Message()
	assert.NoError(t, err, "EIP712Message should not return an error")

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types: apitypes.Types{
			tx.EIP712Type():     tx.EIP712Types(),
			domain.EIP712Type(): domain.EIP712Types(),
		},
		PrimaryType: tx.EIP712Type(),
		Domain:      domain.EIP712Domain(),
		Message:     eip712Msg,
	})
	assert.NoError(t, err, "TypedDataAndHash should not return an error")

	sig, err := SignPayloadWithECDSA(context.Background(), hash, PrivateKey1, nil)
	assert.NoError(t, err, "SignPayloadWithECDSA should not return an error")
	assert.Equal(t, signature, hexutil.Encode(sig), "Signatures must be the same")
}

func TestSignPayloadWithECDSASignMessage(t *testing.T) {
	signature := "0x7c15eb760c394b0ca49496e71d841378d8bfd4f9fb67e930eb5531485329ab7c67068d1f8ef4b480ec327214ee6ed203687e3fbe74b92367b259281e340d16fd1c"

	sig, err := SignPayloadWithECDSA(context.Background(), accounts.TextHash([]byte("Hello World!")), PrivateKey1, nil)
	assert.NoError(t, err, "SignPayloadWithECDSA should not return an error")
	assert.Equal(t, signature, hexutil.Encode(sig), "Signatures must be the same")
}

func TestSignPayloadWithECDSASignTypedData(t *testing.T) {
	signature := "0xbcaf0673c0c2b0e120165d207d42281d0c6e85f0a7f6b8044b0578a91cf5bda66b4aeb62aca4ae17012a38d71c9943e27285792fa7d788d848f849e3ea2e614b1b"

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    "Example",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(270),
		},
		Types: apitypes.Types{
			"Person": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "age", Type: "uint8"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
		},
		PrimaryType: "Person",
		Message: apitypes.TypedDataMessage{
			"name": "John",
			"age":  hexutil.EncodeUint64(30),
		},
	})
	assert.NoError(t, err, "TypedDataAndHash should not return an error")

	sig, err := SignPayloadWithECDSA(context.Background(), hash, PrivateKey1, nil)
	assert.NoError(t, err, "SignPayloadWithECDSA should not return an error")
	assert.Equal(t, signature, hexutil.Encode(sig), "Signatures must be the same")
}

func TestSignPayloadWithMultipleECDSASignTransaction(t *testing.T) {
	signature := "0x475e207d1e5da85721e37118cea54b2a3ac8e5dcd79cd7935c59bdd5cbc71e9824d4ab9dbaa5f8542e51588f4187c406fc4311c2ce9a9aa2a269f14298e5777d1b4ff4f280885d2dd0b2234d82cacec8ba94bd6659b64b1d516668b4ca79faf58a58c469fd95590e2541ca01866e312e56c7e38a74b4a8b72fdb07a69a3b34c19f1c"

	tx := types.Transaction712{
		Nonce:     big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasFeeCap: big.NewInt(1_000_000_000),
		Gas:       big.NewInt(1_000_000_000),
		To:        &Address2,
		Value:     big.NewInt(7_000_000_000),
		Data:      hexutil.Bytes{},
		ChainID:   big.NewInt(270),
		From:      &Address1,
		Meta: &types.Eip712Meta{
			GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
		},
	}

	domain := eip712.ZkSyncEraEIP712Domain(270)

	eip712Msg, err := tx.EIP712Message()
	assert.NoError(t, err, "EIP712Message should not return an error")

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types: apitypes.Types{
			tx.EIP712Type():     tx.EIP712Types(),
			domain.EIP712Type(): domain.EIP712Types(),
		},
		PrimaryType: tx.EIP712Type(),
		Domain:      domain.EIP712Domain(),
		Message:     eip712Msg,
	})
	assert.NoError(t, err, "TypedDataAndHash should not return an error")

	sig, err := SignPayloadWithMultipleECDSA(context.Background(), hash, []string{PrivateKey1, PrivateKey2}, nil)
	assert.NoError(t, err, "SignPayloadWithMultipleECDSA should not return an error")
	assert.Equal(t, signature, hexutil.Encode(sig), "Signatures must be the same")
}

func TestSignPayloadWithMultipleECDSASignMessage(t *testing.T) {
	signature := "0x7c15eb760c394b0ca49496e71d841378d8bfd4f9fb67e930eb5531485329ab7c67068d1f8ef4b480ec327214ee6ed203687e3fbe74b92367b259281e340d16fd1c2f2f4a312d23de1bcadff9c547fe670a9e21beae16a7c9688fc10b97ba2e286574de339c2b70bd3f02bd021c270a1405942cc3e1268bf3f7a7a419a1c7aea2db1c"

	sig, err := SignPayloadWithMultipleECDSA(
		context.Background(),
		accounts.TextHash([]byte("Hello World!")),
		[]string{PrivateKey1, PrivateKey2},
		nil)
	assert.NoError(t, err, "SignPayloadWithMultipleECDSA should not return an error")
	assert.Equal(t, signature, hexutil.Encode(sig), "Signatures must be the same")
}

func TestSignPayloadWithMultipleECDSASignTypedData(t *testing.T) {
	signature := "0xbcaf0673c0c2b0e120165d207d42281d0c6e85f0a7f6b8044b0578a91cf5bda66b4aeb62aca4ae17012a38d71c9943e27285792fa7d788d848f849e3ea2e614b1b8231ec20acfc86483b908e8f1e88c917b244465c7e73202b6f2643377a6e54f5640f0d3e2f5902695faec96668b2e998148c49a4de613bb7bc4325a3c855cf6a1b"

	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    "Example",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(270),
		},
		Types: apitypes.Types{
			"Person": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "age", Type: "uint8"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
		},
		PrimaryType: "Person",
		Message: apitypes.TypedDataMessage{
			"name": "John",
			"age":  hexutil.EncodeUint64(30),
		},
	})
	assert.NoError(t, err, "TypedDataAndHash should not return an error")

	sig, err := SignPayloadWithMultipleECDSA(context.Background(), hash, []string{PrivateKey1, PrivateKey2}, nil)
	assert.NoError(t, err, "SignPayloadWithMultipleECDSA should not return an error")
	assert.Equal(t, signature, hexutil.Encode(sig), "Signatures must be the same")
}
