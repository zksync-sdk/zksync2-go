#!/bin/bash

# Fetch the contracts from https://github.com/matter-labs/era-contracts
# and generates ABI json files and puts them in abi folder.

docker create --name abi-gen --entrypoint /usr/local/bin/entrypoint.sh node:18
docker cp entrypoint.sh abi-gen:/usr/local/bin/entrypoint.sh
docker start -i abi-gen
docker cp abi-gen:/abi/. ../../abi
docker rm abi-gen