package imp

import "fmt"

func Send(form, to, amount string) {
fmt.Printf("从地址：%s ，转到地址：%s %s钱 \n",form,to,amount)
}