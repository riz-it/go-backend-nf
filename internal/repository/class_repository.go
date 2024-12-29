package repository

import (
	"github.com/sirupsen/logrus"
	"riz.it/nurul-faizah/internal/entity"
)

type ClassRepository struct {
	Repository[entity.Class]
	Log *logrus.Logger
}

func NewClass(log *logrus.Logger) *ClassRepository {
	return &ClassRepository{
		Log: log,
	}
}
