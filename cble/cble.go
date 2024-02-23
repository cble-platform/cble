package cble

import (
	"context"
	"fmt"
	"sync"

	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/engine"
	"github.com/cble-platform/cble-backend/internal/database"
	"github.com/cble-platform/cble-backend/internal/defaultadmin"
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
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	//-----//
	// ENT //
	//-----//

	client, err := database.Initialize(ctx, cbleConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
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
		Config:     cbleConfig,
		Ent:        client,
		Webserver:  w,
		GRPCServer: grpcServer,
	}, nil
}

// Initialize should be called before Run
func (s *CBLEServer) Initialize(ctx context.Context) error {
	//-------------//
	// Permissions //
	//-------------//

	// for _, perm := range s.Config.Initialization.Permissions {
	// 	_, err := s.PermissionEngine.RegisterPermission(ctx, perm.Key, perm.Component, perm.Description)
	// 	if err != nil {
	// 		logrus.Warnf("failed to register permission %s: %v", perm.Key, err)
	// 	}
	// }

	//---------------//
	// Default Admin //
	//---------------//

	err := defaultadmin.InitializeDefaultAdminUserGroup(ctx, s.Ent, s.Config)
	if err != nil {
		return fmt.Errorf("failed to initialize default admin user/group: %v", err)
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

	if s.Config.Providers.AutoLoad == nil || *s.Config.Providers.AutoLoad {
		err := s.GRPCServer.RunAllProviders(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"component": "GRPC_SERVER",
			}).Errorf("failed to initialize providers: %v", err)
		}
	}

	//-----------//
	// Webserver //
	//-----------//

	wg.Add(1)
	go s.Webserver.Listen(ctx, wg)

	//----------//
	// Runtimes //
	//----------//

	go engine.AutoSuspendDeploymentWatchdog(ctx, s.Ent, s.GRPCServer, s.Config.Deployments.AutoSuspendTime)
	go engine.AutoDestroyExpiredDeploymentWatchdog(ctx, s.Ent, s.GRPCServer)
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
