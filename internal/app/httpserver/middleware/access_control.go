package middleware

import (
	"fmt"
	"net/http"

	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func AccessControlMiddleware(cfg *config.Values, logger logger.LogClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("jwt")

		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			logger.ErrorCtx(c, fmt.Errorf("can't get jwt token from cookie: %w", err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		decodedJwtToken, err := decodeJwtToken(jwtToken, cfg.JwtSecret)

		if err != nil {
			logger.ErrorCtx(c, fmt.Errorf("can't decode jwt token: %w", err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if decodedJwtToken.UserID == "" {
			logger.WarnCtx(c, "user_id is absent in jwt")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(constants.ContextUserID, decodedJwtToken.UserID)

		c.Next()
	}
}

func decodeJwtToken(jwtToken string, jwtSecret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("can't parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return token.Claims.(*Claims), nil
}
