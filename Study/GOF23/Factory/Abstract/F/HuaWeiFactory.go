package F

import (
	"Factory/Abstract/I/Phone"
	"Factory/Abstract/I/Router"
	"Factory/Abstract/C/HuaWei"
	_"Factory/Abstract/I"
)

type HuaWeiFactory struct {

}


/*
制作手机
*/
func (this *HuaWeiFactory) MakePhone() Phone.IPhoneProduct {
	var huaWeiPhone Phone.IPhoneProduct = new(HuaWei.HuaWeiPhone)
	return huaWeiPhone
}

/*
制作路由器
*/
func (this *HuaWeiFactory) MakeRouter() Router.IRouterProduct {
	var huaWeiRouter Router.IRouterProduct = new(HuaWei.HuaWeiRouter)
	return huaWeiRouter
}
 	