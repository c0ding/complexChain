package imp

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
)

func CreateGenesis(address string) {
	BLC.NewBlockchainWithGenesisBlock(address)
}
