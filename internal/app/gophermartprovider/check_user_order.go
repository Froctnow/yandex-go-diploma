package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) CheckUserOrder(
	ctx context.Context,
	tx pgclient.Transaction,
	orderNumber string,
	userID string,
) (bool, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"CheckUserOrder",
		nil,
		tx,
		orderNumber,
		userID,
	)
	if err != nil {
		return false, fmt.Errorf("can't execute CheckUserOrder: %w", err)
	}

	err = rows.Err()

	if err != nil {
		return false, fmt.Errorf("can't execute CheckUserOrder: %w", err)
	}

	return pgclient.ValueFromRows[bool](rows)
}
