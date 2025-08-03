package event

import (
	"github.com/putradwinandap/tinygo-micro/shared-lib/message_broker/iface"

	log "github.com/sirupsen/logrus"
)

type EventURLVisitLogUseCase struct {
	rabbitComs iface.EventConsumer
}

func NewEventURLVisitLogUseCase(rabbitComs iface.EventConsumer) *EventURLVisitLogUseCase {
	return &EventURLVisitLogUseCase{
		rabbitComs: rabbitComs,
	}
}

func (e *EventURLVisitLogUseCase) Execute() {
	e.rabbitComs.Consume("shorturl.resolved", func(payload any) {
		// casting if needed
		// data := payload.(map[string]interface{})
		// Simpan ke DB, logging, dll
		log.Info("Received URL visit log event:", payload)
	})
}
