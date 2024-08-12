/**
 * @dev Calculates storage location for ERC-7201
 *
 * Run this script with the namespace ID as the CLI argument.
 *
 * Ex:
 *
 * ❯ node examples/erc7201.js camino.messenger.storage.CMAccount
 * NAMESPACE ID: camino.messenger.storage.CMAccount
 * Storage Location: 0x0c7b73796c7cc89b9f849b9056a93200eba741881e57a1b03b9bedb2c0e07100
 */

const ethers = require("ethers");

// Check if a namespace ID is provided
if (process.argv.length < 3) {
    console.error("Please provide the namespace ID as a CLI argument.");
    process.exit(1);
}

// Get the namespace ID from CLI arguments
const NAMESPACE_ID = process.argv[2];

//const NAMESPACE_ID = "camimo.messenger.storage.PartnerConfiguration";

// Calculate storage location according to ERC-7201 specification.
//
// Same calculation on Solidity:
// keccak256(abi.encode(uint256(keccak256("camino.messenger.storage.PartnerConfiguration")) - 1)) & ~bytes32(uint256(0xff));

const hash = ethers.keccak256(ethers.toUtf8Bytes(NAMESPACE_ID));
const hashAsBigInt = BigInt(hash);
const subtractedHash = hashAsBigInt - BigInt(1);
const abiEncoded = ethers.AbiCoder.defaultAbiCoder().encode(["uint256"], [subtractedHash]);
const finalHash = ethers.keccak256(abiEncoded);
const finalHashAsBigInt = BigInt(finalHash);

// Equivalent to 0xffff...ff00 with last byte zeroed out. Used by ERC-7201 to align
// the storage location to 256 bytes.
const mask = ~BigInt(0xff);

const maskedFinalHash = finalHashAsBigInt & mask;

const storageLocation = ethers.toBeHex(maskedFinalHash);

console.log("NAMESPACE ID:", NAMESPACE_ID);
console.log("Storage Location:", storageLocation);
