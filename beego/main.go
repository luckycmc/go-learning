package main

import (
	"beego/controllers"
	_ "beego/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.ErrorController(&controllers.ErrorController{})
	web.Run()
}
