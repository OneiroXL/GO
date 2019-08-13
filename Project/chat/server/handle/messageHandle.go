package handle

import (
	"chat/model/message"
	"chat/model/base"
	"chat/common/tools"
	"encoding/json"
)

type MessageHandle struct {

}

func NewMessageHandle() *MessageHandle {
	messageHandle := &MessageHandle{}
	return messageHandle
}


/*
群发消息
*/
func (this *MessageHandle) GroupSendMessage(groupSendMessage message.GroupSendMessage) {
	interactive := base.Interactive{}

	interactive.Status = 10000
	interactive.Type = 201
	interactive.Message = "成功"

	groupSendMessageJson,_ := json.Marshal(groupSendMessage)
	interactive.Data = string(groupSendMessageJson)
	interactiveJson,_ := json.Marshal(interactive)
	//发送消息
	tcpTool := tools.TcpTool{}
	for _, val := range userMap {
		tcpTool.Conn = val.Conn
		tcpTool.Write(interactiveJson)
	}
}