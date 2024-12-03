package main

import (
	"beego/controllers"
	_ "beego/routers"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/redis"
)

func main() {
	web.ErrorController(&controllers.ErrorController{})
	web.Run()
}
