package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"
	userusecaseerrors "github.com/Froctnow/yandex-go-diploma/internal/app/usecase/user/errors"
)

func (r *userRouter) Withdraw(ctx *gin.Context) {
	var req httpmodels.WithdrawRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	errs := r.validator.UserWithdraw(req.Order)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	userID := ctx.GetString(constants.ContextUserID)

	err := r.userUseCase.Withdraw(ctx, userID, req.Sum, req.Order)

	if errors.As(err, &userusecaseerrors.UserNotEnoughBalance{}) {
		ctx.AbortWithStatus(http.StatusPaymentRequired)
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
