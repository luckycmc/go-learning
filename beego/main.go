package main

import (
	"beego/controllers"
	_ "beego/routers"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
	_ "github.com/beego/beego/v2/server/web/session/mysql"
	"log"
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
		ProviderConfig:  "root:root@tcp(localhost:3306)/beego?charset=utf8",
	}
	var err error
	globalSessions, err = session.NewManager("mysql", sessionConfig)
	if err != nil {
		log.Fatalf("Failed to initialize session manager: %v", err)
	}
	log.Println("Session manager initialized successfully")
	go globalSessions.GC()
}

func main() {
	web.ErrorController(&controllers.ErrorController{})
	web.Run()
	// 设置session驱动
	// redis
	// beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "192.168.72.130:6379"

	// Mysql
	// beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "root:root@tcp(192.168.72.130:3306)/beego?charset=utf8"
}
