package validator

import (
	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"

	"github.com/gobuffalo/validate"
)

func (v *validator) UserRegister(data *httpmodels.RegisterRequest) *validate.Errors {
	checks := []validate.Validator{
		&StringLenGreaterThenValidator{
			Name:  "Login",
			Field: data.Login,
			Min:   1,
		},
		&StringLenGreaterThenValidator{
			Name:  "Password",
			Field: data.Password,
		},
	}
	errors := validate.Validate(checks...)
	return errors
}
