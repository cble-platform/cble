package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	SSL      bool   `yaml:"ssl"`
}

func LoadConfig(cfgFile string) (*Config, error) {
	if cfgFile == "" {
		if f, err := os.Stat("config.local.yaml"); err == nil {
			cfgFile = f.Name()
		} else if f, err := os.Stat("config.yaml"); err == nil {
			cfgFile = f.Name()
		} else {
			return nil, fmt.Errorf("could not find any config file (one of: config.yaml, config.local.yaml)")
		}
	}

	configBytes, err := os.ReadFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %v", cfgFile, err)
	}

	loadedConfig := Config{}
	err = yaml.Unmarshal(configBytes, &loadedConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML syntax in %s: %v", cfgFile, err)
	}
	return &loadedConfig, nil
}
