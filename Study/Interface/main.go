package main

import (
	"fmt"
)

type USB interface{
	Start()
	Stop()
}

type Phone struct{
	
}

func (this Phone) Start(){
	fmt.Println("手机开始工作")
}

func (this Phone) Stop(){
	fmt.Println("手机停止工作")
}

func (this Phone) Call(){
	fmt.Println("打电话....")
}


type Camera struct{

}

func (this Camera) Start(){
	fmt.Println("相机开始工作")
}

func (this Camera) Stop(){
	fmt.Println("相机停止工作")
}

type Computer struct{

}

func (this Computer) Working(usb USB){
	usb.Start()
	if p,flag := usb.(Phone);flag{ //类型断言
		p.Call()
	}
	usb.Stop()
}




func main()  {
	phone := Phone{}
	camera := Camera{}

	computer := Computer{}
	computer.Working(phone)
	computer.Working(camera)
}