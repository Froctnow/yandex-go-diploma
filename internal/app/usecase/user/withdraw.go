package user

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/usecase/user/errors"
)

func (u *userUseCase) Withdraw(ctx context.Context, userID string, sum float32, orderNumber string) error {
	tx, err := u.provider.BeginTransaction()

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't begin transaction: %w", err))
		return err
	}

	balance, err := u.provider.GetUserBalanceForUpdate(ctx, tx, userID)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't get user balance: %w", err))
		u.provider.RollbackTransaction(tx, u.logger)
		return err
	}

	if balance < sum {
		u.logger.ErrorCtx(ctx, fmt.Errorf("not enough money on balance"))
		u.provider.RollbackTransaction(tx, u.logger)
		return errors.UserNotEnoughBalance{}
	}

	err = u.provider.CreateWithdrawTransaction(ctx, tx, userID, sum, orderNumber)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't create withdraw transaction: %w", err))
		u.provider.RollbackTransaction(tx, u.logger)
		return err
	}

	err = u.provider.UpdateUserBalance(ctx, tx, userID, balance-sum)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't update user balance: %w", err))
		u.provider.RollbackTransaction(tx, u.logger)
		return err
	}

	err = u.provider.CommitTransaction(tx)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't commit transaction: %w", err))

		u.provider.RollbackTransaction(tx, u.logger)

		return err
	}

	return nil
}
