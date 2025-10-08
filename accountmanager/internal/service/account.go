package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github/com/CargoMan0/GoPay/accountmanager/internal/entity"
	"github/com/CargoMan0/GoPay/accountmanager/internal/repository"
)

type Repository interface {
	GetAccountByID(ctx context.Context, id uuid.UUID) (*entity.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error)

	SaveAccountAndEventInTx(ctx context.Context, account *entity.Account) error
	UpdateAccountPassword(ctx context.Context, id uuid.UUID, hash string) error
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hash, password string) (bool, error)
}

type AccountService struct {
	repo   Repository
	hasher PasswordHasher
}

func NewAccountService(
	repo Repository,
	hasher PasswordHasher,
) *AccountService {
	return &AccountService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *AccountService) NewAccount(ctx context.Context, data *entity.NewAccountData) (*entity.NewAccountResult, error) {
	account := &entity.Account{
		ID:       uuid.New(),
		Username: data.Username,
		Email:    data.Email,
	}

	err := s.repo.SaveAccountAndEventInTx(ctx, account)
	if err != nil {
		return nil, err
	}

	return &entity.NewAccountResult{
		ID: account.ID,
	}, nil
}

func (s *AccountService) GetAccount(ctx context.Context, id uuid.UUID) (*entity.Account, error) {
	account, err := s.repo.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrAccountNotFound
		}
		return nil, fmt.Errorf("repository: get account: %w", err)
	}

	return account, nil
}
