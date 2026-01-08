package impl

import (
	"database/sql"
	"errors"
	"fmt"
)

type transactionRepositoryImpl struct {
	tx         *sql.Tx
	txFinished bool
}

func NewTransactionRepositoryImpl(tx *sql.Tx) *transactionRepositoryImpl {
	return &transactionRepositoryImpl{
		tx:         tx,
		txFinished: false,
	}
}

func (t *transactionRepositoryImpl) Commit() error {
	if t.txFinished {
		return nil
	}

	err := t.tx.Commit()
	if err != nil {
		if errors.Is(err, sql.ErrTxDone) {
			return nil
		}
		return fmt.Errorf("commit transaction: %w", err)
	}

	t.txFinished = true
	return nil
}

func (t *transactionRepositoryImpl) Rollback() error {
	if t.txFinished {
		return nil
	}

	err := t.tx.Rollback()
	if err != nil {
		if errors.Is(err, sql.ErrTxDone) {
			return nil
		}

		return fmt.Errorf("rollback transaction: %w", err)
	}

	t.txFinished = true
	return nil
}
