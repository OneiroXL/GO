package center

import (
	"fmt"
	"net"
	"chat/model/user"
	"chat/client/handle"
	"chat/common/tools"
)

type ControllerCenter struct {
	Conn net.Conn
	TcpTool tools.TcpTool
	UserID int
}


func (this *ControllerCenter) Login(){
	fmt.Println("----------登录----------")

	var userID int
	var passcode string

	fmt.Printf("请输入用户ID:")
	fmt.Scanf("%d\n",&userID)
	fmt.Printf("请输入密码:")
	fmt.Scanf("%s\n",&passcode)

	userLoginModel := user.UserLoginModel{}
	userLoginModel.UserID = userID
	userLoginModel.Passcode = passcode

	userHandle := handle.UserHandle{
		Conn:this.Conn,
		TcpTool:this.TcpTool,
	}

	err := userHandle.UserLogin(userLoginModel)

	if(err != nil){
		
	}
	this.UserID = userID
	go Process(this.Conn,userID)
	this.MainMenu()
}


func (this *ControllerCenter) MainMenu() {
	for {
		fmt.Println("----------欢迎来到聊天系统----------")
		fmt.Println("1:消息管理")
		fmt.Println("0:退出")
		
		fmt.Printf("请选择:")
		var v string
		fmt.Scanf("%s\n",&v)

		switch(v){
			case "1":{
				this.MessageMenu()
			}
			case "0":{
				fmt.Println("Bye!")
				return
			}
			default:{
				fmt.Println("输入正确的参数")
			}
		}
	}
}


func (this *ControllerCenter) MessageMenu(){
	for {
		fmt.Println("----------消息管理----------")
		fmt.Println("1:发送消息")
		fmt.Println("0:退出")
		
		fmt.Printf("请选择:")
		var v string
		fmt.Scanf("%s\n",&v)

		switch(v){
			case "1":{
				fmt.Printf("请输入您要发送的消息:")

				var msg string
				fmt.Scanf("%s\n",&msg)

				messageHandle := handle.NewMessageHandle(this.Conn,this.UserID)
				messageHandle.GroupSendMessage(msg)
			}
			case "0":{
				return
			}
			default:{
				fmt.Println("输入正确的参数")
			}
		}
	}
}
