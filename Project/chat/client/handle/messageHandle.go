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
	UserID int
	TcpTool tools.TcpTool
}


func NewMessageHandle(conn net.Conn,userID int) *MessageHandle {
	messageHandle := &MessageHandle{
		Conn:conn,
		UserID:userID,
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
	groupSendMessage.SendUserID = this.UserID
	groupSendMessage.Message = msg
	groupSendMessageJSON,_ := json.Marshal(groupSendMessage)

	//交互消息体
	interactive := base.Interactive{}
	interactive.Type = 200
	interactive.Status = 10000
	interactive.Message = "群发消息"
	interactive.Data = string(groupSendMessageJSON)


	interactiveJSON,_ := json.Marshal(interactive)

	this.TcpTool.Write(interactiveJSON)
}

/*
展示群发消息
*/
func (this *MessageHandle) ShowGroupSendMessage(groupSendMessage message.GroupSendMessage) {
	if(groupSendMessage.SendUserID != this.UserID){
		fmt.Printf("用户[%v]:%s\n",groupSendMessage.SendUserID,groupSendMessage.Message)
	}	
}