package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"riz.it/nurul-faizah/internal/domain"
	"riz.it/nurul-faizah/internal/dto"
	"riz.it/nurul-faizah/internal/entity"
	"riz.it/nurul-faizah/internal/utils"
)

type ClassUseCase struct {
	DB              *gorm.DB
	Validate        *validator.Validate
	ClassRepository domain.ClassRepository
	Log             *logrus.Logger
}

func NewClassUseCase(db *gorm.DB, log *logrus.Logger, v *validator.Validate, classRepository domain.ClassRepository) domain.ClassUsecase {
	return &ClassUseCase{
		DB:              db,
		Validate:        v,
		Log:             log,
		ClassRepository: classRepository,
	}
}

// Create implements domain.ClassUsecase.
func (c *ClassUseCase) Create(ctx context.Context, req *dto.CreateClassRequest) (*dto.ClassResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	class := &entity.Class{
		Name: req.Name,
	}

	if err := c.ClassRepository.Create(tx, class); err != nil {
		c.Log.Warnf("failed create class to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return utils.ClassToResponse(class), nil
}
