package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/cble-platform/cble-backend/cble"
	"github.com/cble-platform/cble-backend/internal/logo"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	ctx, closeContext := context.WithCancel(context.Background())

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

	// Run all runtimes
	var wg sync.WaitGroup
	cbleServer.Run(ctx, &wg)

	// Close global context on signal receive
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		s := <-sigCh
		logrus.Warnf("Received signal %v, attempting graceful shutdown...", s)
		closeContext()
		wg.Done()
	}()

	// Wait for all runtimes to shutdown
	wg.Wait()

	// Shutdown
	if err := cbleServer.Shutdown(); err != nil {
		logrus.Fatalf("failed to initialize CBLE server: %v", err)
	}

	logrus.Infof("CBLE shutdown successful")
}
