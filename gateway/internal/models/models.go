package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type NewTransferData struct {
	ID          uuid.UUID
	Amount      decimal.Decimal
	FromAddress string
	ToAddress   string
}
