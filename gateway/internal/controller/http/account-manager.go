package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github/com/CargoMan0/GoPay/gateway/internal/models"
)

type AccountManagerClient interface {
	NewAccount(ctx context.Context, data *models.NewAccountData) (*models.NewAccountResult, error)
}

type AccountManagerController struct {
	accountManagerClient AccountManagerClient
}

func NewAccountManagerController(accountManagerClient AccountManagerClient) *AccountManagerController {
	return &AccountManagerController{
		accountManagerClient: accountManagerClient,
	}
}

func (a *AccountManagerController) CreateAccount(c *fiber.Ctx) error {
	var req newAccountRequest

	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	if len(req.Username) <= 0 || len(req.Username) >= 16 {
		err = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username length must be between 0 and 16",
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return nil
		}
		return nil
	}

	if len(req.Password) <= 0 || len(req.Password) >= 16 {
		err = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password length must be between 0 and 16",
		})
	}

	// TODO: valid email

	data := &models.NewAccountData{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	ctx := c.Context()
	account, err := a.accountManagerClient.NewAccount(ctx, data)
	if err != nil {
		// TODO: handle error
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	resp := newAccountResponse{
		WalletAddress: account.WalletAddress,
		AccessToken:   account.AccessToken,
		RefreshToken:  account.RefreshToken,
	}

	err = c.JSON(resp)
	if err != nil {
		return fmt.Errorf("write json: %w", err)
	}

	return nil
}

func (a *AccountManagerController) GetAccount(c *fiber.Ctx) error {
	return nil
}

func (a *AccountManagerController) UpdateAccount(c *fiber.Ctx) error {
	return nil
}
