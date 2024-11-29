const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers } = require("hardhat");

// Fixtures
const {
    setupSigners,
    developerFeeBp,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployCMAccountWithDepositFixture,
    deployBookingTokenFixture,
    deployBookingTokenWithNullUSDFixture,
} = require("./utils/fixtures");

describe("BookingToken", function () {
    describe("Main", function () {
        it("should deploy correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            expect(await bookingToken.hasRole(await bookingToken.DEFAULT_ADMIN_ROLE(), signers.btAdmin.address)).to.be
                .true;
            expect(await bookingToken.hasRole(await bookingToken.UPGRADER_ROLE(), signers.btUpgrader.address)).to.be
                .true;
            expect(await bookingToken.isCMAccount(supplierCMAccount.getAddress())).to.be.true;
        });
    });

    describe("Mint", function () {
        it("Native: should revert if not called from a CMAccount", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            await expect(
                bookingToken
                    .connect(signers.btAdmin)
                    ["safeMintWithReservation(address,string,uint256,uint256,address,uint256,bool)"](
                        distributorCMAccount.getAddress(), // reservedFor
                        tokenURI, // tokenURI
                        expirationTimestamp, // expiration
                        price, // price
                        ethers.ZeroAddress, // paymentToken: zero address, means native coin
                        0, // off chain payment currency
                        true,
                    ),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount") // Caller is not a CMAccount
                .withArgs(signers.btAdmin.address);
        });

        it("Native: should revert if reservedFor is not a CMAccount", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    signers.otherAccount1.address, // set reservedFor to a non-CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount")
                .withArgs(signers.otherAccount1.address); // reservedFor address
        });

        it("Native: should mint a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Check cancellable flag
            expect(await bookingToken.isCancellable(0n)).to.equal(true);

            // Mint again to make sure the token id is incremented
            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    1n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(1n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(1n)).to.equal(1); // Reserved == 1
        });

        it("ERC20: should mint a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("120");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n, // token id
                    distributorCMAccount.getAddress(), // reservedFor
                    supplierCMAccount.getAddress(), // supplier
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                );

            // Sanity check
            expect(await bookingToken.getReservationPrice(0n)).to.be.deep.equal([price, await nullUSD.getAddress()]);
        });
    });

    describe("Buy", function () {
        it("Native: should buy a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [price, -price]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 3
        });

        it("Native: should buy a booking token with zero price correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [price, -price]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Off-chain: should buy a booking token with off-chain payment correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseUnits("559.99", 5);

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off-chain payment marker
            const offChainPaymentMarker = await bookingToken.OFFCHAIN_PAYMENT();

            // address(1)
            const OneAddress = ethers.getAddress("0x0000000000000000000000000000000000000001");

            expect(offChainPaymentMarker).to.equal(OneAddress);

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances, balances should not change as it is off-chain payment
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0n, 0n]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Off-chain: should buy a booking token with off-chain payment with zero price correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseUnits("0", 5);

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off-chain payment marker
            const offChainPaymentMarker = await bookingToken.OFFCHAIN_PAYMENT();

            // address(1)
            const OneAddress = ethers.getAddress("0x0000000000000000000000000000000000000001");

            expect(offChainPaymentMarker).to.equal(OneAddress);

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances, balances should not change as it is off-chain payment
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0n, 0n]);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Off-chain: should revert with off-chain payment and msg.value > 0", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.5");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off-chain payment marker
            const offChainPaymentMarker = await bookingToken.OFFCHAIN_PAYMENT();

            // address(1)
            const OneAddress = ethers.getAddress("0x0000000000000000000000000000000000000001");

            expect(offChainPaymentMarker).to.equal(OneAddress);

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/
            // Impersonate the CMAccount contract
            await network.provider.request({
                method: "hardhat_impersonateAccount",
                params: [await distributorCMAccount.getAddress()],
            });

            // Give it some CAM balance
            await network.provider.send("hardhat_setBalance", [
                await distributorCMAccount.getAddress(),
                ethers.toBeHex(price + ethers.parseEther("100")),
            ]);

            // Get the impersonated signer
            const impersonatedSigner = await ethers.getSigner(await distributorCMAccount.getAddress());

            // Try to buy the token with CAM - should revert
            await expect(bookingToken.connect(impersonatedSigner).buyReservedToken(0n, { value: price }))
                .to.be.revertedWithCustomError(bookingToken, "UnexpectedNativePayment")
                .withArgs(price);
        });

        it("ERC20: should buy a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("500");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            // CAM
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0, 0]);
            // Token: NullUSD
            await expect(buyTx).to.changeTokenBalances(
                nullUSD,
                [supplierCMAccount, distributorCMAccount],
                [price, -price],
            );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("ERC20: should revert if msg.value > 0", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("500");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Impersonate the CMAccount contract
            await network.provider.request({
                method: "hardhat_impersonateAccount",
                params: [await distributorCMAccount.getAddress()],
            });

            // Give it some CAM balance
            await network.provider.send("hardhat_setBalance", [
                await distributorCMAccount.getAddress(),
                ethers.toBeHex(price + ethers.parseEther("100")),
            ]);

            // Get the impersonated signer
            const impersonatedSigner = await ethers.getSigner(await distributorCMAccount.getAddress());

            // Try to buy the token with CAM - should revert
            await expect(bookingToken.connect(impersonatedSigner).buyReservedToken(0n, { value: price }))
                .to.be.revertedWithCustomError(bookingToken, "UnexpectedNativePayment")
                .withArgs(price);
        });

        it("ERC20: should buy a booking token with zero price correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check balances
            // CAM
            await expect(buyTx).to.changeEtherBalances([supplierCMAccount, distributorCMAccount], [0, 0]);
            // Token: NullUSD
            await expect(buyTx).to.changeTokenBalances(
                nullUSD,
                [supplierCMAccount, distributorCMAccount],
                [price, -price],
            );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());
        });

        it("Native: should revert when trying to buy a booking token reserved for another CMAccount", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    supplierCMAccount.getAddress(), // set reservedFor address to NOT distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    supplierCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx)
                .to.be.revertedWithCustomError(bookingToken, "ReservationMismatch")
                .withArgs(supplierCMAccount.getAddress(), distributorCMAccount.getAddress());
        });
    });

    describe("Transfer", function () {
        it("should revert transfer a booking token if the token is reserved", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Try to transfer the token
            await expect(
                supplierCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount1.address, 0n),
            )
                .to.be.revertedWithCustomError(bookingToken, "TokenIsReserved")
                .withArgs(0n, await distributorCMAccount.getAddress());
        });
        it("should transfer a booking token after token is bought", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 1

            // Set WITHDRAWER_ROLE
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            // Try to transfer the token, should not revert
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount1.address, 0n),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(distributorCMAccount.getAddress(), signers.otherAccount1.address, 0n);
        });
    });
    describe("Expiration", function () {
        it("should record a booking token as expired correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Try to expire the token before the expiration timestamp
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.be.revertedWithCustomError(bookingToken, "TokenIsReserved")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Advance time by 24 hours, token should can be expired after
            await network.provider.send("evm_increaseTime", [24 * 60 * 60]);
            await network.provider.send("evm_mine");

            // Expire the token
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.emit(bookingToken, "TokenExpired")
                .withArgs(0n);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(2); // Expired == 2
        });
        it("should revert recording a as expired if it's bought already", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 3

            // Try to expire the token, should revert with InvalidTokenStatus
            await expect(distributorCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 3); // Bought == 3
        });
    });
    describe("Cancellation", function () {
        it("should get cancellable flag correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Get cancellable flag
            expect(await bookingToken.isCancellable(0n)).to.equal(true);

            // Mint one with isCancellable set to false
            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    false,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    1n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    false,
                );

            // Check token booking status
            expect(await bookingToken.getBookingStatus(1n)).to.equal(1); // Reserved == 1

            // Get cancellable flag
            expect(await bookingToken.isCancellable(1n)).to.equal(false);
        });

        it("should initiate cancellation of a booking token correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Wait for the transaction to be mined
            await expect(buyTx).to.not.be.reverted;

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(
                supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 42n, 1n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Test cancelling a cancellation proposal
            await expect(
                supplierCMAccount.connect(signers.btAdmin).withdrawCancellationProposal(
                    token_id,
                    42, // Reason
                    1, // Reason Version
                ),
            )
                .to.emit(bookingToken, "CancellationWithdrawn")
                .withArgs(token_id, await supplierCMAccount.getAddress(), 42n, 1n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                0n,
                ethers.ZeroAddress,
                0n, // Unspecified
                0n, // Cancellation Reason
                0n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Initiate the cancellation with the distributor
            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(
                    0n,
                    refundAmount,
                    42n, // Cancellation Reason
                    1n, // Cancellation Reason Version
                ),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, await distributorCMAccount.getAddress(), refundAmount, 42n, 1n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                await distributorCMAccount.getAddress(),
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Test cancelling a cancellation proposal
            await expect(
                distributorCMAccount.connect(signers.btAdmin).withdrawCancellationProposal(
                    0n,
                    42n, // Reason
                    1n, // Reason Version
                ),
            )
                .to.emit(bookingToken, "CancellationWithdrawn")
                .withArgs(token_id, await distributorCMAccount.getAddress(), 42n, 1n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                0n,
                ethers.ZeroAddress,
                0n, // Status: Unspecified
                0n, // Cancellation Reason
                0n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);
        });

        it("should revert initiating a proposal if token state is reserved or expired", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(
                supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 1n); // tokenID == 0, Reserved == 1

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                0n,
                ethers.ZeroAddress,
                0n, // NoProposal == 0
                0n, // Cancellation Reason
                0n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Expire the token

            // Advance time by 24 hours, token should can be expired after
            await network.provider.send("evm_increaseTime", [24 * 60 * 60]);
            await network.provider.send("evm_mine");

            // Expire the token
            await expect(supplierCMAccount.connect(signers.btAdmin).recordExpiration(0n))
                .to.emit(bookingToken, "TokenExpired")
                .withArgs(0n);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(2); // Expired == 2

            // Try to cancel the token
            await expect(
                supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(0n, 2n); // tokenID == 0, Expired == 2
        });
        it("Native: should accept a cancellation correctly and pay the refund", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx)
                .to.be.emit(bookingToken, "TokenBought")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 42n, 1n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Try to accept the cancellation with an incorrect refund amount
            await expect(
                distributorCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id, refundAmount + 1n),
            )
                .to.revertedWithCustomError(bookingToken, "IncorrectAmount")
                .withArgs(refundAmount + 1n, refundAmount);

            // Accept the cancellation, this should send the refund from supplier to distributor
            const acceptTx = supplierCMAccount
                .connect(signers.btAdmin)
                .acceptCancellationProposal(token_id, refundAmount);

            // Wait for the txn
            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);

            // Check balances
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [refundAmount, -refundAmount, 0n],
            );
        });
        it("Native: should revert accepting if the caller is not authorized", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx).to.not.be.reverted;

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(3); // Bought == 1

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            // Initiate the cancellation by the supplier
            await expect(
                supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 0n, 0n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 0n, 0n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer, // supplier
                1n, // Pending == 1
                0n, // Cancellation Reason
                0n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Try to accept the cancellation proposed by the supplier with the
            // supplier cm account, this should revert due to being not authorized
            await expect(supplierCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id, refundAmount))
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToAcceptCancellation")
                .withArgs(0n, await supplierCMAccount.getAddress());

            // Try to accept the cancellation proposed by the supplier with another account
            await expect(bookingToken.connect(signers.otherAccount3).acceptCancellationProposal(token_id, refundAmount))
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount")
                .withArgs(signers.otherAccount3.address);

            // Accept the supplier's proposal with the distributor cm account
            await expect(
                distributorCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id, refundAmount),
            )
                .to.emit(bookingToken, "CancellationAcceptedByTheOwner")
                .withArgs(token_id, await distributorCMAccount.getAddress(), refundAmount);

            // Now the supplier's proposal is accepted by the owner/distributor. Try
            // to accept it again with the distributor. This should revert with the
            // same error
            await expect(
                distributorCMAccount.connect(signers.btAdmin).acceptCancellationProposal(token_id, refundAmount),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToAcceptCancellation")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Sanity check: Try to accept the cancellation proposed by the owner with another account
            await expect(bookingToken.connect(signers.otherAccount2).acceptCancellationProposal(token_id, refundAmount))
                .to.be.revertedWithCustomError(bookingToken, "NotCMAccount")
                .withArgs(signers.otherAccount2.address);

            // Finally, accept the "owner accepted proposal" with the supplier cm account
            const acceptTx = await supplierCMAccount
                .connect(signers.btAdmin)
                .acceptCancellationProposal(token_id, refundAmount);

            // Check balances
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [refundAmount, -refundAmount, 0n],
            );

            // Check events
            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);
        });

        it("ERC20: should accept a cancellation correctly and pay the refund", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("500");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    nullUSD.getAddress(), // nullUSD address
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("450");

            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 123n, 2n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 123n, 2n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                123n, // Cancellation Reason
                2n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Accept the cancellation, this should send the refund from supplier to distributor
            const acceptTx = supplierCMAccount
                .connect(signers.btAdmin)
                .acceptCancellationProposal(token_id, refundAmount);

            // Sanity check
            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);

            // Check balances
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [0n, 0n, 0n],
            );

            await expect(acceptTx).to.changeTokenBalances(
                nullUSD,
                [distributorCMAccount, supplierCMAccount],
                [refundAmount, -refundAmount],
            );
        });

        it("Off-chain: should accept a cancellation correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("500");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Off-chain payment marker
            const offChainPaymentMarker = await bookingToken.OFFCHAIN_PAYMENT();

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    offChainPaymentMarker, // off-chain payment marker, address(1)
                    6, // off chain payment currency, 6 == Euro
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Try to buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            // Check emitted events
            await expect(buyTx).to.be.emit(bookingToken, "TokenBought").withArgs(0n, distributorCMAccount.getAddress());

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await distributorCMAccount.getAddress());

            // Try to cancel the token
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("450");

            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 123n, 2n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 123n, 2n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                123n, // Cancellation Reason
                2n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Accept the cancellation, this should send the refund from supplier to distributor
            const acceptTx = supplierCMAccount
                .connect(signers.btAdmin)
                .acceptCancellationProposal(token_id, refundAmount);

            // Sanity check
            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);

            // Check balances, balances should not change as it is off-chain payment
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [0n, 0n, 0n],
            );

            // Sanity check
            await expect(acceptTx).to.changeTokenBalances(nullUSD, [distributorCMAccount, supplierCMAccount], [0n, 0n]);
        });

        it("should revert transferring if there is an cancellation is active or cancelled", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx)
                .to.be.emit(bookingToken, "TokenBought")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Initiate the proposal
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 42n, 1n);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // THERE IS AN ACTIVE CANCELLATION PROPOSAL

            // Set WITHDRAWER_ROLE
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            // Try to transfer the token, should revert.
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
            )
                .to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled")
                .withArgs(0n);

            // Sanity check
            await expect(
                bookingToken
                    .connect(signers.withdrawer)
                    .safeTransferFrom(await distributorCMAccount.getAddress(), signers.otherAccount2.address, 0n),
            ).to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled");

            await expect(
                bookingToken
                    .connect(signers.withdrawer)
                    .transferFrom(await distributorCMAccount.getAddress(), signers.otherAccount2.address, 0n),
            ).to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled");

            // Accept the cancellation
            await expect(supplierCMAccount.connect(signers.btAdmin).acceptCancellationProposal(0n, refundAmount)).to.be
                .not.reverted;

            // Sanity check
            // Check token ownership, should revert
            await expect(bookingToken.ownerOf(0n))
                .to.be.revertedWithCustomError(bookingToken, "ERC721NonexistentToken")
                .withArgs(0n);

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(4); // Cancelled == 4

            // Check the proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                4n, // Accepted == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Try to transfer the token, should revert with InvalidTokenStatus
            // because token is burned when cancellation is accepted
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
            ).to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus");
        });
        it("should reject correctly and allow transfer if cancellation is rejected", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx)
                .to.be.emit(bookingToken, "TokenBought")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Initiate the cancellation
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 42n, 1n);

            // Check proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // TRY TRANSFER, SHOULD REVERT

            // Set WITHDRAWER_ROLE
            const WITHDRAWER_ROLE = await distributorCMAccount.WITHDRAWER_ROLE();
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(WITHDRAWER_ROLE, signers.withdrawer.address),
            ).to.not.reverted;

            // Try to transfer the token, should revert.
            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
            )
                .to.be.revertedWithCustomError(bookingToken, "TokenHasActiveCancellationProposalOrCancelled")
                .withArgs(0n);

            // REJECT THE CANCELLATION

            // Try to reject with distributor, should revert
            await expect(
                distributorCMAccount.connect(signers.btAdmin).rejectCancellationProposal(
                    token_id,
                    37n, // Rejection Reason
                    2n, // Rejection Reason Version
                ),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToRejectCancellation")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Actual reject with supplier
            await expect(
                supplierCMAccount.connect(signers.btAdmin).rejectCancellationProposal(
                    token_id,
                    37n, // Rejection Reason
                    2n, // Rejection Reason Version
                ),
            )
                .to.emit(bookingToken, "CancellationRejected")
                .withArgs(token_id, await supplierCMAccount.getAddress(), 37n, 2n);

            // Check proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                2n, // Rejected == 2
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                37n, // Rejection Reason
                2n, // Rejection Reason Version: Unspecified
            ]);

            // TRY TRANSFER AGAIN, SHOULD NOT REVERT

            await expect(
                distributorCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(await distributorCMAccount.getAddress(), signers.otherAccount2.address, 0n);
        });
        it("counter proposal: should do counter proposals correctly", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx)
                .to.be.emit(bookingToken, "TokenBought")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Initiate the cancellation
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 42n, 1n);

            // Check proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // COUNTER PROPOSALS

            const newRefundAmount = ethers.parseEther("0.030");

            // Try counter proposal with distributor cm account, should revert
            await expect(
                distributorCMAccount.connect(signers.btAdmin).counterCancellationProposal(token_id, refundAmount),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToCounterCancellation")
                .withArgs(token_id, await distributorCMAccount.getAddress());

            // Counter Cancellation Proposal with CM account, should emit CancellationCountered event
            await expect(
                supplierCMAccount.connect(signers.btAdmin).counterCancellationProposal(token_id, newRefundAmount),
            )
                .to.emit(bookingToken, "CancellationCountered")
                .withArgs(token_id, await supplierCMAccount.getAddress(), newRefundAmount);

            // Check proposal, should have new refund amount and status of Countered
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                newRefundAmount,
                proposer, // new proposer is the supplier now
                3n, // Countered == 3
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // DETOUR: Test "withdraw counter proposal" by distributor cm account

            await expect(
                distributorCMAccount.connect(signers.btAdmin).withdrawCancellationProposal(
                    token_id,
                    42n, // Reason
                    1n, // Reason Version
                ),
            )
                .to.emit(bookingToken, "CancellationWithdrawn")
                .withArgs(token_id, await distributorCMAccount.getAddress(), 42n, 1n);

            // Check proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                0n,
                ethers.ZeroAddress,
                0n,
                0n, // Cancellation Reason
                0n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // BACK TO COUNTER PROPOSAL: Recreate the cancellation
            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 42n, 1n);

            await expect(
                supplierCMAccount.connect(signers.btAdmin).counterCancellationProposal(token_id, newRefundAmount),
            )
                .to.emit(bookingToken, "CancellationCountered")
                .withArgs(token_id, await supplierCMAccount.getAddress(), newRefundAmount);
            // END BACK TO COUNTER PROPOSAL

            // Try to accept the countered cancellation proposal with supplier cm account, should revert
            await expect(
                supplierCMAccount
                    .connect(signers.btAdmin)
                    .acceptCounteredCancellationProposal(token_id, newRefundAmount),
            )
                .to.be.revertedWithCustomError(bookingToken, "NotAuthorizedToAcceptCounterProposal")
                .withArgs(token_id, await supplierCMAccount.getAddress());

            // Accept the countered cancellation proposal with distributor cm account
            await expect(
                distributorCMAccount
                    .connect(signers.btAdmin)
                    .acceptCounteredCancellationProposal(token_id, newRefundAmount),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, newRefundAmount, 42n, 1n);

            // Check proposal, should be with the new refund amount and status to back to Pending
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                newRefundAmount,
                proposer,
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Finally, accept the proposal with the supplier cm account and transferring the refund
            const acceptTx = supplierCMAccount
                .connect(signers.btAdmin)
                .acceptCancellationProposal(token_id, newRefundAmount);

            await expect(acceptTx)
                .to.emit(bookingToken, "CancellationAccepted")
                .withArgs(token_id, await supplierCMAccount.getAddress(), newRefundAmount);

            // Check balances
            await expect(acceptTx).to.changeEtherBalances(
                [distributorCMAccount, supplierCMAccount, bookingToken],
                [newRefundAmount, -newRefundAmount, 0n],
            );

            // Distributor: Try to initiate cancellation again, should revert
            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(
                    token_id,
                    4n, // Cancelled == 4
                );

            // Supplier: Try to initiate cancellation again, should revert
            await expect(
                supplierCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.be.revertedWithCustomError(bookingToken, "InvalidTokenStatus")
                .withArgs(
                    token_id,
                    4n, // Cancelled == 4
                );

            // Distributor: Try to counter cancellation again, should revert
            await expect(
                distributorCMAccount.connect(signers.btAdmin).counterCancellationProposal(token_id, refundAmount),
            )
                .to.be.revertedWithCustomError(bookingToken, "NoPendingCancellationProposal")
                .withArgs(token_id);

            // Supplier: Try to counter cancellation again, should revert
            await expect(supplierCMAccount.connect(signers.btAdmin).counterCancellationProposal(token_id, refundAmount))
                .to.be.revertedWithCustomError(bookingToken, "NoPendingCancellationProposal")
                .withArgs(token_id);
        });
        it("rejected: should do counter proposals for rejected cancellations", async function () {
            const { cmAccountManager, supplierCMAccount, distributorCMAccount, bookingToken } =
                await loadFixture(deployBookingTokenFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            /***************************************************
             *                   SUPPLIER                      *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierCMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // paymentToken: zero address, means native coin
                    0, // off chain payment currency
                    true,
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Check token booking status
            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved == 1

            /***************************************************
             *                  DISTRIBUTOR                    *
             ***************************************************/

            // Grant BOOKING_OPERATOR_ROLE
            await expect(
                distributorCMAccount
                    .connect(signers.cmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            // Buy the token
            const buyTx = distributorCMAccount.connect(signers.btAdmin).buyBookingToken(0n);

            await expect(buyTx)
                .to.be.emit(bookingToken, "TokenBought")
                .withArgs(0n, await distributorCMAccount.getAddress());

            // Initiate the cancellation
            const token_id = 0n;
            const proposer = await distributorCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");

            await expect(
                distributorCMAccount.connect(signers.btAdmin).initiateCancellationProposal(0n, refundAmount, 42n, 1n),
            )
                .to.emit(bookingToken, "CancellationPending")
                .withArgs(token_id, proposer, refundAmount, 42n, 1n);

            // Check proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                1n, // Pending == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason: Unspecified
                0n, // Rejection Reason Version: Unspecified
            ]);

            // Reject the cancellation
            await expect(supplierCMAccount.connect(signers.btAdmin).rejectCancellationProposal(token_id, 37n, 1n))
                .to.emit(bookingToken, "CancellationRejected")
                .withArgs(token_id, await supplierCMAccount.getAddress(), 37n, 1n);

            // Check proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                2n, // Rejected == 2
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                37n, // Rejection Reason
                1n, // Rejection Reason Version
            ]);

            // Revive the rejected cancellation proposal with a counter proposal

            await expect(supplierCMAccount.connect(signers.btAdmin).counterCancellationProposal(token_id, refundAmount))
                .to.emit(bookingToken, "CancellationCountered")
                .withArgs(token_id, await supplierCMAccount.getAddress(), refundAmount);

            // Check proposal
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                proposer,
                3n, // Countered == 1
                42n, // Cancellation Reason
                1n, // Cancellation Reason Version
                0n, // Rejection Reason
                0n, // Rejection Reason Version
            ]);
        });
    });
});
