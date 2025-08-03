package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
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

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload.(string)),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
