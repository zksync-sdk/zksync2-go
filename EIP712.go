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
	FeeToken       common.Address `json:"feeToken"`
	ErgsPerPubdata *hexutil.Big   `json:"ergsPerPubdata,omitempty"`
	ErgsPerStorage *hexutil.Big   `json:"ergsPerStorage,omitempty" rlp:"-"`
	FactoryDeps    [][]byte       `json:"factoryDeps"`
	AAParams       *AAParams      `json:"aaParams"`
}

type AAParams struct {
	From      string `json:"from"`
	Signature []byte `json:"signature"`
}

type Eip712Domain struct {
	Name              string         `json:"name"`
	Version           string         `json:"version"`
	ChainId           *big.Int       `json:"chainId"`
	VerifyingContract common.Address `json:"verifyingContract"`
}

func (d *Eip712Domain) GetEIP712Type() string {
	return "EIP712Domain"
}

func (d *Eip712Domain) GetEIP712Types() []apitypes.Type {
	return []apitypes.Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		//{Name: "verifyingContract", Type: "address"},
	}
}

func (d *Eip712Domain) GetEIP712Domain() apitypes.TypedDataDomain {
	return apitypes.TypedDataDomain{
		Name:    d.Name,
		Version: d.Version,
		ChainId: math.NewHexOrDecimal256(d.ChainId.Int64()),
		//VerifyingContract: d.VerifyingContract.String(),
	}
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
		VerifyingContract: common.Address{},
	}
}
