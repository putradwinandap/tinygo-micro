package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type RabbitMQBroker struct {
	Conn *amqp.Connection
}

func NewRabbitMQBroker(dsn string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Warn("Failed to connect to RabbitMQ:", err)
		return nil, err
	}

	return conn, nil
}
