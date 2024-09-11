# Solidity API

## CMAccount

A CM Account manages funds, minting/buying of booking tokens, provided
or wanted services, and multiple bots for distributors and suppliers on
Camino Messenger ecosystem.

Registering bots is done by role based access control. Bot's with
`CHEQUE_OPERATOR_ROLE` can issue cheques to paid by the {CMAccount} contract.
Bot can also have `GAS_WITHDRAWER_ROLE` and `BOOKING_OPERATOR_ROLE`.

`GAS_WITHDRAWER_ROLE` enables a bot to withdraw native coins (CAM) from the
contract to be used as gas money. This restricted with a `limit`
(wei/aCAM) and `period` (seconds) by the `BOT_ADMIN_ROLE`. Default starting
values are 10 CAM per 24 hours.

`BOOKING_OPERATOR_ROLE` enables a bot to mint and buy Booking Tokens by
calling the corresponding functions on the {BookingToken} contract. The buy
operation pays the price of the Booking Token with the funds on the
{CMAccount} contract.

_This contract uses UUPS style upgradeability. The authorization function
`_authorizeUpgrade(address)` can be called by the `UPGRADER_ROLE` and is
restricted to only upgrade to the implementation address registered on the
{CMAccountManager} contract._

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

Upgrader role can upgrade the contract to a new implementation.

### BOT_ADMIN_ROLE

```solidity
bytes32 BOT_ADMIN_ROLE
```

Bot admin role can add & remove bots and set gas money withdrawal
parameters.

### CHEQUE_OPERATOR_ROLE

```solidity
bytes32 CHEQUE_OPERATOR_ROLE
```

Cheque operator role can issue cheques to be paid by this CMAccount
contract.

### GAS_WITHDRAWER_ROLE

```solidity
bytes32 GAS_WITHDRAWER_ROLE
```

Gas withdrawer role can withdraw gas money from the contract. This is
intended to be used by the bots and is granted when `addMessengerBot` is
called.

### WITHDRAWER_ROLE

```solidity
bytes32 WITHDRAWER_ROLE
```

Withdrawer role can withdraw funds from the contract.

### BOOKING_OPERATOR_ROLE

```solidity
bytes32 BOOKING_OPERATOR_ROLE
```

Booking operator role can mint and buy booking tokens using the
functions on this contract. This is generally used by the bots. The
price for the booking token is paid by this contract.

### SERVICE_ADMIN_ROLE

```solidity
bytes32 SERVICE_ADMIN_ROLE
```

Service admin role can add & remove supported & wanted services.

### CMAccountStorage

```solidity
struct CMAccountStorage {
    address _manager;
    address _bookingToken;
    uint256 _prefundAmount;
}
```

### CMAccountUpgraded

```solidity
event CMAccountUpgraded(address oldImplementation, address newImplementation)
```

CMAccount upgrade event. Emitted when the CMAccount implementation is upgraded.

### Deposit

```solidity
event Deposit(address sender, uint256 amount)
```

Deposit event, emitted when there is a new deposit

### Withdraw

```solidity
event Withdraw(address receiver, uint256 amount)
```

Withdraw event, emitted when there is a new withdrawal

### MessengerBotAdded

```solidity
event MessengerBotAdded(address bot)
```

Messenger bot added

### MessengerBotRemoved

```solidity
event MessengerBotRemoved(address bot)
```

Messenger bot removed

### CMAccountImplementationMismatch

```solidity
error CMAccountImplementationMismatch(address latestImplementation, address newImplementation)
```

CMAccount implementation address does not match the one in the manager

### CMAccountNoUpgradeNeeded

```solidity
error CMAccountNoUpgradeNeeded(address oldImplementation, address newImplementation)
```

New implementation is the same as the current implementation, no update needed

### DepositorNotAllowed

```solidity
error DepositorNotAllowed(address sender)
```

Error to revert with if depositer is not allowed

### ZeroValueDeposit

```solidity
error ZeroValueDeposit(address sender)
```

Error to revert zero value deposits

### PrefundNotSpentYet

```solidity
error PrefundNotSpentYet(uint256 withdrawableAmount, uint256 prefundLeft, uint256 amount)
```

Error to revert with if the prefund is not spent yet

### TransferToZeroAddress

```solidity
error TransferToZeroAddress()
```

Error to revert if transfer to zero address

### constructor

```solidity
constructor() public
```

### initialize

```solidity
function initialize(address manager, address bookingToken, uint256 prefundAmount, address defaultAdmin, address upgrader) public
```

### receive

```solidity
receive() external payable
```

### getManagerAddress

```solidity
function getManagerAddress() public view returns (address)
```

Returns the CMAccountManager address.

#### Return Values

| Name | Type    | Description              |
| ---- | ------- | ------------------------ |
| [0]  | address | CMAccountManager address |

### getBookingTokenAddress

```solidity
function getBookingTokenAddress() public view returns (address)
```

Returns the booking token address.

#### Return Values

| Name | Type    | Description          |
| ---- | ------- | -------------------- |
| [0]  | address | BookingToken address |

### getPrefundAmount

```solidity
function getPrefundAmount() public view returns (uint256)
```

Returns the prefund amount.

#### Return Values

| Name | Type    | Description    |
| ---- | ------- | -------------- |
| [0]  | uint256 | prefund amount |

### \_authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

Authorizes the upgrade of the CMAccount.

Reverts if the new implementation is the same as the old one.

Reverts if the new implementation does not match the implementation address
in the manager. Only implementations registered at the manager are allowed.

_Emits a {CMAccountUpgraded} event._

#### Parameters

| Name              | Type    | Description                    |
| ----------------- | ------- | ------------------------------ |
| newImplementation | address | The new implementation address |

### isBotAllowed

```solidity
function isBotAllowed(address bot) public view returns (bool)
```

Returns true if an address is authorized to sign cheques

#### Parameters

| Name | Type    | Description       |
| ---- | ------- | ----------------- |
| bot  | address | The bot's address |

### withdraw

```solidity
function withdraw(address payable recipient, uint256 amount) external
```

Withdraw CAM from the CMAccount

This function reverts if the amount is bigger then the prefund left to spend. This is to prevent
spam by forcing user to spend the full prefund for cheques, so they can not just create an account
and withdraw the prefund.

### mintBookingToken

```solidity
function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken) external
```

Mints booking token.

#### Parameters

| Name                | Type            | Description                                  |
| ------------------- | --------------- | -------------------------------------------- |
| reservedFor         | address         | The account to reserve the token for         |
| uri                 | string          | The URI of the token                         |
| expirationTimestamp | uint256         | The expiration timestamp                     |
| price               | uint256         | The price of the token                       |
| paymentToken        | contract IERC20 | The payment token, if address(0) then native |

### buyBookingToken

```solidity
function buyBookingToken(uint256 tokenId) external
```

Buys booking token.

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

### onERC721Received

```solidity
function onERC721Received(address, address, uint256, bytes) public virtual returns (bytes4)
```

Always returns `IERC721Receiver.onERC721Received.selector`.

_See {IERC721Receiver-onERC721Received}._

### transferERC20

```solidity
function transferERC20(contract IERC20 token, address to, uint256 amount) external
```

Transfers ERC20 tokens.

This function reverts if `to` is the zero address.

#### Parameters

| Name   | Type            | Description                           |
| ------ | --------------- | ------------------------------------- |
| token  | contract IERC20 | The ERC20 token                       |
| to     | address         | The address to transfer the tokens to |
| amount | uint256         | The amount of tokens to transfer      |

