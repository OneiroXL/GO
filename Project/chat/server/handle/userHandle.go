package handle

import (
	"chat/model/user"
	"chat/common/tools"
	"chat/model/base"
	"chat/server/bll/accountInfoBLL"
	"net"
)

var (
	userMap map[string]user.UserInfo = make(map[string]user.UserInfo,3)
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
用户注册
*/
func (this *UserHandle) UserRegister(userRegisterModel user.UserRegisterModel){
	userBLL := accountInfoBLL.UserBLL{}
	
	interactiveResponse := base.InteractiveResponse{}
	interactiveResponse.Type = 103

	//是否已存在用户
	isExistUser := userBLL.IsExistUser(userRegisterModel.UserCode);
	if(isExistUser){
		interactiveResponse.Message = "用户已存在"
		interactiveResponse.Status = 10001
		this.TcpTool.Respose(interactiveResponse);
		return
	}
	//添加用户
	err := userBLL.AddUser(userRegisterModel)
	if(err != nil){
		interactiveResponse.Message = "注册出错"
		interactiveResponse.Status = 10002
		this.TcpTool.Respose(interactiveResponse);
		return
	} 

	interactiveResponse.Message = "注册成功"
	interactiveResponse.Status = 10000
	this.TcpTool.Respose(interactiveResponse);
}

/*
用户登录
*/
func (this *UserHandle) UserLogin(userLoginModel user.UserLoginModel){

	userBLL := accountInfoBLL.UserBLL{}

	interactiveResponse := base.InteractiveResponse{}
	interactiveResponse.Type = 101

	isOk,mes,err := userBLL.LoginVerification(userLoginModel)
	if(err != nil){
		interactiveResponse.Message = "登录错误"
		interactiveResponse.Status = 10002
	} else if(isOk){
		userInfo := user.UserInfo{
			UserCode:userLoginModel.UserCode,
			Conn:this.Conn,
		}
		this.AddUser(userInfo)

		interactiveResponse.Message = mes
		interactiveResponse.Status = 10000
	} else {
		interactiveResponse.Message = mes
		interactiveResponse.Status = 10001
	}
	
	this.TcpTool.Respose(interactiveResponse);
}


/*
添加用户
*/
func (this *UserHandle) AddUser(userInfo user.UserInfo){
	userMap[userInfo.UserCode] = userInfo
}

/*
移除用户
*/
func (this *UserHandle) RemoveUser(userCode string){
	delete(userMap,userCode)
}