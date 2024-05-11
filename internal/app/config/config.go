package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Values struct {
	Address         string `env:"SERVER_ADDRESS" envSeparator:":"`
	Hostname        string `env:"BASE_URL" envSeparator:":"`
	LogLevel        string `env:"LOG_LEVEL" envSeparator:":"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envSeparator:":"`
	DatabaseDSN     string `env:"DATABASE_DSN" envSeparator:":"`
	StorageMode     string
	JwtSecret       string
	JwtTokenExpire  time.Duration
}

func NewConfig(isUseFlags bool) (*Values, error) {
	var cfg Values
	var address = new(string)
	var hostname = new(string)
	var logLevel = new(string)
	var fileStoragePath = new(string)
	var databaseDSN = new(string)

	err := env.Parse(&cfg)

	if err != nil {
		panic(fmt.Errorf("can't parse env %w", err))
	}

	if isUseFlags {
		address = flag.String("a", "", "address of service")
		hostname = flag.String("b", "", "hostname of service")
		logLevel = flag.String("loglevel", "", "level of logs")
		fileStoragePath = flag.String("f", "", "file path to the storage file")
		databaseDSN = flag.String("d", "", "database DSN")
		// разбор командной строки
		flag.Parse()
	}

	if cfg.Address == "" {
		if *address == "" {
			*address = fmt.Sprintf(`:%d`, 8080)
		}

		cfg.Address = *address
	}
	if cfg.Hostname == "" {
		if *hostname == "" {
			*hostname = "http://localhost:8080"
		}

		cfg.Hostname = *hostname
	}

	if cfg.LogLevel == "" {
		if *logLevel == "" {
			*logLevel = "info"
		}

		cfg.LogLevel = *logLevel
	}

	if cfg.FileStoragePath == "" {
		if *fileStoragePath == "" {
			*fileStoragePath = ""
		}

		cfg.FileStoragePath = *fileStoragePath
	}

	if cfg.DatabaseDSN == "" {
		if *databaseDSN == "" {
			*databaseDSN = ""
		}

		cfg.DatabaseDSN = *databaseDSN
	}

	cfg.JwtSecret = "supersecretkey"
	cfg.JwtTokenExpire = time.Hour * 3

	fmt.Println(cfg.LogLevel)

	return &cfg, nil
}
