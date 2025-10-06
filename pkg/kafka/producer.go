package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type ProducerConfig struct {
	Brokers []string
	Topic   string
}

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(cfg ProducerConfig) *Producer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Brokers...),
		Topic:        cfg.Topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}

	return &Producer{writer: w}
}

func (p *Producer) SendMessage(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
