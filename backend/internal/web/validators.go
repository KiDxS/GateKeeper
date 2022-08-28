package web

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

// Create a custom error message.
func msgForTag(fe validator.FieldError, fieldName string) (err string) {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required to contain values.", fieldName)
	case "min":
		return fmt.Sprintf("This %s field must have a minimum length of 12 characters.", fieldName)
	case "eqfield":
		switch fieldName {
		case "confirm_password":
			return fmt.Sprintf("The %s field must be equal to the new_password field", fieldName)
		}

	}
	return
}

// Validate a field
func Validate(fieldStruct interface{}) (e string) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(fieldStruct)
	if err == nil {
		return
	}

	validationErrors := err.(validator.ValidationErrors)
	validationErr := validationErrors[0]
	fieldName := validationErr.Field()
	log.Warn().Msg(validationErrors.Error())
	e = msgForTag(validationErr, fieldName)
	return
}