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

// XENCryptoMintInfo is an auto generated low-level Go binding around an user-defined struct.
type XENCryptoMintInfo struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}

// XENCryptoStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type XENCryptoStakeInfo struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}

// XenMetaData contains all meta data concerning the Xen contract.
var XenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"MintClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"name\":\"RankClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAYS_IN_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_PM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_PM_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_RANK_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_RANK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PENALTY_PCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TERM_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TERM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TERM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TERM_AMPLIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TERM_AMPLIFIER_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_WINDOW_DAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_DAYS_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_MIN_BURN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeMinters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimMintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"other\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimRank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentAMP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEAAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentMaxTerm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rankDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaa\",\"type\":\"uint256\"}],\"name\":\"getGrossReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserMint\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"internalType\":\"structXENCrypto.MintInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserStake\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"internalType\":\"structXENCrypto.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalRank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalXenStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBurns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// XenABI is the input ABI used to generate the binding from.
// Deprecated: Use XenMetaData.ABI instead.
var XenABI = XenMetaData.ABI

// Xen is an auto generated Go binding around an Ethereum contract.
type Xen struct {
	XenCaller     // Read-only binding to the contract
	XenTransactor // Write-only binding to the contract
	XenFilterer   // Log filterer for contract events
}

