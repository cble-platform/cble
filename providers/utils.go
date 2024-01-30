package providers

import (
	"context"
	"fmt"

	providerGRPC "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (ps *CBLEServer) RunAllProviders(ctx context.Context) error {
	entProviders, err := ps.entClient.Provider.Query().All(ctx)
	if err != nil {
		return fmt.Errorf("failed to query providers: %v", err)
	}

	for _, entProvider := range entProviders {
		// Queue the provider to load
		ps.QueueLoadProvider(entProvider.ID.String())
	}

	return nil
}

// // Converts all deployment maps into protobuf-friendly structs
// func DeploymentMapsToStructs(entDeployment *ent.Deployment) (*structpb.Struct, *structpb.Struct, *structpb.Struct, error) {
// 	// Convert maps into protobuf-friendly structs
// 	templateVarsStruct, err := structpb.NewStruct(entDeployment.TemplateVars)
// 	if err != nil {
// 		return nil, nil, nil, fmt.Errorf("failed to parse template vars into structpb: %v", err)
// 	}
// 	deploymentVarsStruct, err := structpb.NewStruct(entDeployment.DeploymentVars)
// 	if err != nil {
// 		return nil, nil, nil, fmt.Errorf("failed to parse deployment vars into structpb: %v", err)
// 	}
// 	// Deployment state is of type map[string]string and needs to be converted to map[string]interface{}
// 	deploymentState := make(map[string]interface{}, len(entDeployment.DeploymentState))
// 	for k, v := range entDeployment.DeploymentState {
// 		deploymentState[k] = v
// 	}
// 	deploymentStateStruct, err := structpb.NewStruct(deploymentState)
// 	if err != nil {
// 		return nil, nil, nil, fmt.Errorf("failed to parse deployment state into structpb: %v", err)
// 	}

// 	return templateVarsStruct, deploymentVarsStruct, deploymentStateStruct, nil
// }

func (ps *CBLEServer) getProviderClient(providerId string) (providerGRPC.ProviderClient, error) {
	clientRaw, ok := ps.providerClients.Load(providerId)
	if !ok {
		return nil, fmt.Errorf("provider is not loaded")
	}
	client, ok := clientRaw.(providerGRPC.ProviderClient)
	if !ok {
		return nil, fmt.Errorf("provider client not stored properly")
	}
	return client, nil
}
