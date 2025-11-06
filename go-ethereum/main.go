package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum/eventQuery"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/OLIHAnkdXwaywPvwWxXiA")
	if err != nil {
		log.Fatal(err)
	}
	//queryBlockHeader.Run(client)
	//queryBlockTX.Run(client)
	//queryReceipt.Run(client)
	//createAdd.Run(client)
	//transfer.Run(client)
	//queryBalance.Run(client)
	//queryCoinBalance.Run(client)
	//subscribeBlock.Run(client)
	//deployContract.AbigenDeploy(client)
	//deployContract.EtherDeploy(client)
	//executeContract.ExecuteContract(client)
	//eventQuery.EventQuery(client)
	eventQuery.SubEvent(client)
}
