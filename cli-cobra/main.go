
package main

import (
	"github.com/c0ding/complexChain/cli-cobra/cmd"
	//BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
)

func main() {

	//var (
	//	blockchain *BLC.Blockchain
	//)
	//
	//if blockchain == nil {
	//	blockchain = BLC.NewBlockchainWithGenesisBlock()
	//}
	//
	//defer blockchain.DB.Close()
	cmd.Execute()

}
