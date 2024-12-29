package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"riz.it/nurul-faizah/internal/domain"
	"riz.it/nurul-faizah/internal/dto"
	"riz.it/nurul-faizah/internal/utils"
)

type ClassController struct {
	ClassUseCase domain.ClassUsecase
	Log          *logrus.Logger
}

func NewClassController(classUseCase domain.ClassUsecase, log *logrus.Logger) *ClassController {
	return &ClassController{
		ClassUseCase: classUseCase,
		Log:          log,
	}
}

func (c *ClassController) Create(ctx *fiber.Ctx) error {

	request := new(dto.CreateClassRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(&dto.ApiResponse[*dto.ClassResponse]{
			Status:  false,
			Message: "Validation failed",
			Errors:  fails,
			Data:    nil,
		})
	}

	response, err := c.ClassUseCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(&dto.ApiResponse[*dto.ClassResponse]{
		Status:  true,
		Message: "Class Created Successfully",
		Data:    &response,
	})
}
