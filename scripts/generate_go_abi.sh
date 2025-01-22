#!/bin/bash
#
# Generate Go ABI bindings for the contracts.
#
# Requires abigen from https://github.com/ethereum/go-ethereum/

set -e

### Variables ###

GO_ETH_VERSION="v1.14.12"
GO_MODULE_NAME="github.com/chain4travel/camino-messenger-contracts/go/contracts"

GEN_PATH="go/contracts"
ARTIFACTS_PATH="artifacts"
ARTIFACTS=(
    "contracts/account/CMAccount.sol/CMAccount.json"
    "contracts/manager/CMAccountManager.sol/CMAccountManager.json"
    "contracts/booking-token/BookingToken.sol/BookingToken.json"
    "contracts/booking-token/BookingTokenOperator.sol/BookingTokenOperator.json"
    "@openzeppelin/contracts/token/ERC20/ERC20.sol/ERC20.json"
    "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol/ERC1967Proxy.json"
)

### END Variables ###

# Parse command-line arguments
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
    --skip-yarn)
        SKIP_YARN=true
        shift
        ;;
    --help)
        echo "Usage: $0 [--skip-yarn]"
        exit 0
        ;;
    *)
        echo "Unknown option: $1"
        exit 1
        ;;
    esac
done

echo "Starting Go ABI generator..."
SECONDS=0
start=$(date +%s.%N)

# Set workdir
if [ -z "${WORKDIR}" ]; then
    # camino-messenger-bot root folder
    WORKDIR=$(
        cd "$(dirname "${BASH_SOURCE[0]}")"
        cd .. && pwd
    )
fi
echo "cd $WORKDIR"
cd "$WORKDIR"

# Text colors
WHITE='\033[1;37m'
BLACK='\033[0;30m'
NC='\033[0m' # No Color

if [ -z "$SKIP_YARN" ]; then
    echo -e -n "Cleaning node modules..."
    rm -rf node_modules && echo -e "${WHITE}done!${NC}"

    echo -e "Install dependencies..."
    yarn install --frozen-lockfile

    echo -e -n "Cleaning artifacts directory..."
    yarn hardhat clean >/dev/null && echo -e "${WHITE}done!${NC}"

    echo -e -n "Compiling contracts..."
    yarn hardhat compile --force >/dev/null && echo -e "${WHITE}done!${NC}"
else
    echo -e "Skipping yarn install and hardhat compile..."
fi

echo "Generating Go ABI bindings..."
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

# Go to the Go module directory
cd "$GEN_PATH"

# Clean go.mod and go.sum
echo -e "Cleaning go.mod and go.sum..."
rm -rfv go.mod go.sum

# Run go mod init
echo -e "Running ${WHITE}go mod init ${GO_MODULE_NAME} ${NC}..."
go mod init ${GO_MODULE_NAME}

# Get ethereum/go-ethereum
echo -e "Running ${WHITE}go get github.com/ethereum/go-ethereum${GO_ETH_VERSION} ${NC}..."
go get github.com/ethereum/go-ethereum@${GO_ETH_VERSION}

# Run go mod tidy
echo -e "Running ${WHITE}go mod tidy ${NC}..."
go mod tidy -v

end=$(date +%s.%N)
runtime=$(echo "$end - $start" | bc -l | awk '{printf "%.2f", $1}')
echo -e "Finished in ${runtime}s"
