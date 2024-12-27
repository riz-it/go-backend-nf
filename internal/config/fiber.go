package config

import (
	"github.com/gofiber/fiber/v2"
	"riz.it/nurul-faizah/internal/dto"
)

func NewFiber(config *Bootstrap) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.Server.Name,
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(&dto.ApiResponse[string]{
			Status:  false,
			Message: err.Error(),
			Errors:  nil,
		})
	}
}
