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

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
}

func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ("", exchangeName, "")
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.FailOnErr(err, "Failed to connect to RabbitMQ")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.FailOnErr(err, "Failed to open a channel")
	return rabbitMQ
}

// PublishPub 生产
func (r *RabbitMQ) PublishPub(message string) {
	// 创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.FailOnErr(err, "Failed to declare an exchange")

	// 发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumeSub 消费
func (r *RabbitMQ) ConsumeSub() {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.FailOnErr(err, "Failed to declare an exchange")
	// 尝试创建队列，注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.FailOnErr(err, "Failed to declare a queue")

	// 绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)

	// 消费消息
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	wg.Add(1)
	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Successfully consumed a message")
	<-forever
}
