package internal

import (
	"context"
	"errors"
	"fmt"
	"study-chat/internal/config"

	"study-chat/pkg/logger"
	"study-chat/pkg/sentry"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"study-chat/pkg/postgres"
)

func Migrate(cfg config.Config) error {
	if err := sentry.Init(cfg.Sentry.DSN, cfg.Sentry.Environment); err != nil {
		return fmt.Errorf("failed to init sentry: %w", err)
	}
	logger.Setup()

	connData, err := postgres.NewConnectionData(
		cfg.Postgres.Hosts,
		cfg.Postgres.Database,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Port,
		cfg.Postgres.SSL,
	)
	if err != nil {
		return fmt.Errorf("failed to init postgres connection data: %w", err)
	}
	cluster, err := postgres.InitCluster(context.Background(), connData)
	if err != nil {
		return fmt.Errorf("failed to init postgres cluster: %w", err)
	}

	masterHost := cluster.Primary().Addr()
	connURL := connData.URL(masterHost)

	m, err := migrate.New(
		cfg.Postgres.MigrationPath,
		connURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
