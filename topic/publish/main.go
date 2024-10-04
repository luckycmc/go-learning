package main

import (
	"fmt"
	RabbitMQ "mysql/topic/rabbit"
	"strconv"
	"time"
)

func main() {
	kevinone := RabbitMQ.NewRabbitMQTopic("kevinTopic", "kevin.topic.one")
	kevintwo := RabbitMQ.NewRabbitMQTopic("kevinTopic", "kevin.topic.two")
	for i := 0; i < 100; i++ {
		kevinone.PublishTopic("hello kevin topic one" + strconv.Itoa(i))
		kevintwo.PublishTopic("hello kevin topic two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
