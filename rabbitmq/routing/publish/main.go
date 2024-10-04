package main

import (
	"fmt"
	RabbitMQ "mysql/rabbitmq/routing/rabbit"
	"strconv"
	"time"
)

func main() {
	kevinone := RabbitMQ.NewRabbitMQRouting("kevin", "kevin_one")
	kevintwo := RabbitMQ.NewRabbitMQRouting("kevin", "kevin_two")
	for i := 0; i <= 100; i++ {
		kevinone.PublishRouting("hello one" + strconv.Itoa(i))
		kevintwo.PublishRouting("hello two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
