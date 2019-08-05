package center

import (
	"fmt"
	"net"
)

type ControllerCenter struct {
	Conn net.Conn
}

func (this *ControllerCenter) MainMenu() {
	for {
		fmt.Println("----------欢迎来到聊天系统----------")
		fmt.Println("1:发送消息")
		fmt.Println("输入'exit'退出")
		
		fmt.Printf("请选择:")
		var v string
		fmt.Scanf("%s",&v)

		switch(v){
			case "1":{
				fmt.Println("")
			}
			case "exit":{
				fmt.Println("Bye!")
				return
			}
			default:{
				fmt.Println("输入正确的参数")
			}
		}
	}
}