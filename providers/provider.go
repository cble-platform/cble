package providers

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"

	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/internal/git"
	providerGRPC "github.com/cble-platform/cble-provider-grpc/pkg/provider"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (ps *CBLEServer) downloadProvider(entProvider *ent.Provider) error {
	providerRepoPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "source")
	logrus.WithFields(logrus.Fields{"repo_path": providerRepoPath}).Debugf("Downloading provider %s", entProvider.ID.String())

	// Clone/checkout the provider from git if needed
	if _, err := os.Stat(providerRepoPath); os.IsNotExist(err) {
		logrus.Debugf("Provider does not exist, cloning repo")
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

	providerBinaryPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "provider")
	logrus.WithFields(logrus.Fields{"binary_path": providerBinaryPath}).Debugf("Compiling provider %s", entProvider.ID.String())

	// Build the provider into a binary
	cmd := exec.Command("sh", "-c", fmt.Sprintf("go get ./... && go build -o %s %s", providerBinaryPath, providerRepoPath))
	cmd.Dir = providerRepoPath
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("command exited with status %d. output: %s", exiterr.ExitCode(), cmdOutput)
		}
		return fmt.Errorf("command run error: %v", err)
	}

	return nil
}

// Runs a provider binary. Should be run as a go routine
func (ps *CBLEServer) runProvider(ctx context.Context, entProvider *ent.Provider, shutdown chan bool) {
	providerBinaryPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "provider")

	// Check the provider is compiled
	if _, err := os.Stat(providerBinaryPath); os.IsNotExist(err) {
		// The provider binary has yet to be built
		logrus.Errorf("failed to run provider server: provider has not been compiled yet")
		return
	}

	logrus.Debugf("Executing provider server binary for %s", entProvider.ID.String())

	// Start the binary with the provider ID as argument
	cmd := exec.Command(providerBinaryPath, entProvider.ID.String())
	if err := cmd.Start(); err != nil {
		logrus.Errorf("failed to run provider server: failed to start provider: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			logrus.Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			cmd.Process.Signal(syscall.SIGTERM)
			return
		case <-shutdown:
			logrus.Warnf("Gracefully shutting down Provider %s", entProvider.DisplayName)
			cmd.Process.Signal(syscall.SIGTERM)
			return
		}
	}
}

func (ps *CBLEServer) startProviderConnection(ctx context.Context, shutdown chan bool, providerId string) {
	registeredProvider, exists := ps.registeredProviders.Load(providerId)
	if !exists {
		logrus.Errorf("attempted to start provider on non-registered provider (%s)", providerId)
		return
	}

	logrus.Debugf("starting provider connection to provider %s with socket ID %s", providerId, registeredProvider.(RegisteredProvider).SocketID)

	providerOpts := &providerGRPC.ProviderClientOptions{
		// TODO: implement TLS for provider connections
		TLS:      false,
		CAFile:   "",
		SocketID: registeredProvider.(RegisteredProvider).SocketID,
	}
	providerConn, err := providerGRPC.Connect(providerOpts)
	if err != nil {
		logrus.Errorf("failed to connect to provider gRPC server (%s): %v", providerId, err)
		return
	}
	client, err := providerGRPC.NewClient(ctx, providerConn)
	if err != nil {
		logrus.Errorf("failed to create client for provider (%s): %v", providerId, err)
		return
	}
	// Store the client reference for synchronous use
	ps.providerClients.Store(providerId, client)

	// Convert providerId to UUID
	providerUuid, err := uuid.Parse(providerId)
	if err != nil {
		logrus.Errorf("failed to parse providerId as UUID: %v", err)
		return
	}

	// Get the provider from ENT
	entProvider, err := ps.entClient.Provider.Get(ctx, providerUuid)
	if err != nil {
		logrus.Errorf("failed to query provider: %v", err)
	}

	// Configure the provider
	reply, err := ps.Configure(ctx, entProvider)
	if err != nil || !reply.Success {
		logrus.Errorf("failed to configure provider %s: %v", providerId, err)
		return
	}
}
