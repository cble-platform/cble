package providers

import (
	"context"
	"fmt"
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
