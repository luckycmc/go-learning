package controllers

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
	_ "github.com/beego/beego/v2/server/web/session/mysql"
	"log"
	"net/http"
)

var globalSessions1 *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "sessionId",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "root:root@tcp(192.168.72.130:3306)/beego?charset=utf8",
	}
	var err error
	globalSessions1, err = session.NewManager("mysql", sessionConfig)
	if err != nil {
		log.Fatalf("Failed to initialize session manager: %v", err)
	}
	log.Println("Session manager initialized successfully")
	go globalSessions1.GC()
}

// Mysql_sessionController operations for Mysql_session
type Mysql_sessionController struct {
	beego.Controller
}

// URLMapping ...
func (c *Mysql_sessionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Get", c.Get)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Mysql_session
// @Param	body		body 	models.Mysql_session	true		"body for Mysql_session content"
// @Success 201 {object} models.Mysql_session
// @Failure 403 body is empty
// @router / [post]
func (c *Mysql_sessionController) Post() {

}

func (c *Mysql_sessionController) Get() {
	logger := logs.GetLogger()
	ctx := context.Background()
	r, _ := http.NewRequest("GET", "/", nil)
	w := http.ResponseWriter(c.Ctx.ResponseWriter)
	sess, err := globalSessions.SessionStart(w, r)
	if err != nil {
		logger.Println("abc")
	}
	defer sess.SessionRelease(ctx, w)
	sess.Set(ctx, "name", "kevin112")

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

// GetOne ...
// @Title GetOne
// @Description get Mysql_session by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Mysql_session
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Mysql_sessionController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Mysql_session
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Mysql_session
// @Failure 403
// @router / [get]
func (c *Mysql_sessionController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Mysql_session
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Mysql_session	true		"body for Mysql_session content"
// @Success 200 {object} models.Mysql_session
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Mysql_sessionController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Mysql_session
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Mysql_sessionController) Delete() {

}
