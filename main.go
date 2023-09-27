package main

import (
	"context"
	"fmt"

	"github.com/cble-platform/backend/config"
	"github.com/cble-platform/backend/internal"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var cfgFile string

func main() {
	fmt.Println(internal.Logo())
	//-------------//
	// Load config //
	//-------------//

	pflag.StringVar(&cfgFile, "config", "", "the path to the config file")
	pflag.Parse()

	cbleConfig, err := config.LoadConfig(cfgFile)
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()

	//-----//
	// ENT //
	//-----//

	client, err := internal.InitializeDatabase(ctx, cbleConfig)
	if err != nil {
		logrus.Fatalf("failed to initialize database: %v", err)
	}
	defer client.Close()

	//---------------//
	// Default Admin //
	//---------------//

	err = internal.InitializeDefaultAdminUserGroup(ctx, client, cbleConfig)
	if err != nil {
		logrus.Fatalf("failed to initialize default admin user/group")
	}

	//-----------//
	// Webserver //
	//-----------//

	w := internal.NewWebserver(cbleConfig, client)
	w.Listen()
}
