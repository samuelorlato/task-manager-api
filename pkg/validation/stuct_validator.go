package validation

import "github.com/go-playground/validator/v10"

func ValidateStruct(d interface{}) error {
	validator := validator.New(validator.WithRequiredStructEnabled())

	err := validator.Struct(d)
	if err != nil {
		return err
	}

	return nil
}
