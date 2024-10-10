package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建路由
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 绑定路由规则
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "kevin")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// 监听端口
	r.Run(":8080")
}
