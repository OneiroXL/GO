package F

import (
	"Factory/Abstract/I/Phone"
	"Factory/Abstract/I/Router"
	"Factory/Abstract/C/XiaoMi"
	_"Factory/Abstract/I"
)

type XiaoMiFactory struct {

}


/*
制作手机
*/
func (this *XiaoMiFactory) MakePhone() Phone.IPhoneProduct {
	var xiaoMiPhone Phone.IPhoneProduct  = new(XiaoMi.XiaoMiPhone)
	return xiaoMiPhone
}

/*
制作路由器
*/
func (this *XiaoMiFactory) MakeRouter() Router.IRouterProduct {
	var xiaoRouter Router.IRouterProduct  = new(XiaoMi.XiaoMiRouter)
	return xiaoRouter
}
 	