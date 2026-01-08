package core

import (
	"github.com/google/uuid"
)

type EventName string

const (
	AccountBalanceCreated        EventName = "AccountBalanceCreated"
	AccountBalanceCreationFailed EventName = "AccountBalanceCreationFailed"
)

type DomainEvent interface {
	Name() EventName
}

type BalanceCreatedEvent struct {
	UserID uuid.UUID
}

func (BalanceCreatedEvent) Name() EventName {
	return AccountBalanceCreated
}

type BalanceCreationFailedEvent struct {
	UserID uuid.UUID
}

func (BalanceCreationFailedEvent) Name() EventName {
	return AccountBalanceCreationFailed
}
