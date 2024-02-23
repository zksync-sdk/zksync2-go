#!/bin/bash

# Generates go bindings from abi folder using abigen tool
# and puts them in contracts folder.

docker create -it --name abigen \
	--entrypoint /usr/local/bin/entrypoint.sh \
	ethereum/client-go:alltools-v1.13.13

docker cp ../../abi abigen:/abi
docker cp entrypoint.sh abigen:/usr/local/bin/entrypoint.sh
docker start -i abigen

docker cp abigen:/contracts/. ../../contracts
docker rm abigen