package main

import (
	"fmt"

	"geth/contract/bep20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	TokenAddress = "0xaec945e04baf28b135fa7c640f624f8d90f1c3a6"
)

func main() {
	client, err := ethclient.Dial("https://bsc.kyberengineering.io/")
	if err != nil {
		panic(err)
	}

	tokenContract, err := bep20.NewBEP20(common.HexToAddress(TokenAddress), client)
	if err != nil {
		panic(err)
	}

	decimals, err := tokenContract.Decimals(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	fmt.Println("decimals", decimals)
}
