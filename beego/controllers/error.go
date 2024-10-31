package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// ErrorController operations for Error
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

// URLMapping ...
func (c *ErrorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Error
// @Param	body		body 	models.Error	true		"body for Error content"
// @Success 201 {object} models.Error
// @Failure 403 body is empty
// @router / [post]
func (c *ErrorController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Error by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Error
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ErrorController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Error
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Error
// @Failure 403
// @router / [get]
func (c *ErrorController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Error
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Error	true		"body for Error content"
// @Success 200 {object} models.Error
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ErrorController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Error
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ErrorController) Delete() {

}
