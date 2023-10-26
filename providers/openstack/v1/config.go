package v1

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type OpenstackProviderConfig struct {
	AuthUrl     string `yaml:"auth_url"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	ProjectID   string `yaml:"project_id"`
	ProjectName string `yaml:"project_name"`
	RegionName  string `yaml:"region_name"`
	DomainName  string `yaml:"domain_name,omitempty"`
	DomainId    string `yaml:"domain_id,omitempty"`
}

func ConfigFromBytes(in []byte) (*OpenstackProviderConfig, error) {
	var config OpenstackProviderConfig
	if err := yaml.Unmarshal(in, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}
	return &config, nil
}
