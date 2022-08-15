package bb

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	WALLET   = "wallet"
	KNC      = "knc"
	USDT     = "usdt"
	PRL      = "prl"
	KSROUTER = "router"
)

var (
	bsc = map[string]common.Address{
		WALLET:   common.HexToAddress("0x54F42802c6B381682c48B530f234De0a896Af1DD"), // ("0x43D5c49A397410d67F26B9455A30f212d1EcC27A"),
		KNC:      common.HexToAddress("0xfe56d5892bdffc7bf58f2e84be1b2c32d21c308b"),
		USDT:     common.HexToAddress("0x55d398326f99059ff775485246999027b3197955"),
		PRL:      common.HexToAddress("0xd07e82440a395f3f3551b42da9210cd1ef4f8b24"),
		KSROUTER: common.HexToAddress("0x00555513Acf282B42882420E5e5bA87b44D8fA6E"),
	}
	ether = map[string]common.Address{
		WALLET:   common.HexToAddress("0x54F42802c6B381682c48B530f234De0a896Af1DD"),
		KNC:      common.HexToAddress("0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202"),
		USDT:     common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
		KSROUTER: common.HexToAddress("0x00555513Acf282B42882420E5e5bA87b44D8fA6E"),
	}
	book = ether
)
