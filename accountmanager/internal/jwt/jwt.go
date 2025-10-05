package jwt

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType uint8

const (
	Refresh TokenType = iota + 1
	Access
)

type TokenManager struct {
	secret []byte
}

func NewTokenManager(secret string) *TokenManager {
	return &TokenManager{
		secret: []byte(secret),
	}
}

func (t *TokenManager) GenerateToken(tokenType TokenType, userID uuid.UUID) (string, error) {
	var expiration time.Duration
	switch tokenType {
	case Access:
		expiration = time.Minute * 15
	case Refresh:
		expiration = time.Hour * 24 * 7
	default:
		return "", ErrInvalidTokenType
	}

	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"exp":  time.Now().Add(expiration).Unix(),
		"iat":  time.Now().Unix(),
		"type": tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.secret)
}

func (t *TokenManager) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return t.secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return token, nil
}
