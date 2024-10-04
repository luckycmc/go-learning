package main

import RabbitMQ "mysql/topic/rabbit"

func main() {
	kevinone := RabbitMQ.NewRabbitMQTopic("kevinTopic", "#")
	kevinone.ReceiveTopic()
}
