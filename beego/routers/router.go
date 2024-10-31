package routers

import (
	"beego/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/user", &controllers.UserController{})
	// 基本路由
	// beego.Get("/", func(ctx *context.Context) {
	// 	ctx.Output.Body([]byte("Welcome to beego"))
	// })
	// 注解路由
	beego.Include(&controllers.UserController{})

	// 自定义handler

	// 正则
	beego.Router("/index/:name:string", &controllers.IndexController{}, "get:GetOne")
	beego.Router("/index", &controllers.IndexController{}, "get:Get")
	beego.Router("/redirect", &controllers.IndexController{}, "get:Index")

	// namespace
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/test", &controllers.ArticleController{}, "get:GetOne"),
		beego.NSNamespace("/article",
			beego.NSRouter("/all", &controllers.ArticleController{}, "get:GetAll"),
			beego.NSRouter("/", &controllers.ArticleController{}, "get:Get;post:Post"),
			beego.NSRouter("/create", &controllers.ArticleController{}, "get:Create"),
		))
	beego.AddNamespace(ns)
}
