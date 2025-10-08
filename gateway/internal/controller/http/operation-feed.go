package http

import (
	"github.com/gofiber/fiber/v2"
)

type OperationFeedController struct {
}

func NewOperationFeedController() *OperationFeedController {
	return &OperationFeedController{}
}

func (c *OperationFeedController) GetTransfer(ctx *fiber.Ctx) error {
	return nil
}

func (c *OperationFeedController) GetTransfers(ctx *fiber.Ctx) error {
	return nil
}
