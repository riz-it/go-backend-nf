package entity

import "database/sql"

type UserAccount struct {
	ID          int          `gorm:"column:id;primaryKey"`
	Email       string       `gorm:"column:email"`
	Password    string       `gorm:"column:password"`
	AccountName string       `gorm:"column:account_name"`
	HashedRt    string       `gorm:"column:hashed_rt"`
	IsActive    bool         `gorm:"column:is_active"`
	CreatedAt   sql.NullTime `gorm:"created_at"`
	UpdatedAt   sql.NullTime `gorm:"updated_at"`
}

func (UserAccount) TableName() string {
	return "user_accounts"
}
