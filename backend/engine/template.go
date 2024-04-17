package engine

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/cble-platform/cble/backend/ent"
	"gopkg.in/yaml.v3"
)

func templateObject(ctx context.Context, entDeploymentNode *ent.DeploymentNode) ([]byte, error) {
	// Get the resource
	entResource, err := entDeploymentNode.QueryResource().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query resource from node: %v", err)
	}
	// Get the deployment
	entDeployment, err := entDeploymentNode.QueryDeployment().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query deployment from node: %v", err)
	}

	// Convert the object to YAML
	objectBytes, err := yaml.Marshal(entResource.Object)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource (%s) object into YAML: %v", entResource.ID, err)
	}

	// Parse the object definition as a template
	t, err := template.New("object").Parse(string(objectBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to parse object as a template: %v", err)
	}

	// Execute the template and subsitute in deployment template variables
	var templatedObject bytes.Buffer
	if err := t.Execute(&templatedObject, entDeployment.TemplateVars); err != nil {
		return nil, fmt.Errorf("failed to template object definition: %v", err)
	}

	return templatedObject.Bytes(), nil
}
