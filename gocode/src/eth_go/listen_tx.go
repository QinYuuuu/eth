package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func GetNewBlock(client *ethclient.Client, sub ethereum.Subscription, headers chan *types.Header) (*types.Block, error) {
	select {
	case err := <-sub.Err():
		log.Fatal(err)
		return nil, err
	case header := <-headers:
		block, err := client.BlockByHash(context.Background(), header.Hash())
		if err != nil {
			log.Fatal(err)
		}
		/*
			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		*/
		return block, err
	}
}

func GetTxData(client *ethclient.Client, block *types.Block) {
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
			fmt.Println(msg.From().Hex())
		}
	}
}

func main() {
	client, err := ethclient.Dial("\\\\.\\pipe\\data/geth1.ipc")
	if err != nil {
		log.Fatal(err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		block, err := GetNewBlock(client, sub, headers)
		if err != nil {
			log.Fatal(err)
		}
		GetTxData(client, block)
	}
}
