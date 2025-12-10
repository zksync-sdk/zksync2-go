#!/bin/bash
echo "Install required tools"

apt-get update && apt-get install -y jq
corepack enable
corepack prepare pnpm@latest --activate


echo "Clone the matter-labs/zksync-sso-clave-contracts repository"
git clone --depth 1 --single-branch --branch main https://github.com/matter-labs/zksync-sso-clave-contracts.git

pushd zksync-sso-clave-contracts || exit 1

echo "Install dependencies"
pnpm install

echo "Building smart contracts"
pnpm build

echo "Extracting ABI"
SRC_DIR="./artifacts-zk/src"
OUT_DIR="/abi"
# Move all *.json one level up because it has structure like ./AAFactory.sol/AAFactory.json
find "$SRC_DIR" -type f -name "*.json" ! -name "*.dbg.json" | while read -r file; do
    dest_dir=$(dirname "$file")       # /$SRC_DIR/<intermediate_path>/AccountProxy.sol
    dest_dir=$(dirname "$dest_dir")   # /$SRC_DIR/<intermediate_path>
    mv $file $dest_dir
done

# Find all *.json files excluding *.dbg.json
find "$SRC_DIR" -type f -name "*.json" ! -name "*.dbg.json" | while read -r file; do
    # Get relative path
    rel_path="${file#$SRC_DIR/}"

    # Compute destination path
    dest_path="$OUT_DIR/$rel_path" # like /abi/zksso/AAFactory.json

    # Create destination directory if needed
    mkdir -p "$(dirname "$dest_path")"

    # Output file content, extract ABI only
    jq '.abi' "$file" > $dest_path
done
