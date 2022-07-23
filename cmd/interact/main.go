package main

import (
	"context"
	"fmt"
	"geth/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

type NewID struct {
	Date *big.Int
	Id   *big.Int
}

func main() {
	// Create an IPC based RPC connection to a remote node
	// NOTE update the path to the ipc file!
	client, err := ethclient.Dial("/Users/nguyenducminh/Library/Ethereum/goerli/geth.ipc")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	//store, err := contract.NewStorage(common.HexToAddress("0xE3718FA212e2F13333d7f1464e83013E22F73BfF"), client)
	//if err != nil {
	//	log.Fatalf("Failed to instantiate Storage contract: %v", err)
	//}
	//
	contractAddress := common.HexToAddress("0xead507e4fa6674ab296f6b8194F31bf24E4ac3C6")
	iContract, err := contract.NewIncreaseEvent(contractAddress, client)
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

	IncreaseIContractID(iContract, auth)

	SubscribeEvent(client, iContract, contractAddress)

	//time.Sleep(5 * time.Second)
	GetIContractID(iContract)
	//Store(store, auth, big.NewInt(500))
	//Retreive(store)
}

func SubscribeEvent(
	client *ethclient.Client,
	iContract *contract.IncreaseEvent,
	contractAddress common.Address) {

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs := make(chan types.Log)
	contractAbi, err := abi.JSON(strings.NewReader(string(contract.IncreaseEventMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal("err", err)
			return
		case vLog := <-logs:
			var event NewID
			err := contractAbi.UnpackIntoInterface(&event, "NewID", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("vlog", "date", event.Date, "id", event.Id) // pointer to event log

			return
		}
	}
}

func IncreaseIContractID(iContract *contract.IncreaseEvent, auth *bind.TransactOpts) {
	tx, err := iContract.Increase(auth)
	if err != nil {
		log.Fatalf("Failed to increase id: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
	fmt.Println("-------------------------------------")
}

func GetIContractID(iContract *contract.IncreaseEvent) *big.Int {
	value, err := iContract.Get(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GetIContractID", value)
	return value
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
	//fmt.Println("Store")

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
