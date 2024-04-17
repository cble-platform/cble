package providers

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/cble-platform/cble/backend/ent"
	"gopkg.in/yaml.v3"
)

// Runs a synchronous GenerateDependencies command
func (ps *CBLEServer) Configure(ctx context.Context, entProvider *ent.Provider) (*pgrpc.ConfigureReply, error) {
	client, err := ps.getProviderClient(entProvider.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider client: %v", err)
	}

	// Create the request
	request := &pgrpc.ConfigureRequest{
		Config: entProvider.ConfigBytes,
	}

	return client.Configure(ctx, request)
}

// Runs a synchronous ExtractResourceMetadata command
func (ps *CBLEServer) ExtractResourceMetadata(ctx context.Context, entProvider *ent.Provider, entResources []*ent.Resource) (*pgrpc.ExtractResourceMetadataReply, error) {
	client, err := ps.getProviderClient(entProvider.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider client: %v", err)
	}

	// Convert resource structs
	resources := make([]*pgrpc.Resource, len(entResources))
	for i, entResource := range entResources {
		// Convert the object to YAML
		objectBytes, err := yaml.Marshal(entResource.Object)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal resource (%s) object into YAML: %v", entResource.ID, err)
		}
		// Convert to gRPC-ready resource
		resources[i] = &pgrpc.Resource{
			Id:     entResource.ID.String(),
			Key:    entResource.Key,
			Object: objectBytes,
		}
	}

	// Create the request
	request := &pgrpc.ExtractResourceMetadataRequest{
		Resources: resources,
	}

	return client.ExtractResourceMetadata(ctx, request)
}

// Runs a synchronous RetrieveData command
func (ps *CBLEServer) RetrieveData(ctx context.Context, entProvider *ent.Provider, entDeploymentNode *ent.DeploymentNode, templatedObject []byte) (*pgrpc.RetrieveDataReply, error) {
	// Get the provider client
	client, err := ps.getProviderClient(entProvider.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider client: %v", err)
	}

	// Get the deployment
	entDeployment, err := entDeploymentNode.QueryDeployment().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query deployment from node: %v", err)
	}
	// Get the resource
	entResource, err := entDeploymentNode.QueryResource().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query resource from node: %v", err)
	}

	// Generate dependency var map
	dependencyVarsMap := make(map[string]*pgrpc.DependencyVars)
	entDependencyNodes, err := entDeploymentNode.QueryPrevNodes().WithResource().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query dependency nodes: %v", err)
	}
	for _, entDependencyNode := range entDependencyNodes {
		// Add the dependency's vars to dependency var map
		dependencyVarsMap[entDependencyNode.Edges.Resource.Key] = &pgrpc.DependencyVars{
			Vars: entDependencyNode.Vars,
		}
	}

	// Create the request
	request := &pgrpc.RetrieveDataRequest{
		Deployment: &pgrpc.Deployment{
			Id:           entDeployment.ID.String(),
			TemplateVars: entDeployment.TemplateVars,
		},
		Resource: &pgrpc.Resource{
			Id:     entResource.ID.String(),
			Key:    entResource.Key,
			Object: templatedObject,
		},
		Vars:           entDeploymentNode.Vars,
		DependencyVars: dependencyVarsMap,
	}

	return client.RetrieveData(ctx, request)
}

