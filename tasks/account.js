require("@nomicfoundation/hardhat-toolbox");
const { types } = require("hardhat/config");

const ACCOUNT_SCOPE = scope("account", "CMAccount Tasks");

// TODO: Get private key from .env or hardhat vars

function getAddressesForNetwork(hre) {
    let addresses;

    if (hre.network.name === "columbus") {
        console.log("Running on columbus");
        addresses = require("../ignition/deployments/chain-501/deployed_addresses.json");
    } else if (hre.network.name === "camino") {
        console.log("Running on camino");
        addresses = require("../ignition/deployments/chain-500/deployed_addresses.json");
    } else if (hre.network.name === "localhost") {
        console.log("Running on localhost");
        addresses = require("../ignition/deployments/chain-31337/deployed_addresses.json");
    } else {
        throw new Error(`Unsupported network: ${hre.network.name}`);
    }

    return addresses;
}

async function getManager(hre) {
    const addresses = getAddressesForNetwork(hre);
    return await ethers.getContractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"]);
}

async function getCMAccount(cmAccountAddress) {
    return await ethers.getContractAt("CMAccount", cmAccountAddress);
}

async function handleRoles(taskArgs, hre, action) {
    const cmAccount = await getCMAccount(taskArgs.cmAccount);
    console.log("CMAccount:", taskArgs.cmAccount);

    try {
        const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

        console.log(
            `${action === "grantRole" ? "Granting" : "Revoking"} role ${taskArgs.role} for address ${taskArgs.address}...`,
        );

        const role = await cmAccount.connect(signer)[taskArgs.role]();
        const tx = await cmAccount.connect(signer)[action](role, taskArgs.address);
        const txReceipt = await tx.wait();
        console.log("Tx:", txReceipt.hash);
    } catch (error) {
        handleTransactionError(error, cmAccount);
    }
}

function handleTransactionError(error, contract) {
    console.error("❌ Transaction failed!");

    if (error.data?.message) {
        console.error(`Reason: ${error.data.message}`);
    } else if (error.data && contract) {
        const decodedError = contract.interface.parseError(error.data);
        console.error("Message:", error.message);
        console.error(`Reason: ${decodedError?.name} (${decodedError?.args})`);
    } else if (error.message?.includes("[taskArgs.role] is not a function")) {
        console.error("Reason: CMAccount does not have this role.");
    } else if (error.message) {
        console.error("Message:", error.message);
    } else {
        // General error logging
        console.error("An unexpected error occurred.");
        console.error("Error:", error);
    }
}

ACCOUNT_SCOPE.task("role:grant", "Grant role")
    .addParam("privateKey", "Private key to use")
    .addParam("role", "Role to grant. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to grant role to")
    .addParam("cmAccount", "CMAccount address")
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "grantRole");
    });

ACCOUNT_SCOPE.task("role:revoke", "Revoke role")
    .addParam("privateKey", "Private key to use")
    .addParam("role", "Role to grant. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to revoke role to")
    .addParam("cmAccount", "CMAccount address")
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "revokeRole");
    });

ACCOUNT_SCOPE.task("role:has", "Check if address has role")
    .addParam("role", "Role to check. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to check")
    .addParam("cmAccount", "CMAccount address")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);

        try {
            console.log("Running on", hre.network.name);
            const role = await cmAccount[taskArgs.role]();
            const hasRole = await cmAccount.hasRole(role, taskArgs.address);

            console.log(`Address ${taskArgs.address} ${hasRole ? "has" : "does not have"} role ${taskArgs.role}`);
            console.log(`${hasRole ? "🟢" : "🔴"}`, hasRole);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("role:members", "List role members")
    .addParam("role", "Role to list. Ex: SERVICE_ADMIN_ROLE")
    .addParam("cmAccount", "CMAccount address")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);
        console.log("Role:", taskArgs.role);

        try {
            const role = await cmAccount[taskArgs.role]();
            const count = await cmAccount.getRoleMemberCount(role);
            console.log("Total Members:", count);

            // Iterate over the members of the role
            const members = [];
            for (let i = 0; i < count; i++) {
                const member = await cmAccount.getRoleMember(role, i);
                members.push(member);
            }
            console.log(members);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("create", "Create CMAccount")
    .addParam("privateKey", "Private key to use")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        try {
            // Get signer from private key
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            console.log("Creating CMAccount...");
            console.log("Signer:", signer.address);
            const tx = await manager
                .connect(signer)
                .createCMAccount(signer.address, signer.address, { value: ethers.parseEther("100") });

            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);

            // Parse event to get the CMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return manager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = manager.interface.parseLog(event);
            const cmAccountAddress = parsedEvent.args.account;

            console.log("CMAccount Address:", cmAccountAddress);
        } catch (error) {
            handleTransactionError(error, manager);
        }
    });

