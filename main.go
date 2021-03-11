package main

import (
	_ "EasyTutor/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/v1/easy-tutor/swagger"] = "swagger"
	}
	beego.Run()
}
