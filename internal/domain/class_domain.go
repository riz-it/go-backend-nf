package domain

import (
	"context"

	"gorm.io/gorm"
	"riz.it/nurul-faizah/internal/dto"
	"riz.it/nurul-faizah/internal/entity"
)

type ClassRepository interface {
	Create(db *gorm.DB, c *entity.Class) error
	// Update(db *gorm.DB, c *entity.Class) error
	// Delete(db *gorm.DB, c *entity.Class) error
	// FindAll(db *gorm.DB, c *[]entity.Class) error
	// FindByID(db *gorm.DB, user *entity.UserAccount, id uint) error
}

type ClassUsecase interface {
	Create(ctx context.Context, req *dto.CreateClassRequest) (*dto.ClassResponse, error)
	// Update(ctx context.Context, req *dto.SignInRequest) (*dto.SignInResponse, error)
	// Delete(ctx context.Context, userID uint) error
	// FindAll(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.SignInResponse, error)
	// FindByID(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.SignInResponse, error)
}