// XenCaller is an auto generated read-only Go binding around an Ethereum contract.
type XenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XenSession struct {
	Contract     *Xen              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XenCallerSession struct {
	Contract *XenCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// XenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XenTransactorSession struct {
	Contract     *XenTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XenRaw is an auto generated low-level Go binding around an Ethereum contract.
type XenRaw struct {
	Contract *Xen // Generic contract binding to access the raw methods on
}

// XenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XenCallerRaw struct {
	Contract *XenCaller // Generic read-only contract binding to access the raw methods on
}

// XenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XenTransactorRaw struct {
	Contract *XenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXen creates a new instance of Xen, bound to a specific deployed contract.
func NewXen(address common.Address, backend bind.ContractBackend) (*Xen, error) {
	contract, err := bindXen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Xen{XenCaller: XenCaller{contract: contract}, XenTransactor: XenTransactor{contract: contract}, XenFilterer: XenFilterer{contract: contract}}, nil
}

// NewXenCaller creates a new read-only instance of Xen, bound to a specific deployed contract.
func NewXenCaller(address common.Address, caller bind.ContractCaller) (*XenCaller, error) {
	contract, err := bindXen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XenCaller{contract: contract}, nil
}

// NewXenTransactor creates a new write-only instance of Xen, bound to a specific deployed contract.
func NewXenTransactor(address common.Address, transactor bind.ContractTransactor) (*XenTransactor, error) {
	contract, err := bindXen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XenTransactor{contract: contract}, nil
}

// NewXenFilterer creates a new log filterer instance of Xen, bound to a specific deployed contract.
func NewXenFilterer(address common.Address, filterer bind.ContractFilterer) (*XenFilterer, error) {
	contract, err := bindXen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XenFilterer{contract: contract}, nil
}

// bindXen binds a generic wrapper to an already deployed contract.
func bindXen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xen *XenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xen.Contract.XenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xen *XenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xen.Contract.XenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xen *XenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xen.Contract.XenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xen *XenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xen *XenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xen *XenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xen.Contract.contract.Transact(opts, method, params...)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Xen *XenCaller) AUTHORS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "AUTHORS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Xen *XenSession) AUTHORS() (string, error) {
	return _Xen.Contract.AUTHORS(&_Xen.CallOpts)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Xen *XenCallerSession) AUTHORS() (string, error) {
	return _Xen.Contract.AUTHORS(&_Xen.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_Xen *XenCaller) DAYSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "DAYS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_Xen *XenSession) DAYSINYEAR() (*big.Int, error) {
	return _Xen.Contract.DAYSINYEAR(&_Xen.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_Xen *XenCallerSession) DAYSINYEAR() (*big.Int, error) {
	return _Xen.Contract.DAYSINYEAR(&_Xen.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_Xen *XenCaller) EAAPMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "EAA_PM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_Xen *XenSession) EAAPMSTART() (*big.Int, error) {
	return _Xen.Contract.EAAPMSTART(&_Xen.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_Xen *XenCallerSession) EAAPMSTART() (*big.Int, error) {
	return _Xen.Contract.EAAPMSTART(&_Xen.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_Xen *XenCaller) EAAPMSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "EAA_PM_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_Xen *XenSession) EAAPMSTEP() (*big.Int, error) {
	return _Xen.Contract.EAAPMSTEP(&_Xen.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_Xen *XenCallerSession) EAAPMSTEP() (*big.Int, error) {
	return _Xen.Contract.EAAPMSTEP(&_Xen.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_Xen *XenCaller) EAARANKSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "EAA_RANK_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_Xen *XenSession) EAARANKSTEP() (*big.Int, error) {
	return _Xen.Contract.EAARANKSTEP(&_Xen.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_Xen *XenCallerSession) EAARANKSTEP() (*big.Int, error) {
	return _Xen.Contract.EAARANKSTEP(&_Xen.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_Xen *XenCaller) GENESISRANK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "GENESIS_RANK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_Xen *XenSession) GENESISRANK() (*big.Int, error) {
	return _Xen.Contract.GENESISRANK(&_Xen.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_Xen *XenCallerSession) GENESISRANK() (*big.Int, error) {
	return _Xen.Contract.GENESISRANK(&_Xen.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_Xen *XenCaller) MAXPENALTYPCT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "MAX_PENALTY_PCT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_Xen *XenSession) MAXPENALTYPCT() (*big.Int, error) {
	return _Xen.Contract.MAXPENALTYPCT(&_Xen.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_Xen *XenCallerSession) MAXPENALTYPCT() (*big.Int, error) {
	return _Xen.Contract.MAXPENALTYPCT(&_Xen.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_Xen *XenCaller) MAXTERMEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "MAX_TERM_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_Xen *XenSession) MAXTERMEND() (*big.Int, error) {
	return _Xen.Contract.MAXTERMEND(&_Xen.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_Xen *XenCallerSession) MAXTERMEND() (*big.Int, error) {
	return _Xen.Contract.MAXTERMEND(&_Xen.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_Xen *XenCaller) MAXTERMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "MAX_TERM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_Xen *XenSession) MAXTERMSTART() (*big.Int, error) {
	return _Xen.Contract.MAXTERMSTART(&_Xen.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_Xen *XenCallerSession) MAXTERMSTART() (*big.Int, error) {
	return _Xen.Contract.MAXTERMSTART(&_Xen.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_Xen *XenCaller) MINTERM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "MIN_TERM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_Xen *XenSession) MINTERM() (*big.Int, error) {
	return _Xen.Contract.MINTERM(&_Xen.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_Xen *XenCallerSession) MINTERM() (*big.Int, error) {
	return _Xen.Contract.MINTERM(&_Xen.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_Xen *XenCaller) REWARDAMPLIFIEREND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "REWARD_AMPLIFIER_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_Xen *XenSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _Xen.Contract.REWARDAMPLIFIEREND(&_Xen.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_Xen *XenCallerSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _Xen.Contract.REWARDAMPLIFIEREND(&_Xen.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_Xen *XenCaller) REWARDAMPLIFIERSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "REWARD_AMPLIFIER_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_Xen *XenSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _Xen.Contract.REWARDAMPLIFIERSTART(&_Xen.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_Xen *XenCallerSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _Xen.Contract.REWARDAMPLIFIERSTART(&_Xen.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Xen *XenCaller) SECONDSINDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "SECONDS_IN_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Xen *XenSession) SECONDSINDAY() (*big.Int, error) {
	return _Xen.Contract.SECONDSINDAY(&_Xen.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Xen *XenCallerSession) SECONDSINDAY() (*big.Int, error) {
	return _Xen.Contract.SECONDSINDAY(&_Xen.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_Xen *XenCaller) TERMAMPLIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "TERM_AMPLIFIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_Xen *XenSession) TERMAMPLIFIER() (*big.Int, error) {
	return _Xen.Contract.TERMAMPLIFIER(&_Xen.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_Xen *XenCallerSession) TERMAMPLIFIER() (*big.Int, error) {
	return _Xen.Contract.TERMAMPLIFIER(&_Xen.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_Xen *XenCaller) TERMAMPLIFIERTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "TERM_AMPLIFIER_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_Xen *XenSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _Xen.Contract.TERMAMPLIFIERTHRESHOLD(&_Xen.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_Xen *XenCallerSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _Xen.Contract.TERMAMPLIFIERTHRESHOLD(&_Xen.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_Xen *XenCaller) WITHDRAWALWINDOWDAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "WITHDRAWAL_WINDOW_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_Xen *XenSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _Xen.Contract.WITHDRAWALWINDOWDAYS(&_Xen.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_Xen *XenCallerSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _Xen.Contract.WITHDRAWALWINDOWDAYS(&_Xen.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_Xen *XenCaller) XENAPYDAYSSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "XEN_APY_DAYS_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_Xen *XenSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _Xen.Contract.XENAPYDAYSSTEP(&_Xen.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_Xen *XenCallerSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _Xen.Contract.XENAPYDAYSSTEP(&_Xen.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_Xen *XenCaller) XENAPYEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "XEN_APY_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_Xen *XenSession) XENAPYEND() (*big.Int, error) {
	return _Xen.Contract.XENAPYEND(&_Xen.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_Xen *XenCallerSession) XENAPYEND() (*big.Int, error) {
	return _Xen.Contract.XENAPYEND(&_Xen.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_Xen *XenCaller) XENAPYSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "XEN_APY_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_Xen *XenSession) XENAPYSTART() (*big.Int, error) {
	return _Xen.Contract.XENAPYSTART(&_Xen.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_Xen *XenCallerSession) XENAPYSTART() (*big.Int, error) {
	return _Xen.Contract.XENAPYSTART(&_Xen.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_Xen *XenCaller) XENMINBURN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "XEN_MIN_BURN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_Xen *XenSession) XENMINBURN() (*big.Int, error) {
	return _Xen.Contract.XENMINBURN(&_Xen.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_Xen *XenCallerSession) XENMINBURN() (*big.Int, error) {
	return _Xen.Contract.XENMINBURN(&_Xen.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_Xen *XenCaller) XENMINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "XEN_MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_Xen *XenSession) XENMINSTAKE() (*big.Int, error) {
	return _Xen.Contract.XENMINSTAKE(&_Xen.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_Xen *XenCallerSession) XENMINSTAKE() (*big.Int, error) {
	return _Xen.Contract.XENMINSTAKE(&_Xen.CallOpts)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_Xen *XenCaller) ActiveMinters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "activeMinters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_Xen *XenSession) ActiveMinters() (*big.Int, error) {
	return _Xen.Contract.ActiveMinters(&_Xen.CallOpts)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_Xen *XenCallerSession) ActiveMinters() (*big.Int, error) {
	return _Xen.Contract.ActiveMinters(&_Xen.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_Xen *XenCaller) ActiveStakes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "activeStakes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_Xen *XenSession) ActiveStakes() (*big.Int, error) {
	return _Xen.Contract.ActiveStakes(&_Xen.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_Xen *XenCallerSession) ActiveStakes() (*big.Int, error) {
	return _Xen.Contract.ActiveStakes(&_Xen.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Xen *XenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Xen *XenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Xen.Contract.Allowance(&_Xen.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Xen *XenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Xen.Contract.Allowance(&_Xen.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Xen *XenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Xen *XenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Xen.Contract.BalanceOf(&_Xen.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Xen *XenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Xen.Contract.BalanceOf(&_Xen.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Xen *XenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Xen *XenSession) Decimals() (uint8, error) {
	return _Xen.Contract.Decimals(&_Xen.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Xen *XenCallerSession) Decimals() (uint8, error) {
	return _Xen.Contract.Decimals(&_Xen.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Xen *XenCaller) GenesisTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "genesisTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Xen *XenSession) GenesisTs() (*big.Int, error) {
	return _Xen.Contract.GenesisTs(&_Xen.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Xen *XenCallerSession) GenesisTs() (*big.Int, error) {
	return _Xen.Contract.GenesisTs(&_Xen.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_Xen *XenCaller) GetCurrentAMP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "getCurrentAMP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_Xen *XenSession) GetCurrentAMP() (*big.Int, error) {
	return _Xen.Contract.GetCurrentAMP(&_Xen.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_Xen *XenCallerSession) GetCurrentAMP() (*big.Int, error) {
	return _Xen.Contract.GetCurrentAMP(&_Xen.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_Xen *XenCaller) GetCurrentAPY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "getCurrentAPY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_Xen *XenSession) GetCurrentAPY() (*big.Int, error) {
	return _Xen.Contract.GetCurrentAPY(&_Xen.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_Xen *XenCallerSession) GetCurrentAPY() (*big.Int, error) {
	return _Xen.Contract.GetCurrentAPY(&_Xen.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_Xen *XenCaller) GetCurrentEAAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "getCurrentEAAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_Xen *XenSession) GetCurrentEAAR() (*big.Int, error) {
	return _Xen.Contract.GetCurrentEAAR(&_Xen.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_Xen *XenCallerSession) GetCurrentEAAR() (*big.Int, error) {
	return _Xen.Contract.GetCurrentEAAR(&_Xen.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_Xen *XenCaller) GetCurrentMaxTerm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "getCurrentMaxTerm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_Xen *XenSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _Xen.Contract.GetCurrentMaxTerm(&_Xen.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_Xen *XenCallerSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _Xen.Contract.GetCurrentMaxTerm(&_Xen.CallOpts)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_Xen *XenCaller) GetGrossReward(opts *bind.CallOpts, rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "getGrossReward", rankDelta, amplifier, term, eaa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_Xen *XenSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _Xen.Contract.GetGrossReward(&_Xen.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_Xen *XenCallerSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _Xen.Contract.GetGrossReward(&_Xen.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Xen *XenCaller) GetUserMint(opts *bind.CallOpts) (XENCryptoMintInfo, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "getUserMint")

	if err != nil {
		return *new(XENCryptoMintInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(XENCryptoMintInfo)).(*XENCryptoMintInfo)

	return out0, err

}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Xen *XenSession) GetUserMint() (XENCryptoMintInfo, error) {
	return _Xen.Contract.GetUserMint(&_Xen.CallOpts)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Xen *XenCallerSession) GetUserMint() (XENCryptoMintInfo, error) {
	return _Xen.Contract.GetUserMint(&_Xen.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_Xen *XenCaller) GetUserStake(opts *bind.CallOpts) (XENCryptoStakeInfo, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "getUserStake")

	if err != nil {
		return *new(XENCryptoStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(XENCryptoStakeInfo)).(*XENCryptoStakeInfo)

	return out0, err

}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_Xen *XenSession) GetUserStake() (XENCryptoStakeInfo, error) {
	return _Xen.Contract.GetUserStake(&_Xen.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_Xen *XenCallerSession) GetUserStake() (XENCryptoStakeInfo, error) {
	return _Xen.Contract.GetUserStake(&_Xen.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_Xen *XenCaller) GlobalRank(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "globalRank")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_Xen *XenSession) GlobalRank() (*big.Int, error) {
	return _Xen.Contract.GlobalRank(&_Xen.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_Xen *XenCallerSession) GlobalRank() (*big.Int, error) {
	return _Xen.Contract.GlobalRank(&_Xen.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xen *XenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xen *XenSession) Name() (string, error) {
	return _Xen.Contract.Name(&_Xen.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xen *XenCallerSession) Name() (string, error) {
	return _Xen.Contract.Name(&_Xen.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xen *XenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xen *XenSession) Symbol() (string, error) {
	return _Xen.Contract.Symbol(&_Xen.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xen *XenCallerSession) Symbol() (string, error) {
	return _Xen.Contract.Symbol(&_Xen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Xen *XenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Xen *XenSession) TotalSupply() (*big.Int, error) {
	return _Xen.Contract.TotalSupply(&_Xen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Xen *XenCallerSession) TotalSupply() (*big.Int, error) {
	return _Xen.Contract.TotalSupply(&_Xen.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_Xen *XenCaller) TotalXenStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "totalXenStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_Xen *XenSession) TotalXenStaked() (*big.Int, error) {
	return _Xen.Contract.TotalXenStaked(&_Xen.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_Xen *XenCallerSession) TotalXenStaked() (*big.Int, error) {
	return _Xen.Contract.TotalXenStaked(&_Xen.CallOpts)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_Xen *XenCaller) UserBurns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "userBurns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_Xen *XenSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _Xen.Contract.UserBurns(&_Xen.CallOpts, arg0)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_Xen *XenCallerSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _Xen.Contract.UserBurns(&_Xen.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_Xen *XenCaller) UserMints(opts *bind.CallOpts, arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "userMints", arg0)

	outstruct := new(struct {
		User       common.Address
		Term       *big.Int
		MaturityTs *big.Int
		Rank       *big.Int
		Amplifier  *big.Int
		EaaRate    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Term = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaturityTs = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rank = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amplifier = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EaaRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_Xen *XenSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _Xen.Contract.UserMints(&_Xen.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_Xen *XenCallerSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _Xen.Contract.UserMints(&_Xen.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_Xen *XenCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	var out []interface{}
	err := _Xen.contract.Call(opts, &out, "userStakes", arg0)

	outstruct := new(struct {
		Term       *big.Int
		MaturityTs *big.Int
		Amount     *big.Int
		Apy        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Term = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaturityTs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Apy = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_Xen *XenSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _Xen.Contract.UserStakes(&_Xen.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_Xen *XenCallerSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _Xen.Contract.UserStakes(&_Xen.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Xen *XenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Xen *XenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Approve(&_Xen.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Xen *XenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Approve(&_Xen.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_Xen *XenTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_Xen *XenSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Burn(&_Xen.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_Xen *XenTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Burn(&_Xen.TransactOpts, user, amount)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_Xen *XenTransactor) ClaimMintReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "claimMintReward")
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_Xen *XenSession) ClaimMintReward() (*types.Transaction, error) {
	return _Xen.Contract.ClaimMintReward(&_Xen.TransactOpts)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_Xen *XenTransactorSession) ClaimMintReward() (*types.Transaction, error) {
	return _Xen.Contract.ClaimMintReward(&_Xen.TransactOpts)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_Xen *XenTransactor) ClaimMintRewardAndShare(opts *bind.TransactOpts, other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "claimMintRewardAndShare", other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_Xen *XenSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.ClaimMintRewardAndShare(&_Xen.TransactOpts, other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_Xen *XenTransactorSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.ClaimMintRewardAndShare(&_Xen.TransactOpts, other, pct)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_Xen *XenTransactor) ClaimMintRewardAndStake(opts *bind.TransactOpts, pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "claimMintRewardAndStake", pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_Xen *XenSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.ClaimMintRewardAndStake(&_Xen.TransactOpts, pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_Xen *XenTransactorSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.ClaimMintRewardAndStake(&_Xen.TransactOpts, pct, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_Xen *XenTransactor) ClaimRank(opts *bind.TransactOpts, term *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "claimRank", term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_Xen *XenSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.ClaimRank(&_Xen.TransactOpts, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_Xen *XenTransactorSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.ClaimRank(&_Xen.TransactOpts, term)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Xen *XenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Xen *XenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.DecreaseAllowance(&_Xen.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Xen *XenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.DecreaseAllowance(&_Xen.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Xen *XenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Xen *XenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.IncreaseAllowance(&_Xen.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Xen *XenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.IncreaseAllowance(&_Xen.TransactOpts, spender, addedValue)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_Xen *XenTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "stake", amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_Xen *XenSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Stake(&_Xen.TransactOpts, amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_Xen *XenTransactorSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Stake(&_Xen.TransactOpts, amount, term)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Xen *XenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Xen *XenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Transfer(&_Xen.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Xen *XenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.Transfer(&_Xen.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Xen *XenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Xen *XenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.TransferFrom(&_Xen.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Xen *XenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Xen.Contract.TransferFrom(&_Xen.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Xen *XenTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xen.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Xen *XenSession) Withdraw() (*types.Transaction, error) {
	return _Xen.Contract.Withdraw(&_Xen.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Xen *XenTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Xen.Contract.Withdraw(&_Xen.TransactOpts)
}

// XenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Xen contract.
type XenApprovalIterator struct {
	Event *XenApproval // Event containing the contract specifics and raw log

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
func (it *XenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XenApproval)
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
		it.Event = new(XenApproval)
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
func (it *XenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XenApproval represents a Approval event raised by the Xen contract.
type XenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Xen *XenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*XenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Xen.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &XenApprovalIterator{contract: _Xen.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Xen *XenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *XenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Xen.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XenApproval)
				if err := _Xen.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Xen *XenFilterer) ParseApproval(log types.Log) (*XenApproval, error) {
	event := new(XenApproval)
	if err := _Xen.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XenMintClaimedIterator is returned from FilterMintClaimed and is used to iterate over the raw logs and unpacked data for MintClaimed events raised by the Xen contract.
type XenMintClaimedIterator struct {
	Event *XenMintClaimed // Event containing the contract specifics and raw log

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
func (it *XenMintClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XenMintClaimed)
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
		it.Event = new(XenMintClaimed)
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
func (it *XenMintClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XenMintClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XenMintClaimed represents a MintClaimed event raised by the Xen contract.
type XenMintClaimed struct {
	User         common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintClaimed is a free log retrieval operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_Xen *XenFilterer) FilterMintClaimed(opts *bind.FilterOpts, user []common.Address) (*XenMintClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.FilterLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &XenMintClaimedIterator{contract: _Xen.contract, event: "MintClaimed", logs: logs, sub: sub}, nil
}

// WatchMintClaimed is a free log subscription operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_Xen *XenFilterer) WatchMintClaimed(opts *bind.WatchOpts, sink chan<- *XenMintClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.WatchLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XenMintClaimed)
				if err := _Xen.contract.UnpackLog(event, "MintClaimed", log); err != nil {
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

// ParseMintClaimed is a log parse operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_Xen *XenFilterer) ParseMintClaimed(log types.Log) (*XenMintClaimed, error) {
	event := new(XenMintClaimed)
	if err := _Xen.contract.UnpackLog(event, "MintClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XenRankClaimedIterator is returned from FilterRankClaimed and is used to iterate over the raw logs and unpacked data for RankClaimed events raised by the Xen contract.
type XenRankClaimedIterator struct {
	Event *XenRankClaimed // Event containing the contract specifics and raw log

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
func (it *XenRankClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XenRankClaimed)
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
		it.Event = new(XenRankClaimed)
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
func (it *XenRankClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XenRankClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XenRankClaimed represents a RankClaimed event raised by the Xen contract.
type XenRankClaimed struct {
	User common.Address
	Term *big.Int
	Rank *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRankClaimed is a free log retrieval operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_Xen *XenFilterer) FilterRankClaimed(opts *bind.FilterOpts, user []common.Address) (*XenRankClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.FilterLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &XenRankClaimedIterator{contract: _Xen.contract, event: "RankClaimed", logs: logs, sub: sub}, nil
}

// WatchRankClaimed is a free log subscription operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_Xen *XenFilterer) WatchRankClaimed(opts *bind.WatchOpts, sink chan<- *XenRankClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.WatchLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XenRankClaimed)
				if err := _Xen.contract.UnpackLog(event, "RankClaimed", log); err != nil {
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

// ParseRankClaimed is a log parse operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_Xen *XenFilterer) ParseRankClaimed(log types.Log) (*XenRankClaimed, error) {
	event := new(XenRankClaimed)
	if err := _Xen.contract.UnpackLog(event, "RankClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XenStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Xen contract.
type XenStakedIterator struct {
	Event *XenStaked // Event containing the contract specifics and raw log

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
func (it *XenStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XenStaked)
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
		it.Event = new(XenStaked)
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
func (it *XenStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XenStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XenStaked represents a Staked event raised by the Xen contract.
type XenStaked struct {
	User   common.Address
	Amount *big.Int
	Term   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_Xen *XenFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*XenStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &XenStakedIterator{contract: _Xen.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_Xen *XenFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *XenStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XenStaked)
				if err := _Xen.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_Xen *XenFilterer) ParseStaked(log types.Log) (*XenStaked, error) {
	event := new(XenStaked)
	if err := _Xen.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Xen contract.
type XenTransferIterator struct {
	Event *XenTransfer // Event containing the contract specifics and raw log

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
func (it *XenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XenTransfer)
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
		it.Event = new(XenTransfer)
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
func (it *XenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XenTransfer represents a Transfer event raised by the Xen contract.
type XenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Xen *XenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*XenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Xen.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &XenTransferIterator{contract: _Xen.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Xen *XenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *XenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Xen.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XenTransfer)
				if err := _Xen.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Xen *XenFilterer) ParseTransfer(log types.Log) (*XenTransfer, error) {
	event := new(XenTransfer)
	if err := _Xen.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XenWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Xen contract.
type XenWithdrawnIterator struct {
	Event *XenWithdrawn // Event containing the contract specifics and raw log

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
func (it *XenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XenWithdrawn)
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
		it.Event = new(XenWithdrawn)
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
func (it *XenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XenWithdrawn represents a Withdrawn event raised by the Xen contract.
type XenWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_Xen *XenFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*XenWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &XenWithdrawnIterator{contract: _Xen.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_Xen *XenFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *XenWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xen.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XenWithdrawn)
				if err := _Xen.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_Xen *XenFilterer) ParseWithdrawn(log types.Log) (*XenWithdrawn, error) {
	event := new(XenWithdrawn)
	if err := _Xen.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
