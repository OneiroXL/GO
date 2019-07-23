package server

import (
	_"zilizili/server/routers"
	"zilizili/model"
	"github.com/astaxie/beego"
)

func Start() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	model.Database(beego.AppConfig.String("sqlconn"))
	beego.Run()
}