### transferERC721

```solidity
function transferERC721(contract IERC721 token, address to, uint256 tokenId) external
```

Transfers ERC721 tokens.

This function reverts if `to` is the zero address.

#### Parameters

| Name    | Type             | Description                           |
| ------- | ---------------- | ------------------------------------- |
| token   | contract IERC721 | The ERC721 token                      |
| to      | address          | The address to transfer the tokens to |
| tokenId | uint256          | The token id of the token             |

### addService

```solidity
function addService(string serviceName, uint256 fee, bool restrictedRate, string[] capabilities) public
```

Adds a service to the account as a supported service.

`serviceName` is defined as pkg + service name in protobuf. For example:

```text
 ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
"cmp.services.accommodation.v1alpha.AccommodationSearchService")
```

_These services are coming from the Camino Messenger Protocol's protobuf
definitions._

#### Parameters

| Name           | Type     | Description                                               |
| -------------- | -------- | --------------------------------------------------------- |
| serviceName    | string   | Service name to add to the account as a supported service |
| fee            | uint256  | Fee of the service in aCAM (wei in ETH terminology)       |
| restrictedRate | bool     |                                                           |
| capabilities   | string[] | Capabilities of the service (if any, optional)            |

### removeService

```solidity
function removeService(string serviceName) public
```

Remove a service from the account by its name

### setServiceFee

```solidity
function setServiceFee(string serviceName, uint256 fee) public
```

Set the fee of a service by name

### setServiceRestrictedRate

```solidity
function setServiceRestrictedRate(string serviceName, bool restrictedRate) public
```

Set the restricted rate of a service by name

### setServiceCapabilities

```solidity
function setServiceCapabilities(string serviceName, string[] capabilities) public
```

Set all capabilities for a service by name

### addServiceCapability

```solidity
function addServiceCapability(string serviceName, string capability) public
```

Add a single capability to the service by name

### removeServiceCapability

```solidity
function removeServiceCapability(string serviceName, string capability) public
```

Remove a single capability from the service by name

### getSupportedServices

```solidity
function getSupportedServices() public view returns (string[] serviceNames, struct PartnerConfiguration.Service[] services)
```

Get all supported services. Return a list of service names and a list of service objects.

### getServiceFee

```solidity
function getServiceFee(string serviceName) public view returns (uint256 fee)
```

Get service fee by name. Overloading the getServiceFee function.

### getServiceRestrictedRate

```solidity
function getServiceRestrictedRate(string serviceName) public view returns (bool restrictedRate)
```

Get service restricted rate by name. Overloading the getServiceRestrictedRate function.

### getServiceCapabilities

```solidity
function getServiceCapabilities(string serviceName) public view returns (string[] capabilities)
```

Get service capabilities by name. Overloading the getServiceCapabilities function.

### addWantedServices

```solidity
function addWantedServices(string[] serviceNames) public
```

Adds wanted services.

#### Parameters

| Name         | Type     | Description           |
| ------------ | -------- | --------------------- |
| serviceNames | string[] | List of service names |

### removeWantedServices

```solidity
function removeWantedServices(string[] serviceNames) public
```

Removes wanted services.

#### Parameters

| Name         | Type     | Description           |
| ------------ | -------- | --------------------- |
| serviceNames | string[] | List of service names |

### getWantedServices

```solidity
function getWantedServices() public view returns (string[] serviceNames)
```

Get all wanted services.

#### Return Values

| Name         | Type     | Description           |
| ------------ | -------- | --------------------- |
| serviceNames | string[] | List of service names |

### setOffChainPaymentSupported

```solidity
function setOffChainPaymentSupported(bool _isSupported) public
```

Sets if off-chain payment is supported.

#### Parameters

| Name          | Type | Description                            |
| ------------- | ---- | -------------------------------------- |
| \_isSupported | bool | true if off-chain payment is supported |

### addSupportedToken

```solidity
function addSupportedToken(address _supportedToken) public
```

Adds a supported payment token.

#### Parameters

| Name             | Type    | Description          |
| ---------------- | ------- | -------------------- |
| \_supportedToken | address | address of the token |

### removeSupportedToken

```solidity
function removeSupportedToken(address _supportedToken) public
```

Removes a supported payment token.

#### Parameters

| Name             | Type    | Description          |
| ---------------- | ------- | -------------------- |
| \_supportedToken | address | address of the token |

### addPublicKey

```solidity
function addPublicKey(address pubKeyAddress, bytes data) public
```

Add public key with address

These public keys are intended to be used with for off-chain encryption of private booking data.

#### Parameters

| Name          | Type    | Description               |
| ------------- | ------- | ------------------------- |
| pubKeyAddress | address | address of the public key |
| data          | bytes   | public key data           |

### removePublicKey

```solidity
function removePublicKey(address pubKeyAddress) public
```

Remove public key by address

### addMessengerBot

```solidity
function addMessengerBot(address bot, uint256 gasMoney) public
```

Adds messenger bot with initial gas money.

### removeMessengerBot

```solidity
function removeMessengerBot(address bot) public
```

Removes messenger bot by revoking the roles.

### withdrawGasMoney

```solidity
function withdrawGasMoney(uint256 amount) public
```

Withdraw gas money. Requires the `GAS_WITHDRAWER_ROLE`.

#### Parameters

| Name   | Type    | Description                          |
| ------ | ------- | ------------------------------------ |
| amount | uint256 | The amount to withdraw in aCAM (wei) |

### setGasMoneyWithdrawal

```solidity
function setGasMoneyWithdrawal(uint256 limit, uint256 period) public
```

Set gas money withdrawal parameters. Requires the `BOT_ADMIN_ROLE`.

#### Parameters

| Name   | Type    | Description                                       |
| ------ | ------- | ------------------------------------------------- |
| limit  | uint256 | Amount of gas money to withdraw in wei per period |
| period | uint256 | Duration of the withdrawal period in seconds      |

## ChequeManager

ChequeManager manages, verifies, and cashes in the messenger cheques.

EIP712 Domain name & version:
DOMAIN_NAME = "CaminoMessenger"
DOMAIN_VERSION= "1"

### MESSENGER_CHEQUE_TYPEHASH

```solidity
bytes32 MESSENGER_CHEQUE_TYPEHASH
```

Pre-computed hash of the MessengerCheque type

```
keccak256(
    "MessengerCheque(address fromCMAccount,address toCMAccount,address toBot,uint256 counter,uint256 amount,uint256 createdAt,uint256 expiresAt)"
);
```

### DOMAIN_TYPEHASH

```solidity
bytes32 DOMAIN_TYPEHASH
```

Pre-computed hash of the EIP712Domain type

```
keccak256("EIP712Domain(string name,string version,uint256 chainId)");
```

### MessengerCheque

Struct representing a Messenger Cheque.

```solidity
struct MessengerCheque {
    address fromCMAccount;
    address toCMAccount;
    address toBot;
    uint256 counter;
    uint256 amount;
    uint256 createdAt;
    uint256 expiresAt;
}
```

### LastCashIn

Struct for tracking the counter, amount and timestamps used for the last
cash-in operation.

```solidity
struct LastCashIn {
    uint256 counter;
    uint256 amount;
    uint256 createdAt;
    uint256 expiresAt;
}
```

### ChequeManagerStorage

```solidity
struct ChequeManagerStorage {
  mapping(address => mapping(address => struct ChequeManager.LastCashIn)) _lastCashIns;
  uint256 _totalChequePayments;
  bytes32 _domainSeparator;
}
```

