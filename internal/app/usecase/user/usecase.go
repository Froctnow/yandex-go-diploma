package user

import (
	"context"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider"
	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/user/models"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

type userUseCase struct {
	logger   logger.LogClient
	provider gophermartprovider.GophermartProvider
}

func NewUseCase(
	logger logger.LogClient,
	provider gophermartprovider.GophermartProvider,
) UseCase {
	return &userUseCase{
		logger:   logger,
		provider: provider,
	}
}

type UseCase interface {
	GetBalance(ctx context.Context, userID string) (models.UseBalance, error)
	Withdraw(ctx context.Context, userID string, sum float32, orderNumber string) error
	GetWithdraws(ctx context.Context, userID string) ([]models.UserWithdraw, error)
}
