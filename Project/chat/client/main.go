package main

import (
	"net"
	"fmt"
	"chat/client/center"
	"chat/common/tools"
)

func main(){
	conn, errConn := net.Dial("tcp", "127.0.0.1:8888")

	if errConn != nil {
		fmt.Println("服务器连接出错:", errConn)
	}
	tcpTool := tools.TcpTool{
		Conn:conn,
	}
	controllerCenter := center.ControllerCenter{
		Conn:conn,
		TcpTool:tcpTool,
	}

	controllerCenter.Index()
}
