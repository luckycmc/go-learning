package controllers

import (
	"beego/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"html/template"
)

// ArticleController operations for Article
type ArticleController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArticleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Get", c.Create)
}

// Create ...
// @router /article/create
func (c *ArticleController) Create() {
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "article/create.tpl"
}

// Post ...
// @Title Create
// @Description create Article
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 201 {object} models.Article
// @Failure 403 body is empty
// @router /article [post]
func (c *ArticleController) Post() {
	title := c.GetString("title")
	content := c.GetString("content")
	c.Ctx.Output.Body([]byte(title))
	c.Ctx.Output.Body([]byte(content))
	article := models.Article{
		Title:   title,
		Content: content,
	}
	id, err := models.AddArticle(&article)
	if err != nil {
		logs.Error(err)
	}
	c.Data["id"] = id
}

// Get ...
// @router /article [get]
func (c *ArticleController) Get() {
	log := logs.GetLogger()
	id := c.Ctx.Input.Param("id")
	log.Println("id: ", id)
	c.TplName = "article/show.tpl"
}

// GetOne ...
// @Title GetOne
// @Description get Article by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :id is empty
// @router /article/:id [get]
func (c *ArticleController) GetOne() {
	log := logs.GetLogger()
	id, err := c.GetInt(":id")
	if err != nil {
		logs.Error(err)
	}
	log.Println("id: ", id)
	// select
	article, err := models.GetArticleById(int64(id))
	if err != nil {
		logs.Error(err)
	}
	c.Data["article"] = &article
	c.TplName = "article/show.tpl"
}

// GetAll ...
// @Title GetAll
// @Description get Article
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Article
// @Failure 403
// @router / [get]
func (c *ArticleController) GetAll() {
	c.TplName = "article/show.tpl"

}

// Put ...
// @Title Put
// @Description update the Article
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 200 {object} models.Article
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArticleController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Article
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArticleController) Delete() {

}
