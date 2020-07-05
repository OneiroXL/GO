package CollyFund

import (
	"fmt"
	"log"
	"strconv"
	"fund/Model"
	"encoding/json"
	"github.com/gocolly/colly"
)

type CrawlingHandle struct {
	C *colly.Collector
	Body string
}

/*
Init
*/
func (this *CrawlingHandle) CrawlingInit()  {
	this.C = colly.NewCollector()
	this.UserAgentInit();


	this.C.OnRequest(func(r *colly.Request) {
		this.HeadersInit(r)
	})
	 
	this.C.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	 
	this.C.OnResponse(func(r *colly.Response) {
		this.Body = string(r.Body)
		//fmt.Println("Data:",this.Body)
	})
	 
	// this.C.OnHTML("a[href]", func(e *colly.HTMLElement) {
	 
	// 	e.Request.Visit(e.Attr("href"))
	 
	// })
	 
	// this.C.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	 
	// 	fmt.Println("First column of a table row:", e.Text)
	 
	// })
	 
	// this.C.OnScraped(func(r *colly.Response) {
	 
	// 	fmt.Println("Finished")
	 
	// })
}

/*
Start
*/
func (this *CrawlingHandle) Start(url string) {
	this.C.Visit(url)
	this.BadyHandle()
}

func (this *CrawlingHandle) BadyHandle()  {
	var bodyData string = this.Body
	var bodyJson = bodyData[42:len(bodyData) -1]
	//fmt.Println(bodyJson)

	fundValuationDetailModel := Model.FundValuationDetailModel{};
	//将Json转化为对象
	err := json.Unmarshal([]byte(bodyJson), &fundValuationDetailModel)

	if err != nil {
		fmt.Println(err)
	}
	
	var fundValuationDetail Model.FundValuationDetail = fundValuationDetailModel.Datas[0]
	actualIncrease,_ := strconv.ParseFloat(fundValuationDetail.ActualIncrease,32) 
	fmt.Println("----------------------------华安成长创新混合----------------------------")
	fmt.Printf(
		"基金编码:%s\n" +
		"基金净值:%s\n" +
		"估计净值:%s\n" +
		"估计涨幅:%s\n" +
		"实际涨幅:%f\n" +
		"更新时间:%s\n",
		fundValuationDetail.FundCode,
		fundValuationDetail.FundNetWorth,
		fundValuationDetail.EstimateNetWorth,
		fundValuationDetail.EstimateIncrease,
		actualIncrease * 100 ,
		fundValuationDetail.UpdateTime,
	)
}


/*
init Headers
*/
func (this *CrawlingHandle) HeadersInit(r *colly.Request) {
	// r.Headers.Set("Host", "query.sse.com.cn")
    // r.Headers.Set("Connection", "keep-alive")
    // r.Headers.Set("Accept", "*/*")
    // r.Headers.Set("Origin", "http://www.sse.com.cn")
    // r.Headers.Set("Referer", "http://www.sse.com.cn/assortment/stock/list/share/") //关键头 如果没有 则返回 错误
    // r.Headers.Set("Accept-Encoding", "gzip, deflate")
    // r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
}

/*
init UserAgent
*/
func (this *CrawlingHandle) UserAgentInit() {
	this.C.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
}