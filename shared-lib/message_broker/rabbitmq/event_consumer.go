package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type RabbitMQConsumer struct {
	conn *amqp.Connection
}

func NewRabbitMQConsumer(conn *amqp.Connection) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		conn: conn,
	}
}

func (p *RabbitMQConsumer) Consume(topic string, handler func(payload any)) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	//defer ch.Close()

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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return err
	}

	/*
		go func() {
			for msg := range msgs {
				var payload any
				if err := json.Unmarshal(msg.Body, &payload); err != nil {
					log.Warn("Failed to unmarshal message:", err)
					continue
				}
				go handler(payload)
			}
		}()
	*/

	for msg := range msgs {
		var payload any
		log.Info("Received message:", string(msg.Body))
		_ = json.Unmarshal(msg.Body, &payload) // optional
		handler(payload)
	}

	return nil
}
