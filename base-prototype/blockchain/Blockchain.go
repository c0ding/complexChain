package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

const (
	dbName         = "blockchain.db"
	blockTableName = "blocks"
)

var (
	G_Blockchain *Blockchain
)

type Blockchain struct {
	Tip []byte //最新的区块的Hash
	DB  *bolt.DB
}

func (blockchain *Blockchain) Iterator() *BlockchainIterator {

	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

// 判断数据库是否存在
func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}

	return true
}

func (b *Blockchain) Printchain() {

	var (
		blockchainIterator *BlockchainIterator
	)

	blockchainIterator = b.Iterator()

	for {

		block := blockchainIterator.Next()

		fmt.Printf("Height：%d\n", block.Height)
		fmt.Printf("PrevBlockHash：%x\n", block.PreBlockHash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Printf("Timestamp：%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x\n", block.Hash)
		fmt.Printf("Nonce：%d\n", block.Nonce)

		fmt.Println()

		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}

}

// 创建单例
func setG_Blockchain(bc *Blockchain) {
	G_Blockchain = bc
}

func NewBlockchainWithGenesisBlock(data string) {
	var (
		db        *bolt.DB
		err       error
		blockHash []byte
	)

	if DBExists() {
		log.Println("创世区块已存在。。。。")
		os.Exit(1)
	}

	log.Println("正在创建创世区块。。。。")

	if db, err = bolt.Open(dbName, 0600, nil); err != nil {
		log.Panic(err)
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		var (
			bucket       *bolt.Bucket
			err          error
			genesisBlock = NewGenesisBlock(data)
		)

		if bucket = tx.Bucket([]byte(blockTableName)); bucket == nil {
			if bucket, err = tx.CreateBucket([]byte(blockTableName)); err != nil {
				log.Panic(err)
			}
		}

		// 存储创世区块到表中
		if err = bucket.Put(genesisBlock.Hash, genesisBlock.Serialize()); err != nil {
			log.Panic(err)
		}

		// 存储最新的区块的hash
		if err = bucket.Put([]byte("l"), genesisBlock.Hash); err != nil {
			log.Panic(err)
		}

		return nil
	}); err != nil {
		log.Panic(err)
	}

	setG_Blockchain(&Blockchain{blockHash, db})
}

func (b *Blockchain) AddBlock(data string) {

	if DBExists() == false {
		log.Println("创世区块不存在，先创建创世区块")
		os.Exit(1)
	}

	var (
		err error
	)

	if err = b.DB.Update(func(tx *bolt.Tx) error {

		var (
			bucket     *bolt.Bucket
			err        error
			blockBytes []byte
			block      *Block
			newBlock   *Block
		)

		//1. 获取表
		if bucket = tx.Bucket([]byte(blockTableName)); bucket == nil {
			log.Panic("获取表失败")
		}

		//2. 创建新区块

		blockBytes = bucket.Get(b.Tip)

		block = DeSerialize(blockBytes)
		newBlock = NewBlock(data, block.Height+1, block.Hash)

		if err = bucket.Put(newBlock.Hash, newBlock.Serialize()); err != nil {
			log.Panic(err)
		}

		//更新数据库里面"l"对应的hash，l存着最新的区块
		if err = bucket.Put([]byte("l"), newBlock.Hash); err != nil {
			log.Panic(err)
		}

		b.Tip = newBlock.Hash

		return nil
	}); err != nil {
		log.Panic(err)
	}

}

// 返回Blockchain对象
func BlockchainObject() *Blockchain {

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			// 读取最新区块的Hash
			tip = b.Get([]byte("l"))

		}

		return nil
	})

	setG_Blockchain(&Blockchain{tip, db})
	return &Blockchain{tip, db}
}
