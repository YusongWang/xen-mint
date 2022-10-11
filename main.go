package main

import (
	"context"
	"crypto/ecdsa"

	"fmt"
	"log"
	"math/big"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

type config struct {
	ChainApi   string `env:"URL"`
	PrivateKey string `env:"PRIVATEKEY,unset"`
}

type Key struct {
	ID         uint   `gorm:"column:id;primarykey"`
	Address    string `gorm:"column:address;index;not null"`
	PrivateKey string `gorm:"column:privateKey;not null"`
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
				Name:    "wallet",
				Aliases: []string{"c"},
				Usage:   "create wallets",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() < 2 {
						fmt.Println("Args less then 2, first is dbnamne,seced is wallet count")
						os.Exit(0)
					}
					db_name := cCtx.Args().First()
					count_str := cCtx.Args().Get(1)
					count, err := strconv.Atoi(count_str)
					if err != nil {
						fmt.Println("error: ", err)
						//ExitProcess()
						os.Exit(0)
					}

					create_wallets(db_name, count)
					//fmt.Println("added task: ", cCtx.Args().First())
					fmt.Println("Create success..")
					return nil
				},
			},
			{
				Name:    "mint",
				Aliases: []string{"m"},
				Usage:   "mint LIKE on bsc",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() < 2 {
						fmt.Println("Args less then 2, first is dbnamne,seced is wallet count")
						os.Exit(0)
					}
					db_name := cCtx.Args().First()
					start_str := cCtx.Args().Get(1)
					start, err := strconv.Atoi(start_str)
					if err != nil {
						fmt.Println("error: ", err)
						//ExitProcess()
						os.Exit(0)
					}
					if start <= 10 {
						start = 1000
					}

					count_str := cCtx.Args().Get(2)
					count, err := strconv.Atoi(count_str)
					if err != nil {
						fmt.Println("error: ", err)
						//ExitProcess()
						os.Exit(0)
					}

					openDatabase(db_name)
					batch_mint(start, count)
					fmt.Printf("success for mint %d wallet address\n", count)
					return nil
				},
			},
			{
				Name:    "transfer",
				Aliases: []string{"t"},
				Usage:   "Use one Wallet and transfer BNB to all address. Max Count is 50",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() < 2 {
						fmt.Println("Args less then 2, first is dbnamne,seced is wallet count")
						os.Exit(0)
					}
					db_name := cCtx.Args().First()
					start_str := cCtx.Args().Get(1)
					start, err := strconv.Atoi(start_str)
					if err != nil {
						fmt.Println("error: ", err)
						//ExitProcess()
						os.Exit(0)
					}
					if start <= 10 {
						start = 1000
					}

					count_str := cCtx.Args().Get(2)
					count, err := strconv.Atoi(count_str)
					if err != nil {
						fmt.Println("error: ", err)
						//ExitProcess()
						os.Exit(0)
					}

					openDatabase(db_name)
					batch_trasfer(start, count)

					fmt.Printf("success for transfer %d wallet address\n", count)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	//batch_mint()
	//batch_trasfer()
}

func batch_mint(start, count int) {
	kdb := Keys{}
	dbkeys, err := kdb.Scan(start, count)
	if err != nil {
		panic(err)
	}

	var wallets_privatekey []string
	for idx := range dbkeys {
		wallets_privatekey = append(wallets_privatekey, dbkeys[idx].PrivateKey)
	}

	for idx := range wallets_privatekey {
		if err := mint(idx, strings.TrimPrefix(wallets_privatekey[idx], "0x")); err != nil {
			fmt.Println("at count panic  plase start with this index ", idx)
			panic(err)
		}
	}
}

func NewKey() (Key, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Key{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return Key{}, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	//contractAddr := crypto.CreateAddress(address, 0)
	privateKeyHex := "0x" + hexutil.Encode(privateKeyBytes)[2:]

	return Key{Address: address.Hex(), PrivateKey: privateKeyHex}, nil
}

func mint(idx int, privateKeyStr string) error {
	key := privateKeyStr
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
	like, err := NewLikeBnb(common.HexToAddress("0xBDE5AbC1c689BaA94ac91eE1328064c59712418B"), conn)
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
	auth.GasLimit = uint64(287959) // in units
	auth.GasPrice = gasPrice

	_, err = like.ClaimRank(auth, big.NewInt(1))
	if err != nil {
		return err
	}

	fmt.Println("done with idx :", idx)
	return nil
}

func batch_trasfer(start, count int) {
	kdb := Keys{}
	dbkeys, err := kdb.Scan(start, count)
	if err != nil {
		panic(err)
	}

	var wallets []common.Address
	for idx := range dbkeys {
		wallets = append(wallets, common.HexToAddress(dbkeys[idx].Address))
	}

	fmt.Println(wallets)
	key := cfg.PrivateKey
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
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
	batchTrasfer, err := NewBatchTrasfer(common.HexToAddress("0xf2062ED85949824b72b1fe581D75f0eb9901b783"), conn)
	if err != nil {
		panic(err)
	}

	nonce, err := client.conn.PendingNonceAt(context.Background(), address)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.conn.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(signkey, chainId)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0).Mul(big.NewInt(1500000000000000), big.NewInt(int64(count))) // in wei
	auth.GasLimit = uint64(3000000)                                                        // in units
	auth.GasPrice = gasPrice

	_, err = batchTrasfer.Transfer(auth, wallets, big.NewInt(1500000000000000))
	if err != nil {
		fmt.Println("trasfer() failed ")
		panic(err)
	}

	//fmt.Println(trasfer)
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

func create_wallets(db_name string, count int) {
	fmt.Println("create address : ", count)
	openDatabase(db_name)
	// 生成地址
	var stopped int32
	keysChan := make(chan Key, BatchSize)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for atomic.LoadInt32(&stopped) == 0 {
				key, err := NewKey()
				if err != nil {
					panic(err)
				}

				keysChan <- key
			}
		}()
	}

	// 保存账户
	start := time.Now()
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		var size int
		keys := make([]Key, 0)
		for key := range keysChan {
			size++
			keys = append(keys, key)
			if len(keys) < BatchSize && size < count {
				continue
			}

			err := Keys{}.Save(keys)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d accounts created, time: %v\n", size, time.Since(start))
			keys = keys[:0]

			if size >= count {
				atomic.StoreInt32(&stopped, 1)
				waitGroup.Done()
				return
			}
		}
	}()

	waitGroup.Wait()
}