ACCOUNT_SCOPE.task("wanted:add", "Add wanted service to CMAccount")
    .addParam("privateKey", "Private key to use")
    .addParam("cmAccount", "CMAccount address")
    .addParam("serviceName", "Name of service to add")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);
        console.log("Service Name:", taskArgs.serviceName);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            console.log("Adding service to CMAccount...");
            console.log("Signer:", signer.address);

            const tx = await cmAccount.connect(signer).addWantedServices([taskArgs.serviceName]);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("wanted:remove", "Remove wanted service from CMAccount")
    .addParam("privateKey", "Private key to use")
    .addParam("cmAccount", "CMAccount address")
    .addParam("serviceName", "Name of service to remove")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);
        console.log("Service Name:", taskArgs.serviceName);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            console.log("Removing service from CMAccount...");
            console.log("Signer:", signer.address);

            const tx = await cmAccount.connect(signer).removeWantedServices([taskArgs.serviceName]);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("wanted:list", "List all wanted service from CMAccount")
    .addParam("cmAccount", "CMAccount address")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);

        try {
            console.log("Listing all wanted services from CMAccount...");

            const wantedServices = await cmAccount.getWantedServices();
            console.log("Wanted Services:");
            console.log(wantedServices);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:add", "Add supported service to CMAccount")
    .addParam("privateKey", "Private key to use")
    .addParam("cmAccount", "CMAccount address")
    .addParam("serviceName", "Name of service to add")
    .addParam("fee", "Fee of the service in aCAM (wei in ETH terminology)")
    .addParam("restrictedRate", "Restricted rate of the service", false, types.boolean)
    .addOptionalParam("capabilities", "Capabilities of the service, comma separated (optional)")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);
        console.log("Service Name:", taskArgs.serviceName);
        console.log("Fee:", taskArgs.fee);
        console.log("Restricted Rate:", taskArgs.restrictedRate);
        console.log("Capabilities:", capabilities);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            const capabilities = taskArgs.capabilities ? taskArgs.capabilities.split(",") : [];

            console.log("Adding service to CMAccount...");
            console.log("Signer:", signer.address);

            const tx = await cmAccount
                .connect(signer)
                .addService(taskArgs.serviceName, taskArgs.fee, taskArgs.restrictedRate, capabilities);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:remove", "Remove wanted service from CMAccount")
    .addParam("privateKey", "Private key to use")
    .addParam("cmAccount", "CMAccount address")
    .addParam("serviceName", "Name of service to remove")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);
        console.log("Service Name:", taskArgs.serviceName);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            console.log("Removing service from CMAccount...");
            console.log("Signer:", signer.address);

            const tx = await cmAccount.connect(signer).removeService(taskArgs.serviceName);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:list", "List supported services from CMAccount")
    .addParam("cmAccount", "CMAccount address")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);

        try {
            console.log("Listing all supported services from CMAccount...");

            const supportedServices = await cmAccount.getSupportedServices();
            const serviceNames = supportedServices[0];
            const serviceDetails = supportedServices[1];
            if (serviceNames.length > 0) {
                console.log("Supported Services:");
                for (let i = 0; i < serviceNames.length; i++) {
                    console.log(`📦 ${serviceNames[i]}`);
                    const feeACAM = serviceDetails[i][0];
                    const feeNCAM = ethers.formatUnits(serviceDetails[i][0], "gwei");
                    const feeCAM = ethers.formatEther(serviceDetails[i][0]);
                    console.log(`\t💰 Fee: ${feeNCAM} nCAM (${feeACAM} aCAM or ${feeCAM} CAM)`);
                    console.log(`\t🔒 Restricted Rate: ${serviceDetails[i][1]} ${serviceDetails[i][1] ? "✅" : "❌"}`);

                    for (let j = 0; j < serviceDetails[i][2].length; j++) {
                        console.log(`\t🔧 ${serviceDetails[i][2][j]}`);
                    }
                }
            } else {
                console.log("🛑 CM Account does not have any supported services!");
            }
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

module.exports = {};
