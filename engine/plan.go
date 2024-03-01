package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/ent/resource"
	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/google/uuid"
)

func CreateDeployment(ctx context.Context, client *ent.Client, entBlueprint *ent.Blueprint, projectID uuid.UUID, templateVars map[string]string, expiryTime time.Time, requester *ent.User) (*ent.Deployment, error) {
	// Get the project from blueprint
	entProject, err := client.Project.Get(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to query project: %v", err)
	}

	// Check we won't exceed the quotas
	newUsages, err := checkProjectQuotas(ctx, entBlueprint, entProject)
	if err != nil {
		return nil, fmt.Errorf("quota exceeded: %v", err)
	}

	// Create the deployment
	entDeployment, err := client.Deployment.Create().
		SetName(entBlueprint.Name).
		SetDescription(entBlueprint.Description).
		SetState(deployment.StateAwaiting).
		SetTemplateVars(templateVars).
		SetExpiresAt(expiryTime).
		SetBlueprint(entBlueprint).
		SetRequester(requester).
		SetProject(entProject).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment: %v", err)
	}

	// Query all "root resources" (resources without dependencies)
	entRootReqsources, err := entBlueprint.QueryResources().Where(
		resource.Not(resource.HasDependsOn()),
	).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query root resources: %v", err)
	}

	// Generate deployment nodes from resources
	nodeMap := make(map[string]*ent.DeploymentNode) // Store created resources for memoization
	for _, entResource := range entRootReqsources {
		_, err := createDeploymentNode(ctx, client, entDeployment, entResource, nodeMap, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create deployment node: %v", err)
		}
	}

	// Update the project quota usage
	err = entProject.Update().
		SetUsageCPU(int(newUsages.Cpu)).
		SetUsageRAM(int(newUsages.Ram)).
		SetUsageDisk(int(newUsages.Disk)).
		SetUsageNetwork(int(newUsages.Network)).
		SetUsageRouter(int(newUsages.Router)).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update project quota usage: %v", err)
	}

	return entDeployment, nil
}

func createDeploymentNode(ctx context.Context, client *ent.Client, entDeployment *ent.Deployment, entResource *ent.Resource, nodeMap map[string]*ent.DeploymentNode, prevNode *ent.DeploymentNode) (*ent.DeploymentNode, error) {
	var entDeploymentNode *ent.DeploymentNode
	var err error
	var ok bool

	// Create the node if it hasn't been created already
	if entDeploymentNode, ok = nodeMap[entResource.Key]; !ok {
		// Create the node
		entDeploymentNode, err = client.DeploymentNode.Create().
			SetState(deploymentnode.StateToDeploy).
			SetVars(map[string]string{}).
			SetDeployment(entDeployment).
			SetResource(entResource).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create deployment node for resource %s: %v", entResource.Key, err)
		}
		// Add it to the node map
		nodeMap[entResource.Key] = entDeploymentNode
	}

	// Add prev node if exists
	if prevNode != nil {
		err = entDeploymentNode.Update().AddPrevNodes(prevNode).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to add prev node to deployment node: %v", err)
		}
	}

	// Create all dependent nodes
	entDependentResource, err := entResource.QueryRequiredBy().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query dependent resources for resource %s: %v", entResource.Key, err)
	}
	for _, entDependentResource := range entDependentResource {
		_, err = createDeploymentNode(ctx, client, entDeployment, entDependentResource, nodeMap, entDeploymentNode)
		if err != nil {
			return nil, err
		}
	}

	return entDeploymentNode, nil
}

func checkProjectQuotas(ctx context.Context, entBlueprint *ent.Blueprint, entProject *ent.Project) (*pgrpc.QuotaRequirements, error) {
	// Query all the resources from blueprint
	entResources, err := entBlueprint.QueryResources().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query resources: %v", err)
	}

	// Add up the new usage level
	newUsageCPU := entProject.UsageCPU
	newUsageRAM := entProject.UsageRAM
	newUsageDisk := entProject.UsageDisk
	newUsageNetwork := entProject.UsageNetwork
	newUsageRouter := entProject.UsageRouter
	for _, entResource := range entResources {
		newUsageCPU += int(entResource.QuotaRequirements.Cpu)
		newUsageRAM += int(entResource.QuotaRequirements.Ram)
		newUsageDisk += int(entResource.QuotaRequirements.Disk)
		newUsageNetwork += int(entResource.QuotaRequirements.Network)
		newUsageRouter += int(entResource.QuotaRequirements.Router)
	}

	// Only check quota if not unlimited and ensure we'll be under the quota limit on this project
	if entProject.QuotaCPU >= 0 && newUsageCPU > entProject.QuotaCPU {
		return nil, fmt.Errorf("not enough CPU quota: requires %d and have %d", newUsageCPU, entProject.QuotaCPU)
	}
	if entProject.QuotaRAM >= 0 && newUsageRAM > entProject.QuotaRAM {
		return nil, fmt.Errorf("not enough RAM quota: requires %d and have %d", newUsageRAM, entProject.QuotaRAM)
	}
	if entProject.QuotaDisk >= 0 && newUsageDisk > entProject.QuotaDisk {
		return nil, fmt.Errorf("not enough Disk quota: requires %d and have %d", newUsageDisk, entProject.QuotaDisk)
	}
	if entProject.QuotaCPU >= 0 && newUsageNetwork > entProject.QuotaNetwork {
		return nil, fmt.Errorf("not enough Network quota: requires %d and have %d", newUsageNetwork, entProject.QuotaNetwork)
	}
	if entProject.QuotaRouter >= 0 && newUsageRouter > entProject.QuotaRouter {
		return nil, fmt.Errorf("not enough Router quota: requires %d and have %d", newUsageRouter, entProject.QuotaRouter)
	}

	return &pgrpc.QuotaRequirements{
		Cpu:     uint64(newUsageCPU),
		Ram:     uint64(newUsageRAM),
		Disk:    uint64(newUsageDisk),
		Network: uint64(newUsageNetwork),
		Router:  uint64(newUsageRouter),
	}, nil
}
