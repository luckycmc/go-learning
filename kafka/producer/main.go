package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_0_0

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_topic"
	msg.Value = sarama.StringEncoder("Hello World")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.72.130:9092"}, config)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	defer client.Close()
	// 发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("Error sending messages:", err)
		return
	}
	fmt.Printf("Partition:%d Offset:%d\n", partition, offset)
}
