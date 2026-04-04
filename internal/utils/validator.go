package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var passwordRegex = regexp.MustCompile(`^[[:graph:]]{8,20}$`)

func ValidatePassword(fl validator.FieldLevel) bool {
	return passwordRegex.MatchString(fl.Field().String())
}
