package server

import (
	_"zilizili/server/routers"
	"zilizili/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/context"
	"log"
)

var FilterUser = func(ctx *context.Context) {
	token := ctx.Input.Header("Token")
	log.Println("token:",token)
    if token != "ZXL" {
        ctx.WriteString("请先登录")
    }
}

func Start() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowAllOrigins:  true,
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)

	model.Database(beego.AppConfig.String("sqlconn"))
	beego.Run()
}
