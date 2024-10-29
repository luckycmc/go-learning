package main

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func main() {
	e := email.NewEmail()
	// 设置发送方邮箱
	e.From = "测试邮件<chenmc2015@163.com>"
	// 设置接收方邮箱
	e.To = []string{"2287232750@qq.com"}
	// 设置主题
	e.Subject = "你好"
	// 设置内容
	e.Text = []byte("每天都要开心呀")
	// 设置服务器相关配置
	err := e.Send(
		"smtp.163.com:25",
		smtp.PlainAuth(
			"",
			"chenmc2015@163.com",
			"JCVAJFRk7PvSi26A",
			"smtp.163.com",
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
