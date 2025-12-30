package core

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type AccountBalance struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CurrencyID uint
	Amount     decimal.Decimal
	UpdatedAt  time.Time
}

type AccountBalanceAudit struct {
	ID        uint
	BalanceID uuid.UUID
	Operation OperationType
	Status    Status
	CreatedAt time.Time
}
