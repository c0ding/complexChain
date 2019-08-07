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

func (b *Blockchain) AddBlock(data string, height int64, preBlockHash []byte) {
	var (
		block *Block
	)
	block = NewBlock(data, height, preBlockHash)
	b.Blocks = append(b.Blocks, block)

}
