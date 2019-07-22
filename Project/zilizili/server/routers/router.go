// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/video",
	// 		beego.NSInclude(
	// 			&controllers.VideoController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
	beego.Router("api/video/AddVideo", &controllers.VideoController{},"post:AddVideo")
	beego.Router("api/video/UpdateVideo", &controllers.VideoController{},"put:UpdateVideo")
	beego.Router("api/video/GetVideo", &controllers.VideoController{},"get:GetVideo")
	beego.Router("api/video/GetVideoList", &controllers.VideoController{},"get:GetVideoList")
	beego.Router("api/video/DeleteVideo", &controllers.VideoController{},"delete:DeleteVideo")
}