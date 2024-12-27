package config

import "github.com/go-playground/validator/v10"

func NewValidator(conf *Bootstrap) *validator.Validate {
	return validator.New()
}
