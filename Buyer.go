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

// BuyerMetaData contains all meta data concerning the Buyer contract.
var BuyerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"day\",\"type\":\"uint8\"},{\"internalType\":\"contractXEN\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SECONDS_IN_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_addr\",\"outputs\":[{\"internalType\":\"contractXEN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_day\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"number\",\"type\":\"uint8\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"number\",\"type\":\"uint8\"}],\"name\":\"claimToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimTracker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"number\",\"type\":\"uint8\"}],\"name\":\"factory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fatoryTracker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAddr\",\"outputs\":[{\"internalType\":\"contractXEN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idTracker\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimday\",\"type\":\"uint256\"},{\"internalType\":\"contractMintForMint\",\"name\":\"mint\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractXEN\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060043610620001185760003560e01c80638da5cb5b11620000a5578063d1d80fdf116200006f578063d1d80fdf14620002b2578063eebb426d14620002d2578063f2fde38b14620002f2578063fdafc62c14620003125762000118565b80638da5cb5b146200022a5780639e382400146200024c578063b25f4202146200026e578063b2bdfa7b14620002905762000118565b806361a52a3611620000e757806361a52a3614620001a357806362ab1a7714620001c55780636350b2c214620001e7578063715018a6146200021e5762000118565b806308e53003146200011d5780631b7fe6b6146200013f57806320bc14d61462000161578063222abc111462000181575b600080fd5b6200012762000332565b60405162000136919062000c43565b60405180910390f35b620001496200035c565b60405162000158919062000c7b565b60405180910390f35b6200017f600480360381019062000179919062000cdb565b62000368565b005b6200018b620004e1565b6040516200019a919062000c7b565b60405180910390f35b620001ad620004ed565b604051620001bc919062000c7b565b60405180910390f35b620001cf620004f4565b604051620001de919062000c43565b60405180910390f35b620002056004803603810190620001ff919062000d3e565b6200051a565b6040516200021592919062000d95565b60405180910390f35b620002286200055e565b005b6200023462000576565b60405162000243919062000de7565b60405180910390f35b620002566200059f565b60405162000265919062000e15565b60405180910390f35b62000278620005b2565b60405162000287919062000c7b565b60405180910390f35b6200029a620005be565b604051620002a9919062000e57565b60405180910390f35b620002d06004803603810190620002ca919062000eb9565b620005e4565b005b620002f06004803603810190620002ea919062000cdb565b62000632565b005b6200031060048036038101906200030a919062000f1c565b6200076b565b005b6200033060048036038101906200032a919062000cdb565b620007f5565b005b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60028060000154905081565b6200037262000a05565b60005b8160ff16811015620004dd5760006200038f600162000a8a565b9050600030600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900460ff16604051620003d59062000b7a565b620003e39392919062000f73565b604051809103906000f08015801562000400573d6000803e3d6000fd5b5090506200040d62000b88565b600081600001818152505081816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600760008581526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050620004c4600162000a98565b5050508080620004d49062000fdf565b91505062000375565b5050565b60038060000154905081565b6201518081565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60076020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b6200056862000a05565b62000574600062000aae565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600560009054906101000a900460ff1681565b60018060000154905081565b600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b620005ee62000a05565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6200063c62000a05565b60005b8160ff168110156200076757600062000659600362000a8a565b9050600060076000838152602001908152602001600020604051806040016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050806020015173ffffffffffffffffffffffffffffffffffffffff16634e71d92d6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156200072a57600080fd5b505af11580156200073f573d6000803e3d6000fd5b505050506200074f600362000a98565b505080806200075e9062000fdf565b9150506200063f565b5050565b6200077562000a05565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620007e7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007de90620010b3565b60405180910390fd5b620007f28162000aae565b50565b620007ff62000a05565b60005b8160ff16811015620009cb5760006200081c600262000a8a565b9050600060076000838152602001908152602001600020604051806040016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050905062015180600560009054906101000a900460ff1660ff16620008c39190620010d5565b42620008d0919062001120565b816000018181525050806020015173ffffffffffffffffffffffffffffffffffffffff1663a6f2ae3a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156200092657600080fd5b505af11580156200093b573d6000803e3d6000fd5b5050505080600760008481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050620009b3600262000a98565b50508080620009c29062000fdf565b91505062000802565b5062015180600560009054906101000a900460ff1660ff16620009ef9190620010d5565b42620009fc919062001120565b60048190555050565b62000a0f62000b72565b73ffffffffffffffffffffffffffffffffffffffff1662000a2f62000576565b73ffffffffffffffffffffffffffffffffffffffff161462000a88576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a7f90620011ab565b60405180910390fd5b565b600081600001549050919050565b6001816000016000828254019250508190555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b610e7b80620011ce83390190565b604051806040016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600062000c0362000bfd62000bf78462000bb8565b62000bd8565b62000bb8565b9050919050565b600062000c178262000be2565b9050919050565b600062000c2b8262000c0a565b9050919050565b62000c3d8162000c1e565b82525050565b600060208201905062000c5a600083018462000c32565b92915050565b6000819050919050565b62000c758162000c60565b82525050565b600060208201905062000c92600083018462000c6a565b92915050565b600080fd5b600060ff82169050919050565b62000cb58162000c9d565b811462000cc157600080fd5b50565b60008135905062000cd58162000caa565b92915050565b60006020828403121562000cf45762000cf362000c98565b5b600062000d048482850162000cc4565b91505092915050565b62000d188162000c60565b811462000d2457600080fd5b50565b60008135905062000d388162000d0d565b92915050565b60006020828403121562000d575762000d5662000c98565b5b600062000d678482850162000d27565b91505092915050565b600062000d7d8262000c0a565b9050919050565b62000d8f8162000d70565b82525050565b600060408201905062000dac600083018562000c6a565b62000dbb602083018462000d84565b9392505050565b600062000dcf8262000bb8565b9050919050565b62000de18162000dc2565b82525050565b600060208201905062000dfe600083018462000dd6565b92915050565b62000e0f8162000c9d565b82525050565b600060208201905062000e2c600083018462000e04565b92915050565b600062000e3f8262000bb8565b9050919050565b62000e518162000e32565b82525050565b600060208201905062000e6e600083018462000e46565b92915050565b600062000e818262000dc2565b9050919050565b62000e938162000e74565b811462000e9f57600080fd5b50565b60008135905062000eb38162000e88565b92915050565b60006020828403121562000ed25762000ed162000c98565b5b600062000ee28482850162000ea2565b91505092915050565b62000ef68162000dc2565b811462000f0257600080fd5b50565b60008135905062000f168162000eeb565b92915050565b60006020828403121562000f355762000f3462000c98565b5b600062000f458482850162000f05565b91505092915050565b600062000f5b8262000c0a565b9050919050565b62000f6d8162000f4e565b82525050565b600060608201905062000f8a600083018662000f62565b62000f99602083018562000e46565b62000fa8604083018462000e04565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600062000fec8262000c60565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362001021576200102062000fb0565b5b600182019050919050565b600082825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006200109b6026836200102c565b9150620010a8826200103d565b604082019050919050565b60006020820190508181036000830152620010ce816200108c565b9050919050565b6000620010e28262000c60565b9150620010ef8362000c60565b9250828202620010ff8162000c60565b9150828204841483151762001119576200111862000fb0565b5b5092915050565b60006200112d8262000c60565b91506200113a8362000c60565b925082820190508082111562001155576200115462000fb0565b5b92915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000620011936020836200102c565b9150620011a0826200115b565b602082019050919050565b60006020820190508181036000830152620011c68162001184565b905091905056fe60806040523480156200001157600080fd5b5060405162000e7b38038062000e7b8339818101604052810190620000379190620002ca565b620000576200004b620000fd60201b60201c565b6200010560201b60201c565b82600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160146101000a81548160ff021916908360ff16021790555050505062000326565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001fb82620001ce565b9050919050565b60006200020f82620001ee565b9050919050565b620002218162000202565b81146200022d57600080fd5b50565b600081519050620002418162000216565b92915050565b60006200025482620001ce565b9050919050565b620002668162000247565b81146200027257600080fd5b50565b60008151905062000286816200025b565b92915050565b600060ff82169050919050565b620002a4816200028c565b8114620002b057600080fd5b50565b600081519050620002c48162000299565b92915050565b600080600060608486031215620002e657620002e5620001c9565b5b6000620002f68682870162000230565b9350506020620003098682870162000275565b92505060406200031c86828701620002b3565b9150509250925092565b610b4580620003366000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80639e382400116100665780639e382400146100e8578063a6f2ae3a14610106578063b2bdfa7b14610110578063ecd0c0c31461012e578063f2fde38b1461014c57610093565b80634355f1ec146100985780634e71d92d146100b6578063715018a6146100c05780638da5cb5b146100ca575b600080fd5b6100a0610168565b6040516100ad91906107b3565b60405180910390f35b6100be61018e565b005b6100c8610354565b005b6100d2610368565b6040516100df91906107ef565b60405180910390f35b6100f0610391565b6040516100fd9190610826565b60405180910390f35b61010e6103a4565b005b61011861051b565b60405161012591906107ef565b60405180910390f35b610136610541565b6040516101439190610862565b60405180910390f35b610166600480360381019061016191906108ae565b610567565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6101966105ea565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166308e530036040518163ffffffff1660e01b8152600401602060405180830381865afa158015610203573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102279190610919565b600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631c560305600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660646040518363ffffffff1660e01b81526004016102e792919061098b565b600060405180830381600087803b15801561030157600080fd5b505af1158015610315573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b61035c6105ea565b6103666000610668565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160149054906101000a900460ff1681565b6103ac6105ea565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166308e530036040518163ffffffff1660e01b8152600401602060405180830381865afa158015610419573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043d9190610919565b600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639ff054df600160149054906101000a900460ff166040518263ffffffff1660e01b81526004016104e791906109e5565b600060405180830381600087803b15801561050157600080fd5b505af1158015610515573d6000803e3d6000fd5b50505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61056f6105ea565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d590610a83565b60405180910390fd5b6105e781610668565b50565b6105f261072c565b73ffffffffffffffffffffffffffffffffffffffff16610610610368565b73ffffffffffffffffffffffffffffffffffffffff1614610666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065d90610aef565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061077961077461076f84610734565b610754565b610734565b9050919050565b600061078b8261075e565b9050919050565b600061079d82610780565b9050919050565b6107ad81610792565b82525050565b60006020820190506107c860008301846107a4565b92915050565b60006107d982610734565b9050919050565b6107e9816107ce565b82525050565b600060208201905061080460008301846107e0565b92915050565b600060ff82169050919050565b6108208161080a565b82525050565b600060208201905061083b6000830184610817565b92915050565b600061084c82610780565b9050919050565b61085c81610841565b82525050565b60006020820190506108776000830184610853565b92915050565b600080fd5b61088b816107ce565b811461089657600080fd5b50565b6000813590506108a881610882565b92915050565b6000602082840312156108c4576108c361087d565b5b60006108d284828501610899565b91505092915050565b60006108e6826107ce565b9050919050565b6108f6816108db565b811461090157600080fd5b50565b600081519050610913816108ed565b92915050565b60006020828403121561092f5761092e61087d565b5b600061093d84828501610904565b91505092915050565b6000819050919050565b6000819050919050565b600061097561097061096b84610946565b610754565b610950565b9050919050565b6109858161095a565b82525050565b60006040820190506109a060008301856107e0565b6109ad602083018461097c565b9392505050565b60006109cf6109ca6109c58461080a565b610754565b610950565b9050919050565b6109df816109b4565b82525050565b60006020820190506109fa60008301846109d6565b92915050565b600082825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610a6d602683610a00565b9150610a7882610a11565b604082019050919050565b60006020820190508181036000830152610a9c81610a60565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610ad9602083610a00565b9150610ae482610aa3565b602082019050919050565b60006020820190508181036000830152610b0881610acc565b905091905056fea2646970667358221220fe4fa72122ac8bbb964ed8a111db1c0d068e258b7150c228434f8ee12581ad7864736f6c63430008110033a2646970667358221220d04a8479595707ef9075c4f659176f9c8b402e538c414f7b820991bdc832ea7264736f6c63430008110033",
}

// BuyerABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyerMetaData.ABI instead.
var BuyerABI = BuyerMetaData.ABI

// BuyerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BuyerMetaData.Bin instead.
var BuyerBin = BuyerMetaData.Bin

// DeployBuyer deploys a new Ethereum contract, binding an instance of Buyer to it.
func DeployBuyer(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address, day uint8, addr common.Address) (common.Address, *types.Transaction, *Buyer, error) {
	parsed, err := BuyerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuyerBin), backend, owner, day, addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Buyer{BuyerCaller: BuyerCaller{contract: contract}, BuyerTransactor: BuyerTransactor{contract: contract}, BuyerFilterer: BuyerFilterer{contract: contract}}, nil
}

// Buyer is an auto generated Go binding around an Ethereum contract.
type Buyer struct {
	BuyerCaller     // Read-only binding to the contract
	BuyerTransactor // Write-only binding to the contract
	BuyerFilterer   // Log filterer for contract events
}

// BuyerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuyerSession struct {
	Contract     *Buyer            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyerCallerSession struct {
	Contract *BuyerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyerTransactorSession struct {
	Contract     *BuyerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyerRaw struct {
	Contract *Buyer // Generic contract binding to access the raw methods on
}

// BuyerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyerCallerRaw struct {
	Contract *BuyerCaller // Generic read-only contract binding to access the raw methods on
}

// BuyerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyerTransactorRaw struct {
	Contract *BuyerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuyer creates a new instance of Buyer, bound to a specific deployed contract.
func NewBuyer(address common.Address, backend bind.ContractBackend) (*Buyer, error) {
	contract, err := bindBuyer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Buyer{BuyerCaller: BuyerCaller{contract: contract}, BuyerTransactor: BuyerTransactor{contract: contract}, BuyerFilterer: BuyerFilterer{contract: contract}}, nil
}

// NewBuyerCaller creates a new read-only instance of Buyer, bound to a specific deployed contract.
func NewBuyerCaller(address common.Address, caller bind.ContractCaller) (*BuyerCaller, error) {
	contract, err := bindBuyer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyerCaller{contract: contract}, nil
}

// NewBuyerTransactor creates a new write-only instance of Buyer, bound to a specific deployed contract.
func NewBuyerTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyerTransactor, error) {
	contract, err := bindBuyer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyerTransactor{contract: contract}, nil
}

