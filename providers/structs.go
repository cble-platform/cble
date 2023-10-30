package providers

import "gopkg.in/yaml.v3"

type Blueprint struct {
	Version string `yaml:"version"`
}

type Object struct {
	Resource   string    `yaml:"resource"`
	Config     yaml.Node `yaml:"config"`
	DependsOn  []string  `yaml:"depends_on,omitempty"`
	RequiredBy []string  `yaml:"-"`
}
