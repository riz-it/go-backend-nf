package entity

import "database/sql"

type Class struct {
	ID        int          `gorm:"column:id;primaryKey"`
	Name      string       `gorm:"column:name"`
	IsActive  bool         `gorm:"column:is_active"`
	Leader    int          `gorm:"column:leader"`
	CreatedAt sql.NullTime `gorm:"created_at"`
	UpdatedAt sql.NullTime `gorm:"updated_at"`
}

func (Class) TableName() string {
	return "classes"
}