// NewBuyerFilterer creates a new log filterer instance of Buyer, bound to a specific deployed contract.
func NewBuyerFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyerFilterer, error) {
	contract, err := bindBuyer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyerFilterer{contract: contract}, nil
}

// bindBuyer binds a generic wrapper to an already deployed contract.
func bindBuyer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuyerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buyer *BuyerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buyer.Contract.BuyerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buyer *BuyerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buyer.Contract.BuyerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buyer *BuyerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buyer.Contract.BuyerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buyer *BuyerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buyer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buyer *BuyerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buyer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buyer *BuyerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buyer.Contract.contract.Transact(opts, method, params...)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Buyer *BuyerCaller) SECONDSINDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "SECONDS_IN_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Buyer *BuyerSession) SECONDSINDAY() (*big.Int, error) {
	return _Buyer.Contract.SECONDSINDAY(&_Buyer.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Buyer *BuyerCallerSession) SECONDSINDAY() (*big.Int, error) {
	return _Buyer.Contract.SECONDSINDAY(&_Buyer.CallOpts)
}

// Addr is a free data retrieval call binding the contract method 0x62ab1a77.
//
// Solidity: function _addr() view returns(address)
func (_Buyer *BuyerCaller) Addr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "_addr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x62ab1a77.
//
// Solidity: function _addr() view returns(address)
func (_Buyer *BuyerSession) Addr() (common.Address, error) {
	return _Buyer.Contract.Addr(&_Buyer.CallOpts)
}

