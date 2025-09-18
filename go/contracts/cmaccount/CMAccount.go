// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cmaccount

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PartnerConfigurationService is an auto generated low-level Go binding around an user-defined struct.
type PartnerConfigurationService struct {
	Fee            *big.Int
	RestrictedRate bool
	Capabilities   []string
}

// CmaccountMetaData contains all meta data concerning the Cmaccount contract.
var CmaccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"latestImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountImplementationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountNoUpgradeNeeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"name\":\"ChequeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"IncorrectValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chequeCounter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastCounter\",\"type\":\"uint256\"}],\"name\":\"InvalidCounter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromCMAccount\",\"type\":\"address\"}],\"name\":\"InvalidFromCMAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"InvalidPaymentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"use\",\"type\":\"uint8\"}],\"name\":\"InvalidPublicKeyUseType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toCMAccount\",\"type\":\"address\"}],\"name\":\"InvalidToCMAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"NotAllowedToSignCheques\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefundLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrefundNotSpentYet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedNativePayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceededForPeriod\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"CMAccountUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromCMAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toCMAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromBot\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toBot\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidChequeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidDeveloperFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"ChequeCashedIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportsOffChainPayment\",\"type\":\"bool\"}],\"name\":\"OffChainPaymentSupportUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceCapabilitiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ServiceFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"ServiceRestrictedRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOKING_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHEQUE_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER_CHEQUE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVICE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasMoney\",\"type\":\"uint256\"}],\"name\":\"addMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"addService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"addServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"addWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"buyBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toBot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"cashInCheque\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasMoneyWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGasMoneyWithdrawalForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromBot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toBot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getLastCashIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastCounter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastExpiresAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKeysAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pubKeyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getService\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service\",\"name\":\"service\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalChequePayments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"getTotalChequePaymentsPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toBot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"hashMessengerCheque\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toBot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"hashTypedDataV4\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"isBotAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"mintBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offChainPaymentSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"removeMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"removePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"removeServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"removeSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"removeWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGasMoneyWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSupported\",\"type\":\"bool\"}],\"name\":\"setOffChainPaymentSupported\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"setServiceCapabilities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"setServiceRestrictedRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toCMAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toBot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifyCheque\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"reason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGasMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516157a4620001046000396000818161308e015281816130b701526133d201526157a46000f3fe6080604052600436106103a85760003560e01c806301ffc9a7146103b457806308564c19146103e95780630ede80d61461040b578063136f50ca1461043b578063150b7a021461045d57806318274da4146104a15780631aca6376146104c15780631c54f0f7146104e35780631c5db99e1461050357806320606b7014610523578063241bbbfc14610545578063248a9ca31461055a5780632a1193801461057a5780632f2ff15d1461059a578063319d13f3146105ba57806333746274146105da57806336568abe146105fc578063372e802b1461061c578063383aba871461063c57806339e4c7051461065e57806341bf7c691461067e57806342072bbd1461069e578063432cf639146106b3578063457006db146106d35780634f1ef286146107015780634f3f4639146107145780635008c8ec1461073657806351889d6b1461075657806352d1902d146107765780635c9889941461078b5780635e07f869146107ab57806363e86cc8146107cb578063658db0af146107ed5780636d69fcaf146108105780636fc22cd114610830578063705977531461085057806374aa20481461088057806374fe60e9146108a05780637512e55b146108c057806376319190146108e05780637eec56c714610900578063852b3ccb14610923578063857cdbb81461094557806385f438c1146109725780638c20f574146109945780638f69347d146109b45780639010d07c146109d457806391d14854146109f45780639db5dbe414610a14578063a217fddf14610a34578063a31aa03914610a49578063a3246ad314610a69578063a73ebdd814610a96578063a7d022f814610ab6578063ad3cb1cc14610ad6578063b512463514610b07578063b82923fb14610b27578063be66718814610b3c578063c162d7da14610b5c578063c6640e6814610b71578063ca15c87314610b91578063ccde65dc14610bb1578063cd9ef91414610bd1578063d09445c214610bf1578063d3c7c2c714610c13578063d547741f14610c28578063da47d85614610c48578063e0b78add14610c75578063e26a61bb14610c95578063e5a6725c14610cb5578063e7bfce9a14610cd5578063e96cf7ad14610cf5578063ea79d07a14610d0a578063eb5ea27314610d1f578063ebc20d2014610d3f578063ecaa76ef14610d5f578063ed24911d14610d7f578063ee3b641f14610d94578063f3fef3a314610db4578063f51acaea14610dd4578063f72c0d8b14610df4578063f7e45f0914610e16578063f8c8765e14610e3657600080fd5b366103af57005b600080fd5b3480156103c057600080fd5b506103d46103cf3660046145fc565b610e56565b60405190151581526020015b60405180910390f35b3480156103f557600080fd5b506103fe610e81565b6040516103e091906146d0565b34801561041757600080fd5b5061042d60008051602061574f83398151915281565b6040519081526020016103e0565b34801561044757600080fd5b50610450610fab565b6040516103e091906146e3565b34801561046957600080fd5b506104886104783660046147ff565b630a85bd0160e11b949350505050565b6040516001600160e01b031990911681526020016103e0565b3480156104ad57600080fd5b5061042d6104bc36600461486a565b610fcb565b3480156104cd57600080fd5b506104e16104dc36600461489e565b610fd9565b005b3480156104ef57600080fd5b506104e16104fe3660046148df565b611086565b34801561050f57600080fd5b506104e161051e36600461499f565b61112a565b34801561052f57600080fd5b5061042d60008051602061570f83398151915281565b34801561055157600080fd5b506103d46111e2565b34801561056657600080fd5b5061042d6105753660046149d3565b6111fa565b34801561058657600080fd5b506104e1610595366004614a03565b61121a565b3480156105a657600080fd5b506104e16105b5366004614a3f565b6112c0565b3480156105c657600080fd5b506103fe6105d536600461486a565b6112e2565b3480156105e657600080fd5b5061042d6000805160206156cf83398151915281565b34801561060857600080fd5b506104e1610617366004614a3f565b6112f0565b34801561062857600080fd5b5061042d610637366004614a6f565b611323565b34801561064857600080fd5b5061042d60008051602061562f83398151915281565b34801561066a57600080fd5b506104e161067936600461499f565b611380565b34801561068a57600080fd5b506104e1610699366004614af4565b611426565b3480156106aa57600080fd5b5061045061149e565b3480156106bf57600080fd5b506104e16106ce366004614b46565b6114b5565b3480156106df57600080fd5b506106f36106ee366004614bbb565b611526565b6040516103e0929190614c69565b6104e161070f366004614c82565b6117e1565b34801561072057600080fd5b50610729611800565b6040516103e09190614cd1565b34801561074257600080fd5b5061042d610751366004614a6f565b61181e565b34801561076257600080fd5b506104e1610771366004614ce5565b61188a565b34801561078257600080fd5b5061042d61195b565b34801561079757600080fd5b506104e16107a63660046149d3565b611978565b3480156107b757600080fd5b506103fe6107c63660046149d3565b6119ad565b3480156107d757600080fd5b5061042d60008051602061564f83398151915281565b3480156107f957600080fd5b50610802611ab5565b6040516103e0929190614d11565b34801561081c57600080fd5b506104e161082b366004614d1f565b611ad7565b34801561083c57600080fd5b506104e161084b3660046148df565b611af8565b34801561085c57600080fd5b5061087061086b366004614d3c565b611b1a565b6040516103e09493929190614d87565b34801561088c57600080fd5b506104e161089b366004614da2565b611b7f565b3480156108ac57600080fd5b506104e16108bb366004614a03565b611c37565b3480156108cc57600080fd5b506104e16108db366004614de8565b611c71565b3480156108ec57600080fd5b506104e16108fb366004614d1f565b611ce0565b34801561090c57600080fd5b50610915611d01565b6040516103e0929190614ebf565b34801561092f57600080fd5b5061042d60008051602061572f83398151915281565b34801561095157600080fd5b50610965610960366004614d1f565b611ec6565b6040516103e09190614f31565b34801561097e57600080fd5b5061042d6000805160206156ef83398151915281565b3480156109a057600080fd5b506104e16109af366004614de8565b611fab565b3480156109c057600080fd5b506103d46109cf3660046149d3565b61201a565b3480156109e057600080fd5b506107296109ef3660046148df565b61204b565b348015610a0057600080fd5b506103d4610a0f366004614a3f565b612079565b348015610a2057600080fd5b506104e1610a2f36600461489e565b6120af565b348015610a4057600080fd5b5061042d600081565b348015610a5557600080fd5b506104e1610a64366004614f44565b612102565b348015610a7557600080fd5b50610a89610a843660046149d3565b612123565b6040516103e09190614f61565b348015610aa257600080fd5b5061042d610ab1366004614d1f565b612150565b348015610ac257600080fd5b506104e1610ad1366004614fa2565b61217d565b348015610ae257600080fd5b50610965604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610b1357600080fd5b506103d4610b2236600461486a565b6121ee565b348015610b3357600080fd5b506104e16121fc565b348015610b4857600080fd5b506104e1610b573660046148df565b612297565b348015610b6857600080fd5b506107296122d1565b348015610b7d57600080fd5b506104e1610b8c366004614d1f565b6122ec565b348015610b9d57600080fd5b5061042d610bac3660046149d3565b612387565b348015610bbd57600080fd5b506104e1610bcc366004614c82565b6123ac565b348015610bdd57600080fd5b506104e1610bec366004614fe8565b6123ce565b348015610bfd57600080fd5b5061042d6000805160206156af83398151915281565b348015610c1f57600080fd5b50610a89612486565b348015610c3457600080fd5b506104e1610c43366004614a3f565b6124a0565b348015610c5457600080fd5b50610c68610c633660046149d3565b6124bc565b6040516103e09190615016565b348015610c8157600080fd5b506103d4610c90366004614d1f565b6125e5565b348015610ca157600080fd5b506104e1610cb0366004615029565b6125ff565b348015610cc157600080fd5b506104e1610cd03660046149d3565b61269d565b348015610ce157600080fd5b506104e1610cf0366004614d1f565b612729565b348015610d0157600080fd5b5061042d61274a565b348015610d1657600080fd5b50610a8961275f565b348015610d2b57600080fd5b5061042d610d3a3660046149d3565b612779565b348015610d4b57600080fd5b506104e1610d5a3660046150bc565b6127a4565b348015610d6b57600080fd5b506104e1610d7a366004614bbb565b612811565b348015610d8b57600080fd5b5061042d612a38565b348015610da057600080fd5b50610802610daf366004614d1f565b612a4d565b348015610dc057600080fd5b506104e1610dcf366004614ce5565b612a88565b348015610de057600080fd5b506104e1610def36600461486a565b612b2e565b348015610e0057600080fd5b5061042d60008051602061566f83398151915281565b348015610e2257600080fd5b506104e1610e31366004614da2565b612b82565b348015610e4257600080fd5b506104e1610e51366004615115565b612bbc565b60006001600160e01b03198216635a05180f60e01b1480610e7b5750610e7b82612d6f565b92915050565b60606000610e8d610fab565b9050600081516001600160401b03811115610eaa57610eaa61473c565b604051908082528060200260200182016040528015610edd57816020015b6060815260200190600190039081610ec85790505b50905060005b8251811015610fa457610ef46122d1565b6001600160a01b0316635a81a626848381518110610f1457610f14615171565b60200260200101516040518263ffffffff1660e01b8152600401610f3a91815260200190565b600060405180830381865afa158015610f57573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f7f9190810190615187565b828281518110610f9157610f91615171565b6020908102919091010152600101610ee3565b5092915050565b60606000610fb7612da4565b9050610fc581600901612dc8565b91505090565b6000610e7b610d3a83612dd5565b6000805160206156ef833981519152610ff181612e4b565b6001600160a01b03831661101857604051633a954ecd60e21b815260040160405180910390fd5b604051632142170760e11b81523060048201526001600160a01b038481166024830152604482018490528516906342842e0e90606401600060405180830381600087803b15801561106857600080fd5b505af115801561107c573d6000803e3d6000fd5b5050505050505050565b60008051602061572f83398151915261109e81612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6321b87f3a6110c0611800565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018690526044810185905260640160006040518083038186803b15801561110d57600080fd5b505af4158015611121573d6000803e3d6000fd5b50505050505050565b6000805160206156af83398151915261114281612e4b565b60005b82518110156111dd57600061117284838151811061116557611165615171565b6020026020010151612dd5565b905061117d81612e55565b83828151811061118f5761118f615171565b60200260200101516040516111a491906151f4565b604051908190038120907f50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f890600090a250600101611145565b505050565b6000806111ed612da4565b6003015460ff1692915050565b600080611205612e93565b60009384526020525050604090206001015490565b60008051602061572f83398151915261123281612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6307e47316611254611800565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024810187905261ffff80871660448301528516606482015260840160006040518083038186803b1580156112ac57600080fd5b505af415801561107c573d6000803e3d6000fd5b6112c9826111fa565b6112d281612e4b565b6112dc8383612eb7565b50505050565b6060610e7b6107c683612dd5565b6001600160a01b03811633146113195760405163334bd91960e11b815260040160405180910390fd5b6111dd8282612ef9565b600061132d612a38565b61133d8a8a8a8a8a8a8a8a61181e565b60405161190160f01b6020820152602281019290925260428201526062015b60405160208183030381529060405280519060200120905098975050505050505050565b6000805160206156af83398151915261139881612e4b565b60005b82518110156111dd5760006113bb84838151811061116557611165615171565b90506113c681612f32565b8382815181106113d8576113d8615171565b60200260200101516040516113ed91906151f4565b604051908190038120907f0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e90600090a25060010161139b565b6000805160206156af83398151915261143e81612e4b565b61145061144a84612dd5565b83612f70565b8260405161145e91906151f4565b604051908190038120838252907fdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41906020015b60405180910390a2505050565b606060006114aa612da4565b9050610fc581612dc8565b6000805160206156af8339815191526114cd81612e4b565b6114e16114d986612dd5565b858486612f99565b846040516114ef91906151f4565b604051908190038120907f763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae937527990600090a25050505050565b6000806001600160a01b038b16301461155d578a604051637b7d696f60e01b81526004016115549190614cd1565b60405180910390fd5b6115656122d1565b6001600160a01b03166312b357b58b6040518263ffffffff1660e01b81526004016115909190614cd1565b602060405180830381865afa1580156115ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d19190615210565b6115f05789604051633f2c64af60e01b81526004016115549190614cd1565b84421061161357604051636453b2e560e01b815260048101869052602401611554565b600061161d6122d1565b6001600160a01b031663d79e59266040518163ffffffff1660e01b8152600401602060405180830381865afa15801561165a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061167e919061522d565b9050806001600160a01b0316856001600160a01b0316146116b65784816040516304dac3ed60e11b815260040161155492919061524a565b6116c78c8c8c8c8c8c8c8c8c613031565b92506116d2836125e5565b6116f1578260405163248d101d60e11b81526004016115549190614cd1565b60006116fb61305f565b90506000816003016000866001600160a01b03166001600160a01b0316815260200190815260200160002060008d6001600160a01b03166001600160a01b031681526020019081526020016000206000886001600160a01b03166001600160a01b03168152602001908152602001600020905080600101548a101561179a5760018101546040516307c83fcf60e41b8152611554918c91600401614d11565b80548b116117bf578054604051632256490160e11b8152611554918d91600401614d11565b60018101546117ce908b61527a565b9350505050995099975050505050505050565b6117e9613083565b6117f282613113565b6117fc8282613256565b5050565b60008061180b61330a565b600101546001600160a01b031692915050565b6040805160008051602061574f83398151915260208201526001600160a01b03808b16928201929092528189166060820152818816608082015260a0810187905260c0810186905260e0810185905261010081018490529082166101208201526000906101400161135c565b6000805160206156cf8339815191526118a281612e4b565b6001600160a01b0383166118c957604051633a954ecd60e21b815260040160405180910390fd5b6118e160008051602061564f83398151915284612eb7565b506118fa60008051602061572f83398151915284612eb7565b5061191360008051602061562f83398151915284612eb7565b506040516001600160a01b038416907fdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a87399490600090a26111dd6001600160a01b0384168361332e565b60006119656133c7565b5060008051602061568f83398151915290565b611980613410565b60008051602061562f83398151915261199881612e4b565b6119a182613446565b506119aa613573565b50565b606060006119b9612da4565b90506119c58382613584565b806002016000848152602001908152602001600020600201805480602002602001604051908101604052809291908181526020016000905b82821015611aa9578382906000526020600020018054611a1c9061528d565b80601f0160208091040260200160405190810160405280929190818152602001828054611a489061528d565b8015611a955780601f10611a6a57610100808354040283529160200191611a95565b820191906000526020600020905b815481529060010190602001808311611a7857829003601f168201915b5050505050815260200190600101906119fd565b50505050915050919050565b6000806000611ac26135ae565b90508060020154816003015492509250509091565b6000805160206156af833981519152611aef81612e4b565b6117fc826135d2565b6000805160206156cf833981519152611b1081612e4b565b6111dd8383613648565b6000806000806000611b2a61305f565b6001600160a01b0398891660009081526003918201602090815260408083209a8c16835299815289822098909a1681529690985250505092909120805460018201546002830154929095015490959193509150565b60008051602061572f833981519152611b9781612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63fd13a43e611bb9611800565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018890526044810187905261ffff80871660648301528516608482015260a40160006040518083038186803b158015611c1857600080fd5b505af4158015611c2c573d6000803e3d6000fd5b505050505050505050565b60008051602061572f833981519152611c4f81612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63793dddac611254611800565b6000805160206156af833981519152611c8981612e4b565b611c9b611c9584612dd5565b836136a2565b82604051611ca991906151f4565b60405180910390207f498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf836040516114919190614f31565b6000805160206156af833981519152611cf881612e4b565b6117fc826136e6565b6060806000611d0e61149e565b9050600081516001600160401b03811115611d2b57611d2b61473c565b604051908082528060200260200182016040528015611d5e57816020015b6060815260200190600190039081611d495790505b509050600082516001600160401b03811115611d7c57611d7c61473c565b604051908082528060200260200182016040528015611db557816020015b611da26144fd565b815260200190600190039081611d9a5790505b50905060005b8351811015611ebb57611dcc6122d1565b6001600160a01b0316635a81a626858381518110611dec57611dec615171565b60200260200101516040518263ffffffff1660e01b8152600401611e1291815260200190565b600060405180830381865afa158015611e2f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611e579190810190615187565b838281518110611e6957611e69615171565b6020026020010181905250611e96848281518110611e8957611e89615171565b60200260200101516124bc565b828281518110611ea857611ea8615171565b6020908102919091010152600101611dbb565b509094909350915050565b60606000611ed2612da4565b9050611ee1600682018461375c565b611f00578260405163ba650b5f60e01b81526004016115549190614cd1565b6001600160a01b038316600090815260088201602052604090208054611f259061528d565b80601f0160208091040260200160405190810160405280929190818152602001828054611f519061528d565b8015611f9e5780601f10611f7357610100808354040283529160200191611f9e565b820191906000526020600020905b815481529060010190602001808311611f8157829003601f168201915b5050505050915050919050565b6000805160206156af833981519152611fc381612e4b565b611fd5611fcf84612dd5565b83613771565b82604051611fe391906151f4565b60405180910390207fba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264836040516114919190614f31565b600080612025612da4565b90506120318382613584565b600092835260020160205250604090206001015460ff1690565b6000806120566138a7565b600085815260208290526040902090915061207190846138cb565b949350505050565b600080612084612e93565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b6000805160206156ef8339815191526120c781612e4b565b6001600160a01b0383166120ee57604051633a954ecd60e21b815260040160405180910390fd5b6112dc6001600160a01b03851684846138d7565b6000805160206156af83398151915261211a81612e4b565b6117fc8261392f565b6060600061212f6138a7565b600084815260208290526040902090915061214990612dc8565b9392505050565b60008061215b61305f565b6001600160a01b03909316600090815260049093016020525050604090205490565b6000805160206156af83398151915261219581612e4b565b6121a76121a184612dd5565b83613986565b826040516121b591906151f4565b6040519081900381208315158252907f23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab90602001611491565b6000610e7b6109cf83612dd5565b6000805160206156af83398151915261221481612e4b565b600061221e611d01565b50905060005b81518110156111dd5761224a61224583838151811061116557611165615171565b6139c0565b81818151811061225c5761225c615171565b602002602001015160405161227191906151f4565b6040519081900381209060008051602061560f83398151915290600090a2600101612224565b60008051602061572f8339815191526122af81612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63348e06dd6110c0611800565b6000806122dc61330a565b546001600160a01b031692915050565b6000805160206156cf83398151915261230481612e4b565b61231c60008051602061564f83398151915283612ef9565b5061233560008051602061572f83398151915283612ef9565b5061234e60008051602061562f83398151915283612ef9565b506040516001600160a01b038316907fd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc6291390600090a25050565b6000806123926138a7565b600084815260208290526040902090915061214990613a29565b6000805160206156af8339815191526123c481612e4b565b6111dd8383613a33565b6123d6613410565b60008051602061572f8339815191526123ee81612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__637adf63b7612410611800565b6040516001600160e01b031960e084901b1681526001600160a01b0391821660048201526024810188905260448101879052908516606482015260840160006040518083038186803b15801561246557600080fd5b505af4158015612479573d6000803e3d6000fd5b50505050506111dd613573565b60606000612492612da4565b9050610fc560048201612dc8565b6124a9826111fa565b6124b281612e4b565b6112dc8383612ef9565b6124c46144fd565b60006124ce612da4565b90506124da8382613584565b6000838152600280830160209081526040808420815160608101835281548152600182015460ff161515818501529381018054835181860281018601855281815295969295938701949192909184015b828210156125d65783829060005260206000200180546125499061528d565b80601f01602080910402602001604051908101604052809291908181526020018280546125759061528d565b80156125c25780601f10612597576101008083540402835291602001916125c2565b820191906000526020600020905b8154815290600101906020018083116125a557829003601f168201915b50505050508152602001906001019061252a565b50505091525090949350505050565b6000610e7b60008051602061564f83398151915283612079565b60008051602061572f83398151915261261781612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63e4c22569612639611800565b8a8a8a8a8a8a8a6040518963ffffffff1660e01b81526004016126639897969594939291906152c7565b60006040518083038186803b15801561267b57600080fd5b505af415801561268f573d6000803e3d6000fd5b505050505050505050505050565b60008051602061572f8339815191526126b581612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63c7bffa966126d7611800565b846040518363ffffffff1660e01b81526004016126f5929190614c69565b60006040518083038186803b15801561270d57600080fd5b505af4158015612721573d6000803e3d6000fd5b505050505050565b6000805160206156af83398151915261274181612e4b565b6117fc82613acf565b60008061275561305f565b6001015492915050565b6060600061276b612da4565b9050610fc581600601612dc8565b600080612784612da4565b90506127908382613584565b600092835260020160205250604090205490565b6000805160206156af8339815191526127bc81612e4b565b6127ce6127c884612dd5565b83613b68565b826040516127dc91906151f4565b604051908190038120907fd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c53137190600090a2505050565b612819613410565b60008061282d8b8b8b8b8b8b8b8b8b611526565b91509150612840828a8a8a8a8a8a613ba6565b600061271061284d6122d1565b6001600160a01b0316633c5559386040518163ffffffff1660e01b8152600401602060405180830381865afa15801561288a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ae9190615325565b6128b8908461533e565b6128c29190615355565b905060006128d0828461527a565b90508b6001600160a01b03168d6001600160a01b03161461292b5760006128f561305f565b6001600160a01b0388166000908152600482016020526040812080549293508692909190612924908490615377565b9091555050505b6129a86129366122d1565b6001600160a01b0316630470d3ac6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612973573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612997919061522d565b6001600160a01b03881690846138d7565b6129bc6001600160a01b0387168d836138d7565b604080516001600160a01b0386811682528d811660208301529181018c9052606081018b90526080810183905260a0810184905287821660c0820152818e16918f16907fa7708e82cb201b0c0dc3d520642d0e0eb290d001b6acec29d01aeeb6af7dab209060e00160405180910390a350505050611c2c613573565b600080612a4361305f565b6002015492915050565b6000806000612a5a6135ae565b6001600160a01b03909416600090815260208581526040808320546001909701909152902054939492505050565b612a90613410565b6000805160206156ef833981519152612aa881612e4b565b6001600160a01b038316612acf57604051633a954ecd60e21b815260040160405180910390fd5b612ae26001600160a01b0384168361332e565b826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436483604051612b1d91815260200190565b60405180910390a2506117fc613573565b6000805160206156af833981519152612b4681612e4b565b612b5261224583612dd5565b81604051612b6091906151f4565b6040519081900381209060008051602061560f83398151915290600090a25050565b60008051602061572f833981519152612b9a81612e4b565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63b54e72d8611bb9611800565b6000612bc6613c1d565b805490915060ff600160401b82041615906001600160401b0316600081158015612bed5750825b90506000826001600160401b03166001148015612c095750303b155b905081158015612c17575080155b15612c355760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315612c5e57845460ff60401b1916600160401b1785555b612c66613c41565b612c6e613c41565b612c76613c49565b612c81600088612eb7565b50612c9a6000805160206156af83398151915288612eb7565b50612cb36000805160206156cf83398151915288612eb7565b50612ccc60008051602061566f83398151915287612eb7565b506000612cd761330a565b80546001600160a01b03808d166001600160a01b0319928316178355600183018054918d16919092161790559050678ac7230489e8000062015180612d1c8282613ce3565b5050508315611c2c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050505050565b60006001600160e01b03198216637965db0b60e01b1480610e7b57506301ffc9a760e01b6001600160e01b0319831614610e7b565b7ff2856e5e1b7689dcde1bb551fd115c3cad8d243ea609d47a46b4d22ee58d300090565b6060600061214983613d06565b6000612ddf6122d1565b6001600160a01b031663352af39a836040518263ffffffff1660e01b8152600401612e0a9190614f31565b602060405180830381865afa158015612e27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7b9190615325565b6119aa8133613d62565b6000612e5f612da4565b90506000612e706009830184613d8d565b9050806111dd57604051631a1e056960e01b815260048101849052602401611554565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b600080612ec26138a7565b90506000612ed08585613d99565b90508015612071576000858152602083905260409020612ef09085613e3a565b50949350505050565b600080612f046138a7565b90506000612f128585613e4f565b90508015612071576000858152602083905260409020612ef09085613ec7565b6000612f3c612da4565b90506000612f4d6009830184613edc565b9050806111dd57604051637eae59f160e11b815260048101849052602401611554565b6000612f7a612da4565b9050612f868382613584565b6000928352600201602052604090912055565b6000612fa3612da4565b90506000612fb18287613d8d565b905080612fd3576040516221e3bb60e31b815260048101879052602401611554565b60408051606081018252868152841515602080830191825282840188815260008b81526002888101845295902084518155925160018401805460ff191691151591909117905551805193949293611c2c938501929190910190614520565b6000806130448b8b8b8b8b8b8b8b611323565b90506130508184613ee8565b9b9a5050505050505050505050565b7f175f7e400d42af44d9ebd24e9efee8a2c4ed78ddf46a83e51a493ae382c8760090565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806130f357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166130e7613f12565b6001600160a01b031614155b156131115760405163703e46dd60e11b815260040160405180910390fd5b565b60008051602061566f83398151915261312b81612e4b565b60006131356122d1565b6001600160a01b0316639d825bc56040518163ffffffff1660e01b8152600401602060405180830381865afa158015613172573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613196919061522d565b905060006131a2613f12565b9050836001600160a01b0316816001600160a01b0316036131da5780846040516382afabc160e01b815260040161155492919061524a565b816001600160a01b0316846001600160a01b03161461321057818460405163699a021f60e11b815260040161155492919061524a565b836001600160a01b0316816001600160a01b03167fa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db460405160405180910390a350505050565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156132b0575060408051601f3d908101601f191682019092526132ad91810190615325565b60015b6132cf5781604051634c9c8ce360e01b81526004016115549190614cd1565b60008051602061568f833981519152811461330057604051632a87526960e21b815260048101829052602401611554565b6111dd8383613f2e565b7f0c7b73796c7cc89b9f849b9056a93200eba741881e57a1b03b9bedb2c0e0710090565b8047101561335357478160405163cf47918160e01b8152600401611554929190614d11565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146133a0576040519150601f19603f3d011682016040523d82523d6000602084013e6133a5565b606091505b50509050806111dd5760405163d6bda27560e01b815260040160405180910390fd5b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146131115760405163703e46dd60e11b815260040160405180910390fd5b600061341a613f84565b80549091506001190161344057604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60006134506135ae565b9050806002015482111561347f57806002015482604051631728bc5b60e31b8152600401611554929190614d11565b600381015433600090815260208390526040902054429161349f91615377565b8111156134c75733600090815260018301602090815260408083208390559084905290208190555b60028201543360009081526001840160205260409020546134e9908590615377565b11156135105781600201548360405163d54b188760e01b8152600401611554929190614d11565b33600090815260018301602052604081208054859290613531908490615377565b909155506135419050338461332e565b60405183815233907fb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c290602001611491565b600061357d613f84565b6001905550565b61358e8183613fa8565b6117fc57604051631e96f6ed60e21b815260048101839052602401611554565b7f99a652063088b6badaeb0c7f680676baf720654b4f86f50167944489af637d0090565b60006135dc612da4565b905060006135ed6004830184613e3a565b90508061360f5782604051632872fbf960e11b81526004016115549190614cd1565b6040516001600160a01b038416907fa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f90600090a2505050565b60006136526135ae565b60028101849055600381018390556040519091507f8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e906136959085908590614d11565b60405180910390a1505050565b60006136ac612da4565b90506136b88382613584565b6000838152600280830160209081526040832090910180546001810182559083529120016112dc83826153e7565b60006136f0612da4565b905060006137016004830184613ec7565b9050806137235782604051631532e67160e21b81526004016115549190614cd1565b6040516001600160a01b038416907f85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af290600090a2505050565b6000612149836001600160a01b038416613fb0565b600061377b612da4565b90506137878382613584565b600083815260028083016020526040822001905b81548110156138a057836040516020016137b591906151f4565b604051602081830303815290604052805190602001208282815481106137dd576137dd615171565b906000526020600020016040516020016137f791906154a0565b604051602081830303815290604052805190602001200361389857815482906138229060019061527a565b8154811061383257613832615171565b9060005260206000200182828154811061384e5761384e615171565b9060005260206000200190816138649190615516565b5081805480613875576138756155e2565b6001900381819060005260206000200160006138919190614576565b90556138a0565b60010161379b565b5050505050565b7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200090565b60006121498383613fc8565b6111dd83846001600160a01b031663a9059cbb85856040516024016138fd929190614c69565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613ff2565b6000613939612da4565b60038101805460ff19168415159081179091556040519081529091507fe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e39060200160405180910390a15050565b6000613990612da4565b905061399c8382613584565b60009283526002016020526040909120600101805460ff1916911515919091179055565b60006139ca612da4565b905060006139d88284613edc565b9050806139fb57604051631e96f6ed60e21b815260048101849052602401611554565b600083815260028084016020526040822082815560018101805460ff1916905591906138a0908301826145b0565b6000610e7b825490565b6000613a3d612da4565b90506000613a4e6006830185613e3a565b905080613a705783604051631a6107e360e31b81526004016115549190614cd1565b6001600160a01b03841660009081526008830160205260409020613a9484826153e7565b506040516001600160a01b038516907f928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe8290600090a250505050565b6000613ad9612da4565b90506000613aea6006830184613ec7565b905080613b0c578260405163ba650b5f60e01b81526004016115549190614cd1565b6001600160a01b03831660009081526008830160205260408120613b2f91614576565b6040516001600160a01b038416907fc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf90600090a2505050565b6000613b72612da4565b9050613b7e8382613584565b600083815260028083016020908152604090922084516112dc93919092019190850190614520565b6000613bb061305f565b604080516080810182529788526020808901978852888201968752606089019586526001600160a01b039a8b166000908152600393840182528281209a8c168152998152818a2094909a168952929098525090942092518355905160018301555160028201559051910155565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b61311161405a565b613c5161405a565b6000613c5b61305f565b604051909150613cc39060008051602061570f833981519152907f964f47ba1d5dc0d1184a60039dba40abb2d0eee53398392e97308cb5f4a12f94907fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6904690602001614d87565b60408051601f198184030181529190528051602090910120600290910155565b613ceb61405a565b6000613cf56135ae565b600281019390935550600390910155565b606081600001805480602002602001604051908101604052809291908181526020018280548015613d5657602002820191906000526020600020905b815481526020019060010190808311613d42575b50505050509050919050565b613d6c8282612079565b6117fc57808260405163e2517d3f60e01b8152600401611554929190614c69565b6000612149838361407f565b600080613da4612e93565b9050613db08484612079565b613e30576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055613de63390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610e7b565b6000915050610e7b565b6000612149836001600160a01b03841661407f565b600080613e5a612e93565b9050613e668484612079565b15613e30576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610e7b565b6000612149836001600160a01b0384166140c9565b600061214983836140c9565b600080600080613ef886866141b2565b925092509250613f0882826141ff565b5090949350505050565b60008051602061568f833981519152546001600160a01b031690565b613f37826142b8565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115613f7c576111dd8282614314565b6117fc61438a565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b600061214983835b60009081526001919091016020526040902054151590565b6000826000018281548110613fdf57613fdf615171565b9060005260206000200154905092915050565b600080602060008451602086016000885af180614015576040513d6000823e3d81fd5b50506000513d9150811561402d57806001141561403a565b6001600160a01b0384163b155b156112dc5783604051635274afe760e01b81526004016115549190614cd1565b6140626143a9565b61311157604051631afcd79f60e31b815260040160405180910390fd5b600061408b8383613fb0565b6140c157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610e7b565b506000610e7b565b60008181526001830160205260408120548015613e305760006140ed60018361527a565b85549091506000906141019060019061527a565b905080821461416657600086600001828154811061412157614121615171565b906000526020600020015490508087600001848154811061414457614144615171565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614177576141776155e2565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610e7b565b600080600083516041036141ec5760208401516040850151606086015160001a6141de888285856143c3565b9550955095505050506141f8565b50508151600091506002905b9250925092565b6000826003811115614213576142136155f8565b0361421c575050565b6001826003811115614230576142306155f8565b0361424e5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115614262576142626155f8565b036142835760405163fce698f760e01b815260048101829052602401611554565b6003826003811115614297576142976155f8565b036117fc576040516335e2f38360e21b815260048101829052602401611554565b806001600160a01b03163b6000036142e55780604051634c9c8ce360e01b81526004016115549190614cd1565b60008051602061568f83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161433191906151f4565b600060405180830381855af49150503d806000811461436c576040519150601f19603f3d011682016040523d82523d6000602084013e614371565b606091505b5091509150614381858383614488565b95945050505050565b34156131115760405163b398979f60e01b815260040160405180910390fd5b60006143b3613c1d565b54600160401b900460ff16919050565b600080806fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411156143f4575060009150600390508261447e565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614448573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166144745750600092506001915082905061447e565b9250600091508190505b9450945094915050565b60608261449d57614498826144d4565b612149565b81511580156144b457506001600160a01b0384163b155b15610fa45783604051639996b31560e01b81526004016115549190614cd1565b8051156144e45780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b604051806060016040528060008152602001600015158152602001606081525090565b828054828255906000526020600020908101928215614566579160200282015b82811115614566578251829061455690826153e7565b5091602001919060010190614540565b506145729291506145ca565b5090565b5080546145829061528d565b6000825580601f10614592575050565b601f0160209004906000526020600020908101906119aa91906145e7565b50805460008255906000526020600020908101906119aa91905b808211156145725760006145de8282614576565b506001016145ca565b5b8082111561457257600081556001016145e8565b60006020828403121561460e57600080fd5b81356001600160e01b03198116811461214957600080fd5b60005b83811015614641578181015183820152602001614629565b50506000910152565b60008151808452614662816020860160208601614626565b601f01601f19169290920160200192915050565b60008282518085526020808601955060208260051b8401016020860160005b848110156146c357601f198684030189526146b183835161464a565b98840198925090830190600101614695565b5090979650505050505050565b6020815260006121496020830184614676565b6020808252825182820181905260009190848201906040850190845b8181101561471b578351835292840192918401916001016146ff565b50909695505050505050565b6001600160a01b03811681146119aa57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561477a5761477a61473c565b604052919050565b60006001600160401b0382111561479b5761479b61473c565b50601f01601f191660200190565b600082601f8301126147ba57600080fd5b81356147cd6147c882614782565b614752565b8181528460208386010111156147e257600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806080858703121561481557600080fd5b843561482081614727565b9350602085013561483081614727565b92506040850135915060608501356001600160401b0381111561485257600080fd5b61485e878288016147a9565b91505092959194509250565b60006020828403121561487c57600080fd5b81356001600160401b0381111561489257600080fd5b612071848285016147a9565b6000806000606084860312156148b357600080fd5b83356148be81614727565b925060208401356148ce81614727565b929592945050506040919091013590565b600080604083850312156148f257600080fd5b50508035926020909101359150565b600082601f83011261491257600080fd5b813560206001600160401b038083111561492e5761492e61473c565b8260051b61493d838201614752565b938452858101830193838101908886111561495757600080fd5b84880192505b85831015614993578235848111156149755760008081fd5b6149838a87838c01016147a9565b835250918401919084019061495d565b98975050505050505050565b6000602082840312156149b157600080fd5b81356001600160401b038111156149c757600080fd5b61207184828501614901565b6000602082840312156149e557600080fd5b5035919050565b803561ffff811681146149fe57600080fd5b919050565b600080600060608486031215614a1857600080fd5b83359250614a28602085016149ec565b9150614a36604085016149ec565b90509250925092565b60008060408385031215614a5257600080fd5b823591506020830135614a6481614727565b809150509250929050565b600080600080600080600080610100898b031215614a8c57600080fd5b8835614a9781614727565b97506020890135614aa781614727565b96506040890135614ab781614727565b9550606089013594506080890135935060a0890135925060c0890135915060e0890135614ae381614727565b809150509295985092959890939650565b60008060408385031215614b0757600080fd5b82356001600160401b03811115614b1d57600080fd5b614b29858286016147a9565b95602094909401359450505050565b80151581146119aa57600080fd5b60008060008060808587031215614b5c57600080fd5b84356001600160401b0380821115614b7357600080fd5b614b7f888389016147a9565b95506020870135945060408701359150614b9882614b38565b90925060608601359080821115614bae57600080fd5b5061485e87828801614901565b60008060008060008060008060006101208a8c031215614bda57600080fd5b8935614be581614727565b985060208a0135614bf581614727565b975060408a0135614c0581614727565b965060608a0135955060808a0135945060a08a0135935060c08a0135925060e08a0135614c3181614727565b91506101008a01356001600160401b03811115614c4d57600080fd5b614c598c828d016147a9565b9150509295985092959850929598565b6001600160a01b03929092168252602082015260400190565b60008060408385031215614c9557600080fd5b8235614ca081614727565b915060208301356001600160401b03811115614cbb57600080fd5b614cc7858286016147a9565b9150509250929050565b6001600160a01b0391909116815260200190565b60008060408385031215614cf857600080fd5b8235614d0381614727565b946020939093013593505050565b918252602082015260400190565b600060208284031215614d3157600080fd5b813561214981614727565b600080600060608486031215614d5157600080fd5b8335614d5c81614727565b92506020840135614d6c81614727565b91506040840135614d7c81614727565b809150509250925092565b93845260208401929092526040830152606082015260800190565b60008060008060808587031215614db857600080fd5b8435935060208501359250614dcf604086016149ec565b9150614ddd606086016149ec565b905092959194509250565b60008060408385031215614dfb57600080fd5b82356001600160401b0380821115614e1257600080fd5b614e1e868387016147a9565b93506020850135915080821115614e3457600080fd5b50614cc7858286016147a9565b600060608301825184526020808401511515602086015260408401516060604087015282815180855260808801915060808160051b890101945060208301925060005b81811015614eb257607f19898703018352614ea086855161464a565b95509284019291840191600101614e84565b5093979650505050505050565b604081526000614ed26040830185614676565b6020838203818501528185518084528284019150828160051b85010183880160005b83811015614f2257601f19878403018552614f10838351614e41565b94860194925090850190600101614ef4565b50909998505050505050505050565b602081526000612149602083018461464a565b600060208284031215614f5657600080fd5b813561214981614b38565b6020808252825182820181905260009190848201906040850190845b8181101561471b5783516001600160a01b031683529284019291840191600101614f7d565b60008060408385031215614fb557600080fd5b82356001600160401b03811115614fcb57600080fd5b614fd7858286016147a9565b9250506020830135614a6481614b38565b600080600060608486031215614ffd57600080fd5b83359250602084013591506040840135614d7c81614727565b6020815260006121496020830184614e41565b600080600080600080600060e0888a03121561504457600080fd5b873561504f81614727565b965060208801356001600160401b0381111561506a57600080fd5b6150768a828b016147a9565b9650506040880135945060608801359350608088013561509581614727565b925060a0880135915060c08801356150ac81614b38565b8091505092959891949750929550565b600080604083850312156150cf57600080fd5b82356001600160401b03808211156150e657600080fd5b6150f2868387016147a9565b9350602085013591508082111561510857600080fd5b50614cc785828601614901565b6000806000806080858703121561512b57600080fd5b843561513681614727565b9350602085013561514681614727565b9250604085013561515681614727565b9150606085013561516681614727565b939692955090935050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561519957600080fd5b81516001600160401b038111156151af57600080fd5b8201601f810184136151c057600080fd5b80516151ce6147c882614782565b8181528560208385010111156151e357600080fd5b614381826020830160208601614626565b60008251615206818460208701614626565b9190910192915050565b60006020828403121561522257600080fd5b815161214981614b38565b60006020828403121561523f57600080fd5b815161214981614727565b6001600160a01b0392831681529116602082015260400190565b634e487b7160e01b600052601160045260246000fd5b81810381811115610e7b57610e7b615264565b600181811c908216806152a157607f821691505b6020821081036152c157634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b0389811682528881166020830152610100604083018190526000916152f58483018b61464a565b6060850199909952608084019790975250509290931660a083015260c082015290151560e0909101529392505050565b60006020828403121561533757600080fd5b5051919050565b8082028115828204841417610e7b57610e7b615264565b60008261537257634e487b7160e01b600052601260045260246000fd5b500490565b80820180821115610e7b57610e7b615264565b601f8211156111dd576000816000526020600020601f850160051c810160208610156153b35750805b601f850160051c820191505b81811015612721578281556001016153bf565b600019600383901b1c191660019190911b1790565b81516001600160401b038111156154005761540061473c565b6154148161540e845461528d565b8461538a565b602080601f83116001811461544357600084156154315750858301515b61543b85826153d2565b865550612721565b600085815260208120601f198616915b8281101561547257888601518255948401946001909101908401615453565b50858210156154905787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008083546154ae8161528d565b600182811680156154c657600181146154db5761550a565b60ff198416875282151583028701945061550a565b8760005260208060002060005b858110156155015781548a8201529084019082016154e8565b50505082870194505b50929695505050505050565b818103615521575050565b61552b825461528d565b6001600160401b038111156155425761554261473c565b6155508161540e845461528d565b6000601f82116001811461557e576000831561556c5750848201545b61557684826153d2565b8555506138a0565b600085815260209020601f19841690600086815260209020845b838110156155b85782860154825560019586019590910190602001615598565b50858310156154905793015460001960f8600387901b161c19169092555050600190811b01905550565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fdfe52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813e562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c954c1d97560aef1d9626af2ccff8a6ac6e245c44c307e795e9c16f7211c611a315189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9a95e87c5af084bf5db8491c3a6515da9dd6da39b24b0eb0af08d7b9cd808d91c6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4c2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e3acdf00ba9ef08b5f2c22768276611b9af078bf6c24fa36b34ec5e9f2eb061fa47a14584cc614c4358a01f9a3731417edd2a8d4528cf486fc8b0489059a33214a26469706673582212204c730cfcd098b3fe777d1a9002b2b756deb7cf7145188595a96ba0b37bb2775064736f6c63430008180033",
}

// CmaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use CmaccountMetaData.ABI instead.
var CmaccountABI = CmaccountMetaData.ABI

// CmaccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CmaccountMetaData.Bin instead.
var CmaccountBin = CmaccountMetaData.Bin

// DeployCmaccount deploys a new Ethereum contract, binding an instance of Cmaccount to it.
func DeployCmaccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cmaccount, error) {
	parsed, err := CmaccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CmaccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cmaccount{CmaccountCaller: CmaccountCaller{contract: contract}, CmaccountTransactor: CmaccountTransactor{contract: contract}, CmaccountFilterer: CmaccountFilterer{contract: contract}}, nil
}

// Cmaccount is an auto generated Go binding around an Ethereum contract.
type Cmaccount struct {
	CmaccountCaller     // Read-only binding to the contract
	CmaccountTransactor // Write-only binding to the contract
	CmaccountFilterer   // Log filterer for contract events
}

// CmaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type CmaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CmaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CmaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CmaccountSession struct {
	Contract     *Cmaccount        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CmaccountCallerSession struct {
	Contract *CmaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CmaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CmaccountTransactorSession struct {
	Contract     *CmaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CmaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type CmaccountRaw struct {
	Contract *Cmaccount // Generic contract binding to access the raw methods on
}

// CmaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CmaccountCallerRaw struct {
	Contract *CmaccountCaller // Generic read-only contract binding to access the raw methods on
}

// CmaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CmaccountTransactorRaw struct {
	Contract *CmaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCmaccount creates a new instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccount(address common.Address, backend bind.ContractBackend) (*Cmaccount, error) {
	contract, err := bindCmaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cmaccount{CmaccountCaller: CmaccountCaller{contract: contract}, CmaccountTransactor: CmaccountTransactor{contract: contract}, CmaccountFilterer: CmaccountFilterer{contract: contract}}, nil
}

// NewCmaccountCaller creates a new read-only instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccountCaller(address common.Address, caller bind.ContractCaller) (*CmaccountCaller, error) {
	contract, err := bindCmaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CmaccountCaller{contract: contract}, nil
}

