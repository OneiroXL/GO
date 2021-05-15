package account

import (
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"HuoBiClient/config"
	"github.com/huobirdcenter/huobi_golang/pkg/model/account"
)


//获取用户信息
func GetUserInfo() []account.AccountInfo {
	client := new(client.AccountClient).Init(config.AccessKey,config.Secret, config.Host)
	resp, err := client.GetAccountInfo()
	if err != nil {
		applogger.Error("Get account error: %s", err)
		return nil
	} 
	return resp
}