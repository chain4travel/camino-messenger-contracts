#!/bin/bash

# Should be run from the project root and requires abigen from
# https://github.com/ethereum/go-ethereum/

ARTIFACTS_PATH="artifacts"
ARTIFACTS=(
    "contracts/account/CMAccount.sol/CMAccount.json"
    "contracts/manager/CMAccountManager.sol/CMAccountManager.json"
    "contracts/booking-token/BookingTokenV2.sol/BookingTokenV2.json"
    "contracts/booking-token/BookingTokenOperator.sol/BookingTokenOperator.json"
    "contracts/utils/KYCUtils.sol/KYCUtils.json"
    "@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json"
    "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol/ERC1967Proxy.json"
)

GEN_PATH="go/contracts"

for CONTRACT in "${ARTIFACTS[@]}"; do

    # Contract base name
    CONTRACT_NAME=$(basename "$CONTRACT" .json)

    # Generate package name from the contract name
    PACKAGE_NAME=$(echo "$CONTRACT_NAME" | awk '{print tolower($0)}')
    PACKAGE_DIR="$GEN_PATH/$PACKAGE_NAME"

    # Create package directory
    mkdir -p "$PACKAGE_DIR"

    # Create temporary bin file from the bytecode in JSON
    TMP_BIN_FILE=$(mktemp)
    jq -r '.bytecode' "$ARTIFACTS_PATH/$CONTRACT" >"$TMP_BIN_FILE"

    # Create temporary ABI file from JSON
    TMP_ABI_FILE=$(mktemp)
    jq -r '.abi' "$ARTIFACTS_PATH/$CONTRACT" >"$TMP_ABI_FILE"

    abigen --abi "$TMP_ABI_FILE" --bin "$TMP_BIN_FILE" --pkg $PACKAGE_NAME --out="$PACKAGE_DIR/${CONTRACT_NAME}.go"

    # Clean up temporary file
    rm "$TMP_BIN_FILE"
    rm "$TMP_ABI_FILE"
done
