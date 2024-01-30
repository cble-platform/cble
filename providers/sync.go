package providers

import (
	"context"

	providerGRPC "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

// Runs a synchronous one-off GetConsole command
func (ps *CBLEServer) GetConsoleSync(ctx context.Context, providerId string, request *providerGRPC.GetConsoleRequest) (*providerGRPC.GetConsoleReply, error) {
	client, err := ps.getProviderClient(providerId)
	if err != nil {
		return nil, err
	}

	return client.GetConsole(ctx, request)
}

// Runs a synchronous one-off GetConsole command
func (ps *CBLEServer) GetResourceListSync(ctx context.Context, providerId string, request *providerGRPC.GetResourceListRequest) (*providerGRPC.GetResourceListReply, error) {
	client, err := ps.getProviderClient(providerId)
	if err != nil {
		return nil, err
	}

	return client.GetResourceList(ctx, request)
}
