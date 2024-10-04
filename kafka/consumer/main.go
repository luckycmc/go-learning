package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"sync"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.72.129:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer:", err)
		return
	}
	partitionList, err := consumer.Partitions("test_topic")
	if err != nil {
		fmt.Println("Failed to get list of partitions:", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test_topic", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		wg.Add(1)
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s \n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			wg.Done()
		}(pc)
		wg.Wait()
	}
}
