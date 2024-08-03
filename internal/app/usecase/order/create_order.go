package order

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/order/errors"
)

func (u *orderUseCase) Create(ctx context.Context, orderNumber string, userID string) error {
	err := u.checkUserOrder(ctx, orderNumber, userID)
	if err != nil {
		u.logger.ErrorCtx(ctx, err)
		return err
	}

	insertedOrderNumber, err := u.provider.CreateOrder(ctx, nil, orderNumber, userID)
	if err != nil {
		err = fmt.Errorf("failed to create order: %w", err)
		u.logger.ErrorCtx(ctx, err)
		return err
	}

	if insertedOrderNumber == "" {
		err = u.checkUserOrder(ctx, orderNumber, userID)
		if err != nil {
			u.logger.ErrorCtx(ctx, err)
			return err
		}

		return errors.OrderAlreadyExists{}
	}

	return nil
}

func (u *orderUseCase) checkUserOrder(ctx context.Context, orderNumber string, userID string) error {
	isUserOrderExists, err := u.provider.CheckUserOrder(ctx, nil, orderNumber, userID)
	if err != nil {
		return fmt.Errorf("failed to check user order: %w", err)
	}

	if isUserOrderExists {
		return errors.UserOrderAlreadyExists{}
	}

	return nil
}
