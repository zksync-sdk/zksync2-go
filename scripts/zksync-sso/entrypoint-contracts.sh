#!/bin/ash

SRC_DIR="/abi"
OUT_DIR="/contracts/zksyncsso"

# Find all *.json files excluding *.dbg.json
find "$SRC_DIR" -type f -name "*.json" ! -name "*.dbg.json" | while read -r file; do
    # Get relative path
    rel_path="${file#$SRC_DIR/}"  # like zksso/interfaces/IGuardianRecoveryValidator.json

    file_name="${rel_path##*/}"       # IGuardianRecoveryValidator.json - file name
    file_name="${file_name%.*}"        # IGuardianRecoveryValidator - strip extension
    file_name_lowercase=$(echo "$file_name" | tr '[:upper:]' '[:lower:]') # iguardianrecoveryvalidator - to lowercase

    base_dir=$(dirname "$rel_path")  # zksso/interfaces/

    out_go_dir="$OUT_DIR"/"$base_dir"/"$file_name_lowercase" # zksso/interfaces/iguardianrecoveryvalidator
    out_go_filepath="$out_go_dir"/"$file_name_lowercase".go  # zksso/interfaces/iguardianrecoveryvalidator/iguardianrecoveryvalidator.go

    mkdir -p "$out_go_dir"

    abigen  \
      --abi "$file" \
      --out "$out_go_filepath" \
      --pkg "$file_name_lowercase" \
      --type "$file_name"

done
