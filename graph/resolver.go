package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cble-platform/cble-backend/auth"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/graph/generated"
	"github.com/cble-platform/cble-backend/internal/permissionengine"
	"github.com/cble-platform/cble-backend/providers"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	ent              *ent.Client
	cbleServer       *providers.CBLEServer
	permissionEngine *permissionengine.PermissionEngine
	// redis *redis.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, cbleServer *providers.CBLEServer, permissionEngine *permissionengine.PermissionEngine) graphql.ExecutableSchema {
	c := generated.Config{
		Resolvers: &Resolver{
			ent:              client,
			cbleServer:       cbleServer,
			permissionEngine: permissionEngine,
			// rdb:           rdb,
		},
	}
	c.Directives.Permission = func(ctx context.Context, obj interface{}, next graphql.Resolver, key string) (interface{}, error) {
		currentUser, err := auth.ForContext(ctx)
		if err != nil {
			return nil, fmt.Errorf("user not authenticated")
		}
		hasPermission, err := permissionEngine.RequestPermission(ctx, currentUser, key)
		if err != nil {
			return nil, err
		}
		if hasPermission {
			return next(ctx)
		}
		return nil, fmt.Errorf("user does not have permission %s", key)
	}
	// c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*model.UserRole) (res interface{}, err error) {
	// 	currentUser, err := auth.ForContext(ctx)
	// 	if err != nil {
	// 		return nil, auth.AUTH_REQUIRED_GQL_ERROR
	// 	}
	// 	for _, role := range roles {
	// 		if role.String() == string(currentUser.Role) {
	// 			return next(ctx)
	// 		}
	// 	}
	// 	return nil, auth.PERMISSION_DENIED_GQL_ERROR
	// }
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
