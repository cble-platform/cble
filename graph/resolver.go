package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/graph/generated"
	"github.com/cble-platform/cble-backend/providers"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	ent        *ent.Client
	cbleServer *providers.CBLEServer
	// redis *redis.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, cbleServer *providers.CBLEServer) graphql.ExecutableSchema {
	c := generated.Config{
		Resolvers: &Resolver{
			ent:        client,
			cbleServer: cbleServer,
			// rdb:           rdb,
		},
	}
	return generated.NewExecutableSchema(c)
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
