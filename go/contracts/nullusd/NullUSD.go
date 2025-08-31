// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nullusd

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

// NullusdMetaData contains all meta data concerning the Nullusd contract.
var NullusdMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405180604001604052806007815260200166139d5b1b1554d160ca1b81525060405180604001604052806004815260200163139554d160e21b8152508160039081620000609190620002d0565b5060046200006f8282620002d0565b505050620000a93362000087620000af60201b60201c565b6200009490600a620004b1565b620000a390620f4240620004c9565b620000b4565b620004f9565b601290565b6001600160a01b038216620000e45760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b620000f260008383620000f6565b5050565b6001600160a01b03831662000125578060026000828254620001199190620004e3565b90915550620001999050565b6001600160a01b038316600090815260208190526040902054818110156200017a5760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620000db565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b038216620001b757600280548290039055620001d6565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200021c91815260200190565b60405180910390a3505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200025457607f821691505b6020821081036200027557634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002cb576000816000526020600020601f850160051c81016020861015620002a65750805b601f850160051c820191505b81811015620002c757828155600101620002b2565b5050505b505050565b81516001600160401b03811115620002ec57620002ec62000229565b6200030481620002fd84546200023f565b846200027b565b602080601f8311600181146200033c5760008415620003235750858301515b600019600386901b1c1916600185901b178555620002c7565b600085815260208120601f198616915b828110156200036d578886015182559484019460019091019084016200034c565b50858210156200038c5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b600181815b80851115620003f3578160001904821115620003d757620003d76200039c565b80851615620003e557918102915b93841c9390800290620003b7565b509250929050565b6000826200040c57506001620004ab565b816200041b57506000620004ab565b81600181146200043457600281146200043f576200045f565b6001915050620004ab565b60ff8411156200045357620004536200039c565b50506001821b620004ab565b5060208310610133831016604e8410600b841016171562000484575081810a620004ab565b620004908383620003b2565b8060001904821115620004a757620004a76200039c565b0290505b92915050565b6000620004c260ff841683620003fb565b9392505050565b8082028115828204841417620004ab57620004ab6200039c565b80820180821115620004ab57620004ab6200039c565b61071180620005096000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806306fdde0314610088578063095ea7b3146100a657806318160ddd146100c957806323b872dd146100db578063313ce567146100ee57806370a08231146100fd57806395d89b4114610126578063a9059cbb1461012e578063dd62ed3e14610141575b600080fd5b610090610154565b60405161009d9190610525565b60405180910390f35b6100b96100b4366004610590565b6101e6565b604051901515815260200161009d565b6002545b60405190815260200161009d565b6100b96100e93660046105ba565b610200565b6040516012815260200161009d565b6100cd61010b3660046105f6565b6001600160a01b031660009081526020819052604090205490565b610090610224565b6100b961013c366004610590565b610233565b6100cd61014f366004610618565b610241565b6060600380546101639061064b565b80601f016020809104026020016040519081016040528092919081815260200182805461018f9061064b565b80156101dc5780601f106101b1576101008083540402835291602001916101dc565b820191906000526020600020905b8154815290600101906020018083116101bf57829003601f168201915b5050505050905090565b6000336101f481858561026c565b60019150505b92915050565b60003361020e85828561027e565b6102198585856102da565b506001949350505050565b6060600480546101639061064b565b6000336101f48185856102da565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6102798383836001610339565b505050565b600061028a8484610241565b905060001981146102d457818110156102c557828183604051637dc7a0d960e11b81526004016102bc93929190610685565b60405180910390fd5b6102d484848484036000610339565b50505050565b6001600160a01b038316610304576000604051634b637e8f60e11b81526004016102bc91906106a6565b6001600160a01b03821661032e57600060405163ec442f0560e01b81526004016102bc91906106a6565b61027983838361040e565b6001600160a01b03841661036357600060405163e602df0560e01b81526004016102bc91906106a6565b6001600160a01b03831661038d576000604051634a1406b160e11b81526004016102bc91906106a6565b6001600160a01b03808516600090815260016020908152604080832093871683529290522082905580156102d457826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161040091815260200190565b60405180910390a350505050565b6001600160a01b03831661043957806002600082825461042e91906106ba565b909155506104989050565b6001600160a01b038316600090815260208190526040902054818110156104795783818360405163391434e360e21b81526004016102bc93929190610685565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166104b4576002805482900390556104d3565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161051891815260200190565b60405180910390a3505050565b60006020808352835180602085015260005b8181101561055357858101830151858201604001528201610537565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461058b57600080fd5b919050565b600080604083850312156105a357600080fd5b6105ac83610574565b946020939093013593505050565b6000806000606084860312156105cf57600080fd5b6105d884610574565b92506105e660208501610574565b9150604084013590509250925092565b60006020828403121561060857600080fd5b61061182610574565b9392505050565b6000806040838503121561062b57600080fd5b61063483610574565b915061064260208401610574565b90509250929050565b600181811c9082168061065f57607f821691505b60208210810361067f57634e487b7160e01b600052602260045260246000fd5b50919050565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0391909116815260200190565b808201808211156101fa57634e487b7160e01b600052601160045260246000fdfea2646970667358221220cc4ccc4ee3f389333f59a16adec6364f02a4499f2847f9bf7ab33883520c640c64736f6c63430008180033",
}

