package logger

import (
	"github.com/Froctnow/yandex-go-diploma/pkg/logger/formatter"

	"github.com/sirupsen/logrus"
)

type ConsoleOptions struct {
	Level string
}

func NewConsole(opts ConsoleOptions) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&formatter.JSONFormatter{})
	level, err := logrus.ParseLevel(opts.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)

	return logger, nil
}
