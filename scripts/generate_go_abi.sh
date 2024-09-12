#!/bin/bash

# Should be run from the project root and requires abigen from
# https://github.com/ethereum/go-ethereum/

CONTRACTS=(
    "account/CMAccount.sol/CMAccount.json"
    "manager/CMAccountManager.sol/CMAccountManager.json"
    "booking-token/BookingToken.sol/BookingToken.json"
    "utils/KYCUtils.sol/KYCUtils.json"
)

ABI_PATH="abi/contracts"
GEN_PATH="go/contracts"

for CONTRACT in "${CONTRACTS[@]}"; do
    CONTRACT_NAME=$(basename "$CONTRACT" .json)

    PACKAGE_NAME=$(echo "$CONTRACT_NAME" | awk '{print tolower($0)}')
    PACKAGE_DIR="$GEN_PATH/$PACKAGE_NAME"

    mkdir -p "$PACKAGE_DIR"

    abigen --abi "$ABI_PATH/${CONTRACT}" --pkg $PACKAGE_NAME --out="$PACKAGE_DIR/${CONTRACT_NAME}.go"
done
