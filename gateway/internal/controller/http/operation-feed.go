package http

import (
	"github.com/gofiber/fiber/v2"
)

type OperationFeedAdapter interface{}

type OperationFeedController struct {
	adapter OperationFeedAdapter
}

func NewOperationFeedController(adapter OperationFeedAdapter) *OperationFeedController {
	return &OperationFeedController{
		adapter: adapter,
	}
}

func (c *OperationFeedController) GetTransfer(ctx *fiber.Ctx) error {
	return nil
}

func (c *OperationFeedController) GetTransfers(ctx *fiber.Ctx) error {
	return nil
}
