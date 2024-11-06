package helpers

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func Validate(data any) (err error) {
	validate := validator.New()
	validate.RegisterValidation("password", validatePassword)
	return validate.Struct(data)
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		hasMinLen  = false
		hasUpper   = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasSpecial
}
