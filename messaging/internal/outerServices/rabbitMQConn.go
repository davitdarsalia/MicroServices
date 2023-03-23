package outerServices

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func MQ(connStr *string) *amqp.Connection {
	conn, err := amqp.Dial(*connStr)
	if err != nil {
		log.Printf("Failed to connect RabbitMQ server: %s", err.Error())
	}

	return conn
}

// AuthQueue - Implementation defined only for auth service. Use this approach for different microservices with defined RoutingKeys and Queues
func AuthQueue(conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to connect to MQ Channel: %s", err.Error())
		return
	}

	err = exchange(ch, "auth_exchange")
	if err != nil {
		log.Printf("Failed to declare an exchange: %s", err.Error())
		return
	}

	err = queue(ch, "auth_queue")
	if err != nil {
		log.Printf("Failed to declare a queue: %s", err.Error())
		return
	}

	err = bindQueue(ch, "auth_queue", "auth", "auth_exchange")
	if err != nil {
		log.Printf("Failed to bind a queue: %s", err.Error())
		return
	}
}

func exchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(name, "direct", true, false, false, false, nil)
}

func queue(ch *amqp.Channel, name string) error {
	_, err := ch.QueueDeclare(name, true, false, false, false, nil)

	return err
}

func bindQueue(ch *amqp.Channel, queueName, routingKey, exchangeName string) error {
	return ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
}
