package repository

import (
	"context"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/core"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/outbox"
	"github.com/google/uuid"
)

type Repository interface {
	BeginTransaction(ctx context.Context) (TransactionRepository, error)

	AccountBalanceRepository
}

type TransactionRepository interface {
	AccountBalanceRepository
	OutboxRepository

	Commit() error
	Rollback() error
}

type OutboxRepository interface {
	InsertOutboxEventMessage(ctx context.Context, message *outbox.Message) error
	UpdateOutboxMessageStatus(ctx context.Context, message *outbox.Message) error
	GetAccountBalanceOutboxMessages(ctx context.Context, batchSize int) ([]outbox.Message, error)
}

type AccountBalanceRepository interface {
	BatchInsertAccountBalance(ctx context.Context, balances []core.AccountBalance) error
	InsertAccountBalanceAudit(ctx context.Context, audit core.AccountBalanceAudit) error
	GetAccountBalances(ctx context.Context, userID uuid.UUID) ([]core.AccountBalance, error)
	GetAllCurrencies(ctx context.Context) ([]core.Currency, error)
}
