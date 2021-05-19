package market

import (
	"github.com/shopspring/decimal"
)

type  CurrentCurrencyQuotationModel struct{

	//调整为新加坡时间的时间戳，单位秒，并以此作为此K线柱的id
	Id int64 

	//以基础币种计量的交易量（以滚动24小时计）
	TransactionAmount decimal.Decimal

	//交易次数（以滚动24小时计）
	TransactionCount int64

	//本阶段开盘价（以滚动24小时计）
	OpenPrice decimal.Decimal

	//本阶段收盘价（以滚动24小时计）
	ClosePrice decimal.Decimal

	//本阶段最低价（以滚动24小时计）
	MostLowPrice decimal.Decimal

	//本阶段最高价（以滚动24小时计）
	MostHightPrice decimal.Decimal

	//以报价币种计量的交易量（以滚动24小时计）
	Vol decimal.Decimal
	
}
