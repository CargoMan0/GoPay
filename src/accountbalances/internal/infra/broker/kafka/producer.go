package kafka

import (
	"context"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/core"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.RoundRobin{},
		},
	}
}

func (p *Producer) Produce(ctx context.Context, key []byte, value []byte, topic core.EventName) error {
	kafkaTopic := fromDomainToKafka(topic)

	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
		Topic: kafkaTopic,
	})
}
