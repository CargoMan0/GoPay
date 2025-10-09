package http

import (
	"errors"
	"fmt"
	"github.com/CargoMan0/GoPay/gateway/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, validator.ErrEmailInvalid):
		err = c.JSON(fiber.Map{
			"error": "Invalid email address",
		})
		if err != nil {
			return fmt.Errorf("write json: %w", err)
		}

		return nil
	case errors.Is(err, validator.ErrPasswordNoLower):
		err = c.JSON(fiber.Map{
			"error": "Password must contain at least one lowercase letter",
		})
		if err != nil {
			return fmt.Errorf("write json: %w", err)
		}

		return nil
	case errors.Is(err, validator.ErrPasswordNoUpper):
		err = c.JSON(fiber.Map{
			"error": "Password must contain at least one uppercase letter",
		})
		if err != nil {
			return fmt.Errorf("write json: %w", err)
		}

		return nil
	case errors.Is(err, validator.ErrPasswordNoDigit):
		err = c.JSON(fiber.Map{
			"error": "Password must contain at least one digit",
		})
		if err != nil {
			return fmt.Errorf("write json: %w", err)
		}

		return nil
	case errors.Is(err, validator.ErrPasswordTooShort):
		err = c.JSON(fiber.Map{
			"error": "Password must be at least 8 characters long",
		})
		if err != nil {
			return fmt.Errorf("write json: %w", err)
		}

		return nil
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unexpected internal error: " + err.Error(),
		})
	}
}
