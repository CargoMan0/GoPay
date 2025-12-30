package service

import "errors"

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")

	ErrWrongPassword = errors.New("wrong password")
)
