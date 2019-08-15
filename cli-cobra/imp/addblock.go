package imp

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
)

func AddDada2Block(txs []*BLC.Transaction) {

	BLC.BlockchainObject()
	BLC.G_Blockchain.AddBlock(txs)
}
