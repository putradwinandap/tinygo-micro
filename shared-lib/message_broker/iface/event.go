package iface

type EventPublisher interface {
	Publish(topic string, payload any) error
}

type EventConsumer interface {
	Consume(topic string, handler func(payload any)) error
}
