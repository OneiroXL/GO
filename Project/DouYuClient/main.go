package main

import (
	"fmt"
	"net"
	"DouYuClient/common/tools"
	"DouYuClient/process"
)

var (
	transferTool tools.TransferTool
)

func main()  {
	conn,errConn := net.Dial("tcp","openbarrage.douyutv.com:8601")

	if(errConn != nil){
		fmt.Println("服务器连接出错:",errConn)
	}
	barrageProcess := process.NewBarrageProcess(conn)
	barrageProcess.Start("213116")
}
