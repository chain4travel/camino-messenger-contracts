// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package servicefeetoken

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

// ServicefeetokenMetaData contains all meta data concerning the Servicefeetoken contract.
var ServicefeetokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newSymbol\",\"type\":\"string\"}],\"name\":\"reinitializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051612335620000fe6000396000818161108a015281816110b301526111fd01526123356000f3fe6080604052600436106101895760003560e01c806301ffc9a71461018e57806306fdde03146101c3578063095ea7b3146101e557806318160ddd1461020557806323b872dd14610228578063248a9ca3146102485780632f2ff15d14610268578063313ce5671461028a5780633644e515146102a657806336568abe146102bb5780633f4ba83a146102db57806340c10f19146102f057806342966c68146103105780634f1ef2861461033057806352d1902d146103435780635c975abb1461035857806370a082311461036d57806379cc67901461038d5780637ecebe00146103ad5780638456cb59146103cd57806384b0196e146103e257806391d148541461040a57806391da124c1461042a57806395d89b411461044a578063a217fddf1461045f578063a9059cbb14610474578063ad3cb1cc14610494578063d505accf146104c5578063d5391393146104e5578063d547741f14610507578063dd62ed3e14610527578063e63ab1e914610547578063f72c0d8b14610569578063f8c8765e1461058b575b600080fd5b34801561019a57600080fd5b506101ae6101a9366004611c03565b6105ab565b60405190151581526020015b60405180910390f35b3480156101cf57600080fd5b506101d86105e2565b6040516101ba9190611c7d565b3480156101f157600080fd5b506101ae610200366004611cac565b610683565b34801561021157600080fd5b5061021a61069b565b6040519081526020016101ba565b34801561023457600080fd5b506101ae610243366004611cd6565b6106b0565b34801561025457600080fd5b5061021a610263366004611d12565b6106d6565b34801561027457600080fd5b50610288610283366004611d2b565b6106f6565b005b34801561029657600080fd5b50604051601281526020016101ba565b3480156102b257600080fd5b5061021a610718565b3480156102c757600080fd5b506102886102d6366004611d2b565b610727565b3480156102e757600080fd5b5061028861075f565b3480156102fc57600080fd5b5061028861030b366004611cac565b610782565b34801561031c57600080fd5b5061028861032b366004611d12565b6107a4565b61028861033e366004611de2565b6107ae565b34801561034f57600080fd5b5061021a6107cd565b34801561036457600080fd5b506101ae6107ea565b34801561037957600080fd5b5061021a610388366004611e43565b6107ff565b34801561039957600080fd5b506102886103a8366004611cac565b61082a565b3480156103b957600080fd5b5061021a6103c8366004611e43565b61083f565b3480156103d957600080fd5b5061028861084a565b3480156103ee57600080fd5b506103f761086a565b6040516101ba9796959493929190611e5e565b34801561041657600080fd5b506101ae610425366004611d2b565b610918565b34801561043657600080fd5b50610288610445366004611f17565b61094e565b34801561045657600080fd5b506101d8610a0b565b34801561046b57600080fd5b5061021a600081565b34801561048057600080fd5b506101ae61048f366004611cac565b610a28565b3480156104a057600080fd5b506101d8604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156104d157600080fd5b506102886104e0366004611f70565b610a36565b3480156104f157600080fd5b5061021a6000805160206122e083398151915281565b34801561051357600080fd5b50610288610522366004611d2b565b610b52565b34801561053357600080fd5b5061021a610542366004611fe3565b610b6e565b34801561055357600080fd5b5061021a6000805160206122c083398151915281565b34801561057557600080fd5b5061021a60008051602061228083398151915281565b34801561059757600080fd5b506102886105a636600461200d565b610baa565b60006001600160e01b03198216637965db0b60e01b14806105dc57506301ffc9a760e01b6001600160e01b03198316145b92915050565b606060006105ee610d8e565b90508060030180546105ff90612061565b80601f016020809104026020016040519081016040528092919081815260200182805461062b90612061565b80156106785780601f1061064d57610100808354040283529160200191610678565b820191906000526020600020905b81548152906001019060200180831161065b57829003601f168201915b505050505091505090565b600033610691818585610db2565b5060019392505050565b6000806106a6610d8e565b6002015492915050565b6000336106be858285610dbf565b6106c9858585610e0c565b60019150505b9392505050565b6000806106e1610e6b565b60009384526020525050604090206001015490565b6106ff826106d6565b61070881610e8f565b6107128383610e99565b50505050565b6000610722610f3a565b905090565b6001600160a01b03811633146107505760405163334bd91960e11b815260040160405180910390fd5b61075a8282610f44565b505050565b6000805160206122c083398151915261077781610e8f565b61077f610fbc565b50565b6000805160206122e083398151915261079a81610e8f565b61075a8383611013565b61077f3382611049565b6107b661107f565b6107bf82611126565b6107c9828261113e565b5050565b60006107d76111f2565b506000805160206122a083398151915290565b6000806107f561123b565b5460ff1692915050565b60008061080a610d8e565b6001600160a01b0390931660009081526020939093525050604090205490565b610835823383610dbf565b6107c98282611049565b60006105dc8261125f565b6000805160206122c083398151915261086281610e8f565b61077f61126a565b60006060806000806000606060006108806112b1565b805490915015801561089457506001810154155b6108dd5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b6108e56112d5565b6108ed6112f2565b60408051600080825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b600080610923610e6b565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b6002600061095a6112fe565b8054909150600160401b900460ff1680610981575080546001600160401b03808416911610155b1561099f5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160481b0319166001600160401b03831617600160401b17815560006109c981610e8f565b6109d38585611322565b50805460ff60401b19168155604051600080516020612260833981519152906109fd90849061209b565b60405180910390a150505050565b60606000610a17610d8e565b90508060040180546105ff90612061565b600033610691818585610e0c565b83421115610a5a5760405163313c898160e11b8152600481018590526024016108d4565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610a898c611334565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610ae482611367565b90506000610af482878787611394565b9050896001600160a01b0316816001600160a01b031614610b3b576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016108d4565b610b468a8a8a610db2565b50505050505050505050565b610b5b826106d6565b610b6481610e8f565b6107128383610f44565b600080610b79610d8e565b6001600160a01b03948516600090815260019190910160209081526040808320959096168252939093525050205490565b6000610bb46112fe565b805490915060ff600160401b82041615906001600160401b0316600081158015610bdb5750825b90506000826001600160401b03166001148015610bf75750303b155b905081158015610c05575080155b15610c235760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315610c4c57845460ff60401b1916600160401b1785555b610ca3604051806040016040528060158152602001742aa9a21029b2b93b34b1b2902332b2902a37b5b2b760591b815250604051806040016040528060088152602001671554d10b9d195cdd60c21b815250611322565b610cab6113c2565b610cb36113ca565b610cbb6113c2565b610ceb6040518060400160405280600f81526020016e29b2b93b34b1b2a332b2aa37b5b2b760891b8152506113da565b610cf36113c2565b610cfe60008a610e99565b50610d176000805160206122c083398151915289610e99565b50610d306000805160206122e083398151915288610e99565b50610d4960008051602061228083398151915287610e99565b508315610d8357845460ff60401b1916855560405160008051602061226083398151915290610d7a9060019061209b565b60405180910390a15b505050505050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0090565b61075a8383836001611405565b6000610dcb8484610b6e565b905060001981146107125781811015610dfd57828183604051637dc7a0d960e11b81526004016108d4939291906120af565b61071284848484036000611405565b6001600160a01b038316610e36576000604051634b637e8f60e11b81526004016108d491906120d0565b6001600160a01b038216610e6057600060405163ec442f0560e01b81526004016108d491906120d0565b61075a8383836114ea565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b61077f81336114f5565b600080610ea4610e6b565b9050610eb08484610918565b610f30576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610ee63390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506105dc565b60009150506105dc565b600061072261152e565b600080610f4f610e6b565b9050610f5b8484610918565b15610f30576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506105dc565b610fc46115a2565b6000610fce61123b565b805460ff1916815590507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405161100891906120d0565b60405180910390a150565b6001600160a01b03821661103d57600060405163ec442f0560e01b81526004016108d491906120d0565b6107c9600083836114ea565b6001600160a01b038216611073576000604051634b637e8f60e11b81526004016108d491906120d0565b6107c9826000836114ea565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061110657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166110fa6000805160206122a0833981519152546001600160a01b031690565b6001600160a01b031614155b156111245760405163703e46dd60e11b815260040160405180910390fd5b565b6000805160206122808339815191526107c981610e8f565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611198575060408051601f3d908101601f19168201909252611195918101906120e4565b60015b6111b75781604051634c9c8ce360e01b81526004016108d491906120d0565b6000805160206122a083398151915281146111e857604051632a87526960e21b8152600481018290526024016108d4565b61075a83836115c7565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146111245760405163703e46dd60e11b815260040160405180910390fd5b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330090565b60008061080a61161d565b611272611641565b600061127c61123b565b805460ff1916600117815590507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610ffb3390565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10090565b606060006112e16112b1565b90508060020180546105ff90612061565b606060006105ee6112b1565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b61132a611667565b6107c9828261168c565b60008061133f61161d565b6001600160a01b03909316600090815260209390935250506040902080546001810190915590565b60006105dc611374610f3a565b8360405161190160f01b8152600281019290925260228201526042902090565b6000806000806113a6888888886116bd565b9250925092506113b68282611782565b50909695505050505050565b611124611667565b6113d2611667565b61112461183b565b6113e2611667565b61077f81604051806040016040528060018152602001603160f81b815250611858565b600061140f610d8e565b90506001600160a01b03851661143b57600060405163e602df0560e01b81526004016108d491906120d0565b6001600160a01b038416611465576000604051634a1406b160e11b81526004016108d491906120d0565b6001600160a01b038086166000908152600183016020908152604080832093881683529290522083905581156114e357836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516114da91815260200190565b60405180910390a35b5050505050565b61075a838383611899565b6114ff8282610918565b6107c95760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016108d4565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6115596118ac565b611561611913565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6115aa6107ea565b61112457604051638dfc202b60e01b815260040160405180910390fd5b6115d082611954565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156116155761075a82826119b0565b6107c9611a26565b7f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0090565b6116496107ea565b156111245760405163d93c066560e01b815260040160405180910390fd5b61166f611a45565b61112457604051631afcd79f60e31b815260040160405180910390fd5b611694611667565b600061169e610d8e565b9050600381016116ae848261214d565b5060048101610712838261214d565b600080806fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411156116ee5750600091506003905082611778565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611742573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661176e57506000925060019150829050611778565b9250600091508190505b9450945094915050565b60008260038111156117965761179661220c565b0361179f575050565b60018260038111156117b3576117b361220c565b036117d15760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156117e5576117e561220c565b036118065760405163fce698f760e01b8152600481018290526024016108d4565b600382600381111561181a5761181a61220c565b036107c9576040516335e2f38360e21b8152600481018290526024016108d4565b611843611667565b600061184d61123b565b805460ff1916905550565b611860611667565b600061186a6112b1565b90506002810161187a848261214d565b5060038101611889838261214d565b5060008082556001909101555050565b6118a1611641565b61075a838383611a5f565b6000806118b76112b1565b905060006118c36112d5565b8051909150156118db57805160209091012092915050565b815480156118ea579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b60008061191e6112b1565b9050600061192a6112f2565b80519091501561194257805160209091012092915050565b600182015480156118ea579392505050565b806001600160a01b03163b6000036119815780604051634c9c8ce360e01b81526004016108d491906120d0565b6000805160206122a083398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516119cd9190612222565b600060405180830381855af49150503d8060008114611a08576040519150601f19603f3d011682016040523d82523d6000602084013e611a0d565b606091505b5091509150611a1d858383611b87565b95945050505050565b34156111245760405163b398979f60e01b815260040160405180910390fd5b6000611a4f6112fe565b54600160401b900460ff16919050565b6000611a69610d8e565b90506001600160a01b038416611a985781816002016000828254611a8d919061223e565b90915550611af79050565b6001600160a01b03841660009081526020829052604090205482811015611ad85784818460405163391434e360e21b81526004016108d4939291906120af565b6001600160a01b03851660009081526020839052604090209083900390555b6001600160a01b038316611b15576002810180548390039055611b34565b6001600160a01b03831660009081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611b7991815260200190565b60405180910390a350505050565b606082611b9c57611b9782611bda565b6106cf565b8151158015611bb357506001600160a01b0384163b155b15611bd35783604051639996b31560e01b81526004016108d491906120d0565b50806106cf565b805115611bea5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b600060208284031215611c1557600080fd5b81356001600160e01b0319811681146106cf57600080fd5b60005b83811015611c48578181015183820152602001611c30565b50506000910152565b60008151808452611c69816020860160208601611c2d565b601f01601f19169290920160200192915050565b6020815260006106cf6020830184611c51565b80356001600160a01b0381168114611ca757600080fd5b919050565b60008060408385031215611cbf57600080fd5b611cc883611c90565b946020939093013593505050565b600080600060608486031215611ceb57600080fd5b611cf484611c90565b9250611d0260208501611c90565b9150604084013590509250925092565b600060208284031215611d2457600080fd5b5035919050565b60008060408385031215611d3e57600080fd5b82359150611d4e60208401611c90565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b0380841115611d8757611d87611d57565b604051601f8501601f19908116603f01168101908282118183101715611daf57611daf611d57565b81604052809350858152868686011115611dc857600080fd5b858560208301376000602087830101525050509392505050565b60008060408385031215611df557600080fd5b611dfe83611c90565b915060208301356001600160401b03811115611e1957600080fd5b8301601f81018513611e2a57600080fd5b611e3985823560208401611d6d565b9150509250929050565b600060208284031215611e5557600080fd5b6106cf82611c90565b60ff60f81b881681526000602060e06020840152611e7f60e084018a611c51565b8381036040850152611e91818a611c51565b606085018990526001600160a01b038816608086015260a0850187905284810360c08601528551808252602080880193509091019060005b81811015611ee557835183529284019291840191600101611ec9565b50909c9b505050505050505050505050565b600082601f830112611f0857600080fd5b6106cf83833560208501611d6d565b60008060408385031215611f2a57600080fd5b82356001600160401b0380821115611f4157600080fd5b611f4d86838701611ef7565b93506020850135915080821115611f6357600080fd5b50611e3985828601611ef7565b600080600080600080600060e0888a031215611f8b57600080fd5b611f9488611c90565b9650611fa260208901611c90565b95506040880135945060608801359350608088013560ff81168114611fc657600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611ff657600080fd5b611fff83611c90565b9150611d4e60208401611c90565b6000806000806080858703121561202357600080fd5b61202c85611c90565b935061203a60208601611c90565b925061204860408601611c90565b915061205660608601611c90565b905092959194509250565b600181811c9082168061207557607f821691505b60208210810361209557634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160401b0391909116815260200190565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0391909116815260200190565b6000602082840312156120f657600080fd5b5051919050565b601f82111561075a576000816000526020600020601f850160051c810160208610156121265750805b601f850160051c820191505b8181101561214557828155600101612132565b505050505050565b81516001600160401b0381111561216657612166611d57565b61217a816121748454612061565b846120fd565b602080601f8311600181146121af57600084156121975750858301515b600019600386901b1c1916600185901b178555612145565b600085815260208120601f198616915b828110156121de578886015182559484019460019091019084016121bf565b50858210156121fc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60008251612234818460208701611c2d565b9190910192915050565b808201808211156105dc57634e487b7160e01b600052601160045260246000fdfec7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6a26469706673582212206a3834dc34c4bd210f776d725f7e817ec554c4f979d6bbafe4fe6c642aeeab6f64736f6c63430008180033",
}

// ServicefeetokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ServicefeetokenMetaData.ABI instead.
var ServicefeetokenABI = ServicefeetokenMetaData.ABI

// ServicefeetokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ServicefeetokenMetaData.Bin instead.
var ServicefeetokenBin = ServicefeetokenMetaData.Bin

// DeployServicefeetoken deploys a new Ethereum contract, binding an instance of Servicefeetoken to it.
func DeployServicefeetoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Servicefeetoken, error) {
	parsed, err := ServicefeetokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ServicefeetokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Servicefeetoken{ServicefeetokenCaller: ServicefeetokenCaller{contract: contract}, ServicefeetokenTransactor: ServicefeetokenTransactor{contract: contract}, ServicefeetokenFilterer: ServicefeetokenFilterer{contract: contract}}, nil
}

// Servicefeetoken is an auto generated Go binding around an Ethereum contract.
type Servicefeetoken struct {
	ServicefeetokenCaller     // Read-only binding to the contract
	ServicefeetokenTransactor // Write-only binding to the contract
	ServicefeetokenFilterer   // Log filterer for contract events
}

// ServicefeetokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServicefeetokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServicefeetokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServicefeetokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServicefeetokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServicefeetokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServicefeetokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServicefeetokenSession struct {
	Contract     *Servicefeetoken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServicefeetokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServicefeetokenCallerSession struct {
	Contract *ServicefeetokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ServicefeetokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServicefeetokenTransactorSession struct {
	Contract     *ServicefeetokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ServicefeetokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServicefeetokenRaw struct {
	Contract *Servicefeetoken // Generic contract binding to access the raw methods on
}

// ServicefeetokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServicefeetokenCallerRaw struct {
	Contract *ServicefeetokenCaller // Generic read-only contract binding to access the raw methods on
}

// ServicefeetokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServicefeetokenTransactorRaw struct {
	Contract *ServicefeetokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServicefeetoken creates a new instance of Servicefeetoken, bound to a specific deployed contract.
func NewServicefeetoken(address common.Address, backend bind.ContractBackend) (*Servicefeetoken, error) {
	contract, err := bindServicefeetoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Servicefeetoken{ServicefeetokenCaller: ServicefeetokenCaller{contract: contract}, ServicefeetokenTransactor: ServicefeetokenTransactor{contract: contract}, ServicefeetokenFilterer: ServicefeetokenFilterer{contract: contract}}, nil
}

// NewServicefeetokenCaller creates a new read-only instance of Servicefeetoken, bound to a specific deployed contract.
func NewServicefeetokenCaller(address common.Address, caller bind.ContractCaller) (*ServicefeetokenCaller, error) {
	contract, err := bindServicefeetoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenCaller{contract: contract}, nil
}

// NewServicefeetokenTransactor creates a new write-only instance of Servicefeetoken, bound to a specific deployed contract.
func NewServicefeetokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ServicefeetokenTransactor, error) {
	contract, err := bindServicefeetoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenTransactor{contract: contract}, nil
}

