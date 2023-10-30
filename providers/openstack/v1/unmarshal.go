package v1

import (
	"fmt"

	"github.com/cble-platform/backend/providers/openstack/internal"
)

const MAJOR_VERSION = "1"

func UnmarshalBlueprintBytes(in []byte) (*OpenstackBlueprint, error) {
	blueprint, err := internal.UnmarshalBlueprintBytes[OpenstackBlueprint](MAJOR_VERSION, in)
	if err != nil {
		return nil, err
	}
	err = unpackObjects(blueprint)
	if err != nil {
		return nil, err
	}
	generateImpliedDependsOn(blueprint)
	return blueprint, nil
}

func UnmarshalBlueprintFile(filepath string) (*OpenstackBlueprint, error) {
	blueprint, err := internal.UnmarshalBlueprintFile[OpenstackBlueprint](MAJOR_VERSION, filepath)
	if err != nil {
		return nil, err
	}
	err = unpackObjects(blueprint)
	if err != nil {
		return nil, err
	}
	generateImpliedDependsOn(blueprint)
	return blueprint, nil
}

func UnmarshalBlueprintBytesWithVars(in []byte, vars map[string]interface{}) (*OpenstackBlueprint, error) {
	blueprint, err := internal.UnmarshalBlueprintBytesWithVars[OpenstackBlueprint](MAJOR_VERSION, in, vars)
	if err != nil {
		return nil, err
	}
	err = unpackObjects(blueprint)
	if err != nil {
		return nil, err
	}
	generateImpliedDependsOn(blueprint)
	return blueprint, nil
}

func UnmarshalBlueprintFileWithVars(filepath string, varFilepath string) (*OpenstackBlueprint, error) {
	blueprint, err := internal.UnmarshalBlueprintFileWithVars[OpenstackBlueprint](MAJOR_VERSION, filepath, varFilepath)
	if err != nil {
		return nil, err
	}
	err = unpackObjects(blueprint)
	if err != nil {
		return nil, err
	}
	generateImpliedDependsOn(blueprint)
	return blueprint, nil
}

func generateImpliedDependsOn(blueprint *OpenstackBlueprint) error {
	for ok, o := range blueprint.Objects {
		switch o.Resource {
		case OpenstackResourceTypeHost:
			// Add all networks host is on as depends_on
			for nk := range o.Host.Networks {
				o.DependsOn = append(o.DependsOn, nk)
				n := blueprint.Objects[nk]
				n.RequiredBy = append(n.RequiredBy, ok)
				blueprint.Objects[nk] = n
			}
			blueprint.Objects[ok] = o
		case OpenstackResourceTypeNetwork:
		case OpenstackResourceTypeRouter:
			// Add all networks host is on as depends_on
			for nk := range o.Router.Networks {
				o.DependsOn = append(o.DependsOn, nk)
				n := blueprint.Objects[nk]
				n.RequiredBy = append(n.RequiredBy, ok)
				blueprint.Objects[nk] = n
			}
			blueprint.Objects[ok] = o
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
