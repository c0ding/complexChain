package main

import "fmt"
import BLC "github.com/c0ding/complexChain/base-prototype/blockchain"

func main() {

	var (
		blockchain  *BLC.Blockchain
		currenBlock *BLC.Block //当前区块，添加节点之前
	)

	blockchain = BLC.NewBlockchainWithGenesisBlock()
	currenBlock = blockchain.Blocks[len(blockchain.Blocks)-1]
	blockchain.AddBlock("100 to me", currenBlock.Height+1, currenBlock.Hash)

	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocks)

}

func test1() {
	//var (
	//	block       *BLC.Block
	//	blockchain  *BLC.Blockchain
	//	currenBlock *BLC.Block //当前区块，添加节点之前
	//)
	//
	//block = BLC.NewGenesisBlock("aaaa")
	//
	//fmt.Println(block)
	//fmt.Println("---------")
	//blockchain = BLC.NewBlockchainWithGenesisBlock()
	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)
	//fmt.Println(blockchain.Blocks[0])
	//
	//fmt.Println("---------")
	//currenBlock = blockchain.Blocks[len(blockchain.Blocks)-1]
	//blockchain.AddBlock("100 to me", currenBlock.Height+1, currenBlock.Hash)
	//
	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)
	//fmt.Println(blockchain.Blocks[0])
	//fmt.Println(blockchain.Blocks[1])
}
