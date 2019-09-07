package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	// 区块高度
	Height int64
	//上一个区块hash
	PreBlockHash []byte
	//交易数据
	Txs []*Transaction
	//时间戳
	TimeStamp int64
	//本区块的hash
	Hash []byte
	// 随机值
	Nonce int64
}

// 需要将Txs转换成[]byte
func (block *Block) HashTransactions() []byte {

	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]

}

func NewBlock(txs []*Transaction, height int64, preBlockHash []byte) *Block {

	var (
		timestamp int64
		hash      []byte
		nonce     int64
		pow       *Pow
	)
	timestamp = time.Now().Unix()
	block := &Block{
		Txs:          txs,
		Height:       height,
		PreBlockHash: preBlockHash,
		TimeStamp:    timestamp,
		Hash:         nil,
		Nonce:        0,
	}

	// 调用Pow，返回 hash , Nonce
	pow = NewPow(block)
	hash, nonce = pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	fmt.Println()
	return block
}

func NewGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func (b *Block) Serialize() []byte {
	var (
		encoder *gob.Encoder
		err     error
		result  bytes.Buffer
	)

	encoder = gob.NewEncoder(&result)
	if err = encoder.Encode(b); err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeSerialize(blockBytes []byte) *Block {

	var (
		block   Block
		decoder *gob.Decoder
		err     error
	)
	decoder = gob.NewDecoder(bytes.NewReader(blockBytes))
	if err = decoder.Decode(&block); err != nil {

		log.Panic(err)

	}
	return &block
}
