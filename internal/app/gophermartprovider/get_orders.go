package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider/models"
	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) GetOrders(
	ctx context.Context,
	tx pgclient.Transaction,
	userID string,
) ([]models.Order, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"GetOrders",
		nil,
		tx,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("can't execute GetOrders: %w", err)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("can't execute GetOrders: %w", err)
	}

	return pgclient.ListValuesFromRows[models.Order](rows)
}
