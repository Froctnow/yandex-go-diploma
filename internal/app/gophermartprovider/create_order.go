package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

// CreateOrder - создает новый заказ в базе данных и возвращает его номер
func (p *GophermartDBProvider) CreateOrder(
	ctx context.Context,
	tx pgclient.Transaction,
	orderNumber string,
	userID string,
) (string, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"CreateOrder",
		nil,
		tx,
		orderNumber,
		userID,
	)
	if err != nil {
		return "", fmt.Errorf("can't execute CreateOrder: %w", err)
	}

	err = rows.Err()

	if err != nil {
		return "", fmt.Errorf("can't execute CreateOrder: %w", err)
	}

	return pgclient.ValueFromRows[string](rows)
}
