package handle

import (
	"net"
	"encoding/json"
	"chat/model/user"
	"chat/model/base"
	"chat/common/tools"
)

type UserHandle struct {
	Conn net.Conn
	TcpTool tools.TcpTool
}

/*
用户登录
*/
func (this *UserHandle) UserLogin(userLoginModel user.UserLoginModel) error {

	userLoginModelJSON,_ := json.Marshal(userLoginModel)

	//交互消息体
	interactive := base.Interactive{}
	interactive.Type = 100
	interactive.Message = "用户登录"
	interactive.Data = string(userLoginModelJSON)

	interactiveJSON,_ := json.Marshal(interactive)

	this.TcpTool.Write(interactiveJSON)
}