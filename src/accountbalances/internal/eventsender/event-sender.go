package eventsender

import (
	"context"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/async"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/outbox"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/repository"
	"log/slog"
	"time"
)

const (
	eventsBatchSize = 100
)

type EventSender struct {
	outboxRepository repository.OutboxRepository
	publisher        *async.EventPublisher
}

func NewEventSender(publisher *async.EventPublisher) *EventSender {
	return &EventSender{publisher: publisher}
}

func (e *EventSender) Start(ctx context.Context) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.Info("Shut down signal received, shutting down event sender")
			return nil
		case <-ticker.C:
			messages, err := e.outboxRepository.GetAccountBalanceOutboxMessages(ctx, eventsBatchSize)
			if err != nil {
				slog.Error("failed to get account balance outbox messages, retrying...",
					slog.String("error", err.Error()),
				)

				continue
			}

			for _, message := range messages {
				err = e.publisher.PublishEvent(ctx, message.Topic, message.Payload)
				if err != nil {
					slog.Error("failed to publish event",
						slog.String("message.Topic", message.Topic),
					)

					continue
				}

				message.Status = outbox.MessageStatusSent
				err = e.outboxRepository.UpdateOutboxMessageStatus(ctx, &message)
				if err != nil {
					slog.Error("failed to update outbox message status",
						slog.String("message.Topic", message.Topic),
						slog.String("error", err.Error()),
					)

					continue
				}
			}
		}
	}
}
