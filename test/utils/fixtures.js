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

// Deploy NullUSD
async function deployNullUSDFixture() {
    await setupSigners();

    const NullUSD = await ethers.getContractFactory("NullUSD");
    const nullUSD = await NullUSD.deploy();

    const nullUSDDecimals = await nullUSD.decimals();

    return { nullUSD, nullUSDDecimals };
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
    const BookingTokenOperator = await ethers.getContractFactory("BookingTokenOperator");
    const bookingTokenOperator = await BookingTokenOperator.deploy();
    const CMAccount = await ethers.getContractFactory("CMAccount", {
        libraries: { BookingTokenOperator: await bookingTokenOperator.getAddress() },
    });
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

    const { nullUSD, nullUSDDecimals } = await loadFixture(deployNullUSDFixture);

    await cmAccountManager.grantRole(
        await cmAccountManager.DEVELOPER_WALLET_ADMIN_ROLE(),
        signers.developerWalletAdmin.address,
    );
    await cmAccountManager.grantRole(await cmAccountManager.FEE_ADMIN_ROLE(), signers.feeAdmin.address);

    // Set Service Fee Token Address
    await cmAccountManager.grantRole(await cmAccountManager.SERVICE_FEE_TOKEN_ADMIN_ROLE(), signers.feeAdmin.address);
    await cmAccountManager.connect(signers.feeAdmin).setServiceFeeToken(await nullUSD.getAddress());

    // Deploy BookingToken

    const BookingToken = await ethers.getContractFactory("BookingToken");
    const bookingToken = await upgrades.deployProxy(
        BookingToken,
        [await cmAccountManager.getAddress(), signers.btAdmin.address, signers.btUpgrader.address],
        { kind: "uups" },
    );

    // Set BookingToken address on the manager
    await cmAccountManager.connect(signers.managerVersioner).setBookingTokenAddress(bookingToken.getAddress());

    // Get pre fund amount
    const prefundAmount = await cmAccountManager.getPrefundAmount();

    // Approve allowance for service fee prefund amount
    await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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

    return {
        cmAccountManager,
        cmAccount,
        bookingToken,
        prefundAmount,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployCMAccountWithDepositFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccount, bookingToken, prefundAmount, nullUSD, nullUSDDecimals } =
        await loadFixture(deployAndConfigureAllFixture);

    // Grant withdrawer role
    const WITHDRAWER_ROLE = await cmAccount.WITHDRAWER_ROLE();
    await cmAccount.connect(signers.cmAccountAdmin).grantRole(WITHDRAWER_ROLE, signers.withdrawer.address);

    // Deposit CAM
    const depositAmount = ethers.parseEther("1");

    const depositTx = {
        to: cmAccount.getAddress(),
        value: depositAmount,
    };

    const txResponse = await signers.depositor.sendTransaction(depositTx);
    await txResponse.wait();

    // Deposit service fee token
    await nullUSD.transfer(cmAccount.getAddress(), ethers.parseUnits("1", nullUSDDecimals));

    return {
        cmAccountManager,
        cmAccount,
        bookingToken,
        prefundAmount,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployBookingTokenFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccount, bookingToken, prefundAmount, nullUSD, nullUSDDecimals } = await loadFixture(
        deployCMAccountWithDepositFixture,
    );

    // Supplier CMAccount with deposit
    const supplierCMAccount = cmAccount;

    // Approve allowance for service fee prefund amount
    await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);

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

    // Deposit service fee token
    await nullUSD.transfer(distributorCMAccount.getAddress(), ethers.parseUnits("1", nullUSDDecimals));

    return {
        cmAccountManager,
        supplierCMAccount,
        distributorCMAccount,
        bookingToken,
        prefundAmount,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployBookingTokenWithNullUSDFixture() {
    // Set up signers
    await setupSigners();

    const {
        cmAccountManager,
        supplierCMAccount,
        distributorCMAccount,
        bookingToken,
        prefundAmount,
        nullUSD,
        nullUSDDecimals,
    } = await loadFixture(deployBookingTokenFixture);

    // Fund NullUSD to the CM accounts
    const fundAmount = ethers.parseEther("1000");
    await nullUSD.transfer(await supplierCMAccount.getAddress(), fundAmount);
    await nullUSD.transfer(await distributorCMAccount.getAddress(), fundAmount);

    return {
        cmAccountManager,
        supplierCMAccount,
        distributorCMAccount,
        bookingToken,
        prefundAmount,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployCancellationSupportFixture() {
    // Set up signers
    await setupSigners();

    const {
        cmAccountManager,
        supplierCMAccount,
        distributorCMAccount,
        bookingToken,
        prefundAmount,
        nullUSD,
        nullUSDDecimals,
    } = await loadFixture(deployBookingTokenWithNullUSDFixture);

    // Set accounts
    const otherBookingOperator = signers.otherAccount1;
    const supplierBookingOperator = signers.otherAccount2;
    const distributorBookingOperator = signers.otherAccount3;

    // Set BOOKING_OPERATOR_ROLE
    // Supplier
    const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
    await supplierCMAccount
        .connect(signers.cmAccountAdmin)
        .grantRole(BOOKING_OPERATOR_ROLE, supplierBookingOperator.address);

    // Distributor
    await distributorCMAccount
        .connect(signers.cmAccountAdmin)
        .grantRole(BOOKING_OPERATOR_ROLE, distributorBookingOperator.address);

    // Mint BOOKING TOKEN with NATIVE PAYMENT -----------------------------------------------------------------------

    const tokenURI = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;
    const price = ethers.parseEther("0.05");

    await supplierCMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorCMAccount.getAddress(), // Reserved for
        tokenURI, // URI
        expirationTimestamp, // Expiration of the reservation
        price, // Price of token in wei
        ethers.ZeroAddress, // paymentToken: zero address, means native coin
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Token with ID 0 minted with native payment
    const tokenWithNativePayment = 0n;

    // Buy the token
    await distributorCMAccount
        .connect(distributorBookingOperator)
        .buyBookingToken(tokenWithNativePayment, price, ethers.ZeroAddress);

    // Mint BOOKING TOKEN with NULLUSD PAYMENT------------------------------------------------------------------------

    const tokenURI2 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp2 = Math.floor(Date.now() / 1000) + 120;
    const price2 = ethers.parseEther("99.95");

    await supplierCMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorCMAccount.getAddress(), // Reserved for
        tokenURI2, // URI
        expirationTimestamp2, // Expiration of the reservation
        price2, // Price of token in wei
        nullUSD.getAddress(), // paymentToken
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Token with ID 1 minted with NullUSD payment
    const tokenWithNullUSDPayment = 1n;

    // Buy the token
    await distributorCMAccount
        .connect(distributorBookingOperator)
        .buyBookingToken(tokenWithNullUSDPayment, price2, await nullUSD.getAddress());

    // Mint BOOKING TOKEN without buying -----------------------------------------------------------------------------

    const tokenURI3 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp3 = Math.floor(Date.now() / 1000) + 600;
    const price3 = ethers.parseEther("0.95");

    await supplierCMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorCMAccount.getAddress(), // Reserved for
        tokenURI3, // URI
        expirationTimestamp3, // Expiration of the reservation
        price3, // Price of token in wei
        ethers.ZeroAddress, // paymentToken
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Token with ID 2 minted without buying
    const tokenWithoutBuying = 2n;

    // Mint BOOKING TOKEN with passed expiration ---------------------------------------------------------------------

    const tokenURI4 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

    // get block time from the chain
    const block = await ethers.provider.getBlock("latest");

    const expirationTimestamp4 = block.timestamp + 70; // min expiration time diff is 60
    const price4 = ethers.parseEther("0.95");

    await supplierCMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorCMAccount.getAddress(), // Reserved for
        tokenURI4, // URI
        expirationTimestamp4, // Expiration of the reservation
        price4, // Price of token in wei
        ethers.ZeroAddress, // paymentToken
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Advance time to after the expiration
    await network.provider.send("evm_increaseTime", [expirationTimestamp4 - block.timestamp + 10]);
    await network.provider.send("evm_mine");

    // Token with ID 3 minted with passed expiration
    const tokenWithPassedExpiration = 3n;

    // Mint BOOKING TOKEN with off chain payment ---------------------------------------------------------------------

    const tokenURI5 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp5 = (await ethers.provider.getBlock("latest")).timestamp + 120;
    const price5 = ethers.parseEther("0.95");
    const offChainPaymentToken = ethers.getAddress("0x0000000000000000000000000000000000000001");
    const offChainPaymentCurrency = 123;

    await supplierCMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorCMAccount.getAddress(), // Reserved for
        tokenURI5, // URI
        expirationTimestamp5, // Expiration of the reservation
        price5, // Price of token in wei
        offChainPaymentToken, // paymentToken
        offChainPaymentCurrency, // offchain payment currency
        true, // cancellable
    );

    // Token with ID 4 minted with off chain payment
    const tokenWithOffChainPayment = 4n;

    // Buy the token
    await distributorCMAccount
        .connect(distributorBookingOperator)
        .buyBookingToken(tokenWithOffChainPayment, price5, offChainPaymentToken);

    /// OTHER CM ACCOUNT ///

    // We also need another CM Account to test for fail cases
    // Create other CMAccount
    await nullUSD.approve(await cmAccountManager.getAddress(), prefundAmount);
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
    const otherCMAccountAddress = parsedEvent.args.account;

    // Get the CMAccount instance at the address
    const otherCMAccount = await ethers.getContractAt("CMAccount", otherCMAccountAddress);

    // Deposit funds to the CMAccount
    const depositAmount = ethers.parseEther("5");
    const depositTx = {
        to: otherCMAccount.getAddress(),
        value: depositAmount,
    };
    const txResponse = await signers.depositor.sendTransaction(depositTx);
    await txResponse.wait();

    // Distributor
    await otherCMAccount.connect(signers.cmAccountAdmin).grantRole(BOOKING_OPERATOR_ROLE, otherBookingOperator.address);

    return {
        supplierCMAccount,
        distributorCMAccount,
        bookingToken,
        nullUSD,
        tokenWithNativePayment,
        tokenWithNullUSDPayment,
        supplierBookingOperator,
        distributorBookingOperator,
        otherCMAccount,
        otherBookingOperator,
        tokenWithoutBuying,
        tokenWithPassedExpiration,
        offChainPaymentToken,
        offChainPaymentCurrency,
        tokenWithOffChainPayment,
    };
}

async function deployAndConfigureAllWithRegisteredServicesFixture() {
    // Set up signers
    await setupSigners();

    const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

    // Grant SERVICE_REGISTRY_ADMIN_ROLE
    const SERVICE_REGISTRY_ADMIN_ROLE = await cmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();
    await cmAccountManager
        .connect(signers.managerAdmin)
        .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.registryAdmin.address);

    // Services to register
    const serviceName1 = "cmp.service.accommodation.v1.AccommodationSearchService";
    const serviceHash1 = ethers.keccak256(ethers.toUtf8Bytes(serviceName1));

    const serviceName2 = "cmp.service.accommodation.v2.AccommodationSearchService";
    const serviceHash2 = ethers.keccak256(ethers.toUtf8Bytes(serviceName2));

    const serviceName3 = "cmp.service.accommodation.v3.AccommodationSearchService";
    const serviceHash3 = ethers.keccak256(ethers.toUtf8Bytes(serviceName3));

    const serviceName4 = "cmp.service.accommodation.v4.AccommodationSearchService";
    const serviceHash4 = ethers.keccak256(ethers.toUtf8Bytes(serviceName4));

    const serviceName5 = "cmp.service.accommodation.v5.AccommodationSearchService";
    const serviceHash5 = ethers.keccak256(ethers.toUtf8Bytes(serviceName5));

    const serviceName6 = "cmp.service.accommodation.v6.AccommodationSearchService";
    const serviceHash6 = ethers.keccak256(ethers.toUtf8Bytes(serviceName6));

    const services = {
        serviceName1,
        serviceHash1,
        serviceName2,
        serviceHash2,
        serviceName3,
        serviceHash3,
        serviceName4,
        serviceHash4,
        serviceName5,
        serviceHash5,
        serviceName6,
        serviceHash6,
    };

    // Register services
    await cmAccountManager.connect(signers.registryAdmin).registerService(serviceName1);
    await cmAccountManager.connect(signers.registryAdmin).registerService(serviceName2);
    await cmAccountManager.connect(signers.registryAdmin).registerService(serviceName3);
    await cmAccountManager.connect(signers.registryAdmin).registerService(serviceName4);
    await cmAccountManager.connect(signers.registryAdmin).registerService(serviceName5);
    await cmAccountManager.connect(signers.registryAdmin).registerService(serviceName6);

    // Get the SERVICE_ADMIN_ROLE
    const SERVICE_ADMIN_ROLE = await cmAccount.SERVICE_ADMIN_ROLE();

    // Grant SERVICE_ADMIN_ROLE to otherAccount1
    await cmAccount.connect(signers.cmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.cmServiceAdmin.address);

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
    deployBookingTokenWithNullUSDFixture,
    deployCancellationSupportFixture,
    deployNullUSDFixture,
};
