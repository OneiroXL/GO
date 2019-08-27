package main

import (
	"fmt"
	"net"
	"chat/common/tools"
	"chat/model/base"
	"encoding/json"
	"chat/server/center"
	"chat/server/dao"
)

func main(){
	lister, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil{
		fmt.Println("监听失败...err: ",err)
		return
	}
	defer lister.Close()//延时关闭listen

	for {
		conn, err := lister.Accept()
		if err != nil{
			fmt.Println("连接Accept() 失败,err: ",err)
		}else{
			fmt.Printf("Accept() suc conn=%v,客户端IP=%v\n",conn,conn.RemoteAddr().String())
		}
		//初始化数据库
		DBBase.Database("root:123456@tcp(127.0.0.1:3306)/chat?charset=utf8&parseTime=True&loc=Local")
		go Process(conn)
	}
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
			fmt.Println("数据读取出错，或连接已断开 Err:",err)
			
			conn.Close();
			return
		}
		interactiveRequest := base.InteractiveRequest{}
		json.Unmarshal([]byte(data),&interactiveRequest)
		interactiveCenter.InteractiveHandle(interactiveRequest)
	}

}