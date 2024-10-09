package eip712

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

// Domain represents the domain parameters used for EIP-712 signing.
type Domain struct {
	Name              string          `json:"name"`              // Name of the domain.
	Version           string          `json:"version"`           // Version of the domain.
	ChainId           *big.Int        `json:"chainId"`           // Chain ID associated with the domain.
	VerifyingContract *common.Address `json:"verifyingContract"` // Address of the verifying contract for the domain.
}

// Type returns the name of the domain field.
func (d *Domain) Type() string {
	return "EIP712Domain"
}

// Types returns the domain field types.
func (d *Domain) Types() []apitypes.Type {
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

// TypedData returns domain typed data.
func (d *Domain) TypedData() apitypes.TypedDataDomain {
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

// ZkSyncEraEIP712Domain represents the EIP-712 domain for ZKsync Era.
func ZkSyncEraEIP712Domain(chainId int64) *Domain {
	return &Domain{
		Name:              DomainDefaultName,
		Version:           DomainDefaultVersion,
		ChainId:           big.NewInt(chainId),
		VerifyingContract: nil,
	}
}
