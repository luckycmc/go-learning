package main

import (
	"beego/controllers"
	_ "beego/routers"
	_ "github.com/astaxie/beego/session/redis"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
	// 设置session驱动
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.72.130:6379"
}
