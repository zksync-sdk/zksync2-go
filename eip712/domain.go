package eip712

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

type TypedData interface {
	GetEIP712Type() string
	GetEIP712Types() []apitypes.Type
	GetEIP712Message() (apitypes.TypedDataMessage, error)
}

type Domain struct {
	Name              string          `json:"name"`
	Version           string          `json:"version"`
	ChainId           *big.Int        `json:"chainId"`
	VerifyingContract *common.Address `json:"verifyingContract"`
}

func (d *Domain) GetEIP712Type() string {
	return "EIP712Domain"
}

func (d *Domain) GetEIP712Types() []apitypes.Type {
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

func (d *Domain) GetEIP712Domain() apitypes.TypedDataDomain {
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
	DomainDefaultName    = `zkSync`
	DomainDefaultVersion = `2`
)

func DefaultEip712Domain(chainId int64) *Domain {
	return &Domain{
		Name:              DomainDefaultName,
		Version:           DomainDefaultVersion,
		ChainId:           big.NewInt(chainId),
		VerifyingContract: nil,
	}
}
