package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, groupID string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			GroupID: groupID,
		}),
	}
}

func (c *Consumer) Consume(ctx context.Context, handler func(msg []byte) error) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				return err
			}
			if err = handler(m.Value); err != nil {

			}
		}
	}
}
