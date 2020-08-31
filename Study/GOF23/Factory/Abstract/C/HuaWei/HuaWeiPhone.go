package HuaWei


import (
	_"Factory/Abstract/I/Phone"
	"fmt"
)


type HuaWeiPhone struct{

}

/*
打开手机
*/
func (this *HuaWeiPhone) Open() {
	fmt.Println("打开华为手机")
}

/*
关闭
*/
func(this *HuaWeiPhone) Close()  {
	fmt.Println("关闭华为手机")
}

/*
打电话
*/
func(this *HuaWeiPhone) Call()  {
	fmt.Println("打华为电话")
}

/*
发送短信
*/
func(this *HuaWeiPhone) SendSMS()  {
	fmt.Println("发送华为短信")
}

