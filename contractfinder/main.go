package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Xennis/3thereum/contractfinder/contracts/uniswap"
)

const (
	envEthURL  = "ETF_URL"
	ethAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
)

type pair struct {
	inSymbol  string
	inAddress common.Address
	outSymbol string
	outAdress common.Address
}

func getExchangeAddress(ctx context.Context, client *ethclient.Client, pair pair) (*common.Address, error) {
	const (
		uniswapFactoryAddress = "0xc0a47dfe034b400b47bdad5fecda2621de6c4d95"
	)

	token, err := uniswap.NewToken(common.HexToAddress(uniswapFactoryAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate uniswap token: %v", err)
	}
	exchange, err := token.GetExchange(&bind.CallOpts{
		From:    pair.inAddress,
		Context: ctx,
	}, pair.outAdress)
	if err != nil {
		return nil, fmt.Errorf("failed to get exchange: %v", err)
	}
	return &exchange, nil
}

func main() {
	ethURL := os.Getenv(envEthURL)
	if ethURL == "" {
		log.Fatalf("env %s is required", envEthURL)
	}

	ctx := context.Background()

	client, err := ethclient.Dial(ethURL)
	if err != nil {
		log.Fatalf("failed to dial eth client: %v", err)
	}

	fmt.Println("Exchange addresses:")

	// Token addresses can be found on https://etherscan.io/tokens
	for _, p := range []pair{
		{
			inSymbol:  "ETH",
			inAddress: common.HexToAddress(ethAddress),
			outSymbol: "DAI",
			outAdress: common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
		},
		{
			inSymbol:  "ETH",
			inAddress: common.HexToAddress(ethAddress),
			outSymbol: "BND",
			outAdress: common.HexToAddress("0xB8c77482e45F1F44dE1745F52C74426C631bDD52"),
		},
		{
			inSymbol:  "ETH",
			inAddress: common.HexToAddress(ethAddress),
			outSymbol: "USD",
			outAdress: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
		},
	} {
		exchange, err := getExchangeAddress(ctx, client, p)
		if err != nil {
			log.Fatalf("failed to get exchange address %s -> %s: %v", p.inSymbol, p.outSymbol, err)
		}
		fmt.Printf("%s -> %s: %s\n", p.inSymbol, p.outSymbol, exchange)
	}
}
