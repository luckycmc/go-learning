package controllers

import (
	"beego/models"
	beego "github.com/beego/beego/v2/server/web"
	"html/template"
	"strconv"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

type user struct {
	Id   int64  `form:"-"`
	Name string `form:"name"`
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Create", c.Create)
}

// Create ...
// @router /user/create [get]
func (c *UserController) Create() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "user/create.tpl"
}

// AddUser ...
// @router /user/add_user [post]
func (c *UserController) AddUser() {
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		return
	}
	// name := c.GetString("name")
	user := models.User{
		Name: u.Name,
	}
	id, err := models.AddUser(&user)
	if err != nil {
		return
	}
	c.Ctx.WriteString("success insert: " + strconv.FormatInt(id, 10))
}

// Post ...
// @Title Create
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {object} models.User
// @Failure 403 body is empty
// @router /user [post]
func (c *UserController) Post() {

}

// Get ...
// @router /user [get]
func (c *UserController) Get() {

	c.TplName = "user/index.tpl"
}

// GetOne ...
// @Title GetOne
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /user/:id:int [get]
func (c *UserController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id
}

// GetAll ...
// @Title GetAll
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router /user [get]
func (c *UserController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /user/:id [put]
func (c *UserController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /user/:id [delete]
func (c *UserController) Delete() {

}
