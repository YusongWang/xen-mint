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

// LikeBNBMintInfo is an auto generated low-level Go binding around an user-defined struct.
type LikeBNBMintInfo struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}

// LikeBNBStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type LikeBNBStakeInfo struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}

// LikeBnbMetaData contains all meta data concerning the LikeBnb contract.
var LikeBnbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"MintClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"name\":\"RankClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAYS_IN_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_PM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_PM_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_RANK_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_RANK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PENALTY_PCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TERM_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TERM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TERM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TERM_AMPLIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TERM_AMPLIFIER_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_WINDOW_DAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_DAYS_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_MIN_BURN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cRank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eeaRate\",\"type\":\"uint256\"}],\"name\":\"_calculateMintReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeMinters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"calculateMintReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimMintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"other\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimRank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentAMP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEAAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentMaxTerm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rankDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaa\",\"type\":\"uint256\"}],\"name\":\"getGrossReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserMint\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"internalType\":\"structLikeBNB.MintInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserStake\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"internalType\":\"structLikeBNB.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalRank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalXenStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBurns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LikeBnbABI is the input ABI used to generate the binding from.
// Deprecated: Use LikeBnbMetaData.ABI instead.
var LikeBnbABI = LikeBnbMetaData.ABI

// LikeBnb is an auto generated Go binding around an Ethereum contract.
type LikeBnb struct {
	LikeBnbCaller     // Read-only binding to the contract
	LikeBnbTransactor // Write-only binding to the contract
	LikeBnbFilterer   // Log filterer for contract events
}

