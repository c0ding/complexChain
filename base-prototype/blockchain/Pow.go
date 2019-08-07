package blockchain

type Pow struct {
	Block *Block
}

func NewPow(block *Block) *Pow {
	return &Pow{Block: block}
}

func (p *Pow) Run() ([]byte, int64) {
	return nil, 0
}

//// 所有字段统一转化为字节数组，在进行hash
//// 封装一个工具方法
//func (b *Block) SetHash() {
//	var (
//		height     []byte
//		time       []byte
//		blockBytes []byte
//		hash       [32]byte
//	)
//
//	height = common.Int2Bytes(b.Height)
//	time = common.Timestamp2Bytes(b.TimeStamp)
//	blockBytes = bytes.Join([][]byte{height, b.PreBlockHash, b.Data, time}, []byte{})
//	hash = sha256.Sum256(blockBytes)
//	b.Hash = hash[:]
//}
