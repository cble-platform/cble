package v1

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cble-platform/backend/ent"
	"github.com/cble-platform/backend/providers"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	"github.com/sirupsen/logrus"
)

func (provider *OpenstackProvider) DeployBlueprint(ctx context.Context, client *ent.Client, entRequester *ent.User, entBlueprint *ent.Blueprint, templateVars map[string]interface{}) error {
	// Create a deployment for this blueprint
	entDeployment, err := client.Deployment.Create().
		SetTemplateVars(templateVars).
		SetBlueprint(entBlueprint).
		SetRequester(entRequester).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create deployment: %v", err)
	}

	// Generate authenticated client session
	authClient, err := provider.newAuthClient()
	if err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	// Parse blueprint into struct
	blueprint, err := UnmarshalBlueprintBytesWithVars(entBlueprint.BlueprintTemplate, templateVars)
	if err != nil {
		return fmt.Errorf("failed to unmarshal blueprint: %v", err)
	}

	// Validate the blueprint is valid
	err = ValidateBlueprint(blueprint)
	if err != nil {
		return fmt.Errorf("blueprint is invalid: %v", err)
	}

	objectsWg := sync.WaitGroup{}
	for k := range blueprint.Objects {
		objectsWg.Add(1)
		go func(key string) {
			// Wait until all depends_on are done
			err := awaitDependsOn(entDeployment, blueprint, key)
			if err != nil {
				logrus.Errorf("failed to deploy network: %v", err)
			} else {
				switch blueprint.Objects[key].Resource {
				case OpenstackResourceTypeHost:
					if err := provider.deployHost(ctx, client, authClient, entDeployment, blueprint, key); err != nil {
						logrus.Errorf("failed to deploy host \"%s\": %v", key, err)
					}
				case OpenstackResourceTypeNetwork:
					if err := provider.deployNetwork(ctx, client, authClient, entDeployment, blueprint, key); err != nil {
						logrus.Errorf("failed to deploy network \"%s\": %v", key, err)
					}
				case OpenstackResourceTypeRouter:
					if err := provider.deployRouter(ctx, client, authClient, entDeployment, blueprint, key); err != nil {
						logrus.Errorf("failed to deploy router \"%s\": %v", key, err)
					}
				}
			}
			objectsWg.Done()
		}(k)
	}
	objectsWg.Wait()

	// // Deploy Networks
	// networksWg := sync.WaitGroup{}
	// for k := range blueprint.Networks {
	// 	networksWg.Add(1)
	// 	go func(key string) {
	// 		// Wait until all depends_on are done
	// 		err := awaitDependsOn(entDeployment, blueprint, key)
	// 		if err != nil {
	// 			logrus.Errorf("failed to deploy network: %v", err)
	// 		} else if err := provider.deployNetwork(ctx, client, authClient, entDeployment, blueprint, key); err != nil {
	// 			logrus.Errorf("failed to deploy network \"%s\": %v", key, err)
	// 		}
	// 		networksWg.Done()
	// 	}(k)
	// }
	// networksWg.Wait()
	// // Deploy Routers
	// routersWg := sync.WaitGroup{}
	// for k := range blueprint.Routers {
	// 	routersWg.Add(1)
	// 	go func(key string) {
	// 		// Wait until all depends_on are done
	// 		err := awaitDependsOn(entDeployment, blueprint, key)
	// 		if err != nil {
	// 			logrus.Errorf("failed to deploy router: %v", err)
	// 		} else if err := provider.deployRouter(ctx, client, authClient, entDeployment, blueprint, key); err != nil {
	// 			logrus.Errorf("failed to deploy router \"%s\": %v", key, err)
	// 		}
	// 		routersWg.Done()
	// 	}(k)
	// }
	// routersWg.Wait()
	// // Deploy Hosts
	// hostsWg := sync.WaitGroup{}
	// for k := range blueprint.Hosts {
	// 	hostsWg.Add(1)
	// 	go func(key string) {
	// 		// Wait until all depends_on are done
	// 		err := awaitDependsOn(entDeployment, blueprint, key)
	// 		if err != nil {
	// 			logrus.Errorf("failed to deploy host: %v", err)
	// 		} else if err := provider.deployHost(ctx, client, authClient, entDeployment, blueprint, key); err != nil {
	// 			logrus.Errorf("failed to deploy host \"%s\": %v", key, err)
	// 		}
	// 		hostsWg.Done()
	// 	}(k)
	// }
	// hostsWg.Wait()
	return nil
}

