fetch-contracts:
	cd scripts/fetch-abi && ./execute.sh && cd ../..

generate-contracts:
	cd scripts/generate-contracts && ./execute.sh && cd ../..

run-tests-on-eth-based-chain:
	go test -v -skip='^.*_NonEthBasedChain_.*$\' ./test  ./accounts ./utils ./types

run-tests-on-non-eth-based-chain:
	ETH_BASED_CHAIN=false go test -v -skip='^.*_EthBasedChain_.*$\'  ./test ./accounts ./utils ./types

check-format:
	cd scripts/ && ./check-format.sh && cd ../..

format:
	gofmt -w .
