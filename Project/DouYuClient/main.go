package main

import (
	"DouYuClient/common/tools"
	"DouYuClient/process"
	"fmt"
	"net"
)

var (
	transferTool tools.TransferTool
)

func main() {
	conn, errConn := net.Dial("tcp", "openbarrage.douyutv.com:8601")

	if errConn != nil {
		fmt.Println("服务器连接出错:", errConn)
	}
	barrageProcess := process.NewBarrageProcess(conn)

	var roomId string
	fmt.Print("请输入房间号:")
	fmt.Scanf("%s/n", &roomId)
	barrageProcess.Start(roomId)
}
