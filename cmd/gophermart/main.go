package main

import (
	"context"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/internal/app/log"
	"github.com/Froctnow/yandex-go-diploma/internal/bootstrap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.NewConfig(true)

	if err != nil {
		panic(fmt.Errorf("config read err %w", err))
	}

	logger, err := log.New(*cfg)

	if err != nil {
		panic(err)
	}

	bootstrap.RunApp(ctx, cfg, logger)
}
