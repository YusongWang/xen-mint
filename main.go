package main

import (
	"context"
	"crypto/ecdsa"

	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

type config struct {
	ChainApi        string `env:"URL"`
	PrivateKey      string `env:"PRIVATEKEY,unset"`
	BuyerFacter     uint8  `env:"FACTER"`
	BuyerNumber     uint8  `env:"NUMBER"`
	GasLimit        uint64 `env:"GAS"`
	ContractAddress string `env:"BUYADDR,unset"`
	XenAddress      string `env:XENADDR`
	Day             uint8  `env:"DAY"`
}

type Key struct {
	ID         uint   `gorm:"column:id;primarykey"`
	Address    string `gorm:"column:address;index;not null"`
	PrivateKey string `gorm:"column:privateKey;not null"`
}

type EthClient struct {
	conn    *ethclient.Client
	chainId *big.Int
}

func NewEthClient(conn *ethclient.Client, chain *big.Int) EthClient {
	return EthClient{
		conn, chain,
	}
}

func (Key) TableName() string {
	return "keys"
}

var cfg config

func main() {
	godotenv.Load()
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	fmt.Println("RPC:", cfg.ChainApi)
	fmt.Println("合约:", cfg.ContractAddress)
	fmt.Println("XEN合约:", cfg.XenAddress)
	fmt.Println("GasLimit:", cfg.GasLimit)
	fmt.Println("每次预先创建多少钱包:", cfg.BuyerFacter)
	fmt.Println("一次mint多少钱包:", cfg.BuyerNumber)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "deploy",
				Aliases: []string{"d"},
				Usage:   "deploy Contract Buyer",
				Action: func(cCtx *cli.Context) error {
					//create_wallets()
					deployContract()
					//fmt.Println("added task: ", cCtx.Args().First())
					fmt.Println("Deploy success..")
					return nil
				},
			},
			{
				Name:    "factor",
				Aliases: []string{"f"},
				Usage:   "PreCreate Contract for Buyer",
				Action: func(cCtx *cli.Context) error {
					batch_mint()
					fmt.Printf("成功创建 %d 个地址\n", cfg.BuyerFacter)
					return nil
				},
			},
			{
				Name:    "buyer",
				Aliases: []string{"b"},
				Usage:   "Mint the Token",
				Action: func(cCtx *cli.Context) error {
					if err := buyer(); err != nil {
						panic(err)
					}
					fmt.Printf("成功购买 %d 个地址\n", cfg.BuyerNumber)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func batch_mint() {
	if err := mint(); err != nil {
		panic(err)
	}
}

func mint() error {
	key := cfg.PrivateKey
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	signkey := privateKey
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	conn, err := ethclient.Dial(cfg.ChainApi)
	if err != nil {
		panic(err)
	}

	chainId, err := conn.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	client := NewEthClient(conn, chainId)
	buyer, err := NewBuyer(common.HexToAddress(cfg.ContractAddress), conn)
	if err != nil {
		return err
	}

	nonce, err := client.conn.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}

	gasPrice, err := client.conn.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(signkey, chainId)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(cfg.GasLimit) // in units
	auth.GasPrice = gasPrice
	tx, err := buyer.Factory(auth, cfg.BuyerFacter)
	if err != nil {
		return err
	}

	fmt.Println(err, tx.Hash())

	return nil
}

func buyer() error {
	key := cfg.PrivateKey
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	signkey := privateKey
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	conn, err := ethclient.Dial(cfg.ChainApi)
	if err != nil {
		panic(err)
	}

	chainId, err := conn.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	client := NewEthClient(conn, chainId)
	buyer, err := NewBuyer(common.HexToAddress(cfg.ContractAddress), conn)
	if err != nil {
		return err
	}

	nonce, err := client.conn.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}

	gasPrice, err := client.conn.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(signkey, chainId)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(cfg.GasLimit) // in units
	auth.GasPrice = gasPrice
	tx, err := buyer.BuyToken(auth, cfg.BuyerNumber)
	if err != nil {
		return err
	}

	fmt.Println(tx.Hash())
	return nil
}

func deployContract() error {
	key := cfg.PrivateKey
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	signkey := privateKey
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	conn, err := ethclient.Dial(cfg.ChainApi)
	if err != nil {
		panic(err)
	}

	chainId, err := conn.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	client := NewEthClient(conn, chainId)
	// buyer, err := NewBuyer(common.HexToAddress(cfg.ContractAddress), conn)
	// if err != nil {
	// 	return err
	// }

	nonce, err := client.conn.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}

	gasPrice, err := client.conn.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(signkey, chainId)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(cfg.GasLimit) // in units
	auth.GasPrice = gasPrice

	xen := common.HexToAddress(cfg.XenAddress)
	xenAddr, _, _, err := DeployBuyer(auth, client.conn, address, cfg.Day, xen)
	if err != nil {
		return err
	}

	fmt.Println(xenAddr.Hex())
	return nil
}
