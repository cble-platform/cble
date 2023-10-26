package providers

import (
	"context"

	"github.com/cble-platform/backend/ent"
)

type Provider interface {
	ID() string
	Name() string
	Description() string
	Author() string
	Version() string
	DeployBlueprint(ctx context.Context, entBlueprint *ent.Blueprint) error
}
