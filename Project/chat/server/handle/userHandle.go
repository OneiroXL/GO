package handle

import (
	"chat/model/user"
	"chat/common/tools"
	"chat/model/base"
	"encoding/json"
	"net"
)

var (
	userMap map[int]user.UserInfo = make(map[int]user.UserInfo,3)
)

type UserHandle struct {
	Conn net.Conn
	TcpTool tools.TcpTool
}

func NewUserHandle(conn net.Conn,tcpTool tools.TcpTool) *UserHandle {
	userHandle := &UserHandle{
		Conn:conn,
		TcpTool:tcpTool,
	}
	return userHandle
}

/*
用户登录
*/
func (this *UserHandle) UserLogin(userLoginModel user.UserLoginModel){
	userInfo := user.UserInfo{
		ID:userLoginModel.UserID,
		Conn:this.Conn,
	}
	this.AddUser(userInfo)

	interactive := base.Interactive{}
	interactive.Type = 100
	interactive.Message = "成功"
	interactive.Status = 10000

	interactiveJSON,_ := json.Marshal(interactive)
	
	this.TcpTool.Write([]byte(interactiveJSON))
}


/*
添加用户
*/
func (this *UserHandle) AddUser(userInfo user.UserInfo){
	userMap[userInfo.ID] = userInfo
}

