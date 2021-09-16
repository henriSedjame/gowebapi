package validations

import "github.com/go-playground/validator"

func IsValid(i interface{}) error {
	return validator.New().Struct(i)
}
