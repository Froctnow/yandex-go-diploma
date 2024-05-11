package auth

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/auth/errors"
)

func (u *authUseCase) Login(ctx context.Context, login string, password string) error {
	user, err := u.provider.GetUserForLogin(ctx, nil, login)

	if err != nil {
		err = fmt.Errorf("can't get user for login: %w", err)
		u.logger.ErrorCtx(ctx, err)
		return err
	}

	if user.ID == "" {
		err = errors.IncorrectLoginPasswordError{}
		u.logger.ErrorCtx(ctx, err)
		return err
	}

	if !u.checkPasswordHash(password, user.Password) {
		err = errors.IncorrectLoginPasswordError{}
		u.logger.ErrorCtx(ctx, err)
		return err
	}

	return nil
}
