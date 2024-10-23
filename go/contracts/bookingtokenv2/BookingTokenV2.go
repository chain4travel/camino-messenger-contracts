// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bookingtokenv2

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

// Bookingtokenv2MetaData contains all meta data concerning the Bookingtokenv2 contract.
var Bookingtokenv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minExpirationTimestampDiff\",\"type\":\"uint256\"}],\"name\":\"ExpirationTimestampTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"IncorrectAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservationPrice\",\"type\":\"uint256\"}],\"name\":\"IncorrectPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumBookingToken.BookingStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidTokenStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoCounteredCancellationProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoPendingCancellationProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotAuthorizedToAcceptCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotAuthorizedToAcceptCounterProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotAuthorizedToCancelProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotAuthorizedToCounterCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotAuthorizedToInitiateCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotAuthorizedToRejectCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotAuthorizedToSetCancellable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotCMAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"name\":\"ReservationExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"ReservationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"}],\"name\":\"SupplierIsNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenHasActiveCancellationProposalOrCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"}],\"name\":\"TokenIsReserved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"acceptedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"CancellationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"counteredBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRefundAmount\",\"type\":\"uint256\"}],\"name\":\"CancellationCountered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"CancellationPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"acceptedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"CancellationProposalAcceptedByTheOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cancelledBy\",\"type\":\"address\"}],\"name\":\"CancellationProposalCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rejectedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumCancellationRejectionReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"CancellationRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isCancellable\",\"type\":\"bool\"}],\"name\":\"TokenCancellableUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isCancellable\",\"type\":\"bool\"}],\"name\":\"TokenReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_EXPIRATION_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkRefundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellationProposal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkRefundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCounteredCancellationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyReservedToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelCancellationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newRefundAmount\",\"type\":\"uint256\"}],\"name\":\"counterCancellationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBookingStatus\",\"outputs\":[{\"internalType\":\"enumBookingToken.BookingStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCancellationProposalRefundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCancellationProposalStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposedBy\",\"type\":\"address\"},{\"internalType\":\"enumCancellationProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumCancellationRejectionReason\",\"name\":\"rejectionReason\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinExpirationTimestampDiff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReservationPaymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReservationPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"initiateCancellationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isCMAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCancellable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCancellationRejectionReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"rejectCancellationProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"safeMintWithReservation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isCancellable\",\"type\":\"bool\"}],\"name\":\"safeMintWithReservation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minExpirationTimestampDiff\",\"type\":\"uint256\"}],\"name\":\"setMinExpirationTimestampDiff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Bookingtokenv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bookingtokenv2MetaData.ABI instead.
var Bookingtokenv2ABI = Bookingtokenv2MetaData.ABI

// Bookingtokenv2 is an auto generated Go binding around an Ethereum contract.
type Bookingtokenv2 struct {
	Bookingtokenv2Caller     // Read-only binding to the contract
	Bookingtokenv2Transactor // Write-only binding to the contract
	Bookingtokenv2Filterer   // Log filterer for contract events
}

// Bookingtokenv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Bookingtokenv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bookingtokenv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Bookingtokenv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bookingtokenv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bookingtokenv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bookingtokenv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bookingtokenv2Session struct {
	Contract     *Bookingtokenv2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bookingtokenv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bookingtokenv2CallerSession struct {
	Contract *Bookingtokenv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Bookingtokenv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bookingtokenv2TransactorSession struct {
	Contract     *Bookingtokenv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Bookingtokenv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Bookingtokenv2Raw struct {
	Contract *Bookingtokenv2 // Generic contract binding to access the raw methods on
}

// Bookingtokenv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bookingtokenv2CallerRaw struct {
	Contract *Bookingtokenv2Caller // Generic read-only contract binding to access the raw methods on
}

// Bookingtokenv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bookingtokenv2TransactorRaw struct {
	Contract *Bookingtokenv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBookingtokenv2 creates a new instance of Bookingtokenv2, bound to a specific deployed contract.
func NewBookingtokenv2(address common.Address, backend bind.ContractBackend) (*Bookingtokenv2, error) {
	contract, err := bindBookingtokenv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2{Bookingtokenv2Caller: Bookingtokenv2Caller{contract: contract}, Bookingtokenv2Transactor: Bookingtokenv2Transactor{contract: contract}, Bookingtokenv2Filterer: Bookingtokenv2Filterer{contract: contract}}, nil
}

// NewBookingtokenv2Caller creates a new read-only instance of Bookingtokenv2, bound to a specific deployed contract.
func NewBookingtokenv2Caller(address common.Address, caller bind.ContractCaller) (*Bookingtokenv2Caller, error) {
	contract, err := bindBookingtokenv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2Caller{contract: contract}, nil
}

// NewBookingtokenv2Transactor creates a new write-only instance of Bookingtokenv2, bound to a specific deployed contract.
func NewBookingtokenv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Bookingtokenv2Transactor, error) {
	contract, err := bindBookingtokenv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2Transactor{contract: contract}, nil
}

// NewBookingtokenv2Filterer creates a new log filterer instance of Bookingtokenv2, bound to a specific deployed contract.
func NewBookingtokenv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Bookingtokenv2Filterer, error) {
	contract, err := bindBookingtokenv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2Filterer{contract: contract}, nil
}

// bindBookingtokenv2 binds a generic wrapper to an already deployed contract.
func bindBookingtokenv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Bookingtokenv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtokenv2 *Bookingtokenv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtokenv2.Contract.Bookingtokenv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtokenv2 *Bookingtokenv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.Bookingtokenv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtokenv2 *Bookingtokenv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.Bookingtokenv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtokenv2 *Bookingtokenv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtokenv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtokenv2 *Bookingtokenv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtokenv2 *Bookingtokenv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bookingtokenv2.Contract.DEFAULTADMINROLE(&_Bookingtokenv2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bookingtokenv2.Contract.DEFAULTADMINROLE(&_Bookingtokenv2.CallOpts)
}

// MINEXPIRATIONADMINROLE is a free data retrieval call binding the contract method 0x2edf5e2c.
//
// Solidity: function MIN_EXPIRATION_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Caller) MINEXPIRATIONADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "MIN_EXPIRATION_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINEXPIRATIONADMINROLE is a free data retrieval call binding the contract method 0x2edf5e2c.
//
// Solidity: function MIN_EXPIRATION_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Session) MINEXPIRATIONADMINROLE() ([32]byte, error) {
	return _Bookingtokenv2.Contract.MINEXPIRATIONADMINROLE(&_Bookingtokenv2.CallOpts)
}

