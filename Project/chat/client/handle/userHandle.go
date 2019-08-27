package handle

import (
	"net"
	"encoding/json"
	"chat/model/user"
	"chat/model/base"
	"chat/common/tools"
	"fmt"
)

type UserHandle struct {
	Conn net.Conn
	TcpTool tools.TcpTool
}

/*
用户注册
*/
func (this *UserHandle) UserRegister(userRegisterModel user.UserRegisterModel) (string,error) {
	data,err := this.TcpTool.Request(userRegisterModel,102)
	return data.Message,err
}

/*
用户登录
*/
func (this *UserHandle) UserLogin(userLoginModel user.UserLoginModel) (error) {

	userLoginModelJSON,_ := json.Marshal(userLoginModel)

	//交互消息体
	interactiveRequest := base.InteractiveRequest{}
	interactiveRequest.Type = 100
	interactiveRequest.Data = string(userLoginModelJSON)

	interactiveRequestJSON,_ := json.Marshal(interactiveRequest)

	data,err := this.TcpTool.WriteAndRead(interactiveRequestJSON)

	if(data.Status == 10000){
		fmt.Println("登录成功")
	}else{
		fmt.Println(data.Message)
	}
	return err
}