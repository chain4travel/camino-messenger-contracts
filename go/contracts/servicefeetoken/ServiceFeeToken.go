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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051612198620000fe60003960008181610faf01528181610fd8015261112201526121986000f3fe60806040526004361061017e5760003560e01c806301ffc9a71461018357806306fdde03146101b8578063095ea7b3146101da57806318160ddd146101fa57806323b872dd1461021d578063248a9ca31461023d5780632f2ff15d1461025d578063313ce5671461027f5780633644e5151461029b57806336568abe146102b05780633f4ba83a146102d057806340c10f19146102e557806342966c68146103055780634f1ef2861461032557806352d1902d146103385780635c975abb1461034d57806370a082311461036257806379cc6790146103825780637ecebe00146103a25780638456cb59146103c257806384b0196e146103d757806391d14854146103ff57806395d89b411461041f578063a217fddf14610434578063a9059cbb14610449578063ad3cb1cc14610469578063d505accf1461049a578063d5391393146104ba578063d547741f146104dc578063dd62ed3e146104fc578063e63ab1e91461051c578063f72c0d8b1461053e578063f8c8765e14610560575b600080fd5b34801561018f57600080fd5b506101a361019e366004611b28565b610580565b60405190151581526020015b60405180910390f35b3480156101c457600080fd5b506101cd6105b7565b6040516101af9190611ba2565b3480156101e657600080fd5b506101a36101f5366004611bd1565b610658565b34801561020657600080fd5b5061020f610670565b6040519081526020016101af565b34801561022957600080fd5b506101a3610238366004611bfb565b610685565b34801561024957600080fd5b5061020f610258366004611c37565b6106ab565b34801561026957600080fd5b5061027d610278366004611c50565b6106cb565b005b34801561028b57600080fd5b50604051601281526020016101af565b3480156102a757600080fd5b5061020f6106ed565b3480156102bc57600080fd5b5061027d6102cb366004611c50565b6106fc565b3480156102dc57600080fd5b5061027d610734565b3480156102f157600080fd5b5061027d610300366004611bd1565b610757565b34801561031157600080fd5b5061027d610320366004611c37565b610779565b61027d610333366004611c92565b610783565b34801561034457600080fd5b5061020f6107a2565b34801561035957600080fd5b506101a36107bf565b34801561036e57600080fd5b5061020f61037d366004611d53565b6107d4565b34801561038e57600080fd5b5061027d61039d366004611bd1565b6107ff565b3480156103ae57600080fd5b5061020f6103bd366004611d53565b610814565b3480156103ce57600080fd5b5061027d61081f565b3480156103e357600080fd5b506103ec61083f565b6040516101af9796959493929190611d6e565b34801561040b57600080fd5b506101a361041a366004611c50565b6108ed565b34801561042b57600080fd5b506101cd610923565b34801561044057600080fd5b5061020f600081565b34801561045557600080fd5b506101a3610464366004611bd1565b610940565b34801561047557600080fd5b506101cd604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156104a657600080fd5b5061027d6104b5366004611e07565b61094e565b3480156104c657600080fd5b5061020f60008051602061214383398151915281565b3480156104e857600080fd5b5061027d6104f7366004611c50565b610a6a565b34801561050857600080fd5b5061020f610517366004611e7a565b610a86565b34801561052857600080fd5b5061020f60008051602061212383398151915281565b34801561054a57600080fd5b5061020f6000805160206120e383398151915281565b34801561056c57600080fd5b5061027d61057b366004611ea4565b610ac2565b60006001600160e01b03198216637965db0b60e01b14806105b157506301ffc9a760e01b6001600160e01b03198316145b92915050565b606060006105c3610cb3565b90508060030180546105d490611ef8565b80601f016020809104026020016040519081016040528092919081815260200182805461060090611ef8565b801561064d5780601f106106225761010080835404028352916020019161064d565b820191906000526020600020905b81548152906001019060200180831161063057829003601f168201915b505050505091505090565b600033610666818585610cd7565b5060019392505050565b60008061067b610cb3565b6002015492915050565b600033610693858285610ce4565b61069e858585610d31565b60019150505b9392505050565b6000806106b6610d90565b60009384526020525050604090206001015490565b6106d4826106ab565b6106dd81610db4565b6106e78383610dbe565b50505050565b60006106f7610e5f565b905090565b6001600160a01b03811633146107255760405163334bd91960e11b815260040160405180910390fd5b61072f8282610e69565b505050565b60008051602061212383398151915261074c81610db4565b610754610ee1565b50565b60008051602061214383398151915261076f81610db4565b61072f8383610f38565b6107543382610f6e565b61078b610fa4565b6107948261104b565b61079e8282611063565b5050565b60006107ac611117565b5060008051602061210383398151915290565b6000806107ca611160565b5460ff1692915050565b6000806107df610cb3565b6001600160a01b0390931660009081526020939093525050604090205490565b61080a823383610ce4565b61079e8282610f6e565b60006105b182611184565b60008051602061212383398151915261083781610db4565b61075461118f565b60006060806000806000606060006108556111d6565b805490915015801561086957506001810154155b6108b25760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b6108ba6111fa565b6108c2611217565b60408051600080825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6000806108f8610d90565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b6060600061092f610cb3565b90508060040180546105d490611ef8565b600033610666818585610d31565b834211156109725760405163313c898160e11b8152600481018590526024016108a9565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886109a18c611223565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e00160405160208183030381529060405280519060200120905060006109fc82611256565b90506000610a0c82878787611283565b9050896001600160a01b0316816001600160a01b031614610a53576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016108a9565b610a5e8a8a8a610cd7565b50505050505050505050565b610a73826106ab565b610a7c81610db4565b6106e78383610e69565b600080610a91610cb3565b6001600160a01b03948516600090815260019190910160209081526040808320959096168252939093525050205490565b6000610acc6112b1565b805490915060ff600160401b82041615906001600160401b0316600081158015610af35750825b90506000826001600160401b03166001148015610b0f5750303b155b905081158015610b1d575080155b15610b3b5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315610b6457845460ff60401b1916600160401b1785555b610bbb604051806040016040528060158152602001742aa9a21029b2b93b34b1b2902332b2902a37b5b2b760591b815250604051806040016040528060088152602001671554d10b9d195cdd60c21b8152506112d5565b610bc36112e7565b610bcb6112ef565b610bd36112e7565b610c036040518060400160405280600f81526020016e29b2b93b34b1b2a332b2aa37b5b2b760891b8152506112ff565b610c0b6112e7565b610c1660008a610dbe565b50610c2f60008051602061212383398151915289610dbe565b50610c4860008051602061214383398151915288610dbe565b50610c616000805160206120e383398151915287610dbe565b508315610ca857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0090565b61072f838383600161132a565b6000610cf08484610a86565b905060001981146106e75781811015610d2257828183604051637dc7a0d960e11b81526004016108a993929190611f32565b6106e78484848403600061132a565b6001600160a01b038316610d5b576000604051634b637e8f60e11b81526004016108a99190611f53565b6001600160a01b038216610d8557600060405163ec442f0560e01b81526004016108a99190611f53565b61072f83838361140f565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b610754813361141a565b600080610dc9610d90565b9050610dd584846108ed565b610e55576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610e0b3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506105b1565b60009150506105b1565b60006106f7611453565b600080610e74610d90565b9050610e8084846108ed565b15610e55576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506105b1565b610ee96114c7565b6000610ef3611160565b805460ff1916815590507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b604051610f2d9190611f53565b60405180910390a150565b6001600160a01b038216610f6257600060405163ec442f0560e01b81526004016108a99190611f53565b61079e6000838361140f565b6001600160a01b038216610f98576000604051634b637e8f60e11b81526004016108a99190611f53565b61079e8260008361140f565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061102b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661101f600080516020612103833981519152546001600160a01b031690565b6001600160a01b031614155b156110495760405163703e46dd60e11b815260040160405180910390fd5b565b6000805160206120e383398151915261079e81610db4565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156110bd575060408051601f3d908101601f191682019092526110ba91810190611f67565b60015b6110dc5781604051634c9c8ce360e01b81526004016108a99190611f53565b600080516020612103833981519152811461110d57604051632a87526960e21b8152600481018290526024016108a9565b61072f83836114ec565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146110495760405163703e46dd60e11b815260040160405180910390fd5b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330090565b6000806107df611542565b611197611566565b60006111a1611160565b805460ff1916600117815590507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610f203390565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10090565b606060006112066111d6565b90508060020180546105d490611ef8565b606060006105c36111d6565b60008061122e611542565b6001600160a01b03909316600090815260209390935250506040902080546001810190915590565b60006105b1611263610e5f565b8360405161190160f01b8152600281019290925260228201526042902090565b6000806000806112958888888861158c565b9250925092506112a58282611651565b50909695505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b6112dd61170a565b61079e828261172f565b61104961170a565b6112f761170a565b611049611760565b61130761170a565b61075481604051806040016040528060018152602001603160f81b81525061177d565b6000611334610cb3565b90506001600160a01b03851661136057600060405163e602df0560e01b81526004016108a99190611f53565b6001600160a01b03841661138a576000604051634a1406b160e11b81526004016108a99190611f53565b6001600160a01b0380861660009081526001830160209081526040808320938816835292905220839055811561140857836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516113ff91815260200190565b60405180910390a35b5050505050565b61072f8383836117be565b61142482826108ed565b61079e5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016108a9565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61147e6117d1565b611486611838565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6114cf6107bf565b61104957604051638dfc202b60e01b815260040160405180910390fd5b6114f582611879565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561153a5761072f82826118d5565b61079e61194b565b7f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb0090565b61156e6107bf565b156110495760405163d93c066560e01b815260040160405180910390fd5b600080806fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411156115bd5750600091506003905082611647565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611611573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661163d57506000925060019150829050611647565b9250600091508190505b9450945094915050565b600082600381111561166557611665611f80565b0361166e575050565b600182600381111561168257611682611f80565b036116a05760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156116b4576116b4611f80565b036116d55760405163fce698f760e01b8152600481018290526024016108a9565b60038260038111156116e9576116e9611f80565b0361079e576040516335e2f38360e21b8152600481018290526024016108a9565b61171261196a565b61104957604051631afcd79f60e31b815260040160405180910390fd5b61173761170a565b6000611741610cb3565b9050600381016117518482611fe6565b50600481016106e78382611fe6565b61176861170a565b6000611772611160565b805460ff1916905550565b61178561170a565b600061178f6111d6565b90506002810161179f8482611fe6565b50600381016117ae8382611fe6565b5060008082556001909101555050565b6117c6611566565b61072f838383611984565b6000806117dc6111d6565b905060006117e86111fa565b80519091501561180057805160209091012092915050565b8154801561180f579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b6000806118436111d6565b9050600061184f611217565b80519091501561186757805160209091012092915050565b6001820154801561180f579392505050565b806001600160a01b03163b6000036118a65780604051634c9c8ce360e01b81526004016108a99190611f53565b60008051602061210383398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516118f291906120a5565b600060405180830381855af49150503d806000811461192d576040519150601f19603f3d011682016040523d82523d6000602084013e611932565b606091505b5091509150611942858383611aac565b95945050505050565b34156110495760405163b398979f60e01b815260040160405180910390fd5b60006119746112b1565b54600160401b900460ff16919050565b600061198e610cb3565b90506001600160a01b0384166119bd57818160020160008282546119b291906120c1565b90915550611a1c9050565b6001600160a01b038416600090815260208290526040902054828110156119fd5784818460405163391434e360e21b81526004016108a993929190611f32565b6001600160a01b03851660009081526020839052604090209083900390555b6001600160a01b038316611a3a576002810180548390039055611a59565b6001600160a01b03831660009081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611a9e91815260200190565b60405180910390a350505050565b606082611ac157611abc82611aff565b6106a4565b8151158015611ad857506001600160a01b0384163b155b15611af85783604051639996b31560e01b81526004016108a99190611f53565b50806106a4565b805115611b0f5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b600060208284031215611b3a57600080fd5b81356001600160e01b0319811681146106a457600080fd5b60005b83811015611b6d578181015183820152602001611b55565b50506000910152565b60008151808452611b8e816020860160208601611b52565b601f01601f19169290920160200192915050565b6020815260006106a46020830184611b76565b80356001600160a01b0381168114611bcc57600080fd5b919050565b60008060408385031215611be457600080fd5b611bed83611bb5565b946020939093013593505050565b600080600060608486031215611c1057600080fd5b611c1984611bb5565b9250611c2760208501611bb5565b9150604084013590509250925092565b600060208284031215611c4957600080fd5b5035919050565b60008060408385031215611c6357600080fd5b82359150611c7360208401611bb5565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611ca557600080fd5b611cae83611bb5565b915060208301356001600160401b0380821115611cca57600080fd5b818501915085601f830112611cde57600080fd5b813581811115611cf057611cf0611c7c565b604051601f8201601f19908116603f01168101908382118183101715611d1857611d18611c7c565b81604052828152886020848701011115611d3157600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600060208284031215611d6557600080fd5b6106a482611bb5565b60ff60f81b881681526000602060e06020840152611d8f60e084018a611b76565b8381036040850152611da1818a611b76565b606085018990526001600160a01b038816608086015260a0850187905284810360c08601528551808252602080880193509091019060005b81811015611df557835183529284019291840191600101611dd9565b50909c9b505050505050505050505050565b600080600080600080600060e0888a031215611e2257600080fd5b611e2b88611bb5565b9650611e3960208901611bb5565b95506040880135945060608801359350608088013560ff81168114611e5d57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611e8d57600080fd5b611e9683611bb5565b9150611c7360208401611bb5565b60008060008060808587031215611eba57600080fd5b611ec385611bb5565b9350611ed160208601611bb5565b9250611edf60408601611bb5565b9150611eed60608601611bb5565b905092959194509250565b600181811c90821680611f0c57607f821691505b602082108103611f2c57634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0391909116815260200190565b600060208284031215611f7957600080fd5b5051919050565b634e487b7160e01b600052602160045260246000fd5b601f82111561072f576000816000526020600020601f850160051c81016020861015611fbf5750805b601f850160051c820191505b81811015611fde57828155600101611fcb565b505050505050565b81516001600160401b03811115611fff57611fff611c7c565b6120138161200d8454611ef8565b84611f96565b602080601f83116001811461204857600084156120305750858301515b600019600386901b1c1916600185901b178555611fde565b600085815260208120601f198616915b8281101561207757888601518255948401946001909101908401612058565b50858210156120955787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600082516120b7818460208701611b52565b9190910192915050565b808201808211156105b157634e487b7160e01b600052601160045260246000fdfe189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6a264697066735822122033908e3a5f8c1d2075fc9508858a878e54c959eb625c0361cbbc90f19a2f85b964736f6c63430008180033",
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
