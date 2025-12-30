package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	entity2 "github/com/CargoMan0/GoPay/src/accountmanager/internal/entity"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SaveAccountAndEventInTx(ctx context.Context, account *entity2.Account) (err error) {
	const (
		saveAccountQuery      = `INSERT INTO account.accounts(id,registration_date,username,password_hash,email,refresh_token_hash) VALUES ($1, $2, $3, $4, $5, $6)`
		saveAccountEventQuery = `INSERT INTO account.account_created_events(event_type,payload) VALUES ($1, $2)`
	)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			err = errors.Join(err, fmt.Errorf("rollback tx: %w", rollbackErr))
		}
	}()

	_, err = tx.ExecContext(ctx, saveAccountQuery, account.ID, account.RegistrationDate, account.Username, account.PasswordHash, account.Email, account.RefreshTokenHash)
	if err != nil {
		return fmt.Errorf("exec sql query: %w", err)
	}

	_, err = tx.ExecContext(ctx, saveAccountEventQuery)
	if err != nil {
		return fmt.Errorf("exec sql query: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (r *repository) GetAccountByID(ctx context.Context, id uuid.UUID) (*entity2.Account, error) {
	const query = `SELECT registration_date,username,password_hash,email,refresh_token_hash FROM account.accounts WHERE id = $1`

	var res = entity2.Account{
		ID: id,
	}

	err := r.db.QueryRowContext(ctx, query, id).Scan(&res.RegistrationDate, &res.Username, &res.PasswordHash, &res.Email, &res.RefreshTokenHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &res, nil
}

func (r *repository) GetAccountByEmail(ctx context.Context, email string) (*entity2.Account, error) {
	const query = `SELECT id,registration_date,username,password_hash,refresh_token_hash FROM account.accounts WHERE email = $1`

	var res = entity2.Account{
		Email: email,
	}

	err := r.db.QueryRowContext(ctx, query, email).Scan(&res.ID, &res.RegistrationDate, &res.Username, &res.PasswordHash, &res.RefreshTokenHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &res, nil
}

func (r *repository) GetEvents(ctx context.Context, limit uint8) (_ []entity2.Event, err error) {
	const query = `SELECT * FROM account.account_created_events LIMIT $1`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("query rows: %w", err)
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close rows: %w", closeErr))
		}
	}()

	var res []entity2.Event
	for rows.Next() {
		var event entity2.Event
		_ = event
		// TODO: finish
	}

	return res, nil
}

func (r *repository) UpdateAccountPassword(ctx context.Context, id uuid.UUID, password string) error {
	const query = `UPDATE account.accounts SET password_hash = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return fmt.Errorf("exec sql query: %w", err)
	}

	return nil
}
