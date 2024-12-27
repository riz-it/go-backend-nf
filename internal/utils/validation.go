package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			res[strings.ToLower(e.Field())] = e.Tag()
		}
	}
	return res
}
