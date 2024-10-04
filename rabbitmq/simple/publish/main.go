package main

import (
	"fmt"
	RabbitMQ "mysql/rabbitmq/simple/rabbit"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("kevin")
	rabbitmq.PublishSimple("Hello World")
	fmt.Println("发送成功")
}
