package providers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	cbleGRPC "github.com/cble-platform/cble-provider-grpc/pkg/cble"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CBLEServer struct {
	cbleGRPC.DefaultCBLEServer

	// ENT Client
	entClient *ent.Client
	// CBLE Config
	providersConfig *config.ProvidersConfig

	// Queue of all providers to start
	providerServerQueue chan string
	// Shutdown channels for each provider
	serverShutdown *sync.Map

	// Register individual provider registrations
	registeredProviders *sync.Map
	// Individual provider client references
	providerClients *sync.Map
	// Queue of providers to connect to after registration
	connectionQueue chan string
}

type RegisteredProvider struct {
	ID       string
	SocketID string
	Features *cbleGRPC.ProviderFeatures
}

func NewServer(entClient *ent.Client, providersConfig *config.ProvidersConfig) *CBLEServer {
	return &CBLEServer{
		entClient:           entClient,
		providersConfig:     providersConfig,
		providerServerQueue: make(chan string),
		serverShutdown:      new(sync.Map),
		registeredProviders: new(sync.Map),
		providerClients:     new(sync.Map),
		connectionQueue:     make(chan string, 10),
	}
}

// Listens on configured port. Runs in a go routine
func (ps *CBLEServer) Listen(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	serveCtx, cancelServeCtx := context.WithCancel(context.Background())

	go func() {
		// Auto exists when the context is cancelled
		if err := cbleGRPC.DefaultServe(serveCtx, ps); err != nil {
			logrus.WithFields(logrus.Fields{
				"component": "GRPC_SERVER",
			}).Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for parent context to close
	<-ctx.Done()
	// Set timeout timestamp
	timeoutTime := time.Now().Add(30 * time.Second)
	// Wait for all providers to unregister
	for {
		// If we reached the timeout, just quit
		if time.Until(timeoutTime) <= 0 {
			break
		}
		logrus.WithFields(logrus.Fields{
			"component": "GRPC_SERVER",
		}).Infof("Waiting for all providers to unregister (%.0fs)...", time.Until(timeoutTime).Seconds())
		providersAreLoaded := false
		ps.registeredProviders.Range(func(key, value any) bool {
			providersAreLoaded = true
			return false
		})
		// If no providers are loaded break
		if !providersAreLoaded {
			break
		}
		time.Sleep(1 * time.Second)
	}
	logrus.WithFields(logrus.Fields{
		"component": "GRPC_SERVER",
	}).Info("All providers unregistered!")
	logrus.WithFields(logrus.Fields{
		"component": "GRPC_SERVER",
	}).Warnf("Gracefully shutting down CBLE gRPC server...")
	// Shutdown gRPC server
	cancelServeCtx()
}

// Runs all provider servers queued to start. Runs in a go routine
func (ps *CBLEServer) RunProviderServers(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// When a provider server is queued, run it in a go routine
	for {
		select {
		case providerId := <-ps.providerServerQueue:
			logrus.WithFields(logrus.Fields{
				"component":  "PROVIDER_ENGINE",
				"providerId": providerId,
			}).Debugf("Running provider server for %s", providerId)
			providerUuid, err := uuid.Parse(providerId)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"component":  "PROVIDER_ENGINE",
					"providerId": providerId,
				}).Errorf("failed start provider server: failed to parse provider UUID %s", providerId)
				continue
			}
			entProvider, err := ps.entClient.Provider.Get(ctx, providerUuid)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"component":  "PROVIDER_ENGINE",
					"providerId": providerId,
				}).Errorf("failed start provider server: failed to find provider with ID %s", providerId)
				continue
			}
			// Ensure the provider is downloaded/updates
			err = ps.downloadProvider(entProvider)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"component":  "PROVIDER_ENGINE",
					"providerId": providerId,
				}).Errorf("failed to start provider server: failed to download/update provider: %v", err)
				continue
			}
			// Create an individual shutdown channel for this provider
			shutdownChan := make(chan bool)
			ps.serverShutdown.Store(providerId, shutdownChan)
			// Run the provider server in a go routine
			go ps.runProvider(ctx, entProvider, shutdownChan)
		case <-ctx.Done():
			logrus.WithFields(logrus.Fields{
				"component": "PROVIDER_ENGINE",
			}).Warn("Gracefully shutting down provider server runtime...")
			return
		}
	}
}

