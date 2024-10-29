package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Product struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// 模拟从数据库查询的数据
var allproduct []*Product = []*Product{
	{1, "苹果"},
	{2, "华为"},
	{3, "OPPO"},
}

var (
	// 生成的html保存目录
	htmlOutPath = "./template"
	// 静态文件模板目录
	templatePath = "./template/"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(c *gin.Context) {
		GetGenerateHtml()
		c.HTML(http.StatusOK, "index.html", gin.H{"allproduct": allproduct})
	})
	r.GET("index2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "html_index.html", gin.H{})
	})
	r.Run()
}

// 生成静态文件的方法
func GetGenerateHtml() {
	// 获取模板
	contentTmp, err := template.ParseFiles(filepath.Join(templatePath, "index.html"))
	if err != nil {
		fmt.Println(err)
	}
	// 获取html生成路径
	fileName := filepath.Join(htmlOutPath, "html_index.html")
	// 生成静态文件
	generateStaticHtml(contentTmp, fileName, gin.H{"allproduct": allproduct})

}

// 生成静态文件
func generateStaticHtml(template *template.Template, fileName string, product map[string]interface{}) {
	// 判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
		}
	}
	// 生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	template.Execute(file, product)
}

func exist(filaName string) bool {
	_, err := os.Stat(filaName)
	return err == nil || os.IsExist(err)
}
