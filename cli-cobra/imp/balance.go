package imp

import (
	"fmt"
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
)

func GetBalance(address string) {
	fmt.Println(address)
	BLC.UnSpentTransationsWithAdress(address)
}
