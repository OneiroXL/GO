package handle

import (
	"net"
	"encoding/json"
	"chat/model/base"
	"chat/model/message"
	"chat/common/tools"
	"fmt"
)

type MessageHandle struct{
	Conn net.Conn
	UserCode string
	TcpTool tools.TcpTool
}


func NewMessageHandle(conn net.Conn,userCode string) *MessageHandle {
	messageHandle := &MessageHandle{
		Conn:conn,
		UserCode:userCode,
		TcpTool:tools.TcpTool{
			Conn:conn,
		},
	}
	return messageHandle
}

/*
群发消息
*/
func (this *MessageHandle) GroupSendMessage(msg string){
	//消息体
	groupSendMessage := message.GroupSendMessage{}
	groupSendMessage.SendUserCode = this.UserCode
	groupSendMessage.Message = msg
	groupSendMessageJSON,_ := json.Marshal(groupSendMessage)

	//交互消息体
	interactiveRequest := base.InteractiveRequest{}
	interactiveRequest.Type = 200
	interactiveRequest.Data = string(groupSendMessageJSON)


	interactiveRequestJSON,_ := json.Marshal(interactiveRequest)

	this.TcpTool.Write(interactiveRequestJSON)
}

/*
展示群发消息
*/
func (this *MessageHandle) ShowGroupSendMessage(groupSendMessage message.GroupSendMessage) {
	if(groupSendMessage.SendUserCode != this.UserCode){
		fmt.Printf("用户[%v]:%s\n",groupSendMessage.SendUserCode,groupSendMessage.Message)
	}	
}