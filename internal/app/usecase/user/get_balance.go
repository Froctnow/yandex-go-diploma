package user

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/user/models"
)

func (u *userUseCase) GetBalance(ctx context.Context, userID string) (models.UseBalance, error) {
	balance, err := u.provider.GetUserBalance(ctx, nil, userID)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't get user balance: %w", err))
		return models.UseBalance{}, err
	}

	withdrawn, err := u.provider.GetUserWithdrawn(ctx, nil, userID)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't get user withdrawn: %w", err))
		return models.UseBalance{}, err
	}

	return models.UseBalance{
		Balance:   balance,
		Withdrawn: withdrawn,
	}, nil
}
