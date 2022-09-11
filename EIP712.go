package zksync2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

type EIP712TypedData interface {
	GetEIP712Type() string
	GetEIP712Types() []apitypes.Type
	GetEIP712Message() apitypes.TypedDataMessage
}

type Eip712Meta struct {
	ErgsPerPubdata  *hexutil.Big     `json:"ergsPerPubdata,omitempty"`
	CustomSignature hexutil.Bytes    `json:"customSignature"`
	FactoryDeps     []hexutil.Bytes  `json:"factoryDeps"`
	PaymasterParams *PaymasterParams `json:"paymasterParams"`
}

type PaymasterParams struct {
	Paymaster      common.Address `json:"paymaster"`
	PaymasterInput hexutil.Bytes  `json:"paymasterInput"`
}

type Eip712Domain struct {
	Name              string          `json:"name"`
	Version           string          `json:"version"`
	ChainId           *big.Int        `json:"chainId"`
	VerifyingContract *common.Address `json:"verifyingContract"`
}

func (d *Eip712Domain) GetEIP712Type() string {
	return "EIP712Domain"
}

func (d *Eip712Domain) GetEIP712Types() []apitypes.Type {
	types := []apitypes.Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
	}
	if d.VerifyingContract != nil {
		types = append(types, apitypes.Type{Name: "verifyingContract", Type: "address"})
	}
	return types
}

func (d *Eip712Domain) GetEIP712Domain() apitypes.TypedDataDomain {
	domain := apitypes.TypedDataDomain{
		Name:    d.Name,
		Version: d.Version,
		ChainId: math.NewHexOrDecimal256(d.ChainId.Int64()),
	}
	if d.VerifyingContract != nil {
		domain.VerifyingContract = d.VerifyingContract.String()
	}
	return domain
}

const (
	Eip712DomainDefaultName    = `zkSync`
	Eip712DomainDefaultVersion = `2`
)

func DefaultEip712Domain(chainId int64) *Eip712Domain {
	return &Eip712Domain{
		Name:              Eip712DomainDefaultName,
		Version:           Eip712DomainDefaultVersion,
		ChainId:           big.NewInt(chainId),
		VerifyingContract: nil,
	}
}
