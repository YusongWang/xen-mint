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

// BatchTrasferMetaData contains all meta data concerning the BatchTrasfer contract.
var BatchTrasferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BatchTrasferABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchTrasferMetaData.ABI instead.
var BatchTrasferABI = BatchTrasferMetaData.ABI

// BatchTrasfer is an auto generated Go binding around an Ethereum contract.
type BatchTrasfer struct {
	BatchTrasferCaller     // Read-only binding to the contract
	BatchTrasferTransactor // Write-only binding to the contract
	BatchTrasferFilterer   // Log filterer for contract events
}

// BatchTrasferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchTrasferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTrasferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchTrasferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTrasferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchTrasferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTrasferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchTrasferSession struct {
	Contract     *BatchTrasfer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchTrasferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchTrasferCallerSession struct {
	Contract *BatchTrasferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BatchTrasferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchTrasferTransactorSession struct {
	Contract     *BatchTrasferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BatchTrasferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchTrasferRaw struct {
	Contract *BatchTrasfer // Generic contract binding to access the raw methods on
}

// BatchTrasferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchTrasferCallerRaw struct {
	Contract *BatchTrasferCaller // Generic read-only contract binding to access the raw methods on
}

// BatchTrasferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchTrasferTransactorRaw struct {
	Contract *BatchTrasferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchTrasfer creates a new instance of BatchTrasfer, bound to a specific deployed contract.
func NewBatchTrasfer(address common.Address, backend bind.ContractBackend) (*BatchTrasfer, error) {
	contract, err := bindBatchTrasfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchTrasfer{BatchTrasferCaller: BatchTrasferCaller{contract: contract}, BatchTrasferTransactor: BatchTrasferTransactor{contract: contract}, BatchTrasferFilterer: BatchTrasferFilterer{contract: contract}}, nil
}

// NewBatchTrasferCaller creates a new read-only instance of BatchTrasfer, bound to a specific deployed contract.
func NewBatchTrasferCaller(address common.Address, caller bind.ContractCaller) (*BatchTrasferCaller, error) {
	contract, err := bindBatchTrasfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTrasferCaller{contract: contract}, nil
}

// NewBatchTrasferTransactor creates a new write-only instance of BatchTrasfer, bound to a specific deployed contract.
func NewBatchTrasferTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchTrasferTransactor, error) {
	contract, err := bindBatchTrasfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTrasferTransactor{contract: contract}, nil
}

// NewBatchTrasferFilterer creates a new log filterer instance of BatchTrasfer, bound to a specific deployed contract.
func NewBatchTrasferFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchTrasferFilterer, error) {
	contract, err := bindBatchTrasfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchTrasferFilterer{contract: contract}, nil
}

// bindBatchTrasfer binds a generic wrapper to an already deployed contract.
func bindBatchTrasfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchTrasferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTrasfer *BatchTrasferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchTrasfer.Contract.BatchTrasferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTrasfer *BatchTrasferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.BatchTrasferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTrasfer *BatchTrasferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.BatchTrasferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTrasfer *BatchTrasferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchTrasfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTrasfer *BatchTrasferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTrasfer *BatchTrasferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0x1939c1ff.
//
// Solidity: function transfer(address[] addresses, uint256 amount) payable returns()
func (_BatchTrasfer *BatchTrasferTransactor) Transfer(opts *bind.TransactOpts, addresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.contract.Transact(opts, "transfer", addresses, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x1939c1ff.
//
// Solidity: function transfer(address[] addresses, uint256 amount) payable returns()
func (_BatchTrasfer *BatchTrasferSession) Transfer(addresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.Transfer(&_BatchTrasfer.TransactOpts, addresses, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x1939c1ff.
//
// Solidity: function transfer(address[] addresses, uint256 amount) payable returns()
func (_BatchTrasfer *BatchTrasferTransactorSession) Transfer(addresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.Transfer(&_BatchTrasfer.TransactOpts, addresses, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xffc3a769.
//
// Solidity: function transfer(address[] addresses, uint256[] amounts) payable returns()
func (_BatchTrasfer *BatchTrasferTransactor) Transfer0(opts *bind.TransactOpts, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.contract.Transact(opts, "transfer0", addresses, amounts)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xffc3a769.
//
// Solidity: function transfer(address[] addresses, uint256[] amounts) payable returns()
func (_BatchTrasfer *BatchTrasferSession) Transfer0(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.Transfer0(&_BatchTrasfer.TransactOpts, addresses, amounts)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xffc3a769.
//
// Solidity: function transfer(address[] addresses, uint256[] amounts) payable returns()
func (_BatchTrasfer *BatchTrasferTransactorSession) Transfer0(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.Transfer0(&_BatchTrasfer.TransactOpts, addresses, amounts)
}

// TransferToken is a paid mutator transaction binding the contract method 0x7a728217.
//
// Solidity: function transferToken(address token, address[] addresses, uint256 amount) returns()
func (_BatchTrasfer *BatchTrasferTransactor) TransferToken(opts *bind.TransactOpts, token common.Address, addresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.contract.Transact(opts, "transferToken", token, addresses, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x7a728217.
//
// Solidity: function transferToken(address token, address[] addresses, uint256 amount) returns()
func (_BatchTrasfer *BatchTrasferSession) TransferToken(token common.Address, addresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.TransferToken(&_BatchTrasfer.TransactOpts, token, addresses, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x7a728217.
//
// Solidity: function transferToken(address token, address[] addresses, uint256 amount) returns()
func (_BatchTrasfer *BatchTrasferTransactorSession) TransferToken(token common.Address, addresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.TransferToken(&_BatchTrasfer.TransactOpts, token, addresses, amount)
}

// TransferToken0 is a paid mutator transaction binding the contract method 0xa62ceef5.
//
// Solidity: function transferToken(address token, address[] addresses, uint256[] amounts) returns()
func (_BatchTrasfer *BatchTrasferTransactor) TransferToken0(opts *bind.TransactOpts, token common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.contract.Transact(opts, "transferToken0", token, addresses, amounts)
}

// TransferToken0 is a paid mutator transaction binding the contract method 0xa62ceef5.
//
// Solidity: function transferToken(address token, address[] addresses, uint256[] amounts) returns()
func (_BatchTrasfer *BatchTrasferSession) TransferToken0(token common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.TransferToken0(&_BatchTrasfer.TransactOpts, token, addresses, amounts)
}

// TransferToken0 is a paid mutator transaction binding the contract method 0xa62ceef5.
//
// Solidity: function transferToken(address token, address[] addresses, uint256[] amounts) returns()
func (_BatchTrasfer *BatchTrasferTransactorSession) TransferToken0(token common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTrasfer.Contract.TransferToken0(&_BatchTrasfer.TransactOpts, token, addresses, amounts)
}
