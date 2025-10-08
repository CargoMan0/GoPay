package grpc

import "github.com/gofiber/fiber/v2"

type OperationFeedClient struct {
}

func NewOperationFeedAdapter() *OperationFeedClient {
	return &OperationFeedClient{}
}

func (a *OperationFeedClient) GetTransfer(c *fiber.Ctx) error {
	return nil
}

func (a *OperationFeedClient) GetTransfers(c *fiber.Ctx) error {
	return nil
}
