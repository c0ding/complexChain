package main

import (
	"github.com/c0ding/complexChain/base-prototype/block"
	"fmt"
)

func main() {
	var (
		b *block.Block
	)

	b = block.NewBlock("asd",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Println(b)
}