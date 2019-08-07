package blockchain

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchainWithGenesisBlock() *Blockchain {
	var (
		genesisBlock = NewGenesisBlock("genesis data")
	)

	return &Blockchain{Blocks: []*Block{genesisBlock}}

}
