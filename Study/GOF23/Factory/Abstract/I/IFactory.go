package I

import (
	"Factory/Abstract/I/Phone"
	"Factory/Abstract/I/Router"
)

type IFactory interface{
	MakePhone() Phone.IPhoneProduct
 	MakeRouter() Router.IRouterProduct
}


