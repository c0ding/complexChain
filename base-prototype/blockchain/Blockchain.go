package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
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

// 创建单例
func setG_Blockchain(bc *Blockchain) {
	G_Blockchain = bc
}

// 返回Blockchain对象
func BlockchainObject() *Blockchain {

	if DBExists() == false {
		log.Println("创世区块不存在，先创建创世区块")
		os.Exit(1)
	}

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

func (b *Blockchain) Printchain() {

	if DBExists() == false {
		log.Println("创世区块不存在，先创建创世区块")
		os.Exit(1)
	}

	var (
		blockchainIterator *BlockchainIterator
	)

	blockchainIterator = b.Iterator()

	for {

		block := blockchainIterator.Next()

		fmt.Printf("Height：%d\n", block.Height)
		fmt.Printf("PrevBlockHash：%x\n", block.PreBlockHash)
		//fmt.Printf("Txs：%s\n", block.Txs)
		fmt.Println("------------------------------")
		fmt.Printf("Txs:")
		for _, tx := range block.Txs {

			fmt.Printf("%x\n", tx.TxHash)
			fmt.Println("Vins:")
			for _, in := range tx.Vins {
				fmt.Printf("%x\n", in.TxHash)
				fmt.Printf("%d\n", in.Vout)
				fmt.Printf("%s\n", in.ScriptSig)
			}

			fmt.Println("Vouts:")
			for _, out := range tx.Vouts {
				fmt.Println(out.Value)
				fmt.Println(out.ScriptPubKey)
			}
		}

		fmt.Println("------------------------------")

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

func NewBlockchainWithGenesisBlock(address string) {
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
			bucket *bolt.Bucket
			err    error

			//genesisBlock = NewGenesisBlock(txs)

			genesisTX = NewCoinbaseTransaction(address)

			genesisBlock = NewGenesisBlock([]*Transaction{genesisTX})
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

func (b *Blockchain) AddBlock(txs []*Transaction) {

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
		newBlock = NewBlock(txs, block.Height+1, block.Hash)

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

// 如果一个地址对应的TXOutput未花费，那么这个Transaction就应该添加到数组中返回
func UnSpentTransationsWithAdress(address string) []*Transaction {

	return nil
}

// 挖掘新的区块
func (blockchain *Blockchain) MineNewBlock(from []string, to []string, amount []string) {

	//	$ ./bc send -from '["juncheng"]' -to '["zhangqiang"]' -amount '["2"]'
	//	[juncheng]
	//	[zhangqiang]
	//	[2]

	//1.建立一笔交易

	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)

	value, _ := strconv.Atoi(amount[0])

	tx := NewSimpleTransaction(from[0], to[0], value)
	fmt.Println(tx)

	//1. 通过相关算法建立Transaction数组

	var txs []*Transaction
	txs = append(txs, tx)

	var block *Block

	blockchain.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b != nil {

			hash := b.Get([]byte("l"))

			blockBytes := b.Get(hash)

			block = DeSerialize(blockBytes)

		}

		return nil
	})

	//2. 建立新的区块
	block = NewBlock(txs, block.Height+1, block.Hash)

	//将新区块存储到数据库
	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {

			b.Put(block.Hash, block.Serialize())

			b.Put([]byte("l"), block.Hash)

			blockchain.Tip = block.Hash

		}
		return nil
	})

}
