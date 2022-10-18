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
	Bin: "0x608060405234801561001057600080fd5b506004361061004c5760003560e01c8063246a95121461005157806330f3f0db1461006d578063883b030014610089578063dcc9f282146100a5575b600080fd5b61006b60048036038101906100669190610710565b6100d5565b005b610087600480360381019061008291906107a2565b610215565b005b6100a3600480360381019061009e91906107cf565b6102af565b005b6100bf60048036038101906100ba9190610852565b6103ec565b6040516100cc91906108a1565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610163576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015a9061093f565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405161018a91906109d0565b6000604051808303816000865af19150503d80600081146101c7576040519150601f19603f3d011682016040523d82523d6000602084013e6101cc565b606091505b5050905080610210576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020790610a33565b60405180910390fd5b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029a90610ac5565b60405180910390fd5b6102ac81610453565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033490610ac5565b60405180910390fd5b60008490505b838561034f9190610b14565b8110156103e557600061036233836103ec565b90508073ffffffffffffffffffffffffffffffffffffffff1663246a951285856040518363ffffffff1660e01b815260040161039f929190610b92565b600060405180830381600087803b1580156103b957600080fd5b505af11580156103cd573d6000803e3d6000fd5b505050505080806103dd90610bc2565b915050610343565b5050505050565b6000808383604051602001610402929190610c73565b604051602081830303815290604052805190602001209050308160005460405160200161043193929190610d21565b6040516020818303038152906040528051906020012060001c91505092915050565b6000733d602d80600a3d3981f3363d3d373d3d3d363d7360601b3060601b6e5af43d82803e903d91602b57fd5bf360881b60405160200161049693929190610e03565b6040516020818303038152906040529050806040516020016104b891906109d0565b60405160208183030381529060405280519060200120600081905550600080600154905060005b8481101561053f5760003383836104f69190610b14565b604051602001610507929190610c73565b604051602081830303815290604052805190602001209050808551602087016000f5935050808061053790610bc2565b9150506104df565b50838161054c9190610b14565b60018190555050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105978261056c565b9050919050565b6105a78161058c565b81146105b257600080fd5b50565b6000813590506105c48161059e565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61061d826105d4565b810181811067ffffffffffffffff8211171561063c5761063b6105e5565b5b80604052505050565b600061064f610558565b905061065b8282610614565b919050565b600067ffffffffffffffff82111561067b5761067a6105e5565b5b610684826105d4565b9050602081019050919050565b82818337600083830152505050565b60006106b36106ae84610660565b610645565b9050828152602081018484840111156106cf576106ce6105cf565b5b6106da848285610691565b509392505050565b600082601f8301126106f7576106f66105ca565b5b81356107078482602086016106a0565b91505092915050565b6000806040838503121561072757610726610562565b5b6000610735858286016105b5565b925050602083013567ffffffffffffffff81111561075657610755610567565b5b610762858286016106e2565b9150509250929050565b6000819050919050565b61077f8161076c565b811461078a57600080fd5b50565b60008135905061079c81610776565b92915050565b6000602082840312156107b8576107b7610562565b5b60006107c68482850161078d565b91505092915050565b600080600080608085870312156107e9576107e8610562565b5b60006107f78782880161078d565b94505060206108088782880161078d565b9350506040610819878288016105b5565b925050606085013567ffffffffffffffff81111561083a57610839610567565b5b610846878288016106e2565b91505092959194509250565b6000806040838503121561086957610868610562565b5b6000610877858286016105b5565b92505060206108888582860161078d565b9150509250929050565b61089b8161058c565b82525050565b60006020820190506108b66000830184610892565b92915050565b600082825260208201905092915050565b7f4f6e6c79206f726967696e616c2063616e2063616c6c20746869732066756e6360008201527f74696f6e2e000000000000000000000000000000000000000000000000000000602082015250565b60006109296025836108bc565b9150610934826108cd565b604082019050919050565b600060208201905081810360008301526109588161091c565b9050919050565b600081519050919050565b600081905092915050565b60005b83811015610993578082015181840152602081019050610978565b60008484015250505050565b60006109aa8261095f565b6109b4818561096a565b93506109c4818560208601610975565b80840191505092915050565b60006109dc828461099f565b915081905092915050565b7f5472616e73616374696f6e206661696c65642e00000000000000000000000000600082015250565b6000610a1d6013836108bc565b9150610a28826109e7565b602082019050919050565b60006020820190508181036000830152610a4c81610a10565b9050919050565b7f4f6e6c79206465706c6f7965722063616e2063616c6c20746869732066756e6360008201527f74696f6e2e000000000000000000000000000000000000000000000000000000602082015250565b6000610aaf6025836108bc565b9150610aba82610a53565b604082019050919050565b60006020820190508181036000830152610ade81610aa2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b1f8261076c565b9150610b2a8361076c565b9250828201905080821115610b4257610b41610ae5565b5b92915050565b600082825260208201905092915050565b6000610b648261095f565b610b6e8185610b48565b9350610b7e818560208601610975565b610b87816105d4565b840191505092915050565b6000604082019050610ba76000830185610892565b8181036020830152610bb98184610b59565b90509392505050565b6000610bcd8261076c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610bff57610bfe610ae5565b5b600182019050919050565b60008160601b9050919050565b6000610c2282610c0a565b9050919050565b6000610c3482610c17565b9050919050565b610c4c610c478261058c565b610c29565b82525050565b6000819050919050565b610c6d610c688261076c565b610c52565b82525050565b6000610c7f8285610c3b565b601482019150610c8f8284610c5c565b6020820191508190509392505050565b600081905092915050565b7fff00000000000000000000000000000000000000000000000000000000000000600082015250565b6000610ce0600183610c9f565b9150610ceb82610caa565b600182019050919050565b6000819050919050565b6000819050919050565b610d1b610d1682610cf6565b610d00565b82525050565b6000610d2c82610cd3565b9150610d388286610c3b565b601482019150610d488285610d0a565b602082019150610d588284610d0a565b602082019150819050949350505050565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b6000819050919050565b610db0610dab82610d69565b610d95565b82525050565b60007fffffffffffffffffffffffffffffff000000000000000000000000000000000082169050919050565b6000819050919050565b610dfd610df882610db6565b610de2565b82525050565b6000610e0f8286610d9f565b601482019150610e1f8285610d9f565b601482019150610e2f8284610dec565b600f8201915081905094935050505056fea2646970667358221220fb44d032cbbb347d7265f4f6fddd79fd0ec557ea87de33ec81a2ac34fdbf749f64736f6c63430008110033",
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
