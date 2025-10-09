package http

import (
	"github.com/gofiber/fiber/v2"
)

type TransferManagerAdapter interface{}

type TransferManagerController struct {
	adapter TransferManagerAdapter
}

func NewTransferManagerController(adapter TransferManagerAdapter) *TransferManagerController {
	return &TransferManagerController{
		adapter: adapter,
	}
}

func (t *TransferManagerController) CreateTransfer(ctx *fiber.Ctx) error {
	return nil
}

func (t *TransferManagerController) CancelTransfer(ctx *fiber.Ctx) error {
	return nil
}
