package main

import RabbitMQ "mysql/rabbitmq/work/rabbit"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQWork("kevin")
	rabbitmq.ConsumeWork()
}
