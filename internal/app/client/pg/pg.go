package pg

import (
	"embed"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
	"github.com/Froctnow/yandex-go-diploma/pkg/pgclient"
)

var (
	//go:embed queries/*
	queryFiles embed.FS

	pathsToDBQueries = []string{"queries/"}
)

//go:generate mockery --srcpkg=vcs.bingo-boom.ru/bb_online/go-modules/pgclient --case=underscore --name=Transaction

func New(cfg *config.Values, log logger.LogClient) (pgclient.PGClient, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid pg config")
	}
	if cfg.DatabaseDSN == "" {
		return nil, nil
	}
	connString := cfg.DatabaseDSN
	configValues := pgclient.PostgreSQL{
		ConnString:     connString,
		PathsToQueries: pathsToDBQueries,
		LogLevel:       pgclient.LogLevelNone,
	}
	return pgclient.New(configValues, log, queryFiles)
}