### ChequeCashedIn

```solidity
event ChequeCashedIn(address fromCMAccount, address toCMAccount, address fromBot, address toBot, uint256 counter, uint256 amount, uint256 paidChequeAmount, uint256 paidDeveloperFee)
```

Cash-in event. Emitted when a cheque is cashed in.

### InvalidFromCMAccount

```solidity
error InvalidFromCMAccount(address fromCMAccount)
```

Invalid CM Account. Cheque's `fromCMAccount` has to be for `address(this)`.

### InvalidToCMAccount

```solidity
error InvalidToCMAccount(address toCMAccount)
```

`toCMAccount` address is not a registered CMAccount on the manager.

### NotAllowedToSignCheques

```solidity
error NotAllowedToSignCheques(address signer)
```

The signer is not allowed to sign cheques

### InvalidCounter

```solidity
error InvalidCounter(uint256 chequeCounter, uint256 lastCounter)
```

Invalid counter for the cheque. The counter on the cheque is not greater then the last
recorded counter.

### InvalidAmount

```solidity
error InvalidAmount(uint256 chequeAmount, uint256 lastAmount)
```

Last recorded amount is lower than the cheque's amount.

### ChequeExpired

```solidity
error ChequeExpired(uint256 expiresAt)
```

The cheque is expired at the given timestamp.

### \_\_ChequeManager_init

```solidity
function __ChequeManager_init() internal
```

Initializes the contract, setting the domain separator with EIP712 domain type hash and
the domain.

EIP712Domain {
string name;
string version;
uint256 chainid;
}

### getDomainSeparator

```solidity
function getDomainSeparator() public view returns (bytes32)
```

Returns the domain separator.

### hashMessengerCheque

```solidity
function hashMessengerCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt) public pure returns (bytes32)
```

Returns the hash of the `MessengerCheque` encoded with
`MESSENGER_CHEQUE_TYPEHASH`.

### hashTypedDataV4

```solidity
function hashTypedDataV4(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt) public view returns (bytes32)
```

Returns the hash of the typed data (cheque) with prefix and domain
separator.

### recoverSigner

```solidity
function recoverSigner(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, bytes signature) internal view returns (address signer)
```

Returns the signer for the given cheque and signature. Uses {ECDSA} library to
recover the signer.

### verifyCheque

```solidity
function verifyCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, bytes signature) public view returns (address signer, uint256 paymentAmount)
```

Returns signer and payment amount if the signature is valid for the
given cheque, the signer is an allowed bot, cheque counter and amounts are
valid according to last cash ins.

Please be aware that `cheque.amount < paymentAmount` for a valid cheque as
long as the last amount is lower than the cheque amount. Only the difference
between the cheque amount and the last recorded amount is paid.

### cashInCheque

```solidity
function cashInCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, bytes signature) public
```

Cash in a cheque by verifying it and paying the difference between the
cheque amount and the last recorded amount for the signer and `toBot` pair.

A percentage of the amount is also paid to the developer wallet.

#### Parameters

| Name          | Type    | Description                                                                         |
| ------------- | ------- | ----------------------------------------------------------------------------------- |
| fromCMAccount | address | The CM Account that will pay the amount. This contract.                             |
| toCMAccount   | address | The CM Account that will receive the amount.                                        |
| toBot         | address | The address of the bot that received the cheque.                                    |
| counter       | uint256 | The counter of the cheque. Should be increased with every cheque.                   |
| amount        | uint256 | The amount on the cheque. Should be greater then or equal the last recorded amount. |
| createdAt     | uint256 | The creation timestamp of the cheque.                                               |
| expiresAt     | uint256 | The expiration timestamp of the cheque.                                             |
| signature     | bytes   | The signature of the cheque.                                                        |

### getLastCashIn

```solidity
function getLastCashIn(address fromBot, address toBot) public view returns (uint256 lastCounter, uint256 lastAmount, uint256 lastCreatedAt, uint256 lastExpiresAt)
```

Returns last cash-in details for given `fromBot` & `toBot` pair.

#### Parameters

| Name    | Type    | Description                                                                                                      |
| ------- | ------- | ---------------------------------------------------------------------------------------------------------------- |
| fromBot | address | The address of the bot that sent the cheque.                                                                     |
| toBot   | address | The address of the bot that received the cheque. Returns (lastCounter, lastAmount, lastCreatedAt, lastExpiresAt) |

#### Return Values

| Name          | Type    | Description                                  |
| ------------- | ------- | -------------------------------------------- |
| lastCounter   | uint256 | The last counter of the cheque.              |
| lastAmount    | uint256 | The last amount of the cheque.               |
| lastCreatedAt | uint256 | The last creation timestamp of the cheque.   |
| lastExpiresAt | uint256 | The last expiration timestamp of the cheque. |

### setLastCashIn

```solidity
function setLastCashIn(address fromBot, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt) internal
```

Sets last cash-in for given `fromBot`, `toBot` pair.

#### Parameters

| Name      | Type    | Description                                      |
| --------- | ------- | ------------------------------------------------ |
| fromBot   | address | The address of the bot that sent the cheque.     |
| toBot     | address | The address of the bot that received the cheque. |
| counter   | uint256 | The counter of the cheque.                       |
| amount    | uint256 | The amount of the cheque.                        |
| createdAt | uint256 | The creation timestamp of the cheque.            |
| expiresAt | uint256 | The expiration timestamp of the cheque.          |

### getTotalChequePayments

```solidity
function getTotalChequePayments() public view returns (uint256)
```

Returns total cheque payments. This is the sum of all cashed in cheques.

#### Return Values

| Name | Type    | Description                                         |
| ---- | ------- | --------------------------------------------------- |
| [0]  | uint256 | totalChequePayments The total cheque payments made. |

### isBotAllowed

```solidity
function isBotAllowed(address bot) public view virtual returns (bool)
```

Abstract function to check if a bot is allowed to sign cheques. This
must be implemented by the inheriting contract.

### getManagerAddress

```solidity
function getManagerAddress() public view virtual returns (address)
```

Abstract function to get the manager address. This must be implemented
by the inheriting contract.

## GasMoneyManager

GasMoneyManager manages gas money withdrawals for a {CMAccount}.

Gas money withdrawals are restricted to a withdrawal limit and period.

### GasMoneyStorage

```solidity
struct GasMoneyStorage {
    mapping(address => uint256) _withdrawalPeriodStart;
    mapping(address => uint256) _withdrawnAmount;
    uint256 _withdrawalLimit;
    uint256 _withdrawalPeriod;
}
```

### GasMoneyWithdrawal

```solidity
event GasMoneyWithdrawal(address withdrawer, uint256 amount)
```

Gas money withdrawal event

#### Parameters

| Name       | Type    | Description                   |
| ---------- | ------- | ----------------------------- |
| withdrawer | address | the address of the withdrawer |
| amount     | uint256 | the amount withdrawn          |

### GasMoneyWithdrawalUpdated

```solidity
event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
```

Gas money withdrawal limit and period updated event

#### Parameters

| Name   | Type    | Description                         |
| ------ | ------- | ----------------------------------- |
| limit  | uint256 | the withdrawal limit for the period |
| period | uint256 | the withdrawal period in seconds    |

### WithdrawalLimitExceeded

```solidity
error WithdrawalLimitExceeded(uint256 limit, uint256 amount)
```

### WithdrawalLimitExceededForPeriod

```solidity
error WithdrawalLimitExceededForPeriod(uint256 limit, uint256 amount)
```

