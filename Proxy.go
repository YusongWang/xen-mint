// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"proxyFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516109da3803806109da83398101604081905261002f9161015d565b306080523360a05261004081610046565b506101f8565b604080517f3d602d80600a3d3981f3363d3d373d3d3d363d7300000000000000000000000060208201523060601b6001600160601b03191660348201526e5af43d82803e903d91602b57fd5bf360881b604882015281516037818303018152605782019092526100ba908290607701610176565b60408051601f1981840301815291905280516020909101206000908155600154815b84811015610149576000336100f184846101c7565b60405160609290921b6001600160601b03191660208301526034820152605401604051602081830303815290604052805190602001209050808551602087016000f59350508080610141906101df565b9150506100dc565b5061015484826101c7565b60015550505050565b60006020828403121561016f57600080fd5b5051919050565b6000825160005b81811015610197576020818601810151858301520161017d565b818111156101a6576000828501525b509190910192915050565b634e487b7160e01b600052601160045260246000fd5b600082198211156101da576101da6101b1565b500190565b6000600182016101f1576101f16101b1565b5060010190565b60805160a0516107b7610223600039600081816101f9015261024d0152600060c601526107b76000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063246a95121461005157806330f3f0db14610066578063883b030014610079578063dcc9f2821461008c575b600080fd5b61006461005f366004610575565b6100bb565b005b6100646100743660046105c3565b6101ee565b6100646100873660046105dc565b610242565b61009f61009a36600461063d565b610327565b6040516001600160a01b03909116815260200160405180910390f35b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101465760405162461bcd60e51b815260206004820152602560248201527f4f6e6c79206f726967696e616c2063616e2063616c6c20746869732066756e636044820152643a34b7b71760d91b60648201526084015b60405180910390fd5b6000826001600160a01b0316826040516101609190610697565b6000604051808303816000865af19150503d806000811461019d576040519150601f19603f3d011682016040523d82523d6000602084013e6101a2565b606091505b50509050806101e95760405162461bcd60e51b81526020600482015260136024820152722a3930b739b0b1ba34b7b7103330b4b632b21760691b604482015260640161013d565b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102365760405162461bcd60e51b815260040161013d906106b3565b61023f816103a8565b50565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461028a5760405162461bcd60e51b815260040161013d906106b3565b835b610296848661070e565b8110156103205760006102a93383610327565b6040516312354a8960e11b81529091506001600160a01b0382169063246a9512906102da9087908790600401610726565b600060405180830381600087803b1580156102f457600080fd5b505af1158015610308573d6000803e3d6000fd5b5050505050808061031890610768565b91505061028c565b5050505050565b604080516001600160601b0319606094851b8116602080840191909152603480840195909552835180840390950185526054830184528451948101949094206000546001600160f81b031960748501523090961b9091166075830152608982015260a9808201949094528151808203909401845260c9019052815191012090565b60408051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b60208201523060601b6001600160601b03191660348201526e5af43d82803e903d91602b57fd5bf360881b60488201528151603781830301815260578201909252610413908290607701610697565b60408051601f1981840301815291905280516020909101206000908155600154815b848110156104a25760003361044a848461070e565b60405160609290921b6001600160601b03191660208301526034820152605401604051602081830303815290604052805190602001209050808551602087016000f5935050808061049a90610768565b915050610435565b506104ad848261070e565b60015550505050565b80356001600160a01b03811681146104cd57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126104f957600080fd5b813567ffffffffffffffff80821115610514576105146104d2565b604051601f8301601f19908116603f0116810190828211818310171561053c5761053c6104d2565b8160405283815286602085880101111561055557600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561058857600080fd5b610591836104b6565b9150602083013567ffffffffffffffff8111156105ad57600080fd5b6105b9858286016104e8565b9150509250929050565b6000602082840312156105d557600080fd5b5035919050565b600080600080608085870312156105f257600080fd5b8435935060208501359250610609604086016104b6565b9150606085013567ffffffffffffffff81111561062557600080fd5b610631878288016104e8565b91505092959194509250565b6000806040838503121561065057600080fd5b610659836104b6565b946020939093013593505050565b60005b8381101561068257818101518382015260200161066a565b83811115610691576000848401525b50505050565b600082516106a9818460208701610667565b9190910192915050565b60208082526025908201527f4f6e6c79206465706c6f7965722063616e2063616c6c20746869732066756e636040820152643a34b7b71760d91b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60008219821115610721576107216106f8565b500190565b60018060a01b03831681526040602082015260008251806040840152610753816060850160208701610667565b601f01601f1916919091016060019392505050565b60006001820161077a5761077a6106f8565b506001019056fea264697066735822122036f54fd50f1832e7a2ade1e495eaff99127058da117e00169a2c0d4bb8ff173b64736f6c634300080f0033",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// ProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProxyMetaData.Bin instead.
var ProxyBin = ProxyMetaData.Bin

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _n *big.Int) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyBin), backend, _n)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// ProxyFor is a free data retrieval call binding the contract method 0xdcc9f282.
//
// Solidity: function proxyFor(address sender, uint256 i) view returns(address proxy)
func (_Proxy *ProxyCaller) ProxyFor(opts *bind.CallOpts, sender common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "proxyFor", sender, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyFor is a free data retrieval call binding the contract method 0xdcc9f282.
//
// Solidity: function proxyFor(address sender, uint256 i) view returns(address proxy)
func (_Proxy *ProxySession) ProxyFor(sender common.Address, i *big.Int) (common.Address, error) {
	return _Proxy.Contract.ProxyFor(&_Proxy.CallOpts, sender, i)
}

// ProxyFor is a free data retrieval call binding the contract method 0xdcc9f282.
//
// Solidity: function proxyFor(address sender, uint256 i) view returns(address proxy)
func (_Proxy *ProxyCallerSession) ProxyFor(sender common.Address, i *big.Int) (common.Address, error) {
	return _Proxy.Contract.ProxyFor(&_Proxy.CallOpts, sender, i)
}

// Callback is a paid mutator transaction binding the contract method 0x246a9512.
//
// Solidity: function callback(address target, bytes data) returns()
func (_Proxy *ProxyTransactor) Callback(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "callback", target, data)
}

