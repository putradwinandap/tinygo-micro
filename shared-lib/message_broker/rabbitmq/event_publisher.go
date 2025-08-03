package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type RabbitMQPublisher struct {
	conn *amqp.Connection
}

func NewRabbitMQPublisher(conn *amqp.Connection) *RabbitMQPublisher {
	return &RabbitMQPublisher{
		conn: conn,
	}
}

func (p *RabbitMQPublisher) Publish(topic string, payload any) error {
	ch, err := p.conn.Channel()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warn("Failed to open a channel")
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return err
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warn("Failed to marshal payload")
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {

		log.WithFields(log.Fields{
			"err": err,
		}).Warn("Failed to publish message")
		return err
	}

	return nil
}