### \_\_GasMoneyManager_init

```solidity
function __GasMoneyManager_init(uint256 withdrawalLimit, uint256 withdrawalPeriod) internal
```

### \_withdrawGasMoney

```solidity
function _withdrawGasMoney(uint256 amount) internal
```

Withdraws gas money.

This functions is intended to be called by the bot to withdraw gas money.
Inheriting contract should restrict who can call this with a public
function.

### \_setGasMoneyWithdrawal

```solidity
function _setGasMoneyWithdrawal(uint256 limit, uint256 period) internal
```

Sets the gas money withdrawal limit and period.

#### Parameters

| Name   | Type    | Description                         |
| ------ | ------- | ----------------------------------- |
| limit  | uint256 | the withdrawal limit for the period |
| period | uint256 | the withdrawal period in seconds    |

### getGasMoneyWithdrawal

```solidity
function getGasMoneyWithdrawal() public view returns (uint256 withdrawalLimit, uint256 withdrawalPeriod)
```

Returns the gas money withdrawal restrictions.

#### Return Values

| Name             | Type    | Description |
| ---------------- | ------- | ----------- |
| withdrawalLimit  | uint256 |             |
| withdrawalPeriod | uint256 |             |

### getGasMoneyWithdrawalForAccount

```solidity
function getGasMoneyWithdrawalForAccount(address account) public view returns (uint256 periodStart, uint256 withdrawnAmount)
```

Returns the gas money withdrawal details for an account.

#### Parameters

| Name    | Type    | Description            |
| ------- | ------- | ---------------------- |
| account | address | address of the account |

#### Return Values

| Name            | Type    | Description                              |
| --------------- | ------- | ---------------------------------------- |
| periodStart     | uint256 | timestamp of the withdrawal period start |
| withdrawnAmount | uint256 | amount withdrawn within the period       |

## BookingTokenOperator

Booking token operator contract is used by the {CMAccount} contract to mint
and buy booking tokens.

We made this a library so that we can use it in the {CMAccount} contract without
increasing the size of the contract.

### TokenApprovalFailed

```solidity
error TokenApprovalFailed(address token, address spender, uint256 amount)
```

_Token approval for the BookingToken address failed._

#### Parameters

| Name    | Type    | Description                                         |
| ------- | ------- | --------------------------------------------------- |
| token   | address | token address                                       |
| spender | address | spender address (the BookingToken contract address) |
| amount  | uint256 | amount of tokens to approve                         |

### mintBookingToken

```solidity
function mintBookingToken(address bookingToken, address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken) public
```

_Mints a booking token._

#### Parameters

| Name                | Type            | Description                                                                  |
| ------------------- | --------------- | ---------------------------------------------------------------------------- |
| bookingToken        | address         | booking token contract address                                               |
| reservedFor         | address         | address of the CM Account that can buy the token (generally the distributor) |
| uri                 | string          | URI of the token                                                             |
| expirationTimestamp | uint256         | expiration timestamp of the token in seconds                                 |
| price               | uint256         | price of the token                                                           |
| paymentToken        | contract IERC20 | payment token address                                                        |

### buyBookingToken

```solidity
function buyBookingToken(address bookingToken, uint256 tokenId) public
```

_Buys a booking token with the specified price and payment token in the
reservation._

#### Parameters

| Name         | Type    | Description                    |
| ------------ | ------- | ------------------------------ |
| bookingToken | address | booking token contract address |
| tokenId      | uint256 | token id                       |

## IBookingToken

### safeMintWithReservation

```solidity
function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken) external
```

### buyReservedToken

```solidity
function buyReservedToken(uint256 tokenId) external payable
```

### getReservationPrice

```solidity
function getReservationPrice(uint256 tokenId) external view returns (uint256 price, contract IERC20 paymentToken)
```

## ICMAccountManager

### getAccountImplementation

```solidity
function getAccountImplementation() external view returns (address)
```

### getDeveloperFeeBp

```solidity
function getDeveloperFeeBp() external view returns (uint256)
```

### getDeveloperWallet

```solidity
function getDeveloperWallet() external view returns (address)
```

### isCMAccount

```solidity
function isCMAccount(address account) external view returns (bool)
```

### getRegisteredServiceHashByName

```solidity
function getRegisteredServiceHashByName(string serviceName) external view returns (bytes32 serviceHash)
```

### getRegisteredServiceNameByHash

```solidity
function getRegisteredServiceNameByHash(bytes32 serviceHash) external view returns (string serviceName)
```

## PartnerConfiguration

Partner Configuration is used by the {CMAccount} contract to register
supported and wanted services by the partner.

### Service

Struct for storing supported service details for suppliers

```solidity
struct Service {
    uint256 _fee;
    bool _restrictedRate;
    string[] _capabilities;
}
```

### PaymentInfo

```solidity
struct PaymentInfo {
  bool _supportsOffChainPayment;
  struct EnumerableSet.AddressSet _supportedTokens;
}
```

### PartnerConfigurationStorage

```solidity
struct PartnerConfigurationStorage {
  struct EnumerableSet.Bytes32Set _servicesHashSet;
  mapping(bytes32 => struct PartnerConfiguration.Service) _supportedServices;
  struct PartnerConfiguration.PaymentInfo _paymentInfo;
  struct EnumerableSet.AddressSet _publicKeyAddressesSet;
  mapping(address => bytes) _publicKeys;
  struct EnumerableSet.Bytes32Set _wantedServicesHashSet;
}
```

### ServiceAlreadyExists

```solidity
error ServiceAlreadyExists(bytes32 serviceHash)
```

### ServiceDoesNotExist

```solidity
error ServiceDoesNotExist(bytes32 serviceHash)
```

### WantedServiceAlreadyExists

```solidity
error WantedServiceAlreadyExists(bytes32 serviceHash)
```

### WantedServiceDoesNotExist

```solidity
error WantedServiceDoesNotExist(bytes32 serviceHash)
```

### PaymentTokenAlreadyExists

```solidity
error PaymentTokenAlreadyExists(address token)
```

### PaymentTokenDoesNotExist

```solidity
error PaymentTokenDoesNotExist(address token)
```

### PublicKeyAlreadyExists

```solidity
error PublicKeyAlreadyExists(address pubKeyAddress)
```

### PublicKeyDoesNotExist

```solidity
error PublicKeyDoesNotExist(address pubKeyAddress)
```

### InvalidPublicKeyUseType

```solidity
error InvalidPublicKeyUseType(uint8 use)
```

### ServiceAdded

```solidity
event ServiceAdded(bytes32 serviceHash)
```

### ServiceRemoved

```solidity
event ServiceRemoved(bytes32 serviceHash)
```

### WantedServiceAdded

```solidity
event WantedServiceAdded(bytes32 serviceHash)
```

### WantedServiceRemoved

```solidity
event WantedServiceRemoved(bytes32 serviceHash)
```

### ServiceFeeUpdated

```solidity
event ServiceFeeUpdated(bytes32 serviceHash, uint256 fee)
```

### ServiceRestrictedRateUpdated

```solidity
event ServiceRestrictedRateUpdated(bytes32 serviceHash, bool restrictedRate)
```

### ServiceCapabilitiesUpdated

```solidity
event ServiceCapabilitiesUpdated(bytes32 serviceHash)
```

### ServiceCapabilityAdded

```solidity
event ServiceCapabilityAdded(bytes32 serviceHash, string capability)
```

### ServiceCapabilityRemoved

```solidity
event ServiceCapabilityRemoved(bytes32 serviceHash, string capability)
```

