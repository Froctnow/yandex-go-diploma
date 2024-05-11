package log

import (
	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

func New(cfg config.Values) (logger.LogClient, error) {
	log, err := logger.New(logger.Options{
		ConsoleOptions: logger.ConsoleOptions{
			Level: cfg.LogLevel,
		},
	})
	if err != nil {
		return nil, err
	}

	return log, err
}
