package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) UpdateUserBalance(
	ctx context.Context,
	tx pgclient.Transaction,
	userID string,
	sum float32,
) error {
	_, err := p.conn.NamedQueryxContext(
		ctx,
		"UpdateUserBalance",
		nil,
		tx,
		sum,
		userID,
	)
	if err != nil {
		return fmt.Errorf("can't execute UpdateUserBalance: %w", err)
	}

	return nil
}
