package hasher

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
}

func NewPasswordHasher() *PasswordHasher {

	return &PasswordHasher{}
}

func (p *PasswordHasher) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hash), nil
}

func (p *PasswordHasher) CompareHashAndPassword(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, fmt.Errorf("failed to compare hash and password: %w", err)
	}

	return true, nil
}
