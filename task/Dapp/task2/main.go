package main

import (
	"context"
	"fmt"
	"go-ethereum/task/Dapp/task2/counter"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddr = "0x457367717a9CAF8DAd634571578Cc54042069ab4"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/OLIHAnkdXwaywPvwWxXiA")
	if err != nil {
		log.Fatal(err)
	}
	counterContract, err := counter.NewCounter(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("143611b795be4190b2c15a47e287a4c512fd71c4cdd9c8a9a28be628d81146b5")
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}

	tx, err := counterContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	value, err := counterContract.Count(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the counter number: ", value)
}
