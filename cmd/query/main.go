package main

import (
	"context"
	"fmt"
	"geth/contract/dai"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
}

// LogApproval ..
type LogApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/c8a0f577c41240ab90d542d4c1f9f1ba")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(15471000),
		ToBlock:   big.NewInt(15477408),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(dai.ContractMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	//logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	//logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		//fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		//fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		//case logTransferSigHash.Hex():
		//	fmt.Printf("Log Name: Transfer\n")
		//
		//	var transferEvent LogTransfer
		//
		//	err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	transferEvent.Src = common.HexToAddress(vLog.Topics[1].Hex())
		//	transferEvent.Dst = common.HexToAddress(vLog.Topics[2].Hex())
		//
		//	fmt.Printf("From: %s\n", transferEvent.Src.Hex())
		//	fmt.Printf("To: %s\n", transferEvent.Dst.Hex())
		//	fmt.Printf("Tokens: %s\n", transferEvent.Wad.String())

		case logApprovalSigHash.Hex():
			fmt.Println("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.Src = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Guy = common.HexToAddress(vLog.Topics[2].Hex())

			//if approvalEvent.Guy.String() != "0x1111111254fb6c44bac0bed2854e76f90643097d" {
			//	continue
			//}
			fmt.Printf("Token Owner: %s\n", approvalEvent.Src.Hex())
			fmt.Printf("Guy: %s\n", approvalEvent.Guy.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Wad.String())
		}
	}
}