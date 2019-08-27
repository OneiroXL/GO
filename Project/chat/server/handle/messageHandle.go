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
	interactiveResponse := base.InteractiveResponse{}

	interactiveResponse.Status = 10000
	interactiveResponse.Type = 201
	interactiveResponse.Message = "成功"

	groupSendMessageJson,_ := json.Marshal(groupSendMessage)
	interactiveResponse.Data = string(groupSendMessageJson)
	interactiveResponseJson,_ := json.Marshal(interactiveResponse)

	//发送消息
	tcpTool := tools.TcpTool{}
	for _, val := range userMap {
		tcpTool.Conn = val.Conn
		tcpTool.Write(interactiveResponseJson)
	}
}