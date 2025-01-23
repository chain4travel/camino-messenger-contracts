// Run with --network columbus on SHA a53b8fc03973025f4ed10edb46aa4d2f0c3e76ee
// Check 00_validate_01_BookingTokenRefactor.js for details
const { ethers, upgrades } = require("hardhat");

async function main() {
    // Existing BookingToken proxy address on Columbus
    const existingAddress = "0xe55E387F5474a012D1b048155E25ea78C7DBfBBC";

    // Current version of the BookingToken contract
    const OldBookingTokenV2 = await ethers.getContractFactory("BookingTokenV2");

    // Force import the existing deployment from chain using the old BookingTokenV2
    const result = await upgrades.forceImport(existingAddress, OldBookingTokenV2, { kind: "uups" });

    console.log(result);

    console.log(`Successfully force-imported contract at ${existingAddress}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
