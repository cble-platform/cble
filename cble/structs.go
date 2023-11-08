package cble

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/internal/database"
	"github.com/cble-platform/cble-backend/internal/permissionengine"
	"github.com/cble-platform/cble-backend/internal/utils"
	"github.com/cble-platform/cble-backend/internal/webserver"
)

type CBLEServer struct {
	Config           *config.Config
	Ent              *ent.Client
	Webserver        *webserver.CBLEWebserver
	PermissionEngine *permissionengine.PermissionEngine
}

func NewServer(ctx context.Context, configFile string) (*CBLEServer, error) {
	//-------------//
	// Load config //
	//-------------//

	cbleConfig, err := config.LoadConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	//-----//
	// ENT //
	//-----//

	client, err := database.Initialize(ctx, cbleConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	//-----------//
	// Webserver //
	//-----------//

	w := webserver.New(cbleConfig, client)

	//-------------------//
	// Permission Engine //
	//-------------------//

	pe, err := permissionengine.New(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create permission engine")
	}

	return &CBLEServer{
		Config:           cbleConfig,
		Ent:              client,
		Webserver:        w,
		PermissionEngine: pe,
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

// Run is a blocking method (for now) which should be called to run the main webserver
func (s *CBLEServer) Run(ctx context.Context) {
	//-----------//
	// Webserver //
	//-----------//

	s.Webserver.Listen()
}

// Shutdown should be called after Run returns
func (s *CBLEServer) Shutdown(ctx context.Context) error {
	//-----//
	// Ent //
	//-----//

	err := s.Ent.Close()
	if err != nil {
		return fmt.Errorf("failed to shutdown ent client: %v", err)
	}

	return nil
}