// MINEXPIRATIONADMINROLE is a free data retrieval call binding the contract method 0x2edf5e2c.
//
// Solidity: function MIN_EXPIRATION_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) MINEXPIRATIONADMINROLE() ([32]byte, error) {
	return _Bookingtokenv2.Contract.MINEXPIRATIONADMINROLE(&_Bookingtokenv2.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Caller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Session) UPGRADERROLE() ([32]byte, error) {
	return _Bookingtokenv2.Contract.UPGRADERROLE(&_Bookingtokenv2.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Bookingtokenv2.Contract.UPGRADERROLE(&_Bookingtokenv2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bookingtokenv2.Contract.UPGRADEINTERFACEVERSION(&_Bookingtokenv2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bookingtokenv2.Contract.UPGRADEINTERFACEVERSION(&_Bookingtokenv2.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bookingtokenv2.Contract.BalanceOf(&_Bookingtokenv2.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bookingtokenv2.Contract.BalanceOf(&_Bookingtokenv2.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bookingtokenv2.Contract.GetApproved(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bookingtokenv2.Contract.GetApproved(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetBookingStatus is a free data retrieval call binding the contract method 0x3c15b31c.
//
// Solidity: function getBookingStatus(uint256 tokenId) view returns(uint8)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetBookingStatus(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getBookingStatus", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetBookingStatus is a free data retrieval call binding the contract method 0x3c15b31c.
//
// Solidity: function getBookingStatus(uint256 tokenId) view returns(uint8)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetBookingStatus(tokenId *big.Int) (uint8, error) {
	return _Bookingtokenv2.Contract.GetBookingStatus(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetBookingStatus is a free data retrieval call binding the contract method 0x3c15b31c.
//
// Solidity: function getBookingStatus(uint256 tokenId) view returns(uint8)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetBookingStatus(tokenId *big.Int) (uint8, error) {
	return _Bookingtokenv2.Contract.GetBookingStatus(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetCancellationProposalRefundAmount is a free data retrieval call binding the contract method 0xa101c67e.
//
// Solidity: function getCancellationProposalRefundAmount(uint256 tokenId) view returns(uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetCancellationProposalRefundAmount(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getCancellationProposalRefundAmount", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCancellationProposalRefundAmount is a free data retrieval call binding the contract method 0xa101c67e.
//
// Solidity: function getCancellationProposalRefundAmount(uint256 tokenId) view returns(uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetCancellationProposalRefundAmount(tokenId *big.Int) (*big.Int, error) {
	return _Bookingtokenv2.Contract.GetCancellationProposalRefundAmount(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetCancellationProposalRefundAmount is a free data retrieval call binding the contract method 0xa101c67e.
//
// Solidity: function getCancellationProposalRefundAmount(uint256 tokenId) view returns(uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetCancellationProposalRefundAmount(tokenId *big.Int) (*big.Int, error) {
	return _Bookingtokenv2.Contract.GetCancellationProposalRefundAmount(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetCancellationProposalStatus is a free data retrieval call binding the contract method 0x454d0db9.
//
// Solidity: function getCancellationProposalStatus(uint256 tokenId) view returns(uint256 refundAmount, address proposedBy, uint8 status, uint8 rejectionReason)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetCancellationProposalStatus(opts *bind.CallOpts, tokenId *big.Int) (struct {
	RefundAmount    *big.Int
	ProposedBy      common.Address
	Status          uint8
	RejectionReason uint8
}, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getCancellationProposalStatus", tokenId)

	outstruct := new(struct {
		RefundAmount    *big.Int
		ProposedBy      common.Address
		Status          uint8
		RejectionReason uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RefundAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProposedBy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.RejectionReason = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetCancellationProposalStatus is a free data retrieval call binding the contract method 0x454d0db9.
//
// Solidity: function getCancellationProposalStatus(uint256 tokenId) view returns(uint256 refundAmount, address proposedBy, uint8 status, uint8 rejectionReason)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetCancellationProposalStatus(tokenId *big.Int) (struct {
	RefundAmount    *big.Int
	ProposedBy      common.Address
	Status          uint8
	RejectionReason uint8
}, error) {
	return _Bookingtokenv2.Contract.GetCancellationProposalStatus(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetCancellationProposalStatus is a free data retrieval call binding the contract method 0x454d0db9.
//
// Solidity: function getCancellationProposalStatus(uint256 tokenId) view returns(uint256 refundAmount, address proposedBy, uint8 status, uint8 rejectionReason)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetCancellationProposalStatus(tokenId *big.Int) (struct {
	RefundAmount    *big.Int
	ProposedBy      common.Address
	Status          uint8
	RejectionReason uint8
}, error) {
	return _Bookingtokenv2.Contract.GetCancellationProposalStatus(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetManagerAddress() (common.Address, error) {
	return _Bookingtokenv2.Contract.GetManagerAddress(&_Bookingtokenv2.CallOpts)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetManagerAddress() (common.Address, error) {
	return _Bookingtokenv2.Contract.GetManagerAddress(&_Bookingtokenv2.CallOpts)
}

// GetMinExpirationTimestampDiff is a free data retrieval call binding the contract method 0x0e75c1a8.
//
// Solidity: function getMinExpirationTimestampDiff() view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetMinExpirationTimestampDiff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getMinExpirationTimestampDiff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinExpirationTimestampDiff is a free data retrieval call binding the contract method 0x0e75c1a8.
//
// Solidity: function getMinExpirationTimestampDiff() view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetMinExpirationTimestampDiff() (*big.Int, error) {
	return _Bookingtokenv2.Contract.GetMinExpirationTimestampDiff(&_Bookingtokenv2.CallOpts)
}

// GetMinExpirationTimestampDiff is a free data retrieval call binding the contract method 0x0e75c1a8.
//
// Solidity: function getMinExpirationTimestampDiff() view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetMinExpirationTimestampDiff() (*big.Int, error) {
	return _Bookingtokenv2.Contract.GetMinExpirationTimestampDiff(&_Bookingtokenv2.CallOpts)
}

// GetReservationPaymentToken is a free data retrieval call binding the contract method 0xb191d092.
//
// Solidity: function getReservationPaymentToken(uint256 tokenId) view returns(address paymentToken)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetReservationPaymentToken(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getReservationPaymentToken", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReservationPaymentToken is a free data retrieval call binding the contract method 0xb191d092.
//
// Solidity: function getReservationPaymentToken(uint256 tokenId) view returns(address paymentToken)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetReservationPaymentToken(tokenId *big.Int) (common.Address, error) {
	return _Bookingtokenv2.Contract.GetReservationPaymentToken(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetReservationPaymentToken is a free data retrieval call binding the contract method 0xb191d092.
//
// Solidity: function getReservationPaymentToken(uint256 tokenId) view returns(address paymentToken)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetReservationPaymentToken(tokenId *big.Int) (common.Address, error) {
	return _Bookingtokenv2.Contract.GetReservationPaymentToken(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetReservationPrice is a free data retrieval call binding the contract method 0x004fdd3c.
//
// Solidity: function getReservationPrice(uint256 tokenId) view returns(uint256 price, address paymentToken)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetReservationPrice(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Price        *big.Int
	PaymentToken common.Address
}, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getReservationPrice", tokenId)

	outstruct := new(struct {
		Price        *big.Int
		PaymentToken common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PaymentToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetReservationPrice is a free data retrieval call binding the contract method 0x004fdd3c.
//
// Solidity: function getReservationPrice(uint256 tokenId) view returns(uint256 price, address paymentToken)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetReservationPrice(tokenId *big.Int) (struct {
	Price        *big.Int
	PaymentToken common.Address
}, error) {
	return _Bookingtokenv2.Contract.GetReservationPrice(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetReservationPrice is a free data retrieval call binding the contract method 0x004fdd3c.
//
// Solidity: function getReservationPrice(uint256 tokenId) view returns(uint256 price, address paymentToken)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetReservationPrice(tokenId *big.Int) (struct {
	Price        *big.Int
	PaymentToken common.Address
}, error) {
	return _Bookingtokenv2.Contract.GetReservationPrice(&_Bookingtokenv2.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bookingtokenv2.Contract.GetRoleAdmin(&_Bookingtokenv2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bookingtokenv2.Contract.GetRoleAdmin(&_Bookingtokenv2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bookingtokenv2.Contract.HasRole(&_Bookingtokenv2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bookingtokenv2.Contract.HasRole(&_Bookingtokenv2.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bookingtokenv2.Contract.IsApprovedForAll(&_Bookingtokenv2.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bookingtokenv2.Contract.IsApprovedForAll(&_Bookingtokenv2.CallOpts, owner, operator)
}

// IsCMAccount is a free data retrieval call binding the contract method 0x12b357b5.
//
// Solidity: function isCMAccount(address account) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Caller) IsCMAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "isCMAccount", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCMAccount is a free data retrieval call binding the contract method 0x12b357b5.
//
// Solidity: function isCMAccount(address account) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Session) IsCMAccount(account common.Address) (bool, error) {
	return _Bookingtokenv2.Contract.IsCMAccount(&_Bookingtokenv2.CallOpts, account)
}

// IsCMAccount is a free data retrieval call binding the contract method 0x12b357b5.
//
// Solidity: function isCMAccount(address account) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) IsCMAccount(account common.Address) (bool, error) {
	return _Bookingtokenv2.Contract.IsCMAccount(&_Bookingtokenv2.CallOpts, account)
}

// IsCancellable is a free data retrieval call binding the contract method 0x2d3a6329.
//
// Solidity: function isCancellable(uint256 tokenId) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Caller) IsCancellable(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "isCancellable", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCancellable is a free data retrieval call binding the contract method 0x2d3a6329.
//
// Solidity: function isCancellable(uint256 tokenId) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Session) IsCancellable(tokenId *big.Int) (bool, error) {
	return _Bookingtokenv2.Contract.IsCancellable(&_Bookingtokenv2.CallOpts, tokenId)
}

// IsCancellable is a free data retrieval call binding the contract method 0x2d3a6329.
//
// Solidity: function isCancellable(uint256 tokenId) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) IsCancellable(tokenId *big.Int) (bool, error) {
	return _Bookingtokenv2.Contract.IsCancellable(&_Bookingtokenv2.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Session) Name() (string, error) {
	return _Bookingtokenv2.Contract.Name(&_Bookingtokenv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) Name() (string, error) {
	return _Bookingtokenv2.Contract.Name(&_Bookingtokenv2.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bookingtokenv2.Contract.OwnerOf(&_Bookingtokenv2.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bookingtokenv2.Contract.OwnerOf(&_Bookingtokenv2.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2Session) ProxiableUUID() ([32]byte, error) {
	return _Bookingtokenv2.Contract.ProxiableUUID(&_Bookingtokenv2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bookingtokenv2.Contract.ProxiableUUID(&_Bookingtokenv2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bookingtokenv2.Contract.SupportsInterface(&_Bookingtokenv2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bookingtokenv2.Contract.SupportsInterface(&_Bookingtokenv2.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Session) Symbol() (string, error) {
	return _Bookingtokenv2.Contract.Symbol(&_Bookingtokenv2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) Symbol() (string, error) {
	return _Bookingtokenv2.Contract.Symbol(&_Bookingtokenv2.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Bookingtokenv2.Contract.TokenByIndex(&_Bookingtokenv2.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Bookingtokenv2.Contract.TokenByIndex(&_Bookingtokenv2.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Bookingtokenv2.Contract.TokenOfOwnerByIndex(&_Bookingtokenv2.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Bookingtokenv2.Contract.TokenOfOwnerByIndex(&_Bookingtokenv2.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Bookingtokenv2.Contract.TokenURI(&_Bookingtokenv2.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Bookingtokenv2.Contract.TokenURI(&_Bookingtokenv2.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtokenv2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2Session) TotalSupply() (*big.Int, error) {
	return _Bookingtokenv2.Contract.TotalSupply(&_Bookingtokenv2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bookingtokenv2 *Bookingtokenv2CallerSession) TotalSupply() (*big.Int, error) {
	return _Bookingtokenv2.Contract.TotalSupply(&_Bookingtokenv2.CallOpts)
}

// AcceptCancellationProposal is a paid mutator transaction binding the contract method 0xde62fe4d.
//
// Solidity: function acceptCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) AcceptCancellationProposal(opts *bind.TransactOpts, tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "acceptCancellationProposal", tokenId, checkRefundAmount)
}

// AcceptCancellationProposal is a paid mutator transaction binding the contract method 0xde62fe4d.
//
// Solidity: function acceptCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) AcceptCancellationProposal(tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.AcceptCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, checkRefundAmount)
}

// AcceptCancellationProposal is a paid mutator transaction binding the contract method 0xde62fe4d.
//
// Solidity: function acceptCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) AcceptCancellationProposal(tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.AcceptCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, checkRefundAmount)
}

// AcceptCounteredCancellationProposal is a paid mutator transaction binding the contract method 0xf024df20.
//
// Solidity: function acceptCounteredCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) AcceptCounteredCancellationProposal(opts *bind.TransactOpts, tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "acceptCounteredCancellationProposal", tokenId, checkRefundAmount)
}

// AcceptCounteredCancellationProposal is a paid mutator transaction binding the contract method 0xf024df20.
//
// Solidity: function acceptCounteredCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) AcceptCounteredCancellationProposal(tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.AcceptCounteredCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, checkRefundAmount)
}

// AcceptCounteredCancellationProposal is a paid mutator transaction binding the contract method 0xf024df20.
//
// Solidity: function acceptCounteredCancellationProposal(uint256 tokenId, uint256 checkRefundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) AcceptCounteredCancellationProposal(tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.AcceptCounteredCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, checkRefundAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.Approve(&_Bookingtokenv2.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.Approve(&_Bookingtokenv2.TransactOpts, to, tokenId)
}

// BuyReservedToken is a paid mutator transaction binding the contract method 0x96591edd.
//
// Solidity: function buyReservedToken(uint256 tokenId) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) BuyReservedToken(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "buyReservedToken", tokenId)
}

// BuyReservedToken is a paid mutator transaction binding the contract method 0x96591edd.
//
// Solidity: function buyReservedToken(uint256 tokenId) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) BuyReservedToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.BuyReservedToken(&_Bookingtokenv2.TransactOpts, tokenId)
}

// BuyReservedToken is a paid mutator transaction binding the contract method 0x96591edd.
//
// Solidity: function buyReservedToken(uint256 tokenId) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) BuyReservedToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.BuyReservedToken(&_Bookingtokenv2.TransactOpts, tokenId)
}

// CancelCancellationProposal is a paid mutator transaction binding the contract method 0x6088f4fd.
//
// Solidity: function cancelCancellationProposal(uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) CancelCancellationProposal(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "cancelCancellationProposal", tokenId)
}

// CancelCancellationProposal is a paid mutator transaction binding the contract method 0x6088f4fd.
//
// Solidity: function cancelCancellationProposal(uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) CancelCancellationProposal(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.CancelCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId)
}

// CancelCancellationProposal is a paid mutator transaction binding the contract method 0x6088f4fd.
//
// Solidity: function cancelCancellationProposal(uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) CancelCancellationProposal(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.CancelCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId)
}

// CounterCancellationProposal is a paid mutator transaction binding the contract method 0x2e0258a0.
//
// Solidity: function counterCancellationProposal(uint256 tokenId, uint256 newRefundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) CounterCancellationProposal(opts *bind.TransactOpts, tokenId *big.Int, newRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "counterCancellationProposal", tokenId, newRefundAmount)
}

// CounterCancellationProposal is a paid mutator transaction binding the contract method 0x2e0258a0.
//
// Solidity: function counterCancellationProposal(uint256 tokenId, uint256 newRefundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) CounterCancellationProposal(tokenId *big.Int, newRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.CounterCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, newRefundAmount)
}

// CounterCancellationProposal is a paid mutator transaction binding the contract method 0x2e0258a0.
//
// Solidity: function counterCancellationProposal(uint256 tokenId, uint256 newRefundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) CounterCancellationProposal(tokenId *big.Int, newRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.CounterCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, newRefundAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.GrantRole(&_Bookingtokenv2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.GrantRole(&_Bookingtokenv2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address manager, address defaultAdmin, address upgrader) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) Initialize(opts *bind.TransactOpts, manager common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "initialize", manager, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address manager, address defaultAdmin, address upgrader) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) Initialize(manager common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.Initialize(&_Bookingtokenv2.TransactOpts, manager, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address manager, address defaultAdmin, address upgrader) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) Initialize(manager common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.Initialize(&_Bookingtokenv2.TransactOpts, manager, defaultAdmin, upgrader)
}

// InitiateCancellationProposal is a paid mutator transaction binding the contract method 0x09b35e00.
//
// Solidity: function initiateCancellationProposal(uint256 tokenId, uint256 refundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) InitiateCancellationProposal(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "initiateCancellationProposal", tokenId, refundAmount)
}

// InitiateCancellationProposal is a paid mutator transaction binding the contract method 0x09b35e00.
//
// Solidity: function initiateCancellationProposal(uint256 tokenId, uint256 refundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) InitiateCancellationProposal(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.InitiateCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, refundAmount)
}

// InitiateCancellationProposal is a paid mutator transaction binding the contract method 0x09b35e00.
//
// Solidity: function initiateCancellationProposal(uint256 tokenId, uint256 refundAmount) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) InitiateCancellationProposal(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.InitiateCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, refundAmount)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) RecordExpiration(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "recordExpiration", tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RecordExpiration(&_Bookingtokenv2.TransactOpts, tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RecordExpiration(&_Bookingtokenv2.TransactOpts, tokenId)
}

// RejectCancellationProposal is a paid mutator transaction binding the contract method 0xc14239f1.
//
// Solidity: function rejectCancellationProposal(uint256 tokenId, uint8 reason) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) RejectCancellationProposal(opts *bind.TransactOpts, tokenId *big.Int, reason uint8) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "rejectCancellationProposal", tokenId, reason)
}

// RejectCancellationProposal is a paid mutator transaction binding the contract method 0xc14239f1.
//
// Solidity: function rejectCancellationProposal(uint256 tokenId, uint8 reason) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) RejectCancellationProposal(tokenId *big.Int, reason uint8) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RejectCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, reason)
}

// RejectCancellationProposal is a paid mutator transaction binding the contract method 0xc14239f1.
//
// Solidity: function rejectCancellationProposal(uint256 tokenId, uint8 reason) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) RejectCancellationProposal(tokenId *big.Int, reason uint8) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RejectCancellationProposal(&_Bookingtokenv2.TransactOpts, tokenId, reason)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RenounceRole(&_Bookingtokenv2.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RenounceRole(&_Bookingtokenv2.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RevokeRole(&_Bookingtokenv2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.RevokeRole(&_Bookingtokenv2.TransactOpts, role, account)
}

// SafeMintWithReservation is a paid mutator transaction binding the contract method 0x5d4badb2.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) SafeMintWithReservation(opts *bind.TransactOpts, reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "safeMintWithReservation", reservedFor, uri, expirationTimestamp, price, paymentToken)
}

// SafeMintWithReservation is a paid mutator transaction binding the contract method 0x5d4badb2.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) SafeMintWithReservation(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeMintWithReservation(&_Bookingtokenv2.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken)
}

// SafeMintWithReservation is a paid mutator transaction binding the contract method 0x5d4badb2.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) SafeMintWithReservation(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeMintWithReservation(&_Bookingtokenv2.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken)
}

// SafeMintWithReservation0 is a paid mutator transaction binding the contract method 0x8c9c3c12.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, bool _isCancellable) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) SafeMintWithReservation0(opts *bind.TransactOpts, reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, _isCancellable bool) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "safeMintWithReservation0", reservedFor, uri, expirationTimestamp, price, paymentToken, _isCancellable)
}

// SafeMintWithReservation0 is a paid mutator transaction binding the contract method 0x8c9c3c12.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, bool _isCancellable) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) SafeMintWithReservation0(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, _isCancellable bool) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeMintWithReservation0(&_Bookingtokenv2.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, _isCancellable)
}

// SafeMintWithReservation0 is a paid mutator transaction binding the contract method 0x8c9c3c12.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, bool _isCancellable) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) SafeMintWithReservation0(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, _isCancellable bool) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeMintWithReservation0(&_Bookingtokenv2.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, _isCancellable)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeTransferFrom(&_Bookingtokenv2.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeTransferFrom(&_Bookingtokenv2.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeTransferFrom0(&_Bookingtokenv2.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SafeTransferFrom0(&_Bookingtokenv2.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SetApprovalForAll(&_Bookingtokenv2.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SetApprovalForAll(&_Bookingtokenv2.TransactOpts, operator, approved)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address manager) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) SetManagerAddress(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "setManagerAddress", manager)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address manager) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) SetManagerAddress(manager common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SetManagerAddress(&_Bookingtokenv2.TransactOpts, manager)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address manager) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) SetManagerAddress(manager common.Address) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SetManagerAddress(&_Bookingtokenv2.TransactOpts, manager)
}

// SetMinExpirationTimestampDiff is a paid mutator transaction binding the contract method 0x516a82b8.
//
// Solidity: function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) SetMinExpirationTimestampDiff(opts *bind.TransactOpts, minExpirationTimestampDiff *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "setMinExpirationTimestampDiff", minExpirationTimestampDiff)
}

// SetMinExpirationTimestampDiff is a paid mutator transaction binding the contract method 0x516a82b8.
//
// Solidity: function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) SetMinExpirationTimestampDiff(minExpirationTimestampDiff *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SetMinExpirationTimestampDiff(&_Bookingtokenv2.TransactOpts, minExpirationTimestampDiff)
}

// SetMinExpirationTimestampDiff is a paid mutator transaction binding the contract method 0x516a82b8.
//
// Solidity: function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) SetMinExpirationTimestampDiff(minExpirationTimestampDiff *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.SetMinExpirationTimestampDiff(&_Bookingtokenv2.TransactOpts, minExpirationTimestampDiff)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.TransferFrom(&_Bookingtokenv2.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.TransferFrom(&_Bookingtokenv2.TransactOpts, from, to, tokenId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bookingtokenv2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.UpgradeToAndCall(&_Bookingtokenv2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bookingtokenv2 *Bookingtokenv2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bookingtokenv2.Contract.UpgradeToAndCall(&_Bookingtokenv2.TransactOpts, newImplementation, data)
}

// Bookingtokenv2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bookingtokenv2 contract.
type Bookingtokenv2ApprovalIterator struct {
	Event *Bookingtokenv2Approval // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2Approval)
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
		it.Event = new(Bookingtokenv2Approval)
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
func (it *Bookingtokenv2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2Approval represents a Approval event raised by the Bookingtokenv2 contract.
type Bookingtokenv2Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Bookingtokenv2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2ApprovalIterator{contract: _Bookingtokenv2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2Approval)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseApproval(log types.Log) (*Bookingtokenv2Approval, error) {
	event := new(Bookingtokenv2Approval)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Bookingtokenv2 contract.
type Bookingtokenv2ApprovalForAllIterator struct {
	Event *Bookingtokenv2ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2ApprovalForAll)
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
		it.Event = new(Bookingtokenv2ApprovalForAll)
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
func (it *Bookingtokenv2ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2ApprovalForAll represents a ApprovalForAll event raised by the Bookingtokenv2 contract.
type Bookingtokenv2ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Bookingtokenv2ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2ApprovalForAllIterator{contract: _Bookingtokenv2.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2ApprovalForAll)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseApprovalForAll(log types.Log) (*Bookingtokenv2ApprovalForAll, error) {
	event := new(Bookingtokenv2ApprovalForAll)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Bookingtokenv2 contract.
type Bookingtokenv2BatchMetadataUpdateIterator struct {
	Event *Bookingtokenv2BatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2BatchMetadataUpdate)
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
		it.Event = new(Bookingtokenv2BatchMetadataUpdate)
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
func (it *Bookingtokenv2BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Bookingtokenv2 contract.
type Bookingtokenv2BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*Bookingtokenv2BatchMetadataUpdateIterator, error) {

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2BatchMetadataUpdateIterator{contract: _Bookingtokenv2.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2BatchMetadataUpdate)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseBatchMetadataUpdate(log types.Log) (*Bookingtokenv2BatchMetadataUpdate, error) {
	event := new(Bookingtokenv2BatchMetadataUpdate)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2CancellationAcceptedIterator is returned from FilterCancellationAccepted and is used to iterate over the raw logs and unpacked data for CancellationAccepted events raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationAcceptedIterator struct {
	Event *Bookingtokenv2CancellationAccepted // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2CancellationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2CancellationAccepted)
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
		it.Event = new(Bookingtokenv2CancellationAccepted)
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
func (it *Bookingtokenv2CancellationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2CancellationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2CancellationAccepted represents a CancellationAccepted event raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationAccepted struct {
	TokenId      *big.Int
	AcceptedBy   common.Address
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCancellationAccepted is a free log retrieval operation binding the contract event 0xdf11499efc7ab0feb9befa7a615c79d3df759d9930b41f31c1f0723cfd9b10f9.
//
// Solidity: event CancellationAccepted(uint256 indexed tokenId, address indexed acceptedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterCancellationAccepted(opts *bind.FilterOpts, tokenId []*big.Int, acceptedBy []common.Address) (*Bookingtokenv2CancellationAcceptedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var acceptedByRule []interface{}
	for _, acceptedByItem := range acceptedBy {
		acceptedByRule = append(acceptedByRule, acceptedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "CancellationAccepted", tokenIdRule, acceptedByRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2CancellationAcceptedIterator{contract: _Bookingtokenv2.contract, event: "CancellationAccepted", logs: logs, sub: sub}, nil
}

// WatchCancellationAccepted is a free log subscription operation binding the contract event 0xdf11499efc7ab0feb9befa7a615c79d3df759d9930b41f31c1f0723cfd9b10f9.
//
// Solidity: event CancellationAccepted(uint256 indexed tokenId, address indexed acceptedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchCancellationAccepted(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2CancellationAccepted, tokenId []*big.Int, acceptedBy []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var acceptedByRule []interface{}
	for _, acceptedByItem := range acceptedBy {
		acceptedByRule = append(acceptedByRule, acceptedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "CancellationAccepted", tokenIdRule, acceptedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2CancellationAccepted)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationAccepted", log); err != nil {
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

// ParseCancellationAccepted is a log parse operation binding the contract event 0xdf11499efc7ab0feb9befa7a615c79d3df759d9930b41f31c1f0723cfd9b10f9.
//
// Solidity: event CancellationAccepted(uint256 indexed tokenId, address indexed acceptedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseCancellationAccepted(log types.Log) (*Bookingtokenv2CancellationAccepted, error) {
	event := new(Bookingtokenv2CancellationAccepted)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2CancellationCounteredIterator is returned from FilterCancellationCountered and is used to iterate over the raw logs and unpacked data for CancellationCountered events raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationCounteredIterator struct {
	Event *Bookingtokenv2CancellationCountered // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2CancellationCounteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2CancellationCountered)
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
		it.Event = new(Bookingtokenv2CancellationCountered)
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
func (it *Bookingtokenv2CancellationCounteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2CancellationCounteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2CancellationCountered represents a CancellationCountered event raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationCountered struct {
	TokenId         *big.Int
	CounteredBy     common.Address
	NewRefundAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancellationCountered is a free log retrieval operation binding the contract event 0x7b6cf4b0eeba58e59d225cabf115102b6f1f5af0515f5f6bb6ec9709bc854d09.
//
// Solidity: event CancellationCountered(uint256 indexed tokenId, address indexed counteredBy, uint256 newRefundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterCancellationCountered(opts *bind.FilterOpts, tokenId []*big.Int, counteredBy []common.Address) (*Bookingtokenv2CancellationCounteredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var counteredByRule []interface{}
	for _, counteredByItem := range counteredBy {
		counteredByRule = append(counteredByRule, counteredByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "CancellationCountered", tokenIdRule, counteredByRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2CancellationCounteredIterator{contract: _Bookingtokenv2.contract, event: "CancellationCountered", logs: logs, sub: sub}, nil
}

// WatchCancellationCountered is a free log subscription operation binding the contract event 0x7b6cf4b0eeba58e59d225cabf115102b6f1f5af0515f5f6bb6ec9709bc854d09.
//
// Solidity: event CancellationCountered(uint256 indexed tokenId, address indexed counteredBy, uint256 newRefundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchCancellationCountered(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2CancellationCountered, tokenId []*big.Int, counteredBy []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var counteredByRule []interface{}
	for _, counteredByItem := range counteredBy {
		counteredByRule = append(counteredByRule, counteredByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "CancellationCountered", tokenIdRule, counteredByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2CancellationCountered)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationCountered", log); err != nil {
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

// ParseCancellationCountered is a log parse operation binding the contract event 0x7b6cf4b0eeba58e59d225cabf115102b6f1f5af0515f5f6bb6ec9709bc854d09.
//
// Solidity: event CancellationCountered(uint256 indexed tokenId, address indexed counteredBy, uint256 newRefundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseCancellationCountered(log types.Log) (*Bookingtokenv2CancellationCountered, error) {
	event := new(Bookingtokenv2CancellationCountered)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationCountered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2CancellationPendingIterator is returned from FilterCancellationPending and is used to iterate over the raw logs and unpacked data for CancellationPending events raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationPendingIterator struct {
	Event *Bookingtokenv2CancellationPending // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2CancellationPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2CancellationPending)
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
		it.Event = new(Bookingtokenv2CancellationPending)
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
func (it *Bookingtokenv2CancellationPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2CancellationPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2CancellationPending represents a CancellationPending event raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationPending struct {
	TokenId      *big.Int
	ProposedBy   common.Address
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCancellationPending is a free log retrieval operation binding the contract event 0x53c665da5aadaba077915db08a54c7222d7b72603f945ccf87ace2212772e1df.
//
// Solidity: event CancellationPending(uint256 indexed tokenId, address indexed proposedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterCancellationPending(opts *bind.FilterOpts, tokenId []*big.Int, proposedBy []common.Address) (*Bookingtokenv2CancellationPendingIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var proposedByRule []interface{}
	for _, proposedByItem := range proposedBy {
		proposedByRule = append(proposedByRule, proposedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "CancellationPending", tokenIdRule, proposedByRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2CancellationPendingIterator{contract: _Bookingtokenv2.contract, event: "CancellationPending", logs: logs, sub: sub}, nil
}

// WatchCancellationPending is a free log subscription operation binding the contract event 0x53c665da5aadaba077915db08a54c7222d7b72603f945ccf87ace2212772e1df.
//
// Solidity: event CancellationPending(uint256 indexed tokenId, address indexed proposedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchCancellationPending(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2CancellationPending, tokenId []*big.Int, proposedBy []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var proposedByRule []interface{}
	for _, proposedByItem := range proposedBy {
		proposedByRule = append(proposedByRule, proposedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "CancellationPending", tokenIdRule, proposedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2CancellationPending)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationPending", log); err != nil {
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

// ParseCancellationPending is a log parse operation binding the contract event 0x53c665da5aadaba077915db08a54c7222d7b72603f945ccf87ace2212772e1df.
//
// Solidity: event CancellationPending(uint256 indexed tokenId, address indexed proposedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseCancellationPending(log types.Log) (*Bookingtokenv2CancellationPending, error) {
	event := new(Bookingtokenv2CancellationPending)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2CancellationProposalAcceptedByTheOwnerIterator is returned from FilterCancellationProposalAcceptedByTheOwner and is used to iterate over the raw logs and unpacked data for CancellationProposalAcceptedByTheOwner events raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationProposalAcceptedByTheOwnerIterator struct {
	Event *Bookingtokenv2CancellationProposalAcceptedByTheOwner // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2CancellationProposalAcceptedByTheOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2CancellationProposalAcceptedByTheOwner)
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
		it.Event = new(Bookingtokenv2CancellationProposalAcceptedByTheOwner)
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
func (it *Bookingtokenv2CancellationProposalAcceptedByTheOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2CancellationProposalAcceptedByTheOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2CancellationProposalAcceptedByTheOwner represents a CancellationProposalAcceptedByTheOwner event raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationProposalAcceptedByTheOwner struct {
	TokenId      *big.Int
	AcceptedBy   common.Address
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCancellationProposalAcceptedByTheOwner is a free log retrieval operation binding the contract event 0x8262f1b4e879bd89552bbc702dc20d0b0552f0a5fe8292e9bf4a3700492e9d98.
//
// Solidity: event CancellationProposalAcceptedByTheOwner(uint256 indexed tokenId, address indexed acceptedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterCancellationProposalAcceptedByTheOwner(opts *bind.FilterOpts, tokenId []*big.Int, acceptedBy []common.Address) (*Bookingtokenv2CancellationProposalAcceptedByTheOwnerIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var acceptedByRule []interface{}
	for _, acceptedByItem := range acceptedBy {
		acceptedByRule = append(acceptedByRule, acceptedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "CancellationProposalAcceptedByTheOwner", tokenIdRule, acceptedByRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2CancellationProposalAcceptedByTheOwnerIterator{contract: _Bookingtokenv2.contract, event: "CancellationProposalAcceptedByTheOwner", logs: logs, sub: sub}, nil
}

// WatchCancellationProposalAcceptedByTheOwner is a free log subscription operation binding the contract event 0x8262f1b4e879bd89552bbc702dc20d0b0552f0a5fe8292e9bf4a3700492e9d98.
//
// Solidity: event CancellationProposalAcceptedByTheOwner(uint256 indexed tokenId, address indexed acceptedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchCancellationProposalAcceptedByTheOwner(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2CancellationProposalAcceptedByTheOwner, tokenId []*big.Int, acceptedBy []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var acceptedByRule []interface{}
	for _, acceptedByItem := range acceptedBy {
		acceptedByRule = append(acceptedByRule, acceptedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "CancellationProposalAcceptedByTheOwner", tokenIdRule, acceptedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2CancellationProposalAcceptedByTheOwner)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationProposalAcceptedByTheOwner", log); err != nil {
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

// ParseCancellationProposalAcceptedByTheOwner is a log parse operation binding the contract event 0x8262f1b4e879bd89552bbc702dc20d0b0552f0a5fe8292e9bf4a3700492e9d98.
//
// Solidity: event CancellationProposalAcceptedByTheOwner(uint256 indexed tokenId, address indexed acceptedBy, uint256 refundAmount)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseCancellationProposalAcceptedByTheOwner(log types.Log) (*Bookingtokenv2CancellationProposalAcceptedByTheOwner, error) {
	event := new(Bookingtokenv2CancellationProposalAcceptedByTheOwner)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationProposalAcceptedByTheOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2CancellationProposalCancelledIterator is returned from FilterCancellationProposalCancelled and is used to iterate over the raw logs and unpacked data for CancellationProposalCancelled events raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationProposalCancelledIterator struct {
	Event *Bookingtokenv2CancellationProposalCancelled // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2CancellationProposalCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2CancellationProposalCancelled)
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
		it.Event = new(Bookingtokenv2CancellationProposalCancelled)
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
func (it *Bookingtokenv2CancellationProposalCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2CancellationProposalCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2CancellationProposalCancelled represents a CancellationProposalCancelled event raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationProposalCancelled struct {
	TokenId     *big.Int
	CancelledBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancellationProposalCancelled is a free log retrieval operation binding the contract event 0x76733a4d6a5b40818eb25829d380052f1de6913f3b4b77517363fc38bf23b897.
//
// Solidity: event CancellationProposalCancelled(uint256 indexed tokenId, address indexed cancelledBy)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterCancellationProposalCancelled(opts *bind.FilterOpts, tokenId []*big.Int, cancelledBy []common.Address) (*Bookingtokenv2CancellationProposalCancelledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "CancellationProposalCancelled", tokenIdRule, cancelledByRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2CancellationProposalCancelledIterator{contract: _Bookingtokenv2.contract, event: "CancellationProposalCancelled", logs: logs, sub: sub}, nil
}

// WatchCancellationProposalCancelled is a free log subscription operation binding the contract event 0x76733a4d6a5b40818eb25829d380052f1de6913f3b4b77517363fc38bf23b897.
//
// Solidity: event CancellationProposalCancelled(uint256 indexed tokenId, address indexed cancelledBy)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchCancellationProposalCancelled(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2CancellationProposalCancelled, tokenId []*big.Int, cancelledBy []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var cancelledByRule []interface{}
	for _, cancelledByItem := range cancelledBy {
		cancelledByRule = append(cancelledByRule, cancelledByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "CancellationProposalCancelled", tokenIdRule, cancelledByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2CancellationProposalCancelled)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationProposalCancelled", log); err != nil {
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

// ParseCancellationProposalCancelled is a log parse operation binding the contract event 0x76733a4d6a5b40818eb25829d380052f1de6913f3b4b77517363fc38bf23b897.
//
// Solidity: event CancellationProposalCancelled(uint256 indexed tokenId, address indexed cancelledBy)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseCancellationProposalCancelled(log types.Log) (*Bookingtokenv2CancellationProposalCancelled, error) {
	event := new(Bookingtokenv2CancellationProposalCancelled)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationProposalCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2CancellationRejectedIterator is returned from FilterCancellationRejected and is used to iterate over the raw logs and unpacked data for CancellationRejected events raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationRejectedIterator struct {
	Event *Bookingtokenv2CancellationRejected // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2CancellationRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2CancellationRejected)
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
		it.Event = new(Bookingtokenv2CancellationRejected)
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
func (it *Bookingtokenv2CancellationRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2CancellationRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2CancellationRejected represents a CancellationRejected event raised by the Bookingtokenv2 contract.
type Bookingtokenv2CancellationRejected struct {
	TokenId    *big.Int
	RejectedBy common.Address
	Reason     uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancellationRejected is a free log retrieval operation binding the contract event 0x378d7694f95db72d4f3dadbeef7aa7254d740e8f9a8382a0895848b87677738b.
//
// Solidity: event CancellationRejected(uint256 indexed tokenId, address indexed rejectedBy, uint8 reason)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterCancellationRejected(opts *bind.FilterOpts, tokenId []*big.Int, rejectedBy []common.Address) (*Bookingtokenv2CancellationRejectedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var rejectedByRule []interface{}
	for _, rejectedByItem := range rejectedBy {
		rejectedByRule = append(rejectedByRule, rejectedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "CancellationRejected", tokenIdRule, rejectedByRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2CancellationRejectedIterator{contract: _Bookingtokenv2.contract, event: "CancellationRejected", logs: logs, sub: sub}, nil
}

// WatchCancellationRejected is a free log subscription operation binding the contract event 0x378d7694f95db72d4f3dadbeef7aa7254d740e8f9a8382a0895848b87677738b.
//
// Solidity: event CancellationRejected(uint256 indexed tokenId, address indexed rejectedBy, uint8 reason)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchCancellationRejected(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2CancellationRejected, tokenId []*big.Int, rejectedBy []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var rejectedByRule []interface{}
	for _, rejectedByItem := range rejectedBy {
		rejectedByRule = append(rejectedByRule, rejectedByItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "CancellationRejected", tokenIdRule, rejectedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2CancellationRejected)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationRejected", log); err != nil {
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

// ParseCancellationRejected is a log parse operation binding the contract event 0x378d7694f95db72d4f3dadbeef7aa7254d740e8f9a8382a0895848b87677738b.
//
// Solidity: event CancellationRejected(uint256 indexed tokenId, address indexed rejectedBy, uint8 reason)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseCancellationRejected(log types.Log) (*Bookingtokenv2CancellationRejected, error) {
	event := new(Bookingtokenv2CancellationRejected)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "CancellationRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bookingtokenv2 contract.
type Bookingtokenv2InitializedIterator struct {
	Event *Bookingtokenv2Initialized // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2Initialized)
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
		it.Event = new(Bookingtokenv2Initialized)
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
func (it *Bookingtokenv2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2Initialized represents a Initialized event raised by the Bookingtokenv2 contract.
type Bookingtokenv2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Bookingtokenv2InitializedIterator, error) {

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2InitializedIterator{contract: _Bookingtokenv2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2Initialized) (event.Subscription, error) {

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2Initialized)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseInitialized(log types.Log) (*Bookingtokenv2Initialized, error) {
	event := new(Bookingtokenv2Initialized)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Bookingtokenv2 contract.
type Bookingtokenv2MetadataUpdateIterator struct {
	Event *Bookingtokenv2MetadataUpdate // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2MetadataUpdate)
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
		it.Event = new(Bookingtokenv2MetadataUpdate)
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
func (it *Bookingtokenv2MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2MetadataUpdate represents a MetadataUpdate event raised by the Bookingtokenv2 contract.
type Bookingtokenv2MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*Bookingtokenv2MetadataUpdateIterator, error) {

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2MetadataUpdateIterator{contract: _Bookingtokenv2.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2MetadataUpdate)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseMetadataUpdate(log types.Log) (*Bookingtokenv2MetadataUpdate, error) {
	event := new(Bookingtokenv2MetadataUpdate)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bookingtokenv2 contract.
type Bookingtokenv2RoleAdminChangedIterator struct {
	Event *Bookingtokenv2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2RoleAdminChanged)
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
		it.Event = new(Bookingtokenv2RoleAdminChanged)
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
func (it *Bookingtokenv2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2RoleAdminChanged represents a RoleAdminChanged event raised by the Bookingtokenv2 contract.
type Bookingtokenv2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Bookingtokenv2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2RoleAdminChangedIterator{contract: _Bookingtokenv2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2RoleAdminChanged)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseRoleAdminChanged(log types.Log) (*Bookingtokenv2RoleAdminChanged, error) {
	event := new(Bookingtokenv2RoleAdminChanged)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bookingtokenv2 contract.
type Bookingtokenv2RoleGrantedIterator struct {
	Event *Bookingtokenv2RoleGranted // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2RoleGranted)
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
		it.Event = new(Bookingtokenv2RoleGranted)
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
func (it *Bookingtokenv2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2RoleGranted represents a RoleGranted event raised by the Bookingtokenv2 contract.
type Bookingtokenv2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Bookingtokenv2RoleGrantedIterator, error) {

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

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2RoleGrantedIterator{contract: _Bookingtokenv2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2RoleGranted)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseRoleGranted(log types.Log) (*Bookingtokenv2RoleGranted, error) {
	event := new(Bookingtokenv2RoleGranted)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bookingtokenv2 contract.
type Bookingtokenv2RoleRevokedIterator struct {
	Event *Bookingtokenv2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2RoleRevoked)
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
		it.Event = new(Bookingtokenv2RoleRevoked)
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
func (it *Bookingtokenv2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2RoleRevoked represents a RoleRevoked event raised by the Bookingtokenv2 contract.
type Bookingtokenv2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Bookingtokenv2RoleRevokedIterator, error) {

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

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2RoleRevokedIterator{contract: _Bookingtokenv2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2RoleRevoked)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseRoleRevoked(log types.Log) (*Bookingtokenv2RoleRevoked, error) {
	event := new(Bookingtokenv2RoleRevoked)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2TokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenBoughtIterator struct {
	Event *Bookingtokenv2TokenBought // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2TokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2TokenBought)
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
		it.Event = new(Bookingtokenv2TokenBought)
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
func (it *Bookingtokenv2TokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2TokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2TokenBought represents a TokenBought event raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenBought struct {
	TokenId *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0xa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb95040.
//
// Solidity: event TokenBought(uint256 indexed tokenId, address indexed buyer)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterTokenBought(opts *bind.FilterOpts, tokenId []*big.Int, buyer []common.Address) (*Bookingtokenv2TokenBoughtIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "TokenBought", tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2TokenBoughtIterator{contract: _Bookingtokenv2.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0xa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb95040.
//
// Solidity: event TokenBought(uint256 indexed tokenId, address indexed buyer)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2TokenBought, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "TokenBought", tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2TokenBought)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenBought", log); err != nil {
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

// ParseTokenBought is a log parse operation binding the contract event 0xa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb95040.
//
// Solidity: event TokenBought(uint256 indexed tokenId, address indexed buyer)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseTokenBought(log types.Log) (*Bookingtokenv2TokenBought, error) {
	event := new(Bookingtokenv2TokenBought)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2TokenCancellableUpdatedIterator is returned from FilterTokenCancellableUpdated and is used to iterate over the raw logs and unpacked data for TokenCancellableUpdated events raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenCancellableUpdatedIterator struct {
	Event *Bookingtokenv2TokenCancellableUpdated // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2TokenCancellableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2TokenCancellableUpdated)
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
		it.Event = new(Bookingtokenv2TokenCancellableUpdated)
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
func (it *Bookingtokenv2TokenCancellableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2TokenCancellableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2TokenCancellableUpdated represents a TokenCancellableUpdated event raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenCancellableUpdated struct {
	TokenId       *big.Int
	IsCancellable bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenCancellableUpdated is a free log retrieval operation binding the contract event 0x014e546de98d6e31d4722b8674630820a5e455f2b43c09805d19b58b64a14708.
//
// Solidity: event TokenCancellableUpdated(uint256 indexed tokenId, bool isCancellable)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterTokenCancellableUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*Bookingtokenv2TokenCancellableUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "TokenCancellableUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2TokenCancellableUpdatedIterator{contract: _Bookingtokenv2.contract, event: "TokenCancellableUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenCancellableUpdated is a free log subscription operation binding the contract event 0x014e546de98d6e31d4722b8674630820a5e455f2b43c09805d19b58b64a14708.
//
// Solidity: event TokenCancellableUpdated(uint256 indexed tokenId, bool isCancellable)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchTokenCancellableUpdated(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2TokenCancellableUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "TokenCancellableUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2TokenCancellableUpdated)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenCancellableUpdated", log); err != nil {
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

// ParseTokenCancellableUpdated is a log parse operation binding the contract event 0x014e546de98d6e31d4722b8674630820a5e455f2b43c09805d19b58b64a14708.
//
// Solidity: event TokenCancellableUpdated(uint256 indexed tokenId, bool isCancellable)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseTokenCancellableUpdated(log types.Log) (*Bookingtokenv2TokenCancellableUpdated, error) {
	event := new(Bookingtokenv2TokenCancellableUpdated)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenCancellableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2TokenExpiredIterator is returned from FilterTokenExpired and is used to iterate over the raw logs and unpacked data for TokenExpired events raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenExpiredIterator struct {
	Event *Bookingtokenv2TokenExpired // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2TokenExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2TokenExpired)
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
		it.Event = new(Bookingtokenv2TokenExpired)
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
func (it *Bookingtokenv2TokenExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2TokenExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2TokenExpired represents a TokenExpired event raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenExpired struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenExpired is a free log retrieval operation binding the contract event 0x492531370c4d9936ebe217e769581e72fb2a02b10df161cd9ccd358f1aa45f9a.
//
// Solidity: event TokenExpired(uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterTokenExpired(opts *bind.FilterOpts, tokenId []*big.Int) (*Bookingtokenv2TokenExpiredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "TokenExpired", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2TokenExpiredIterator{contract: _Bookingtokenv2.contract, event: "TokenExpired", logs: logs, sub: sub}, nil
}

// WatchTokenExpired is a free log subscription operation binding the contract event 0x492531370c4d9936ebe217e769581e72fb2a02b10df161cd9ccd358f1aa45f9a.
//
// Solidity: event TokenExpired(uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchTokenExpired(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2TokenExpired, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "TokenExpired", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2TokenExpired)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenExpired", log); err != nil {
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

// ParseTokenExpired is a log parse operation binding the contract event 0x492531370c4d9936ebe217e769581e72fb2a02b10df161cd9ccd358f1aa45f9a.
//
// Solidity: event TokenExpired(uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseTokenExpired(log types.Log) (*Bookingtokenv2TokenExpired, error) {
	event := new(Bookingtokenv2TokenExpired)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2TokenReservedIterator is returned from FilterTokenReserved and is used to iterate over the raw logs and unpacked data for TokenReserved events raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenReservedIterator struct {
	Event *Bookingtokenv2TokenReserved // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2TokenReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2TokenReserved)
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
		it.Event = new(Bookingtokenv2TokenReserved)
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
func (it *Bookingtokenv2TokenReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2TokenReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2TokenReserved represents a TokenReserved event raised by the Bookingtokenv2 contract.
type Bookingtokenv2TokenReserved struct {
	TokenId             *big.Int
	ReservedFor         common.Address
	Supplier            common.Address
	ExpirationTimestamp *big.Int
	Price               *big.Int
	PaymentToken        common.Address
	IsCancellable       bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokenReserved is a free log retrieval operation binding the contract event 0x458ab0460080169903d98b67072b89a34af76e5da8a99d1046a66de159966086.
//
// Solidity: event TokenReserved(uint256 indexed tokenId, address indexed reservedFor, address indexed supplier, uint256 expirationTimestamp, uint256 price, address paymentToken, bool isCancellable)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterTokenReserved(opts *bind.FilterOpts, tokenId []*big.Int, reservedFor []common.Address, supplier []common.Address) (*Bookingtokenv2TokenReservedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var reservedForRule []interface{}
	for _, reservedForItem := range reservedFor {
		reservedForRule = append(reservedForRule, reservedForItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "TokenReserved", tokenIdRule, reservedForRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2TokenReservedIterator{contract: _Bookingtokenv2.contract, event: "TokenReserved", logs: logs, sub: sub}, nil
}

// WatchTokenReserved is a free log subscription operation binding the contract event 0x458ab0460080169903d98b67072b89a34af76e5da8a99d1046a66de159966086.
//
// Solidity: event TokenReserved(uint256 indexed tokenId, address indexed reservedFor, address indexed supplier, uint256 expirationTimestamp, uint256 price, address paymentToken, bool isCancellable)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchTokenReserved(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2TokenReserved, tokenId []*big.Int, reservedFor []common.Address, supplier []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var reservedForRule []interface{}
	for _, reservedForItem := range reservedFor {
		reservedForRule = append(reservedForRule, reservedForItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "TokenReserved", tokenIdRule, reservedForRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2TokenReserved)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenReserved", log); err != nil {
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

// ParseTokenReserved is a log parse operation binding the contract event 0x458ab0460080169903d98b67072b89a34af76e5da8a99d1046a66de159966086.
//
// Solidity: event TokenReserved(uint256 indexed tokenId, address indexed reservedFor, address indexed supplier, uint256 expirationTimestamp, uint256 price, address paymentToken, bool isCancellable)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseTokenReserved(log types.Log) (*Bookingtokenv2TokenReserved, error) {
	event := new(Bookingtokenv2TokenReserved)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "TokenReserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bookingtokenv2 contract.
type Bookingtokenv2TransferIterator struct {
	Event *Bookingtokenv2Transfer // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2Transfer)
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
		it.Event = new(Bookingtokenv2Transfer)
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
func (it *Bookingtokenv2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2Transfer represents a Transfer event raised by the Bookingtokenv2 contract.
type Bookingtokenv2Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Bookingtokenv2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2TransferIterator{contract: _Bookingtokenv2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2Transfer)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseTransfer(log types.Log) (*Bookingtokenv2Transfer, error) {
	event := new(Bookingtokenv2Transfer)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bookingtokenv2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bookingtokenv2 contract.
type Bookingtokenv2UpgradedIterator struct {
	Event *Bookingtokenv2Upgraded // Event containing the contract specifics and raw log

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
func (it *Bookingtokenv2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bookingtokenv2Upgraded)
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
		it.Event = new(Bookingtokenv2Upgraded)
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
func (it *Bookingtokenv2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bookingtokenv2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bookingtokenv2Upgraded represents a Upgraded event raised by the Bookingtokenv2 contract.
type Bookingtokenv2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Bookingtokenv2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenv2UpgradedIterator{contract: _Bookingtokenv2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bookingtokenv2 *Bookingtokenv2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Bookingtokenv2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bookingtokenv2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bookingtokenv2Upgraded)
				if err := _Bookingtokenv2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Bookingtokenv2 *Bookingtokenv2Filterer) ParseUpgraded(log types.Log) (*Bookingtokenv2Upgraded, error) {
	event := new(Bookingtokenv2Upgraded)
	if err := _Bookingtokenv2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
