package main

import (
	"context"
	"crypto/ecdsa"
	"strings"
	"time"

	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
)

func ToDecimal(value *big.Int, decimals int) decimal.Decimal {

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

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
	proxy, err := NewProxy(common.HexToAddress(cfg.ContractAddress), conn)
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
	tx, err := proxy.Increase(auth, big.NewInt(int64(cfg.BuyerFacter)))
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
	buyer, err := NewProxy(common.HexToAddress(cfg.ContractAddress), conn)
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

	xenABI, err := abi.JSON(strings.NewReader(ABI))

	if err != nil {
		panic(err)
	}

	data, err := xenABI.Pack("claimRank", big.NewInt(int64(cfg.Day)))
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	tx, err := buyer.Execute(auth, big.NewInt(0), big.NewInt(100), common.HexToAddress(cfg.XenAddress), data)
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

	blockNumber, err := conn.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	balance, err := conn.BalanceAt(context.Background(), address, big.NewInt(int64(blockNumber)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Blance:%s \n", ToDecimal(balance, 18))

	//conn.SyncProgress(ctx context.Context)
	client := NewEthClient(conn, chainId)

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

	xenAddr, tx, _, err := DeployProxy(auth, client.conn, big.NewInt(int64(cfg.BuyerNumber)))
	if err != nil {
		return err
	}

	for {
		time.Sleep(time.Second * 3)
		_, isPeading, err := conn.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			panic(err)
		}

		if isPeading {
			fmt.Println("Wait for confim! ...")
		} else {
			fmt.Printf("confim: %s \n address: %s\n", tx.Hash(), xenAddr.Hex())
			break
		}
	}

	blockNumber, err = conn.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	newbalance, err := conn.BalanceAt(context.Background(), address, big.NewInt(int64(blockNumber)))
	if err != nil {
		panic(err)
	}

	gasUse := big.NewInt(0).Sub(balance, newbalance)
	fmt.Printf("Blance:%s Gas Use : %s \n", ToDecimal(newbalance, 18), ToDecimal(gasUse, 18).String())
	return nil
}
