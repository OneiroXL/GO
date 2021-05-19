package market

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)


//获取当前币种行情
func GetCurrentCurrencyBySymbols(symbols string) *market.CandlestickAskBid {

	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestickAskBid(symbols)
	if err != nil {
		applogger.Error(err.Error())
		return nil
	} 
	return resp
}