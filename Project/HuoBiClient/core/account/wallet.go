package account

import (
	"HuoBiClient/model/account"
	"strconv"
	HBSDKAccount "HuoBiClient/HBSDK/account"
)


var (
	//金融信息map[币种][状态]{信息}
	balanceInfoMap map[string]map[string]*account.BalanceModel = make(map[string]map[string]*account.BalanceModel,50)
)

//初始化用户信息
func WalletInit(){

	//获取用户信息
	userInfo := GetUserByType("spot");
	//请求用户信息
	accountBalanceInfo := HBSDKAccount.GetBalanceInfo(userInfo.Id);

	if(accountBalanceInfo != nil){
		for _,balance := range accountBalanceInfo.List {

			balanceModel := new(account.BalanceModel)

			balanceModel.Currency = balance.Currency;
			balanceModel.Type = balance.Type;

			balanceFloat,_ := strconv.ParseFloat(balance.Balance,64)
			balanceModel.Amount = balanceFloat

			if(balanceInfoMap[balanceModel.Currency] == nil){
				balanceInfoMap[balanceModel.Currency] = make(map[string]*account.BalanceModel,2)
			}

			balanceInfoMap[balanceModel.Currency][balanceModel.Type] = balanceModel
		}
	}

}

//获取全部金额信息
func GetAllBalanceInfo() map[string]map[string]*account.BalanceModel {
	return balanceInfoMap
}

//获取金额信息依据类型
func GetBalanceInfoByCurrency(currency string,statusType string) *account.BalanceModel {
	return balanceInfoMap[currency][statusType]
}

//获取金额大于X的金额信息
func GetAmountGTXBalanceInfoSlice(x float64,statusType string) []*account.BalanceModel {

	balanceModelSlice := make([]*account.BalanceModel,0); 

	for _,v := range balanceInfoMap{
		if(v[statusType].Amount > x){
			balanceModelSlice = append(balanceModelSlice,v[statusType])
		}
	}

	return balanceModelSlice
}
