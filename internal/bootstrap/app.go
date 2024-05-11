package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/internal/app/httpserver"
	"github.com/Froctnow/yandex-go-diploma/internal/app/validator"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
)

func RunApp(ctx context.Context, cfg *config.Values, logger logger.LogClient) {
	ginEngine := NewGinEngine()
	httpServer, err := RunHTTPServer(ginEngine, cfg)
	if err != nil {
		panic(fmt.Errorf("http server can't start %w", err))
	}

	validatorInstance := validator.New()

	_ = httpserver.NewGophermartServer(
		ginEngine,
		logger,
		validatorInstance,
		cfg,
	)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	fmt.Println("app is ready")
	select {
	case v := <-exit:
		fmt.Printf("signal.Notify: %v\n\n", v)
	case done := <-ctx.Done():
		fmt.Println(fmt.Errorf("ctx.Done: %v", done))
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server Exited Properly")
}
