package imp

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
	"github.com/c0ding/complexChain/base-prototype/common"
)

func Send(form, to, amount string) {

	//fmt.Printf("从地址：%s ，转到地址：%s %s钱 \n", form, to, amount)
	//fmt.Println(common.JSONToArray(form))
	//fmt.Println(common.JSONToArray(to))
	//fmt.Println(common.JSONToArray(amount))

	forms := common.JSONToArray(form)
	tos := common.JSONToArray(to)
	amounts := common.JSONToArray(amount)

	BLC.BlockchainObject()
	BLC.G_Blockchain.MineNewBlock(forms, tos, amounts)

}
