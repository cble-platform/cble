package main

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble-backend/cble"
	"github.com/cble-platform/cble-backend/internal/logo"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	ctx := context.Background()

	// Print the logo
	fmt.Println(logo.Print())

	// Parse the config path from command line args
	var cfgFile string
	pflag.StringVar(&cfgFile, "config", "", "the path to the config file")
	pflag.Parse()

	// Create the CBLEServer instance
	cbleServer, err := cble.NewServer(ctx, cfgFile)
	if err != nil {
		logrus.Fatalf("failed to create CBLE server: %v", err)
	}

	// Do the CBLE lifetime

	// Initialize
	if err := cbleServer.Initialize(ctx); err != nil {
		logrus.Fatalf("failed to initialize CBLE server: %v", err)
	}
	// Run (blocking)
	cbleServer.Run(ctx)
	// Shutdown
	if err := cbleServer.Shutdown(ctx); err != nil {
		logrus.Fatalf("failed to initialize CBLE server: %v", err)
	}
}
