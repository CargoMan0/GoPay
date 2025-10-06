package eventsender

import (
	"context"
	"github/com/CargoMan0/GoPay/accountmanager/internal/config"
	"github/com/CargoMan0/GoPay/accountmanager/internal/entity"
	"log/slog"
	"time"
)

type Producer interface {
	SendMessage(ctx context.Context, key, value []byte) error
}

type Repository interface {
	GetEvents(ctx context.Context, limit uint8) ([]entity.Event, error)
}

type EventSender struct {
	producer Producer
	repo     Repository

	handlePeriod time.Duration
	maxBatchSize uint8
}

func New(cfg config.EvenSender, producer Producer) *EventSender {
	return &EventSender{
		producer:     producer,
		handlePeriod: time.Duration(cfg.HandlePeriodSeconds) * time.Second,
		maxBatchSize: cfg.MaxBatchSize,
	}
}

func (s *EventSender) RunEventProcessing(ctx context.Context) {
	ticker := time.NewTicker(s.handlePeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.Info("EventSender: Context canceled")
			return
		case <-ticker.C:
			events, err := s.repo.GetEvents(ctx, s.maxBatchSize)
			if err != nil {
				slog.Info("EventSender: Error getting events: ", slog.String("error", err.Error()))
				continue
			}

			for event := range events {
				_ = event
				// TODO: finish
			}
		}
	}
}

func (s *EventSender) SendEvent(ctx context.Context, event *entity.Event) error {
	return nil
}
