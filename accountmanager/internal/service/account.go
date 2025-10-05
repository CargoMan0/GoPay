package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github/com/CargoMan0/GoPay/accountmanager/internal/entity"
	"github/com/CargoMan0/GoPay/accountmanager/internal/jwt"
	"github/com/CargoMan0/GoPay/accountmanager/internal/repository"
)

type Repository interface {
	SaveAccount(ctx context.Context, account *entity.Account) error
	UpdateAccountPassword(ctx context.Context, id uuid.UUID, hash string) error
	GetAccountByID(ctx context.Context, id uuid.UUID) (*entity.Account, error)
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hash, password string) (bool, error)
}

type TokenManager interface {
	GenerateToken(tokenType jwt.TokenType, userID uuid.UUID) (string, error)
	ValidateToken(token string) (bool, error)
}

type AccountService struct {
	repo         Repository
	hasher       PasswordHasher
	tokenManager TokenManager
}

func NewAccountService(
	repo Repository,
	hasher PasswordHasher,
	tokenManager TokenManager,
) *AccountService {
	return &AccountService{
		repo:         repo,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}

func (s *AccountService) NewAccount(ctx context.Context, data *entity.NewAccountData) (uuid.UUID, error) {
	userID := uuid.New()

	accessToken, err := s.tokenManager.GenerateToken(jwt.Access, userID)
	if err != nil {
		return uuid.Nil, err
	}

	account := &entity.Account{
		ID:       userID,
		Username: data.Username,
		Email:    data.Email,
	}

	hashed, err := s.hasher.HashPassword(data.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to hash password: %w", err)
	}
	account.PasswordHash = hashed

	err = s.repo.SaveAccount(ctx, account)
	if err != nil {
		return uuid.Nil, err
	}

	// TODO: Save account`s balance also
	// TODO: Transactional Outbox pattern

	return account.ID, nil
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

func (s *AccountService) ChangePassword(ctx context.Context, data *entity.ChangePasswordData) error {
	account, err := s.repo.GetAccountByID(ctx, data.ID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrAccountNotFound
		}
	}

	match, err := s.hasher.CompareHashAndPassword(account.PasswordHash, data.OldPassword)
	if err != nil {
		return fmt.Errorf("hasher: compare hash and password: %w", err)
	}
	if !match {
		return ErrWrongPassword
	}

	hashed, err := s.hasher.HashPassword(data.NewPassword)
	if err != nil {
		return fmt.Errorf("hasher: hash password: %w", err)
	}

	err = s.repo.UpdateAccountPassword(ctx, data.ID, hashed)
	if err != nil {
		return fmt.Errorf("repository: update account password: %w", err)
	}

	return nil
}
