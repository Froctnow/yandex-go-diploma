package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Values struct {
	Address              string `env:"RUN_ADDRESS" envSeparator:":"`
	LogLevel             string `env:"LOG_LEVEL" envSeparator:":"`
	DatabaseURI          string `env:"DATABASE_URI" envSeparator:":"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS" envSeparator:":"`
	JwtSecret            string
	JwtTokenExpire       time.Duration
}

func NewConfig(isUseFlags bool) (*Values, error) {
	var cfg Values
	address := new(string)
	logLevel := new(string)
	databaseURI := new(string)
	accrualSystemAddress := new(string)

	err := env.Parse(&cfg)
	if err != nil {
		panic(fmt.Errorf("can't parse env %w", err))
	}

	if isUseFlags {
		address = flag.String("a", "", "address of service")
		databaseURI = flag.String("d", "", "database URI")
		accrualSystemAddress = flag.String("r", "", "accrual system address")
		// разбор командной строки
		flag.Parse()
	}

	if cfg.Address == "" {
		cfg.Address = *address
	}

	if cfg.LogLevel == "" {
		if *logLevel == "" {
			*logLevel = "info"
		}

		cfg.LogLevel = *logLevel
	}

	if cfg.DatabaseURI == "" {
		cfg.DatabaseURI = *databaseURI
	}

	if cfg.AccrualSystemAddress == "" {
		cfg.AccrualSystemAddress = *accrualSystemAddress
	}

	cfg.JwtSecret = "supersecretkey"
	cfg.JwtTokenExpire = time.Hour * 3

	return &cfg, nil
}
