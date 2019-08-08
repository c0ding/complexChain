package blockchain

import (
	"github.com/boltdb/bolt"
	"log"
)

const (
	dbName         = "blockchain.db"
	blockTableName = "blocks"
)

type Blockchain struct {
	Tip []byte //最新的区块的Hash
	DB  *bolt.DB
}

func NewBlockchainWithGenesisBlock() *Blockchain {
	var (
		genesisBlock = NewGenesisBlock("genesis data")
		db           *bolt.DB
		err          error
		blockHash    []byte
	)

	if db, err = bolt.Open(dbName, 0600, nil); err != nil {
		log.Panic(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		var (
			bucket *bolt.Bucket
			err    error
		)

		if bucket = tx.Bucket([]byte(blockTableName)); bucket == nil {
			if bucket, err = tx.CreateBucket([]byte(blockTableName)); err != nil {
				log.Panic(err)
			}
		}
		// 存储最新的区块的hash
		if err = bucket.Put([]byte("l"), genesisBlock.Hash); err != nil {
			log.Panic(err)
		}

		blockHash = genesisBlock.Hash
		return nil
	})

	return &Blockchain{blockHash, db}

}

//func (b *Blockchain) AddBlock(data string, height int64, preBlockHash []byte) {
//	var (
//		block *Block
//	)
//	block = NewBlock(data, height, preBlockHash)
//	b.Blocks = append(b.Blocks, block)
//
//}
