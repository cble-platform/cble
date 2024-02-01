package providers

// // Runs a synchronous one-off GetConsole command
// func (ps *CBLEServer) GetConsoleSync(ctx context.Context, entProvider *ent.Provider, deploymentNode *ent.DeploymentNode) (*providerGRPC.GetConsoleReply, error) {
// 	entDeployment, err := deploymentNode.QueryDeployment().Only(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	entResource, err := deploymentNode.QueryResource().Only(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	request := &providerGRPC.GetConsoleRequest{
// 		DeploymentId: entDeployment.ID.String(),
// 		HostKey:      entResource.Key,
// 		// Vars:           &structpb.Struct{},
// 		// DeploymentVars: &structpb.Struct{},
// 	}

// 	client, err := ps.getProviderClient(entProvider.ID.String())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.GetConsole(ctx, request)
// }

// // Runs a synchronous one-off GetConsole command
// func (ps *CBLEServer) GetResourceListSync(ctx context.Context, entProvider *ent.Provider, request *providerGRPC.GetResourceListRequest) (*providerGRPC.GetResourceListReply, error) {
// 	client, err := ps.getProviderClient(entProvider.ID.String())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.GetResourceList(ctx, request)
// }