### PaymentTokenAdded

```solidity
event PaymentTokenAdded(address token)
```

### PaymentTokenRemoved

```solidity
event PaymentTokenRemoved(address token)
```

### OffChainPaymentSupportUpdated

```solidity
event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
```

### PublicKeyAdded

```solidity
event PublicKeyAdded(address pubKeyAddress)
```

### PublicKeyRemoved

```solidity
event PublicKeyRemoved(address pubKeyAddress)
```

### \_\_PartnerConfiguration_init

```solidity
function __PartnerConfiguration_init() internal
```

### \_\_PartnerConfiguration_init_unchained

```solidity
function __PartnerConfiguration_init_unchained() internal
```

### \_addService

```solidity
function _addService(bytes32 serviceHash, uint256 fee, string[] capabilities, bool restrictedRate) internal virtual
```

Adds a supported Service object for a given hash.

#### Parameters

| Name           | Type     | Description                                   |
| -------------- | -------- | --------------------------------------------- |
| serviceHash    | bytes32  | Hash of the service                           |
| fee            | uint256  | Fee for the service                           |
| capabilities   | string[] | Capabilities for the service                  |
| restrictedRate | bool     | If the service is restricted to pre-agreement |

### \_removeService

```solidity
function _removeService(bytes32 serviceHash) internal virtual
```

Removes a supported Service object for a given hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### \_setServiceFee

```solidity
function _setServiceFee(bytes32 serviceHash, uint256 fee) internal virtual
```

Sets the Service fee for a given hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |
| fee         | uint256 | Fee                 |

### \_setServiceRestrictedRate

```solidity
function _setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) internal virtual
```

Sets the Service restricted rate for a given hash.

#### Parameters

| Name           | Type    | Description         |
| -------------- | ------- | ------------------- |
| serviceHash    | bytes32 | Hash of the service |
| restrictedRate | bool    | Restricted rate     |

### \_setServiceCapabilities

```solidity
function _setServiceCapabilities(bytes32 serviceHash, string[] capabilities) internal virtual
```

Sets the Service capabilities for a given hash.

#### Parameters

| Name         | Type     | Description         |
| ------------ | -------- | ------------------- |
| serviceHash  | bytes32  | Hash of the service |
| capabilities | string[] | Capabilities        |

### \_addServiceCapability

```solidity
function _addServiceCapability(bytes32 serviceHash, string capability) internal virtual
```

Adds a capability to the service.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |
| capability  | string  | Capability          |

### \_removeServiceCapability

```solidity
function _removeServiceCapability(bytes32 serviceHash, string capability) internal virtual
```

Removes a capability from the service.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |
| capability  | string  | Capability          |

### getAllServiceHashes

```solidity
function getAllServiceHashes() public view returns (bytes32[] serviceHashes)
```

Returns all supported service hashes.

### getService

```solidity
function getService(bytes32 serviceHash) public view virtual returns (struct PartnerConfiguration.Service service)
```

Returns the Service object for a given hash. Service object contains fee and capabilities.

`serviceHash` is keccak256 hash of the pkg + service name as:

```text
           ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
keccak256("cmp.services.accommodation.v1alpha.AccommodationSearchService")
```

_These services are coming from the Camino Messenger Protocol's protobuf
definitions._

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getServiceFee

```solidity
function getServiceFee(bytes32 serviceHash) public view virtual returns (uint256 fee)
```

Returns the fee for a given service hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getServiceRestrictedRate

```solidity
function getServiceRestrictedRate(bytes32 serviceHash) public view virtual returns (bool restrictedRate)
```

Returns the restricted rate for a given service hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getServiceCapabilities

```solidity
function getServiceCapabilities(bytes32 serviceHash) public view virtual returns (string[] capabilities)
```

Returns the capabilities for a given service hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### \_addWantedService

```solidity
function _addWantedService(bytes32 serviceHash) internal virtual
```

Adds a wanted service hash to the wanted services set.

Reverts if the service already exists.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### \_removeWantedService

```solidity
function _removeWantedService(bytes32 serviceHash) internal virtual
```

Removes a wanted service hash from the wanted services set.

Reverts if the service does not exist.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getWantedServiceHashes

```solidity
function getWantedServiceHashes() public view virtual returns (bytes32[] serviceHashes)
```

Returns all wanted service hashes.

#### Return Values

| Name          | Type      | Description           |
| ------------- | --------- | --------------------- |
| serviceHashes | bytes32[] | Wanted service hashes |

### \_addSupportedToken

```solidity
function _addSupportedToken(address _token) internal virtual
```

Adds a supported payment token.

#### Parameters

| Name    | Type    | Description                       |
| ------- | ------- | --------------------------------- |
| \_token | address | Payment token address to be added |

### \_removeSupportedToken

```solidity
function _removeSupportedToken(address _token) internal virtual
```

Removes a supported payment token.

#### Parameters

| Name    | Type    | Description                         |
| ------- | ------- | ----------------------------------- |
| \_token | address | Payment token address to be removed |

### getSupportedTokens

```solidity
function getSupportedTokens() public view virtual returns (address[] tokens)
```

Returns supported token addresses.

#### Return Values

| Name   | Type      | Description               |
| ------ | --------- | ------------------------- |
| tokens | address[] | Supported token addresses |

### \_setOffChainPaymentSupported

```solidity
function _setOffChainPaymentSupported(bool _supportsOffChainPayment) internal virtual
```

Sets the off-chain payment support is supported.

### offChainPaymentSupported

```solidity
function offChainPaymentSupported() public view virtual returns (bool)
```

Returns true if off-chain payment is supported for the given service.

### \_addPublicKey

```solidity
function _addPublicKey(address pubKeyAddress, bytes publicKeyData) internal virtual
```

Adds public key with an address. Reverts if the public key already
exists.

Beware: This functions does not check if the public key is actually for the
given address.

### \_removePublicKey

```solidity
function _removePublicKey(address pubKeyAddress) internal virtual
```

Removes the public key for a given address

Reverts if the public key does not exist

### getPublicKeysAddresses

```solidity
function getPublicKeysAddresses() public view virtual returns (address[] pubKeyAddresses)
```

Returns the addresses of all public keys. These can then be used to
retrieve the public keys the `getPublicKey(address)` function.

### getPublicKey

```solidity
function getPublicKey(address pubKeyAddress) public view virtual returns (bytes data)
```

Returns the public key for a given address.

Reverts if the public key does not exist

#### Parameters

| Name          | Type    | Description               |
| ------------- | ------- | ------------------------- |
| pubKeyAddress | address | Address of the public key |

## ICMAccount

### initialize

```solidity
function initialize(address manager, address bookingToken, uint256 prefundAmount, address owner, address upgrader) external
```

## BookingToken

Booking Token contract represents a booking done on the Camino Messenger.

Suppliers can mint Booking Tokens and reserve them for a distributor address to
buy.

Booking Tokens can have zero price, meaning that the payment will be done
off-chain.

When a token is minted with a reservation, it can not transferred until the
expiration timestamp is reached or the token is bought.

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

Upgrader role can upgrade the contract to a new implementation.

### MIN_EXPIRATION_ADMIN_ROLE

```solidity
bytes32 MIN_EXPIRATION_ADMIN_ROLE
```

This role can set the mininum allowed expiration timestamp difference.

### TokenReservation

```solidity
struct TokenReservation {
  address reservedFor;
  address supplier;
  uint256 expirationTimestamp;
  uint256 price;
  contract IERC20 paymentToken;
}
```

