package server

import "github.com/gofiber/fiber/v2"

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
		},
	)

	ctrl.SetupRoutes(app)

	return app
}
