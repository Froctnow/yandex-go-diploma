package auth

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/auth/errors"
)

func (u *authUseCase) Register(ctx context.Context, login string, password string) (string, error) {
	hashPassword, err := u.hashPassword(password)
	if err != nil {
		err = fmt.Errorf("can't hash password: %w", err)
		u.logger.ErrorCtx(ctx, err)
		return "", err
	}

	userID, err := u.provider.CreateUser(ctx, nil, login, hashPassword)
	if err != nil {
		err = fmt.Errorf("can't create user: %w", err)
		u.logger.ErrorCtx(ctx, err)
		return "", err
	}

	if userID == "" {
		err = errors.UserAlreadyExistsError{}
		u.logger.ErrorCtx(ctx, err)
		return "", err
	}

	return userID, nil
}
