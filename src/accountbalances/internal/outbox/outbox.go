package outbox

import "context"

type EventSender struct {
}

func NewEventSender() *EventSender {
	return &EventSender{}
}

func (e *EventSender) Run(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	}
}
