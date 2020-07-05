package main

import (
	"fund/CollyFund"
	"time"
	_"fmt"
)

func main(){
	ch := CollyFund.CrawlingHandle{}
	ch.CrawlingInit()
	for {
		ch.Start("https://fundmobapi.eastmoney.com/FundMApi/FundValuationDetail.ashx?callback=jQuery311023667329794458625_1592546620210&FCODE=007460&deviceid=Wap&plat=Wap&product=EFund&version=2.0.0&Uid=&_=1592546620211")
		time.Sleep(2 *time.Second)
	}
	
}

