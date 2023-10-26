package v1

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

func (provider OpenstackProvider) newAuthClient() (*gophercloud.ProviderClient, error) {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: provider.config.AuthUrl,
		Username:         provider.config.Username,
		Password:         provider.config.Password,
		TenantID:         provider.config.ProjectID,
		TenantName:       provider.config.ProjectName,
	}
	if provider.config.DomainName != "" {
		authOpts.DomainName = provider.config.DomainName
	} else {
		authOpts.DomainID = provider.config.DomainId
	}
	return openstack.AuthenticatedClient(authOpts)
}
