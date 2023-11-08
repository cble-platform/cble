package v1

import (
	"context"
	"sync"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

func (provider OpenstackProvider) newAuthClient() (*gophercloud.ProviderClient, error) {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: provider.config.AuthUrl,
		Username:         provider.config.Username,
		Password:         provider.config.Password,
		TenantID:         provider.config.ProjectID,
		TenantName:       provider.config.ProjectName,
	}
	if provider.config.DomainName != "" {
		authOpts.DomainName = provider.config.DomainName
	} else {
		authOpts.DomainID = provider.config.DomainId
	}
	return openstack.AuthenticatedClient(authOpts)
}

func loadState(entDeployment *ent.Deployment, stateMap *sync.Map) {
	for k, v := range entDeployment.DeploymentState {
		stateMap.Store(k, providers.DeploymentState(v))
	}
}

func loadVars(entDeployment *ent.Deployment, varMap *sync.Map) {
	for k, v := range entDeployment.DeploymentVars {
		varMap.Store(k, v)
	}
}

func saveState(ctx context.Context, entDeployment *ent.Deployment, stateMap *sync.Map) error {
	newState := make(map[string]int)
	stateMap.Range(func(key, value any) bool {
		newState[key.(string)] = int(value.(providers.DeploymentState))
		return true
	})
	return entDeployment.Update().SetDeploymentState(newState).Exec(ctx)
}

func saveVars(ctx context.Context, entDeployment *ent.Deployment, varMap *sync.Map) error {
	newVars := make(map[string]interface{})
	varMap.Range(func(key, value any) bool {
		newVars[key.(string)] = value
		return true
	})
	return entDeployment.Update().SetDeploymentVars(newVars).Exec(ctx)
}

func setDeploymentState(ctx context.Context, entDeployment *ent.Deployment, stateMap *sync.Map, key string, state providers.DeploymentState) error {
	stateMap.Store(key, state)
	return saveState(ctx, entDeployment, stateMap)
}

func saveDeploymentVar(ctx context.Context, entDeployment *ent.Deployment, varMap *sync.Map, varName string, value interface{}) error {
	varMap.Store(varName, value)
	return saveVars(ctx, entDeployment, varMap)
}

func deleteDeploymentVar(ctx context.Context, entDeployment *ent.Deployment, varMap *sync.Map, varName string) error {
	varMap.Delete(varName)
	return saveVars(ctx, entDeployment, varMap)
}
