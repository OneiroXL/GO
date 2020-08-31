package XiaoMi


import (
	_"Factory/Abstract/I/Phone"
	"fmt"
)

type XiaoMiPhone struct {

}


/*
打开手机
*/
func (this *XiaoMiPhone) Open() {
	fmt.Println("打开小米手机")
}

/*
关闭
*/
func(this *XiaoMiPhone) Close()  {
	fmt.Println("关闭小米手机")
}

/*
打电话
*/
func(this *XiaoMiPhone) Call()  {
	fmt.Println("打小米电话")
}

/*
发送短信
*/
func(this *XiaoMiPhone) SendSMS()  {
	fmt.Println("发送小米短信")
}

