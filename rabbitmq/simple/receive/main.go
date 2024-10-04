package main

import RabbitMQ "mysql/rabbitmq/simple/rabbit"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "kevin")
	rabbitmq.ConsumeSimple()
}
