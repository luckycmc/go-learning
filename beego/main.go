package main

import (
	"beego/controllers"
	_ "beego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
