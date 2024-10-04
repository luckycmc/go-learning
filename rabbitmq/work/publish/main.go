package main

import (
	"fmt"
	RabbitMQ "mysql/rabbitmq/work/rabbit"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQWork("kevin")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishWork("Hello World" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
