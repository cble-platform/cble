package v1

import "fmt"

func ValidateBlueprint(blueprint *OpenstackBlueprint) error {
	for k, o := range blueprint.Objects {
		// Validate individual dependent types
		switch o.Resource {
		case OpenstackResourceTypeHost:
			if err := validateHost(blueprint, k); err != nil {
				return fmt.Errorf("invalid host \"%s\": %v", k, err)
			}
		case OpenstackResourceTypeNetwork:
			if err := validateNetwork(blueprint, k); err != nil {
				return fmt.Errorf("invalid network \"%s\": %v", k, err)
			}
		case OpenstackResourceTypeRouter:
			if err := validateRouter(blueprint, k); err != nil {
				return fmt.Errorf("invalid router \"%s\": %v", k, err)
			}
		}
		// Validate dependencies
		for _, d := range o.DependsOn {
			// Check that not self-dependent
			if d == k {
				return fmt.Errorf("object \"%s\" dependent on self", k)
			}
			// Check dependency is valid object
			if _, exists := blueprint.Objects[d]; !exists {
				return fmt.Errorf("object \"%s\" dependency \"%s\" is undefined", k, d)
			}
		}
	}
	return nil
}

func validateHost(blueprint *OpenstackBlueprint, key string) error {
	for networkKey, networkAttachment := range blueprint.Hosts[key].Networks {
		// Check that the network key we're attaching to is defined
		network, exists := blueprint.Networks[networkKey]
		if !exists {
			return fmt.Errorf("network object \"%s\" is not defined", networkKey)
		}
		// If not DHCP, check for valid IP address
		if !networkAttachment.DHCP && networkAttachment.IP != nil {
			if !network.Subnet.Contains(*networkAttachment.IP) {
				return fmt.Errorf("ip of %s on network \"%s\" (%s) is not valid", networkAttachment.IP, networkKey, network.Subnet)
			}
		}
	}
	return nil
}

func validateNetwork(blueprint *OpenstackBlueprint, key string) error {
	// Check that the gateway address is in the subnet
	if blueprint.Networks[key].Gateway != nil {
		if !blueprint.Networks[key].Subnet.Contains(*blueprint.Networks[key].Gateway) {
			return fmt.Errorf("gateway of %s not in subnet %s", blueprint.Networks[key].Gateway, blueprint.Networks[key].Subnet)
		}
	}
	// If DHCP ranges set, validate them
	if blueprint.Networks[key].DHCP != nil {
		for _, dhcp := range blueprint.Networks[key].DHCP {
			// Check the DHCP ranges are proper in subnet
			if !blueprint.Networks[key].Subnet.Contains(dhcp.Start) || !blueprint.Networks[key].Subnet.Contains(dhcp.End) {
				return fmt.Errorf("invalid dhcp range %s - %s: range not in subnet %s", dhcp.Start, dhcp.End, blueprint.Networks[key].Subnet)
			}
			rangeCmp := dhcp.Start.Compare(dhcp.End)
			// Check the start < end
			if rangeCmp == 0 {
				return fmt.Errorf("invalid dhcp range %s - %s: must contain at least 2 IP addresses", dhcp.Start, dhcp.End)
			}
			// Check the start < end
			if rangeCmp > 0 {
				return fmt.Errorf("invalid dhcp range %s - %s: start IP must come before end IP", dhcp.Start, dhcp.End)
			}
		}
	}
	return nil
}

func validateRouter(blueprint *OpenstackBlueprint, key string) error {
	for networkKey, networkAttachment := range blueprint.Routers[key].Networks {
		// Check that the network key we're attaching to is defined
		network, exists := blueprint.Networks[networkKey]
		if !exists {
			return fmt.Errorf("network object \"%s\" is not defined", networkKey)
		}
		// If not DHCP, check for valid IP address
		if !networkAttachment.DHCP && networkAttachment.IP != nil {
			if !network.Subnet.Contains(*networkAttachment.IP) {
				return fmt.Errorf("ip of %s on network \"%s\" (%s) is not valid", networkAttachment.IP, networkKey, network.Subnet)
			}
		}
	}
	return nil
}
