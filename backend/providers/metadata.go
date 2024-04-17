package providers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func ParseMetadata(providerRepoPath string) (*ProviderMetadata, error) {
	metaFile := ""
	filepath.WalkDir(providerRepoPath, func(path string, d fs.DirEntry, err error) error {
		if filepath.Base(path) == "cble-metadata.yaml" || filepath.Base(path) == "cble-metadata.yml" {
			metaFile = path
			return nil
		}
		return nil
	})
	if metaFile == "" {
		return nil, fmt.Errorf("no cble-metadata.yaml file found in provider repo")
	}

	metaBytes, err := os.ReadFile(metaFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %v", metaFile, err)
	}

	loadedMetadata := ProviderMetadata{}
	err = yaml.Unmarshal(metaBytes, &loadedMetadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML syntax in %s: %v", metaFile, err)
	}

	logrus.WithField("component", "PROVIDER_METADATA").Debugf("Parsed metadata file %s", metaFile)

	return &loadedMetadata, nil
}
