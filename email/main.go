package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"
)

func main() {
	ch := make(chan *email.Email, 10)
	p, err := email.NewPool("smtp.163.com:25",
		4,
		smtp.PlainAuth(
			"",
			"chenmc2015@163.com",
			"JCVAJFRk7PvSi26A",
			"smtp.163.com",
		))
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := p.Send(e, 10*time.Second)
				if err != nil {
					fmt.Fprintf(os.Stderr, "email:%v failed to send email:%v\n", e, err)
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		e := email.NewEmail()
		// 设置发送方邮箱
		e.From = "测试邮件<chenmc2015@163.com>"
		// 设置接收方邮箱
		e.To = []string{"2287232750@qq.com"}
		// 设置抄送
		e.Cc = []string{"kevinchenqq1314@gmail.com"}
		// 设置秘密抄送
		e.Bcc = []string{"kevinchenqq1314@gmail.com"}
		// 设置主题
		e.Subject = "你好"
		// 设置内容
		e.Text = []byte(fmt.Sprintf("每天都要开心呀%d", i+1))
		// 发送html
		// e.HTML = []byte(`
		// <h1>幸运每一天</h1>
		// `)
		// 发送附件
		e.AttachFile("./test.txt")
		ch <- e
	}
	close(ch)
	wg.Wait()
}
