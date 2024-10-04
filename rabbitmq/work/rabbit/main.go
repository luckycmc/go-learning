package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var wg sync.WaitGroup

const MQURL = "amqp://kevin:123456@192.168.72.130:5672/kevin"

// RabbitMQ rabbitmq结构体
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

func (r *RabbitMQ) FailOnErr(err error, messages string) {
	if err != nil {
		log.Fatalf("%s:%s", messages, err)
		panic(fmt.Sprintf("%s:%s", messages, err))
	}
}

func NewRabbitMQWork(queueName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.FailOnErr(err, "Failed to connect to RabbitMQ")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.FailOnErr(err, "Failed to open a channel")
	return rabbitMQ
}

func (r *RabbitMQ) PublishWork(message string) {
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

func (r *RabbitMQ) ConsumeWork() {
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
