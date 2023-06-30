package eip712

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

// TypedData represents typed data as defined by EIP-712.
type TypedData interface {
	// EIP712Type returns the EIP-712 type.
	EIP712Type() string
	// EIP712Types return the supported types.
	EIP712Types() []apitypes.Type
	// EIP712Message returns the EIP-712 message.
	EIP712Message() (apitypes.TypedDataMessage, error)
}

// Domain represents the domain parameters used for EIP-712 signing.
type Domain struct {
	Name              string          `json:"name"`              // Name of the domain.
	Version           string          `json:"version"`           // Version of the domain.
	ChainId           *big.Int        `json:"chainId"`           // Chain ID associated with the domain.
	VerifyingContract *common.Address `json:"verifyingContract"` // Address of the verifying contract for the domain.
}

func (d *Domain) EIP712Type() string {
	return "EIP712Domain"
}

func (d *Domain) EIP712Types() []apitypes.Type {
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

func (d *Domain) EIP712Domain() apitypes.TypedDataDomain {
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

func ZkSyncEraEIP712Domain(chainId int64) *Domain {
	return &Domain{
		Name:              DomainDefaultName,
		Version:           DomainDefaultVersion,
		ChainId:           big.NewInt(chainId),
		VerifyingContract: nil,
	}
}
