package transaction

import (
	AccountCore "HuoBiClient/core/account"
	HBSDKTransaction "HuoBiClient/HBSDK/transaction"
)

//下单
func  PlaceOrder( orderType string ,symbol string,num float64,price float64)  {
		//获取用户信息
		userInfo := AccountCore.GetUserByType("spot");

		//下单
		HBSDKTransaction.PlaceOrder(userInfo.Id,orderType,symbol,num,price)
}

// 取消订单
func CancelOrderById(orderId string)  {
	HBSDKTransaction.CancelOrderById(orderId)
}