### BookingTokenStorage

```solidity
struct BookingTokenStorage {
  address _manager;
  uint256 _nextTokenId;
  uint256 _minExpirationTimestampDiff;
  mapping(uint256 => struct BookingToken.TokenReservation) _reservations;
}
```

### TokenReserved

```solidity
event TokenReserved(uint256 tokenId, address reservedFor, address supplier, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken)
```

Event emitted when a token is reserved.

#### Parameters

| Name                | Type            | Description           |
| ------------------- | --------------- | --------------------- |
| tokenId             | uint256         | token id              |
| reservedFor         | address         | reserved for address  |
| supplier            | address         | supplier address      |
| expirationTimestamp | uint256         | expiration timestamp  |
| price               | uint256         | price of the token    |
| paymentToken        | contract IERC20 | payment token address |

### TokenBought

```solidity
event TokenBought(uint256 tokenId, address buyer)
```

Event emitted when a token is bought.

#### Parameters

| Name    | Type    | Description   |
| ------- | ------- | ------------- |
| tokenId | uint256 | token id      |
| buyer   | address | buyer address |

### ExpirationTimestampTooSoon

```solidity
error ExpirationTimestampTooSoon(uint256 expirationTimestamp, uint256 minExpirationTimestampDiff)
```

Error for expiration timestamp too soon. It must be at least
`_minExpirationTimestampDiff` seconds in the future.

### NotCMAccount

```solidity
error NotCMAccount(address account)
```

Address is not a CM Account.

#### Parameters

| Name    | Type    | Description     |
| ------- | ------- | --------------- |
| account | address | account address |

### ReservationMismatch

```solidity
error ReservationMismatch(address reservedFor, address buyer)
```

ReservedFor and buyer mismatch.

#### Parameters

| Name        | Type    | Description          |
| ----------- | ------- | -------------------- |
| reservedFor | address | reserved for address |
| buyer       | address | buyer address        |

### ReservationExpired

```solidity
error ReservationExpired(uint256 tokenId, uint256 expirationTimestamp)
```

Reservation expired.

#### Parameters

| Name                | Type    | Description          |
| ------------------- | ------- | -------------------- |
| tokenId             | uint256 | token id             |
| expirationTimestamp | uint256 | expiration timestamp |

### IncorrectPrice

```solidity
error IncorrectPrice(uint256 price, uint256 reservationPrice)
```

Incorrect price.

#### Parameters

| Name             | Type    | Description        |
| ---------------- | ------- | ------------------ |
| price            | uint256 | price of the token |
| reservationPrice | uint256 | reservation price  |

### SupplierIsNotOwner

```solidity
error SupplierIsNotOwner(uint256 tokenId, address supplier)
```

Supplier is not the owner.

#### Parameters

| Name     | Type    | Description      |
| -------- | ------- | ---------------- |
| tokenId  | uint256 | token id         |
| supplier | address | supplier address |

### TokenIsReserved

```solidity
error TokenIsReserved(uint256 tokenId, address reservedFor)
```

Token is reserved and can not be transferred.

#### Parameters

| Name        | Type    | Description          |
| ----------- | ------- | -------------------- |
| tokenId     | uint256 | token id             |
| reservedFor | address | reserved for address |

### InsufficientAllowance

```solidity
error InsufficientAllowance(address sender, contract IERC20 paymentToken, uint256 price, uint256 allowance)
```

Insufficient allowance to transfer the ERC20 token to the supplier.

#### Parameters

| Name         | Type            | Description           |
| ------------ | --------------- | --------------------- |
| sender       | address         | msg.sender            |
| paymentToken | contract IERC20 | payment token address |
| price        | uint256         | price of the token    |
| allowance    | uint256         | allowance amount      |

### onlyCMAccount

```solidity
modifier onlyCMAccount(address account)
```

Only CMAccount modifier.

### initialize

```solidity
function initialize(address manager, address defaultAdmin, address upgrader) public
```

### \_authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

Function to authorize an upgrade for UUPS proxy.

### safeMintWithReservation

```solidity
function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken) public
```

Mints a new token with a reservation for a specific address.

#### Parameters

| Name                | Type            | Description                                                           |
| ------------------- | --------------- | --------------------------------------------------------------------- |
| reservedFor         | address         | The CM Account address that can buy the token                         |
| uri                 | string          | The URI of the token                                                  |
| expirationTimestamp | uint256         | The expiration timestamp                                              |
| price               | uint256         | The price of the token                                                |
| paymentToken        | contract IERC20 | The token used to pay for the reservation. If address(0) then native. |

### buyReservedToken

```solidity
function buyReservedToken(uint256 tokenId) external payable
```

Buys a reserved token. The reservation must be for the message sender.

Also the message sender should set allowance for the payment token to this
contract to at least the reservation price. (only for ERC20 tokens)

For native coin, the message sender should send the exact amount.

Only CM Accounts can call this function

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

### \_reserve

```solidity
function _reserve(uint256 tokenId, address reservedFor, address supplier, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken) internal
```

Reserve a token for a specific address with an expiration timestamp

### checkTransferable

```solidity
function checkTransferable(uint256 tokenId) internal
```

Check if the token is transferable

### isCMAccount

```solidity
function isCMAccount(address account) public view returns (bool)
```

Checks if an address is a CM Account.

#### Parameters

| Name    | Type    | Description          |
| ------- | ------- | -------------------- |
| account | address | The address to check |

#### Return Values

| Name | Type | Description                         |
| ---- | ---- | ----------------------------------- |
| [0]  | bool | true if the address is a CM Account |

### requireCMAccount

```solidity
function requireCMAccount(address account) internal view
```

Checks if the address is a CM Account and reverts if not.

#### Parameters

| Name    | Type    | Description          |
| ------- | ------- | -------------------- |
| account | address | The address to check |

### setManagerAddress

```solidity
function setManagerAddress(address manager) public
```

Sets for the manager address.

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| manager | address | The address of the manager |

### getManagerAddress

```solidity
function getManagerAddress() public view returns (address)
```

Returns for the manager address.

### setMinExpirationTimestampDiff

```solidity
function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) public
```

Sets minimum expiration timestamp difference in seconds.

#### Parameters

| Name                       | Type    | Description                                        |
| -------------------------- | ------- | -------------------------------------------------- |
| minExpirationTimestampDiff | uint256 | Minimum expiration timestamp difference in seconds |

### getMinExpirationTimestampDiff

```solidity
function getMinExpirationTimestampDiff() public view returns (uint256)
```

Returns minimum expiration timestamp difference in seconds.

### getReservationPrice

```solidity
function getReservationPrice(uint256 tokenId) public view returns (uint256 price, contract IERC20 paymentToken)
```

Returns the token reservation price for a specific token.

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

### transferFrom

```solidity
function transferFrom(address from, address to, uint256 tokenId) public
```

Override transferFrom to check if token is reserved. It reverts if
the token is reserved.

### safeTransferFrom

```solidity
function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) public
```

Override safeTransferFrom to check if token is reserved. It reverts if
the token is reserved.

### \_update

```solidity
function _update(address to, uint256 tokenId, address auth) internal returns (address)
```

### \_increaseBalance

```solidity
function _increaseBalance(address account, uint128 value) internal
```

### tokenURI

```solidity
function tokenURI(uint256 tokenId) public view returns (string)
```

### supportsInterface

```solidity
function supportsInterface(bytes4 interfaceId) public view returns (bool)
```

## CMAccountManager

