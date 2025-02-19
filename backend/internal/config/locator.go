package config

import (
	"context"
	"fmt"
	hasql "golang.yandex/hasql/sqlx"
	"study-chat/pkg/locator"
	"study-chat/pkg/postgres"
	"study-chat/pkg/validator"
)

const (
	ConfigServiceKey    = "config"
	ValidatorServiceKey = "validator"
	ClusterServiceKey   = "cluster"
)

func InitLocator(cfg Config) (locator.ServiceLocator, error) {
	serviceLoc := locator.NewLocator()
	validatorCmp := validator.NewRuValidator()

	cluster, err := initCluster(cfg)
	if err != nil {
		return nil, err
	}

	serviceLoc.Add(ValidatorServiceKey, validatorCmp)
	serviceLoc.Add(ConfigServiceKey, cfg)
	serviceLoc.Add(ClusterServiceKey, cluster)

	return serviceLoc, nil
}

func initCluster(cfg Config) (*hasql.Cluster, error) {
	connData, err := postgres.NewConnectionData(
		cfg.Postgres.Hosts,
		cfg.Postgres.Database,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Port,
		cfg.Postgres.SSL,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to init postgres connection data: %w", err)
	}
	cluster, err := postgres.InitCluster(context.Background(), connData)
	if err != nil {
		return nil, fmt.Errorf("failed to init postgres cluster: %w", err)
	}

	return cluster, nil
}
