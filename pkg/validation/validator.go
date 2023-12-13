package validation

import "gopkg.in/validator.v2"

type Validator struct {}

func (v *Validator) Validate(d interface{}) error {
	err := validator.Validate(d)
	if err != nil {
		return err
	}

	return nil
}