package Model

import (
	"fund/Model/Base"
)

type FundValuationDetail struct {
	FundCode string `json:"fundcode"` //基金编码
	FundNetWorth string `json:"dwjz"` //基金净值
	EstimateNetWorth string `json:"gsz"` //估计净值
	EstimateIncrease string `json:"gszzl"` //估计涨幅
	ActualIncrease string `json:"gszze"` //实际涨幅
	UpdateTime string `json:"gztime"` //更新时间
}

type FundValuationDetailModel struct {
	Base.ResultPageModel
	Datas []FundValuationDetail
}
