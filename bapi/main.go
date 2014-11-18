package main

import (
	_ "github.com/duoping.github.com/bapi/docs"
	_ "github.com/duoping.github.com/bapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