func (provider *OpenstackProvider) deployNetwork(ctx context.Context, client *ent.Client, authClient *gophercloud.ProviderClient, entDeployment *ent.Deployment, blueprint *OpenstackBlueprint, networkKey string) error {
	// Get the network from blueprint
	network, exist := blueprint.Networks[networkKey]
	if !exist {
		return fmt.Errorf("network \"%s\" is not defined", networkKey)
	}

	// Generate the Network V2 client
	endpointOpts := gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: provider.config.RegionName,
	}
	networkClient, err := openstack.NewNetworkV2(authClient, endpointOpts)
	if err != nil {
		return fmt.Errorf("failed to create openstack network client: %v", err)
	}

	networkName := networkKey
	if network.Name != nil {
		networkName = *network.Name
	}

	// Create the network
	deployedNetwork, err := networks.Create(networkClient, networks.CreateOpts{
		Name:         networkName,
		AdminStateUp: gophercloud.Enabled,
	}).Extract()
	if err != nil {
		return fmt.Errorf("failed to create network: %v", err)
	}

	// Save the deployed network id into vars
	entDeployment.DeploymentVars[networkKey+"_id"] = deployedNetwork.ID
	err = entDeployment.Update().SetDeploymentVars(entDeployment.DeploymentVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save deployment vars: %v", err)
	}

	// Configure the subnet on the network
	var gatewayIp *string = nil
	if network.Gateway != nil {
		gatewayString := network.Gateway.String()
		gatewayIp = &gatewayString
	}
	dhcpPools := []subnets.AllocationPool{}
	for _, dhcp := range network.DHCP {
		dhcpPools = append(dhcpPools, subnets.AllocationPool{
			Start: dhcp.Start.String(),
			End:   dhcp.End.String(),
		})
	}
	dnsServers := []string{}
	for _, resolverIP := range network.Resolvers {
		dnsServers = append(dnsServers, resolverIP.String())
	}

	// Create openstack subnet on network
	deployedSubnet, err := subnets.Create(networkClient, subnets.CreateOpts{
		NetworkID:       deployedNetwork.ID,
		CIDR:            network.Subnet.String(),
		Name:            networkName,
		Description:     fmt.Sprintf("%s Subnet for Network \"%s\"", network.Subnet.String(), networkName),
		AllocationPools: dhcpPools,
		GatewayIP:       gatewayIp,
		IPVersion:       gophercloud.IPv4,
		EnableDHCP:      gophercloud.Enabled,
		DNSNameservers:  dnsServers,
	}).Extract()
	if err != nil {
		return fmt.Errorf("failed to create subnet: %v", err)
	}

	// Save the deployed network subnet id into vars
	entDeployment.DeploymentVars[networkKey+"_subnet_id"] = deployedSubnet.ID
	err = entDeployment.Update().SetDeploymentVars(entDeployment.DeploymentVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save deployment vars: %v", err)
	}

	// Set network as active for dependencies
	entDeployment.IsActive[networkKey] = providers.DeploySUCCEEDED
	err = entDeployment.Update().SetIsActive(entDeployment.IsActive).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save is_active map: %s", err)
	}

	logrus.Debugf("Successfully deployed network %s as network %s (%s)", networkKey, deployedNetwork.Name, deployedNetwork.ID)
	return nil
}

