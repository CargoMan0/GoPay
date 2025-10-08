package controller

import (
	"github.com/gofiber/fiber/v2"
	"github/com/CargoMan0/GoPay/gateway/internal/controller/http"
)

type controller struct {
	transferManagerCtrl *http.TransferManagerController
	operationFeedCtrl   *http.OperationFeedController
	accountManagerCtrl  *http.AccountManagerController
	authServiceCtrl     *http.AuthServiceController
}

func New(
	tmc *http.TransferManagerController,
	ofc *http.OperationFeedController,
) *controller {
	return &controller{
		transferManagerCtrl: tmc,
		operationFeedCtrl:   ofc,
	}
}

func (c *controller) SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	auth := apiV1.Group("/auth")
	auth.Post("/register", c.authServiceCtrl.Register)
	auth.Post("/login", c.authServiceCtrl.Login)

	protected := apiV1.Group("", authMiddleware(c.authServiceCtrl))

	accounts := protected.Group("/accounts")
	accounts.Get("/:id", c.accountManagerCtrl.GetAccount)
	accounts.Patch("/:id", c.accountManagerCtrl.UpdateAccount)

	transfers := protected.Group("/transfers")
	transfers.Post("/", c.transferManagerCtrl.CreateTransfer)
	transfers.Delete("/:id", c.transferManagerCtrl.CancelTransfer)

	operations := protected.Group("/operations")
	operations.Get("/transfers/:id", c.operationFeedCtrl.GetTransfer)
	operations.Get("/transfers", c.operationFeedCtrl.GetTransfers)
}
