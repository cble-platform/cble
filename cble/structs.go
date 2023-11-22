package cble

import (
	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/internal/permissionengine"
	"github.com/cble-platform/cble-backend/internal/webserver"
	"github.com/cble-platform/cble-backend/providers"
)

type CBLEServer struct {
	Config           *config.Config
	Ent              *ent.Client
	Webserver        *webserver.CBLEWebserver
	PermissionEngine *permissionengine.PermissionEngine
	GRPCServer       *providers.CBLEServer
}
