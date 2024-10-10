package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

type User struct {
	Id    uint
	Name  string
	Email string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(192.168.72.130:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&User{})

	// 设置连接池
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)

	var startTime = time.Now()
	wg.Add(1)
	// 创建记录
	go func() {
		for i := 1; i <= 100; i++ {
			user := User{
				uint(i),
				"test_user" + strconv.Itoa(i),
				"test_user" + strconv.Itoa(i) + "@email.com",
			}
			db.Create(&user)
		}
		wg.Done()
	}()
	wg.Wait()

	var endTime = time.Now()
	fmt.Println("create_time: ", endTime.Sub(startTime))

	// 查询
	var user = new(User)
	db.First(user)
	fmt.Printf("user:%#v\n", user)

	var uu User
	db.Find(&uu, "email = ?", "test_user2@email.com")
	fmt.Printf("user:%#v\n", uu)

	// 更新
	db.Model(&user).Update("email", "test11@email.com")

	// 删除
	db.Delete(&uu)
}