// Addr is a free data retrieval call binding the contract method 0x62ab1a77.
//
// Solidity: function _addr() view returns(address)
func (_Buyer *BuyerCallerSession) Addr() (common.Address, error) {
	return _Buyer.Contract.Addr(&_Buyer.CallOpts)
}

// Day is a free data retrieval call binding the contract method 0x9e382400.
//
// Solidity: function _day() view returns(uint8)
func (_Buyer *BuyerCaller) Day(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "_day")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Day is a free data retrieval call binding the contract method 0x9e382400.
//
// Solidity: function _day() view returns(uint8)
func (_Buyer *BuyerSession) Day() (uint8, error) {
	return _Buyer.Contract.Day(&_Buyer.CallOpts)
}

// Day is a free data retrieval call binding the contract method 0x9e382400.
//
// Solidity: function _day() view returns(uint8)
func (_Buyer *BuyerCallerSession) Day() (uint8, error) {
	return _Buyer.Contract.Day(&_Buyer.CallOpts)
}

// ProjectOwner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Buyer *BuyerCaller) ProjectOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProjectOwner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Buyer *BuyerSession) ProjectOwner() (common.Address, error) {
	return _Buyer.Contract.ProjectOwner(&_Buyer.CallOpts)
}

// ProjectOwner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Buyer *BuyerCallerSession) ProjectOwner() (common.Address, error) {
	return _Buyer.Contract.ProjectOwner(&_Buyer.CallOpts)
}

