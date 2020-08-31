package XiaoMi


import (
	_"Factory/Abstract/I/Router"
	"fmt"
)


type XiaoMiRouter struct{

}

/*
打开
*/
func (this *XiaoMiRouter) Open() {
	fmt.Println("打开小米路由器")
}

/*
关闭
*/
func(this *XiaoMiRouter) Close()  {
	fmt.Println("关闭小米路由器")
}

/*
设置
*/
func (this *XiaoMiRouter) Set() {
	fmt.Println("设置小米路由器")
}