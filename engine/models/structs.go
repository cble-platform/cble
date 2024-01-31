package models

import (
	"gopkg.in/yaml.v3"
)

type Blueprint struct {
	Version string             `yaml:"version" json:"version"`
	Objects map[string]*Object `yaml:",inline" json:"objects"`
}

type Object struct {
	Resource  string    `yaml:"resource" json:"resource"`
	Config    yaml.Node `yaml:"config" json:"config"`
	DependsOn []string  `yaml:"depends_on,omitempty" json:"depends_on,omitempty"`
}

type BlueprintVariableType string

const (
	BlueprintVariableType_STRING BlueprintVariableType = "STRING"
	BlueprintVariableType_INT    BlueprintVariableType = "INT"
)