// ClaimTracker is a free data retrieval call binding the contract method 0x222abc11.
//
// Solidity: function claimTracker() view returns(uint256 _value)
func (_Buyer *BuyerCaller) ClaimTracker(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "claimTracker")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimTracker is a free data retrieval call binding the contract method 0x222abc11.
//
// Solidity: function claimTracker() view returns(uint256 _value)
func (_Buyer *BuyerSession) ClaimTracker() (*big.Int, error) {
	return _Buyer.Contract.ClaimTracker(&_Buyer.CallOpts)
}

// ClaimTracker is a free data retrieval call binding the contract method 0x222abc11.
//
// Solidity: function claimTracker() view returns(uint256 _value)
func (_Buyer *BuyerCallerSession) ClaimTracker() (*big.Int, error) {
	return _Buyer.Contract.ClaimTracker(&_Buyer.CallOpts)
}

// FatoryTracker is a free data retrieval call binding the contract method 0xb25f4202.
//
// Solidity: function fatoryTracker() view returns(uint256 _value)
func (_Buyer *BuyerCaller) FatoryTracker(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "fatoryTracker")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FatoryTracker is a free data retrieval call binding the contract method 0xb25f4202.
//
// Solidity: function fatoryTracker() view returns(uint256 _value)
func (_Buyer *BuyerSession) FatoryTracker() (*big.Int, error) {
	return _Buyer.Contract.FatoryTracker(&_Buyer.CallOpts)
}

// FatoryTracker is a free data retrieval call binding the contract method 0xb25f4202.
//
// Solidity: function fatoryTracker() view returns(uint256 _value)
func (_Buyer *BuyerCallerSession) FatoryTracker() (*big.Int, error) {
	return _Buyer.Contract.FatoryTracker(&_Buyer.CallOpts)
}

// FetchAddr is a free data retrieval call binding the contract method 0x08e53003.
//
// Solidity: function fetchAddr() view returns(address)
func (_Buyer *BuyerCaller) FetchAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "fetchAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FetchAddr is a free data retrieval call binding the contract method 0x08e53003.
//
// Solidity: function fetchAddr() view returns(address)
func (_Buyer *BuyerSession) FetchAddr() (common.Address, error) {
	return _Buyer.Contract.FetchAddr(&_Buyer.CallOpts)
}

