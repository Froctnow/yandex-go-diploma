package user

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/middleware"
	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/auth"
	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/order"
	"github.com/Froctnow/yandex-go-diploma/internal/app/validator"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

type Router interface {
	Register(c *gin.Context)
}

type userRouter struct {
	authUseCase  auth.UseCase
	validator    validator.Validator
	cfg          *config.Values
	logger       logger.LogClient
	orderUseCase order.UseCase
}

func NewRouter(
	ginGroup *gin.RouterGroup,
	authUseCase auth.UseCase,
	validator validator.Validator,
	cfg *config.Values,
	logger logger.LogClient,
	orderUseCase order.UseCase,
) Router {
	router := &userRouter{
		authUseCase:  authUseCase,
		validator:    validator,
		cfg:          cfg,
		logger:       logger,
		orderUseCase: orderUseCase,
	}

	userGroup := ginGroup.Group("/user")
	userGroup.POST("/register", router.Register)
	userGroup.POST("/login", router.Login)
	userGroup.POST("/orders", middleware.AccessControlMiddleware(cfg, logger), router.CreateOrder)
	userGroup.GET("/orders", middleware.AccessControlMiddleware(cfg, logger), router.GetOrders)

	return router
}

func (r *userRouter) checkHeaderContentType(value string) bool {
	isTextPlain := strings.Contains(value, "application/json")
	isXGzip := strings.Contains(value, "application/x-gzip")

	return isTextPlain || isXGzip
}

// authorizeUser авторизует пользователя и устанавливает куки
func (r *userRouter) authorizeUser(c *gin.Context, userID string) error {
	token, err := buildJWTString(r.cfg.JwtSecret, r.cfg.JwtTokenExpire, userID)
	if err != nil {
		r.logger.ErrorCtx(c, fmt.Errorf("can't build jwt token: %w", err))
		return err
	}

	r.logger.DebugCtx(c, "jwt token", token)

	c.SetCookie("jwt", token, int(r.cfg.JwtTokenExpire.Seconds()), "/", "", false, true)
	c.Set(constants.ContextUserID, userID)

	return nil
}

// buildJWTString создаёт строку токена
func buildJWTString(jwtSecret string, jwtTokenExpire time.Duration, userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtTokenExpire)),
		},
		UserID: userID,
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("can't sign token: %w", err)
	}

	// возвращаем строку токена
	return tokenString, nil
}
