package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider/models"
	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) GetUserWithdraws(
	ctx context.Context,
	tx pgclient.Transaction,
	userID string,
) ([]models.Withdraw, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"GetUserWithdraws",
		nil,
		tx,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("can't execute GetUserWithdraws: %w", err)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("can't execute GetUserWithdraws: %w", err)
	}

	return pgclient.ListValuesFromRows[models.Withdraw](rows)
}
