package imp

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
)

func AddDada2Block(data string) {
	BLC.G_Blockchain.AddBlock(data)
}
