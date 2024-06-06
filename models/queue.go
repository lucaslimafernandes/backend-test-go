package models

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitMQChannel *amqp.Channel
var NotifyQueue amqp.Queue
var UnsafeQueue amqp.Queue

func RConn() {

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_HOST"))
	if err != nil {
		log.Fatalln("Failed to connect RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("Failed to open a RabbitMQ channel:", err)
	}

	q, err := ch.QueueDeclare(
		"notify",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalln("Failed to declare a queue:", err)
	}

	u, err := ch.QueueDeclare(
		"unsafe",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalln("Failed to declare a queue:", err)
	}

	RabbitMQChannel = ch
	NotifyQueue = q
	UnsafeQueue = u

}
