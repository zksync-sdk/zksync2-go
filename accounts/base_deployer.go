package accounts

import (
	"github.com/ethereum/go-ethereum/common"
)

// BaseDeployer implements the Deployer interface.
type BaseDeployer struct {
	adapter *AdapterL2
}

// NewBaseDeployer creates an instance of BaseDeployer.
func NewBaseDeployer(adapter *AdapterL2) *BaseDeployer {
	return &BaseDeployer{adapter}
}

func (a *BaseDeployer) Deploy(auth *TransactOpts, tx Create2Transaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployContract, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return (*a.adapter).SendTransaction(opts.Context, preparedTx)
}

func (a *BaseDeployer) DeployWithCreate(auth *TransactOpts, tx CreateTransaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployContract, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return (*a.adapter).SendTransaction(opts.Context, preparedTx)
}

func (a *BaseDeployer) DeployAccount(auth *TransactOpts, tx Create2Transaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployAccount, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return (*a.adapter).SendTransaction(opts.Context, preparedTx)
}

func (a *BaseDeployer) DeployAccountWithCreate(auth *TransactOpts, tx CreateTransaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployAccount, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return (*a.adapter).SendTransaction(opts.Context, preparedTx)
}
