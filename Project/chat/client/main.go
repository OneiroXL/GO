package main

import (
	"net"
	"fmt"
	"chat/client/center"
	"chat/common/tools"
	"chat/model/base"
)

func main(){
	conn, errConn := net.Dial("tcp", "127.0.0.1:8888")

	if errConn != nil {
		fmt.Println("服务器连接出错:", errConn)
	}

	controllerCenter := center.ControllerCenter{
		Conn:conn,
	}

	controllerCenter.MainMenu()
}


func Process(conn net.Conn){
	tcpTool := tools.TcpTool{
		Conn:conn,
	}
	interactiveCenter := center.InteractiveCenter{
		Conn:conn,
		TcpTool:tcpTool,
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