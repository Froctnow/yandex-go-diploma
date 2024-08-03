package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) CreateWithdrawTransaction(
	ctx context.Context,
	tx pgclient.Transaction,
	userID string,
	sum float32,
	orderNumber string,
) error {
	_, err := p.conn.Exec(
		ctx,
		"CreateWithdrawTransaction",
		nil,
		tx,
		userID,
		sum,
		orderNumber,
	)
	if err != nil {
		return fmt.Errorf("can't execute CreateWithdrawTransaction: %w", err)
	}

	return nil
}
