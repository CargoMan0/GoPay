package validator

import "errors"

var (
	ErrEmailInvalid      = errors.New("email is invalid")
	ErrPasswordTooShort  = errors.New("password must be at least 8 characters long")
	ErrPasswordNoUpper   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoLower   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordNoDigit   = errors.New("password must contain at least one digit")
	ErrPasswordNoSpecial = errors.New("password must contain at least one special character")
)
