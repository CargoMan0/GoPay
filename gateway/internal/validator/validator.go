package validator

import (
	"regexp"
	"unicode"
)

type Validator struct {
	minPasswordLength int
}

func New() *Validator {
	return &Validator{
		minPasswordLength: 8,
	}
}

func (v *Validator) WithMinPasswordLength(length int) *Validator {
	v.minPasswordLength = length
	return v
}

func (v *Validator) ValidateEmail(email string) error {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil || !matched {
		return ErrEmailInvalid
	}
	return nil
}

func (v *Validator) ValidatePassword(password string) error {
	if len(password) < v.minPasswordLength {
		return ErrPasswordTooShort
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return ErrPasswordNoUpper
	}
	if !hasLower {
		return ErrPasswordNoLower
	}
	if !hasDigit {
		return ErrPasswordNoDigit
	}
	if !hasSpecial {
		return ErrPasswordNoSpecial
	}

	return nil
}
