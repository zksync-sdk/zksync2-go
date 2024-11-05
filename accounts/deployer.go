package accounts

import (
	"github.com/ethereum/go-ethereum/common"
)

// Deployer is associated with an account and provides deployment of smart contracts
// and smart accounts on L2 network for the associated account.
type Deployer struct {
	runner *WalletL2
}

// NewDeployer creates an instance of BaseDeployer.
func NewDeployer(adapter *WalletL2) *Deployer {
	return &Deployer{adapter}
}

// Deploy deploys smart contract using CREATE2 opcode.
func (a *Deployer) Deploy(auth *TransactOptsL1, tx Create2Transaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployContract, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return a.runner.SendTransaction(opts.Context, preparedTx)
}

// DeployWithCreate deploys smart contract using CREATE opcode.
func (a *Deployer) DeployWithCreate(auth *TransactOptsL1, tx CreateTransaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployContract, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return a.runner.SendTransaction(opts.Context, preparedTx)
}

// DeployAccount deploys smart account using CREATE2 opcode.
func (a *Deployer) DeployAccount(auth *TransactOptsL1, tx Create2Transaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployAccount, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return a.runner.SendTransaction(opts.Context, preparedTx)
}

// DeployAccountWithCreate deploys smart account using CREATE opcode.
func (a *Deployer) DeployAccountWithCreate(auth *TransactOptsL1, tx CreateTransaction) (common.Hash, error) {
	opts := ensureTransactOpts(auth)
	preparedTx, err := tx.ToTransaction(DeployAccount, opts)
	if err != nil {
		return [32]byte{}, err
	}
	return a.runner.SendTransaction(opts.Context, preparedTx)
}
