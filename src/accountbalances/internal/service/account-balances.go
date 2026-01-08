package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/core"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/outbox"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Service struct {
	repository repository.Repository
}

func New(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateAccountBalance(ctx context.Context, userID uuid.UUID) (err error) {
	tx, err := s.repository.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("repository: begin transaction for account balance creating: %w", err)
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = errors.Join(err, fmt.Errorf("rollback error: %w", rollbackErr))
			}
		}
	}()

	currencies, err := tx.GetAllCurrencies(ctx)
	if err != nil {
		err = fmt.Errorf("tx: account balance repository: get all currencies: %w", err)
		return err
	}

	var (
		balances = make([]core.AccountBalance, 0, len(currencies))
		balance  core.AccountBalance
	)

	timeNow := time.Now()
	for _, currency := range currencies {
		balance = core.AccountBalance{
			ID:         uuid.New(),
			UserID:     userID,
			Amount:     decimal.Zero,
			CurrencyID: currency.ID,
			UpdatedAt:  timeNow,
		}

		balances = append(balances, balance)
	}

	err = tx.BatchInsertAccountBalance(ctx, balances)
	if err != nil {
		err = fmt.Errorf("tx: account balance repository: batch insert account balance: %w", err)
		return err
	}

	event := core.BalanceCreatedEvent{
		UserID: userID,
	}
	message := &outbox.Message{
		Name:    event.Name(),
		Status:  outbox.MessageStatusPending,
		Payload: []byte(userID.String()),
	}
	err = tx.InsertOutboxEventMessage(ctx, message)
	if err != nil {
		err = fmt.Errorf("tx: outbox repository: insert account balance outbox message: %w", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("tx repository: commit account balance creation: %w", err)
		return err
	}

	return nil
}

/*
func (s *Service) CreateAccountBalanceAudit(ctx context.Context) error {
	audit := core.AccountBalanceAudit{
		ЩзStatus:    core.StatusPending,
		CreatedAt: time.Now(),
	}

	err := s.repository.InsertAccountBalanceAudit(ctx, audit)
	if err != nil {
		return fmt.Errorf("repository: create account balance audit: %w", err)
	}

	return nil
}

func (s *Service) GetAllAccountBalances(ctx context.Context, accountID uuid.UUID) ([]core.AccountBalance, error) {
	accountBalances, err := s.repository.GetAccountBalances(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("repository: get account balances: %w", err)
	}

	return accountBalances, nil
}
*/
