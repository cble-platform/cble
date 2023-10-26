package v1

import (
	"fmt"

	"github.com/cble-platform/backend/providers/openstack/internal"
)

const MAJOR_VERSION = "1"

func UnmarshalBlueprint(filepath string) (*OpenstackBlueprint, error) {
	blueprint, err := internal.UnmarshalBlueprint[OpenstackBlueprint](MAJOR_VERSION, filepath)
	if err != nil {
		return nil, err
	}
	err = unpackObjects(blueprint)
	if err != nil {
		return nil, err
	}
	return blueprint, nil
}

func UnmarshalBlueprintWithVars(filepath string, varFilepath string) (*OpenstackBlueprint, error) {
	blueprint, err := internal.UnmarshalBlueprintWithVars[OpenstackBlueprint](MAJOR_VERSION, filepath, varFilepath)
	if err != nil {
		return nil, err
	}
	err = unpackObjects(blueprint)
	if err != nil {
		return nil, err
	}
	return blueprint, nil
}

func ValidateBlueprint(blueprint *OpenstackBlueprint) error {
	for k, o := range blueprint.Objects {
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
	}
	return nil
}

func unpackObjects(blueprint *OpenstackBlueprint) error {
	// Initialize the native maps
	blueprint.Hosts = make(map[string]OpenstackHost)
	blueprint.Networks = make(map[string]OpenstackNetwork)
	blueprint.Routers = make(map[string]OpenstackRouter)

	for k, o := range blueprint.Objects {
		switch o.Resource {
		case OpenstackResourceTypeHost:
			if o.Host == nil {
				return fmt.Errorf("host \"%s\" has no host config", k)
			}
			blueprint.Hosts[k] = *o.Host
		case OpenstackResourceTypeNetwork:
			if o.Network == nil {
				return fmt.Errorf("network \"%s\" has no host config", k)
			}
			blueprint.Networks[k] = *o.Network
		case OpenstackResourceTypeRouter:
			if o.Router == nil {
				return fmt.Errorf("router \"%s\" has no host config", k)
			}
			blueprint.Routers[k] = *o.Router
		}
	}
	return nil
}
