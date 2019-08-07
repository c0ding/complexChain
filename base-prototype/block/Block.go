package block

import (
	"bytes"
	"crypto/sha256"
	"github.com/c0ding/complexChain/base-prototype/common"
	"time"
)

type Block struct {
	// 区块高度
	Height int64
	//上一个区块hash
	PreBlockHash []byte
	//交易数据
	Data []byte
	//时间戳
	TimeStamp int64
	//本区块的hash
	Hash []byte
}

func NewBlock(data string, height int64, preBlockHash []byte) *Block {

	var (
		timestamp int64
	)
	timestamp = time.Now().Unix()
	block := &Block{
		Data:         []byte(data),
		Height:       height,
		PreBlockHash: preBlockHash,
		TimeStamp:    timestamp,
		Hash:         nil,
	}

	block.SetHash()
	return block

}
func NewGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

// 所有字段统一转化为字节数组，在进行hash
// 封装一个工具方法
func (b *Block) SetHash() {
	var (
		height     []byte
		time       []byte
		blockBytes []byte
		hash       [32]byte
	)

	height = common.Int2Bytes(b.Height)
	time = common.Timestamp2Bytes(b.TimeStamp)
	blockBytes = bytes.Join([][]byte{height, b.PreBlockHash, b.Data, time}, []byte{})
	hash = sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}
