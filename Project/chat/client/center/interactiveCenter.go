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
	UserCode string
}


func Process(conn net.Conn,userCode string){
	tcpTool := tools.TcpTool{
		Conn:conn,
	}
	interactiveCenter := InteractiveCenter{
		Conn:conn,
		TcpTool:tcpTool,
		UserCode:userCode,
	}
	for {
		data,err := tcpTool.Read()
		if(err != nil){
			fmt.Println("服务器错误")
			return
		}
		interactiveResponse := base.InteractiveResponse{}
		json.Unmarshal([]byte(data),&interactiveResponse)
		interactiveCenter.InteractiveHandle(interactiveResponse)
	}
}

/*
处理处理
*/
func (this *InteractiveCenter) InteractiveHandle(interactiveResponse base.InteractiveResponse) error {
	switch interactiveResponse.Type {
		case 101:{
			
		}
		case 201:{
			groupSendMessage := message.GroupSendMessage{}
			json.Unmarshal([]byte(interactiveResponse.Data),&groupSendMessage)

			messageHandle := handle.NewMessageHandle(this.Conn,this.UserCode)
			messageHandle.ShowGroupSendMessage(groupSendMessage)
		}
	}
	return nil
}