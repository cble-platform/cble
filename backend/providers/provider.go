package providers

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/git"
	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (ps *CBLEServer) downloadProvider(entProvider *ent.Provider) error {
	providerRepoPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "source")
	logrus.WithFields(logrus.Fields{"repo_path": providerRepoPath}).Debugf("Downloading provider %s", entProvider.ID.String())

	// Clone/checkout the provider from git if needed
	if _, err := os.Stat(providerRepoPath); os.IsNotExist(err) {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": entProvider.ID,
		}).Debugf("Provider does not exist, cloning repo")
		// Provider dir doesn't exist so clone repo
		err := git.CloneProvider(providerRepoPath, entProvider)
		if err != nil {
			return fmt.Errorf("failed to clone provider repo: %v", err)
		}
	}
	// Checkout requested version
	err := git.CheckoutProvider(providerRepoPath, entProvider)
	if err != nil {
		return fmt.Errorf("failed to checkout provider repo: %v", err)
	}

	// Build the provider
	err = BuildProvider(providerRepoPath, entProvider)
	if err != nil {
		return fmt.Errorf("failed to build provider: %v", err)
	}

	return nil
}

// Runs a provider binary. Should be run as a go routine
func (ps *CBLEServer) runProvider(ctx context.Context, entProvider *ent.Provider, shutdown chan bool) {
	logrus.WithFields(logrus.Fields{
		"component":  "PROVIDER_ENGINE",
		"providerId": entProvider.ID,
	}).Debugf("Executing provider server binary for %s", entProvider.ID.String())

	// Start the provider (blocking)
	ExecuteProvider(ctx, ps.providersConfig.CacheDir, entProvider, shutdown)
}

func (ps *CBLEServer) startProviderConnection(ctx context.Context, providerId string) {
	registeredProvider, exists := ps.registeredProviders.Load(providerId)
	if !exists {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": providerId,
		}).Errorf("attempted to start provider on non-registered provider (%s)", providerId)
		return
	}

	logrus.WithFields(logrus.Fields{
		"component":  "PROVIDER_ENGINE",
		"providerId": providerId,
	}).Debugf("starting provider connection to provider %s with socket ID %s", providerId, registeredProvider.(RegisteredProvider).SocketID)

	providerOpts := &pgrpc.ProviderClientOptions{
		// TODO: implement TLS for provider connections
		TLS:      false,
		CAFile:   "",
		SocketID: registeredProvider.(RegisteredProvider).SocketID,
	}
	providerConn, err := pgrpc.Connect(providerOpts)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": providerId,
		}).Errorf("failed to connect to provider gRPC server (%s): %v", providerId, err)
		return
	}
	client, err := pgrpc.NewClient(ctx, providerConn)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": providerId,
		}).Errorf("failed to create client for provider (%s): %v", providerId, err)
		return
	}
	// Store the client reference for synchronous use
	ps.providerClients.Store(providerId, client)

	// Convert providerId to UUID
	providerUuid, err := uuid.Parse(providerId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": providerId,
		}).Errorf("failed to parse providerId as UUID: %v", err)
		return
	}

	// Get the provider from ENT
	entProvider, err := ps.entClient.Provider.Get(ctx, providerUuid)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": providerId,
		}).Errorf("failed to query provider: %v", err)
	}

	// Configure the provider
	reply, err := ps.Configure(ctx, entProvider)
	if err != nil || !reply.Success {
		logrus.WithFields(logrus.Fields{
			"component":  "PROVIDER_ENGINE",
			"providerId": providerId,
		}).Errorf("failed to configure provider %s: %v", providerId, err)
		return
	}
}
