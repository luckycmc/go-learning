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
	_, err = c.Do("MSet", "name", "kevin", "email", "test@email.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Strings(c.Do("MGet", "name", "email"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range r {
		fmt.Println(v)
	}
}
