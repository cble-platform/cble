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
	"github.com/sirupsen/logrus"
)

type DeploymentState int

const (
	DeployFAILED     DeploymentState = -1
	DeploySUCCEEDED  DeploymentState = 0
	DeployINPROGRESS DeploymentState = 1
	DeployDESTROYED  DeploymentState = 2
)

type CommandType int

const (
	CommandCONFIGURE CommandType = 1
	CommandDEPLOY    CommandType = 2
	CommandDESTROY   CommandType = 3
)

type ProviderCommand struct {
	// Type of command
	Type CommandType

	// Channel to send error messages back over
	err chan error
	// Channel to send reply back over
	reply chan interface{}

	// Used for CONFIGURE
	ConfigureRequest *providerGRPC.ConfigureRequest
	// Used for DEPLOY
	DeployRequest *providerGRPC.DeployRequest
	// Used for DESTROY
	DestroyRequest *providerGRPC.DestroyRequest
}

func (ps *CBLEServer) downloadProvider(entProvider *ent.Provider) error {
	providerRepoPath := path.Join(ps.providersConfig.CacheDir, entProvider.ID.String(), "source")
	logrus.WithFields(logrus.Fields{"repo_path": providerRepoPath}).Debugf("Downloading provider %s", entProvider.ID.String())

	// Clone/checkout the provider from git
	if _, err := os.Stat(providerRepoPath); os.IsNotExist(err) {
		logrus.Debugf("Provider does not exist, cloning repo")
		// Provider dir doesn't exist so clone repo
		err := git.CloneProvider(providerRepoPath, entProvider)
		if err != nil {
			return fmt.Errorf("failed to clone provider repo: %v", err)
		}
	} else {
		logrus.Debugf("Provider exists, checking out version")
		// Provider dir exists so just checkout new version
		err := git.CheckoutProvider(providerRepoPath, entProvider)
		if err != nil {
			return fmt.Errorf("failed to checkout provider repo: %v", err)
		}
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

func (ps *CBLEServer) startProviderConnection(ctx context.Context, shutdown chan bool, providerKey string, commandQueue chan ProviderCommand) {
	provider, exists := ps.registeredProviders.Load(providerKey)
	if !exists {
		logrus.Errorf("attempted to start provider on non-registered provider (%s)", providerKey)
		return
	}

	providerOpts := &providerGRPC.ProviderClientOptions{
		// TODO: implement TLS for provider connections
		TLS:      false,
		CAFile:   "",
		SocketID: provider.(RegisteredProvider).SocketID,
	}
	providerConn, err := providerGRPC.Connect(providerOpts)
	if err != nil {
		logrus.Errorf("failed to connect to provider gRPC server (%s): %v", providerKey, err)
		return
	}
	client, err := providerGRPC.NewClient(ctx, providerConn)
	if err != nil {
		logrus.Errorf("failed to create client for provider (%s): %v", providerKey, err)
		return
	}

	// Provider connection event loop
	for {
		select {
		case command := <-commandQueue:
			go ps.handleProviderCommand(ctx, client, &command)
		case <-shutdown:
			logrus.Warnf("Gracefully shutting down provider client %s", providerKey)
			return
		case <-ctx.Done():
			logrus.Warnf("Gracefully shutting down provider client %s", providerKey)
			return
		}
	}
}

func (ps *CBLEServer) handleProviderCommand(ctx context.Context, client providerGRPC.ProviderClient, command *ProviderCommand) {
	switch command.Type {
	case CommandCONFIGURE:
		// Check the configuration request is populated
		if command.ConfigureRequest == nil {
			command.err <- fmt.Errorf("configure request is nil")
			return
		}
		// Send the configuration request
		reply, err := client.Configure(ctx, command.ConfigureRequest)
		if err != nil {
			command.err <- err
		}
		command.reply <- reply
	case CommandDEPLOY:
		// Check the deployment request is populated
		if command.DeployRequest == nil {
			command.err <- fmt.Errorf("deploy request is nil")
			return
		}
		// Send the deploy request
		reply, err := client.Deploy(ctx, command.DeployRequest)
		if err != nil {
			command.err <- err
		}
		command.reply <- reply
	case CommandDESTROY:
		// Check the destroy request is populated
		if command.DestroyRequest == nil {
			command.err <- fmt.Errorf("destroy request is nil")
			return
		}
		// Send the destroy request
		reply, err := client.Destroy(ctx, command.DestroyRequest)
		if err != nil {
			command.err <- err
		}
		command.reply <- reply
	}
}
