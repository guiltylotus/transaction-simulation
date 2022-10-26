package main

import (
	"bytes"
	"fmt"
	metaaggregator "geth/contract/meta_aggregator"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	hexResult := "0x000000000000000000000000000000000000000000000000014a41f8a9544181000000000000000000000000000000000000000000000000000000000007d0cd"

	ab, err := abi.JSON(bytes.NewBufferString(metaaggregator.MetaAggregatorMetaData.ABI))
	if err != nil {
		panic(err)
	}
	res, err := ab.Unpack("swapSimpleMode", hexutil.MustDecode(hexResult))
	if err != nil {
		panic(err)
	}
	fmt.Println("res", res)
}
