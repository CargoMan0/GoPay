package controller

import (
	"github.com/CargoMan0/GoPay/gateway/internal/controller/http"
	"github.com/gofiber/fiber/v2"
)

func authMiddleware(authCtrl *http.AuthServiceController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
