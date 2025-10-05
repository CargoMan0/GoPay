package entity

import "github.com/google/uuid"

type Account struct {
	ID           uuid.UUID
	Username     string
	Email        string
	PasswordHash string
}

type NewAccountData struct {
	Username string
	Email    string
	Password string
}

type ChangePasswordData struct {
	ID          uuid.UUID
	OldPassword string
	NewPassword string
}
