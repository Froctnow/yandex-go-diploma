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
	// GetOrders - берем все заказы пользователя из нашей базы данных
	GetOrders(ctx context.Context, tx pgclient.Transaction, userID string) ([]models.Order, error)
	// ExpandOrder - добавить информацию о заказе из accrual сервиса
	ExpandOrder(ctx context.Context, tx pgclient.Transaction, status string, accrual *float32, userID string) error
	// GetUserBalance - возвращает баланс пользователя
	GetUserBalance(ctx context.Context, tx pgclient.Transaction, userID string) (float32, error)
	// GetUserWithdrawn - возвращает сумму потраченных баллов пользователя
	GetUserWithdrawn(ctx context.Context, tx pgclient.Transaction, userID string) (float32, error)
	// GetUserBalanceForUpdate - возвращает баланс пользователя для обновления
	GetUserBalanceForUpdate(ctx context.Context, tx pgclient.Transaction, userID string) (float32, error)
	// CreateWithdrawTransaction - создает транзакцию списания средств
	CreateWithdrawTransaction(ctx context.Context, tx pgclient.Transaction, userID string, sum float32, orderNumber string) error
	// UpdateUserBalance - обновляет баланс пользователя
	UpdateUserBalance(ctx context.Context, tx pgclient.Transaction, userID string, sum float32) error
	// GetUserWithdraws - возвращает все списания пользователя
	GetUserWithdraws(ctx context.Context, tx pgclient.Transaction, userID string) ([]models.Withdraw, error)
	// IncreaseUserBalance - увеличивает баланс пользователя
	IncreaseUserBalance(ctx context.Context, tx pgclient.Transaction, userID string, sum float32) error
}
