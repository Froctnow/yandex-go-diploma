package user

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	httpmodels "github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/models"
	ordererrors "github.com/Froctnow/yandex-go-diploma/internal/app/usecase/order/errors"
	"github.com/gin-gonic/gin"

	"github.com/pkg/errors"
)

func (r *userRouter) CreateOrder(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := checkHeaderContentType(headerContentType)

	if !isCorrectHeaderContentType {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "Something went wrong"})
		return
	}

	orderNumber := string(body)

	errs := r.validator.UserCreateOrder(orderNumber)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	userID := ctx.GetString(constants.ContextUserID)
	err = r.orderUseCase.Create(ctx, orderNumber, userID)

	if err != nil && errors.As(err, &ordererrors.UserOrderAlreadyExists{}) {
		ctx.AbortWithStatus(http.StatusOK)
		return
	}

	if err != nil && errors.As(err, &ordererrors.OrderAlreadyExists{}) {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	if err != nil {
		r.logger.ErrorCtx(ctx, fmt.Errorf("failed create order"))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func checkHeaderContentType(value string) bool {
	isTextPlain := strings.Contains(value, "text/plain")

	return isTextPlain
}