This contract manages the creation of the Camino Messenger accounts by
deploying {ERC1967Proxy} proxies that point to the{CMAccount} implementation
address.

Create CM Account: Users who want to create an account should call
`createCMAccount(address admin, address upgrader)` function with addresses of
the accounts admin and upgrader roles and also send the pre fund amount,
which is currently set as 100 CAMs. When the manager contract is paused,
account creation is stopped.

Developer Fee: This contracts also keeps the info about the developer wallet
and fee basis points. Which are used during the cheque cash in to pay for the
developer fee.

Service Registry: {CMAccountManager} also acts as a registry for the services
that {CMAccount} contracts add as a supported or wanted service. Registry
works by hashing (keccak256) the service name (string) and creating a mapping
as keccak256(serviceName) => serviceName. And provides functions that
{CMAccount} function uses to register services. The {CMAccount} only keeps
the hashes (byte32) of the registered services.

### PAUSER_ROLE

```solidity
bytes32 PAUSER_ROLE
```

Pauser role can pause the contract. Currently this only affects the
creation of CM Accounts. When paused, account creation is stopped.

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

Upgrader role can upgrade the contract to a new implementation.

### VERSIONER_ROLE

```solidity
bytes32 VERSIONER_ROLE
```

Versioner role can set new {CMAccount} implementation address. When a
new implementation address is set, it is used for the new {CMAccount}
creations.

The old {CMAccount} contracts are not affected by this. Owners of those
should do the upgrade manually by calling the `upgradeToAndCall(address)`
function on the account.

### FEE_ADMIN_ROLE

```solidity
bytes32 FEE_ADMIN_ROLE
```

Fee admin role can set the developer fee basis points which used for
calculating the developer fee that is cut from the cheque payments.

### DEVELOPER_WALLET_ADMIN_ROLE

```solidity
bytes32 DEVELOPER_WALLET_ADMIN_ROLE
```

Developer wallet admin role can set the developer wallet address
which is used to receive the developer fee.

### PREFUND_ADMIN_ROLE

```solidity
bytes32 PREFUND_ADMIN_ROLE
```

Prefund admin role can set the mandatory prefund amount for {CMAccount}
contracts.

### SERVICE_REGISTRY_ADMIN_ROLE

```solidity
bytes32 SERVICE_REGISTRY_ADMIN_ROLE
```

Service registry admin role can add and remove services to the service
registry mapping. Implemented by {ServiceRegistry} contract.

### CMACCOUNT_ROLE

```solidity
bytes32 CMACCOUNT_ROLE
```

This role is granted to the created CM Accounts. It is used to keep
an enumerable list of CM Accounts.

### CMAccountInfo

CMAccount info struct, to keep track of created CM Accounts and their
creators.

```solidity
struct CMAccountInfo {
    bool isCMAccount;
    address creator;
}
```

### CMAccountManagerStorage

```solidity
struct CMAccountManagerStorage {
  address _latestAccountImplementation;
  uint256 _prefundAmount;
  address _developerWallet;
  uint256 _developerFeeBp;
  address _bookingToken;
  mapping(address => struct CMAccountManager.CMAccountInfo) _cmAccountInfo;
}
```

### CMAccountCreated

```solidity
event CMAccountCreated(address account)
```

CM Account created event.

#### Parameters

| Name    | Type    | Description                      |
| ------- | ------- | -------------------------------- |
| account | address | The address of the new CMAccount |

### CMAccountImplementationUpdated

```solidity
event CMAccountImplementationUpdated(address oldImplementation, address newImplementation)
```

CM Account implementation address updated event.

#### Parameters

| Name              | Type    | Description                    |
| ----------------- | ------- | ------------------------------ |
| oldImplementation | address | The old implementation address |
| newImplementation | address | The new implementation address |

### DeveloperWalletUpdated

```solidity
event DeveloperWalletUpdated(address oldDeveloperWallet, address newDeveloperWallet)
```

Developer wallet address updated event.

#### Parameters

| Name               | Type    | Description                      |
| ------------------ | ------- | -------------------------------- |
| oldDeveloperWallet | address | The old developer wallet address |
| newDeveloperWallet | address | The new developer wallet address |

### DeveloperFeeBpUpdated

```solidity
event DeveloperFeeBpUpdated(uint256 oldDeveloperFeeBp, uint256 newDeveloperFeeBp)
```

Developer fee basis points updated event.

#### Parameters

| Name              | Type    | Description                        |
| ----------------- | ------- | ---------------------------------- |
| oldDeveloperFeeBp | uint256 | The old developer fee basis points |
| newDeveloperFeeBp | uint256 | The new developer fee basis points |

### BookingTokenAddressUpdated

```solidity
event BookingTokenAddressUpdated(address oldBookingToken, address newBookingToken)
```

Booking token address updated event.

#### Parameters

| Name            | Type    | Description                   |
| --------------- | ------- | ----------------------------- |
| oldBookingToken | address | The old booking token address |
| newBookingToken | address | The new booking token address |

### CMAccountInvalidImplementation

```solidity
error CMAccountInvalidImplementation(address implementation)
```

The implementation of the CMAccount is invalid.

#### Parameters

| Name           | Type    | Description                                 |
| -------------- | ------- | ------------------------------------------- |
| implementation | address | The implementation address of the CMAccount |

### CMAccountInvalidAdmin

```solidity
error CMAccountInvalidAdmin(address admin)
```

The admin address is invalid.

#### Parameters

| Name  | Type    | Description       |
| ----- | ------- | ----------------- |
| admin | address | The admin address |

### InvalidDeveloperWallet

```solidity
error InvalidDeveloperWallet(address developerWallet)
```

Invalid developer address.

#### Parameters

| Name            | Type    | Description                  |
| --------------- | ------- | ---------------------------- |
| developerWallet | address | The developer wallet address |

### InvalidBookingTokenAddress

```solidity
error InvalidBookingTokenAddress(address bookingToken)
```

Invalid booking token address.

#### Parameters

| Name         | Type    | Description               |
| ------------ | ------- | ------------------------- |
| bookingToken | address | The booking token address |

### IncorrectPrefundAmount

```solidity
error IncorrectPrefundAmount(uint256 expected, uint256 sended)
```

Incorrect pre fund amount.

#### Parameters

| Name     | Type    | Description                  |
| -------- | ------- | ---------------------------- |
| expected | uint256 | The expected pre fund amount |
| sended   | uint256 |                              |

### constructor

```solidity
constructor() public
```

### initialize

```solidity
function initialize(address defaultAdmin, address pauser, address upgrader, address versioner, address developerWallet, uint256 developerFeeBp) public
```

### pause

```solidity
function pause() public
```

Pauses the CMAccountManager contract. Currently this only affects the
creation of CMAccount. When paused, account creation is stopped.

### unpause

```solidity
function unpause() public
```

Unpauses the CMAccountManager contract.

### \_authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

Authorization for the CMAccountManager contract upgrade.

### createCMAccount

```solidity
function createCMAccount(address admin, address upgrader) external payable returns (address)
```

Creates CMAccount by deploying a ERC1967Proxy with the CMAccount
implementation from the manager.

Because this function is deploying a contract, it reverts if the caller is
not KYC or KYB verified. (For EOAs only)

Caller must send the pre-fund amount with the transaction.

_Emits a {CMAccountCreated} event._

### \_setCMAccountInfo

```solidity
function _setCMAccountInfo(address account, struct CMAccountManager.CMAccountInfo info) internal
```

### getCMAccountCreator

