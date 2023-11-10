package providers

import (
	"context"
	"fmt"
	"sync"

	"github.com/cble-platform/cble-backend/ent"
	cbleGRPC "github.com/cble-platform/cble-provider-grpc/pkg/cble"
	"github.com/cble-platform/cble-provider-grpc/pkg/common"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CBLEServer struct {
	cbleGRPC.DefaultCBLEServer

	// ENT Client
	entClient *ent.Client

	// Send shutdown signal to individual routines
	shutdown *sync.Map
	// Channels to send commands to individual providers
	commandQueues *sync.Map

	// Register individual provider registrations
	registeredProviders *sync.Map
	// Queue of providers to connect to after registration
	connectionQueue chan string
	// Shutdown channels for each provider
	providerShutdown *sync.Map
}

type RegisteredProvider struct {
	ID       string
	SocketID string
	Features map[string]bool
}

func NewServer(entClient *ent.Client) *CBLEServer {
	return &CBLEServer{
		entClient:           entClient,
		shutdown:            new(sync.Map),
		commandQueues:       new(sync.Map),
		registeredProviders: new(sync.Map),
		connectionQueue:     make(chan string, 10),
		providerShutdown:    new(sync.Map),
	}
}

// Listens on configured port
func (ps *CBLEServer) Listen() {
	if err := cbleGRPC.DefaultServe(ps); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func (ps *CBLEServer) SendCommandToProvider(ctx context.Context, entVirtualizationProvider *ent.VirtualizationProvider, command *ProviderCommand) error {
	commandQueue, ok := ps.commandQueues.Load(entVirtualizationProvider.ID.String())
	if !ok {
		return fmt.Errorf("no command queue registered for provider %s", entVirtualizationProvider.ID.String())
	}
	commandQueue.(chan ProviderCommand) <- *command
	return nil
}

func (ps *CBLEServer) RunProviderClients(ctx context.Context) {
	providerCtx, providerKillAll := context.WithCancel(ctx)
	// Wait for a connection in the queue start provider routine
	for {
		select {
		case providerKey := <-ps.connectionQueue:
			shutdownChan, ok := ps.shutdown.Load(providerKey)
			if !ok {
				logrus.Errorf("attempted to start provider connection without a shutdown channel (%s)", providerKey)
				continue
			}
			commandQueue, ok := ps.commandQueues.Load(providerKey)
			if !ok {
				logrus.Errorf("attempted to start provider connection without a command queue (%s)", providerKey)
				continue
			}
			go ps.StartProviderConnection(providerCtx, shutdownChan.(chan bool), providerKey, commandQueue.(chan ProviderCommand))
		case <-ctx.Done():
			providerKillAll()
		}
	}
}

func (ps *CBLEServer) RegisterProvider(ctx context.Context, request *cbleGRPC.RegistrationRequest) (*cbleGRPC.RegistrationReply, error) {
	logrus.Debugf("Registration request from %s@%s (%s)", request.Name, request.Version, request.Id)
	// Check if a provider with this ID is already registered
	if _, exist := ps.registeredProviders.Load(request.Id); exist {
		return nil, fmt.Errorf("provider with same ID (%s) already registered", request.Id)
	}
	// Check this is a valid UUID
	virtualizationProviderUuid, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, fmt.Errorf("provider did not supply a valid ID: %v", err)
	}
	// Check this UUID maps to a valid ENT virtualization provider
	entVirtualizationProvider, err := ps.entClient.VirtualizationProvider.Get(ctx, virtualizationProviderUuid)
	if err != nil {
		return nil, fmt.Errorf("virtualization provider not found with ID %s: %v", request.Id, err)
	}
	// Generate random UUID for socket
	socketId := uuid.NewString()
	// Map the port
	ps.registeredProviders.Store(request.Id, RegisteredProvider{
		ID:       request.Id,
		SocketID: socketId,
		Features: request.Features,
	})
	// Create shutdown and command queue for provider
	ps.shutdown.Store(request.Id, make(chan bool, 1))
	ps.commandQueues.Store(request.Id, make(chan ProviderCommand, 100)) // TODO: measeure the necessary queue buffer size to better help concurrency
	// Add provider to the queue to be connected to
	ps.connectionQueue <- request.Id
	// Set the provider as loaded in ENT
	err = entVirtualizationProvider.Update().SetIsLoaded(true).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set virtualization provider is_loaded state: %v", err)
	}
	// Reply to the provider
	return &cbleGRPC.RegistrationReply{
		Status:   common.RPCStatus_SUCCESS,
		SocketId: socketId,
	}, nil
}

func (ps *CBLEServer) UnregisterProvider(ctx context.Context, request *cbleGRPC.UnregistrationRequest) (*cbleGRPC.UnregistrationReply, error) {
	logrus.Debugf("Unregistration request from %s@%s (%s)", request.Name, request.Version, request.Id)
	// Check to make sure this provider is registered
	prov, exists := ps.registeredProviders.Load(request.Id)
	if !exists {
		return &cbleGRPC.UnregistrationReply{
			Status: common.RPCStatus_FAILURE,
		}, nil
	}
	// Check this is a valid UUID
	virtualizationProviderUuid, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, fmt.Errorf("provider did not supply a valid ID: %v", err)
	}
	// Check this UUID maps to a valid ENT virtualization provider
	entVirtualizationProvider, err := ps.entClient.VirtualizationProvider.Get(ctx, virtualizationProviderUuid)
	if err != nil {
		return nil, fmt.Errorf("virtualization provider not found with ID %s: %v", request.Id, err)
	}
	// Make sure the unregister request is coming with the right ID... super basic security check :)
	if prov.(RegisteredProvider).ID != request.Id {
		return &cbleGRPC.UnregistrationReply{
			Status: common.RPCStatus_FAILURE,
		}, nil
	}
	// If all that passes, unregister the provider
	ps.registeredProviders.Delete(request.Id)
	// Set the provider as loaded in ENT
	err = entVirtualizationProvider.Update().SetIsLoaded(true).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set virtualization provider is_loaded state: %v", err)
	}
	return &cbleGRPC.UnregistrationReply{
		Status: common.RPCStatus_SUCCESS,
	}, nil
}
