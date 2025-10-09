package grpc

import "github.com/gofiber/fiber/v2"

type OperationFeedAdapter struct {
}

func NewOperationFeedAdapter() *OperationFeedAdapter {
	return &OperationFeedAdapter{}
}

func (a *OperationFeedAdapter) GetTransfer(c *fiber.Ctx) error {
	return nil
}

func (a *OperationFeedAdapter) GetTransfers(c *fiber.Ctx) error {
	return nil
}
