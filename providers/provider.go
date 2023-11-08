package providers

import (
	"context"

	"github.com/cble-platform/cble-backend/ent"
)

type Provider interface {
	ID() string
	Name() string
	Description() string
	Author() string
	Version() string
	DeployBlueprint(ctx context.Context, client *ent.Client, entRequester *ent.User, entBlueprint *ent.Blueprint, templateVars map[string]interface{}) (*ent.Deployment, error)
	DestroyBlueprint(ctx context.Context, client *ent.Client, entRequester *ent.User, entDeployment *ent.Deployment, templateVars map[string]interface{}) (*ent.Deployment, error)
}

type DeploymentState int

const (
	DeployFAILED     DeploymentState = -1
	DeploySUCCEEDED  DeploymentState = 0
	DeployINPROGRESS DeploymentState = 1
	DeployDESTROYED  DeploymentState = 2
)
