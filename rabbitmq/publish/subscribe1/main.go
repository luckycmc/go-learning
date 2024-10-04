package main

import RabbitMQ "mysql/rabbitmq/publish/rabbit"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitmq.ConsumeSub()
}
