package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type field interface{}

// Create a custom error message.
func msgForTag(fe validator.FieldError, fieldName string) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required to contain values.", fieldName)
	case "min":
		return fmt.Sprintf("This %s field must have a minimum length of 12 characters.", fieldName)
	case "eqfield":
		if fieldName == "confirm_password" {
			return fmt.Sprintf("The %s field must be equal to the new_password field", fieldName)
		}
	}
	return ""
}

// Validate a field
func Validate(field interface{}) string {
	validate := validator.New()
	e := ""
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(field)
	if err == nil {
		return e
	}

	validationErrors := err.(validator.ValidationErrors)
	validationErr := validationErrors[0]
	fieldName := validationErr.Field()
	e = msgForTag(validationErr, fieldName)
	return e
}
