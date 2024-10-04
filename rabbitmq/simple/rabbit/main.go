package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var wg sync.WaitGroup

const MQURL = "amqp://kevin:123456@192.168.72.130:5672/kevin"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
}

// Destroy 断开channel和connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// NewRabbitMQSimple 创建simple rabbitmq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// 创建rabbitmq实例
	rabbitMQ := NewRabbitMQ(queueName, "", "")
	var err error
	// 获取connection
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.failOnErr(err, "Failed to connect to RabbitMQ")
	// 获取channel
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "Failed to open a channel")
	return rabbitMQ
}

// PublishSimple simple模式生产
func (r *RabbitMQ) PublishSimple(message string) {
	// 申请队列，如果队列不存在自动创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 调用channel发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// simple模式消费
func (r *RabbitMQ) ConsumeSimple() {
	// 申请队列
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	wg.Add(1)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
		wg.Done()
	}()
	wg.Wait()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
