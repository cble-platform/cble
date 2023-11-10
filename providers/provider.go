package providers

import (
	"context"
	"fmt"

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

func (ps *CBLEServer) StartProviderConnection(ctx context.Context, shutdown chan bool, providerKey string, commandQueue chan ProviderCommand) {
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
			go ps.HandleProviderCommand(ctx, client, &command)
		case <-shutdown:
			return
		case <-ctx.Done():
			return
		}
	}
}

func (ps *CBLEServer) HandleProviderCommand(ctx context.Context, client providerGRPC.ProviderClient, command *ProviderCommand) {
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
