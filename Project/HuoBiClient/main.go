package main

import(
	"HuoBiClient/core/account"
	_"HuoBiClient/base/constant"
	_"HuoBiClient/core/transaction"
	"fmt"
	"HuoBiClient/core/market"
	"github.com/shopspring/decimal"
	"time"
)



func main(){
	account.UserInit()
	account.WalletInit()

	//下单
	// transaction.PlaceOrder("buy-limit","dogeusdt",15,0.4)

	// transaction.CancelOrderById("277478979474781")

	// account.WalletInit()
	// //获取大于0的余额币种
	// balanceInfoSlice := account.GetAmountGTXBalanceInfoSlice(0,constant.TRADE)
	// for _,v := range balanceInfoSlice{
	// 	fmt.Printf("币种:%v  状态:%s  金额:%v \n", v.Currency,v.Type,v.Amount)
	// }

	Test();
}

func Test(){

	for {
		currentCurrencyQuotationModel  := market.GetCurrentCurrencyBySymbols("dogeusdt")

		if(currentCurrencyQuotationModel == nil){
			continue
		}

		fmt.Printf("%+v \n",currentCurrencyQuotationModel)
	
		//计算涨幅
		upOrDownPrice := currentCurrencyQuotationModel.ClosePrice.Sub(currentCurrencyQuotationModel.OpenPrice)
	
		upOrDownRate := upOrDownPrice.Div(currentCurrencyQuotationModel.OpenPrice)
	
		fmt.Printf("涨幅:%v  涨幅率:%v%% \n",upOrDownPrice,upOrDownRate.Mul(decimal.NewFromInt(100)))

		time.Sleep(time.Second * 2)
	}


}

