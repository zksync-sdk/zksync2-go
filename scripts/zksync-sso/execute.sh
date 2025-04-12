#!/bin/bash
set -x  # or set -ex for error + command trace

# Fetch the contracts from https://github.com/matter-labs/zksync-sso-clave-contracts
# and generates ABI json files and puts them in abi folder.
chmod +x entrypoint-abi.sh
docker create --name fetch-abi-zksso --entrypoint /usr/local/bin/entrypoint-abi.sh node:23
docker cp entrypoint-abi.sh fetch-abi-zksso:/usr/local/bin/entrypoint-abi.sh
docker start -i fetch-abi-zksso
mkdir -p ../../abi/zksso
docker cp fetch-abi-zksso:/abi/. ../../abi/zksso
docker rm fetch-abi-zksso

# Generates go bindings from abi folder using abigen tool
# and puts them in contracts folder.
chmod +x entrypoint-contracts.sh
docker create -it --name generate-go-contracts-zksso --entrypoint /usr/local/bin/entrypoint-contracts.sh ethereum/client-go:alltools-v1.15.7
docker cp entrypoint-contracts.sh generate-go-contracts-zksso:/usr/local/bin/entrypoint-contracts.sh
docker cp ../../abi/zksso generate-go-contracts-zksso:/abi
docker start -i generate-go-contracts-zksso
docker cp generate-go-contracts-zksso:/contracts/. ../../contracts
docker rm generate-go-contracts-zksso