```solidity
function getCMAccountCreator(address account) public view returns (address)
```

Returns the given account's creator.

#### Parameters

| Name    | Type    | Description         |
| ------- | ------- | ------------------- |
| account | address | The account address |

### isCMAccount

```solidity
function isCMAccount(address account) public view returns (bool)
```

Check if an address is CMAccount created by the manager.

#### Parameters

| Name    | Type    | Description                  |
| ------- | ------- | ---------------------------- |
| account | address | The account address to check |

### getAccountImplementation

```solidity
function getAccountImplementation() public view returns (address)
```

Returns the CMAccount implementation address.

### setAccountImplementation

```solidity
function setAccountImplementation(address newImplementation) public
```

Set a new CMAccount implementation address.

#### Parameters

| Name              | Type    | Description                    |
| ----------------- | ------- | ------------------------------ |
| newImplementation | address | The new implementation address |

### \_setAccountImplementation

```solidity
function _setAccountImplementation(address newImplementation) internal
```

### getPrefundAmount

```solidity
function getPrefundAmount() public view returns (uint256)
```

Returns the prefund amount.

### setPrefundAmount

```solidity
function setPrefundAmount(uint256 newPrefundAmount) public
```

Sets the prefund amount.

### getBookingTokenAddress

```solidity
function getBookingTokenAddress() public view returns (address)
```

Returns the booking token address.

### setBookingTokenAddress

```solidity
function setBookingTokenAddress(address token) public
```

Sets booking token address.

### \_setBookingTokenAddress

```solidity
function _setBookingTokenAddress(address token) internal
```

### getDeveloperWallet

```solidity
function getDeveloperWallet() public view returns (address developerWallet)
```

Returns developer wallet address.

### setDeveloperWallet

```solidity
function setDeveloperWallet(address developerWallet) public
```

Sets developer wallet address.

### getDeveloperFeeBp

```solidity
function getDeveloperFeeBp() public view returns (uint256 developerFeeBp)
```

Returns developer fee in basis points.

### setDeveloperFeeBp

```solidity
function setDeveloperFeeBp(uint256 bp) public
```

Sets developer fee in basis points.

A basis point (bp) is one hundredth of 1 percentage point.

1 bp = 0.01%, 1/10,000⁠, or 0.0001.
10 bp = 0.1%, 1/1,000⁠, or 0.001.
100 bp = 1%, ⁠1/100⁠, or 0.01.

### registerService

```solidity
function registerService(string serviceName) public
```

Registers a given service name. CM Accounts can only register services
if they are also registered in the service registry on the manager contract.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### unregisterService

```solidity
function unregisterService(string serviceName) public
```

Unregisters a given service name. CM Accounts will not be able to register
the service anymore.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

## CMAccountManagerV2

### getVersion

```solidity
function getVersion() public pure returns (string)
```

## ServiceRegistry

Service registry is used by the {CMAccountManager} contract to register
services by hashing (keccak256) the service name (string) and creating a mapping
as keccak256(serviceName) => serviceName.

### ServiceRegistryStorage

```solidity
struct ServiceRegistryStorage {
  struct EnumerableSet.Bytes32Set _servicesHashSet;
  mapping(bytes32 => string) _serviceNameByHash;
  mapping(string => bytes32) _hashByServiceName;
}
```

### ServiceRegistered

```solidity
event ServiceRegistered(string serviceName, bytes32 serviceHash)
```

### ServiceUnregistered

```solidity
event ServiceUnregistered(string serviceName, bytes32 serviceHash)
```

### ServiceAlreadyRegistered

```solidity
error ServiceAlreadyRegistered(string serviceName)
```

### ServiceNotRegistered

```solidity
error ServiceNotRegistered()
```

### \_\_ServiceRegistry_init

```solidity
function __ServiceRegistry_init() internal
```

### \_\_ServiceRegistry_init_unchained

```solidity
function __ServiceRegistry_init_unchained() internal
```

### \_registerServiceName

```solidity
function _registerServiceName(string serviceName) internal virtual
```

Adds a new service by its name. This function calculates the hash of the
service name and adds it to the registry

{serviceName} is the pkg + service name as:

```text
 ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
"cmp.services.accommodation.v1alpha.AccommodationSearchService"
```

_These services are coming from the Camino Messenger Protocol's protobuf
definitions._

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### \_unregisterServiceName

```solidity
function _unregisterServiceName(string serviceName) internal virtual
```

Removes a service by its name. This function calculates the hash of the
service name and removes it from the registry.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### getRegisteredServiceNameByHash

```solidity
function getRegisteredServiceNameByHash(bytes32 serviceHash) public view returns (string serviceName)
```

Returns the name of a service by its hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getRegisteredServiceHashByName

```solidity
function getRegisteredServiceHashByName(string serviceName) public view returns (bytes32 serviceHash)
```

Returns the hash of a service by its name.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### getAllRegisteredServiceHashes

```solidity
function getAllRegisteredServiceHashes() public view returns (bytes32[] services)
```

Returns all registered service **hashes**.

#### Return Values

| Name     | Type      | Description                   |
| -------- | --------- | ----------------------------- |
| services | bytes32[] | All registered service hashes |

### getAllRegisteredServiceNames

```solidity
function getAllRegisteredServiceNames() public view returns (string[] services)
```

Returns all registered service **names**.

#### Return Values

| Name     | Type     | Description                  |
| -------- | -------- | ---------------------------- |
| services | string[] | All registered service names |

## Dummy

### getVersion

```solidity
function getVersion() public pure returns (string)
```

## NullUSD

### constructor

```solidity
constructor() public
```

## ICaminoAdmin

### getKycState

```solidity
function getKycState(address account) external view returns (uint256)
```

## KYCUtils

### ADMIN_ADDR

```solidity
address ADMIN_ADDR
```

Admin contract address

### KYC_VERIFIED

```solidity
uint256 KYC_VERIFIED
```

Constants for KYC states

### KYC_EXPIRED

```solidity
uint256 KYC_EXPIRED
```

### KYB_VERIFIED

```solidity
uint256 KYB_VERIFIED
```

### NotKYCVerified

```solidity
error NotKYCVerified(address account)
```

Errors

### NotKYBVerified

```solidity
error NotKYBVerified(address account)
```

### NotVerified

```solidity
error NotVerified(address account)
```

### getKYCState

```solidity
function getKYCState(address account) internal view returns (uint256)
```

_Returns KYC state from the CaminoAdmin contract_

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| account | address | address to check the state |

### isKYCVerified

```solidity
function isKYCVerified(address account) internal view returns (bool)
```

_Returns true if the address is KYC verified_

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| account | address | address to check the state |

### isKYBVerified

```solidity
function isKYBVerified(address account) internal view returns (bool)
```

_Returns true if the address is KYB verified_

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| account | address | address to check the state |

### isVerified

```solidity
function isVerified(address account) internal view returns (bool)
```

_Returns true if the address is KYC or KYB verified_

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| account | address | address to check the state |

### requireKYCVerified

```solidity
function requireKYCVerified(address account) internal view
```

_Reverts with `NotKYCVerified(account)` if the account is not KYC verified._

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| account | address | address to check the state |

### requireKYBVerified

```solidity
function requireKYBVerified(address account) internal view
```

_Reverts with `NotKYBVerified(account)` if the account is not KYB verified._

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| account | address | address to check the state |

### requireVerified

```solidity
function requireVerified(address account) internal view
```

_Reverts with `NotVerified(account)` if the account is not KYC or KYB verified._

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| account | address | address to check the state |
