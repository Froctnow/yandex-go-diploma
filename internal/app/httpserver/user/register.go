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

func (r *userRouter) Register(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := r.checkHeaderContentType(headerContentType)

	if !isCorrectHeaderContentType {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	var req httpmodels.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	errs := r.validator.UserRegister(&req)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	err := r.authUseCase.Register(ctx, req.Login, req.Password)

	if err != nil && errors.As(err, &autherrors.UserAlreadyExistsError{}) {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	if err != nil {
		r.logger.ErrorCtx(ctx, fmt.Errorf("failed register"))
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
