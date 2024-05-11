package gophermartprovider

import (
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

type GophermartProvider interface {
	BeginTransaction() (pgclient.Transaction, error)
	RollbackTransaction(tx pgclient.Transaction, log logger.LogClient)
	CommitTransaction(tx pgclient.Transaction) error
}
