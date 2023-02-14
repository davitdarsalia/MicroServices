package outerServices

import (
	mq "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func MqConnection() *mq.Connection {
	conn, err := mq.Dial(os.Getenv("RABBITMQ_CONN"))

	if err != nil {
		log.Printf("Failed to connect RabbitMQ server: %s", err.Error())
	}
	return conn
}
