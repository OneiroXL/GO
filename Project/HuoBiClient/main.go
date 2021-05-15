package main

import(
	"HuoBiClient/core/account"
	"HuoBiClient/base/constant"
	"fmt"
)



func main(){
	account.UserInit()
	account.WalletInit()

	//获取所有币种信息
	allBalanceInfo := account.GetAllBalanceInfo();
	fmt.Println(len(allBalanceInfo))
	// for _,v := range allBalanceInfo{
	// 	fmt.Println(v)
	// }

	//获取大于0的余额币种
	balanceInfoSlice := account.GetAmountGTXBalanceInfoSlice(0,constant.TRADE)
	for _,v := range balanceInfoSlice{
		fmt.Printf("币种:%v  状态:%s  金额:%v \n", v.Currency,v.Type,v.Amount)
	}
}

