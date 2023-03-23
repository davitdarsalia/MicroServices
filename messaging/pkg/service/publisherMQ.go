package service

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func publishToMQ(mqConn *amqp.Connection, exchange string, value interface{}) {
	ch, err := mqConn.Channel()

	if err != nil {
		log.Printf("Failed to publish a message to mq: %s", err.Error())
		return
	}

	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true, false, false, false, nil,
	)

	ch.Close()
}
