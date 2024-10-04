package main

import RabbitMQ "mysql/rabbitmq/routing/rabbit"

func main() {
	kevintwo := RabbitMQ.NewRabbitMQRouting("kevin", "kevin_two")
	kevintwo.ReceiveRouting()
}
