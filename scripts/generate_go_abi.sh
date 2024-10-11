#!/bin/bash

# Should be run from the project root and requires abigen from
# https://github.com/ethereum/go-ethereum/

CONTRACTS=(
    "contracts/account/CMAccount.sol/CMAccount.json"
    "contracts/manager/CMAccountManager.sol/CMAccountManager.json"
    "contracts/booking-token/BookingToken.sol/BookingToken.json"
    "contracts/utils/KYCUtils.sol/KYCUtils.json"
    "@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json"
)

ABI_PATH="abi"
GEN_PATH="go/contracts"

for CONTRACT in "${CONTRACTS[@]}"; do
    CONTRACT_NAME=$(basename "$CONTRACT" .json)

    PACKAGE_NAME=$(echo "$CONTRACT_NAME" | awk '{print tolower($0)}')
    PACKAGE_DIR="$GEN_PATH/$PACKAGE_NAME"

    mkdir -p "$PACKAGE_DIR"

    abigen --abi "$ABI_PATH/${CONTRACT}" --pkg $PACKAGE_NAME --out="$PACKAGE_DIR/${CONTRACT_NAME}.go"
done
