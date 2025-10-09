package models

import (
	"github.com/google/uuid"
	"time"
)

type NewAccountData struct {
	Username string
	Email    string
	Password string
}

type NewAccountResult struct {
	WalletAddress string
	AccessToken   string
	RefreshToken  string
}

type RegisterData struct {
	Email     string
	Password  string
	Username  string
	IP        string
	UserAgent string
}

type RegisterResult struct {
	UserID    uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
}

type LoginResult struct {
	SessionID string
	ExpiresAt time.Time
}

type ValidateSessionResult struct {
	Username   string
	Email      string
	ExpiresAt  time.Time
	LastUsedAt time.Time
}
