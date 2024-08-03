package user

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/user/models"
)

func (u *userUseCase) GetWithdraws(ctx context.Context, userID string) ([]models.UserWithdraw, error) {
	withdraws, err := u.provider.GetUserWithdraws(ctx, nil, userID)
	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't get user balance: %w", err))
		return nil, err
	}

	result := make([]models.UserWithdraw, 0, len(withdraws))

	for _, withdraw := range withdraws {
		result = append(result, models.UserWithdraw{
			OrderNumber: withdraw.OrderNumber,
			Sum:         withdraw.Sum,
			ProcessedAt: withdraw.ProcessedAt,
		})
	}

	return result, nil
}
