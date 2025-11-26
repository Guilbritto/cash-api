package errors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		return nil
	}
	validationErrros := err.(validator.ValidationErrors)
	validationError := validationErrros[0]

	switch validationError.Tag() {
	case "required":
		return errors.New("field " + validationError.Field() + " is required")
	case "min":
		return errors.New("field " + validationError.Field() + " must be at least " + validationError.Param() + " characters long")
	case "max":
		return errors.New("field " + validationError.Field() + " must be at most " + validationError.Param() + " characters long")
	default:
		return errors.New("field " + validationError.Field() + " is invalid")
	}

	return nil
}
