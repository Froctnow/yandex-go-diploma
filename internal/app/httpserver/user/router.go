package user

import (
	"github.com/Froctnow/yandex-go-diploma/internal/app/validator"
	"github.com/gin-gonic/gin"
)

type Router interface {
}

type userRouter struct {
	urlUseCase auth.UseCase
	validator  validator.Validator
}

func NewRouter(
	ginGroup *gin.RouterGroup,
	urlUseCase auth.UseCase,
	validator validator.Validator,
) Router {
	router := &userRouter{
		urlUseCase: urlUseCase,
		validator:  validator,
	}

	_ = ginGroup.Group("/api")

	return router
}
