package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) IncreaseUserBalance(
	ctx context.Context,
	tx pgclient.Transaction,
	userID string,
	sum float32,
) error {
	_, err := p.conn.Exec(
		ctx,
		"IncreaseUserBalance",
		nil,
		tx,
		sum,
		userID,
	)
	if err != nil {
		return fmt.Errorf("can't execute IncreaseUserBalance: %w", err)
	}

	return nil
}
