package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/ent/resource"
)

func CreateDeployment(ctx context.Context, client *ent.Client, entBlueprint *ent.Blueprint, templateVars map[string]string, expiryTime time.Time, requester *ent.User) (*ent.Deployment, error) {
	// Create the deployment
	entDeployment, err := client.Deployment.Create().
		SetName(entBlueprint.Name).
		SetDescription(entBlueprint.Description).
		SetState(deployment.StateAwaiting).
		SetTemplateVars(templateVars).
		SetExpiresAt(expiryTime).
		SetBlueprint(entBlueprint).
		SetRequester(requester).
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
