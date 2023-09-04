package main

import (
	"fmt"

	"github.com/cble-platform/backend/config"
	"github.com/cble-platform/backend/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var cfgFile string

func main() {
	// CLI flags
	pflag.StringVar(&cfgFile, "config", "", "the path to the config file")
	pflag.Parse()

	fmt.Println(internal.Logo())
	_, err := config.LoadConfig(cfgFile)
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}
}
