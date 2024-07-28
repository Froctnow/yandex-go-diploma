package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) GetUserBalance(
	ctx context.Context,
	tx pgclient.Transaction,
	userID string,
) (float32, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"GetUserBalance",
		nil,
		tx,
		userID,
	)
	if err != nil {
		return 0, fmt.Errorf("can't execute GetUserBalance: %w", err)
	}

	err = rows.Err()
	if err != nil {
		return 0, fmt.Errorf("can't execute GetUserBalance: %w", err)
	}

	return pgclient.ValueFromRows[float32](rows)
}