// FetchAddr is a free data retrieval call binding the contract method 0x08e53003.
//
// Solidity: function fetchAddr() view returns(address)
func (_Buyer *BuyerCallerSession) FetchAddr() (common.Address, error) {
	return _Buyer.Contract.FetchAddr(&_Buyer.CallOpts)
}

// IdTracker is a free data retrieval call binding the contract method 0x1b7fe6b6.
//
// Solidity: function idTracker() view returns(uint256 _value)
func (_Buyer *BuyerCaller) IdTracker(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "idTracker")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdTracker is a free data retrieval call binding the contract method 0x1b7fe6b6.
//
// Solidity: function idTracker() view returns(uint256 _value)
func (_Buyer *BuyerSession) IdTracker() (*big.Int, error) {
	return _Buyer.Contract.IdTracker(&_Buyer.CallOpts)
}

// IdTracker is a free data retrieval call binding the contract method 0x1b7fe6b6.
//
// Solidity: function idTracker() view returns(uint256 _value)
func (_Buyer *BuyerCallerSession) IdTracker() (*big.Int, error) {
	return _Buyer.Contract.IdTracker(&_Buyer.CallOpts)
}

// Mints is a free data retrieval call binding the contract method 0x6350b2c2.
//
// Solidity: function mints(uint256 ) view returns(uint256 claimday, address mint)
func (_Buyer *BuyerCaller) Mints(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Claimday *big.Int
	Mint     common.Address
}, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "mints", arg0)

	outstruct := new(struct {
		Claimday *big.Int
		Mint     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Claimday = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Mint = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Mints is a free data retrieval call binding the contract method 0x6350b2c2.
//
// Solidity: function mints(uint256 ) view returns(uint256 claimday, address mint)
func (_Buyer *BuyerSession) Mints(arg0 *big.Int) (struct {
	Claimday *big.Int
	Mint     common.Address
}, error) {
	return _Buyer.Contract.Mints(&_Buyer.CallOpts, arg0)
}

// Mints is a free data retrieval call binding the contract method 0x6350b2c2.
//
// Solidity: function mints(uint256 ) view returns(uint256 claimday, address mint)
func (_Buyer *BuyerCallerSession) Mints(arg0 *big.Int) (struct {
	Claimday *big.Int
	Mint     common.Address
}, error) {
	return _Buyer.Contract.Mints(&_Buyer.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buyer *BuyerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buyer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buyer *BuyerSession) Owner() (common.Address, error) {
	return _Buyer.Contract.Owner(&_Buyer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buyer *BuyerCallerSession) Owner() (common.Address, error) {
	return _Buyer.Contract.Owner(&_Buyer.CallOpts)
}

// BuyToken is a paid mutator transaction binding the contract method 0xfdafc62c.
//
// Solidity: function buyToken(uint8 number) returns()
func (_Buyer *BuyerTransactor) BuyToken(opts *bind.TransactOpts, number uint8) (*types.Transaction, error) {
	return _Buyer.contract.Transact(opts, "buyToken", number)
}

// BuyToken is a paid mutator transaction binding the contract method 0xfdafc62c.
//
// Solidity: function buyToken(uint8 number) returns()
func (_Buyer *BuyerSession) BuyToken(number uint8) (*types.Transaction, error) {
	return _Buyer.Contract.BuyToken(&_Buyer.TransactOpts, number)
}

// BuyToken is a paid mutator transaction binding the contract method 0xfdafc62c.
//
// Solidity: function buyToken(uint8 number) returns()
func (_Buyer *BuyerTransactorSession) BuyToken(number uint8) (*types.Transaction, error) {
	return _Buyer.Contract.BuyToken(&_Buyer.TransactOpts, number)
}

// ClaimToken is a paid mutator transaction binding the contract method 0xeebb426d.
//
// Solidity: function claimToken(uint8 number) returns()
func (_Buyer *BuyerTransactor) ClaimToken(opts *bind.TransactOpts, number uint8) (*types.Transaction, error) {
	return _Buyer.contract.Transact(opts, "claimToken", number)
}

// ClaimToken is a paid mutator transaction binding the contract method 0xeebb426d.
//
// Solidity: function claimToken(uint8 number) returns()
func (_Buyer *BuyerSession) ClaimToken(number uint8) (*types.Transaction, error) {
	return _Buyer.Contract.ClaimToken(&_Buyer.TransactOpts, number)
}

// ClaimToken is a paid mutator transaction binding the contract method 0xeebb426d.
//
// Solidity: function claimToken(uint8 number) returns()
func (_Buyer *BuyerTransactorSession) ClaimToken(number uint8) (*types.Transaction, error) {
	return _Buyer.Contract.ClaimToken(&_Buyer.TransactOpts, number)
}

// Factory is a paid mutator transaction binding the contract method 0x20bc14d6.
//
// Solidity: function factory(uint8 number) returns()
func (_Buyer *BuyerTransactor) Factory(opts *bind.TransactOpts, number uint8) (*types.Transaction, error) {
	return _Buyer.contract.Transact(opts, "factory", number)
}

// Factory is a paid mutator transaction binding the contract method 0x20bc14d6.
//
// Solidity: function factory(uint8 number) returns()
func (_Buyer *BuyerSession) Factory(number uint8) (*types.Transaction, error) {
	return _Buyer.Contract.Factory(&_Buyer.TransactOpts, number)
}

// Factory is a paid mutator transaction binding the contract method 0x20bc14d6.
//
// Solidity: function factory(uint8 number) returns()
func (_Buyer *BuyerTransactorSession) Factory(number uint8) (*types.Transaction, error) {
	return _Buyer.Contract.Factory(&_Buyer.TransactOpts, number)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buyer *BuyerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buyer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buyer *BuyerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Buyer.Contract.RenounceOwnership(&_Buyer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buyer *BuyerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Buyer.Contract.RenounceOwnership(&_Buyer.TransactOpts)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd1d80fdf.
//
// Solidity: function setAddr(address addr) returns()
func (_Buyer *BuyerTransactor) SetAddr(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Buyer.contract.Transact(opts, "setAddr", addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd1d80fdf.
//
// Solidity: function setAddr(address addr) returns()
func (_Buyer *BuyerSession) SetAddr(addr common.Address) (*types.Transaction, error) {
	return _Buyer.Contract.SetAddr(&_Buyer.TransactOpts, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd1d80fdf.
//
// Solidity: function setAddr(address addr) returns()
func (_Buyer *BuyerTransactorSession) SetAddr(addr common.Address) (*types.Transaction, error) {
	return _Buyer.Contract.SetAddr(&_Buyer.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buyer *BuyerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Buyer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buyer *BuyerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Buyer.Contract.TransferOwnership(&_Buyer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buyer *BuyerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Buyer.Contract.TransferOwnership(&_Buyer.TransactOpts, newOwner)
}

// BuyerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Buyer contract.
type BuyerOwnershipTransferredIterator struct {
	Event *BuyerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BuyerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyerOwnershipTransferred)
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
		it.Event = new(BuyerOwnershipTransferred)
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
func (it *BuyerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyerOwnershipTransferred represents a OwnershipTransferred event raised by the Buyer contract.
type BuyerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buyer *BuyerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BuyerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Buyer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BuyerOwnershipTransferredIterator{contract: _Buyer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buyer *BuyerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BuyerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Buyer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyerOwnershipTransferred)
				if err := _Buyer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buyer *BuyerFilterer) ParseOwnershipTransferred(log types.Log) (*BuyerOwnershipTransferred, error) {
	event := new(BuyerOwnershipTransferred)
	if err := _Buyer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
