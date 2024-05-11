package migration

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Froctnow/yandex-go-diploma/internal/app/config"
	"github.com/Froctnow/yandex-go-diploma/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ExecuteMigrations(cfg *config.Values, log logger.LogClient) error {
	db, err := sql.Open("postgres", cfg.DatabaseURI)

	if err != nil {
		return fmt.Errorf("failed connect to database for migrations %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("failed create postgres instance for migrate %w", err)
	}

	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/app/migration/migrations",
		"postgres", driver)

	if err != nil {
		return fmt.Errorf("failed create migration instance %w", err)
	}

	currentVersion, err := getMigrationsVersion(m)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Current version database is %d", currentVersion))

	err = m.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed up mirationgs %w", err)
	}

	newVersion, err := getMigrationsVersion(m)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Current version database is %d", newVersion))

	return nil
}

func getMigrationsVersion(m *migrate.Migrate) (uint, error) {
	version, dirty, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		return 0, fmt.Errorf("failed get version of migrations %w", err)
	}

	if dirty {
		return 0, fmt.Errorf("database is dirty")
	}

	return version, nil
}