// Callback is a paid mutator transaction binding the contract method 0x246a9512.
//
// Solidity: function callback(address target, bytes data) returns()
func (_Proxy *ProxySession) Callback(target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Callback(&_Proxy.TransactOpts, target, data)
}

// Callback is a paid mutator transaction binding the contract method 0x246a9512.
//
// Solidity: function callback(address target, bytes data) returns()
func (_Proxy *ProxyTransactorSession) Callback(target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Callback(&_Proxy.TransactOpts, target, data)
}

// Execute is a paid mutator transaction binding the contract method 0x883b0300.
//
// Solidity: function execute(uint256 _start, uint256 _count, address target, bytes data) returns()
func (_Proxy *ProxyTransactor) Execute(opts *bind.TransactOpts, _start *big.Int, _count *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "execute", _start, _count, target, data)
}

// Execute is a paid mutator transaction binding the contract method 0x883b0300.
//
// Solidity: function execute(uint256 _start, uint256 _count, address target, bytes data) returns()
func (_Proxy *ProxySession) Execute(_start *big.Int, _count *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Execute(&_Proxy.TransactOpts, _start, _count, target, data)
}

// Execute is a paid mutator transaction binding the contract method 0x883b0300.
//
// Solidity: function execute(uint256 _start, uint256 _count, address target, bytes data) returns()
func (_Proxy *ProxyTransactorSession) Execute(_start *big.Int, _count *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Execute(&_Proxy.TransactOpts, _start, _count, target, data)
}

// Increase is a paid mutator transaction binding the contract method 0x30f3f0db.
//
// Solidity: function increase(uint256 _n) returns()
func (_Proxy *ProxyTransactor) Increase(opts *bind.TransactOpts, _n *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "increase", _n)
}

// Increase is a paid mutator transaction binding the contract method 0x30f3f0db.
//
// Solidity: function increase(uint256 _n) returns()
func (_Proxy *ProxySession) Increase(_n *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Increase(&_Proxy.TransactOpts, _n)
}

// Increase is a paid mutator transaction binding the contract method 0x30f3f0db.
//
// Solidity: function increase(uint256 _n) returns()
func (_Proxy *ProxyTransactorSession) Increase(_n *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.Increase(&_Proxy.TransactOpts, _n)
}
