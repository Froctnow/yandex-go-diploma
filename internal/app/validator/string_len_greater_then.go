package validator

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gobuffalo/validate"
)

type StringLenGreaterThenValidator struct {
	Name    string
	Field   string
	Min     int
	Message string
}

func (v *StringLenGreaterThenValidator) IsValid(errors *validate.Errors) {
	strLength := utf8.RuneCountInString(v.Field)
	if v.Message == "" {
		v.Message = fmt.Sprintf("длина поля должна быть больше %d", v.Min)
	}
	if strLength <= v.Min {
		errors.Add(strings.ToLower(v.Name), v.Message)
	}
}
