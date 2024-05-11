package gophermartprovider

import (
	"context"

	"github.com/Froctnow/yandex-go-diploma/internal/app/gophermartprovider/models"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

type GophermartProvider interface {
	BeginTransaction() (pgclient.Transaction, error)
	RollbackTransaction(tx pgclient.Transaction, log logger.LogClient)
	CommitTransaction(tx pgclient.Transaction) error

	// CreateUser - создает нового пользователя в базе данных и возвращает его
	CreateUser(ctx context.Context, tx pgclient.Transaction, login string, password string) (models.User, error)
	// GetUserForLogin - возвращает пользователя по логину и паролю
	GetUserForLogin(ctx context.Context, tx pgclient.Transaction, login string) (models.User, error)
}
