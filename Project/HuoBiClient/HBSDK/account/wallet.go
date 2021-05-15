package account

import (
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"HuoBiClient/config"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
	"strconv"
)


//获取金融信息
func GetBalanceInfo(accountId int64) *account.AccountBalance {
	client := new(client.AccountClient).Init(config.AccessKey,config.Secret, config.Host)
	resp, err := client.GetAccountBalance(strconv.FormatInt(accountId,10))
	if err != nil {
		applogger.Error("Get account balance error: %s", err)
		return nil
	}
	return resp
}