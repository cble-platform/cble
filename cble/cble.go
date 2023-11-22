package cble

import (
	"context"
	"fmt"
	"sync"

	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/internal/database"
	"github.com/cble-platform/cble-backend/internal/permissionengine"
	"github.com/cble-platform/cble-backend/internal/utils"
	"github.com/cble-platform/cble-backend/internal/webserver"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/sirupsen/logrus"
)

func NewServer(ctx context.Context, configFile string) (*CBLEServer, error) {
	//-------------//
	// Load config //
	//-------------//

	cbleConfig, err := config.LoadConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	// Enable debug logging in debug mode
	if cbleConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	//-----//
	// ENT //
	//-----//

	client, err := database.Initialize(ctx, cbleConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	//-------------------//
	// Permission Engine //
	//-------------------//

	pe, err := permissionengine.New(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create permission engine")
	}

	//------------------//
	// gRPC CBLE Server //
	//------------------//

	grpcServer := providers.NewServer(client, &cbleConfig.Providers)

	//-----------//
	// Webserver //
	//-----------//

	w := webserver.New(cbleConfig, client, grpcServer)

	return &CBLEServer{
		Config:           cbleConfig,
		Ent:              client,
		Webserver:        w,
		PermissionEngine: pe,
		GRPCServer:       grpcServer,
	}, nil
}

// Initialize should be called before Run
func (s *CBLEServer) Initialize(ctx context.Context) error {
	//---------------//
	// Default Admin //
	//---------------//

	err := utils.InitializeDefaultAdminUserGroup(ctx, s.Ent, s.Config)
	if err != nil {
		return fmt.Errorf("failed to initialize default admin user/group")
	}

	return nil
}

// Run is called to start all of the CBLE runtimes
func (s *CBLEServer) Run(ctx context.Context, wg *sync.WaitGroup) {
	//------------------//
	// gRPC CBLE Server //
	//------------------//
	wg.Add(3)
	go s.GRPCServer.Listen(ctx, wg)
	go s.GRPCServer.RunProviderServers(ctx, wg)
	go s.GRPCServer.RunProviderClients(ctx, wg)

	//-----------//
	// Providers //
	//-----------//

	err := s.GRPCServer.RunAllProviders(ctx)
	if err != nil {
		logrus.Errorf("failed to initialize providers: %v", err)
	}

	//-----------//
	// Webserver //
	//-----------//

	wg.Add(1)
	go s.Webserver.Listen(ctx, wg)
}

// Shutdown should be called after Run returns
func (s *CBLEServer) Shutdown() error {
	//-----//
	// Ent //
	//-----//

	err := s.Ent.Close()
	if err != nil {
		return fmt.Errorf("failed to shutdown ent client: %v", err)
	}

	return nil
}
