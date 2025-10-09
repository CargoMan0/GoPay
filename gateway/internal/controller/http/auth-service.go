package http

import (
	"github.com/gofiber/fiber/v2"
	"github/com/CargoMan0/GoPay/gateway/internal/models"
)

type Validator interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
}

type AuthServiceClient interface {
}

type AuthServiceController struct {
	authServiceClient AuthServiceClient
	validator         Validator
}

func NewAuthServiceController(client AuthServiceClient, validator Validator) *AuthServiceController {
	return &AuthServiceController{
		authServiceClient: client,
		validator:         validator,
	}
}

func (a *AuthServiceController) Login(c *fiber.Ctx) error {
	var req registerRequest

	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	if len(req.Username) >= 16 || len(req.Username) <= 0 {
		err = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong username size, should be between 0 and 16",
		})
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
		}

		return nil
	}

	err = a.validator.ValidateEmail(req.Email)
	if err != nil {

	}
	err = a.validator.ValidatePassword(req.Password)
	if err != nil {

	}

	data := &models.RegisterData{
		Email:     req.Email,
		Username:  req.Username,
		Password:  req.Password,
		IP:        c.IP(),
		UserAgent: c.GetRespHeader(fiber.HeaderUserAgent),
	}

	ctx := c.Context()
	resp, err := a.authServiceClient.Login(ctx, data)
	if err != nil {
		handleError(c, err)
	}

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
