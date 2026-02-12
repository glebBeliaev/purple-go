package request

import (
	"github.com/go-playground/validator/v10"
)

func ValidatePayload[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}
	return nil
}
