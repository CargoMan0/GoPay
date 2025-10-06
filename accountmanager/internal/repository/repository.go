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
	const query = `INSERT INTO account.accounts(id,registration_date,username,password_hash,email,refresh_token_hash) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query, account.ID, account.RegistrationDate, account.Username, account.PasswordHash, account.Email, account.RefreshTokenHash)
	if err != nil {
		return fmt.Errorf("exec sql query: %w", err)
	}

	return nil
}

func (r *repository) GetAccountByID(ctx context.Context, id uuid.UUID) (*entity.Account, error) {
	const query = `SELECT registration_date,username,password_hash,email,refresh_token_hash FROM account.accounts WHERE id = $1`

	var res entity.Account

	err := r.db.QueryRowContext(ctx, query, id).Scan(&res.RegistrationDate, &res.Username, &res.PasswordHash, &res.Email, &res.RefreshTokenHash)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &res, nil
}

func (r *repository) UpdateAccountPassword(ctx context.Context, id uuid.UUID, password string) error {
	const query = `UPDATE account.accounts SET password_hash = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return fmt.Errorf("exec sql query: %w", err)
	}

	return nil
}
