require("@nomicfoundation/hardhat-toolbox");
const { types } = require("hardhat/config");

const ACCOUNT_SCOPE = scope("account", "CM Account Tasks");

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
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addParam("role", "Role to grant. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to grant role to")
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "grantRole");
    });

ACCOUNT_SCOPE.task("role:revoke", "Revoke role")
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addParam("role", "Role to grant. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to revoke role to")
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "revokeRole");
    });

ACCOUNT_SCOPE.task("role:has", "Check if address has role")
    .addParam("role", "Role to check. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to check")
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
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

ACCOUNT_SCOPE.task("bot:add", "Add bot to the CMAccount")
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
    .addParam("bot", "Bot address")
    .addOptionalParam(
        "gasMoney",
        "Gas money in CAM. This amount will be transferred from the CMAccount to the bot address (Ex: 1 or 0.1)",
        "0",
        types.string,
    )
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);
        console.log("Bot:", taskArgs.bot);
        console.log(
            "Gas:",
            taskArgs.gasMoney,
            "(This amount will be transferred from the CMAccount to the bot address)",
        );

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Adding bot to CMAccount...");
            console.log("Signer:", signer.address);

            const gasMoney = ethers.parseEther(taskArgs.gasMoney);

            const tx = await cmAccount.connect(signer).addMessengerBot(taskArgs.bot, gasMoney);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("bot:remove", "Remove bot from the CMAccount")
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
    .addParam("bot", "Bot address")
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);
        console.log("Bot:", taskArgs.bot);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Removing bot from CMAccount...");
            console.log("Signer:", signer.address);

            const tx = await cmAccount.connect(signer).removeMessengerBot(taskArgs.bot);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

ACCOUNT_SCOPE.task("bot:list", "List all bots from CMAccount")
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
    .setAction(async (taskArgs, hre) => {
        console.log("CMAccount:", taskArgs.cmAccount, "\n");

        console.log("📢 A bot is an address that has been granted some special roles on the CMAccount.");

        const role1 = "CHEQUE_OPERATOR_ROLE";
        console.log("\n🤖", role1, "(Can sign cheques that are valid for the CMAccount)");
        console.log("======================================================");
        await hre.run({ scope: "account", task: "role:members" }, { role: role1, cmAccount: taskArgs.cmAccount });

        const role2 = "BOOKING_OPERATOR_ROLE";
        console.log("\n🤖", role2, "(Can mint and buy Booking Tokens for the CMAccount)");
        console.log("======================================================");
        await hre.run({ scope: "account", task: "role:members" }, { role: role2, cmAccount: taskArgs.cmAccount });

        const role3 = "GAS_WITHDRAWER_ROLE";
        console.log("\n🤖", role3, "(Can withdraw gas from the CMAccount)");
        console.log("======================================================");
        await hre.run({ scope: "account", task: "role:members" }, { role: role3, cmAccount: taskArgs.cmAccount });
    });

ACCOUNT_SCOPE.task("wanted:add", "Add wanted service to CMAccount")
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
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

ACCOUNT_SCOPE.task("upgrade", "Upgrade CMAccount to latest implementation")
    .addOptionalParam("privateKey", "Private key to use", process.env.CMACCOUNT_PK)
    .addOptionalParam("cmAccount", "CMAccount address", process.env.CMACCOUNT_ADDRESS)
    .setAction(async (taskArgs, hre) => {
        const cmAccount = await getCMAccount(taskArgs.cmAccount);
        console.log("CMAccount:", taskArgs.cmAccount);

        // Get new implementation
        const manager = await getManager(hre);
        const implementation = await manager.getAccountImplementation();
        console.log("New Implementation on the Manager:", implementation);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Upgrading CMAccount implementation...");
            console.log("Signer:", signer.address);
            const tx = await cmAccount.connect(signer).upgradeToAndCall(implementation, "0x");
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, cmAccount);
        }
    });

module.exports = {};