// NewServicefeetokenFilterer creates a new log filterer instance of Servicefeetoken, bound to a specific deployed contract.
func NewServicefeetokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ServicefeetokenFilterer, error) {
	contract, err := bindServicefeetoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenFilterer{contract: contract}, nil
}

// bindServicefeetoken binds a generic wrapper to an already deployed contract.
func bindServicefeetoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ServicefeetokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Servicefeetoken *ServicefeetokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Servicefeetoken.Contract.ServicefeetokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Servicefeetoken *ServicefeetokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.ServicefeetokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Servicefeetoken *ServicefeetokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.ServicefeetokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Servicefeetoken *ServicefeetokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Servicefeetoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Servicefeetoken *ServicefeetokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Servicefeetoken *ServicefeetokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.DEFAULTADMINROLE(&_Servicefeetoken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.DEFAULTADMINROLE(&_Servicefeetoken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Servicefeetoken.Contract.DOMAINSEPARATOR(&_Servicefeetoken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Servicefeetoken.Contract.DOMAINSEPARATOR(&_Servicefeetoken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenSession) MINTERROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.MINTERROLE(&_Servicefeetoken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.MINTERROLE(&_Servicefeetoken.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenSession) PAUSERROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.PAUSERROLE(&_Servicefeetoken.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.PAUSERROLE(&_Servicefeetoken.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenSession) UPGRADERROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.UPGRADERROLE(&_Servicefeetoken.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Servicefeetoken.Contract.UPGRADERROLE(&_Servicefeetoken.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Servicefeetoken *ServicefeetokenCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Servicefeetoken *ServicefeetokenSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Servicefeetoken.Contract.UPGRADEINTERFACEVERSION(&_Servicefeetoken.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Servicefeetoken *ServicefeetokenCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Servicefeetoken.Contract.UPGRADEINTERFACEVERSION(&_Servicefeetoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Servicefeetoken.Contract.Allowance(&_Servicefeetoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Servicefeetoken.Contract.Allowance(&_Servicefeetoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Servicefeetoken.Contract.BalanceOf(&_Servicefeetoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Servicefeetoken.Contract.BalanceOf(&_Servicefeetoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Servicefeetoken *ServicefeetokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Servicefeetoken *ServicefeetokenSession) Decimals() (uint8, error) {
	return _Servicefeetoken.Contract.Decimals(&_Servicefeetoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Servicefeetoken *ServicefeetokenCallerSession) Decimals() (uint8, error) {
	return _Servicefeetoken.Contract.Decimals(&_Servicefeetoken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Servicefeetoken *ServicefeetokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Servicefeetoken *ServicefeetokenSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Servicefeetoken.Contract.Eip712Domain(&_Servicefeetoken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Servicefeetoken *ServicefeetokenCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Servicefeetoken.Contract.Eip712Domain(&_Servicefeetoken.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Servicefeetoken.Contract.GetRoleAdmin(&_Servicefeetoken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Servicefeetoken.Contract.GetRoleAdmin(&_Servicefeetoken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Servicefeetoken *ServicefeetokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Servicefeetoken *ServicefeetokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Servicefeetoken.Contract.HasRole(&_Servicefeetoken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Servicefeetoken *ServicefeetokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Servicefeetoken.Contract.HasRole(&_Servicefeetoken.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Servicefeetoken *ServicefeetokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Servicefeetoken *ServicefeetokenSession) Name() (string, error) {
	return _Servicefeetoken.Contract.Name(&_Servicefeetoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Servicefeetoken *ServicefeetokenCallerSession) Name() (string, error) {
	return _Servicefeetoken.Contract.Name(&_Servicefeetoken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Servicefeetoken.Contract.Nonces(&_Servicefeetoken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Servicefeetoken.Contract.Nonces(&_Servicefeetoken.CallOpts, owner)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Servicefeetoken *ServicefeetokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Servicefeetoken *ServicefeetokenSession) Paused() (bool, error) {
	return _Servicefeetoken.Contract.Paused(&_Servicefeetoken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Servicefeetoken *ServicefeetokenCallerSession) Paused() (bool, error) {
	return _Servicefeetoken.Contract.Paused(&_Servicefeetoken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenSession) ProxiableUUID() ([32]byte, error) {
	return _Servicefeetoken.Contract.ProxiableUUID(&_Servicefeetoken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Servicefeetoken *ServicefeetokenCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Servicefeetoken.Contract.ProxiableUUID(&_Servicefeetoken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Servicefeetoken *ServicefeetokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Servicefeetoken *ServicefeetokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Servicefeetoken.Contract.SupportsInterface(&_Servicefeetoken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Servicefeetoken *ServicefeetokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Servicefeetoken.Contract.SupportsInterface(&_Servicefeetoken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Servicefeetoken *ServicefeetokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Servicefeetoken *ServicefeetokenSession) Symbol() (string, error) {
	return _Servicefeetoken.Contract.Symbol(&_Servicefeetoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Servicefeetoken *ServicefeetokenCallerSession) Symbol() (string, error) {
	return _Servicefeetoken.Contract.Symbol(&_Servicefeetoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Servicefeetoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Servicefeetoken *ServicefeetokenSession) TotalSupply() (*big.Int, error) {
	return _Servicefeetoken.Contract.TotalSupply(&_Servicefeetoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Servicefeetoken *ServicefeetokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Servicefeetoken.Contract.TotalSupply(&_Servicefeetoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Approve(&_Servicefeetoken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Approve(&_Servicefeetoken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_Servicefeetoken *ServicefeetokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Burn(&_Servicefeetoken.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Burn(&_Servicefeetoken.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_Servicefeetoken *ServicefeetokenSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.BurnFrom(&_Servicefeetoken.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.BurnFrom(&_Servicefeetoken.TransactOpts, account, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Servicefeetoken *ServicefeetokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.GrantRole(&_Servicefeetoken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.GrantRole(&_Servicefeetoken.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address defaultAdmin, address pauser, address minter, address upgrader) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) Initialize(opts *bind.TransactOpts, defaultAdmin common.Address, pauser common.Address, minter common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "initialize", defaultAdmin, pauser, minter, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address defaultAdmin, address pauser, address minter, address upgrader) returns()
func (_Servicefeetoken *ServicefeetokenSession) Initialize(defaultAdmin common.Address, pauser common.Address, minter common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Initialize(&_Servicefeetoken.TransactOpts, defaultAdmin, pauser, minter, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address defaultAdmin, address pauser, address minter, address upgrader) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) Initialize(defaultAdmin common.Address, pauser common.Address, minter common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Initialize(&_Servicefeetoken.TransactOpts, defaultAdmin, pauser, minter, upgrader)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Servicefeetoken *ServicefeetokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Mint(&_Servicefeetoken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Mint(&_Servicefeetoken.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Servicefeetoken *ServicefeetokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Servicefeetoken *ServicefeetokenSession) Pause() (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Pause(&_Servicefeetoken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Pause(&_Servicefeetoken.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Servicefeetoken *ServicefeetokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Permit(&_Servicefeetoken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Permit(&_Servicefeetoken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// ReinitializeV2 is a paid mutator transaction binding the contract method 0x91da124c.
//
// Solidity: function reinitializeV2(string newName, string newSymbol) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) ReinitializeV2(opts *bind.TransactOpts, newName string, newSymbol string) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "reinitializeV2", newName, newSymbol)
}

// ReinitializeV2 is a paid mutator transaction binding the contract method 0x91da124c.
//
// Solidity: function reinitializeV2(string newName, string newSymbol) returns()
func (_Servicefeetoken *ServicefeetokenSession) ReinitializeV2(newName string, newSymbol string) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.ReinitializeV2(&_Servicefeetoken.TransactOpts, newName, newSymbol)
}

// ReinitializeV2 is a paid mutator transaction binding the contract method 0x91da124c.
//
// Solidity: function reinitializeV2(string newName, string newSymbol) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) ReinitializeV2(newName string, newSymbol string) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.ReinitializeV2(&_Servicefeetoken.TransactOpts, newName, newSymbol)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Servicefeetoken *ServicefeetokenSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.RenounceRole(&_Servicefeetoken.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.RenounceRole(&_Servicefeetoken.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Servicefeetoken *ServicefeetokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Servicefeetoken *ServicefeetokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.RevokeRole(&_Servicefeetoken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.RevokeRole(&_Servicefeetoken.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Transfer(&_Servicefeetoken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Transfer(&_Servicefeetoken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.TransferFrom(&_Servicefeetoken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Servicefeetoken *ServicefeetokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.TransferFrom(&_Servicefeetoken.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Servicefeetoken *ServicefeetokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Servicefeetoken *ServicefeetokenSession) Unpause() (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Unpause(&_Servicefeetoken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Servicefeetoken.Contract.Unpause(&_Servicefeetoken.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Servicefeetoken *ServicefeetokenTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Servicefeetoken.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Servicefeetoken *ServicefeetokenSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.UpgradeToAndCall(&_Servicefeetoken.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Servicefeetoken *ServicefeetokenTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Servicefeetoken.Contract.UpgradeToAndCall(&_Servicefeetoken.TransactOpts, newImplementation, data)
}

// ServicefeetokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Servicefeetoken contract.
type ServicefeetokenApprovalIterator struct {
	Event *ServicefeetokenApproval // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenApproval)
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
		it.Event = new(ServicefeetokenApproval)
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
func (it *ServicefeetokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenApproval represents a Approval event raised by the Servicefeetoken contract.
type ServicefeetokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ServicefeetokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenApprovalIterator{contract: _Servicefeetoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ServicefeetokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenApproval)
				if err := _Servicefeetoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Servicefeetoken *ServicefeetokenFilterer) ParseApproval(log types.Log) (*ServicefeetokenApproval, error) {
	event := new(ServicefeetokenApproval)
	if err := _Servicefeetoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Servicefeetoken contract.
type ServicefeetokenEIP712DomainChangedIterator struct {
	Event *ServicefeetokenEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenEIP712DomainChanged)
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
		it.Event = new(ServicefeetokenEIP712DomainChanged)
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
func (it *ServicefeetokenEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the Servicefeetoken contract.
type ServicefeetokenEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Servicefeetoken *ServicefeetokenFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ServicefeetokenEIP712DomainChangedIterator, error) {

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenEIP712DomainChangedIterator{contract: _Servicefeetoken.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Servicefeetoken *ServicefeetokenFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ServicefeetokenEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenEIP712DomainChanged)
				if err := _Servicefeetoken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Servicefeetoken *ServicefeetokenFilterer) ParseEIP712DomainChanged(log types.Log) (*ServicefeetokenEIP712DomainChanged, error) {
	event := new(ServicefeetokenEIP712DomainChanged)
	if err := _Servicefeetoken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Servicefeetoken contract.
type ServicefeetokenInitializedIterator struct {
	Event *ServicefeetokenInitialized // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenInitialized)
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
		it.Event = new(ServicefeetokenInitialized)
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
func (it *ServicefeetokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenInitialized represents a Initialized event raised by the Servicefeetoken contract.
type ServicefeetokenInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*ServicefeetokenInitializedIterator, error) {

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenInitializedIterator{contract: _Servicefeetoken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ServicefeetokenInitialized) (event.Subscription, error) {

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenInitialized)
				if err := _Servicefeetoken.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Servicefeetoken *ServicefeetokenFilterer) ParseInitialized(log types.Log) (*ServicefeetokenInitialized, error) {
	event := new(ServicefeetokenInitialized)
	if err := _Servicefeetoken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Servicefeetoken contract.
type ServicefeetokenPausedIterator struct {
	Event *ServicefeetokenPaused // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenPaused)
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
		it.Event = new(ServicefeetokenPaused)
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
func (it *ServicefeetokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenPaused represents a Paused event raised by the Servicefeetoken contract.
type ServicefeetokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterPaused(opts *bind.FilterOpts) (*ServicefeetokenPausedIterator, error) {

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenPausedIterator{contract: _Servicefeetoken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ServicefeetokenPaused) (event.Subscription, error) {

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenPaused)
				if err := _Servicefeetoken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Servicefeetoken *ServicefeetokenFilterer) ParsePaused(log types.Log) (*ServicefeetokenPaused, error) {
	event := new(ServicefeetokenPaused)
	if err := _Servicefeetoken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Servicefeetoken contract.
type ServicefeetokenRoleAdminChangedIterator struct {
	Event *ServicefeetokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenRoleAdminChanged)
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
		it.Event = new(ServicefeetokenRoleAdminChanged)
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
func (it *ServicefeetokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenRoleAdminChanged represents a RoleAdminChanged event raised by the Servicefeetoken contract.
type ServicefeetokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ServicefeetokenRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenRoleAdminChangedIterator{contract: _Servicefeetoken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ServicefeetokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenRoleAdminChanged)
				if err := _Servicefeetoken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Servicefeetoken *ServicefeetokenFilterer) ParseRoleAdminChanged(log types.Log) (*ServicefeetokenRoleAdminChanged, error) {
	event := new(ServicefeetokenRoleAdminChanged)
	if err := _Servicefeetoken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Servicefeetoken contract.
type ServicefeetokenRoleGrantedIterator struct {
	Event *ServicefeetokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenRoleGranted)
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
		it.Event = new(ServicefeetokenRoleGranted)
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
func (it *ServicefeetokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenRoleGranted represents a RoleGranted event raised by the Servicefeetoken contract.
type ServicefeetokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ServicefeetokenRoleGrantedIterator, error) {

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

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenRoleGrantedIterator{contract: _Servicefeetoken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ServicefeetokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenRoleGranted)
				if err := _Servicefeetoken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Servicefeetoken *ServicefeetokenFilterer) ParseRoleGranted(log types.Log) (*ServicefeetokenRoleGranted, error) {
	event := new(ServicefeetokenRoleGranted)
	if err := _Servicefeetoken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Servicefeetoken contract.
type ServicefeetokenRoleRevokedIterator struct {
	Event *ServicefeetokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenRoleRevoked)
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
		it.Event = new(ServicefeetokenRoleRevoked)
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
func (it *ServicefeetokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenRoleRevoked represents a RoleRevoked event raised by the Servicefeetoken contract.
type ServicefeetokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ServicefeetokenRoleRevokedIterator, error) {

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

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenRoleRevokedIterator{contract: _Servicefeetoken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ServicefeetokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenRoleRevoked)
				if err := _Servicefeetoken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Servicefeetoken *ServicefeetokenFilterer) ParseRoleRevoked(log types.Log) (*ServicefeetokenRoleRevoked, error) {
	event := new(ServicefeetokenRoleRevoked)
	if err := _Servicefeetoken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Servicefeetoken contract.
type ServicefeetokenTransferIterator struct {
	Event *ServicefeetokenTransfer // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenTransfer)
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
		it.Event = new(ServicefeetokenTransfer)
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
func (it *ServicefeetokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenTransfer represents a Transfer event raised by the Servicefeetoken contract.
type ServicefeetokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ServicefeetokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenTransferIterator{contract: _Servicefeetoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ServicefeetokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenTransfer)
				if err := _Servicefeetoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Servicefeetoken *ServicefeetokenFilterer) ParseTransfer(log types.Log) (*ServicefeetokenTransfer, error) {
	event := new(ServicefeetokenTransfer)
	if err := _Servicefeetoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Servicefeetoken contract.
type ServicefeetokenUnpausedIterator struct {
	Event *ServicefeetokenUnpaused // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenUnpaused)
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
		it.Event = new(ServicefeetokenUnpaused)
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
func (it *ServicefeetokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenUnpaused represents a Unpaused event raised by the Servicefeetoken contract.
type ServicefeetokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ServicefeetokenUnpausedIterator, error) {

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenUnpausedIterator{contract: _Servicefeetoken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ServicefeetokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenUnpaused)
				if err := _Servicefeetoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Servicefeetoken *ServicefeetokenFilterer) ParseUnpaused(log types.Log) (*ServicefeetokenUnpaused, error) {
	event := new(ServicefeetokenUnpaused)
	if err := _Servicefeetoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServicefeetokenUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Servicefeetoken contract.
type ServicefeetokenUpgradedIterator struct {
	Event *ServicefeetokenUpgraded // Event containing the contract specifics and raw log

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
func (it *ServicefeetokenUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServicefeetokenUpgraded)
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
		it.Event = new(ServicefeetokenUpgraded)
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
func (it *ServicefeetokenUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServicefeetokenUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServicefeetokenUpgraded represents a Upgraded event raised by the Servicefeetoken contract.
type ServicefeetokenUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Servicefeetoken *ServicefeetokenFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ServicefeetokenUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Servicefeetoken.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ServicefeetokenUpgradedIterator{contract: _Servicefeetoken.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Servicefeetoken *ServicefeetokenFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ServicefeetokenUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Servicefeetoken.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServicefeetokenUpgraded)
				if err := _Servicefeetoken.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Servicefeetoken *ServicefeetokenFilterer) ParseUpgraded(log types.Log) (*ServicefeetokenUpgraded, error) {
	event := new(ServicefeetokenUpgraded)
	if err := _Servicefeetoken.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