// Runs a synchronous DeployResource command
func (ps *CBLEServer) DeployResource(ctx context.Context, entProvider *ent.Provider, entDeploymentNode *ent.DeploymentNode, templatedObject []byte) (*pgrpc.DeployResourceReply, error) {
	// Get the provider client
	client, err := ps.getProviderClient(entProvider.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider client: %v", err)
	}

	// Get the deployment
	entDeployment, err := entDeploymentNode.QueryDeployment().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query deployment from node: %v", err)
	}
	// Get the resource
	entResource, err := entDeploymentNode.QueryResource().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query resource from node: %v", err)
	}

	// Generate dependency var map
	dependencyVarsMap := make(map[string]*pgrpc.DependencyVars)
	entDependencyNodes, err := entDeploymentNode.QueryPrevNodes().WithResource().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query dependency nodes: %v", err)
	}
	for _, entDependencyNode := range entDependencyNodes {
		// Add the dependency's vars to dependency var map
		dependencyVarsMap[entDependencyNode.Edges.Resource.Key] = &pgrpc.DependencyVars{
			Vars: entDependencyNode.Vars,
		}
	}

	// Create the request
	request := &pgrpc.DeployResourceRequest{
		Deployment: &pgrpc.Deployment{
			Id:           entDeployment.ID.String(),
			TemplateVars: entDeployment.TemplateVars,
		},
		Resource: &pgrpc.Resource{
			Id:     entResource.ID.String(),
			Key:    entResource.Key,
			Object: templatedObject,
		},
		Vars:           entDeploymentNode.Vars,
		DependencyVars: dependencyVarsMap,
	}

	return client.DeployResource(ctx, request)
}

// Runs a synchronous DestroyResource command
func (ps *CBLEServer) DestroyResource(ctx context.Context, entProvider *ent.Provider, entDeploymentNode *ent.DeploymentNode) (*pgrpc.DestroyResourceReply, error) {
	// Get the provider client
	client, err := ps.getProviderClient(entProvider.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider client: %v", err)
	}

	// Get the deployment
	entDeployment, err := entDeploymentNode.QueryDeployment().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query deployment from node: %v", err)
	}
	// Get the resource
	entResource, err := entDeploymentNode.QueryResource().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query resource from node: %v", err)
	}

	// Convert the object to YAML
	objectBytes, err := yaml.Marshal(entResource.Object)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource (%s) object into YAML: %v", entResource.ID, err)
	}

	// Create the request
	request := &pgrpc.DestroyResourceRequest{
		Deployment: &pgrpc.Deployment{
			Id:           entDeployment.ID.String(),
			TemplateVars: entDeployment.TemplateVars,
		},
		Resource: &pgrpc.Resource{
			Id:     entResource.ID.String(),
			Key:    entResource.Key,
			Object: objectBytes,
		},
		Vars: entDeploymentNode.Vars,
	}

	return client.DestroyResource(ctx, request)
}

// Runs a synchronous GetConsole command
func (ps *CBLEServer) GetConsole(ctx context.Context, entProvider *ent.Provider, entDeploymentNode *ent.DeploymentNode) (*pgrpc.GetConsoleReply, error) {
	// Get the provider client
	client, err := ps.getProviderClient(entProvider.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider client: %v", err)
	}

	// Get the resource
	entResource, err := entDeploymentNode.QueryResource().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query resource from node: %v", err)
	}

	// Convert the object to YAML
	objectBytes, err := yaml.Marshal(entResource.Object)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource (%s) object into YAML: %v", entResource.ID, err)
	}

	// Create the request
	request := &pgrpc.GetConsoleRequest{
		Resource: &pgrpc.Resource{
			Id:     entResource.ID.String(),
			Key:    entResource.Key,
			Object: objectBytes,
		},
		Vars: entDeploymentNode.Vars,
	}

	return client.GetConsole(ctx, request)
}

// Runs a synchronous ResourcePower command
func (ps *CBLEServer) ResourcePower(ctx context.Context, entProvider *ent.Provider, entDeploymentNode *ent.DeploymentNode, state pgrpc.PowerState) (*pgrpc.ResourcePowerReply, error) {
	// Get the provider client
	client, err := ps.getProviderClient(entProvider.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider client: %v", err)
	}

	// Get the resource
	entResource, err := entDeploymentNode.QueryResource().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query resource from node: %v", err)
	}

	// Convert the object to YAML
	objectBytes, err := yaml.Marshal(entResource.Object)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource (%s) object into YAML: %v", entResource.ID, err)
	}

	// Create the request
	request := &pgrpc.ResourcePowerRequest{
		Resource: &pgrpc.Resource{
			Id:     entResource.ID.String(),
			Key:    entResource.Key,
			Object: objectBytes,
		},
		Vars:  entDeploymentNode.Vars,
		State: state,
	}

	return client.ResourcePower(ctx, request)
}
