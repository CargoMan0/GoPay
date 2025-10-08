package entity

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID               uuid.UUID
	RegistrationDate time.Time
	Username         string
	Email            string
}

type NewAccountData struct {
	Username string
	Email    string
	Password string
}

type NewAccountResult struct {
	ID               uuid.UUID
	RegistrationDate time.Time
	AccessToken      string
	RefreshToken     string
}

type ChangePasswordData struct {
	ID          uuid.UUID
	OldPassword string
	NewPassword string
}

type LoginAccountResult struct {
	AccessToken  string
	RefreshToken string
}
