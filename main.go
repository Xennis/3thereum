package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Xennis/3thereum/contract/bancorcontractregistry"
	"github.com/Xennis/3thereum/contract/bancornetwork"
	"github.com/Xennis/3thereum/contract/kyberrate"
	"github.com/Xennis/3thereum/contract/uniswapexchange"
	"github.com/Xennis/3thereum/contract/uniswapfactory"
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
	amout     *big.Int
}

//func etherToWei(val *big.Int) *big.Int {
//	return new(big.Int).Mul(val, big.NewInt(params.Ether))
//}

func weiToEther(val *big.Int) *big.Int {
	return new(big.Int).Div(val, big.NewInt(params.Ether))
}

func uniswapRate(ctx context.Context, client *ethclient.Client, pair pair) (*big.Int, error) {
	const (
		uniswapFactoryAddress = "0xc0a47dfe034b400b47bdad5fecda2621de6c4d95"
	)

	factoryToken, err := uniswapfactory.NewToken(common.HexToAddress(uniswapFactoryAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create factory token: %v", err)
	}
	exchangeAddress, err := factoryToken.GetExchange(&bind.CallOpts{
		Context: ctx,
	}, pair.outAdress)
	if err != nil {
		return nil, fmt.Errorf("failed to get exchange: %v", err)
	}

	exchangeToken, err := uniswapexchange.NewToken(exchangeAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create exchange token: %v", err)
	}
	rate, err := exchangeToken.GetEthToTokenInputPrice(&bind.CallOpts{
		Context: ctx,
	}, pair.amout)
	if err != nil {
		return nil, fmt.Errorf("failed to get input price: %v", err)
	}
	return rate, nil
}

func kyberRate(ctx context.Context, client *ethclient.Client, pair pair) (*big.Int, *big.Int, error) {
	const (
		kyberRateAddress = "0x9AAb3f75489902f3a48495025729a0AF77d4b11e"
	)

	token, err := kyberrate.NewToken(common.HexToAddress(kyberRateAddress), client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create rate token: %v", err)
	}
	rate, err := token.GetExpectedRate(&bind.CallOpts{
		Context: ctx,
	}, pair.inAddress, pair.outAdress, pair.amout)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get expected rate: %v", err)
	}
	expectedRate := weiToEther(rate.ExpectedRate)
	worstRate := weiToEther(rate.WorstRate)
	return expectedRate.Mul(expectedRate, pair.amout), worstRate.Mul(worstRate, pair.amout), nil
}

func bancorRate(ctx context.Context, client *ethclient.Client, pair pair) (*big.Int, error) {
	const (
		// https://docs.bancor.network/ethereum-contracts/addresses
		bancorContractRegistry = "0x52Ae12ABe5D8BD778BD5397F99cA900624CfADD4"

		// https://docs.bancor.network/ethereum-contracts/addresses
		bancorNetworkContract = "BancorNetwork"
	)

	crToken, err := bancorcontractregistry.NewToken(common.HexToAddress(bancorContractRegistry), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create network token: %v", err)
	}

	var networkContractAddr [32]byte
	copy(networkContractAddr[:], []byte(bancorNetworkContract))

	adr, err := crToken.AddressOf(&bind.CallOpts{
		Context: ctx,
	}, networkContractAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create network token: %v", err)
	}

	networkToken, err := bancornetwork.NewToken(adr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create network token: %v", err)
	}

	path, err := networkToken.ConversionPath(&bind.CallOpts{
		Context: ctx,
	}, pair.inAddress, pair.outAdress)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversation path: %v", err)
	}

	rate, err := networkToken.RateByPath(&bind.CallOpts{
		Context: ctx,
	}, path, pair.amout)
	if err != nil {
		return nil, fmt.Errorf("failed to get rate by path: %v", err)
	}

	return rate, nil
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

	// Token addresses can be found on https://etherscan.io/tokens
	for _, p := range []pair{
		{
			inSymbol:  "ETH",
			inAddress: common.HexToAddress(ethAddress),
			outSymbol: "DAI",
			outAdress: common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
			amout:     big.NewInt(1),
		},
		{
			inSymbol:  "ETH",
			inAddress: common.HexToAddress(ethAddress),
			outSymbol: "BND",
			outAdress: common.HexToAddress("0xB8c77482e45F1F44dE1745F52C74426C631bDD52"),
			amout:     big.NewInt(1),
		},
		{
			inSymbol:  "ETH",
			inAddress: common.HexToAddress(ethAddress),
			outSymbol: "UNI",
			outAdress: common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"),
			amout:     big.NewInt(1),
		},
	} {
		uniswap, err := uniswapRate(ctx, client, p)
		if err != nil {
			log.Printf("%s -> %s failed to get uniswap rate: %v", p.inSymbol, p.outSymbol, err)
			continue
		}
		kyberExpected, kyberWorst, err := kyberRate(ctx, client, p)
		if err != nil {
			log.Printf("%s -> %s failed to get kyber rate: %v", p.inSymbol, p.outSymbol, err)
			continue
		}
		bancorRate, err := bancorRate(ctx, client, p)
		if err != nil {
			log.Printf("%s -> %s failed to get bancor rate: %v", p.inSymbol, p.outSymbol, err)
			continue
		}
		fmt.Printf("%s -> %s: uniswap=%s, bancor=%s, kyber=%s/%s (expected/worst)\n", p.inSymbol, p.outSymbol, uniswap, bancorRate, kyberExpected, kyberWorst)
	}
}
