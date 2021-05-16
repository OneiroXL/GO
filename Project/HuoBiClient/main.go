package main

import(
	"HuoBiClient/core/account"
	_"HuoBiClient/base/constant"
	"HuoBiClient/core/transaction"
	_"fmt"
)



func main(){
	account.UserInit()
	account.WalletInit()

	//transaction.PlaceOrder("buy-limit","dogeusdt",15,0.4)

	transaction.CancelOrderById("277478979474781")

	// account.WalletInit()
	// //获取大于0的余额币种
	// balanceInfoSlice := account.GetAmountGTXBalanceInfoSlice(0,constant.TRADE)
	// for _,v := range balanceInfoSlice{
	// 	fmt.Printf("币种:%v  状态:%s  金额:%v \n", v.Currency,v.Type,v.Amount)
	// }


}

