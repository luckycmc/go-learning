package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.72.130:6379", redis.DialPassword("secret_redis"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	_, err = c.Do("HSet", "names", "kevin", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("HGet", "names", "kevin"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
