package httpserver

import (
	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/middleware"
	"github.com/Froctnow/yandex-go-diploma/internal/app/validator"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
	"github.com/gin-gonic/gin"
)

type GophermartServer interface {
}

type gophermartServer struct {
}

func NewGophermartServer(
	ginEngine *gin.Engine,
	logger logger.LogClient,
	validator validator.Validator,
	cfg *config.Values,
) GophermartServer {
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/api")
	apiGroup.Use(middleware.AccessControlMiddleware(cfg, logger))
	apiGroup.Use(middleware.LoggingMiddleware(logger))
	apiGroup.Use(middleware.DecompressMiddleware(logger))
	apiGroup.Use(middleware.CompressMiddleware())

	return &gophermartServer{}
}
