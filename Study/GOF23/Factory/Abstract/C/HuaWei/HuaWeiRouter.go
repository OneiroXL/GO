package HuaWei


import (
	_"Factory/Abstract/I/Router"
	"fmt"
)


type HuaWeiRouter struct{

}

/*
打开
*/
func (this *HuaWeiRouter) Open() {
	fmt.Println("打开华为路由器")
}

/*
关闭
*/
func(this *HuaWeiRouter) Close()  {
	fmt.Println("关闭华为路由器")
}

/*
设置
*/
func (this *HuaWeiRouter) Set() {
	fmt.Println("设置华为路由器")
}