package util

import "github.com/go-playground/validator/v10"

var Validator *validator.Validate

func InitValidator() {
	Validator = validator.New()
}

func Validate(s interface{}) error {
	return Validator.Struct(s)
}
