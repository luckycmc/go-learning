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

	beego.Router("/mysql_session", &controllers.Mysql_sessionController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "get:Login")

	// namespace
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/test", &controllers.ArticleController{}, "get:GetOne"),
		beego.NSNamespace("/article",
			beego.NSRouter("/all", &controllers.ArticleController{}, "get:GetAll"),
			beego.NSRouter("/", &controllers.ArticleController{}, "get:Get;post:Post"),
			beego.NSRouter("/create", &controllers.ArticleController{}, "get:Create"),
			beego.NSRouter("/:id", &controllers.ArticleController{}, "get:GetOne"),
			beego.NSRouter("/update/:id", &controllers.ArticleController{}, "post:Put;get:Edit"),
			beego.NSRouter("/delete/:id", &controllers.ArticleController{}, "get:Delete"),
			beego.NSRouter("/:title", &controllers.ArticleController{}, "get:GetByTitle"),
		))
	beego.AddNamespace(ns)
}
