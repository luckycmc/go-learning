package main

import RabbitMQ "mysql/topic/rabbit"

func main() {
	kevintwo := RabbitMQ.NewRabbitMQTopic("kevinTopic", "kevin.*.two")
	kevintwo.ReceiveTopic()
}