// NullusdABI is the input ABI used to generate the binding from.
// Deprecated: Use NullusdMetaData.ABI instead.
var NullusdABI = NullusdMetaData.ABI

// NullusdBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NullusdMetaData.Bin instead.
var NullusdBin = NullusdMetaData.Bin

// DeployNullusd deploys a new Ethereum contract, binding an instance of Nullusd to it.
func DeployNullusd(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nullusd, error) {
	parsed, err := NullusdMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NullusdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nullusd{NullusdCaller: NullusdCaller{contract: contract}, NullusdTransactor: NullusdTransactor{contract: contract}, NullusdFilterer: NullusdFilterer{contract: contract}}, nil
}

// Nullusd is an auto generated Go binding around an Ethereum contract.
type Nullusd struct {
	NullusdCaller     // Read-only binding to the contract
	NullusdTransactor // Write-only binding to the contract
	NullusdFilterer   // Log filterer for contract events
}

// NullusdCaller is an auto generated read-only Go binding around an Ethereum contract.
type NullusdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NullusdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NullusdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NullusdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NullusdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NullusdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NullusdSession struct {
	Contract     *Nullusd          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NullusdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NullusdCallerSession struct {
	Contract *NullusdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NullusdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NullusdTransactorSession struct {
	Contract     *NullusdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NullusdRaw is an auto generated low-level Go binding around an Ethereum contract.
type NullusdRaw struct {
	Contract *Nullusd // Generic contract binding to access the raw methods on
}

// NullusdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NullusdCallerRaw struct {
	Contract *NullusdCaller // Generic read-only contract binding to access the raw methods on
}

// NullusdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NullusdTransactorRaw struct {
	Contract *NullusdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNullusd creates a new instance of Nullusd, bound to a specific deployed contract.
func NewNullusd(address common.Address, backend bind.ContractBackend) (*Nullusd, error) {
	contract, err := bindNullusd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nullusd{NullusdCaller: NullusdCaller{contract: contract}, NullusdTransactor: NullusdTransactor{contract: contract}, NullusdFilterer: NullusdFilterer{contract: contract}}, nil
}

// NewNullusdCaller creates a new read-only instance of Nullusd, bound to a specific deployed contract.
func NewNullusdCaller(address common.Address, caller bind.ContractCaller) (*NullusdCaller, error) {
	contract, err := bindNullusd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NullusdCaller{contract: contract}, nil
}

// NewNullusdTransactor creates a new write-only instance of Nullusd, bound to a specific deployed contract.
func NewNullusdTransactor(address common.Address, transactor bind.ContractTransactor) (*NullusdTransactor, error) {
	contract, err := bindNullusd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NullusdTransactor{contract: contract}, nil
}

// NewNullusdFilterer creates a new log filterer instance of Nullusd, bound to a specific deployed contract.
func NewNullusdFilterer(address common.Address, filterer bind.ContractFilterer) (*NullusdFilterer, error) {
	contract, err := bindNullusd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NullusdFilterer{contract: contract}, nil
}

// bindNullusd binds a generic wrapper to an already deployed contract.
func bindNullusd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NullusdMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nullusd *NullusdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nullusd.Contract.NullusdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nullusd *NullusdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nullusd.Contract.NullusdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nullusd *NullusdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nullusd.Contract.NullusdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nullusd *NullusdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nullusd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nullusd *NullusdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nullusd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nullusd *NullusdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nullusd.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nullusd *NullusdCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nullusd *NullusdSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nullusd.Contract.Allowance(&_Nullusd.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nullusd *NullusdCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nullusd.Contract.Allowance(&_Nullusd.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nullusd *NullusdCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nullusd *NullusdSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Nullusd.Contract.BalanceOf(&_Nullusd.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nullusd *NullusdCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Nullusd.Contract.BalanceOf(&_Nullusd.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nullusd *NullusdCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nullusd *NullusdSession) Decimals() (uint8, error) {
	return _Nullusd.Contract.Decimals(&_Nullusd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nullusd *NullusdCallerSession) Decimals() (uint8, error) {
	return _Nullusd.Contract.Decimals(&_Nullusd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nullusd *NullusdCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nullusd *NullusdSession) Name() (string, error) {
	return _Nullusd.Contract.Name(&_Nullusd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nullusd *NullusdCallerSession) Name() (string, error) {
	return _Nullusd.Contract.Name(&_Nullusd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nullusd *NullusdCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nullusd *NullusdSession) Symbol() (string, error) {
	return _Nullusd.Contract.Symbol(&_Nullusd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nullusd *NullusdCallerSession) Symbol() (string, error) {
	return _Nullusd.Contract.Symbol(&_Nullusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nullusd *NullusdCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nullusd *NullusdSession) TotalSupply() (*big.Int, error) {
	return _Nullusd.Contract.TotalSupply(&_Nullusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nullusd *NullusdCallerSession) TotalSupply() (*big.Int, error) {
	return _Nullusd.Contract.TotalSupply(&_Nullusd.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Nullusd *NullusdSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Approve(&_Nullusd.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Approve(&_Nullusd.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Nullusd *NullusdSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Transfer(&_Nullusd.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Transfer(&_Nullusd.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Nullusd *NullusdSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.TransferFrom(&_Nullusd.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.TransferFrom(&_Nullusd.TransactOpts, from, to, value)
}

// NullusdApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nullusd contract.
type NullusdApprovalIterator struct {
	Event *NullusdApproval // Event containing the contract specifics and raw log

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
func (it *NullusdApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NullusdApproval)
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
		it.Event = new(NullusdApproval)
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
func (it *NullusdApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NullusdApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NullusdApproval represents a Approval event raised by the Nullusd contract.
type NullusdApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nullusd *NullusdFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NullusdApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nullusd.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NullusdApprovalIterator{contract: _Nullusd.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nullusd *NullusdFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NullusdApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nullusd.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NullusdApproval)
				if err := _Nullusd.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Nullusd *NullusdFilterer) ParseApproval(log types.Log) (*NullusdApproval, error) {
	event := new(NullusdApproval)
	if err := _Nullusd.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NullusdTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nullusd contract.
type NullusdTransferIterator struct {
	Event *NullusdTransfer // Event containing the contract specifics and raw log

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
func (it *NullusdTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NullusdTransfer)
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
		it.Event = new(NullusdTransfer)
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
func (it *NullusdTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NullusdTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NullusdTransfer represents a Transfer event raised by the Nullusd contract.
type NullusdTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nullusd *NullusdFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NullusdTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nullusd.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NullusdTransferIterator{contract: _Nullusd.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nullusd *NullusdFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NullusdTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nullusd.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NullusdTransfer)
				if err := _Nullusd.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Nullusd *NullusdFilterer) ParseTransfer(log types.Log) (*NullusdTransfer, error) {
	event := new(NullusdTransfer)
	if err := _Nullusd.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
