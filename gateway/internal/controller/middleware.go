package controller

import (
	"github.com/gofiber/fiber/v2"
	"github/com/CargoMan0/GoPay/gateway/internal/controller/http"
)

func authMiddleware(authCtrl *http.AuthServiceController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
