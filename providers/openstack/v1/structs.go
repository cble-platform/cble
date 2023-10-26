package v1

import (
	"fmt"
	"net/netip"

	"gopkg.in/yaml.v3"
)

type OpenstackResourceType string

const (
	OpenstackResourceTypeHost    OpenstackResourceType = "openstack.v1.host"
	OpenstackResourceTypeNetwork OpenstackResourceType = "openstack.v1.network"
	OpenstackResourceTypeRouter  OpenstackResourceType = "openstack.v1.router"
)

type OpenstackBlueprint struct {
	Version  string                     `yaml:"version"`
	Objects  map[string]OpenstackObject `yaml:",inline"`
	Hosts    map[string]OpenstackHost
	Networks map[string]OpenstackNetwork
	Routers  map[string]OpenstackRouter
}

type OpenstackObject struct {
	Resource OpenstackResourceType `yaml:"resource"`
	Config   yaml.Node             `yaml:"config"`
	Host     *OpenstackHost        `yaml:"-"`
	Network  *OpenstackNetwork     `yaml:"-"`
	Router   *OpenstackRouter      `yaml:"-"`
}

func (o *OpenstackObject) UnmarshalYAML(n *yaml.Node) error {
	type O OpenstackObject
	type T struct {
		*O `yaml:",inline"`
	}

	obj := &T{O: (*O)(o)}
	if err := n.Decode(obj); err != nil {
		return err
	}

	switch o.Resource {
	case OpenstackResourceTypeHost:
		o.Host = new(OpenstackHost)
		return obj.Config.Decode(o.Host)
	case OpenstackResourceTypeNetwork:
		o.Network = new(OpenstackNetwork)
		return obj.Config.Decode(o.Network)
	case OpenstackResourceTypeRouter:
		o.Router = new(OpenstackRouter)
		return obj.Config.Decode(o.Router)
	default:
		return fmt.Errorf("unknown resource type \"%s\"", o.Resource)
	}
}

type OpenstackHost struct {
	// Openstack instance name
	Name *string `yaml:"name,omitempty"`
	// Openstack instance description
	Description *string `yaml:"description,omitempty"`
	// Hostname of the host
	Hostname string `yaml:"hostname"`
	// Image of the host
	Image string `yaml:"image"`
	// Flavor of the host
	Flavor string `yaml:"flavor"`
	// Disk size of the host (in MB)
	DiskSize int `yaml:"disk_size"`
	// Networks to attach this host to
	Networks map[string]OpenstackNetworkAttachment `yaml:"networks"`
}

type OpenstackNetworkAttachment struct {
	// Should this interface get IP via DHCP (overrides IP setting if set)
	DHCP bool `yaml:"dhcp,omitempty"`
	// IPv4 address to use for the interface
	IP netip.Addr `yaml:"ip,omitempty"`
}

type OpenstackNetwork struct {
	// The subnet CIDR for the network
	Subnet netip.Prefix `yaml:"subnet"`
	// The gateway for the network
	Gateway netip.Addr `yaml:"gateway,omitempty"`
	// DHCP ranges for the network (omit to disable DHCP)
	DHCP []OpenstackNetworkDHCP `yaml:"dhcp,omitempty"`
}

type OpenstackNetworkDHCP struct {
	// The start IP address for the DHCP range
	Start netip.Addr `yaml:"start"`
	// The end IP address for the DHCP range
	End netip.Addr `yaml:"end"`
}

type OpenstackRouter struct {
	// Openstack instance name
	Name *string `yaml:"name,omitempty"`
	// The ID or Name of the external Openstack network to attach this router to
	ExternalNetwork string `yaml:"external_network"`
	// Networks to attach this host to
	Networks map[string]OpenstackNetworkAttachment `yaml:"networks"`
}