// NewCmaccountTransactor creates a new write-only instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*CmaccountTransactor, error) {
	contract, err := bindCmaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CmaccountTransactor{contract: contract}, nil
}

// NewCmaccountFilterer creates a new log filterer instance of Cmaccount, bound to a specific deployed contract.
func NewCmaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*CmaccountFilterer, error) {
	contract, err := bindCmaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CmaccountFilterer{contract: contract}, nil
}

// bindCmaccount binds a generic wrapper to an already deployed contract.
func bindCmaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CmaccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cmaccount *CmaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cmaccount.Contract.CmaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cmaccount *CmaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.Contract.CmaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cmaccount *CmaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cmaccount.Contract.CmaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cmaccount *CmaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cmaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cmaccount *CmaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cmaccount *CmaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cmaccount.Contract.contract.Transact(opts, method, params...)
}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) BOOKINGOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "BOOKING_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) BOOKINGOPERATORROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOOKINGOPERATORROLE(&_Cmaccount.CallOpts)
}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) BOOKINGOPERATORROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOOKINGOPERATORROLE(&_Cmaccount.CallOpts)
}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) BOTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "BOT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) BOTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOTADMINROLE(&_Cmaccount.CallOpts)
}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) BOTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.BOTADMINROLE(&_Cmaccount.CallOpts)
}

