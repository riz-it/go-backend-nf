package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"riz.it/nurul-faizah/internal/domain"
	"riz.it/nurul-faizah/internal/dto"
	"riz.it/nurul-faizah/internal/utils"
)

type AuthController struct {
	AuthUseCase domain.AuthUseCase
	Log         *logrus.Logger
}

func NewAuthController(authUseCase domain.AuthUseCase, log *logrus.Logger) *AuthController {
	return &AuthController{
		AuthUseCase: authUseCase,
		Log:         log,
	}
}

func (c *AuthController) SignUp(ctx *fiber.Ctx) error {

	request := new(dto.SignUpRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(&dto.ApiResponse[*dto.SignUpResponse]{
			Status:  false,
			Message: "Validation failed",
			Errors:  fails,
			Data:    nil,
		})
	}

	response, err := c.AuthUseCase.SignUp(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(&dto.ApiResponse[*dto.SignUpResponse]{
		Status:  true,
		Message: "Sign Up Success",
		Data:    &response,
	})
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {

	request := new(dto.SignInRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(&dto.ApiResponse[*dto.SignInResponse]{
			Status:  false,
			Message: "Validation failed",
			Errors:  fails,
			Data:    nil,
		})
	}

	response, err := c.AuthUseCase.SignIn(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(&dto.ApiResponse[*dto.SignInResponse]{
		Status:  true,
		Message: "Sign In Success",
		Data:    &response,
	})
}
