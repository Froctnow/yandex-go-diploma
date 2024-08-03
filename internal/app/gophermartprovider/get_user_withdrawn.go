package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) GetUserWithdrawn(
	ctx context.Context,
	tx pgclient.Transaction,
	userID string,
) (float32, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"GetUserWithdrawn",
		nil,
		tx,
		userID,
	)
	if err != nil {
		return 0, fmt.Errorf("can't execute GetUserWithdrawn: %w", err)
	}

	err = rows.Err()
	if err != nil {
		return 0, fmt.Errorf("can't execute GetUserWithdrawn: %w", err)
	}

	return pgclient.ValueFromRows[float32](rows)
}
