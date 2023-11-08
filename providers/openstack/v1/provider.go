package v1

import "fmt"

type OpenstackProvider struct {
	config *OpenstackProviderConfig
}

const (
	id          = "openstack_v1"
	name        = "Openstack"
	description = "Builder that interfaces with Openstack"
	author      = "Bradley Harker <github.com/BradHacker>"
	version     = "1.0"
)

func NewProvider(config *OpenstackProviderConfig) (*OpenstackProvider, error) {
	// Create the provider with passed config
	provider := OpenstackProvider{
		config: config,
	}
	// Test the connection
	if _, err := provider.newAuthClient(); err != nil {
		return nil, fmt.Errorf("connection test failed: %v", err)
	}

	return &provider, nil
}

func (provider *OpenstackProvider) ID() string {
	return id
}

func (provider *OpenstackProvider) Name() string {
	return name
}

func (provider *OpenstackProvider) Description() string {
	return description
}

func (provider *OpenstackProvider) Author() string {
	return author
}

func (provider *OpenstackProvider) Version() string {
	return version
}
