package main

import (
	"context"
	"fmt"
	"geth/contract"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	b, err := ioutil.ReadFile("./wallet/UTC--2022-07-21T07-57-06.751826000Z--7ca04051b273a8ce59ebcc260bb2c10da93d2059")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "minh")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("/Users/nguyenducminh/Library/Ethereum/goerli/geth.ipc")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	//privateKey := key.PrivateKey
	//fmt.Println(privateKey)

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := contract.DeployIncreaseEvent(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-----------------------------------")
}
