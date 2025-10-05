package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github/com/CargoMan0/GoPay/accountmanager/internal/entity"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SaveAccount(ctx context.Context, account *entity.Account) error {
	const query = `INSERT INTO todo: finish`

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("exec sql query: %w", err)
	}

	return nil
}

func (r *repository) GetAccountByID(ctx context.Context, id uuid.UUID) (*entity.Account, error) {
	const query = `SELECT`

	var res entity.Account

	err := r.db.QueryRowContext(ctx, query, id).Scan()
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &res, nil
}

func (r *repository) UpdateAccountPassword(ctx context.Context, id uuid.UUID, password string) error {
	const query = `UPDATE SET password = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return fmt.Errorf("exec sql query: %w", err)
	}

	return nil
}
