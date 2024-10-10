package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建路由
	r := gin.Default()
	// 绑定路由规则
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/")
	// 监听端口，默认8080
	r.Run(":8080")
}
