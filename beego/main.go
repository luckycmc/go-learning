package main

import (
	"beego/controllers"
	_ "beego/routers"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "192.168.72.130:6379,100,secret_redis",
	}
	globalSessions, _ = session.NewManager("redis", sessionConfig)
	go globalSessions.GC()
}

func main() {
	web.ErrorController(&controllers.ErrorController{})
	web.Run()
}
