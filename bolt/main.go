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

func main() {
	usageOpen()
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
		if bucket, err = tx.CreateBucket([]byte("bucket01")); err != nil {
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
