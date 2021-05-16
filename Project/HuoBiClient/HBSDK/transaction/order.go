package transaction

import (
	"HuoBiClient/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/order"
	"strconv"
	"fmt"
)

//下单
/*
方向：

buy: 买
sell: 卖
行为类型：

market：市价单，该类型订单仅需指定下单金额或下单数量，不需要指定价格，订单在进入撮合时，会直接与对手方进行成交，直至金额或数量低于最小成交金额或成交数量为止。
limit：限价单，该类型订单需指定下单价格，下单数量。
limit-maker：只做maker单，即限价挂单，该订单在进入撮合时，只能作为maker进入市场深度，若订单会作为taker成交，则会被直接取消。
ioc：立即成交或取消（immediately or cancel），该订单在进入撮合后，若不能直接成交，则会被直接取消（部分成交后，剩余部分也会被取消）。
limit-fok： 立即完全成交否则完全取消（fill or kill），该订单在进入撮合后，若不能立即完全成交，则会被完全取消。
market-grid：网格交易市价单（暂不支持API下单）。
limit-grid：网格交易限价单（暂不支持API下单）。
stop-limit：止盈止损单，设置高于或低于市场价格的订单，当订单到达触发价格后，才会正式的进入撮合队列。该类型已被策略委托订单代替，请使用策略委托订单。
订单来源 (source)：

spot-api：现货API交易
margin-api：逐仓杠杆API交易
super-margin-api：全仓杠杆API交易
c2c-margin-api：C2C杠杆API交易
grid-trading-sys：网格交易（暂不支持API下单）
订单状态 (state):

created：已创建，该状态订单尚未进入撮合队列。
submitted : 已挂单等待成交，该状态订单已进入撮合队列当中。
partial-filled : 部分成交，该状态订单在撮合队列当中，订单的部分数量已经被市场成交，等待剩余部分成交。
filled : 已成交。该状态订单不在撮合队列中，订单的全部数量已经被市场成交。
partial-canceled : 部分成交撤销。该状态订单不在撮合队列中，此状态由partial-filled转化而来，订单数量有部分被成交，但是被撤销。
canceling : 撤销中。该状态订单正在被撤销的过程中，因订单最终需在撮合队列中剔除才会被真正撤销，所以此状态为中间过渡态。
canceled : 已撤销。该状态订单不在撮合订单中，此状态订单没有任何成交数量，且被成功撤销。
相关ID

order-id : 订单的唯一编号
client-order-id : 客户自定义ID，该ID在下单时传入，与下单成功后返回的order-id对应，该ID 24小时内有效。 允许的字符包括字母(大小写敏感)、数字、下划线 (_)和横线(-)，最长64位
match-id : 订单在撮合中的顺序编号
trade-id : 成交的唯一编号
*/
func PlaceOrder(accountId int64, orderType string ,symbol string,num float64,price float64) {
	client := new(client.OrderClient).Init(config.AccessKey,config.Secret, config.Host)
	request := order.PlaceOrderRequest{
		AccountId: strconv.FormatInt(accountId,10),
		Type:      orderType,//"buy-limit",
		Source:    "spot-api",
		Symbol:    symbol,
		Price:     strconv.FormatFloat(price, 'E', -1, 32 ),
		Amount:    strconv.FormatFloat(num, 'E', -1, 32),
	}
	resp, err := client.PlaceOrder(&request)
	fmt.Printf("%+v \n",request)
	fmt.Printf("%+v \n",resp)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			applogger.Info("Place order successfully, order id: %s", resp.Data)
		case "error":
			applogger.Error("Place order error: %s", resp.ErrorMessage)
		}
	}
}

// 取消订单
func CancelOrderById(orderId string)  {
	client := new(client.OrderClient).Init(config.AccessKey,config.Secret, config.Host)
	resp, err := client.CancelOrderById(orderId)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		switch resp.Status {
		case "ok":
			applogger.Info("Cancel order successfully, order id: %s", resp.Data)
		case "error":
			applogger.Info("Cancel order error: %s", resp.ErrorMessage)
		}
	}
}