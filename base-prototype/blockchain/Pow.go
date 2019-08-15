package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/c0ding/complexChain/base-prototype/common"
	"math/big"
)

const (
	targetBit = 16
)

type Pow struct {
	Block  *Block
	target *big.Int
}

func NewPow(block *Block) *Pow {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)

	return &Pow{Block: block, target: target}
}

func (p *Pow) prepareDate(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			p.Block.PreBlockHash,
			p.Block.HashTransactions(),
			common.Int2Bytes(int64(p.Block.Height)),
			common.Timestamp2Bytes(p.Block.TimeStamp),
			common.Int2Bytes(int64(nonce)),
			common.Int2Bytes(int64(targetBit)),
		},
		[]byte{},
	)

	return data
}

func (p *Pow) Run() ([]byte, int64) {

	var (
		nonce   int
		hashInt big.Int
		hash    [32]byte
	)
	nonce = 0
	fmt.Println()
	for {
		dataBytes := p.prepareDate(nonce)
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if p.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1
	}
	return hash[:], int64(nonce)
}

func (p *Pow) IsValid() bool {
	var (
		hashInt big.Int
	)

	hashInt.SetBytes(p.Block.Hash)

	if p.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}
