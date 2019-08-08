package center

import (
	"chat/model/base"
	"net"
	"chat/common/tools"
	"fmt"
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
处理处理
*/
func (this *InteractiveCenter) InteractiveHandle(interactive base.Interactive) error {
	switch interactive.Type {
		case 100:{
			userHandle := handle.NewUserHandle(this.Conn,this.TcpTool)

			userLoginModel := user.UserLoginModel{}
			json.Unmarshal([]byte(interactive.Data),&userLoginModel)
			userHandle.UserLogin(userLoginModel)
		}
	}
	return nil
}

