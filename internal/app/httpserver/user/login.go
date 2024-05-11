package user

import (
	"fmt"
	"net/http"

	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"
	autherrors "github.com/Froctnow/yandex-go-diploma/internal/app/usecase/auth/errors"
	"github.com/gin-gonic/gin"

	"github.com/pkg/errors"
)

func (r *userRouter) Login(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := r.checkHeaderContentType(headerContentType)

	if !isCorrectHeaderContentType {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	var req httpmodels.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	errs := r.validator.UserLogin(&req)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	err := r.authUseCase.Login(ctx, req.Login, req.Password)

	if err != nil && errors.As(err, &autherrors.IncorrectLoginPasswordError{}) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err != nil {
		r.logger.ErrorCtx(ctx, fmt.Errorf("failed login"))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	err = r.authorizeUser(ctx, req.Login)

	if err != nil {
		r.logger.ErrorCtx(ctx, fmt.Errorf("failed authorize user"))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
