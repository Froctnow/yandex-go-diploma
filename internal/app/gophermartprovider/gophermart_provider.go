package gophermartprovider

import (
	"fmt"
	"reflect"

	"github.com/Froctnow/yandex-go-diploma/pkg/logger"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

type GophermartDBProvider struct {
	conn pgclient.PGClient
}

func NewGophermartProvider(dbConn pgclient.PGClient) GophermartProvider {
	return &GophermartDBProvider{
		conn: dbConn,
	}
}

func (p *GophermartDBProvider) BeginTransaction() (pgclient.Transaction, error) {
	return p.conn.BeginTransaction()
}

func (p *GophermartDBProvider) RollbackTransaction(tx pgclient.Transaction, log logger.LogClient) {
	if tx == nil || reflect.ValueOf(tx).IsNil() {
		return
	}
	txErr := tx.Rollback()
	if txErr != nil {
		log.Error(txErr)
	}
}

func (p *GophermartDBProvider) CommitTransaction(tx pgclient.Transaction) error {
	if tx == nil || reflect.ValueOf(tx).IsNil() {
		return fmt.Errorf("nil transaction pointer in CommitTransaction")
	}
	return tx.Commit()
}