// CHEQUEOPERATORROLE is a free data retrieval call binding the contract method 0x63e86cc8.
//
// Solidity: function CHEQUE_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) CHEQUEOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "CHEQUE_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHEQUEOPERATORROLE is a free data retrieval call binding the contract method 0x63e86cc8.
//
// Solidity: function CHEQUE_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) CHEQUEOPERATORROLE() ([32]byte, error) {
	return _Cmaccount.Contract.CHEQUEOPERATORROLE(&_Cmaccount.CallOpts)
}

// CHEQUEOPERATORROLE is a free data retrieval call binding the contract method 0x63e86cc8.
//
// Solidity: function CHEQUE_OPERATOR_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) CHEQUEOPERATORROLE() ([32]byte, error) {
	return _Cmaccount.Contract.CHEQUEOPERATORROLE(&_Cmaccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.DEFAULTADMINROLE(&_Cmaccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.DEFAULTADMINROLE(&_Cmaccount.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cmaccount *CmaccountSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cmaccount.Contract.DOMAINTYPEHASH(&_Cmaccount.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cmaccount.Contract.DOMAINTYPEHASH(&_Cmaccount.CallOpts)
}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) GASWITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "GAS_WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) GASWITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.GASWITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) GASWITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.GASWITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// MESSENGERCHEQUETYPEHASH is a free data retrieval call binding the contract method 0x0ede80d6.
//
// Solidity: function MESSENGER_CHEQUE_TYPEHASH() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) MESSENGERCHEQUETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "MESSENGER_CHEQUE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSENGERCHEQUETYPEHASH is a free data retrieval call binding the contract method 0x0ede80d6.
//
// Solidity: function MESSENGER_CHEQUE_TYPEHASH() view returns(bytes32)
func (_Cmaccount *CmaccountSession) MESSENGERCHEQUETYPEHASH() ([32]byte, error) {
	return _Cmaccount.Contract.MESSENGERCHEQUETYPEHASH(&_Cmaccount.CallOpts)
}

// MESSENGERCHEQUETYPEHASH is a free data retrieval call binding the contract method 0x0ede80d6.
//
// Solidity: function MESSENGER_CHEQUE_TYPEHASH() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) MESSENGERCHEQUETYPEHASH() ([32]byte, error) {
	return _Cmaccount.Contract.MESSENGERCHEQUETYPEHASH(&_Cmaccount.CallOpts)
}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) SERVICEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "SERVICE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) SERVICEADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.SERVICEADMINROLE(&_Cmaccount.CallOpts)
}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) SERVICEADMINROLE() ([32]byte, error) {
	return _Cmaccount.Contract.SERVICEADMINROLE(&_Cmaccount.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) UPGRADERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.UPGRADERROLE(&_Cmaccount.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.UPGRADERROLE(&_Cmaccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccount *CmaccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccount *CmaccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Cmaccount.Contract.UPGRADEINTERFACEVERSION(&_Cmaccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Cmaccount *CmaccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Cmaccount.Contract.UPGRADEINTERFACEVERSION(&_Cmaccount.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountSession) WITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.WITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _Cmaccount.Contract.WITHDRAWERROLE(&_Cmaccount.CallOpts)
}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCaller) GetAllServiceHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getAllServiceHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountSession) GetAllServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetAllServiceHashes(&_Cmaccount.CallOpts)
}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCallerSession) GetAllServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetAllServiceHashes(&_Cmaccount.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccount *CmaccountCaller) GetBookingTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getBookingTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccount *CmaccountSession) GetBookingTokenAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetBookingTokenAddress(&_Cmaccount.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Cmaccount *CmaccountCallerSession) GetBookingTokenAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetBookingTokenAddress(&_Cmaccount.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Cmaccount *CmaccountSession) GetDomainSeparator() ([32]byte, error) {
	return _Cmaccount.Contract.GetDomainSeparator(&_Cmaccount.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Cmaccount.Contract.GetDomainSeparator(&_Cmaccount.CallOpts)
}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Cmaccount *CmaccountCaller) GetGasMoneyWithdrawal(opts *bind.CallOpts) (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getGasMoneyWithdrawal")

	outstruct := new(struct {
		WithdrawalLimit  *big.Int
		WithdrawalPeriod *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawalLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawalPeriod = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Cmaccount *CmaccountSession) GetGasMoneyWithdrawal() (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawal(&_Cmaccount.CallOpts)
}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Cmaccount *CmaccountCallerSession) GetGasMoneyWithdrawal() (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawal(&_Cmaccount.CallOpts)
}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Cmaccount *CmaccountCaller) GetGasMoneyWithdrawalForAccount(opts *bind.CallOpts, account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getGasMoneyWithdrawalForAccount", account)

	outstruct := new(struct {
		PeriodStart     *big.Int
		WithdrawnAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PeriodStart = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawnAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Cmaccount *CmaccountSession) GetGasMoneyWithdrawalForAccount(account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawalForAccount(&_Cmaccount.CallOpts, account)
}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Cmaccount *CmaccountCallerSession) GetGasMoneyWithdrawalForAccount(account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	return _Cmaccount.Contract.GetGasMoneyWithdrawalForAccount(&_Cmaccount.CallOpts, account)
}

// GetLastCashIn is a free data retrieval call binding the contract method 0x70597753.
//
// Solidity: function getLastCashIn(address fromBot, address toBot, address paymentToken) view returns(uint256 lastCounter, uint256 lastAmount, uint256 lastCreatedAt, uint256 lastExpiresAt)
func (_Cmaccount *CmaccountCaller) GetLastCashIn(opts *bind.CallOpts, fromBot common.Address, toBot common.Address, paymentToken common.Address) (struct {
	LastCounter   *big.Int
	LastAmount    *big.Int
	LastCreatedAt *big.Int
	LastExpiresAt *big.Int
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getLastCashIn", fromBot, toBot, paymentToken)

	outstruct := new(struct {
		LastCounter   *big.Int
		LastAmount    *big.Int
		LastCreatedAt *big.Int
		LastExpiresAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastCounter = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastCreatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastExpiresAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLastCashIn is a free data retrieval call binding the contract method 0x70597753.
//
// Solidity: function getLastCashIn(address fromBot, address toBot, address paymentToken) view returns(uint256 lastCounter, uint256 lastAmount, uint256 lastCreatedAt, uint256 lastExpiresAt)
func (_Cmaccount *CmaccountSession) GetLastCashIn(fromBot common.Address, toBot common.Address, paymentToken common.Address) (struct {
	LastCounter   *big.Int
	LastAmount    *big.Int
	LastCreatedAt *big.Int
	LastExpiresAt *big.Int
}, error) {
	return _Cmaccount.Contract.GetLastCashIn(&_Cmaccount.CallOpts, fromBot, toBot, paymentToken)
}

// GetLastCashIn is a free data retrieval call binding the contract method 0x70597753.
//
// Solidity: function getLastCashIn(address fromBot, address toBot, address paymentToken) view returns(uint256 lastCounter, uint256 lastAmount, uint256 lastCreatedAt, uint256 lastExpiresAt)
func (_Cmaccount *CmaccountCallerSession) GetLastCashIn(fromBot common.Address, toBot common.Address, paymentToken common.Address) (struct {
	LastCounter   *big.Int
	LastAmount    *big.Int
	LastCreatedAt *big.Int
	LastExpiresAt *big.Int
}, error) {
	return _Cmaccount.Contract.GetLastCashIn(&_Cmaccount.CallOpts, fromBot, toBot, paymentToken)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Cmaccount *CmaccountCaller) GetManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Cmaccount *CmaccountSession) GetManagerAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetManagerAddress(&_Cmaccount.CallOpts)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Cmaccount *CmaccountCallerSession) GetManagerAddress() (common.Address, error) {
	return _Cmaccount.Contract.GetManagerAddress(&_Cmaccount.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Cmaccount *CmaccountCaller) GetPublicKey(opts *bind.CallOpts, pubKeyAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getPublicKey", pubKeyAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Cmaccount *CmaccountSession) GetPublicKey(pubKeyAddress common.Address) ([]byte, error) {
	return _Cmaccount.Contract.GetPublicKey(&_Cmaccount.CallOpts, pubKeyAddress)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Cmaccount *CmaccountCallerSession) GetPublicKey(pubKeyAddress common.Address) ([]byte, error) {
	return _Cmaccount.Contract.GetPublicKey(&_Cmaccount.CallOpts, pubKeyAddress)
}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Cmaccount *CmaccountCaller) GetPublicKeysAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getPublicKeysAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Cmaccount *CmaccountSession) GetPublicKeysAddresses() ([]common.Address, error) {
	return _Cmaccount.Contract.GetPublicKeysAddresses(&_Cmaccount.CallOpts)
}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Cmaccount *CmaccountCallerSession) GetPublicKeysAddresses() ([]common.Address, error) {
	return _Cmaccount.Contract.GetPublicKeysAddresses(&_Cmaccount.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccount *CmaccountCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccount *CmaccountSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cmaccount.Contract.GetRoleAdmin(&_Cmaccount.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cmaccount.Contract.GetRoleAdmin(&_Cmaccount.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccount *CmaccountCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccount *CmaccountSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Cmaccount.Contract.GetRoleMember(&_Cmaccount.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Cmaccount *CmaccountCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Cmaccount.Contract.GetRoleMember(&_Cmaccount.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccount *CmaccountCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccount *CmaccountSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetRoleMemberCount(&_Cmaccount.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Cmaccount *CmaccountCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetRoleMemberCount(&_Cmaccount.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Cmaccount *CmaccountCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Cmaccount *CmaccountSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Cmaccount.Contract.GetRoleMembers(&_Cmaccount.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Cmaccount *CmaccountCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Cmaccount.Contract.GetRoleMembers(&_Cmaccount.CallOpts, role)
}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((uint256,bool,string[]) service)
func (_Cmaccount *CmaccountCaller) GetService(opts *bind.CallOpts, serviceHash [32]byte) (PartnerConfigurationService, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getService", serviceHash)

	if err != nil {
		return *new(PartnerConfigurationService), err
	}

	out0 := *abi.ConvertType(out[0], new(PartnerConfigurationService)).(*PartnerConfigurationService)

	return out0, err

}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((uint256,bool,string[]) service)
func (_Cmaccount *CmaccountSession) GetService(serviceHash [32]byte) (PartnerConfigurationService, error) {
	return _Cmaccount.Contract.GetService(&_Cmaccount.CallOpts, serviceHash)
}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((uint256,bool,string[]) service)
func (_Cmaccount *CmaccountCallerSession) GetService(serviceHash [32]byte) (PartnerConfigurationService, error) {
	return _Cmaccount.Contract.GetService(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCaller) GetServiceCapabilities(opts *bind.CallOpts, serviceName string) ([]string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceCapabilities", serviceName)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Cmaccount *CmaccountSession) GetServiceCapabilities(serviceName string) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCallerSession) GetServiceCapabilities(serviceName string) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCaller) GetServiceCapabilities0(opts *bind.CallOpts, serviceHash [32]byte) ([]string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceCapabilities0", serviceHash)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Cmaccount *CmaccountSession) GetServiceCapabilities0(serviceHash [32]byte) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities0(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Cmaccount *CmaccountCallerSession) GetServiceCapabilities0(serviceHash [32]byte) ([]string, error) {
	return _Cmaccount.Contract.GetServiceCapabilities0(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceFee is a free data retrieval call binding the contract method 0x18274da4.
//
// Solidity: function getServiceFee(string serviceName) view returns(uint256 fee)
func (_Cmaccount *CmaccountCaller) GetServiceFee(opts *bind.CallOpts, serviceName string) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceFee", serviceName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetServiceFee is a free data retrieval call binding the contract method 0x18274da4.
//
// Solidity: function getServiceFee(string serviceName) view returns(uint256 fee)
func (_Cmaccount *CmaccountSession) GetServiceFee(serviceName string) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceFee is a free data retrieval call binding the contract method 0x18274da4.
//
// Solidity: function getServiceFee(string serviceName) view returns(uint256 fee)
func (_Cmaccount *CmaccountCallerSession) GetServiceFee(serviceName string) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceFee0 is a free data retrieval call binding the contract method 0xeb5ea273.
//
// Solidity: function getServiceFee(bytes32 serviceHash) view returns(uint256 fee)
func (_Cmaccount *CmaccountCaller) GetServiceFee0(opts *bind.CallOpts, serviceHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceFee0", serviceHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetServiceFee0 is a free data retrieval call binding the contract method 0xeb5ea273.
//
// Solidity: function getServiceFee(bytes32 serviceHash) view returns(uint256 fee)
func (_Cmaccount *CmaccountSession) GetServiceFee0(serviceHash [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee0(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceFee0 is a free data retrieval call binding the contract method 0xeb5ea273.
//
// Solidity: function getServiceFee(bytes32 serviceHash) view returns(uint256 fee)
func (_Cmaccount *CmaccountCallerSession) GetServiceFee0(serviceHash [32]byte) (*big.Int, error) {
	return _Cmaccount.Contract.GetServiceFee0(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCaller) GetServiceRestrictedRate(opts *bind.CallOpts, serviceHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceRestrictedRate", serviceHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountSession) GetServiceRestrictedRate(serviceHash [32]byte) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCallerSession) GetServiceRestrictedRate(serviceHash [32]byte) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate(&_Cmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCaller) GetServiceRestrictedRate0(opts *bind.CallOpts, serviceName string) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getServiceRestrictedRate0", serviceName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountSession) GetServiceRestrictedRate0(serviceName string) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate0(&_Cmaccount.CallOpts, serviceName)
}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Cmaccount *CmaccountCallerSession) GetServiceRestrictedRate0(serviceName string) (bool, error) {
	return _Cmaccount.Contract.GetServiceRestrictedRate0(&_Cmaccount.CallOpts, serviceName)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (uint256,bool,string[])[] services)
func (_Cmaccount *CmaccountCaller) GetSupportedServices(opts *bind.CallOpts) (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getSupportedServices")

	outstruct := new(struct {
		ServiceNames []string
		Services     []PartnerConfigurationService
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ServiceNames = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Services = *abi.ConvertType(out[1], new([]PartnerConfigurationService)).(*[]PartnerConfigurationService)

	return *outstruct, err

}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (uint256,bool,string[])[] services)
func (_Cmaccount *CmaccountSession) GetSupportedServices() (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	return _Cmaccount.Contract.GetSupportedServices(&_Cmaccount.CallOpts)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (uint256,bool,string[])[] services)
func (_Cmaccount *CmaccountCallerSession) GetSupportedServices() (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	return _Cmaccount.Contract.GetSupportedServices(&_Cmaccount.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Cmaccount *CmaccountCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Cmaccount *CmaccountSession) GetSupportedTokens() ([]common.Address, error) {
	return _Cmaccount.Contract.GetSupportedTokens(&_Cmaccount.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Cmaccount *CmaccountCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _Cmaccount.Contract.GetSupportedTokens(&_Cmaccount.CallOpts)
}

// GetTotalChequePayments is a free data retrieval call binding the contract method 0xe96cf7ad.
//
// Solidity: function getTotalChequePayments() view returns(uint256)
func (_Cmaccount *CmaccountCaller) GetTotalChequePayments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getTotalChequePayments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalChequePayments is a free data retrieval call binding the contract method 0xe96cf7ad.
//
// Solidity: function getTotalChequePayments() view returns(uint256)
func (_Cmaccount *CmaccountSession) GetTotalChequePayments() (*big.Int, error) {
	return _Cmaccount.Contract.GetTotalChequePayments(&_Cmaccount.CallOpts)
}

// GetTotalChequePayments is a free data retrieval call binding the contract method 0xe96cf7ad.
//
// Solidity: function getTotalChequePayments() view returns(uint256)
func (_Cmaccount *CmaccountCallerSession) GetTotalChequePayments() (*big.Int, error) {
	return _Cmaccount.Contract.GetTotalChequePayments(&_Cmaccount.CallOpts)
}

// GetTotalChequePaymentsPerToken is a free data retrieval call binding the contract method 0xa73ebdd8.
//
// Solidity: function getTotalChequePaymentsPerToken(address paymentToken) view returns(uint256)
func (_Cmaccount *CmaccountCaller) GetTotalChequePaymentsPerToken(opts *bind.CallOpts, paymentToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getTotalChequePaymentsPerToken", paymentToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalChequePaymentsPerToken is a free data retrieval call binding the contract method 0xa73ebdd8.
//
// Solidity: function getTotalChequePaymentsPerToken(address paymentToken) view returns(uint256)
func (_Cmaccount *CmaccountSession) GetTotalChequePaymentsPerToken(paymentToken common.Address) (*big.Int, error) {
	return _Cmaccount.Contract.GetTotalChequePaymentsPerToken(&_Cmaccount.CallOpts, paymentToken)
}

// GetTotalChequePaymentsPerToken is a free data retrieval call binding the contract method 0xa73ebdd8.
//
// Solidity: function getTotalChequePaymentsPerToken(address paymentToken) view returns(uint256)
func (_Cmaccount *CmaccountCallerSession) GetTotalChequePaymentsPerToken(paymentToken common.Address) (*big.Int, error) {
	return _Cmaccount.Contract.GetTotalChequePaymentsPerToken(&_Cmaccount.CallOpts, paymentToken)
}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCaller) GetWantedServiceHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getWantedServiceHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountSession) GetWantedServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetWantedServiceHashes(&_Cmaccount.CallOpts)
}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Cmaccount *CmaccountCallerSession) GetWantedServiceHashes() ([][32]byte, error) {
	return _Cmaccount.Contract.GetWantedServiceHashes(&_Cmaccount.CallOpts)
}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Cmaccount *CmaccountCaller) GetWantedServices(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "getWantedServices")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Cmaccount *CmaccountSession) GetWantedServices() ([]string, error) {
	return _Cmaccount.Contract.GetWantedServices(&_Cmaccount.CallOpts)
}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Cmaccount *CmaccountCallerSession) GetWantedServices() ([]string, error) {
	return _Cmaccount.Contract.GetWantedServices(&_Cmaccount.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccount *CmaccountCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccount *CmaccountSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cmaccount.Contract.HasRole(&_Cmaccount.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cmaccount *CmaccountCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cmaccount.Contract.HasRole(&_Cmaccount.CallOpts, role, account)
}

// HashMessengerCheque is a free data retrieval call binding the contract method 0x5008c8ec.
//
// Solidity: function hashMessengerCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken) pure returns(bytes32)
func (_Cmaccount *CmaccountCaller) HashMessengerCheque(opts *bind.CallOpts, fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "hashMessengerCheque", fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashMessengerCheque is a free data retrieval call binding the contract method 0x5008c8ec.
//
// Solidity: function hashMessengerCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken) pure returns(bytes32)
func (_Cmaccount *CmaccountSession) HashMessengerCheque(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address) ([32]byte, error) {
	return _Cmaccount.Contract.HashMessengerCheque(&_Cmaccount.CallOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken)
}

// HashMessengerCheque is a free data retrieval call binding the contract method 0x5008c8ec.
//
// Solidity: function hashMessengerCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken) pure returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) HashMessengerCheque(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address) ([32]byte, error) {
	return _Cmaccount.Contract.HashMessengerCheque(&_Cmaccount.CallOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x372e802b.
//
// Solidity: function hashTypedDataV4(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken) view returns(bytes32)
func (_Cmaccount *CmaccountCaller) HashTypedDataV4(opts *bind.CallOpts, fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "hashTypedDataV4", fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x372e802b.
//
// Solidity: function hashTypedDataV4(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken) view returns(bytes32)
func (_Cmaccount *CmaccountSession) HashTypedDataV4(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address) ([32]byte, error) {
	return _Cmaccount.Contract.HashTypedDataV4(&_Cmaccount.CallOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken)
}

// HashTypedDataV4 is a free data retrieval call binding the contract method 0x372e802b.
//
// Solidity: function hashTypedDataV4(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken) view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) HashTypedDataV4(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address) ([32]byte, error) {
	return _Cmaccount.Contract.HashTypedDataV4(&_Cmaccount.CallOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken)
}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Cmaccount *CmaccountCaller) IsBotAllowed(opts *bind.CallOpts, bot common.Address) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "isBotAllowed", bot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Cmaccount *CmaccountSession) IsBotAllowed(bot common.Address) (bool, error) {
	return _Cmaccount.Contract.IsBotAllowed(&_Cmaccount.CallOpts, bot)
}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Cmaccount *CmaccountCallerSession) IsBotAllowed(bot common.Address) (bool, error) {
	return _Cmaccount.Contract.IsBotAllowed(&_Cmaccount.CallOpts, bot)
}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Cmaccount *CmaccountCaller) OffChainPaymentSupported(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "offChainPaymentSupported")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Cmaccount *CmaccountSession) OffChainPaymentSupported() (bool, error) {
	return _Cmaccount.Contract.OffChainPaymentSupported(&_Cmaccount.CallOpts)
}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Cmaccount *CmaccountCallerSession) OffChainPaymentSupported() (bool, error) {
	return _Cmaccount.Contract.OffChainPaymentSupported(&_Cmaccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccount *CmaccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccount *CmaccountSession) ProxiableUUID() ([32]byte, error) {
	return _Cmaccount.Contract.ProxiableUUID(&_Cmaccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Cmaccount *CmaccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Cmaccount.Contract.ProxiableUUID(&_Cmaccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccount *CmaccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccount *CmaccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cmaccount.Contract.SupportsInterface(&_Cmaccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cmaccount *CmaccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cmaccount.Contract.SupportsInterface(&_Cmaccount.CallOpts, interfaceId)
}

// VerifyCheque is a free data retrieval call binding the contract method 0x457006db.
//
// Solidity: function verifyCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken, bytes signature) view returns(address signer, uint256 paymentAmount)
func (_Cmaccount *CmaccountCaller) VerifyCheque(opts *bind.CallOpts, fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address, signature []byte) (struct {
	Signer        common.Address
	PaymentAmount *big.Int
}, error) {
	var out []interface{}
	err := _Cmaccount.contract.Call(opts, &out, "verifyCheque", fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken, signature)

	outstruct := new(struct {
		Signer        common.Address
		PaymentAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PaymentAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VerifyCheque is a free data retrieval call binding the contract method 0x457006db.
//
// Solidity: function verifyCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken, bytes signature) view returns(address signer, uint256 paymentAmount)
func (_Cmaccount *CmaccountSession) VerifyCheque(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address, signature []byte) (struct {
	Signer        common.Address
	PaymentAmount *big.Int
}, error) {
	return _Cmaccount.Contract.VerifyCheque(&_Cmaccount.CallOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken, signature)
}

// VerifyCheque is a free data retrieval call binding the contract method 0x457006db.
//
// Solidity: function verifyCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken, bytes signature) view returns(address signer, uint256 paymentAmount)
func (_Cmaccount *CmaccountCallerSession) VerifyCheque(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address, signature []byte) (struct {
	Signer        common.Address
	PaymentAmount *big.Int
}, error) {
	return _Cmaccount.Contract.VerifyCheque(&_Cmaccount.CallOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken, signature)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactor) AcceptCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "acceptCancellation", tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AcceptCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactorSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AcceptCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Cmaccount *CmaccountTransactor) AddMessengerBot(opts *bind.TransactOpts, bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addMessengerBot", bot, gasMoney)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Cmaccount *CmaccountSession) AddMessengerBot(bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddMessengerBot(&_Cmaccount.TransactOpts, bot, gasMoney)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Cmaccount *CmaccountTransactorSession) AddMessengerBot(bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddMessengerBot(&_Cmaccount.TransactOpts, bot, gasMoney)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Cmaccount *CmaccountTransactor) AddPublicKey(opts *bind.TransactOpts, pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addPublicKey", pubKeyAddress, data)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Cmaccount *CmaccountSession) AddPublicKey(pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddPublicKey(&_Cmaccount.TransactOpts, pubKeyAddress, data)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Cmaccount *CmaccountTransactorSession) AddPublicKey(pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddPublicKey(&_Cmaccount.TransactOpts, pubKeyAddress, data)
}

// AddService is a paid mutator transaction binding the contract method 0x432cf639.
//
// Solidity: function addService(string serviceName, uint256 fee, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactor) AddService(opts *bind.TransactOpts, serviceName string, fee *big.Int, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addService", serviceName, fee, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x432cf639.
//
// Solidity: function addService(string serviceName, uint256 fee, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountSession) AddService(serviceName string, fee *big.Int, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddService(&_Cmaccount.TransactOpts, serviceName, fee, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x432cf639.
//
// Solidity: function addService(string serviceName, uint256 fee, bool restrictedRate, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactorSession) AddService(serviceName string, fee *big.Int, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddService(&_Cmaccount.TransactOpts, serviceName, fee, restrictedRate, capabilities)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactor) AddServiceCapability(opts *bind.TransactOpts, serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addServiceCapability", serviceName, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountSession) AddServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactorSession) AddServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactor) AddSupportedToken(opts *bind.TransactOpts, _supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addSupportedToken", _supportedToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountSession) AddSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactorSession) AddSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactor) AddWantedServices(opts *bind.TransactOpts, serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "addWantedServices", serviceNames)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountSession) AddWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactorSession) AddWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.AddWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Cmaccount *CmaccountTransactor) BuyBookingToken(opts *bind.TransactOpts, tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "buyBookingToken", tokenId, expectedPrice, expectedPaymentToken)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Cmaccount *CmaccountSession) BuyBookingToken(tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.BuyBookingToken(&_Cmaccount.TransactOpts, tokenId, expectedPrice, expectedPaymentToken)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Cmaccount *CmaccountTransactorSession) BuyBookingToken(tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.BuyBookingToken(&_Cmaccount.TransactOpts, tokenId, expectedPrice, expectedPaymentToken)
}

// CashInCheque is a paid mutator transaction binding the contract method 0xecaa76ef.
//
// Solidity: function cashInCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken, bytes signature) returns()
func (_Cmaccount *CmaccountTransactor) CashInCheque(opts *bind.TransactOpts, fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address, signature []byte) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "cashInCheque", fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken, signature)
}

// CashInCheque is a paid mutator transaction binding the contract method 0xecaa76ef.
//
// Solidity: function cashInCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken, bytes signature) returns()
func (_Cmaccount *CmaccountSession) CashInCheque(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address, signature []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.CashInCheque(&_Cmaccount.TransactOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken, signature)
}

// CashInCheque is a paid mutator transaction binding the contract method 0xecaa76ef.
//
// Solidity: function cashInCheque(address fromCMAccount, address toCMAccount, address toBot, uint256 counter, uint256 amount, uint256 createdAt, uint256 expiresAt, address paymentToken, bytes signature) returns()
func (_Cmaccount *CmaccountTransactorSession) CashInCheque(fromCMAccount common.Address, toCMAccount common.Address, toBot common.Address, counter *big.Int, amount *big.Int, createdAt *big.Int, expiresAt *big.Int, paymentToken common.Address, signature []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.CashInCheque(&_Cmaccount.TransactOpts, fromCMAccount, toCMAccount, toBot, counter, amount, createdAt, expiresAt, paymentToken, signature)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) CounterCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "counterCancellation", tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Cmaccount *CmaccountSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.CounterCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.CounterCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactor) FinalizeCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "finalizeCancellation", tokenId, refundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountSession) FinalizeCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.FinalizeCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Cmaccount *CmaccountTransactorSession) FinalizeCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.FinalizeCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.GrantRole(&_Cmaccount.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.GrantRole(&_Cmaccount.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Cmaccount *CmaccountTransactor) Initialize(opts *bind.TransactOpts, manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "initialize", manager, bookingToken, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Cmaccount *CmaccountSession) Initialize(manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.Initialize(&_Cmaccount.TransactOpts, manager, bookingToken, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Cmaccount *CmaccountTransactorSession) Initialize(manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.Initialize(&_Cmaccount.TransactOpts, manager, bookingToken, defaultAdmin, upgrader)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) InitiateCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "initiateCancellation", tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Cmaccount *CmaccountSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.InitiateCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.InitiateCancellation(&_Cmaccount.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Cmaccount *CmaccountTransactor) MintBookingToken(opts *bind.TransactOpts, reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "mintBookingToken", reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Cmaccount *CmaccountSession) MintBookingToken(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.MintBookingToken(&_Cmaccount.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Cmaccount *CmaccountTransactorSession) MintBookingToken(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.MintBookingToken(&_Cmaccount.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Cmaccount *CmaccountTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Cmaccount *CmaccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.OnERC721Received(&_Cmaccount.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Cmaccount *CmaccountTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.OnERC721Received(&_Cmaccount.TransactOpts, arg0, arg1, arg2, arg3)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactor) RecordExpiration(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "recordExpiration", tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Cmaccount *CmaccountSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.RecordExpiration(&_Cmaccount.TransactOpts, tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactorSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.RecordExpiration(&_Cmaccount.TransactOpts, tokenId)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) RejectCancellation(opts *bind.TransactOpts, tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "rejectCancellation", tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Cmaccount *CmaccountSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.RejectCancellation(&_Cmaccount.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.RejectCancellation(&_Cmaccount.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Cmaccount *CmaccountTransactor) RemoveAllServices(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeAllServices")
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Cmaccount *CmaccountSession) RemoveAllServices() (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveAllServices(&_Cmaccount.TransactOpts)
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveAllServices() (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveAllServices(&_Cmaccount.TransactOpts)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Cmaccount *CmaccountTransactor) RemoveMessengerBot(opts *bind.TransactOpts, bot common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeMessengerBot", bot)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Cmaccount *CmaccountSession) RemoveMessengerBot(bot common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveMessengerBot(&_Cmaccount.TransactOpts, bot)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveMessengerBot(bot common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveMessengerBot(&_Cmaccount.TransactOpts, bot)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Cmaccount *CmaccountTransactor) RemovePublicKey(opts *bind.TransactOpts, pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removePublicKey", pubKeyAddress)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Cmaccount *CmaccountSession) RemovePublicKey(pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemovePublicKey(&_Cmaccount.TransactOpts, pubKeyAddress)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Cmaccount *CmaccountTransactorSession) RemovePublicKey(pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemovePublicKey(&_Cmaccount.TransactOpts, pubKeyAddress)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Cmaccount *CmaccountTransactor) RemoveService(opts *bind.TransactOpts, serviceName string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeService", serviceName)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Cmaccount *CmaccountSession) RemoveService(serviceName string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveService(&_Cmaccount.TransactOpts, serviceName)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveService(serviceName string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveService(&_Cmaccount.TransactOpts, serviceName)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactor) RemoveServiceCapability(opts *bind.TransactOpts, serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeServiceCapability", serviceName, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountSession) RemoveServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveServiceCapability(&_Cmaccount.TransactOpts, serviceName, capability)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactor) RemoveSupportedToken(opts *bind.TransactOpts, _supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeSupportedToken", _supportedToken)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountSession) RemoveSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveSupportedToken(&_Cmaccount.TransactOpts, _supportedToken)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactor) RemoveWantedServices(opts *bind.TransactOpts, serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "removeWantedServices", serviceNames)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountSession) RemoveWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Cmaccount *CmaccountTransactorSession) RemoveWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.RemoveWantedServices(&_Cmaccount.TransactOpts, serviceNames)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccount *CmaccountTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccount *CmaccountSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RenounceRole(&_Cmaccount.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Cmaccount *CmaccountTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RenounceRole(&_Cmaccount.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RevokeRole(&_Cmaccount.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cmaccount *CmaccountTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cmaccount.Contract.RevokeRole(&_Cmaccount.TransactOpts, role, account)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Cmaccount *CmaccountTransactor) SetGasMoneyWithdrawal(opts *bind.TransactOpts, limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setGasMoneyWithdrawal", limit, period)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Cmaccount *CmaccountSession) SetGasMoneyWithdrawal(limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetGasMoneyWithdrawal(&_Cmaccount.TransactOpts, limit, period)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Cmaccount *CmaccountTransactorSession) SetGasMoneyWithdrawal(limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetGasMoneyWithdrawal(&_Cmaccount.TransactOpts, limit, period)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Cmaccount *CmaccountTransactor) SetOffChainPaymentSupported(opts *bind.TransactOpts, _isSupported bool) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setOffChainPaymentSupported", _isSupported)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Cmaccount *CmaccountSession) SetOffChainPaymentSupported(_isSupported bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetOffChainPaymentSupported(&_Cmaccount.TransactOpts, _isSupported)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Cmaccount *CmaccountTransactorSession) SetOffChainPaymentSupported(_isSupported bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetOffChainPaymentSupported(&_Cmaccount.TransactOpts, _isSupported)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactor) SetServiceCapabilities(opts *bind.TransactOpts, serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setServiceCapabilities", serviceName, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Cmaccount *CmaccountSession) SetServiceCapabilities(serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceCapabilities(&_Cmaccount.TransactOpts, serviceName, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Cmaccount *CmaccountTransactorSession) SetServiceCapabilities(serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceCapabilities(&_Cmaccount.TransactOpts, serviceName, capabilities)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x41bf7c69.
//
// Solidity: function setServiceFee(string serviceName, uint256 fee) returns()
func (_Cmaccount *CmaccountTransactor) SetServiceFee(opts *bind.TransactOpts, serviceName string, fee *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setServiceFee", serviceName, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x41bf7c69.
//
// Solidity: function setServiceFee(string serviceName, uint256 fee) returns()
func (_Cmaccount *CmaccountSession) SetServiceFee(serviceName string, fee *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceFee(&_Cmaccount.TransactOpts, serviceName, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x41bf7c69.
//
// Solidity: function setServiceFee(string serviceName, uint256 fee) returns()
func (_Cmaccount *CmaccountTransactorSession) SetServiceFee(serviceName string, fee *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceFee(&_Cmaccount.TransactOpts, serviceName, fee)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Cmaccount *CmaccountTransactor) SetServiceRestrictedRate(opts *bind.TransactOpts, serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "setServiceRestrictedRate", serviceName, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Cmaccount *CmaccountSession) SetServiceRestrictedRate(serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceRestrictedRate(&_Cmaccount.TransactOpts, serviceName, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Cmaccount *CmaccountTransactorSession) SetServiceRestrictedRate(serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Cmaccount.Contract.SetServiceRestrictedRate(&_Cmaccount.TransactOpts, serviceName, restrictedRate)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactor) TransferERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "transferERC20", token, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Cmaccount *CmaccountSession) TransferERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC20(&_Cmaccount.TransactOpts, token, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactorSession) TransferERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC20(&_Cmaccount.TransactOpts, token, to, amount)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactor) TransferERC721(opts *bind.TransactOpts, token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "transferERC721", token, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Cmaccount *CmaccountSession) TransferERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC721(&_Cmaccount.TransactOpts, token, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Cmaccount *CmaccountTransactorSession) TransferERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.TransferERC721(&_Cmaccount.TransactOpts, token, to, tokenId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccount *CmaccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccount *CmaccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.UpgradeToAndCall(&_Cmaccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Cmaccount *CmaccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Cmaccount.Contract.UpgradeToAndCall(&_Cmaccount.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "withdraw", recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Cmaccount *CmaccountSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.Withdraw(&_Cmaccount.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Cmaccount *CmaccountTransactorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.Withdraw(&_Cmaccount.TransactOpts, recipient, amount)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Cmaccount *CmaccountTransactor) WithdrawCancellation(opts *bind.TransactOpts, tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "withdrawCancellation", tokenId, reason, reasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Cmaccount *CmaccountSession) WithdrawCancellation(tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawCancellation(&_Cmaccount.TransactOpts, tokenId, reason, reasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Cmaccount *CmaccountTransactorSession) WithdrawCancellation(tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawCancellation(&_Cmaccount.TransactOpts, tokenId, reason, reasonVersion)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Cmaccount *CmaccountTransactor) WithdrawGasMoney(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.contract.Transact(opts, "withdrawGasMoney", amount)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Cmaccount *CmaccountSession) WithdrawGasMoney(amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawGasMoney(&_Cmaccount.TransactOpts, amount)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Cmaccount *CmaccountTransactorSession) WithdrawGasMoney(amount *big.Int) (*types.Transaction, error) {
	return _Cmaccount.Contract.WithdrawGasMoney(&_Cmaccount.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cmaccount *CmaccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cmaccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cmaccount *CmaccountSession) Receive() (*types.Transaction, error) {
	return _Cmaccount.Contract.Receive(&_Cmaccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cmaccount *CmaccountTransactorSession) Receive() (*types.Transaction, error) {
	return _Cmaccount.Contract.Receive(&_Cmaccount.TransactOpts)
}

// CmaccountCMAccountUpgradedIterator is returned from FilterCMAccountUpgraded and is used to iterate over the raw logs and unpacked data for CMAccountUpgraded events raised by the Cmaccount contract.
type CmaccountCMAccountUpgradedIterator struct {
	Event *CmaccountCMAccountUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountCMAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountCMAccountUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountCMAccountUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountCMAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountCMAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountCMAccountUpgraded represents a CMAccountUpgraded event raised by the Cmaccount contract.
type CmaccountCMAccountUpgraded struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCMAccountUpgraded is a free log retrieval operation binding the contract event 0xa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db4.
//
// Solidity: event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccount *CmaccountFilterer) FilterCMAccountUpgraded(opts *bind.FilterOpts, oldImplementation []common.Address, newImplementation []common.Address) (*CmaccountCMAccountUpgradedIterator, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "CMAccountUpgraded", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountCMAccountUpgradedIterator{contract: _Cmaccount.contract, event: "CMAccountUpgraded", logs: logs, sub: sub}, nil
}

// WatchCMAccountUpgraded is a free log subscription operation binding the contract event 0xa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db4.
//
// Solidity: event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccount *CmaccountFilterer) WatchCMAccountUpgraded(opts *bind.WatchOpts, sink chan<- *CmaccountCMAccountUpgraded, oldImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "CMAccountUpgraded", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountCMAccountUpgraded)
				if err := _Cmaccount.contract.UnpackLog(event, "CMAccountUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCMAccountUpgraded is a log parse operation binding the contract event 0xa3d484f827e1c900ce24494bfdb214bcbad08472a9f0571fb5beac779a682db4.
//
// Solidity: event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Cmaccount *CmaccountFilterer) ParseCMAccountUpgraded(log types.Log) (*CmaccountCMAccountUpgraded, error) {
	event := new(CmaccountCMAccountUpgraded)
	if err := _Cmaccount.contract.UnpackLog(event, "CMAccountUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountChequeCashedInIterator is returned from FilterChequeCashedIn and is used to iterate over the raw logs and unpacked data for ChequeCashedIn events raised by the Cmaccount contract.
type CmaccountChequeCashedInIterator struct {
	Event *CmaccountChequeCashedIn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountChequeCashedInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountChequeCashedIn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountChequeCashedIn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountChequeCashedInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountChequeCashedInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountChequeCashedIn represents a ChequeCashedIn event raised by the Cmaccount contract.
type CmaccountChequeCashedIn struct {
	FromCMAccount    common.Address
	ToCMAccount      common.Address
	FromBot          common.Address
	ToBot            common.Address
	Counter          *big.Int
	Amount           *big.Int
	PaidChequeAmount *big.Int
	PaidDeveloperFee *big.Int
	PaymentToken     common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChequeCashedIn is a free log retrieval operation binding the contract event 0xa7708e82cb201b0c0dc3d520642d0e0eb290d001b6acec29d01aeeb6af7dab20.
//
// Solidity: event ChequeCashedIn(address indexed fromCMAccount, address indexed toCMAccount, address fromBot, address toBot, uint256 counter, uint256 amount, uint256 paidChequeAmount, uint256 paidDeveloperFee, address paymentToken)
func (_Cmaccount *CmaccountFilterer) FilterChequeCashedIn(opts *bind.FilterOpts, fromCMAccount []common.Address, toCMAccount []common.Address) (*CmaccountChequeCashedInIterator, error) {

	var fromCMAccountRule []interface{}
	for _, fromCMAccountItem := range fromCMAccount {
		fromCMAccountRule = append(fromCMAccountRule, fromCMAccountItem)
	}
	var toCMAccountRule []interface{}
	for _, toCMAccountItem := range toCMAccount {
		toCMAccountRule = append(toCMAccountRule, toCMAccountItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ChequeCashedIn", fromCMAccountRule, toCMAccountRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountChequeCashedInIterator{contract: _Cmaccount.contract, event: "ChequeCashedIn", logs: logs, sub: sub}, nil
}

// WatchChequeCashedIn is a free log subscription operation binding the contract event 0xa7708e82cb201b0c0dc3d520642d0e0eb290d001b6acec29d01aeeb6af7dab20.
//
// Solidity: event ChequeCashedIn(address indexed fromCMAccount, address indexed toCMAccount, address fromBot, address toBot, uint256 counter, uint256 amount, uint256 paidChequeAmount, uint256 paidDeveloperFee, address paymentToken)
func (_Cmaccount *CmaccountFilterer) WatchChequeCashedIn(opts *bind.WatchOpts, sink chan<- *CmaccountChequeCashedIn, fromCMAccount []common.Address, toCMAccount []common.Address) (event.Subscription, error) {

	var fromCMAccountRule []interface{}
	for _, fromCMAccountItem := range fromCMAccount {
		fromCMAccountRule = append(fromCMAccountRule, fromCMAccountItem)
	}
	var toCMAccountRule []interface{}
	for _, toCMAccountItem := range toCMAccount {
		toCMAccountRule = append(toCMAccountRule, toCMAccountItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ChequeCashedIn", fromCMAccountRule, toCMAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountChequeCashedIn)
				if err := _Cmaccount.contract.UnpackLog(event, "ChequeCashedIn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChequeCashedIn is a log parse operation binding the contract event 0xa7708e82cb201b0c0dc3d520642d0e0eb290d001b6acec29d01aeeb6af7dab20.
//
// Solidity: event ChequeCashedIn(address indexed fromCMAccount, address indexed toCMAccount, address fromBot, address toBot, uint256 counter, uint256 amount, uint256 paidChequeAmount, uint256 paidDeveloperFee, address paymentToken)
func (_Cmaccount *CmaccountFilterer) ParseChequeCashedIn(log types.Log) (*CmaccountChequeCashedIn, error) {
	event := new(CmaccountChequeCashedIn)
	if err := _Cmaccount.contract.UnpackLog(event, "ChequeCashedIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Cmaccount contract.
type CmaccountDepositIterator struct {
	Event *CmaccountDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountDeposit represents a Deposit event raised by the Cmaccount contract.
type CmaccountDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Cmaccount *CmaccountFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*CmaccountDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountDepositIterator{contract: _Cmaccount.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Cmaccount *CmaccountFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CmaccountDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountDeposit)
				if err := _Cmaccount.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Cmaccount *CmaccountFilterer) ParseDeposit(log types.Log) (*CmaccountDeposit, error) {
	event := new(CmaccountDeposit)
	if err := _Cmaccount.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountGasMoneyWithdrawalIterator is returned from FilterGasMoneyWithdrawal and is used to iterate over the raw logs and unpacked data for GasMoneyWithdrawal events raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawalIterator struct {
	Event *CmaccountGasMoneyWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountGasMoneyWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountGasMoneyWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountGasMoneyWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountGasMoneyWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountGasMoneyWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountGasMoneyWithdrawal represents a GasMoneyWithdrawal event raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawal struct {
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasMoneyWithdrawal is a free log retrieval operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Cmaccount *CmaccountFilterer) FilterGasMoneyWithdrawal(opts *bind.FilterOpts, withdrawer []common.Address) (*CmaccountGasMoneyWithdrawalIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "GasMoneyWithdrawal", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountGasMoneyWithdrawalIterator{contract: _Cmaccount.contract, event: "GasMoneyWithdrawal", logs: logs, sub: sub}, nil
}

// WatchGasMoneyWithdrawal is a free log subscription operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Cmaccount *CmaccountFilterer) WatchGasMoneyWithdrawal(opts *bind.WatchOpts, sink chan<- *CmaccountGasMoneyWithdrawal, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "GasMoneyWithdrawal", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountGasMoneyWithdrawal)
				if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGasMoneyWithdrawal is a log parse operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Cmaccount *CmaccountFilterer) ParseGasMoneyWithdrawal(log types.Log) (*CmaccountGasMoneyWithdrawal, error) {
	event := new(CmaccountGasMoneyWithdrawal)
	if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountGasMoneyWithdrawalUpdatedIterator is returned from FilterGasMoneyWithdrawalUpdated and is used to iterate over the raw logs and unpacked data for GasMoneyWithdrawalUpdated events raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawalUpdatedIterator struct {
	Event *CmaccountGasMoneyWithdrawalUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountGasMoneyWithdrawalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountGasMoneyWithdrawalUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountGasMoneyWithdrawalUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountGasMoneyWithdrawalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountGasMoneyWithdrawalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountGasMoneyWithdrawalUpdated represents a GasMoneyWithdrawalUpdated event raised by the Cmaccount contract.
type CmaccountGasMoneyWithdrawalUpdated struct {
	Limit  *big.Int
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGasMoneyWithdrawalUpdated is a free log retrieval operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Cmaccount *CmaccountFilterer) FilterGasMoneyWithdrawalUpdated(opts *bind.FilterOpts) (*CmaccountGasMoneyWithdrawalUpdatedIterator, error) {

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "GasMoneyWithdrawalUpdated")
	if err != nil {
		return nil, err
	}
	return &CmaccountGasMoneyWithdrawalUpdatedIterator{contract: _Cmaccount.contract, event: "GasMoneyWithdrawalUpdated", logs: logs, sub: sub}, nil
}

// WatchGasMoneyWithdrawalUpdated is a free log subscription operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Cmaccount *CmaccountFilterer) WatchGasMoneyWithdrawalUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountGasMoneyWithdrawalUpdated) (event.Subscription, error) {

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "GasMoneyWithdrawalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountGasMoneyWithdrawalUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawalUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGasMoneyWithdrawalUpdated is a log parse operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Cmaccount *CmaccountFilterer) ParseGasMoneyWithdrawalUpdated(log types.Log) (*CmaccountGasMoneyWithdrawalUpdated, error) {
	event := new(CmaccountGasMoneyWithdrawalUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "GasMoneyWithdrawalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cmaccount contract.
type CmaccountInitializedIterator struct {
	Event *CmaccountInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountInitialized represents a Initialized event raised by the Cmaccount contract.
type CmaccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccount *CmaccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*CmaccountInitializedIterator, error) {

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CmaccountInitializedIterator{contract: _Cmaccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccount *CmaccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CmaccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountInitialized)
				if err := _Cmaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Cmaccount *CmaccountFilterer) ParseInitialized(log types.Log) (*CmaccountInitialized, error) {
	event := new(CmaccountInitialized)
	if err := _Cmaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountMessengerBotAddedIterator is returned from FilterMessengerBotAdded and is used to iterate over the raw logs and unpacked data for MessengerBotAdded events raised by the Cmaccount contract.
type CmaccountMessengerBotAddedIterator struct {
	Event *CmaccountMessengerBotAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountMessengerBotAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountMessengerBotAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountMessengerBotAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountMessengerBotAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountMessengerBotAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountMessengerBotAdded represents a MessengerBotAdded event raised by the Cmaccount contract.
type CmaccountMessengerBotAdded struct {
	Bot common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMessengerBotAdded is a free log retrieval operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Cmaccount *CmaccountFilterer) FilterMessengerBotAdded(opts *bind.FilterOpts, bot []common.Address) (*CmaccountMessengerBotAddedIterator, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "MessengerBotAdded", botRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountMessengerBotAddedIterator{contract: _Cmaccount.contract, event: "MessengerBotAdded", logs: logs, sub: sub}, nil
}

// WatchMessengerBotAdded is a free log subscription operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Cmaccount *CmaccountFilterer) WatchMessengerBotAdded(opts *bind.WatchOpts, sink chan<- *CmaccountMessengerBotAdded, bot []common.Address) (event.Subscription, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "MessengerBotAdded", botRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountMessengerBotAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessengerBotAdded is a log parse operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Cmaccount *CmaccountFilterer) ParseMessengerBotAdded(log types.Log) (*CmaccountMessengerBotAdded, error) {
	event := new(CmaccountMessengerBotAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountMessengerBotRemovedIterator is returned from FilterMessengerBotRemoved and is used to iterate over the raw logs and unpacked data for MessengerBotRemoved events raised by the Cmaccount contract.
type CmaccountMessengerBotRemovedIterator struct {
	Event *CmaccountMessengerBotRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountMessengerBotRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountMessengerBotRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountMessengerBotRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountMessengerBotRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountMessengerBotRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountMessengerBotRemoved represents a MessengerBotRemoved event raised by the Cmaccount contract.
type CmaccountMessengerBotRemoved struct {
	Bot common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMessengerBotRemoved is a free log retrieval operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Cmaccount *CmaccountFilterer) FilterMessengerBotRemoved(opts *bind.FilterOpts, bot []common.Address) (*CmaccountMessengerBotRemovedIterator, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "MessengerBotRemoved", botRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountMessengerBotRemovedIterator{contract: _Cmaccount.contract, event: "MessengerBotRemoved", logs: logs, sub: sub}, nil
}

// WatchMessengerBotRemoved is a free log subscription operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Cmaccount *CmaccountFilterer) WatchMessengerBotRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountMessengerBotRemoved, bot []common.Address) (event.Subscription, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "MessengerBotRemoved", botRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountMessengerBotRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessengerBotRemoved is a log parse operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Cmaccount *CmaccountFilterer) ParseMessengerBotRemoved(log types.Log) (*CmaccountMessengerBotRemoved, error) {
	event := new(CmaccountMessengerBotRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "MessengerBotRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountOffChainPaymentSupportUpdatedIterator is returned from FilterOffChainPaymentSupportUpdated and is used to iterate over the raw logs and unpacked data for OffChainPaymentSupportUpdated events raised by the Cmaccount contract.
type CmaccountOffChainPaymentSupportUpdatedIterator struct {
	Event *CmaccountOffChainPaymentSupportUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountOffChainPaymentSupportUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountOffChainPaymentSupportUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountOffChainPaymentSupportUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountOffChainPaymentSupportUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountOffChainPaymentSupportUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountOffChainPaymentSupportUpdated represents a OffChainPaymentSupportUpdated event raised by the Cmaccount contract.
type CmaccountOffChainPaymentSupportUpdated struct {
	SupportsOffChainPayment bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOffChainPaymentSupportUpdated is a free log retrieval operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Cmaccount *CmaccountFilterer) FilterOffChainPaymentSupportUpdated(opts *bind.FilterOpts) (*CmaccountOffChainPaymentSupportUpdatedIterator, error) {

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "OffChainPaymentSupportUpdated")
	if err != nil {
		return nil, err
	}
	return &CmaccountOffChainPaymentSupportUpdatedIterator{contract: _Cmaccount.contract, event: "OffChainPaymentSupportUpdated", logs: logs, sub: sub}, nil
}

// WatchOffChainPaymentSupportUpdated is a free log subscription operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Cmaccount *CmaccountFilterer) WatchOffChainPaymentSupportUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountOffChainPaymentSupportUpdated) (event.Subscription, error) {

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "OffChainPaymentSupportUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountOffChainPaymentSupportUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "OffChainPaymentSupportUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOffChainPaymentSupportUpdated is a log parse operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Cmaccount *CmaccountFilterer) ParseOffChainPaymentSupportUpdated(log types.Log) (*CmaccountOffChainPaymentSupportUpdated, error) {
	event := new(CmaccountOffChainPaymentSupportUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "OffChainPaymentSupportUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPaymentTokenAddedIterator is returned from FilterPaymentTokenAdded and is used to iterate over the raw logs and unpacked data for PaymentTokenAdded events raised by the Cmaccount contract.
type CmaccountPaymentTokenAddedIterator struct {
	Event *CmaccountPaymentTokenAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountPaymentTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPaymentTokenAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountPaymentTokenAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountPaymentTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPaymentTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPaymentTokenAdded represents a PaymentTokenAdded event raised by the Cmaccount contract.
type CmaccountPaymentTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenAdded is a free log retrieval operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Cmaccount *CmaccountFilterer) FilterPaymentTokenAdded(opts *bind.FilterOpts, token []common.Address) (*CmaccountPaymentTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PaymentTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPaymentTokenAddedIterator{contract: _Cmaccount.contract, event: "PaymentTokenAdded", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenAdded is a free log subscription operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Cmaccount *CmaccountFilterer) WatchPaymentTokenAdded(opts *bind.WatchOpts, sink chan<- *CmaccountPaymentTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PaymentTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPaymentTokenAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentTokenAdded is a log parse operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Cmaccount *CmaccountFilterer) ParsePaymentTokenAdded(log types.Log) (*CmaccountPaymentTokenAdded, error) {
	event := new(CmaccountPaymentTokenAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPaymentTokenRemovedIterator is returned from FilterPaymentTokenRemoved and is used to iterate over the raw logs and unpacked data for PaymentTokenRemoved events raised by the Cmaccount contract.
type CmaccountPaymentTokenRemovedIterator struct {
	Event *CmaccountPaymentTokenRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountPaymentTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPaymentTokenRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountPaymentTokenRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountPaymentTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPaymentTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPaymentTokenRemoved represents a PaymentTokenRemoved event raised by the Cmaccount contract.
type CmaccountPaymentTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenRemoved is a free log retrieval operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Cmaccount *CmaccountFilterer) FilterPaymentTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*CmaccountPaymentTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PaymentTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPaymentTokenRemovedIterator{contract: _Cmaccount.contract, event: "PaymentTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenRemoved is a free log subscription operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Cmaccount *CmaccountFilterer) WatchPaymentTokenRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountPaymentTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PaymentTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPaymentTokenRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentTokenRemoved is a log parse operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Cmaccount *CmaccountFilterer) ParsePaymentTokenRemoved(log types.Log) (*CmaccountPaymentTokenRemoved, error) {
	event := new(CmaccountPaymentTokenRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "PaymentTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPublicKeyAddedIterator is returned from FilterPublicKeyAdded and is used to iterate over the raw logs and unpacked data for PublicKeyAdded events raised by the Cmaccount contract.
type CmaccountPublicKeyAddedIterator struct {
	Event *CmaccountPublicKeyAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountPublicKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPublicKeyAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountPublicKeyAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountPublicKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPublicKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPublicKeyAdded represents a PublicKeyAdded event raised by the Cmaccount contract.
type CmaccountPublicKeyAdded struct {
	PubKeyAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyAdded is a free log retrieval operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) FilterPublicKeyAdded(opts *bind.FilterOpts, pubKeyAddress []common.Address) (*CmaccountPublicKeyAddedIterator, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PublicKeyAdded", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPublicKeyAddedIterator{contract: _Cmaccount.contract, event: "PublicKeyAdded", logs: logs, sub: sub}, nil
}

// WatchPublicKeyAdded is a free log subscription operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) WatchPublicKeyAdded(opts *bind.WatchOpts, sink chan<- *CmaccountPublicKeyAdded, pubKeyAddress []common.Address) (event.Subscription, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PublicKeyAdded", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPublicKeyAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePublicKeyAdded is a log parse operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) ParsePublicKeyAdded(log types.Log) (*CmaccountPublicKeyAdded, error) {
	event := new(CmaccountPublicKeyAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountPublicKeyRemovedIterator is returned from FilterPublicKeyRemoved and is used to iterate over the raw logs and unpacked data for PublicKeyRemoved events raised by the Cmaccount contract.
type CmaccountPublicKeyRemovedIterator struct {
	Event *CmaccountPublicKeyRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountPublicKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountPublicKeyRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountPublicKeyRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountPublicKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountPublicKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountPublicKeyRemoved represents a PublicKeyRemoved event raised by the Cmaccount contract.
type CmaccountPublicKeyRemoved struct {
	PubKeyAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyRemoved is a free log retrieval operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) FilterPublicKeyRemoved(opts *bind.FilterOpts, pubKeyAddress []common.Address) (*CmaccountPublicKeyRemovedIterator, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "PublicKeyRemoved", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountPublicKeyRemovedIterator{contract: _Cmaccount.contract, event: "PublicKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchPublicKeyRemoved is a free log subscription operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) WatchPublicKeyRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountPublicKeyRemoved, pubKeyAddress []common.Address) (event.Subscription, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "PublicKeyRemoved", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountPublicKeyRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePublicKeyRemoved is a log parse operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Cmaccount *CmaccountFilterer) ParsePublicKeyRemoved(log types.Log) (*CmaccountPublicKeyRemoved, error) {
	event := new(CmaccountPublicKeyRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "PublicKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Cmaccount contract.
type CmaccountRoleAdminChangedIterator struct {
	Event *CmaccountRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountRoleAdminChanged represents a RoleAdminChanged event raised by the Cmaccount contract.
type CmaccountRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccount *CmaccountFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CmaccountRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountRoleAdminChangedIterator{contract: _Cmaccount.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccount *CmaccountFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CmaccountRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountRoleAdminChanged)
				if err := _Cmaccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cmaccount *CmaccountFilterer) ParseRoleAdminChanged(log types.Log) (*CmaccountRoleAdminChanged, error) {
	event := new(CmaccountRoleAdminChanged)
	if err := _Cmaccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Cmaccount contract.
type CmaccountRoleGrantedIterator struct {
	Event *CmaccountRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountRoleGranted represents a RoleGranted event raised by the Cmaccount contract.
type CmaccountRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CmaccountRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountRoleGrantedIterator{contract: _Cmaccount.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CmaccountRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountRoleGranted)
				if err := _Cmaccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) ParseRoleGranted(log types.Log) (*CmaccountRoleGranted, error) {
	event := new(CmaccountRoleGranted)
	if err := _Cmaccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Cmaccount contract.
type CmaccountRoleRevokedIterator struct {
	Event *CmaccountRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountRoleRevoked represents a RoleRevoked event raised by the Cmaccount contract.
type CmaccountRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CmaccountRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountRoleRevokedIterator{contract: _Cmaccount.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CmaccountRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountRoleRevoked)
				if err := _Cmaccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cmaccount *CmaccountFilterer) ParseRoleRevoked(log types.Log) (*CmaccountRoleRevoked, error) {
	event := new(CmaccountRoleRevoked)
	if err := _Cmaccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceAddedIterator is returned from FilterServiceAdded and is used to iterate over the raw logs and unpacked data for ServiceAdded events raised by the Cmaccount contract.
type CmaccountServiceAddedIterator struct {
	Event *CmaccountServiceAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountServiceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountServiceAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountServiceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceAdded represents a ServiceAdded event raised by the Cmaccount contract.
type CmaccountServiceAdded struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceAdded is a free log retrieval operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterServiceAdded(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceAddedIterator{contract: _Cmaccount.contract, event: "ServiceAdded", logs: logs, sub: sub}, nil
}

// WatchServiceAdded is a free log subscription operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchServiceAdded(opts *bind.WatchOpts, sink chan<- *CmaccountServiceAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceAdded is a log parse operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseServiceAdded(log types.Log) (*CmaccountServiceAdded, error) {
	event := new(CmaccountServiceAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceCapabilitiesUpdatedIterator is returned from FilterServiceCapabilitiesUpdated and is used to iterate over the raw logs and unpacked data for ServiceCapabilitiesUpdated events raised by the Cmaccount contract.
type CmaccountServiceCapabilitiesUpdatedIterator struct {
	Event *CmaccountServiceCapabilitiesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountServiceCapabilitiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceCapabilitiesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountServiceCapabilitiesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountServiceCapabilitiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceCapabilitiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceCapabilitiesUpdated represents a ServiceCapabilitiesUpdated event raised by the Cmaccount contract.
type CmaccountServiceCapabilitiesUpdated struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilitiesUpdated is a free log retrieval operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterServiceCapabilitiesUpdated(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceCapabilitiesUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceCapabilitiesUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceCapabilitiesUpdatedIterator{contract: _Cmaccount.contract, event: "ServiceCapabilitiesUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilitiesUpdated is a free log subscription operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchServiceCapabilitiesUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountServiceCapabilitiesUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceCapabilitiesUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceCapabilitiesUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilitiesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceCapabilitiesUpdated is a log parse operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseServiceCapabilitiesUpdated(log types.Log) (*CmaccountServiceCapabilitiesUpdated, error) {
	event := new(CmaccountServiceCapabilitiesUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilitiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceCapabilityAddedIterator is returned from FilterServiceCapabilityAdded and is used to iterate over the raw logs and unpacked data for ServiceCapabilityAdded events raised by the Cmaccount contract.
type CmaccountServiceCapabilityAddedIterator struct {
	Event *CmaccountServiceCapabilityAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountServiceCapabilityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceCapabilityAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountServiceCapabilityAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountServiceCapabilityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceCapabilityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceCapabilityAdded represents a ServiceCapabilityAdded event raised by the Cmaccount contract.
type CmaccountServiceCapabilityAdded struct {
	ServiceName common.Hash
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityAdded is a free log retrieval operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) FilterServiceCapabilityAdded(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceCapabilityAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceCapabilityAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceCapabilityAddedIterator{contract: _Cmaccount.contract, event: "ServiceCapabilityAdded", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityAdded is a free log subscription operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) WatchServiceCapabilityAdded(opts *bind.WatchOpts, sink chan<- *CmaccountServiceCapabilityAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceCapabilityAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceCapabilityAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceCapabilityAdded is a log parse operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) ParseServiceCapabilityAdded(log types.Log) (*CmaccountServiceCapabilityAdded, error) {
	event := new(CmaccountServiceCapabilityAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceCapabilityRemovedIterator is returned from FilterServiceCapabilityRemoved and is used to iterate over the raw logs and unpacked data for ServiceCapabilityRemoved events raised by the Cmaccount contract.
type CmaccountServiceCapabilityRemovedIterator struct {
	Event *CmaccountServiceCapabilityRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountServiceCapabilityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceCapabilityRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountServiceCapabilityRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountServiceCapabilityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceCapabilityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceCapabilityRemoved represents a ServiceCapabilityRemoved event raised by the Cmaccount contract.
type CmaccountServiceCapabilityRemoved struct {
	ServiceName common.Hash
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityRemoved is a free log retrieval operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) FilterServiceCapabilityRemoved(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceCapabilityRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceCapabilityRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceCapabilityRemovedIterator{contract: _Cmaccount.contract, event: "ServiceCapabilityRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityRemoved is a free log subscription operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) WatchServiceCapabilityRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountServiceCapabilityRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceCapabilityRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceCapabilityRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceCapabilityRemoved is a log parse operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Cmaccount *CmaccountFilterer) ParseServiceCapabilityRemoved(log types.Log) (*CmaccountServiceCapabilityRemoved, error) {
	event := new(CmaccountServiceCapabilityRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceCapabilityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceFeeUpdatedIterator is returned from FilterServiceFeeUpdated and is used to iterate over the raw logs and unpacked data for ServiceFeeUpdated events raised by the Cmaccount contract.
type CmaccountServiceFeeUpdatedIterator struct {
	Event *CmaccountServiceFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountServiceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountServiceFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountServiceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceFeeUpdated represents a ServiceFeeUpdated event raised by the Cmaccount contract.
type CmaccountServiceFeeUpdated struct {
	ServiceName common.Hash
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeUpdated is a free log retrieval operation binding the contract event 0xdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41.
//
// Solidity: event ServiceFeeUpdated(string indexed serviceName, uint256 fee)
func (_Cmaccount *CmaccountFilterer) FilterServiceFeeUpdated(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceFeeUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceFeeUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceFeeUpdatedIterator{contract: _Cmaccount.contract, event: "ServiceFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceFeeUpdated is a free log subscription operation binding the contract event 0xdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41.
//
// Solidity: event ServiceFeeUpdated(string indexed serviceName, uint256 fee)
func (_Cmaccount *CmaccountFilterer) WatchServiceFeeUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountServiceFeeUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceFeeUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceFeeUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceFeeUpdated is a log parse operation binding the contract event 0xdd6c54a4503e1d8a1e75d73648f77d8fe66234b437ce30e20edd51563116ec41.
//
// Solidity: event ServiceFeeUpdated(string indexed serviceName, uint256 fee)
func (_Cmaccount *CmaccountFilterer) ParseServiceFeeUpdated(log types.Log) (*CmaccountServiceFeeUpdated, error) {
	event := new(CmaccountServiceFeeUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceRemovedIterator is returned from FilterServiceRemoved and is used to iterate over the raw logs and unpacked data for ServiceRemoved events raised by the Cmaccount contract.
type CmaccountServiceRemovedIterator struct {
	Event *CmaccountServiceRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountServiceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountServiceRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountServiceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceRemoved represents a ServiceRemoved event raised by the Cmaccount contract.
type CmaccountServiceRemoved struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceRemoved is a free log retrieval operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterServiceRemoved(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceRemovedIterator{contract: _Cmaccount.contract, event: "ServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceRemoved is a free log subscription operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchServiceRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountServiceRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceRemoved is a log parse operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseServiceRemoved(log types.Log) (*CmaccountServiceRemoved, error) {
	event := new(CmaccountServiceRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountServiceRestrictedRateUpdatedIterator is returned from FilterServiceRestrictedRateUpdated and is used to iterate over the raw logs and unpacked data for ServiceRestrictedRateUpdated events raised by the Cmaccount contract.
type CmaccountServiceRestrictedRateUpdatedIterator struct {
	Event *CmaccountServiceRestrictedRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountServiceRestrictedRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountServiceRestrictedRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountServiceRestrictedRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountServiceRestrictedRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountServiceRestrictedRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountServiceRestrictedRateUpdated represents a ServiceRestrictedRateUpdated event raised by the Cmaccount contract.
type CmaccountServiceRestrictedRateUpdated struct {
	ServiceName    common.Hash
	RestrictedRate bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterServiceRestrictedRateUpdated is a free log retrieval operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Cmaccount *CmaccountFilterer) FilterServiceRestrictedRateUpdated(opts *bind.FilterOpts, serviceName []string) (*CmaccountServiceRestrictedRateUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "ServiceRestrictedRateUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountServiceRestrictedRateUpdatedIterator{contract: _Cmaccount.contract, event: "ServiceRestrictedRateUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceRestrictedRateUpdated is a free log subscription operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Cmaccount *CmaccountFilterer) WatchServiceRestrictedRateUpdated(opts *bind.WatchOpts, sink chan<- *CmaccountServiceRestrictedRateUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "ServiceRestrictedRateUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountServiceRestrictedRateUpdated)
				if err := _Cmaccount.contract.UnpackLog(event, "ServiceRestrictedRateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceRestrictedRateUpdated is a log parse operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Cmaccount *CmaccountFilterer) ParseServiceRestrictedRateUpdated(log types.Log) (*CmaccountServiceRestrictedRateUpdated, error) {
	event := new(CmaccountServiceRestrictedRateUpdated)
	if err := _Cmaccount.contract.UnpackLog(event, "ServiceRestrictedRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Cmaccount contract.
type CmaccountUpgradedIterator struct {
	Event *CmaccountUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountUpgraded represents a Upgraded event raised by the Cmaccount contract.
type CmaccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccount *CmaccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CmaccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountUpgradedIterator{contract: _Cmaccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccount *CmaccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CmaccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountUpgraded)
				if err := _Cmaccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Cmaccount *CmaccountFilterer) ParseUpgraded(log types.Log) (*CmaccountUpgraded, error) {
	event := new(CmaccountUpgraded)
	if err := _Cmaccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountWantedServiceAddedIterator is returned from FilterWantedServiceAdded and is used to iterate over the raw logs and unpacked data for WantedServiceAdded events raised by the Cmaccount contract.
type CmaccountWantedServiceAddedIterator struct {
	Event *CmaccountWantedServiceAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountWantedServiceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountWantedServiceAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountWantedServiceAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountWantedServiceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountWantedServiceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountWantedServiceAdded represents a WantedServiceAdded event raised by the Cmaccount contract.
type CmaccountWantedServiceAdded struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceAdded is a free log retrieval operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterWantedServiceAdded(opts *bind.FilterOpts, serviceName []string) (*CmaccountWantedServiceAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "WantedServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountWantedServiceAddedIterator{contract: _Cmaccount.contract, event: "WantedServiceAdded", logs: logs, sub: sub}, nil
}

// WatchWantedServiceAdded is a free log subscription operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchWantedServiceAdded(opts *bind.WatchOpts, sink chan<- *CmaccountWantedServiceAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "WantedServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountWantedServiceAdded)
				if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWantedServiceAdded is a log parse operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseWantedServiceAdded(log types.Log) (*CmaccountWantedServiceAdded, error) {
	event := new(CmaccountWantedServiceAdded)
	if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountWantedServiceRemovedIterator is returned from FilterWantedServiceRemoved and is used to iterate over the raw logs and unpacked data for WantedServiceRemoved events raised by the Cmaccount contract.
type CmaccountWantedServiceRemovedIterator struct {
	Event *CmaccountWantedServiceRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountWantedServiceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountWantedServiceRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountWantedServiceRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountWantedServiceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountWantedServiceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountWantedServiceRemoved represents a WantedServiceRemoved event raised by the Cmaccount contract.
type CmaccountWantedServiceRemoved struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceRemoved is a free log retrieval operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) FilterWantedServiceRemoved(opts *bind.FilterOpts, serviceName []string) (*CmaccountWantedServiceRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "WantedServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountWantedServiceRemovedIterator{contract: _Cmaccount.contract, event: "WantedServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchWantedServiceRemoved is a free log subscription operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) WatchWantedServiceRemoved(opts *bind.WatchOpts, sink chan<- *CmaccountWantedServiceRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "WantedServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountWantedServiceRemoved)
				if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWantedServiceRemoved is a log parse operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Cmaccount *CmaccountFilterer) ParseWantedServiceRemoved(log types.Log) (*CmaccountWantedServiceRemoved, error) {
	event := new(CmaccountWantedServiceRemoved)
	if err := _Cmaccount.contract.UnpackLog(event, "WantedServiceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CmaccountWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Cmaccount contract.
type CmaccountWithdrawIterator struct {
	Event *CmaccountWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CmaccountWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CmaccountWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CmaccountWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CmaccountWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CmaccountWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CmaccountWithdraw represents a Withdraw event raised by the Cmaccount contract.
type CmaccountWithdraw struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Cmaccount *CmaccountFilterer) FilterWithdraw(opts *bind.FilterOpts, receiver []common.Address) (*CmaccountWithdrawIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Cmaccount.contract.FilterLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return &CmaccountWithdrawIterator{contract: _Cmaccount.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Cmaccount *CmaccountFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CmaccountWithdraw, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Cmaccount.contract.WatchLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CmaccountWithdraw)
				if err := _Cmaccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Cmaccount *CmaccountFilterer) ParseWithdraw(log types.Log) (*CmaccountWithdraw, error) {
	event := new(CmaccountWithdraw)
	if err := _Cmaccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
