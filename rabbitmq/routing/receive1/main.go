package main

import RabbitMQ "mysql/rabbitmq/routing/rabbit"

func main() {
	kevinone := RabbitMQ.NewRabbitMQRouting("kevin", "kevin_one")
	kevinone.ReceiveRouting()
}
