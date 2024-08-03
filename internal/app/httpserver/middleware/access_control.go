package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver/constants"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

type TokenIsInvalid struct{}

func (e TokenIsInvalid) Error() string {
	return "token is invalid"
}

func AccessControlMiddleware(cfg *config.Values, logger logger.LogClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("jwt")

		if err != nil && errors.Is(err, http.ErrNoCookie) {
			logger.WarnCtx(c, "jwt cookie is absent")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if err != nil {
			logger.ErrorCtx(c, fmt.Errorf("can't get jwt token from cookie: %w", err))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		decodedJwtToken, err := decodeJwtToken(jwtToken, cfg.JwtSecret)

		if err != nil && errors.As(err, &TokenIsInvalid{}) {
			logger.WarnCtx(c, "jwt token is invalid")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

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

	if err != nil && errors.Is(err, jwt.ErrSignatureInvalid) {
		return nil, TokenIsInvalid{}
	}

	if err != nil {
		return nil, fmt.Errorf("can't parse token: %w", err)
	}

	if !token.Valid {
		return nil, TokenIsInvalid{}
	}

	return token.Claims.(*Claims), nil
}
