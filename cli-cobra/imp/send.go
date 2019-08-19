package imp

import (
	BLC "github.com/c0ding/complexChain/base-prototype/blockchain"
	"github.com/c0ding/complexChain/base-prototype/common"
)

func Send(form, to, amount string) {

	//fmt.Printf("从地址：%s ，转到地址：%s %s钱 \n", form, to, amount)

	forms := common.JSONToArray(form)
	tos := common.JSONToArray(to)
	amounts := common.JSONToArray(amount)

	//fmt.Println(forms)
	//fmt.Println(tos)
	//fmt.Println(amounts)

	BLC.BlockchainObject()
	BLC.G_Blockchain.MineNewBlock(forms, tos, amounts)

}
