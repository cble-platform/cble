package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/cble-platform/cble/backend/server"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	ctx, closeContext := context.WithCancel(context.Background())

	// Print the logo
	fmt.Println(PrintLogo())

	// Parse the config path from command line args
	var cfgFile string
	pflag.StringVar(&cfgFile, "config", "", "the path to the config file")
	pflag.Parse()

	// Create the CBLEServer instance
	cbleServer, err := server.NewServer(ctx, cfgFile)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component": "CBLE",
		}).Fatalf("failed to create CBLE server: %v", err)
	}

	// Do the CBLE lifetime

	// Initialize
	if err := cbleServer.Initialize(ctx); err != nil {
		logrus.WithFields(logrus.Fields{
			"component": "CBLE",
		}).Fatalf("failed to initialize CBLE server: %v", err)
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
		logrus.WithFields(logrus.Fields{
			"component": "CBLE",
		}).Warnf("Received signal %v, attempting graceful shutdown...", s)
		closeContext()
		wg.Done()
	}()

	// Wait for all runtimes to shutdown
	wg.Wait()

	// Shutdown
	if err := cbleServer.Shutdown(); err != nil {
		logrus.WithFields(logrus.Fields{
			"component": "CBLE",
		}).Fatalf("failed to initialize CBLE server: %v", err)
	}

	logrus.WithFields(logrus.Fields{
		"component": "CBLE",
	}).Infof("CBLE shutdown successful")
}
