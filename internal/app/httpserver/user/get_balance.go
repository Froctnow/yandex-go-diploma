package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"
)

func (r *userRouter) GetBalance(ctx *gin.Context) {
	userID := ctx.GetString(constants.ContextUserID)

	userBalance, err := r.userUseCase.GetBalance(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	response := httpmodels.GetUserBalanceResponse{
		Current:   userBalance.Balance,
		Withdrawn: userBalance.Withdrawn,
	}

	ctx.JSON(http.StatusOK, response)
}