func (provider *OpenstackProvider) deployRouter(ctx context.Context, client *ent.Client, authClient *gophercloud.ProviderClient, entDeployment *ent.Deployment, blueprint *OpenstackBlueprint, routerKey string) error {
	// Get the router from blueprint
	router, exist := blueprint.Routers[routerKey]
	if !exist {
		return fmt.Errorf("router \"%s\" is not defined", routerKey)
	}

	// Generate the Network V2 client
	endpointOpts := gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: provider.config.RegionName,
	}
	networkClient, err := openstack.NewNetworkV2(authClient, endpointOpts)
	if err != nil {
		return fmt.Errorf("failed to create openstack network client: %v", err)
	}

	// Find the external network
	var routerExternalNetwork *networks.Network = nil
	allNetworkPages, err := networks.List(networkClient, nil).AllPages()
	if err != nil {
		return fmt.Errorf("failed to get router external network \"%s\": %v", router.ExternalNetwork, err)
	}
	allNetworks, err := networks.ExtractNetworks(allNetworkPages)
	if err != nil {
		return fmt.Errorf("failed to get router external network \"%s\": %v", router.ExternalNetwork, err)
	}
	for _, net := range allNetworks {
		if net.Name == router.ExternalNetwork || net.ID == router.ExternalNetwork {
			routerExternalNetwork = &net
			break
		}
	}
	if routerExternalNetwork == nil {
		return fmt.Errorf("failed to get router external network \"%s\": network not found", router.ExternalNetwork)
	}

	routerConfig := routers.CreateOpts{
		Name:         routerKey,
		Description:  "",
		AdminStateUp: gophercloud.Enabled,
		GatewayInfo: &routers.GatewayInfo{
			NetworkID: routerExternalNetwork.ID,
		},
	}
	if router.Name != nil {
		routerConfig.Name = *router.Name
	}
	if router.Description != nil {
		routerConfig.Description = *router.Description
	}

	// Deploy the router
	deployedRouter, err := routers.Create(networkClient, routerConfig).Extract()
	if err != nil {
		return fmt.Errorf("failed to create router: %v", err)
	}

	// Save the deployed router into vars
	entDeployment.DeploymentVars[routerKey+"_id"] = deployedRouter.ID
	err = entDeployment.Update().SetDeploymentVars(entDeployment.DeploymentVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save deployment vars: %v", err)
	}

	// Connect router to all attached networks

	for k, networkAttachment := range router.Networks {
		networkId, exists := entDeployment.DeploymentVars[k+"_id"]
		if !exists {
			return fmt.Errorf("ID unknown for network \"%s\"", k)
		}
		networkSubnetId, exists := entDeployment.DeploymentVars[k+"_subnet_id"]
		if !exists {
			return fmt.Errorf("ID unknown for network \"%s\" subnet", k)
		}
		// Create Openstack port for router on subnet
		osPort, err := ports.Create(networkClient, ports.CreateOpts{
			NetworkID:    networkId.(string),
			AdminStateUp: gophercloud.Enabled,
			FixedIPs: []ports.IP{{
				SubnetID:  networkSubnetId.(string),
				IPAddress: networkAttachment.IP.String(),
			}},
		}).Extract()
		if err != nil {
			return fmt.Errorf("failed to create port for router: %v", err)
		}

		// Save the deployed router network port into vars
		entDeployment.DeploymentVars[routerKey+"_"+k+"_port_id"] = osPort.ID
		err = entDeployment.Update().SetDeploymentVars(entDeployment.DeploymentVars).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to save deployment vars: %v", err)
		}

		osInterface, err := routers.AddInterface(networkClient, deployedRouter.ID, routers.AddInterfaceOpts{
			PortID: osPort.ID,
		}).Extract()
		if err != nil {
			return fmt.Errorf("failed to create router interface: %v", err)
		}

		// Save the deployed router network interface into vars
		entDeployment.DeploymentVars[routerKey+"_"+k+"_interface_id"] = osInterface.ID
		err = entDeployment.Update().SetDeploymentVars(entDeployment.DeploymentVars).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to save deployment vars: %v", err)
		}
	}

	// Set network as active for dependencies
	entDeployment.IsActive[routerKey] = providers.DeploySUCCEEDED
	err = entDeployment.Update().SetIsActive(entDeployment.IsActive).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save is_active map: %s", err)
	}

	logrus.Debugf("Successfully deployed router %s as router %s (%s)", routerKey, deployedRouter.Name, deployedRouter.ID)
	return nil
}

