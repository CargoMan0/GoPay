package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
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

const (
	TokenTypeAccess  = "ACCESS"
	TokenTypeRefresh = "REFRESH"
)

type TokenManager interface {
	GenerateToken(tokenType string, userID uuid.UUID) (string, error)
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

func (s *AccountService) NewAccount(ctx context.Context, data *entity.NewAccountData) (*entity.NewAccountResult, error) {
	userID := uuid.New()

	refreshToken, err := s.tokenManager.GenerateToken(TokenTypeRefresh, userID)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	hashed, err := s.hasher.HashPassword(data.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	account := &entity.Account{
		ID:               userID,
		Username:         data.Username,
		Email:            data.Email,
		RefreshTokenHash: hashToken(refreshToken),
		PasswordHash:     hashed,
	}

	err = s.repo.SaveAccountAndEventInTx(ctx, account)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenManager.GenerateToken(TokenTypeAccess, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &entity.NewAccountResult{
		ID:           account.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
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

func (s *AccountService) LoginAccount(ctx context.Context, email, password string) (*entity.LoginAccountResult, error) {
	account, err := s.repo.GetAccountByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrAccountNotFound
		}
		return nil, fmt.Errorf("repository: get account by email: %w", err)
	}

	matches, err := s.hasher.CompareHashAndPassword(account.PasswordHash, password)
	if err != nil {
		return nil, fmt.Errorf("hasher: compare hash and password: %w", err)
	}
	if !matches {
		return nil, ErrWrongPassword
	}

	accessToken, err := s.tokenManager.GenerateToken(TokenTypeAccess, account.ID)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := s.tokenManager.GenerateToken(TokenTypeRefresh, account.ID)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	return &entity.LoginAccountResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
