#!/bin/bash
#
# Should be run from the project root and requires abigen from
# https://github.com/ethereum/go-ethereum/
#
# Note that if you are switching between branches with different yarn.lock file you
# may need to clean node_modules and yarn install --freeze-lockfile before running
# this script.

set -e

echo "Generating Go ABI bindings..."
SECONDS=0
start=$(date +%s.%N)

# Text colors
WHITE='\033[1;37m'
BLACK='\033[0;30m'
NC='\033[0m' # No Color

ARTIFACTS_PATH="artifacts"
ARTIFACTS=(
    "contracts/account/CMAccount.sol/CMAccount.json"
    "contracts/manager/CMAccountManager.sol/CMAccountManager.json"
    #"contracts/booking-token/BookingTokenV2.sol/BookingTokenV2.json"
    "contracts/booking-token/BookingToken.sol/BookingToken.json"
    "contracts/booking-token/BookingTokenOperator.sol/BookingTokenOperator.json"
    "contracts/utils/KYCUtils.sol/KYCUtils.json"
    "@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json"
    "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol/ERC1967Proxy.json"
)

GEN_PATH="go/contracts"

echo -e -n "Cleaning artifacts directory..."
yarn hardhat clean >/dev/null && echo -e "${WHITE}done!${NC}"

echo -e -n "Compiling contracts..."
yarn hardhat compile --force >/dev/null && echo -e "${WHITE}done!${NC}"

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

    echo -e "Generating $PACKAGE_DIR/${WHITE}${CONTRACT_NAME}.go${NC}"
    ABIGEN_CMD="abigen --abi "$TMP_ABI_FILE" --bin "$TMP_BIN_FILE" --pkg $PACKAGE_NAME --out="$PACKAGE_DIR/${CONTRACT_NAME}.go""
    echo -e "  └─ ${BLACK}$ABIGEN_CMD${NC}"
    $ABIGEN_CMD

    # Clean up temporary file
    rm "$TMP_BIN_FILE"
    rm "$TMP_ABI_FILE"
done

end=$(date +%s.%N)
runtime=$(echo "$end - $start" | bc -l | awk '{printf "%.2f", $1}')
echo -e "Finished in ${runtime}s"
