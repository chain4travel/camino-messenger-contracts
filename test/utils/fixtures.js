/**
 * @dev Fixtures
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { ethers, upgrades } = require("hardhat");

const developerFeeBp = 100;

async function setupSigners() {
    const [
        managerAdmin,
        managerPauser,
        managerUpgrader,
        managerVersioner,
        cmAccountAdmin,
        cmAccountUpgrader,
        cmServiceAdmin,
        developerWallet,
        developerWalletAdmin,
        feeAdmin,
        chequeOperator,
        depositor,
        withdrawer,
        btAdmin,
        btUpgrader,
        registryAdmin,
        otherAccount1,
        otherAccount2,
        otherAccount3,
    ] = await ethers.getSigners();

    signers = {
        managerAdmin,
        managerPauser,
        managerUpgrader,
        managerVersioner,
        cmAccountAdmin,
        cmAccountUpgrader,
        cmServiceAdmin,
        developerWallet,
        developerWalletAdmin,
        feeAdmin,
        chequeOperator,
        depositor,
        withdrawer,
        btAdmin,
        btUpgrader,
        registryAdmin,
        otherAccount1,
        otherAccount2,
        otherAccount3,
    };
}

async function deployCMAccountManagerFixture() {
    // Set up signers
    await setupSigners();

    const CMAccountManager = await ethers.getContractFactory("CMAccountManager");
    const cmAccountManager = await upgrades.deployProxy(
        CMAccountManager,
        [
            signers.managerAdmin.address,
            signers.managerPauser.address,
            signers.managerUpgrader.address,
            signers.managerVersioner.address,
            signers.developerWallet.address,
            developerFeeBp,
        ],
        { kind: "uups" },
    );
    return { cmAccountManager };
}

async function deployCMAccountImplFixture() {
    const CMAccount = await ethers.getContractFactory("CMAccount");
    const cmAccountImpl = await CMAccount.deploy();
    await cmAccountImpl.waitForDeployment();

    return { cmAccountImpl };
}

async function deployCMAccountManagerWithCMAccountImplFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);
    const { cmAccountImpl } = await loadFixture(deployCMAccountImplFixture);

    const cmAccountImplAddress = await cmAccountImpl.getAddress();

    await cmAccountManager.grantRole(await cmAccountManager.VERSIONER_ROLE(), signers.managerVersioner.address);
    await cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(cmAccountImplAddress);

    return { cmAccountManager, cmAccountImplAddress };
}

async function deployAndConfigureAllFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
        deployCMAccountManagerWithCMAccountImplFixture,
    );

    await cmAccountManager.grantRole(
        await cmAccountManager.DEVELOPER_WALLET_ADMIN_ROLE(),
        signers.developerWalletAdmin.address,
    );
    await cmAccountManager.grantRole(await cmAccountManager.FEE_ADMIN_ROLE(), signers.feeAdmin.address);

    // Deploy BookingToken
    const BookingToken = await ethers.getContractFactory("BookingToken");
    const bookingToken = await upgrades.deployProxy(
        BookingToken,
        [await cmAccountManager.getAddress(), signers.btAdmin.address, signers.btUpgrader.address],
        { kind: "uups" },
    );

    // Set BookingToken address on the manager
    await cmAccountManager.connect(signers.managerVersioner).setBookingToken(bookingToken.getAddress());

    // Get pre fund amount
    const prefundAmount = await cmAccountManager.getPrefundAmount();

    const tx = await cmAccountManager.createCMAccount(
        signers.cmAccountAdmin.address,
        signers.cmAccountUpgrader.address,
        { value: prefundAmount },
    );

    const receipt = await tx.wait();

    // Parse event to get the CMAccount address (this is the UUPS proxy address)
    const event = receipt.logs.find((log) => {
        try {
            return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
        } catch (e) {
            return false;
        }
    });

    const parsedEvent = cmAccountManager.interface.parseLog(event);
    const cmAccountAddress = parsedEvent.args.account;

    // Get the CMAccount instance at the address
    const cmAccount = await ethers.getContractAt("CMAccount", cmAccountAddress);

    return { cmAccountManager, cmAccount, bookingToken, prefundAmount };
}

async function deployCMAccountWithDepositFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccount, bookingToken, prefundAmount } =
        await loadFixture(deployAndConfigureAllFixture);

    // Grant withdrawer role
    const WITHDRAWER_ROLE = await cmAccount.WITHDRAWER_ROLE();
    await cmAccount.connect(signers.cmAccountAdmin).grantRole(WITHDRAWER_ROLE, signers.withdrawer.address);

    const depositAmount = ethers.parseEther("1");

    const depositTx = {
        to: cmAccount.getAddress(),
        value: depositAmount,
    };

    const txResponse = await signers.depositor.sendTransaction(depositTx);
    await txResponse.wait();

    return { cmAccountManager, cmAccount, bookingToken, prefundAmount };
}

async function deployBookingTokenFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccount, bookingToken, prefundAmount } = await loadFixture(
        deployCMAccountWithDepositFixture,
    );

    // Supplier CMAccount with deposit
    const supplierCMAccount = cmAccount;

    // Create distributor CMAccount
    const tx = await cmAccountManager.createCMAccount(
        signers.cmAccountAdmin.address,
        signers.cmAccountUpgrader.address,
        { value: prefundAmount },
    );

    const receipt = await tx.wait();

    // Parse event to get the CMAccount address (this is the UUPS proxy address)
    const event = receipt.logs.find((log) => {
        try {
            return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
        } catch (e) {
            return false;
        }
    });

    const parsedEvent = cmAccountManager.interface.parseLog(event);
    const distributorCMAccountAddress = parsedEvent.args.account;

    // Get the CMAccount instance at the address
    const distributorCMAccount = await ethers.getContractAt("CMAccount", distributorCMAccountAddress);

    // Deposit funds to distributor CMAccount
    const depositAmount = ethers.parseEther("1");
    const depositTx = {
        to: distributorCMAccount.getAddress(),
        value: depositAmount,
    };
    const txResponse = await signers.depositor.sendTransaction(depositTx);
    await txResponse.wait();

    return { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, prefundAmount };
}

async function deployAndConfigureAllWithRegisteredServicesFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

    // Grant SERVICE_REGISTRY_ADMIN_ROLE
    const SERVICE_REGISTRY_ADMIN_ROLE = await cmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();
    cmAccountManager
        .connect(signers.managerAdmin)
        .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.registryAdmin.address);

    // Services to register
    const serviceName1 = "cmp.service.accommodation.v1.AccommodationSearchService";
    const serviceHash1 = ethers.keccak256(ethers.toUtf8Bytes(serviceName1));

    const serviceName2 = "cmp.service.accommodation.v2.AccommodationSearchService";
    const serviceHash2 = ethers.keccak256(ethers.toUtf8Bytes(serviceName2));

    const serviceName3 = "cmp.service.accommodation.v3.AccommodationSearchService";
    const serviceHash3 = ethers.keccak256(ethers.toUtf8Bytes(serviceName3));

    const services = {
        serviceName1,
        serviceHash1,
        serviceName2,
        serviceHash2,
        serviceName3,
        serviceHash3,
    };

    // Register services
    cmAccountManager.connect(signers.registryAdmin).registerService(serviceName1);
    cmAccountManager.connect(signers.registryAdmin).registerService(serviceName2);
    cmAccountManager.connect(signers.registryAdmin).registerService(serviceName3);

    // Get the SERVICE_ADMIN_ROLE
    const SERVICE_ADMIN_ROLE = await cmAccount.SERVICE_ADMIN_ROLE();

    // Grant SERVICE_ADMIN_ROLE to otherAccount1
    cmAccount.connect(signers.cmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.cmServiceAdmin.address);

    return { cmAccountManager, cmAccount, services };
}

module.exports = {
    setupSigners,
    developerFeeBp,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployCMAccountWithDepositFixture,
    deployBookingTokenFixture,
    deployAndConfigureAllWithRegisteredServicesFixture,
};
