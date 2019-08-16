package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type TXInput struct {
	// 1. 交易的Hash
	TxHash      []byte
	// 2. 存储TXOutput在Vout里面的索引
	Vout      int
	// 3. 用户名
	ScriptSig string
}

type TXOutput struct {
	Value int64
	ScriptPubKey string  //用户名
}



type Transaction struct {

	//1. 交易hash
	TxHash []byte

	//2. 输入
	Vins []*TXInput

	//3. 输出
	Vouts []*TXOutput

}


func NewCoinbaseTransaction(address string) *Transaction {

	//代表消费
	txInput := &TXInput{[]byte{},-1,"Genesis Data"}

	txOutput := &TXOutput{10,address}

	txCoinbase := &Transaction{[]byte{},[]*TXInput{txInput},[]*TXOutput{txOutput}}

	//设置hash值
	txCoinbase.HashTransaction()


	return txCoinbase
}



func (tx *Transaction) HashTransaction()  {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}
