package validator

import (
	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"
	"github.com/gobuffalo/validate"
)

type validator struct{}

type Validator interface {
	UserRegister(request *httpmodels.RegisterRequest) *validate.Errors
	UserLogin(request *httpmodels.LoginRequest) *validate.Errors
	UserCreateOrder(orderNumber string) *validate.Errors
}

func New() Validator {
	return &validator{}
}