// Runs all provider clients queued to connect. Runs in a go routine
func (ps *CBLEServer) RunProviderClients(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// Wait for a connection in the queue start provider routine
	for {
		select {
		case providerId := <-ps.connectionQueue:
			logrus.WithFields(logrus.Fields{
				"component": "PROVIDER_ENGINE",
			}).Debugf("Running provider client for %s", providerId)
			go ps.startProviderConnection(ctx, providerId)
		case <-ctx.Done():
			logrus.WithFields(logrus.Fields{
				"component": "PROVIDER_ENGINE",
			}).Warn("Gracefully shutting down provider client runtime...")
			return
		}
	}
}

func (ps *CBLEServer) QueueLoadProvider(id string) {
	logrus.WithFields(logrus.Fields{
		"component": "PROVIDER_ENGINE",
	}).Debugf("Loading provider %s", id)
	ps.providerServerQueue <- id
}

func (ps *CBLEServer) QueueUnloadProvider(id string) error {
	logrus.WithFields(logrus.Fields{
		"component": "PROVIDER_ENGINE",
	}).Debugf("Unloading provider %s", id)
	// Check that the server shutdown channel exists
	serverShutdown, ok := ps.serverShutdown.Load(id)
	if !ok {
		return fmt.Errorf("provider server has no shutdown channel")
	}
	// Send the shutdown signal to the provider server
	serverShutdown.(chan bool) <- true
	return nil
}

// func (ps *CBLEServer) StopAllProviders(ctx context.Context) {}

func (ps *CBLEServer) RegisterProvider(ctx context.Context, request *cbleGRPC.RegistrationRequest) (*cbleGRPC.RegistrationReply, error) {
	logrus.WithFields(logrus.Fields{
		"component": "PROVIDER_ENGINE",
	}).Debugf("Registration request from %s@%s (%s)", request.Name, request.Version, request.Id)
	// Check if a provider with this ID is already registered
	if _, exist := ps.registeredProviders.Load(request.Id); exist {
		return nil, fmt.Errorf("provider with same ID (%s) already registered", request.Id)
	}
	// Check this is a valid UUID
	providerUuid, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, fmt.Errorf("provider did not supply a valid ID: %v", err)
	}
	// Check this UUID maps to a valid ENT provider
	entProvider, err := ps.entClient.Provider.Get(ctx, providerUuid)
	if err != nil {
		return nil, fmt.Errorf("provider not found with ID %s: %v", request.Id, err)
	}
	// Generate random UUID for socket
	socketId := uuid.NewString()
	// Map the port
	ps.registeredProviders.Store(request.Id, RegisteredProvider{
		ID:       request.Id,
		SocketID: socketId,
		Features: request.Features,
	})
	// Add provider to the queue to be connected to
	ps.connectionQueue <- request.Id
	// Set the provider as loaded in ENT
	err = entProvider.Update().SetIsLoaded(true).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set provider is_loaded state: %v", err)
	}
	logrus.WithFields(logrus.Fields{
		"component": "PROVIDER_ENGINE",
	}).Debugf("requesting provider %s start listening on socket ID %s", request.Id, socketId)
	// Reply to the provider
	return &cbleGRPC.RegistrationReply{
		Success:  true,
		SocketId: socketId,
	}, nil
}

func (ps *CBLEServer) UnregisterProvider(ctx context.Context, request *cbleGRPC.UnregistrationRequest) (*cbleGRPC.UnregistrationReply, error) {
	logrus.WithFields(logrus.Fields{
		"component": "PROVIDER_ENGINE",
	}).Debugf("Unregistration request from %s@%s (%s)", request.Name, request.Version, request.Id)
	// Check to make sure this provider is registered
	prov, exists := ps.registeredProviders.Load(request.Id)
	if !exists {
		return &cbleGRPC.UnregistrationReply{
			Success: false,
		}, nil
	}
	// Check this is a valid UUID
	providerUuid, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, fmt.Errorf("provider did not supply a valid ID: %v", err)
	}
	// Check this UUID maps to a valid ENT provider
	entProvider, err := ps.entClient.Provider.Get(ctx, providerUuid)
	if err != nil {
		return nil, fmt.Errorf("provider not found with ID %s: %v", request.Id, err)
	}
	// Make sure the unregister request is coming with the right ID... super basic security check :)
	if prov.(RegisteredProvider).ID != request.Id {
		return &cbleGRPC.UnregistrationReply{
			Success: false,
		}, nil
	}
	// If all that passes, unregister the provider
	ps.registeredProviders.Delete(request.Id)
	// Set the provider as loaded in ENT
	err = entProvider.Update().SetIsLoaded(false).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set provider is_loaded state: %v", err)
	}
	return &cbleGRPC.UnregistrationReply{
		Success: true,
	}, nil
}