func (provider *OpenstackProvider) deployHost(ctx context.Context, client *ent.Client, authClient *gophercloud.ProviderClient, entDeployment *ent.Deployment, blueprint *OpenstackBlueprint, hostKey string) error {
	// Get the host from blueprint
	host, exist := blueprint.Hosts[hostKey]
	if !exist {
		return fmt.Errorf("host \"%s\" is not defined", hostKey)
	}

	// Generate the Compute V2 client
	endpointOpts := gophercloud.EndpointOpts{
		Region: provider.config.RegionName,
	}
	computeClient, err := openstack.NewComputeV2(authClient, endpointOpts)
	if err != nil {
		return fmt.Errorf("failed to create compute v2 client: %v", err)
	}

	var hostFlavor *flavors.Flavor = nil
	allFlavorPages, err := flavors.ListDetail(computeClient, nil).AllPages()
	if err != nil {
		return fmt.Errorf("failed to get host flavor \"%s\": %v", host.Flavor, err)
	}
	allFlavors, err := flavors.ExtractFlavors(allFlavorPages)
	if err != nil {
		return fmt.Errorf("failed to get host flavor \"%s\": %v", host.Flavor, err)
	}
	for _, fl := range allFlavors {
		if fl.Name == host.Flavor || fl.ID == host.Flavor {
			hostFlavor = &fl
			break
		}
	}
	if hostFlavor == nil {
		return fmt.Errorf("failed to get host flavor \"%s\": flavor not found", host.Flavor)
	}

	logrus.Debugf("got flavor %s (%s)", hostFlavor.Name, hostFlavor.ID)

	var hostImage *images.Image = nil
	allImagePages, err := images.ListDetail(computeClient, nil).AllPages()
	if err != nil {
		return fmt.Errorf("failed to get host image \"%s\": %v", host.Image, err)
	}
	allImages, err := images.ExtractImages(allImagePages)
	if err != nil {
		return fmt.Errorf("failed to get host image \"%s\": %v", host.Image, err)
	}
	for _, img := range allImages {
		if img.Name == host.Image || img.ID == host.Image {
			hostImage = &img
			break
		}
	}
	if hostImage == nil {
		return fmt.Errorf("failed to get host image \"%s\": image not found", host.Image)
	}

	// Check if the image requires more space than provided
	if host.DiskSize < hostImage.MinDisk {
		return fmt.Errorf("host disk size is too small for image (minimum %dGB required)", hostImage.MinDisk)
	}

	logrus.Debugf("got image %s (%s)", hostImage.Name, hostImage.ID)

	// Use either key or provided name as instance name
	instanceName := host.Hostname
	if host.Name != nil {
		instanceName = *host.Name
	}

	// Configure the volume to clone from the image
	blockOps := []bootfromvolume.BlockDevice{
		{
			UUID:                hostImage.ID,
			BootIndex:           0,
			DeleteOnTermination: true,
			DestinationType:     bootfromvolume.DestinationVolume,
			SourceType:          bootfromvolume.SourceImage,
			VolumeSize:          host.DiskSize,
		},
	}

	// Create network mappings for each network attachment
	hostNetworks := []servers.Network{}
	for k, networkAttachment := range host.Networks {
		_, exists := blueprint.Networks[k]
		if !exists {
			return fmt.Errorf("network \"%s\" is not defined", k)
		}
		networkId, exists := entDeployment.DeploymentVars[k+"_id"]
		if !exists {
			return fmt.Errorf("ID unknown for network \"%s\"", k)
		}
		PLACEHOLDER_NET := servers.Network{
			UUID: networkId.(string),
		}
		if !networkAttachment.DHCP && networkAttachment.IP != nil {
			PLACEHOLDER_NET.FixedIP = networkAttachment.IP.String()
		}
		hostNetworks = append(hostNetworks, PLACEHOLDER_NET)
	}

	// Configure the instance options
	hostOps := servers.CreateOpts{
		Name:      instanceName,
		ImageRef:  hostImage.ID,
		FlavorRef: hostFlavor.ID,
		UserData:  host.UserData,
		Networks:  hostNetworks,
	}

	// Create the host
	createOpts := bootfromvolume.CreateOptsExt{
		CreateOptsBuilder: hostOps,
		BlockDevice:       blockOps,
	}
	deployedServer, err := bootfromvolume.Create(computeClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("failed to deploy host: %v", err)
	}

	// Save the deployed host into vars
	entDeployment.DeploymentVars[hostKey+"_id"] = deployedServer.ID
	err = entDeployment.Update().SetDeploymentVars(entDeployment.DeploymentVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save deployment vars: %v", err)
	}

	// Wait for server to be in ACTIVE state
	for {
		// Get the updated server from Openstack
		deployedServer, err = servers.Get(computeClient, deployedServer.ID).Extract()
		if err != nil {
			return fmt.Errorf("failed to get openstack server status: %v", err)
		}
		if deployedServer.Status == "ERROR" {
			// Something happened and this failed
			return fmt.Errorf("failed to deploy host: server in ERROR state")
		}
		if deployedServer.Status == "ACTIVE" {
			// Server deployed properly
			break
		}
		// Wait 5 seconds before checking again
		time.Sleep(5 * time.Second)
	}

	// Set host as active for dependencies
	entDeployment.IsActive[hostKey] = providers.DeploySUCCEEDED
	err = entDeployment.Update().SetIsActive(entDeployment.IsActive).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to save is_active map: %s", err)
	}

	logrus.Debugf("Successfully deployed host %s as server %s (%s)", hostKey, deployedServer.Name, deployedServer.ID)
	return nil
}

// Blocks execution until all depends_on are done.
func awaitDependsOn(entDeployment *ent.Deployment, blueprint *OpenstackBlueprint, key string) error {
	// Check on dependencies
	for {
		waitingOnDependents := false
		for _, dependsOnKey := range blueprint.Objects[key].DependsOn {
			dependsOnIsActive, exists := entDeployment.IsActive[dependsOnKey]
			if !exists {
				logrus.Debugf("\"%s\" is waiting on \"%s\"", key, dependsOnKey)
				waitingOnDependents = true
				// early break since no need to check others if a single dependency is still inactive
				break
			}
			if dependsOnIsActive == providers.DeployFAILED {
				return fmt.Errorf("\"%s\" dependency \"%s\" failed", key, dependsOnKey)
			}
		}
		// If all depends on objects are done
		if !waitingOnDependents {
			break
		}
		// Wait 5 secs before checking again
		time.Sleep(5 * time.Second)
	}
	return nil
}
