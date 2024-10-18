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
                bookingToken.connect(signers.btAdmin).safeMintWithReservation(
                    distributorCMAccount.getAddress(), // reservedFor
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                );

            // Mint again to make sure the token id is incremented
            await expect(
                await supplierCMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorCMAccount.getAddress(), // set reservedFor address to distributor CMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    1n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                );
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
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    supplierCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
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
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Try to transfer the token
            await expect(
                supplierCMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount1.address, 0n),
            )
                .to.be.revertedWithCustomError(bookingToken, "TokenIsReserved")
                .withArgs(0n, await distributorCMAccount.getAddress());
        });
    });
    describe("Cancellation", function () {
        it("supplier: should initiate cancellation of a booking token correctly", async function () {
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
                    ethers.ZeroAddress, // zero address
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorCMAccount.getAddress(),
                    supplierCMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                );

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await supplierCMAccount.getAddress());

            // Try to cancel the token

            const token_id = 0n;
            const initiator = await supplierCMAccount.getAddress();
            const refundAmount = ethers.parseEther("0.045");
            const refundCurrency = ethers.ZeroAddress;

            await expect(
                supplierCMAccount
                    .connect(signers.cmAccountAdmin)
                    .initiateCancellation(0n, refundAmount, refundCurrency),
            )
                .to.emit(bookingToken, "CancellationInitiated")
                .withArgs(token_id, initiator, refundAmount, refundCurrency);

            // Sanity check
            expect(await bookingToken.getCancellationProposalStatus(token_id)).to.be.deep.equal([
                refundAmount,
                refundCurrency,
                initiator,
                true,
            ]);
        });
        // FIXME: add tests for other cases
    });
});
