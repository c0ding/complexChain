package main

import (
	"fmt"
	"github.com/c0ding/complexChain/base-prototype/blockchain"
)

func main() {
	var (
		block             *blockchain.Block
		genesisBlockchain *blockchain.Blockchain
	)

	block = blockchain.NewGenesisBlock("aaaa")

	fmt.Println(block)
	fmt.Println("---------")
	genesisBlockchain = blockchain.NewBlockchainWithGenesisBlock()
	fmt.Println(genesisBlockchain)
	fmt.Println(genesisBlockchain.Blocks)
	fmt.Println(genesisBlockchain.Blocks[0])
}
