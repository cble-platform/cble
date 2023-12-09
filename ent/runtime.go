// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/cble-platform/cble-backend/ent/blueprint"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/group"
	"github.com/cble-platform/cble-backend/ent/permission"
	"github.com/cble-platform/cble-backend/ent/permissionpolicy"
	"github.com/cble-platform/cble-backend/ent/provider"
	"github.com/cble-platform/cble-backend/ent/providercommand"
	"github.com/cble-platform/cble-backend/ent/schema"
	"github.com/cble-platform/cble-backend/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blueprintFields := schema.Blueprint{}.Fields()
	_ = blueprintFields
	// blueprintDescID is the schema descriptor for id field.
	blueprintDescID := blueprintFields[0].Descriptor()
	// blueprint.DefaultID holds the default value on creation for the id field.
	blueprint.DefaultID = blueprintDescID.Default.(func() uuid.UUID)
	deploymentFields := schema.Deployment{}.Fields()
	_ = deploymentFields
	// deploymentDescTemplateVars is the schema descriptor for template_vars field.
	deploymentDescTemplateVars := deploymentFields[2].Descriptor()
	// deployment.DefaultTemplateVars holds the default value on creation for the template_vars field.
	deployment.DefaultTemplateVars = deploymentDescTemplateVars.Default.(map[string]interface{})
	// deploymentDescDeploymentVars is the schema descriptor for deployment_vars field.
	deploymentDescDeploymentVars := deploymentFields[3].Descriptor()
	// deployment.DefaultDeploymentVars holds the default value on creation for the deployment_vars field.
	deployment.DefaultDeploymentVars = deploymentDescDeploymentVars.Default.(map[string]interface{})
	// deploymentDescDeploymentState is the schema descriptor for deployment_state field.
	deploymentDescDeploymentState := deploymentFields[4].Descriptor()
	// deployment.DefaultDeploymentState holds the default value on creation for the deployment_state field.
	deployment.DefaultDeploymentState = deploymentDescDeploymentState.Default.(map[string]string)
	// deploymentDescID is the schema descriptor for id field.
	deploymentDescID := deploymentFields[0].Descriptor()
	// deployment.DefaultID holds the default value on creation for the id field.
	deployment.DefaultID = deploymentDescID.Default.(func() uuid.UUID)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionFields[0].Descriptor()
	// permission.DefaultID holds the default value on creation for the id field.
	permission.DefaultID = permissionDescID.Default.(func() uuid.UUID)
	permissionpolicyFields := schema.PermissionPolicy{}.Fields()
	_ = permissionpolicyFields
	// permissionpolicyDescIsInherited is the schema descriptor for is_inherited field.
	permissionpolicyDescIsInherited := permissionpolicyFields[2].Descriptor()
	// permissionpolicy.DefaultIsInherited holds the default value on creation for the is_inherited field.
	permissionpolicy.DefaultIsInherited = permissionpolicyDescIsInherited.Default.(bool)
	// permissionpolicyDescID is the schema descriptor for id field.
	permissionpolicyDescID := permissionpolicyFields[0].Descriptor()
	// permissionpolicy.DefaultID holds the default value on creation for the id field.
	permissionpolicy.DefaultID = permissionpolicyDescID.Default.(func() uuid.UUID)
	providerFields := schema.Provider{}.Fields()
	_ = providerFields
	// providerDescIsLoaded is the schema descriptor for is_loaded field.
	providerDescIsLoaded := providerFields[5].Descriptor()
	// provider.DefaultIsLoaded holds the default value on creation for the is_loaded field.
	provider.DefaultIsLoaded = providerDescIsLoaded.Default.(bool)
	// providerDescID is the schema descriptor for id field.
	providerDescID := providerFields[0].Descriptor()
	// provider.DefaultID holds the default value on creation for the id field.
	provider.DefaultID = providerDescID.Default.(func() uuid.UUID)
	providercommandFields := schema.ProviderCommand{}.Fields()
	_ = providercommandFields
	// providercommandDescOutput is the schema descriptor for output field.
	providercommandDescOutput := providercommandFields[5].Descriptor()
	// providercommand.DefaultOutput holds the default value on creation for the output field.
	providercommand.DefaultOutput = providercommandDescOutput.Default.(string)
	// providercommandDescError is the schema descriptor for error field.
	providercommandDescError := providercommandFields[6].Descriptor()
	// providercommand.DefaultError holds the default value on creation for the error field.
	providercommand.DefaultError = providercommandDescError.Default.(string)
	// providercommandDescID is the schema descriptor for id field.
	providercommandDescID := providercommandFields[0].Descriptor()
	// providercommand.DefaultID holds the default value on creation for the id field.
	providercommand.DefaultID = providercommandDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
