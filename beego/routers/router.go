package routers

import (
	"beego/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	// 基本路由
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Welcome to beego"))
	})

	// 自定义handler
}
