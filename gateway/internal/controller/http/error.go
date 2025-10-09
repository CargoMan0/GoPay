package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) {
	switch {
	case errors.Is(err, fiber.ErrInternalServerError):
	}
}
