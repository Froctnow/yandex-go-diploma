package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider/models"
	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

// CreateUser - создает нового пользователя в базе данных и возвращает его
func (p *GophermartDBProvider) CreateUser(
	ctx context.Context,
	tx pgclient.Transaction,
	login string,
	password string,
) (models.User, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"CreateUser",
		nil,
		tx,
		login,
		password,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("can't execute CreateUser: %w", err)
	}

	err = rows.Err()

	if err != nil {
		return models.User{}, fmt.Errorf("can't execute CreateUser: %w", err)
	}

	return pgclient.StructValueFromRows[models.User](rows)
}
