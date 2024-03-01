package server

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
)

func DatabaseConnect(ctx context.Context, cbleConfig *config.Config) (*ent.Client, error) {
	pgPort := 5432
	if cbleConfig.Database.Port != nil {
		pgPort = *cbleConfig.Database.Port
	}
	pgDatabase := "cble"
	if cbleConfig.Database.Database != nil {
		pgDatabase = *cbleConfig.Database.Database
	}
	pgSslMode := "disable"
	if cbleConfig.Database.SSL != nil && *cbleConfig.Database.SSL {
		pgSslMode = "require"
	}
	pgConnStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cbleConfig.Database.Host,
		pgPort,
		cbleConfig.Database.Username,
		pgDatabase,
		cbleConfig.Database.Password,
		pgSslMode,
	)
	client, err := ent.Open(dialect.Postgres, pgConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return client, nil
}
