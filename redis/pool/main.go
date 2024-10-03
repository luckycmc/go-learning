package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func init() {
	// 实例化连接池
	pool = &redis.Pool{
		MaxIdle:     16,                // 最初连接数量
		MaxActive:   16,                // 最大连接数量
		IdleTimeout: 240 * time.Second, // 连接关闭时间
		Dial: func() (redis.Conn, error) { // 要连接的redis
			return redis.Dial("tcp", "192.168.72.130:6379", redis.DialPassword("secret_redis"))
		},
	}
}

func main() {
	// 从连接池获取一个连接
	c := pool.Get()
	// 函数运行结束后把连接放回
	defer c.Close()
	_, err := c.Do("Set", "name", "kevin")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
	// 关闭连接池
	pool.Close()
}
