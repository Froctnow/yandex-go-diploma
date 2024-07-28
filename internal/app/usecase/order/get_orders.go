package order

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	providermodels "github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider/models"
	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/order/models"
)

func (u *orderUseCase) GetOrders(ctx context.Context, userID string) ([]models.Order, error) {
	orders, err := u.provider.GetOrders(ctx, nil, userID)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't get orders: %w", err))
		return nil, err
	}

	u.logger.InfoCtx(ctx, "orders from database", "orders", orders)

	errs, errGroupCtx := errgroup.WithContext(ctx)

	resultOrders := make([]models.Order, 0, len(orders))

	for _, order := range orders {
		order := order
		errs.Go(func() error {
			result, err := u.expandOrder(errGroupCtx, order, userID)

			if err != nil {
				return err
			}

			resultOrders = append(resultOrders, result)

			return nil
		})
	}

	err = errs.Wait()

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't expand orders: %w", err))
		return nil, err
	}

	return resultOrders, nil
}

// Добавить в order информацию из сервиса по accrual
func (u *orderUseCase) expandOrder(ctx context.Context, order providermodels.Order, userID string) (models.Order, error) {
	accrualOrderInfo, err := u.accrualService.GetOrder(ctx, order.Number)
	expandedOrder := models.Order{
		Number:     order.Number,
		Status:     order.Status,
		Accrual:    order.Accrual,
		UploadedAt: order.UploadedAt,
	}

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't accrual: %w", err))
		return models.Order{}, nil
	}

	if accrualOrderInfo == nil {
		u.logger.InfoCtx(ctx, "no accrual info", "order", order)
		return expandedOrder, nil
	}

	if order.Status == accrualOrderInfo.Status {
		u.logger.InfoCtx(ctx, "order already expanded", "order", order)
		return expandedOrder, nil
	}

	expandedOrder.Accrual = accrualOrderInfo.Accrual
	expandedOrder.Status = accrualOrderInfo.Status

	err = u.provider.ExpandOrder(ctx, nil, order.Status, order.Accrual, userID)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't expand order: %w", err))
		return models.Order{}, err
	}

	u.logger.InfoCtx(ctx, "order expanded", "order", expandedOrder)

	return expandedOrder, nil
}
