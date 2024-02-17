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
