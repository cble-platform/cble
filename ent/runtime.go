// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

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
	// blueprintDescCreatedAt is the schema descriptor for created_at field.
	blueprintDescCreatedAt := blueprintFields[1].Descriptor()
	// blueprint.DefaultCreatedAt holds the default value on creation for the created_at field.
	blueprint.DefaultCreatedAt = blueprintDescCreatedAt.Default.(func() time.Time)
	// blueprintDescUpdatedAt is the schema descriptor for updated_at field.
	blueprintDescUpdatedAt := blueprintFields[2].Descriptor()
	// blueprint.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	blueprint.DefaultUpdatedAt = blueprintDescUpdatedAt.Default.(func() time.Time)
	// blueprint.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	blueprint.UpdateDefaultUpdatedAt = blueprintDescUpdatedAt.UpdateDefault.(func() time.Time)
	// blueprintDescID is the schema descriptor for id field.
	blueprintDescID := blueprintFields[0].Descriptor()
	// blueprint.DefaultID holds the default value on creation for the id field.
	blueprint.DefaultID = blueprintDescID.Default.(func() uuid.UUID)
	deploymentFields := schema.Deployment{}.Fields()
	_ = deploymentFields
	// deploymentDescCreatedAt is the schema descriptor for created_at field.
	deploymentDescCreatedAt := deploymentFields[1].Descriptor()
	// deployment.DefaultCreatedAt holds the default value on creation for the created_at field.
	deployment.DefaultCreatedAt = deploymentDescCreatedAt.Default.(func() time.Time)
	// deploymentDescUpdatedAt is the schema descriptor for updated_at field.
	deploymentDescUpdatedAt := deploymentFields[2].Descriptor()
	// deployment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deployment.DefaultUpdatedAt = deploymentDescUpdatedAt.Default.(func() time.Time)
	// deployment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deployment.UpdateDefaultUpdatedAt = deploymentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// deploymentDescDescription is the schema descriptor for description field.
	deploymentDescDescription := deploymentFields[4].Descriptor()
	// deployment.DefaultDescription holds the default value on creation for the description field.
	deployment.DefaultDescription = deploymentDescDescription.Default.(string)
	// deploymentDescTemplateVars is the schema descriptor for template_vars field.
	deploymentDescTemplateVars := deploymentFields[5].Descriptor()
	// deployment.DefaultTemplateVars holds the default value on creation for the template_vars field.
	deployment.DefaultTemplateVars = deploymentDescTemplateVars.Default.(map[string]interface{})
	// deploymentDescDeploymentVars is the schema descriptor for deployment_vars field.
	deploymentDescDeploymentVars := deploymentFields[6].Descriptor()
	// deployment.DefaultDeploymentVars holds the default value on creation for the deployment_vars field.
	deployment.DefaultDeploymentVars = deploymentDescDeploymentVars.Default.(map[string]interface{})
	// deploymentDescDeploymentState is the schema descriptor for deployment_state field.
	deploymentDescDeploymentState := deploymentFields[7].Descriptor()
	// deployment.DefaultDeploymentState holds the default value on creation for the deployment_state field.
	deployment.DefaultDeploymentState = deploymentDescDeploymentState.Default.(map[string]string)
	// deploymentDescID is the schema descriptor for id field.
	deploymentDescID := deploymentFields[0].Descriptor()
	// deployment.DefaultID holds the default value on creation for the id field.
	deployment.DefaultID = deploymentDescID.Default.(func() uuid.UUID)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[1].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupFields[2].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionFields[1].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(func() time.Time)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionFields[2].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(func() time.Time)
	// permission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	permission.UpdateDefaultUpdatedAt = permissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionFields[0].Descriptor()
	// permission.DefaultID holds the default value on creation for the id field.
	permission.DefaultID = permissionDescID.Default.(func() uuid.UUID)
	permissionpolicyFields := schema.PermissionPolicy{}.Fields()
	_ = permissionpolicyFields
	// permissionpolicyDescCreatedAt is the schema descriptor for created_at field.
	permissionpolicyDescCreatedAt := permissionpolicyFields[1].Descriptor()
	// permissionpolicy.DefaultCreatedAt holds the default value on creation for the created_at field.
	permissionpolicy.DefaultCreatedAt = permissionpolicyDescCreatedAt.Default.(func() time.Time)
	// permissionpolicyDescUpdatedAt is the schema descriptor for updated_at field.
	permissionpolicyDescUpdatedAt := permissionpolicyFields[2].Descriptor()
	// permissionpolicy.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permissionpolicy.DefaultUpdatedAt = permissionpolicyDescUpdatedAt.Default.(func() time.Time)
	// permissionpolicy.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	permissionpolicy.UpdateDefaultUpdatedAt = permissionpolicyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// permissionpolicyDescIsInherited is the schema descriptor for is_inherited field.
	permissionpolicyDescIsInherited := permissionpolicyFields[4].Descriptor()
	// permissionpolicy.DefaultIsInherited holds the default value on creation for the is_inherited field.
	permissionpolicy.DefaultIsInherited = permissionpolicyDescIsInherited.Default.(bool)
	// permissionpolicyDescID is the schema descriptor for id field.
	permissionpolicyDescID := permissionpolicyFields[0].Descriptor()
	// permissionpolicy.DefaultID holds the default value on creation for the id field.
	permissionpolicy.DefaultID = permissionpolicyDescID.Default.(func() uuid.UUID)
	providerFields := schema.Provider{}.Fields()
	_ = providerFields
	// providerDescCreatedAt is the schema descriptor for created_at field.
	providerDescCreatedAt := providerFields[1].Descriptor()
	// provider.DefaultCreatedAt holds the default value on creation for the created_at field.
	provider.DefaultCreatedAt = providerDescCreatedAt.Default.(func() time.Time)
	// providerDescUpdatedAt is the schema descriptor for updated_at field.
	providerDescUpdatedAt := providerFields[2].Descriptor()
	// provider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	provider.DefaultUpdatedAt = providerDescUpdatedAt.Default.(func() time.Time)
	// provider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	provider.UpdateDefaultUpdatedAt = providerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// providerDescIsLoaded is the schema descriptor for is_loaded field.
	providerDescIsLoaded := providerFields[7].Descriptor()
	// provider.DefaultIsLoaded holds the default value on creation for the is_loaded field.
	provider.DefaultIsLoaded = providerDescIsLoaded.Default.(bool)
	// providerDescID is the schema descriptor for id field.
	providerDescID := providerFields[0].Descriptor()
	// provider.DefaultID holds the default value on creation for the id field.
	provider.DefaultID = providerDescID.Default.(func() uuid.UUID)
	providercommandFields := schema.ProviderCommand{}.Fields()
	_ = providercommandFields
	// providercommandDescCreatedAt is the schema descriptor for created_at field.
	providercommandDescCreatedAt := providercommandFields[1].Descriptor()
	// providercommand.DefaultCreatedAt holds the default value on creation for the created_at field.
	providercommand.DefaultCreatedAt = providercommandDescCreatedAt.Default.(func() time.Time)
	// providercommandDescUpdatedAt is the schema descriptor for updated_at field.
	providercommandDescUpdatedAt := providercommandFields[2].Descriptor()
	// providercommand.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	providercommand.DefaultUpdatedAt = providercommandDescUpdatedAt.Default.(func() time.Time)
	// providercommand.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	providercommand.UpdateDefaultUpdatedAt = providercommandDescUpdatedAt.UpdateDefault.(func() time.Time)
	// providercommandDescOutput is the schema descriptor for output field.
	providercommandDescOutput := providercommandFields[7].Descriptor()
	// providercommand.DefaultOutput holds the default value on creation for the output field.
	providercommand.DefaultOutput = providercommandDescOutput.Default.(string)
	// providercommandDescError is the schema descriptor for error field.
	providercommandDescError := providercommandFields[8].Descriptor()
	// providercommand.DefaultError holds the default value on creation for the error field.
	providercommand.DefaultError = providercommandDescError.Default.(string)
	// providercommandDescID is the schema descriptor for id field.
	providercommandDescID := providercommandFields[0].Descriptor()
	// providercommand.DefaultID holds the default value on creation for the id field.
	providercommand.DefaultID = providercommandDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
