package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
	"io"
	"net/http"
	"strings"
)

type TokenInfo struct {
	ID       string
	Symbol   string
	Address  string
	Decimals int
	Chain    int
}

type GetCoinListResponse struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}

func main() {
	fmt.Println("future-----------------")
	fclient := futures.NewClient("", "")
	resp, err := fclient.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	tokenMap := make(map[string]*TokenInfo)
	for _, symbol := range resp.Symbols {
		if strings.ToLower(symbol.QuoteAsset) != "usdt" {
			continue
		}
		fmt.Println("symbol", symbol.Symbol)
		tokenInfo := TokenInfo{
			Symbol: symbol.BaseAsset,
			Chain:  56,
		}
		tokenMap[strings.ToLower(symbol.BaseAsset)] = &tokenInfo
	}

	fmt.Println("len future", len(tokenMap))

	httpClient := http.DefaultClient
	coinResp, err := httpClient.Get("https://api.coingecko.com/api/v3/coins/list?include_platform=true")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(coinResp.Body)
	if err != nil {
		panic(err)
	}
	_ = coinResp.Body.Close()
	var cResp []GetCoinListResponse
	if err := json.Unmarshal(data, &cResp); err != nil {
		panic(err)
	}

	fmt.Println("len coin", len(cResp))
	for _, coin := range cResp {
		token, ok := tokenMap[strings.ToLower(coin.Symbol)]
		if !ok {
			continue
		}
		platform, ok := coin.Platforms["binance-smart-chain"]
		if !ok {
			continue
		}
		token.Address = platform
		token.ID = coin.ID
	}

	for _, value := range tokenMap {
		if value.Address == "" {
			continue
		}
		fmt.Println(
			"name", fmt.Sprintf("'%s'", value.ID),
			"chain", fmt.Sprintf("'%s'", "56"),
			"symbol", fmt.Sprintf("'%s'", value.Symbol),
			"address", fmt.Sprintf("'%s'", value.Address),
			"decimals", 18)
	}
}
