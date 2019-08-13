package center

import (
	"chat/model/base"
	"chat/model/message"
	"net"
	"chat/common/tools"
	"chat/client/handle"
	"encoding/json"
	"fmt"
)

/*
交互中心
*/
type InteractiveCenter struct{
	Conn net.Conn
	TcpTool tools.TcpTool
	UserID int
}


func Process(conn net.Conn,userID int){
	tcpTool := tools.TcpTool{
		Conn:conn,
	}
	interactiveCenter := InteractiveCenter{
		Conn:conn,
		TcpTool:tcpTool,
		UserID:userID,
	}
	for {
		data,err := tcpTool.Read()
		if(err != nil){
			fmt.Println("服务器错误")
			return
		}
		interactive := base.Interactive{}
		json.Unmarshal([]byte(data),&interactive)
		interactiveCenter.InteractiveHandle(interactive)
	}
}

/*
处理处理
*/
func (this *InteractiveCenter) InteractiveHandle(interactive base.Interactive) error {
	switch interactive.Type {
		case 101:{
			
		}
		case 201:{
			groupSendMessage := message.GroupSendMessage{}
			json.Unmarshal([]byte(interactive.Data),&groupSendMessage)

			messageHandle := handle.NewMessageHandle(this.Conn,this.UserID)
			messageHandle.ShowGroupSendMessage(groupSendMessage)
		}
	}
	return nil
}