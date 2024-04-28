#!/bin/ash

mkdir -p /contracts/bridgehub
abigen  \
	--abi /abi/IBridgehub.json \
	--out /contracts/bridgehub/bridgehub.go \
	--pkg bridgehub \
	--type IBridgehub

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

mkdir -p /contracts/l1erc20bridge
abigen  \
	--abi /abi/IL1ERC20Bridge.json \
	--out /contracts/l1erc20bridge/l1_erc20_bridge.go \
	--pkg l1erc20bridge \
	--type IL1ERC20Bridge

mkdir -p /contracts/l1messenger
abigen  \
	--abi /abi/IL1Messenger.json \
	--out /contracts/l1messenger/l1_messenger.go \
	--pkg l1messenger \
	--type IL1Messenger

mkdir -p /contracts/l1sharedbridge
abigen  \
	--abi /abi/IL1SharedBridge.json \
	--out /contracts/l1sharedbridge/l1_shared_bridge.go \
	--pkg l1sharedbridge \
	--type IL1SharedBridge

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

mkdir -p /contracts/testneterc20token
abigen  \
	--abi /abi/ITestnetERC20Token.json \
	--out /contracts/testneterc20token/testnet_erc20_token.go \
	--pkg testneterc20token \
	--type ITestnetERC20Token

mkdir -p /contracts/zksyncstatetransition
abigen  \
	--abi /abi/IZkSyncStateTransition.json \
	--out /contracts/zksyncstatetransition/zksync_state_transition.go \
	--pkg zksyncstatetransition \
	--type IZkSyncStateTransition

echo "Folder content"
ls -alhR /contracts
