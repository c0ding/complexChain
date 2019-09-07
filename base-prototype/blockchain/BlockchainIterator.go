package blockchain

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockchainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockchainIterator *BlockchainIterator) Next() *Block {

	var (
		block *Block
		err   error
	)

	if err = blockchainIterator.DB.View(func(tx *bolt.Tx) error {

		var (
			bucket *bolt.Bucket
		)

		bucket = tx.Bucket([]byte(blockTableName))

		if bucket != nil {
			currentBloclBytes := bucket.Get(blockchainIterator.CurrentHash)
			//  获取到当前迭代器里面的currentHash所对应的区块
			block = DeSerialize(currentBloclBytes)

			// 更新迭代器里面CurrentHash
			blockchainIterator.CurrentHash = block.PreBlockHash
		}

		return nil
	}); err != nil {
		log.Panic(err)
	}

	return block

}
