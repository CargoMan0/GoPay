package jwt

import "errors"

var (
	ErrInvalidTokenType = errors.New("invalid token type")
	ErrInvalidToken     = errors.New("invalid token")
)
