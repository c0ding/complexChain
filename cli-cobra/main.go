package main

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
	"github.com/c0ding/complexChain/cli-cobra/cmd"
)

func main() {

	var (
		blockchain *BLC.Blockchain
	)

	if blockchain == nil {
		blockchain = BLC.NewBlockchainWithGenesisBlock()
	}

	defer blockchain.DB.Close()
	cmd.Execute()

}
