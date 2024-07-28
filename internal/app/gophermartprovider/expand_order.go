package gophermartprovider

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

func (p *GophermartDBProvider) ExpandOrder(
	ctx context.Context,
	tx pgclient.Transaction,
	status string,
	accrual *uint32,
	userID string,
) error {
	_, err := p.conn.Exec(
		ctx,
		"ExpandOrder",
		nil,
		tx,
		status,
		accrual,
		userID,
	)
	if err != nil {
		return fmt.Errorf("can't execute ExpandOrder: %w", err)
	}

	return nil
}
