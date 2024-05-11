package validator

import "github.com/gobuffalo/validate"
import httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"

type validator struct{}

type Validator interface {
	UserRegister(request *httpmodels.RegisterRequest) *validate.Errors
}

func New() Validator {
	return &validator{}
}
