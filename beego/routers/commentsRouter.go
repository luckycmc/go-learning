package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["beego/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["beego/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["beego/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["beego/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["beego/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/article`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["beego/controllers:ArticleController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/article`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/user`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/user/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/user/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/user/:id:int`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "AddUser",
			Router:           `/user/add_user`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego/controllers:UserController"] = append(beego.GlobalControllerRouter["beego/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/user/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
