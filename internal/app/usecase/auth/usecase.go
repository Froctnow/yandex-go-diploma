package auth

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	logger   logger.LogClient
	provider gophermartprovider.GophermartProvider
}

func NewUseCase(
	logger logger.LogClient,
	provider gophermartprovider.GophermartProvider,
) UseCase {
	return &authUseCase{
		logger:   logger,
		provider: provider,
	}
}

type UseCase interface {
	Register(ctx context.Context, login string, password string) (string, error)
	Login(ctx context.Context, login string, password string) (string, error)
}

func (u *authUseCase) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.logger.Debug(fmt.Sprintf("hashing password from %s to %s", password, string(bytes)))
	return string(bytes), err
}

func (u *authUseCase) checkPasswordHash(password string, hash string) bool {
	u.logger.Debug(fmt.Sprintf("checking password %s with %s", password, hash))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
