package controllers

import (
	"context"
	"github.com/astaxie/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	"log"
	"net/http"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "sessionId",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "192.168.72.130:6379,100,secret_redis",
	}
	var err error
	globalSessions, err = session.NewManager("redis", sessionConfig)
	if err != nil {
		log.Fatalf("Failed to initialize session manager: %v", err)
	}
	log.Println("Session manager initialized successfully")
	go globalSessions.GC()
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
}

// IndexController operations for Index
type IndexController struct {
	beego.Controller
}

// URLMapping ...
func (c *IndexController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Index
// @Param	body		body 	models.Index	true		"body for Index content"
// @Success 201 {object} models.Index
// @Failure 403 body is empty
// @router / [post]
func (c *IndexController) Post() {

}

func (c *IndexController) Get() {
	logger := logs.GetLogger()
	ctx := context.Background()
	r, _ := http.NewRequest("GET", "/", nil)
	w := http.ResponseWriter(c.Ctx.ResponseWriter)
	logs.GetLogger("ORM").Println("hello")
	sess, err := globalSessions.SessionStart(w, r)
	if err != nil {
		logger.Println("abc")
	}
	defer sess.SessionRelease(ctx, w)
	sess.Set(ctx, "name", "kevin113")

	if err != nil {
		logger.Println("abcd")
	}

	c.Data["name"] = sess.Get(ctx, "name")
	if c.Data["name"] == nil {
		logger.Println("abcde")
	}

	/*v := c.GetSession("name")
	if v == nil {
		c.SetSession("name", "kevin")
		c.Data["name"] = "kevin"
	} else {
		c.SetSession("name", "kevin1")
		c.Data["name"] = v
	}*/
	c.TplName = "index/index.tpl"
}

func (c *IndexController) Index() {
	// c.Abort("401")
	c.Redirect("/", 302)
}

// GetOne ...
// @Title GetOne
// @Description get Index by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Index
// @Failure 403 :id is empty
// @router /:id [get]
func (c *IndexController) GetOne() {
	name := c.Ctx.Input.Param(":name")
	c.Data["name"] = name
	c.TplName = "index/index.tpl"
}

// GetAll ...
// @Title GetAll
// @Description get Index
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Index
// @Failure 403
// @router / [get]
func (c *IndexController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Index
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Index	true		"body for Index content"
// @Success 200 {object} models.Index
// @Failure 403 :id is not int
// @router /:id [put]
func (c *IndexController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Index
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *IndexController) Delete() {

}
