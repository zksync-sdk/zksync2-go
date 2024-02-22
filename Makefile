fetch-contracts:
	cd scripts/fetch-abi && ./execute.sh && cd ../..

generate-contracts:
	cd scripts/generate-contracts && ./execute.sh && cd ../..

run-tests:
	go test ./utils ./test

check-format:
	cd scripts/ && ./check-format.sh && cd ../..

format:
	gofmt -w .
