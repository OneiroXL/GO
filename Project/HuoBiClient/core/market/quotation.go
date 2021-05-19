package market

import (
	"HuoBiClient/model/market"
	HBSDKMarket "HuoBiClient/HBSDK/market"
)


//获取当前币种行情
func GetCurrentCurrencyBySymbols(symbols string) *market.CurrentCurrencyQuotationModel {

	currentCurrencyQuotationModel := new(market.CurrentCurrencyQuotationModel)

	//获取行情数据
	candlestickAskBid := HBSDKMarket.GetCurrentCurrencyBySymbols(symbols)
	
	if candlestickAskBid == nil {
		return nil
	}

	currentCurrencyQuotationModel.Id = candlestickAskBid.Id;
	currentCurrencyQuotationModel.ClosePrice =candlestickAskBid.Close;
	currentCurrencyQuotationModel.MostHightPrice = candlestickAskBid.High;
	currentCurrencyQuotationModel.MostLowPrice = candlestickAskBid.Low;
	currentCurrencyQuotationModel.OpenPrice = candlestickAskBid.Open;
	currentCurrencyQuotationModel.TransactionAmount = candlestickAskBid.Amount
	currentCurrencyQuotationModel.TransactionCount = candlestickAskBid.Count
	currentCurrencyQuotationModel.Vol = candlestickAskBid.Vol

	return currentCurrencyQuotationModel
}