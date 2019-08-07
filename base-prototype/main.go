package main

import (
	"fmt"
	"github.com/c0ding/complexChain/base-prototype/block"
)

func main() {
	var (
		b *block.Block
	)

	b = block.NewGenesisBlock("aaaa")

	fmt.Println(b)
}
