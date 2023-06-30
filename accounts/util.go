package accounts

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}

// ensureCallOpts is a helper method to ensure a CallOpts is not nil, even if the
// user specified it as such.
func ensureCallOpts(opts *CallOpts) *CallOpts {
	if opts == nil {
		return &CallOpts{Context: context.Background()}
	}
	return opts
}

func ensureTransactOpts(auth *TransactOpts) *TransactOpts {
	if auth == nil {
		return &TransactOpts{
			Context: context.Background(),
		}
	}
	if auth.Context == nil {
		auth.Context = context.Background()
	}
	return auth
}

func newTransactorWithSigner(signer *Signer, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, bind.ErrNoChainID
	}
	keyAddr := (*signer).Address()
	latestSigner := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}
			signature, err := (*signer).SignHash(latestSigner.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(latestSigner, signature)
		},
	}, nil
}
