package main

import (
	"Factory/Abstract/F"
	"Factory/Abstract/I/Phone"
	"Factory/Abstract/I/Router"
	"Factory/Abstract/I"
)


func main(){
	var huaWeiFactory I.IFactory = new(F.HuaWeiFactory)


	var huaweiPhone Phone.IPhoneProduct = huaWeiFactory.MakePhone()
	var huaweiRouter Router.IRouterProduct = huaWeiFactory.MakeRouter()

	huaweiPhone.Open()
	huaweiPhone.Call()
	huaweiPhone.Close()

	huaweiRouter.Open()
	huaweiRouter.Set()
	huaweiRouter.Close()


	var xiaoMiFactory I.IFactory = new(F.XiaoMiFactory)

	var xiaomiPhone Phone.IPhoneProduct = xiaoMiFactory.MakePhone()
	var xiaomiRouter Router.IRouterProduct = xiaoMiFactory.MakeRouter()

	xiaomiPhone.Open()
	xiaomiPhone.Call()
	xiaomiPhone.Close()

	xiaomiRouter.Open()
	xiaomiRouter.Set()
	xiaomiRouter.Close()

}