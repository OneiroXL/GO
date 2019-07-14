package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
	"../common/tools"
)

func main()  {
	conn,errConn := net.Dial("tcp","openbarrage.douyutv.com:8601")

	if(errConn != nil){
		fmt.Println("服务区链接出错",errConn)
	}
	transferTool :=  tools.TransferTool{
		Conn:conn,
	}
	var login string = "type@=loginreq/roomid@=30191/\\0"
	loginErr := transferTool.Write([]byte(login))
	if(loginErr != nil){
		fmt.Println("发送数据出错：",loginErr)
	}
	var joingroup string = "type@=joingroup/rid@=30191/gid@=-9999/\\0"
	joingroupErr := transferTool.Write([]byte(joingroup))

	if(joingroupErr != nil){
		fmt.Println("发送数据出错:",joingroupErr)
	}
	KeepLive(transferTool)

	for(true){
		data,err := transferTool.Read();
		if(err != nil){
			fmt.Println("接受数据出错",err)
			break
		}
		fmt.Println("接收到的数据",data)
	}

	if(conn != nil){
		
	}
}

func KeepLive(transferTool tools.TransferTool){
	for {
		var timestamp string = strconv.Itoa(int(time.Now().Unix()))
		var keepliveInfo string = "type@=keeplive/tick@=" + timestamp + "/\\0"
		// time.Sleep(time.Second * 45)
		err :=  transferTool.Write([]byte(keepliveInfo))
		if(err != nil){
			fmt.Println("心跳检测阵亡！！！")
			break
		}
		break
	}
}