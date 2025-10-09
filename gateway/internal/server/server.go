package server

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Controller interface {
	SetupRoutes(app *fiber.App)
}

func NewFiberApp(ctrl Controller) *fiber.App {
	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
			Prefork:               false,
			CaseSensitive:         true,
			AppName:               "Gateway",
			ReadTimeout:           5 * time.Second,
			WriteTimeout:          5 * time.Second,
			ErrorHandler: func(c *fiber.Ctx, err error) error {

			},
		},
	)

	ctrl.SetupRoutes(app)

	return app
}
