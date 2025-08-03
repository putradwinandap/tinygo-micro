package http

import (
	"analytics/config"

	"github.com/putradwinandap/tinygo-micro/shared-lib/message_broker/rabbitmq"

	usecase "analytics/internal/usecase/event"

	log "github.com/sirupsen/logrus"
)

func SetupRouter(cfg *config.Config) {

	rabbitConn, err := rabbitmq.NewRabbitMQBroker(cfg.RabbitMQURL)
	if err != nil {
		log.Fatal("RabbitMQ connection failed:", err)
	}

	rabbitComs := rabbitmq.NewRabbitMQConsumer(rabbitConn)

	EventURLVisitLogUseCase := usecase.NewEventURLVisitLogUseCase(rabbitComs)

	EventURLVisitLogUseCase.Execute()

	select {}

	/*
		r := gin.Default()

		return r
	*/
}
