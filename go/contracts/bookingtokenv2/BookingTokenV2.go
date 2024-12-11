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
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161460b61003e600039600081816125eb01528181612614015261275e015261460b6000f3fe6080604052600436106102595760003560e01c80624fdd3c1461025e57806301ffc9a71461029557806306fdde03146102c5578063081812fc146102e7578063095ea7b3146103145780630e75c1a81461033657806312b357b51461035957806318160ddd1461037957806323b872dd1461038e578063248a9ca3146103ae5780632d3a6329146103ce5780632e0258a0146103ee5780632edf5e2c1461040e5780632f2ff15d146104305780632f745c591461045057806336568abe146104705780633c15b31c1461049057806341431908146104bd57806342842e0e146104dd578063454d0db9146104fd5780634f1ef286146105305780634f6ccce714610543578063516a82b81461056357806352d1902d146105835780635d4badb2146105985780636352211e146105b85780636872b7e6146105d85780636fabeaed146105f857806370a082311461061857806391d148541461063857806391da124c1461065857806395d89b411461067857806396591edd1461068d578063a0f07c74146106a0578063a101c67e146106b5578063a217fddf146106d5578063a22cb465146106ea578063a6dbb9cb1461070a578063ad3cb1cc14610775578063b191d092146107a6578063b47177a1146107c6578063b88d4fde146107e6578063bfb26c0614610806578063c0c53b8b1461081b578063c162d7da1461083b578063c87b56dd14610850578063d547741f14610870578063db2b268214610890578063de62fe4d146108b0578063e5a6725c146108c3578063e985e9c5146108e3578063f024df2014610903578063f72c0d8b14610923575b600080fd5b34801561026a57600080fd5b5061027e610279366004613b5e565b610945565b60405161028c929190613b77565b60405180910390f35b3480156102a157600080fd5b506102b56102b0366004613ba4565b610982565b604051901515815260200161028c565b3480156102d157600080fd5b506102da610993565b60405161028c9190613c11565b3480156102f357600080fd5b50610307610302366004613b5e565b610a34565b60405161028c9190613c24565b34801561032057600080fd5b5061033461032f366004613c4d565b610a49565b005b34801561034257600080fd5b5061034b610a58565b60405190815260200161028c565b34801561036557600080fd5b506102b5610374366004613c79565b610a6d565b34801561038557600080fd5b5061034b610ae3565b34801561039a57600080fd5b506103346103a9366004613c96565b610aee565b3480156103ba57600080fd5b5061034b6103c9366004613b5e565b610b07565b3480156103da57600080fd5b506102b56103e9366004613b5e565b610b27565b3480156103fa57600080fd5b50610334610409366004613cd7565b610b48565b34801561041a57600080fd5b5061034b60008051602061459683398151915281565b34801561043c57600080fd5b5061033461044b366004613cf9565b610d27565b34801561045c57600080fd5b5061034b61046b366004613c4d565b610d49565b34801561047c57600080fd5b5061033461048b366004613cf9565b610dab565b34801561049c57600080fd5b506104b06104ab366004613b5e565b610dde565b60405161028c9190613d5d565b3480156104c957600080fd5b506103346104d8366004613c79565b610e01565b3480156104e957600080fd5b506103346104f8366004613c96565b610e39565b34801561050957600080fd5b5061051d610518366004613b5e565b610e54565b60405161028c9796959493929190613d70565b61033461053e366004613e63565b610f49565b34801561054f57600080fd5b5061034b61055e366004613b5e565b610f64565b34801561056f57600080fd5b5061033461057e366004613b5e565b610fc5565b34801561058f57600080fd5b5061034b610ff2565b3480156105a457600080fd5b506103346105b3366004613eb2565b61100f565b3480156105c457600080fd5b506103076105d3366004613b5e565b611026565b3480156105e457600080fd5b506103346105f3366004613f40565b611031565b34801561060457600080fd5b50610334610613366004613f40565b61124c565b34801561062457600080fd5b5061034b610633366004613c79565b611377565b34801561064457600080fd5b506102b5610653366004613cf9565b6113cf565b34801561066457600080fd5b50610334610673366004613f7c565b611405565b34801561068457600080fd5b506102da6114c2565b61033461069b366004613b5e565b6114df565b3480156106ac57600080fd5b50610307600081565b3480156106c157600080fd5b5061034b6106d0366004613b5e565b61167e565b3480156106e157600080fd5b5061034b600081565b3480156106f657600080fd5b50610334610705366004613fe3565b611699565b34801561071657600080fd5b5061072a610725366004613b5e565b6116a4565b60405161028c919081516001600160a01b0390811682526020808401518216908301526040808401519083015260608084015190830152608092830151169181019190915260a00190565b34801561078157600080fd5b506102da604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156107b257600080fd5b506103076107c1366004613b5e565b611736565b3480156107d257600080fd5b506103346107e1366004614011565b611760565b3480156107f257600080fd5b50610334610801366004614057565b611973565b34801561081257600080fd5b50610307600181565b34801561082757600080fd5b506103346108363660046140c2565b611988565b34801561084757600080fd5b50610307611b28565b34801561085c57600080fd5b506102da61086b366004613b5e565b611b43565b34801561087c57600080fd5b5061033461088b366004613cf9565b611b4e565b34801561089c57600080fd5b506103346108ab36600461410d565b611b6a565b6103346108be366004613cd7565b611cc9565b3480156108cf57600080fd5b506103346108de366004613b5e565b611fce565b3480156108ef57600080fd5b506102b56108fe3660046141a0565b6120d6565b34801561090f57600080fd5b5061033461091e366004613cd7565b612115565b34801561092f57600080fd5b5061034b60008051602061455683398151915281565b6000806000610952612243565b6000948552600390810160205260409094209384015460049094015493946001600160a01b039094169392505050565b600061098d82612267565b92915050565b6060600061099f61228c565b90508060000180546109b0906141ce565b80601f01602080910402602001604051908101604052809291908181526020018280546109dc906141ce565b8015610a295780601f106109fe57610100808354040283529160200191610a29565b820191906000526020600020905b815481529060010190602001808311610a0c57829003601f168201915b505050505091505090565b6000610a3f826122b0565b5061098d826122e8565b610a54828233612311565b5050565b600080610a63612243565b6002015492915050565b6000610a77611b28565b6001600160a01b03166312b357b5836040518263ffffffff1660e01b8152600401610aa29190613c24565b602060405180830381865afa158015610abf573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098d9190614208565b600080610a6361231e565b610af781612342565b610b028383836123d5565b505050565b600080610b12612446565b60009384526020525050604090206001015490565b6000610b3161246a565b600092835260010160205250604090205460ff1690565b6000610b5261246a565b6000848152602082905260409020909150600180820154600160a01b900460ff166004811115610b8457610b84613d29565b14158015610bb2575060026001820154600160a01b900460ff166004811115610baf57610baf613d29565b14155b15610bd85760405163a45a187b60e01b8152600481018590526024015b60405180910390fd5b6000610be2612243565b60008681526003820160205260408120600101549192506001600160a01b0390911690610c0e876122b0565b9050336001600160a01b038316141580610c38575060018401546001600160a01b03828116911614155b15610c5a578633604051639259da7360e01b8152600401610bcf929190613b77565b600087815260208690526040902086905560026001850154600160a01b900460ff166004811115610c8d57610c8d613d29565b03610cc55760008781526020869052604081206002015490610cb19082908061248e565b600089815260208890526040902060020155505b60008781526020868152604091829020600101805460ff60a01b1916600360a01b1790559051878152339189917f7b6cf4b0eeba58e59d225cabf115102b6f1f5af0515f5f6bb6ec9709bc854d0991015b60405180910390a350505050505050565b610d3082610b07565b610d39816124bd565b610d4383836124c7565b50505050565b600080610d5461231e565b9050610d5f84611377565b8310610d8257838360405163295f44f760e21b8152600401610bcf929190614225565b6001600160a01b0384166000908152602091825260408082208583529092522054905092915050565b6001600160a01b0381163314610dd45760405163334bd91960e11b815260040160405180910390fd5b610b028282612568565b600080610de9612243565b60009384526004016020525050604090205460ff1690565b6000610e0c816124bd565b6000610e16612243565b80546001600160a01b0319166001600160a01b0394909416939093179092555050565b610b0283838360405180602001604052806000815250611973565b600080600080600080600080610e6861246a565b60008a81526020828152604080832081516080810183528154815260018201546001600160a01b038116948201949094529495509293929190830190600160a01b900460ff166004811115610ebf57610ebf613d29565b6004811115610ed057610ed0613d29565b81526020016002820154815250509050600080600080610f0f856060015161ffff81811692601083901c821692602081901c83169260309190911c1690565b9350935093509350846000015185602001518660400151868686869c509c509c509c509c509c509c50505050505050919395979092949650565b610f516125e0565b610f5a82612687565b610a54828261269f565b600080610f6f61231e565b9050610f79610ae3565b8310610f9d5760008360405163295f44f760e21b8152600401610bcf929190614225565b806002018381548110610fb257610fb261423e565b9060005260206000200154915050919050565b600080516020614596833981519152610fdd816124bd565b6000610fe7612243565b600201929092555050565b6000610ffc612753565b5060008051602061457683398151915290565b61101f8585858585600080611b6a565b5050505050565b600061098d826122b0565b600061103b61246a565b60008581526020828152604080832081516080810183528154815260018201546001600160a01b038116948201949094529495509293929190830190600160a01b900460ff16600481111561109257611092613d29565b60048111156110a3576110a3613d29565b81526002919091015460209091015290506001816040015160048111156110cc576110cc613d29565b146110ed5760405163a45a187b60e01b815260048101869052602401610bcf565b60006110f7612243565b6000878152600380830160209081526040808420815160a08101835281546001600160a01b039081168252600183015481169482019490945260028201549281019290925292830154606082015260049092015416608082015291925061115d886122b0565b905081602001516001600160a01b0316336001600160a01b031614158061119a5750806001600160a01b031684602001516001600160a01b031614155b156111bc57873360405163765a32d760e11b8152600401610bcf929190613b77565b600088815260208690526040902060018101805460ff60a01b1916600160a11b179055600201546111ee81898961248e565b60008a81526020889052604090819020600201919091555133908a907f258940b052b1b31e944c2415ac9684f3f5f5e2294871b8956dd432bc7c7edca390611239908c908c90614254565b60405180910390a3505050505050505050565b600061125661246a565b60008581526020829052604090206001810154919250906001600160a01b0316331461129957843360405163a97065df60e01b8152600401610bcf929190613b77565b600180820154600160a01b900460ff1660048111156112ba576112ba613d29565b141580156112e8575060036001820154600160a01b900460ff1660048111156112e5576112e5613d29565b14155b156113095760405163a45a187b60e01b815260048101869052602401610bcf565b6000858152602083905260408082208281556001810180546001600160a81b03191690556002019190915551339086907f13a82951767f2a7efeb88a92151eb59d8be7a42817832b1ba9a8f7d8960c91fb906113689088908890614254565b60405180910390a35050505050565b60008061138261228c565b90506001600160a01b0383166113ae5760006040516322718ad960e21b8152600401610bcf9190613c24565b6001600160a01b039092166000908152600390920160205250604090205490565b6000806113da612446565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b6002600061141161279c565b8054909150600160401b900460ff1680611438575080546001600160401b03808416911610155b156114565760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160481b0319166001600160401b03831617600160401b1781556000611480816124bd565b61148a85856127c0565b50805460ff60401b19168155604051600080516020614536833981519152906114b4908490614269565b60405180910390a150505050565b606060006114ce61228c565b90508060010180546109b0906141ce565b6114e76127d2565b336114f181612808565b60006114fb612243565b600084815260038083016020908152604092839020835160a08101855281546001600160a01b039081168083526001840154821694830194909452600283015495820195909552928101546060840152600401549092166080820152919250331461158d578051604051632663a9c560e11b81526001600160a01b039091166004820152336024820152604401610bcf565b80604001514211156115ba5783816040015160405163293d73b760e11b8152600401610bcf92919061427d565b60006115c585611026565b905081602001516001600160a01b0316816001600160a01b03161461160457602082015160405163103d145960e31b8152610bcf918791600401613b77565b61161382602001513387612830565b61162a8260800151836060015184602001516128cc565b6000858152600484016020526040808220805460ff1916600317905551339187917fa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb950409190a35050505061167b61297a565b50565b600061168861246a565b600092835260205250604090205490565b610a5433838361298b565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526116d7612243565b600092835260039081016020908152604093849020845160a08101865281546001600160a01b03908116825260018301548116938201939093526002820154958101959095529182015460608501526004909101541660808301525090565b6000611740612243565b60009283526003016020525060409020600401546001600160a01b031690565b600061176a612243565b905060036000868152600480840160205260409091205460ff169081111561179457611794613d29565b146117cb576000858152600480830160205260409182902054915163e4e3b53b60e01b8152610bcf92889260ff909116910161428b565b60006117d6866122b0565b60008781526003840160205260409020600101549091506001600160a01b0390811690821633148015906118135750336001600160a01b03821614155b1561183357336040516343d11b6760e11b8152600401610bcf9190613c24565b600061183d61246a565b905060008089815260208390526040902060010154600160a01b900460ff16600481111561186d5761186d613d29565b1461188e576040516333cba99b60e21b815260048101899052602401610bcf565b604080516080810182528881523360208201526001918101919091526060810161ffff8816601088901b63ffff0000161790526000898152602083815260409182902083518155908301516001820180546001600160a01b039092166001600160a01b031983168117825593850151929390916001600160a81b03191617600160a01b83600481111561192357611923613d29565b021790555060608201518160020155905050336001600160a01b0316886000805160206145b6833981519152898989604051611961939291906142a8565b60405180910390a35050505050505050565b61197c82612342565b610d4384848484612a30565b600061199261279c565b805490915060ff600160401b82041615906001600160401b03166000811580156119b95750825b90506000826001600160401b031660011480156119d55750303b155b9050811580156119e3575080155b15611a015760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315611a2a57845460ff60401b1916600160401b1785555b611a746040518060400160405280600c81526020016b2137b7b5b4b733aa37b5b2b760a11b815250604051806040016040528060048152602001630545249560e41b8152506127c0565b611a7c612a48565b611a84612a48565b611a8c612a48565b611a94612a48565b611a9f6000886124c7565b50611ab8600080516020614556833981519152876124c7565b506000611ac3612243565b80546001600160a01b0319166001600160a01b038b16178155603c600290910155508315611b1e57845460ff60401b1916855560405160008051602061453683398151915290611b1590600190614269565b60405180910390a15b5050505050505050565b600080611b33612243565b546001600160a01b031692915050565b606061098d82612a50565b611b5782610b07565b611b60816124bd565b610d438383612568565b33611b7481612808565b611b7d88612808565b6000611b87612243565b6002810154909150611b9981426142d8565b8811611bbc578781604051630999f7d760e41b8152600401610bcf92919061427d565b60018201805460009182611bcf836142eb565b919050559050611bdf3382612b70565b611be9818b612b8a565b611bf7818c338c8c8c612be8565b60008181526004840160205260409020805460ff1916600117905584611c1b61246a565b600083815260019190910160205260409020805460ff191691151591909117905585611c4561246a565b600083815260029190910160209081526040918290209290925580518b81529182018a90526001600160a01b038981168383015260608301899052871515608084015290513392918e169184917f1424af4f4cb40d8a1a2d00b2324cb122ba73eac426f98b62c33ff31ca045f0679181900360a00190a45050505050505050505050565b611cd16127d2565b33611cdb81612808565b6000611ce561246a565b60008581526020828152604080832081516080810183528154815260018201546001600160a01b038116948201949094529495509293929190830190600160a01b900460ff166004811115611d3c57611d3c613d29565b6004811115611d4d57611d4d613d29565b8152602001600282015481525050905080600001518414611d85578051604051630f81919b60e21b8152610bcf91869160040161427d565b600181604001516004811115611d9d57611d9d613d29565b14611dbe5760405163a45a187b60e01b815260048101869052602401610bcf565b6000611dc8612243565b6000878152600380830160209081526040808420815160a08101835281546001600160a01b0390811682526001830154811694820194909452600282015492810192909252928301546060820152600490920154166080820152919250611e2e886122b0565b905081602001516001600160a01b031684602001516001600160a01b0316148015611e615750336001600160a01b038216145b15611ed4576000888152602086815260409182902060010180546001600160a01b0319166001600160a01b03851690811790915586519251928352918a917fab9781c7afb2bca22dd3691027e5ef637de104315f60d4d2c4328bc9161c2668910160405180910390a35050505050611fc5565b81602001516001600160a01b0316336001600160a01b0316141580611f0f5750806001600160a01b031684602001516001600160a01b031614155b15611f3157873360405163e363bce560e01b8152600401610bcf929190613b77565b6000888152600484810160209081526040808420805460ff19169093179092558790529020600101805460ff60a01b1916600160a21b179055611f7388612c89565b8351604051908152339089907fdf11499efc7ab0feb9befa7a615c79d3df759d9930b41f31c1f0723cfd9b10f99060200160405180910390a3611fbf82608001518560000151836128cc565b50505050505b50610a5461297a565b6000611fd8612243565b6000838152600382016020908152604080832060048501909252909120549192509060ff16600281600481111561201157612011613d29565b148061202e5750600381600481111561202c5761202c613d29565b145b8061204a5750600481600481111561204857612048613d29565b145b1561206c57838160405163e4e3b53b60e01b8152600401610bcf92919061428b565b81600201544211156120ad576000848152600484016020526040808220805460ff1916600217905551859160008051602061451683398151915291a2610d43565b815460405163d4cde2af60e01b8152610bcf9186916001600160a01b0390911690600401613b77565b6000806120e161228c565b6001600160a01b03948516600090815260059190910160209081526040808320959096168252939093525050205460ff1690565b6000612120836122b0565b9050336001600160a01b0382161461214f57823360405163d699eeab60e01b8152600401610bcf929190613b77565b600061215961246a565b60008581526020829052604090208054919250908414612190578054604051630f81919b60e21b8152610bcf91869160040161427d565b60036001820154600160a01b900460ff1660048111156121b2576121b2613d29565b146121d357604051632070575960e01b815260048101869052602401610bcf565b6000858152602083905260408120600181018054600160a01b60ff60a01b1990911617905560020154819061ffff8082169160101c16600185015485546040519395509193506001600160a01b03169189916000805160206145b683398151915291610d169190879087906142a8565b7f9db9d405bf15683ce835607b1f0b423dc1484d44bb9d5af64a483fa4afd8290090565b60006001600160e01b03198216637965db0b60e01b148061098d575061098d82612cc4565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930090565b6000806122bc83612ce9565b90506001600160a01b03811661098d57604051637e27328960e01b815260048101849052602401610bcf565b6000806122f361228c565b6000938452600401602052505060409020546001600160a01b031690565b610b028383836001612d12565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0090565b61234b81612e1c565b600061235561246a565b600083815260208290526040902090915060026001820154600160a01b900460ff16600481111561238857612388613d29565b0361239257505050565b60006001820154600160a01b900460ff1660048111156123b4576123b4613d29565b14610b02576040516333cba99b60e21b815260048101849052602401610bcf565b6001600160a01b0382166123ff576000604051633250574960e11b8152600401610bcf9190613c24565b600061240c838333612f29565b9050836001600160a01b0316816001600160a01b031614610d43578382826040516364283d7b60e01b8152600401610bcf93929190614304565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b7f56ee42015b616c256a09657decaf7aa5d877decbe489a4b4b22f8bb47660050090565b61ffff60301b60309190911b1661ffff60201b60209290921b919091161763ffffffff60201b19919091161790565b61167b8133612f40565b6000806124d2612446565b90506124de84846113cf565b61255e576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556125143390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061098d565b600091505061098d565b600080612573612446565b905061257f84846113cf565b1561255e576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061098d565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061266757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661265b600080516020614576833981519152546001600160a01b031690565b6001600160a01b031614155b156126855760405163703e46dd60e11b815260040160405180910390fd5b565b600080516020614556833981519152610a54816124bd565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156126f9575060408051601f3d908101601f191682019092526126f691810190614327565b60015b6127185781604051634c9c8ce360e01b8152600401610bcf9190613c24565b600080516020614576833981519152811461274957604051632a87526960e21b815260048101829052602401610bcf565b610b028383612f6b565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146126855760405163703e46dd60e11b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b6127c8612fc1565b610a548282612fe6565b60006127dc613014565b80549091506001190161280257604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b61281181610a6d565b61167b5780604051638014575360e01b8152600401610bcf9190613c24565b6001600160a01b03821661285a576000604051633250574960e11b8152600401610bcf9190613c24565b600061286883836000612f29565b90506001600160a01b03811661289457604051637e27328960e01b815260048101839052602401610bcf565b836001600160a01b0316816001600160a01b031614610d43578382826040516364283d7b60e01b8152600401610bcf93929190614304565b6001600160a01b038316612911578134146128fe573482604051630145611560e21b8152600401610bcf92919061427d565b610b026001600160a01b03821634613038565b6000196001600160a01b03841601612944573415610b02576040516347d6729960e01b8152346004820152602401610bcf565b3415612965576040516347d6729960e01b8152346004820152602401610bcf565b610b026001600160a01b0384163383856130d1565b6000612984613014565b6001905550565b600061299561228c565b90506001600160a01b0383166129c05782604051630b61174360e31b8152600401610bcf9190613c24565b6001600160a01b038481166000818152600584016020908152604080832094881680845294825291829020805460ff191687151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a350505050565b612a3b848484610aee565b610d43338585858561312b565b612685612fc1565b60606000612a5c613243565b9050612a67836122b0565b5060008381526020829052604081208054612a81906141ce565b80601f0160208091040260200160405190810160405280929190818152602001828054612aad906141ce565b8015612afa5780601f10612acf57610100808354040283529160200191612afa565b820191906000526020600020905b815481529060010190602001808311612add57829003601f168201915b505050505090506000612b1860408051602081019091526000815290565b90508051600003612b2b57509392505050565b815115612b5e578082604051602001612b45929190614340565b6040516020818303038152906040529350505050919050565b612b6785613267565b95945050505050565b610a548282604051806020016040528060008152506132db565b6000612b94613243565b6000848152602082905260409020909150612baf83826143b7565b506040518381527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a1505050565b6000612bf2612243565b6040805160a0810182526001600160a01b0398891681529688166020808901918252888301978852606089019687529489166080890190815260009a8b5260039384019095529820955186546001600160a01b031990811691891691909117875597516001870180548a16918916919091179055935160028601555090519183019190915551600490910180549093169116179055565b6000612c986000836000612f29565b90506001600160a01b038116610a5457604051637e27328960e01b815260048101839052602401610bcf565b60006001600160e01b03198216632483248360e11b148061098d575061098d826132f3565b600080612cf461228c565b6000938452600201602052505060409020546001600160a01b031690565b6000612d1c61228c565b90508180612d3257506001600160a01b03831615155b15612deb576000612d42856122b0565b90506001600160a01b03841615801590612d6e5750836001600160a01b0316816001600160a01b031614155b8015612d815750612d7f81856120d6565b155b15612da1578360405163a9fbf51f60e01b8152600401610bcf9190613c24565b8215612de95784866001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b600093845260040160205250506040902080546001600160a01b0319166001600160a01b0392909216919091179055565b6000612e26612243565b600083815260048201602052604090205490915060ff166003816004811115612e5157612e51613d29565b1480612e6e57506002816004811115612e6c57612e6c613d29565b145b15612e7857505050565b6004816004811115612e8c57612e8c613d29565b03612eae57828160405163e4e3b53b60e01b8152600401610bcf92919061428b565b600083815260038301602052604090206002810154421115612f00576000848152600484016020526040808220805460ff1916600217905551859160008051602061451683398151915291a250505050565b805460405163d4cde2af60e01b8152610bcf9186916001600160a01b0390911690600401613b77565b6000612f36848484613318565b90505b9392505050565b612f4a82826113cf565b610a5457808260405163e2517d3f60e01b8152600401610bcf929190614225565b612f74826133a6565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115612fb957610b028282613402565b610a5461346f565b612fc961348e565b61268557604051631afcd79f60e31b815260040160405180910390fd5b612fee612fc1565b6000612ff861228c565b90508061300584826143b7565b5060018101610d4383826143b7565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b8047101561305d57478160405163cf47918160e01b8152600401610bcf92919061427d565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146130aa576040519150601f19603f3d011682016040523d82523d6000602084013e6130af565b606091505b5050905080610b025760405163d6bda27560e01b815260040160405180910390fd5b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610d439085906134a8565b6001600160a01b0383163b1561101f57604051630a85bd0160e11b81526001600160a01b0384169063150b7a029061316d908890889087908790600401614476565b6020604051808303816000875af19250505080156131a8575060408051601f3d908101601f191682019092526131a5918101906144b3565b60015b613208573d8080156131d6576040519150601f19603f3d011682016040523d82523d6000602084013e6131db565b606091505b5080516000036132005783604051633250574960e11b8152600401610bcf9190613c24565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b1461323b5783604051633250574960e11b8152600401610bcf9190613c24565b505050505050565b7f0542a41881ee128a365a727b282c86fa859579490b9bb45aab8503648c8e790090565b6060613272826122b0565b50600061328a60408051602081019091526000815290565b905060008151116132aa5760405180602001604052806000815250612f39565b806132b484613510565b6040516020016132c5929190614340565b6040516020818303038152906040529392505050565b6132e583836135a2565b610b0233600085858561312b565b60006001600160e01b0319821663780e9d6360e01b148061098d575061098d82613607565b600080613326858585613657565b90506001600160a01b0381166133445761333f8461375e565b613367565b846001600160a01b0316816001600160a01b031614613367576133678185613797565b6001600160a01b0385166133835761337e84613830565b612f36565b846001600160a01b0316816001600160a01b031614612f3657612f3685856138f9565b806001600160a01b03163b6000036133d35780604051634c9c8ce360e01b8152600401610bcf9190613c24565b60008051602061457683398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161341f91906144d0565b600060405180830381855af49150503d806000811461345a576040519150601f19603f3d011682016040523d82523d6000602084013e61345f565b606091505b5091509150612b67858383613951565b34156126855760405163b398979f60e01b815260040160405180910390fd5b600061349861279c565b54600160401b900460ff16919050565b600080602060008451602086016000885af1806134cb576040513d6000823e3d81fd5b50506000513d915081156134e35780600114156134f0565b6001600160a01b0384163b155b15610d435783604051635274afe760e01b8152600401610bcf9190613c24565b6060600061351d836139a4565b60010190506000816001600160401b0381111561353c5761353c613dc1565b6040519080825280601f01601f191660200182016040528015613566576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461357057509392505050565b6001600160a01b0382166135cc576000604051633250574960e11b8152600401610bcf9190613c24565b60006135da83836000612f29565b90506001600160a01b03811615610b025760006040516339e3563760e11b8152600401610bcf9190613c24565b60006001600160e01b031982166380ac58cd60e01b148061363857506001600160e01b03198216635b5e139f60e01b145b8061098d57506301ffc9a760e01b6001600160e01b031983161461098d565b60008061366261228c565b9050600061366f85612ce9565b90506001600160a01b0384161561368b5761368b818587613a7a565b6001600160a01b038116156136cb576136a8600086600080612d12565b6001600160a01b0381166000908152600383016020526040902080546000190190555b6001600160a01b038616156136fc576001600160a01b03861660009081526003830160205260409020805460010190555b600085815260028301602052604080822080546001600160a01b0319166001600160a01b038a811691821790925591518893918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a495945050505050565b600061376861231e565b600281018054600085815260039093016020908152604084208290556001820183559183529120019190915550565b60006137a161231e565b905060006137ae84611377565b60008481526001840160209081526040808320546001600160a01b03891684529186905290912091925090818314613808576000838152602082815260408083205485845281842081905583526001870190915290208290555b6000948552600190930160209081526040808620869055928552929092528220919091555050565b600061383a61231e565b6002810154909150600090613851906001906144ec565b600084815260038401602052604081205460028501805493945090928490811061387d5761387d61423e565b90600052602060002001549050808460020183815481106138a0576138a061423e565b6000918252602080832090910192909255828152600386019091526040808220849055868252812055600284018054806138dc576138dc6144ff565b600190038181906000526020600020016000905590555050505050565b600061390361231e565b90506000600161391285611377565b61391c91906144ec565b6001600160a01b0390941660009081526020838152604080832087845282528083208690559482526001909301909252502055565b6060826139665761396182613ad0565b612f39565b815115801561397d57506001600160a01b0384163b155b1561399d5783604051639996b31560e01b8152600401610bcf9190613c24565b5080612f39565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106139e35772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6904ee2d6d415b85acef8160201b8310613a0d576904ee2d6d415b85acef8160201b830492506020015b662386f26fc100008310613a2b57662386f26fc10000830492506010015b6305f5e1008310613a43576305f5e100830492506008015b6127108310613a5757612710830492506004015b60648310613a69576064830492506002015b600a831061098d5760010192915050565b613a85838383613af9565b610b02576001600160a01b038316613ab357604051637e27328960e01b815260048101829052602401610bcf565b818160405163177e802f60e01b8152600401610bcf929190614225565b805115613ae05780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b60006001600160a01b03831615801590612f365750826001600160a01b0316846001600160a01b03161480613b335750613b3384846120d6565b80612f365750826001600160a01b0316613b4c836122e8565b6001600160a01b031614949350505050565b600060208284031215613b7057600080fd5b5035919050565b9182526001600160a01b0316602082015260400190565b6001600160e01b03198116811461167b57600080fd5b600060208284031215613bb657600080fd5b8135612f3981613b8e565b60005b83811015613bdc578181015183820152602001613bc4565b50506000910152565b60008151808452613bfd816020860160208601613bc1565b601f01601f19169290920160200192915050565b602081526000612f396020830184613be5565b6001600160a01b0391909116815260200190565b6001600160a01b038116811461167b57600080fd5b60008060408385031215613c6057600080fd5b8235613c6b81613c38565b946020939093013593505050565b600060208284031215613c8b57600080fd5b8135612f3981613c38565b600080600060608486031215613cab57600080fd5b8335613cb681613c38565b92506020840135613cc681613c38565b929592945050506040919091013590565b60008060408385031215613cea57600080fd5b50508035926020909101359150565b60008060408385031215613d0c57600080fd5b823591506020830135613d1e81613c38565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b6005811061167b57634e487b7160e01b600052602160045260246000fd5b60208101613d6a83613d3f565b91905290565b8781526001600160a01b038716602082015260e08101613d8f87613d3f565b604082019690965261ffff9485166060820152928416608084015290831660a083015290911660c09091015292915050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112613de857600080fd5b81356001600160401b0380821115613e0257613e02613dc1565b604051601f8301601f19908116603f01168101908282118183101715613e2a57613e2a613dc1565b81604052838152866020858801011115613e4357600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215613e7657600080fd5b8235613e8181613c38565b915060208301356001600160401b03811115613e9c57600080fd5b613ea885828601613dd7565b9150509250929050565b600080600080600060a08688031215613eca57600080fd5b8535613ed581613c38565b945060208601356001600160401b03811115613ef057600080fd5b613efc88828901613dd7565b94505060408601359250606086013591506080860135613f1b81613c38565b809150509295509295909350565b803561ffff81168114613f3b57600080fd5b919050565b600080600060608486031215613f5557600080fd5b83359250613f6560208501613f29565b9150613f7360408501613f29565b90509250925092565b60008060408385031215613f8f57600080fd5b82356001600160401b0380821115613fa657600080fd5b613fb286838701613dd7565b93506020850135915080821115613fc857600080fd5b50613ea885828601613dd7565b801515811461167b57600080fd5b60008060408385031215613ff657600080fd5b823561400181613c38565b91506020830135613d1e81613fd5565b6000806000806080858703121561402757600080fd5b843593506020850135925061403e60408601613f29565b915061404c60608601613f29565b905092959194509250565b6000806000806080858703121561406d57600080fd5b843561407881613c38565b9350602085013561408881613c38565b92506040850135915060608501356001600160401b038111156140aa57600080fd5b6140b687828801613dd7565b91505092959194509250565b6000806000606084860312156140d757600080fd5b83356140e281613c38565b925060208401356140f281613c38565b9150604084013561410281613c38565b809150509250925092565b600080600080600080600060e0888a03121561412857600080fd5b873561413381613c38565b965060208801356001600160401b0381111561414e57600080fd5b61415a8a828b01613dd7565b9650506040880135945060608801359350608088013561417981613c38565b925060a0880135915060c088013561419081613fd5565b8091505092959891949750929550565b600080604083850312156141b357600080fd5b82356141be81613c38565b91506020830135613d1e81613c38565b600181811c908216806141e257607f821691505b60208210810361420257634e487b7160e01b600052602260045260246000fd5b50919050565b60006020828403121561421a57600080fd5b8151612f3981613fd5565b6001600160a01b03929092168252602082015260400190565b634e487b7160e01b600052603260045260246000fd5b61ffff92831681529116602082015260400190565b6001600160401b0391909116815260200190565b918252602082015260400190565b8281526040810161429b83613d3f565b8260208301529392505050565b92835261ffff918216602084015216604082015260600190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561098d5761098d6142c2565b6000600182016142fd576142fd6142c2565b5060010190565b6001600160a01b0393841681526020810192909252909116604082015260600190565b60006020828403121561433957600080fd5b5051919050565b60008351614352818460208801613bc1565b835190830190614366818360208801613bc1565b01949350505050565b601f821115610b02576000816000526020600020601f850160051c810160208610156143985750805b601f850160051c820191505b8181101561323b578281556001016143a4565b81516001600160401b038111156143d0576143d0613dc1565b6143e4816143de84546141ce565b8461436f565b602080601f83116001811461441957600084156144015750858301515b600019600386901b1c1916600185901b17855561323b565b600085815260208120601f198616915b8281101561444857888601518255948401946001909101908401614429565b50858210156144665787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906144a990830184613be5565b9695505050505050565b6000602082840312156144c557600080fd5b8151612f3981613b8e565b600082516144e2818460208701613bc1565b9190910192915050565b8181038181111561098d5761098d6142c2565b634e487b7160e01b600052603160045260246000fdfe492531370c4d9936ebe217e769581e72fb2a02b10df161cd9ccd358f1aa45f9ac7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc3ae8648b97d3fd425d26286fc6bb1d50724a93a6a5763921dd2b90405a83b4a479b98b853a6c04dfc02f383069b458c36ac4135883cdbf149b5aff904078e57ca264697066735822122050903aaff41241a8df365ff063c90c8de6ff174c581e8f71c26723d4689a6d5964736f6c63430008180033",
}

// Bookingtokenv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bookingtokenv2MetaData.ABI instead.
var Bookingtokenv2ABI = Bookingtokenv2MetaData.ABI

// Bookingtokenv2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Bookingtokenv2MetaData.Bin instead.
var Bookingtokenv2Bin = Bookingtokenv2MetaData.Bin

// DeployBookingtokenv2 deploys a new Ethereum contract, binding an instance of Bookingtokenv2 to it.
func DeployBookingtokenv2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bookingtokenv2, error) {
	parsed, err := Bookingtokenv2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Bookingtokenv2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bookingtokenv2{Bookingtokenv2Caller: Bookingtokenv2Caller{contract: contract}, Bookingtokenv2Transactor: Bookingtokenv2Transactor{contract: contract}, Bookingtokenv2Filterer: Bookingtokenv2Filterer{contract: contract}}, nil
}

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