// LikeBnbCaller is an auto generated read-only Go binding around an Ethereum contract.
type LikeBnbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeBnbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LikeBnbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeBnbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LikeBnbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LikeBnbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LikeBnbSession struct {
	Contract     *LikeBnb          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LikeBnbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LikeBnbCallerSession struct {
	Contract *LikeBnbCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LikeBnbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LikeBnbTransactorSession struct {
	Contract     *LikeBnbTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LikeBnbRaw is an auto generated low-level Go binding around an Ethereum contract.
type LikeBnbRaw struct {
	Contract *LikeBnb // Generic contract binding to access the raw methods on
}

// LikeBnbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LikeBnbCallerRaw struct {
	Contract *LikeBnbCaller // Generic read-only contract binding to access the raw methods on
}

// LikeBnbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LikeBnbTransactorRaw struct {
	Contract *LikeBnbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLikeBnb creates a new instance of LikeBnb, bound to a specific deployed contract.
func NewLikeBnb(address common.Address, backend bind.ContractBackend) (*LikeBnb, error) {
	contract, err := bindLikeBnb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LikeBnb{LikeBnbCaller: LikeBnbCaller{contract: contract}, LikeBnbTransactor: LikeBnbTransactor{contract: contract}, LikeBnbFilterer: LikeBnbFilterer{contract: contract}}, nil
}

// NewLikeBnbCaller creates a new read-only instance of LikeBnb, bound to a specific deployed contract.
func NewLikeBnbCaller(address common.Address, caller bind.ContractCaller) (*LikeBnbCaller, error) {
	contract, err := bindLikeBnb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LikeBnbCaller{contract: contract}, nil
}

// NewLikeBnbTransactor creates a new write-only instance of LikeBnb, bound to a specific deployed contract.
func NewLikeBnbTransactor(address common.Address, transactor bind.ContractTransactor) (*LikeBnbTransactor, error) {
	contract, err := bindLikeBnb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LikeBnbTransactor{contract: contract}, nil
}

// NewLikeBnbFilterer creates a new log filterer instance of LikeBnb, bound to a specific deployed contract.
func NewLikeBnbFilterer(address common.Address, filterer bind.ContractFilterer) (*LikeBnbFilterer, error) {
	contract, err := bindLikeBnb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LikeBnbFilterer{contract: contract}, nil
}

// bindLikeBnb binds a generic wrapper to an already deployed contract.
func bindLikeBnb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LikeBnbMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeBnb *LikeBnbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeBnb.Contract.LikeBnbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeBnb *LikeBnbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeBnb.Contract.LikeBnbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeBnb *LikeBnbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeBnb.Contract.LikeBnbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LikeBnb *LikeBnbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LikeBnb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LikeBnb *LikeBnbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeBnb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LikeBnb *LikeBnbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LikeBnb.Contract.contract.Transact(opts, method, params...)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_LikeBnb *LikeBnbCaller) AUTHORS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "AUTHORS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_LikeBnb *LikeBnbSession) AUTHORS() (string, error) {
	return _LikeBnb.Contract.AUTHORS(&_LikeBnb.CallOpts)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_LikeBnb *LikeBnbCallerSession) AUTHORS() (string, error) {
	return _LikeBnb.Contract.AUTHORS(&_LikeBnb.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) DAYSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "DAYS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_LikeBnb *LikeBnbSession) DAYSINYEAR() (*big.Int, error) {
	return _LikeBnb.Contract.DAYSINYEAR(&_LikeBnb.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) DAYSINYEAR() (*big.Int, error) {
	return _LikeBnb.Contract.DAYSINYEAR(&_LikeBnb.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) EAAPMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "EAA_PM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_LikeBnb *LikeBnbSession) EAAPMSTART() (*big.Int, error) {
	return _LikeBnb.Contract.EAAPMSTART(&_LikeBnb.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) EAAPMSTART() (*big.Int, error) {
	return _LikeBnb.Contract.EAAPMSTART(&_LikeBnb.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) EAAPMSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "EAA_PM_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbSession) EAAPMSTEP() (*big.Int, error) {
	return _LikeBnb.Contract.EAAPMSTEP(&_LikeBnb.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) EAAPMSTEP() (*big.Int, error) {
	return _LikeBnb.Contract.EAAPMSTEP(&_LikeBnb.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) EAARANKSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "EAA_RANK_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbSession) EAARANKSTEP() (*big.Int, error) {
	return _LikeBnb.Contract.EAARANKSTEP(&_LikeBnb.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) EAARANKSTEP() (*big.Int, error) {
	return _LikeBnb.Contract.EAARANKSTEP(&_LikeBnb.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) GENESISRANK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "GENESIS_RANK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_LikeBnb *LikeBnbSession) GENESISRANK() (*big.Int, error) {
	return _LikeBnb.Contract.GENESISRANK(&_LikeBnb.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GENESISRANK() (*big.Int, error) {
	return _LikeBnb.Contract.GENESISRANK(&_LikeBnb.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) MAXPENALTYPCT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "MAX_PENALTY_PCT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_LikeBnb *LikeBnbSession) MAXPENALTYPCT() (*big.Int, error) {
	return _LikeBnb.Contract.MAXPENALTYPCT(&_LikeBnb.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) MAXPENALTYPCT() (*big.Int, error) {
	return _LikeBnb.Contract.MAXPENALTYPCT(&_LikeBnb.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) MAXTERMEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "MAX_TERM_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_LikeBnb *LikeBnbSession) MAXTERMEND() (*big.Int, error) {
	return _LikeBnb.Contract.MAXTERMEND(&_LikeBnb.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) MAXTERMEND() (*big.Int, error) {
	return _LikeBnb.Contract.MAXTERMEND(&_LikeBnb.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) MAXTERMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "MAX_TERM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_LikeBnb *LikeBnbSession) MAXTERMSTART() (*big.Int, error) {
	return _LikeBnb.Contract.MAXTERMSTART(&_LikeBnb.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) MAXTERMSTART() (*big.Int, error) {
	return _LikeBnb.Contract.MAXTERMSTART(&_LikeBnb.CallOpts)
}

// MINSTART is a free data retrieval call binding the contract method 0xca0246d6.
//
// Solidity: function MIN_START() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) MINSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "MIN_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTART is a free data retrieval call binding the contract method 0xca0246d6.
//
// Solidity: function MIN_START() view returns(uint256)
func (_LikeBnb *LikeBnbSession) MINSTART() (*big.Int, error) {
	return _LikeBnb.Contract.MINSTART(&_LikeBnb.CallOpts)
}

// MINSTART is a free data retrieval call binding the contract method 0xca0246d6.
//
// Solidity: function MIN_START() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) MINSTART() (*big.Int, error) {
	return _LikeBnb.Contract.MINSTART(&_LikeBnb.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) MINTERM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "MIN_TERM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_LikeBnb *LikeBnbSession) MINTERM() (*big.Int, error) {
	return _LikeBnb.Contract.MINTERM(&_LikeBnb.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) MINTERM() (*big.Int, error) {
	return _LikeBnb.Contract.MINTERM(&_LikeBnb.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) REWARDAMPLIFIEREND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "REWARD_AMPLIFIER_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_LikeBnb *LikeBnbSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _LikeBnb.Contract.REWARDAMPLIFIEREND(&_LikeBnb.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _LikeBnb.Contract.REWARDAMPLIFIEREND(&_LikeBnb.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) REWARDAMPLIFIERSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "REWARD_AMPLIFIER_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_LikeBnb *LikeBnbSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _LikeBnb.Contract.REWARDAMPLIFIERSTART(&_LikeBnb.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _LikeBnb.Contract.REWARDAMPLIFIERSTART(&_LikeBnb.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) SECONDSINDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "SECONDS_IN_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_LikeBnb *LikeBnbSession) SECONDSINDAY() (*big.Int, error) {
	return _LikeBnb.Contract.SECONDSINDAY(&_LikeBnb.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) SECONDSINDAY() (*big.Int, error) {
	return _LikeBnb.Contract.SECONDSINDAY(&_LikeBnb.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) TERMAMPLIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "TERM_AMPLIFIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_LikeBnb *LikeBnbSession) TERMAMPLIFIER() (*big.Int, error) {
	return _LikeBnb.Contract.TERMAMPLIFIER(&_LikeBnb.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) TERMAMPLIFIER() (*big.Int, error) {
	return _LikeBnb.Contract.TERMAMPLIFIER(&_LikeBnb.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) TERMAMPLIFIERTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "TERM_AMPLIFIER_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_LikeBnb *LikeBnbSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _LikeBnb.Contract.TERMAMPLIFIERTHRESHOLD(&_LikeBnb.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _LikeBnb.Contract.TERMAMPLIFIERTHRESHOLD(&_LikeBnb.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) WITHDRAWALWINDOWDAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "WITHDRAWAL_WINDOW_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_LikeBnb *LikeBnbSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _LikeBnb.Contract.WITHDRAWALWINDOWDAYS(&_LikeBnb.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _LikeBnb.Contract.WITHDRAWALWINDOWDAYS(&_LikeBnb.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) XENAPYDAYSSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "XEN_APY_DAYS_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _LikeBnb.Contract.XENAPYDAYSSTEP(&_LikeBnb.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _LikeBnb.Contract.XENAPYDAYSSTEP(&_LikeBnb.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) XENAPYEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "XEN_APY_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_LikeBnb *LikeBnbSession) XENAPYEND() (*big.Int, error) {
	return _LikeBnb.Contract.XENAPYEND(&_LikeBnb.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) XENAPYEND() (*big.Int, error) {
	return _LikeBnb.Contract.XENAPYEND(&_LikeBnb.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) XENAPYSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "XEN_APY_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_LikeBnb *LikeBnbSession) XENAPYSTART() (*big.Int, error) {
	return _LikeBnb.Contract.XENAPYSTART(&_LikeBnb.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) XENAPYSTART() (*big.Int, error) {
	return _LikeBnb.Contract.XENAPYSTART(&_LikeBnb.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) XENMINBURN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "XEN_MIN_BURN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_LikeBnb *LikeBnbSession) XENMINBURN() (*big.Int, error) {
	return _LikeBnb.Contract.XENMINBURN(&_LikeBnb.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) XENMINBURN() (*big.Int, error) {
	return _LikeBnb.Contract.XENMINBURN(&_LikeBnb.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) XENMINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "XEN_MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_LikeBnb *LikeBnbSession) XENMINSTAKE() (*big.Int, error) {
	return _LikeBnb.Contract.XENMINSTAKE(&_LikeBnb.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) XENMINSTAKE() (*big.Int, error) {
	return _LikeBnb.Contract.XENMINSTAKE(&_LikeBnb.CallOpts)
}

// CalculateMintReward is a free data retrieval call binding the contract method 0x74aa43d8.
//
// Solidity: function _calculateMintReward(uint256 cRank, uint256 term, uint256 maturityTs, uint256 amplifier, uint256 eeaRate) view returns(uint256)
func (_LikeBnb *LikeBnbCaller) CalculateMintReward(opts *bind.CallOpts, cRank *big.Int, term *big.Int, maturityTs *big.Int, amplifier *big.Int, eeaRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "_calculateMintReward", cRank, term, maturityTs, amplifier, eeaRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMintReward is a free data retrieval call binding the contract method 0x74aa43d8.
//
// Solidity: function _calculateMintReward(uint256 cRank, uint256 term, uint256 maturityTs, uint256 amplifier, uint256 eeaRate) view returns(uint256)
func (_LikeBnb *LikeBnbSession) CalculateMintReward(cRank *big.Int, term *big.Int, maturityTs *big.Int, amplifier *big.Int, eeaRate *big.Int) (*big.Int, error) {
	return _LikeBnb.Contract.CalculateMintReward(&_LikeBnb.CallOpts, cRank, term, maturityTs, amplifier, eeaRate)
}

// CalculateMintReward is a free data retrieval call binding the contract method 0x74aa43d8.
//
// Solidity: function _calculateMintReward(uint256 cRank, uint256 term, uint256 maturityTs, uint256 amplifier, uint256 eeaRate) view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) CalculateMintReward(cRank *big.Int, term *big.Int, maturityTs *big.Int, amplifier *big.Int, eeaRate *big.Int) (*big.Int, error) {
	return _LikeBnb.Contract.CalculateMintReward(&_LikeBnb.CallOpts, cRank, term, maturityTs, amplifier, eeaRate)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) ActiveMinters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "activeMinters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_LikeBnb *LikeBnbSession) ActiveMinters() (*big.Int, error) {
	return _LikeBnb.Contract.ActiveMinters(&_LikeBnb.CallOpts)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) ActiveMinters() (*big.Int, error) {
	return _LikeBnb.Contract.ActiveMinters(&_LikeBnb.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) ActiveStakes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "activeStakes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_LikeBnb *LikeBnbSession) ActiveStakes() (*big.Int, error) {
	return _LikeBnb.Contract.ActiveStakes(&_LikeBnb.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) ActiveStakes() (*big.Int, error) {
	return _LikeBnb.Contract.ActiveStakes(&_LikeBnb.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LikeBnb *LikeBnbCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LikeBnb *LikeBnbSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.Allowance(&_LikeBnb.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.Allowance(&_LikeBnb.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LikeBnb *LikeBnbCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LikeBnb *LikeBnbSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.BalanceOf(&_LikeBnb.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.BalanceOf(&_LikeBnb.CallOpts, account)
}

// InnerCalculateMintReward is a free data retrieval call binding the contract method 0x92d77e80.
//
// Solidity: function calculateMintReward(address sender) view returns(uint256 rewardAmount)
func (_LikeBnb *LikeBnbCaller) InnerCalculateMintReward(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "calculateMintReward", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InnerCalculateMintReward is a free data retrieval call binding the contract method 0x92d77e80.
//
// Solidity: function calculateMintReward(address sender) view returns(uint256 rewardAmount)
func (_LikeBnb *LikeBnbSession) InnerCalculateMintReward(sender common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.InnerCalculateMintReward(&_LikeBnb.CallOpts, sender)
}

// InnerCalculateMintReward is a free data retrieval call binding the contract method 0x92d77e80.
//
// Solidity: function calculateMintReward(address sender) view returns(uint256 rewardAmount)
func (_LikeBnb *LikeBnbCallerSession) InnerCalculateMintReward(sender common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.InnerCalculateMintReward(&_LikeBnb.CallOpts, sender)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LikeBnb *LikeBnbCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LikeBnb *LikeBnbSession) Decimals() (uint8, error) {
	return _LikeBnb.Contract.Decimals(&_LikeBnb.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LikeBnb *LikeBnbCallerSession) Decimals() (uint8, error) {
	return _LikeBnb.Contract.Decimals(&_LikeBnb.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) GenesisTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "genesisTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_LikeBnb *LikeBnbSession) GenesisTs() (*big.Int, error) {
	return _LikeBnb.Contract.GenesisTs(&_LikeBnb.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GenesisTs() (*big.Int, error) {
	return _LikeBnb.Contract.GenesisTs(&_LikeBnb.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) GetCurrentAMP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "getCurrentAMP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_LikeBnb *LikeBnbSession) GetCurrentAMP() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentAMP(&_LikeBnb.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GetCurrentAMP() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentAMP(&_LikeBnb.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) GetCurrentAPY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "getCurrentAPY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_LikeBnb *LikeBnbSession) GetCurrentAPY() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentAPY(&_LikeBnb.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GetCurrentAPY() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentAPY(&_LikeBnb.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) GetCurrentEAAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "getCurrentEAAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_LikeBnb *LikeBnbSession) GetCurrentEAAR() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentEAAR(&_LikeBnb.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GetCurrentEAAR() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentEAAR(&_LikeBnb.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) GetCurrentMaxTerm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "getCurrentMaxTerm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_LikeBnb *LikeBnbSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentMaxTerm(&_LikeBnb.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _LikeBnb.Contract.GetCurrentMaxTerm(&_LikeBnb.CallOpts)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_LikeBnb *LikeBnbCaller) GetGrossReward(opts *bind.CallOpts, rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "getGrossReward", rankDelta, amplifier, term, eaa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_LikeBnb *LikeBnbSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _LikeBnb.Contract.GetGrossReward(&_LikeBnb.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _LikeBnb.Contract.GetGrossReward(&_LikeBnb.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_LikeBnb *LikeBnbCaller) GetUserMint(opts *bind.CallOpts) (LikeBNBMintInfo, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "getUserMint")

	if err != nil {
		return *new(LikeBNBMintInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LikeBNBMintInfo)).(*LikeBNBMintInfo)

	return out0, err

}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_LikeBnb *LikeBnbSession) GetUserMint() (LikeBNBMintInfo, error) {
	return _LikeBnb.Contract.GetUserMint(&_LikeBnb.CallOpts)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_LikeBnb *LikeBnbCallerSession) GetUserMint() (LikeBNBMintInfo, error) {
	return _LikeBnb.Contract.GetUserMint(&_LikeBnb.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_LikeBnb *LikeBnbCaller) GetUserStake(opts *bind.CallOpts) (LikeBNBStakeInfo, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "getUserStake")

	if err != nil {
		return *new(LikeBNBStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LikeBNBStakeInfo)).(*LikeBNBStakeInfo)

	return out0, err

}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_LikeBnb *LikeBnbSession) GetUserStake() (LikeBNBStakeInfo, error) {
	return _LikeBnb.Contract.GetUserStake(&_LikeBnb.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_LikeBnb *LikeBnbCallerSession) GetUserStake() (LikeBNBStakeInfo, error) {
	return _LikeBnb.Contract.GetUserStake(&_LikeBnb.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) GlobalRank(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "globalRank")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_LikeBnb *LikeBnbSession) GlobalRank() (*big.Int, error) {
	return _LikeBnb.Contract.GlobalRank(&_LikeBnb.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) GlobalRank() (*big.Int, error) {
	return _LikeBnb.Contract.GlobalRank(&_LikeBnb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikeBnb *LikeBnbCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikeBnb *LikeBnbSession) Name() (string, error) {
	return _LikeBnb.Contract.Name(&_LikeBnb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LikeBnb *LikeBnbCallerSession) Name() (string, error) {
	return _LikeBnb.Contract.Name(&_LikeBnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikeBnb *LikeBnbCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikeBnb *LikeBnbSession) Symbol() (string, error) {
	return _LikeBnb.Contract.Symbol(&_LikeBnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LikeBnb *LikeBnbCallerSession) Symbol() (string, error) {
	return _LikeBnb.Contract.Symbol(&_LikeBnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikeBnb *LikeBnbSession) TotalSupply() (*big.Int, error) {
	return _LikeBnb.Contract.TotalSupply(&_LikeBnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) TotalSupply() (*big.Int, error) {
	return _LikeBnb.Contract.TotalSupply(&_LikeBnb.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_LikeBnb *LikeBnbCaller) TotalXenStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "totalXenStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_LikeBnb *LikeBnbSession) TotalXenStaked() (*big.Int, error) {
	return _LikeBnb.Contract.TotalXenStaked(&_LikeBnb.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) TotalXenStaked() (*big.Int, error) {
	return _LikeBnb.Contract.TotalXenStaked(&_LikeBnb.CallOpts)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_LikeBnb *LikeBnbCaller) UserBurns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "userBurns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_LikeBnb *LikeBnbSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.UserBurns(&_LikeBnb.CallOpts, arg0)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_LikeBnb *LikeBnbCallerSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _LikeBnb.Contract.UserBurns(&_LikeBnb.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_LikeBnb *LikeBnbCaller) UserMints(opts *bind.CallOpts, arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "userMints", arg0)

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
func (_LikeBnb *LikeBnbSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _LikeBnb.Contract.UserMints(&_LikeBnb.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_LikeBnb *LikeBnbCallerSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _LikeBnb.Contract.UserMints(&_LikeBnb.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_LikeBnb *LikeBnbCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	var out []interface{}
	err := _LikeBnb.contract.Call(opts, &out, "userStakes", arg0)

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
func (_LikeBnb *LikeBnbSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _LikeBnb.Contract.UserStakes(&_LikeBnb.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_LikeBnb *LikeBnbCallerSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _LikeBnb.Contract.UserStakes(&_LikeBnb.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Approve(&_LikeBnb.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Approve(&_LikeBnb.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_LikeBnb *LikeBnbTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_LikeBnb *LikeBnbSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Burn(&_LikeBnb.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_LikeBnb *LikeBnbTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Burn(&_LikeBnb.TransactOpts, user, amount)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_LikeBnb *LikeBnbTransactor) ClaimMintReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "claimMintReward")
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_LikeBnb *LikeBnbSession) ClaimMintReward() (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimMintReward(&_LikeBnb.TransactOpts)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_LikeBnb *LikeBnbTransactorSession) ClaimMintReward() (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimMintReward(&_LikeBnb.TransactOpts)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_LikeBnb *LikeBnbTransactor) ClaimMintRewardAndShare(opts *bind.TransactOpts, other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "claimMintRewardAndShare", other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_LikeBnb *LikeBnbSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimMintRewardAndShare(&_LikeBnb.TransactOpts, other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_LikeBnb *LikeBnbTransactorSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimMintRewardAndShare(&_LikeBnb.TransactOpts, other, pct)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_LikeBnb *LikeBnbTransactor) ClaimMintRewardAndStake(opts *bind.TransactOpts, pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "claimMintRewardAndStake", pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_LikeBnb *LikeBnbSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimMintRewardAndStake(&_LikeBnb.TransactOpts, pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_LikeBnb *LikeBnbTransactorSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimMintRewardAndStake(&_LikeBnb.TransactOpts, pct, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_LikeBnb *LikeBnbTransactor) ClaimRank(opts *bind.TransactOpts, term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "claimRank", term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_LikeBnb *LikeBnbSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimRank(&_LikeBnb.TransactOpts, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_LikeBnb *LikeBnbTransactorSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.ClaimRank(&_LikeBnb.TransactOpts, term)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LikeBnb *LikeBnbTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LikeBnb *LikeBnbSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.DecreaseAllowance(&_LikeBnb.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LikeBnb *LikeBnbTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.DecreaseAllowance(&_LikeBnb.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LikeBnb *LikeBnbTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LikeBnb *LikeBnbSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.IncreaseAllowance(&_LikeBnb.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LikeBnb *LikeBnbTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.IncreaseAllowance(&_LikeBnb.TransactOpts, spender, addedValue)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_LikeBnb *LikeBnbTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "stake", amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_LikeBnb *LikeBnbSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Stake(&_LikeBnb.TransactOpts, amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_LikeBnb *LikeBnbTransactorSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Stake(&_LikeBnb.TransactOpts, amount, term)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Transfer(&_LikeBnb.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.Transfer(&_LikeBnb.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.TransferFrom(&_LikeBnb.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LikeBnb *LikeBnbTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LikeBnb.Contract.TransferFrom(&_LikeBnb.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LikeBnb *LikeBnbTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LikeBnb.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LikeBnb *LikeBnbSession) Withdraw() (*types.Transaction, error) {
	return _LikeBnb.Contract.Withdraw(&_LikeBnb.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_LikeBnb *LikeBnbTransactorSession) Withdraw() (*types.Transaction, error) {
	return _LikeBnb.Contract.Withdraw(&_LikeBnb.TransactOpts)
}

// LikeBnbApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LikeBnb contract.
type LikeBnbApprovalIterator struct {
	Event *LikeBnbApproval // Event containing the contract specifics and raw log

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
func (it *LikeBnbApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeBnbApproval)
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
		it.Event = new(LikeBnbApproval)
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
func (it *LikeBnbApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeBnbApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeBnbApproval represents a Approval event raised by the LikeBnb contract.
type LikeBnbApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LikeBnb *LikeBnbFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LikeBnbApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LikeBnb.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LikeBnbApprovalIterator{contract: _LikeBnb.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LikeBnb *LikeBnbFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LikeBnbApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LikeBnb.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeBnbApproval)
				if err := _LikeBnb.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LikeBnb *LikeBnbFilterer) ParseApproval(log types.Log) (*LikeBnbApproval, error) {
	event := new(LikeBnbApproval)
	if err := _LikeBnb.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeBnbMintClaimedIterator is returned from FilterMintClaimed and is used to iterate over the raw logs and unpacked data for MintClaimed events raised by the LikeBnb contract.
type LikeBnbMintClaimedIterator struct {
	Event *LikeBnbMintClaimed // Event containing the contract specifics and raw log

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
func (it *LikeBnbMintClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeBnbMintClaimed)
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
		it.Event = new(LikeBnbMintClaimed)
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
func (it *LikeBnbMintClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeBnbMintClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeBnbMintClaimed represents a MintClaimed event raised by the LikeBnb contract.
type LikeBnbMintClaimed struct {
	User         common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintClaimed is a free log retrieval operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_LikeBnb *LikeBnbFilterer) FilterMintClaimed(opts *bind.FilterOpts, user []common.Address) (*LikeBnbMintClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.FilterLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &LikeBnbMintClaimedIterator{contract: _LikeBnb.contract, event: "MintClaimed", logs: logs, sub: sub}, nil
}

// WatchMintClaimed is a free log subscription operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_LikeBnb *LikeBnbFilterer) WatchMintClaimed(opts *bind.WatchOpts, sink chan<- *LikeBnbMintClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.WatchLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeBnbMintClaimed)
				if err := _LikeBnb.contract.UnpackLog(event, "MintClaimed", log); err != nil {
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
func (_LikeBnb *LikeBnbFilterer) ParseMintClaimed(log types.Log) (*LikeBnbMintClaimed, error) {
	event := new(LikeBnbMintClaimed)
	if err := _LikeBnb.contract.UnpackLog(event, "MintClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeBnbRankClaimedIterator is returned from FilterRankClaimed and is used to iterate over the raw logs and unpacked data for RankClaimed events raised by the LikeBnb contract.
type LikeBnbRankClaimedIterator struct {
	Event *LikeBnbRankClaimed // Event containing the contract specifics and raw log

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
func (it *LikeBnbRankClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeBnbRankClaimed)
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
		it.Event = new(LikeBnbRankClaimed)
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
func (it *LikeBnbRankClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeBnbRankClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeBnbRankClaimed represents a RankClaimed event raised by the LikeBnb contract.
type LikeBnbRankClaimed struct {
	User common.Address
	Term *big.Int
	Rank *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRankClaimed is a free log retrieval operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_LikeBnb *LikeBnbFilterer) FilterRankClaimed(opts *bind.FilterOpts, user []common.Address) (*LikeBnbRankClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.FilterLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &LikeBnbRankClaimedIterator{contract: _LikeBnb.contract, event: "RankClaimed", logs: logs, sub: sub}, nil
}

// WatchRankClaimed is a free log subscription operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_LikeBnb *LikeBnbFilterer) WatchRankClaimed(opts *bind.WatchOpts, sink chan<- *LikeBnbRankClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.WatchLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeBnbRankClaimed)
				if err := _LikeBnb.contract.UnpackLog(event, "RankClaimed", log); err != nil {
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
func (_LikeBnb *LikeBnbFilterer) ParseRankClaimed(log types.Log) (*LikeBnbRankClaimed, error) {
	event := new(LikeBnbRankClaimed)
	if err := _LikeBnb.contract.UnpackLog(event, "RankClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeBnbStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the LikeBnb contract.
type LikeBnbStakedIterator struct {
	Event *LikeBnbStaked // Event containing the contract specifics and raw log

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
func (it *LikeBnbStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeBnbStaked)
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
		it.Event = new(LikeBnbStaked)
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
func (it *LikeBnbStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeBnbStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeBnbStaked represents a Staked event raised by the LikeBnb contract.
type LikeBnbStaked struct {
	User   common.Address
	Amount *big.Int
	Term   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_LikeBnb *LikeBnbFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*LikeBnbStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &LikeBnbStakedIterator{contract: _LikeBnb.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_LikeBnb *LikeBnbFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LikeBnbStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeBnbStaked)
				if err := _LikeBnb.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_LikeBnb *LikeBnbFilterer) ParseStaked(log types.Log) (*LikeBnbStaked, error) {
	event := new(LikeBnbStaked)
	if err := _LikeBnb.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeBnbTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LikeBnb contract.
type LikeBnbTransferIterator struct {
	Event *LikeBnbTransfer // Event containing the contract specifics and raw log

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
func (it *LikeBnbTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeBnbTransfer)
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
		it.Event = new(LikeBnbTransfer)
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
func (it *LikeBnbTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeBnbTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeBnbTransfer represents a Transfer event raised by the LikeBnb contract.
type LikeBnbTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LikeBnb *LikeBnbFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LikeBnbTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LikeBnb.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LikeBnbTransferIterator{contract: _LikeBnb.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LikeBnb *LikeBnbFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LikeBnbTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LikeBnb.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeBnbTransfer)
				if err := _LikeBnb.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LikeBnb *LikeBnbFilterer) ParseTransfer(log types.Log) (*LikeBnbTransfer, error) {
	event := new(LikeBnbTransfer)
	if err := _LikeBnb.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LikeBnbWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the LikeBnb contract.
type LikeBnbWithdrawnIterator struct {
	Event *LikeBnbWithdrawn // Event containing the contract specifics and raw log

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
func (it *LikeBnbWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LikeBnbWithdrawn)
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
		it.Event = new(LikeBnbWithdrawn)
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
func (it *LikeBnbWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LikeBnbWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LikeBnbWithdrawn represents a Withdrawn event raised by the LikeBnb contract.
type LikeBnbWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_LikeBnb *LikeBnbFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*LikeBnbWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &LikeBnbWithdrawnIterator{contract: _LikeBnb.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_LikeBnb *LikeBnbFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *LikeBnbWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LikeBnb.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LikeBnbWithdrawn)
				if err := _LikeBnb.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_LikeBnb *LikeBnbFilterer) ParseWithdrawn(log types.Log) (*LikeBnbWithdrawn, error) {
	event := new(LikeBnbWithdrawn)
	if err := _LikeBnb.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
