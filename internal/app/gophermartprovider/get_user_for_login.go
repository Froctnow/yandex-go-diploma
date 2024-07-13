package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider/models"
	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) GetUserForLogin(
	ctx context.Context,
	tx pgclient.Transaction,
	login string,
) (models.User, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"GetUserForLogin",
		nil,
		tx,
		login,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("can't execute GetUserForLogin: %w", err)
	}

	err = rows.Err()
	if err != nil {
		return models.User{}, fmt.Errorf("can't execute GetUserForLogin: %w", err)
	}

	return pgclient.StructValueFromRows[models.User](rows)
}
