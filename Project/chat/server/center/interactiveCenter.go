package center

import (
	"chat/model/base"
	"chat/model/message"
	"net"
	"chat/common/tools"
	"chat/server/handle"
	"encoding/json"
	"chat/model/user"
)

/*
交互中心
*/
type InteractiveCenter struct{
	Conn net.Conn
	TcpTool tools.TcpTool
}

/*
交互处理
*/
func (this *InteractiveCenter) InteractiveHandle(interactiveRequest base.InteractiveRequest) error {
	switch interactiveRequest.Type {
		case 100:{
			userHandle := handle.NewUserHandle(this.Conn,this.TcpTool)

			userLoginModel := user.UserLoginModel{}
			json.Unmarshal([]byte(interactiveRequest.Data),&userLoginModel)
			userHandle.UserLogin(userLoginModel)
		}
		case 102:{
			userHandle := handle.NewUserHandle(this.Conn,this.TcpTool)

			userRegisterModel := user.UserRegisterModel{}
			json.Unmarshal([]byte(interactiveRequest.Data),&userRegisterModel)
			userHandle.UserRegister(userRegisterModel)
		}
		case 200:{
			messageHandle := handle.NewMessageHandle()
			
			groupSendMessage := message.GroupSendMessage{}
			json.Unmarshal([]byte(interactiveRequest.Data),&groupSendMessage)
			messageHandle.GroupSendMessage(groupSendMessage)
		}
	}
	return nil
}