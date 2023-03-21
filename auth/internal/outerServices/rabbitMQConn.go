package outerServices

import (
	mq "github.com/rabbitmq/amqp091-go"
	"log"
)

func MqConnection(connStr *string) *mq.Connection {
	conn, err := mq.Dial(*connStr)

	if err != nil {
		log.Printf("Failed to connect RabbitMQ server: %s", err.Error())
	}

	return conn
}
