package imp

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
)

func CreateGenesis(data string) {
	BLC.NewBlockchainWithGenesisBlock(data)
}
