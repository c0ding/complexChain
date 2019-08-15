package imp

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
)

func CreateGenesis(txs []*BLC.Transaction) {
	BLC.NewBlockchainWithGenesisBlock(txs)
}
