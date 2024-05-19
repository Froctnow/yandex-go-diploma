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

	// CreateUser - создает нового пользователя в базе данных и возвращает его ID
	CreateUser(ctx context.Context, tx pgclient.Transaction, login string, password string) (string, error)
	// GetUserForLogin - возвращает пользователя по логину и паролю
	GetUserForLogin(ctx context.Context, tx pgclient.Transaction, login string) (models.User, error)
	// CreateOrder - создает новый заказ в базе данных и возвращает его номер
	CreateOrder(ctx context.Context, tx pgclient.Transaction, orderNumber string, userID string) (string, error)
	// CheckUserOrder - проверят есть ли в базе номер заказа у пользователя
	CheckUserOrder(ctx context.Context, tx pgclient.Transaction, orderNumber string, userID string) (bool, error)
}
