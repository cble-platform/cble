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
	DeployBlueprint(ctx context.Context, client *ent.Client, entRequester *ent.User, entBlueprint *ent.Blueprint, templateVars map[string]interface{}) error
}

const (
	DeploySUCCEEDED  = 0
	DeployINPROGRESS = 1
	DeployFAILED     = 2
)
