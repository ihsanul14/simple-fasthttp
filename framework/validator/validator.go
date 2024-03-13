package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	Validate *validator.Validate
}

type IValidator interface {
	AddRules(map[string]string, interface{})
	Check(interface{}) error
}

func NewValidator() IValidator {
	validate := validator.New()
	return &Validator{Validate: validate}
}

func (v *Validator) AddRules(rule map[string]string, data interface{}) {
	v.Validate.RegisterStructValidationMapRules(rule, data)
}

func (v *Validator) Check(data interface{}) error {
	return v.Validate.Struct(data)
}
