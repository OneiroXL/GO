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
	UserCode string
}

/*
首菜单
*/
func (this *ControllerCenter) Index(){
	for {
		fmt.Println("----------登录注册----------")
		fmt.Println("1:登录")
		fmt.Println("2:注册")
		fmt.Println("0:退出")

		fmt.Printf("请选择:")
		var v string
		fmt.Scanf("%s\n",&v)

		switch(v){
			case "1":{
				this.Login()
			}
			case "2":{
				this.UserRegister()
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

func (this *ControllerCenter) UserRegister(){
	fmt.Println("----------注册----------")

	var userCode string
	var passcode string
	var userName string

	fmt.Printf("请输入用户注册ID:")
	fmt.Scanf("%s\n",&userCode)
	fmt.Printf("请输入密码:")
	fmt.Scanf("%s\n",&passcode)
	fmt.Printf("请输入用户姓名:")
	fmt.Scanf("%s\n",&userName)

	userRegisterModel := user.UserRegisterModel{}

	userRegisterModel.UserCode = userCode
	userRegisterModel.Passcode = passcode
	userRegisterModel.UserName = userName

	userHandle := handle.UserHandle{
		Conn:this.Conn,
		TcpTool:this.TcpTool,
	}

	mes,err := userHandle.UserRegister(userRegisterModel)

	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(mes)
}

func (this *ControllerCenter) Login(){
	fmt.Println("----------登录----------")

	var userCode string
	var passcode string

	fmt.Printf("请输入用户ID:")
	fmt.Scanf("%s\n",&userCode)
	fmt.Printf("请输入密码:")
	fmt.Scanf("%s\n",&passcode)

	userLoginModel := user.UserLoginModel{}
	userLoginModel.UserCode = userCode
	userLoginModel.Passcode = passcode

	userHandle := handle.UserHandle{
		Conn:this.Conn,
		TcpTool:this.TcpTool,
	}

	err := userHandle.UserLogin(userLoginModel)

	if(err != nil){
		
	}
	this.UserCode = userCode
	go Process(this.Conn,userCode)
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

				messageHandle := handle.NewMessageHandle(this.Conn,this.UserCode)
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
