package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func importKs(filePath String) {

}

func main() {
	client, err := ethclient.Dial("\\\\.\\pipe\\data/geth1.ipc")
	if err != nil {
		log.Fatal(err)
	}
	keyFile := "D:\\blockchain\\node1\\data\\keystore\\UTC--2023-02-22T01-47-12.020787200Z--cdb60b2bd9962915f8b6c545efadb8fd0fdedab1"
	toAddress := "0x1c79ccf180a6508149a460a7730a62344673fee3"
	pwd := "123456"
	privateKey, fromAddress, err := keyStoreToPrivateKey(keyFile, "123456")
	if err != nil {
		log.Fatal(err)
	}
}
