package async

import (
	"context"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/infra/broker/kafka"
)

type EventPublisher struct {
	producer *kafka.Producer
}

func NewEventPublisher(pr *kafka.Producer) *EventPublisher {
	return &EventPublisher{producer: pr}
}

func (e *EventPublisher) PublishEvent(ctx context.Context, topic string, payload []byte) error {
	return e.producer.Produce(ctx, []byte{}, payload, topic)
}
