package http

import (
	"github.com/gofiber/fiber/v2"
)

type AccountManagerAdapter interface {
}

type AccountManagerController struct {
	accountManagerAdapter AccountManagerAdapter
}

func NewAccountManagerController(accountManagerClient AccountManagerAdapter) *AccountManagerController {
	return &AccountManagerController{
		accountManagerAdapter: accountManagerClient,
	}
}

func (a *AccountManagerController) GetAccount(c *fiber.Ctx) error {
	return nil
}

func (a *AccountManagerController) UpdateAccount(c *fiber.Ctx) error {
	return nil
}
