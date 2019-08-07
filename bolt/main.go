package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var (
	db  *bolt.DB
	err error
)

const bucketname = "bucket01"

func main() {
	usageOpen()
	usageUpdateDB()
	usageViewDB()
}

func usageOpen() {

	if db, err = bolt.Open("my01.db", 0600, nil); err != nil {
		log.Panic(err)
	}

}

func usageUpdateDB() {
	if err = db.Update(func(tx *bolt.Tx) error {
		var (
			err    error
			bucket *bolt.Bucket
		)

		if bucket = tx.Bucket([]byte(bucketname)); bucket != nil {
			return nil
		}

		if bucket, err = tx.CreateBucket([]byte(bucketname)); err != nil {
			return fmt.Errorf("创建表失败%s", err)
		}

		if err = bucket.Put([]byte("l"), []byte("给你 100 BTC")); err != nil {
			return fmt.Errorf("写入数据失败%s", err)
		}

		return nil
	}); err != nil {
		log.Panic(err)
	}
}

func usageViewDB() {

	if err = db.View(func(tx *bolt.Tx) error {
		var (
			bucket *bolt.Bucket
			result []byte
		)
		if bucket = tx.Bucket([]byte(bucketname)); bucket == nil {
			return fmt.Errorf("获取表对象失败")
		}

		result = bucket.Get([]byte("l"))
		fmt.Println(string(result))

		return nil
	}); err != nil {
		log.Panic(err)
	}
}
