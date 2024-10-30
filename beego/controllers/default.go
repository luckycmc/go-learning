package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "LearnBeego"
	c.Data["Email"] = "kevinqq1314@gmail.com"
	c.TplName = "index.tpl"
}
