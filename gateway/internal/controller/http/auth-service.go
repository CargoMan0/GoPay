package http

import (
	"context"
	"fmt"
	"github.com/CargoMan0/GoPay/gateway/internal/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Validator interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
}

type AuthServiceAdapter interface {
	Register(ctx context.Context, data *models.RegisterData) (*models.RegisterResult, error)
	Login(ctx context.Context, data *models.RegisterData) (*models.LoginResult, error)
	Logout(ctx context.Context, sessionID string) error
}

type AuthServiceController struct {
	adapter   AuthServiceAdapter
	validator Validator
}

func NewAuthServiceController(adapter AuthServiceAdapter, validator Validator) *AuthServiceController {
	return &AuthServiceController{
		adapter:   adapter,
		validator: validator,
	}
}

func (a *AuthServiceController) Login(c *fiber.Ctx) error {
	var req registerRequest

	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	if len(req.Username) >= 16 || len(req.Username) <= 0 {
		err = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong username size, should be between 0 and 16",
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
		}

		return nil
	}

	err = a.validator.ValidateEmail(req.Email)
	if err != nil {
		return handleError(c, err)
	}
	err = a.validator.ValidatePassword(req.Password)
	if err != nil {
		return handleError(c, err)
	}

	data := &models.RegisterData{
		Email:     req.Email,
		Username:  req.Username,
		Password:  req.Password,
		IP:        c.IP(),
		UserAgent: c.GetRespHeader(fiber.HeaderUserAgent),
	}

	ctx := c.Context()
	res, err := a.adapter.Login(ctx, data)
	if err != nil {
		return handleError(c, err)
	}

	resp := loginResponse{
		SessionID: res.SessionID,
		ExpiresAt: res.ExpiresAt.Format(time.RFC3339),
	}

	err = c.JSON(resp)
	if err != nil {
		return fmt.Errorf("write json: %w", err)
	}

	return nil
}

func (a *AuthServiceController) Register(c *fiber.Ctx) error {
	var req registerRequest

	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	if len(req.Username) >= 16 || len(req.Username) <= 0 {
		err = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong username size, should be between 0 and 16",
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
		}

		return nil
	}

	err = a.validator.ValidateEmail(req.Email)
	if err != nil {
		return handleError(c, err)
	}
	err = a.validator.ValidatePassword(req.Password)
	if err != nil {
		return handleError(c, err)
	}

	data := &models.RegisterData{
		Email:     req.Email,
		Username:  req.Username,
		Password:  req.Password,
		IP:        c.IP(),
		UserAgent: c.GetRespHeader(fiber.HeaderUserAgent),
	}

	ctx := c.Context()
	res, err := a.adapter.Register(ctx, data)
	if err != nil {
		return handleError(c, err)
	}

	resp := registerResponse{
		Email:     req.Email,
		Username:  req.Username,
		UserID:    res.UserID.String(),
		CreatedAt: res.CreatedAt.Format(time.RFC3339),
	}

	err = c.JSON(resp)
	if err != nil {
		return fmt.Errorf("write json: %w", err)
	}

	return nil
}

func (a *AuthServiceController) Logout(c *fiber.Ctx) error {
	var req logoutRequest

	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	ctx := c.Context()
	err = a.adapter.Logout(ctx, req.SessionID)
	if err != nil {
		return handleError(c, err)
	}

	c.Status(fiber.StatusOK)
	return nil
}
