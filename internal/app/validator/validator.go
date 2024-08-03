package validator

import (
	"github.com/gobuffalo/validate"

	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"
)

type validator struct{}

type Validator interface {
	UserRegister(request *httpmodels.RegisterRequest) *validate.Errors
	UserLogin(request *httpmodels.LoginRequest) *validate.Errors
	UserCreateOrder(orderNumber string) *validate.Errors
	UserWithdraw(orderNumber string) *validate.Errors
}

func New() Validator {
	return &validator{}
}
