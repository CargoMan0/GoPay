package http

import "github.com/gofiber/fiber/v2"

type AuthServiceController struct {
}

func NewAuthServiceController() *AuthServiceController {
	return &AuthServiceController{}
}

func (a *AuthServiceController) Login(ctx *fiber.Ctx) error {
	return nil
}

func (a *AuthServiceController) Register(ctx *fiber.Ctx) error {
	return nil
}

func (a *AuthServiceController) RefreshToken(ctx *fiber.Ctx) error {
	return nil
}

func (a *AuthServiceController) Logout(ctx *fiber.Ctx) error {
	return nil
}
