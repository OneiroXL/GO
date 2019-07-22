package main

import (
	_"./routers"
	"../model"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	model.Database(beego.AppConfig.String("sqlconn"))
	beego.Run()
}
