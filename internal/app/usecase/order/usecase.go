package order

import (
	"context"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

type orderUseCase struct {
	logger   logger.LogClient
	provider gophermartprovider.GophermartProvider
}

func NewUseCase(
	logger logger.LogClient,
	provider gophermartprovider.GophermartProvider,
) UseCase {
	return &orderUseCase{
		logger:   logger,
		provider: provider,
	}
}

type UseCase interface {
	Create(ctx context.Context, orderNumber string, userID string) error
}
