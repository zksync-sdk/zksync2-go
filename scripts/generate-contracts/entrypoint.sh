#!/bin/ash

mkdir -p /contracts/contractdeployer
abigen  \
	--abi /abi/IContractDeployer.json \
	--out /contracts/contractdeployer/contract_deployer.go \
	--pkg contractdeployer \
	--type IContractDeployer

mkdir -p /contracts/ethtoken
abigen  \
	--abi /abi/IEthToken.json \
	--out /contracts/ethtoken/eth_token.go \
	--pkg ethtoken \
	--type IEthToken

mkdir -p /contracts/l1bridge
abigen  \
	--abi /abi/IL1Bridge.json \
	--out /contracts/l1bridge/l1_bridge.go \
	--pkg l1bridge \
	--type IL1Bridge

mkdir -p /contracts/l1messenger
abigen  \
	--abi /abi/IL1Messenger.json \
	--out /contracts/l1messenger/l1_messenger.go \
	--pkg l1messenger \
	--type IL1Messenger

mkdir -p /contracts/l2bridge
abigen  \
	--abi /abi/IL2Bridge.json \
	--out /contracts/l2bridge/l2_bridge.go \
	--pkg l2bridge \
	--type IL2Bridge

mkdir -p /contracts/nonceholder
abigen  \
	--abi /abi/INonceHolder.json \
	--out /contracts/nonceholder/nonce_holder.go \
	--pkg nonceholder \
	--type INonceHolder

mkdir -p /contracts/paymasterflow
abigen  \
	--abi /abi/IPaymasterFlow.json \
	--out /contracts/paymasterflow/paymaster_flow.go \
	--pkg paymasterflow \
	--type IPaymasterFlow

mkdir -p /contracts/zksync
abigen  \
	--abi /abi/IZkSync.json \
	--out /contracts/zksync/zk_sync.go \
	--pkg zksync \
	--type IZkSync

echo "Folder content"
ls -alhR /contracts