package order

import (
	"context"

	"github.com/Froctnow/yandex-go-diploma/internal/app/accrualservice"
	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider"
	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/order/models"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

type orderUseCase struct {
	logger         logger.LogClient
	provider       gophermartprovider.GophermartProvider
	accrualService accrualservice.Service
}

func NewUseCase(
	logger logger.LogClient,
	provider gophermartprovider.GophermartProvider,
	accrualsService accrualservice.Service,
) UseCase {
	return &orderUseCase{
		logger:         logger,
		provider:       provider,
		accrualService: accrualsService,
	}
}

type UseCase interface {
	Create(ctx context.Context, orderNumber string, userID string) error
	GetOrders(ctx context.Context, userID string) ([]models.Order, error)
}
