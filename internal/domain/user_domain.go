package domain

import (
	"gorm.io/gorm"
	"riz.it/nurul-faizah/internal/entity"
)

type UserAccountRepository interface {
	Create(db *gorm.DB, c *entity.UserAccount) error
	Update(db *gorm.DB, c *entity.UserAccount) error
	FindByEmail(db *gorm.DB, user *entity.UserAccount, email string) error
	CountByEmail(db *gorm.DB, email string) (int64, error)
}
