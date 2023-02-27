package main

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
)

type Client struct {
	c *ethclient.Client
}

func keyStoreToPrivateKey(filePath *string, pwd *string) (*ecdsa.PrivateKey, common.Address, error) {
	keyJSON, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
	}
	unlockedKey, err := keystore.DecryptKey(keyJSON, *pwd)
	//fmt.Println(unlockedKey)
	privateKey := unlockedKey.PrivateKey
	address := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privateKey, address, err
}

func (ec *Client) new0ETHTransaction(privateKey *ecdsa.PrivateKey, data []byte, fromAddress common.Address, toAddress common.Address) *types.Transaction {
	value := big.NewInt(0)
	gasLimit := uint64(21000)
	nonce, err := ec.c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := ec.c.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	return tx
}

func (ec *Client) signTransaction(tx *types.Transaction, privateKey *ecdsa.PrivateKey) *types.Transaction {
	chainID, err := ec.c.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx
}

func main() {

	ac := new(Client)
	client, err := ethclient.Dial("\\\\.\\pipe\\data/geth1.ipc")
	if err != nil {
		log.Fatal(err)
	}
	ac.c = client

	keyFile := "D:\\blockchain\\node1\\data\\keystore\\UTC--2023-02-22T01-47-12.020787200Z--cdb60b2bd9962915f8b6c545efadb8fd0fdedab1"
	pwd := "123456"
	privateKey, fromAddress, err := keyStoreToPrivateKey(&keyFile, &pwd)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(privateKey, fromAddress)

	//生成0ETH的交易，包含想发送的数据
	toAddress := common.HexToAddress("0x1c79ccf180a6508149a460a7730a62344673fee3")
	var data []byte
	//data := []byte("test")
	tx := ac.new0ETHTransaction(privateKey, data, fromAddress, toAddress)

	//交易签名
	signedTx := ac.signTransaction(tx, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//发送交易
	err = ac.c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
}
