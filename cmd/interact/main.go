package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"testgeth/contract"
)

func main() {
	// Create an IPC based RPC connection to a remote node
	// NOTE update the path to the ipc file!
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	store, err := contract.NewStorage(common.HexToAddress("0xE3718FA212e2F13333d7f1464e83013E22F73BfF"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	key := GetKey()
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	Store(store, auth, big.NewInt(200))
	Retreive(store)
}

func GetKey() *keystore.Key {
	b, err := ioutil.ReadFile("./wallet/UTC--2022-07-21T07-57-06.751826000Z--7ca04051b273a8ce59ebcc260bb2c10da93d2059")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "minh")
	if err != nil {
		log.Fatal(err)
	}

	return key
}

func Store(store *contract.Storage, auth *bind.TransactOpts, number *big.Int) {
	tx, err := store.Store(auth, number)
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
}

func Retreive(store *contract.Storage) *big.Int {
	value, err := store.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve value: %v", err)
	}
	fmt.Println("Int", value)

	return value
}
