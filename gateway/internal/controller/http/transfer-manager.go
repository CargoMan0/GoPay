package http

import (
	"github.com/CargoMan0/GoPay/gateway/internal/adapter/grpc"
	"github.com/gofiber/fiber/v2"
)

type TransferManagerController struct {
	// ...
}

func NewTransferManagerController() *TransferManagerController {
	return &TransferManagerController{}

}

func (t *TransferManagerController) CreateTransfer(ctx *fiber.Ctx) error {
	return nil
}

func (t *TransferManagerController) CancelTransfer(ctx *fiber.Ctx) error {
	return nil
}
