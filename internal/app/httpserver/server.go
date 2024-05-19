package httpserver

import (
	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/middleware"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/user"
	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/auth"
	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/order"
	"github.com/Froctnow/yandex-go-diploma/internal/app/validator"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
	"github.com/gin-gonic/gin"
)

type GophermartServer interface {
}

type gophermartServer struct {
	userRouter user.Router
}

func NewGophermartServer(
	ginEngine *gin.Engine,
	logger logger.LogClient,
	validator validator.Validator,
	cfg *config.Values,
	authUseCase auth.UseCase,
	orderUseCase order.UseCase,
) GophermartServer {
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/api")
	apiGroup.Use(middleware.LoggingMiddleware(logger))
	apiGroup.Use(middleware.DecompressMiddleware(logger))
	apiGroup.Use(middleware.CompressMiddleware())

	return &gophermartServer{
		user.NewRouter(apiGroup, authUseCase, validator, cfg, logger, orderUseCase),
	}
}
