package internal

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func UnmarshalBlueprintBytes[T interface{}](majorVersion string, blueprintBytes []byte) (*T, error) {
	// Check the major version of the blueprint matches
	var genericBlueprint OpenstackGenericBlueprint
	if err := yaml.Unmarshal(blueprintBytes, &genericBlueprint); err != nil {
		return nil, fmt.Errorf("failed to unmarshal generic blueprint: %v", err)
	}
	if strings.Split(genericBlueprint.Version, ".")[0] != majorVersion {
		return nil, fmt.Errorf("found version \"%s\" when expected version 1.x", genericBlueprint.Version)
	}
	// Unmarshal the blueprint YAML
	var blueprint T
	if err := yaml.Unmarshal(blueprintBytes, &blueprint); err != nil {
		return nil, fmt.Errorf("failed to unmarshal blueprint: %v", err)
	}
	return &blueprint, nil
}

func UnmarshalBlueprintFile[T interface{}](majorVersion string, filepath string) (*T, error) {
	// Open the blueprint file
	blueprintFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	// Read into a buffer for unmarshalling
	blueprintBytes, err := io.ReadAll(blueprintFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read blueprint file: %v", err)
	}
	return UnmarshalBlueprintBytes[T](majorVersion, blueprintBytes)
}

func UnmarshalBlueprintBytesWithVars[T interface{}](majorVersion string, blueprintBytes []byte, templateVars map[string]interface{}) (*T, error) {
	/// Parse the blueprint as a template
	t, err := template.New("blueprint").Parse(string(blueprintBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to parse blueprint bytes into template: %v", err)
	}
	// Execute the template to subsitute variables
	var blueprintTemplated bytes.Buffer
	if err := t.Execute(&blueprintTemplated, templateVars); err != nil {
		return nil, fmt.Errorf("failed to template vars into blueprint: %v", err)
	}
	return UnmarshalBlueprintBytes[T](majorVersion, blueprintTemplated.Bytes())
}

func UnmarshalBlueprintFileWithVars[T interface{}](majorVersion string, filepath string, varFilepath string) (*T, error) {
	// Open the vars file
	varsFile, err := os.Open(varFilepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open vars file: %v", err)
	}
	// Read vars into a buffer for unmarshalling
	varsBytes, err := io.ReadAll(varsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read vars file: %v", err)
	}
	// Marshal vars into a template context
	blueprintContext := make(map[string]interface{})
	if err := yaml.Unmarshal(varsBytes, blueprintContext); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vars file: %v", err)
	}
	// Parse the blueprint as a template
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse blueprint file into template: %v", err)
	}
	// Execute the template to subsitute variables
	var blueprintTemplated bytes.Buffer
	if err := t.Execute(&blueprintTemplated, blueprintContext); err != nil {
		return nil, fmt.Errorf("failed to template vars into blueprint: %v", err)
	}
	return UnmarshalBlueprintBytes[T](majorVersion, blueprintTemplated.Bytes())
}
