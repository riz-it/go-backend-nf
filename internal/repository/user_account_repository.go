package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"riz.it/nurul-faizah/internal/entity"
)

type UserAccountRepository struct {
	Repository[entity.UserAccount]
	Log *logrus.Logger
}

func NewUserAccount(log *logrus.Logger) *UserAccountRepository {
	return &UserAccountRepository{
		Log: log,
	}
}

// FindByEmail implements domain.UserAccountRepository.
func (uA *UserAccountRepository) FindByEmail(db *gorm.DB, user *entity.UserAccount, email string) error {
	return db.Model(&entity.UserAccount{}).Where("email = ?", email).First(&user).Error
}

func (r *UserAccountRepository) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Model(&entity.UserAccount{}).Where("email = ?", email).Count(&total).Error
	return total, err
}
