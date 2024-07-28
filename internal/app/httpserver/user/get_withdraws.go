package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"
)

func (r *userRouter) GetWithdraws(ctx *gin.Context) {
	userID := ctx.GetString(constants.ContextUserID)

	response, err := r.userUseCase.GetWithdraws(ctx, userID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
