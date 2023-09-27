package permissionengine

import (
	"context"
	"fmt"

	"github.com/cble-platform/backend/ent"
	"github.com/cble-platform/backend/ent/group"
	"github.com/cble-platform/backend/ent/permission"
	"github.com/cble-platform/backend/ent/permissionpolicy"
	"github.com/cble-platform/backend/internal/utils"
)

type PermissionEngine struct {
	ent *ent.Client
}

func New(entClient *ent.Client) (*PermissionEngine, error) {
	if entClient == nil {
		return nil, fmt.Errorf("ent client must not be nil")
	}
	return &PermissionEngine{
		ent: entClient,
	}, nil
}

func (pe *PermissionEngine) RegisterPermission(ctx context.Context, key string, component string, description string) (*ent.Permission, error) {
	entPermission, err := pe.ent.Permission.Query().Where(permission.KeyEQ(key)).Only(ctx)
	if ent.IsNotFound(err) {
		// Permission has yet to be registered, create it
		entPermission, err = pe.ent.Permission.Create().
			SetKey(key).
			SetComponent(component).
			SetDescription(description).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to register permission: %v", err)
		}
		return entPermission, nil
	} else if err != nil {
		// Some other error happened
		return nil, err
	} else if entPermission.Component != component {
		// Some other component registered this same permission, return an error
		return nil, fmt.Errorf("failed to register permission: component \"%s\" has already registered permission \"%s\"", entPermission.Component, entPermission.Key)
	} else {
		// This permission has laready been registered by this component, so return it
		return entPermission, nil
	}
}

func (pe *PermissionEngine) SetPermissionPolicy(ctx context.Context, policyType permissionpolicy.Type, entPermission *ent.Permission, entGroup *ent.Group) (*ent.PermissionPolicy, error) {
	// Create a transactional client
	tx, err := pe.ent.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactional client: %v", err)
	}
	txClient := tx.Client()

	// Search if this policy already exists
	entPermissionPolicy, err := txClient.PermissionPolicy.Query().Where(
		permissionpolicy.And(
			permissionpolicy.HasGroupWith(group.IDEQ(entGroup.ID)),
			permissionpolicy.HasPermissionWith(permission.IDEQ(entPermission.ID)),
		),
	).Only(ctx)
	if ent.IsNotFound(err) {
		// If not found, create the policy
		entPermissionPolicy, err = txClient.PermissionPolicy.Create().
			SetType(policyType).
			SetGroup(entGroup).
			SetPermission(entPermission).
			Save(ctx)
		if err != nil {
			err = fmt.Errorf("failed to create permission policy: %v", err)
			return nil, utils.RollbackWithErr(tx, err)
		}
	} else if err != nil {
		// Some other error happened
		err = fmt.Errorf("failed to query for existing permission policy: %v", err)
		return nil, utils.RollbackWithErr(tx, err)
	} else {
		// Check if the policy in database matches the requested policy type
		if entPermissionPolicy.Type != policyType {
			// Update the policy to match the requested policy type
			entPermissionPolicy, err = entPermissionPolicy.Update().SetType(policyType).Save(ctx)
			if err != nil {
				err = fmt.Errorf("failed to update existing permission policy: %v", err)
				return nil, utils.RollbackWithErr(tx, err)
			}
		}
		// Convert from an inherited policy if previously inherited
		if entPermissionPolicy.IsInherited {
			entPermissionPolicy, err = entPermissionPolicy.Update().SetIsInherited(false).Save(ctx)
			if err != nil {
				err = fmt.Errorf("failed to update existing permission policy: %v", err)
				return nil, utils.RollbackWithErr(tx, err)
			}
		}
	}

	// Inherit this policy to all of the group's children
	entChildren, err := entGroup.QueryChildren().All(ctx)
	if err != nil {
		err = fmt.Errorf("failed to query group children: %v", err)
		return nil, utils.RollbackWithErr(tx, err)
	}
	for _, entChild := range entChildren {
		err = inheritPermissionPolicy(ctx, txClient, policyType, entPermission, entChild)
		if err != nil {
			err = fmt.Errorf("failed to inherit permission policy to children: %v", err)
			return nil, utils.RollbackWithErr(tx, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("failed to commit transaction: %v", err)
		return nil, utils.RollbackWithErr(tx, err)
	}

	return entPermissionPolicy, nil
}

func inheritPermissionPolicy(ctx context.Context, txClient *ent.Client, policyType permissionpolicy.Type, entPermission *ent.Permission, entGroup *ent.Group) error {
	// See if a policy is already set for this permission on this group
	entPermissionPolicy, err := txClient.PermissionPolicy.Query().Where(
		permissionpolicy.And(
			permissionpolicy.HasGroupWith(group.IDEQ(entGroup.ID)),
			permissionpolicy.HasPermissionWith(permission.IDEQ(entPermission.ID)),
		),
	).Only(ctx)
	if ent.IsNotFound(err) {
		// If the policy doesn't exist yet, create it and set it as inherited
		entPermissionPolicy, err = txClient.PermissionPolicy.Create().
			SetType(policyType).
			SetIsInherited(true).
			SetPermission(entPermission).
			SetGroup(entGroup).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create permission policy: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to query for permission policy: %v", err)
	}

	// If it exists, check if it is already an inherited policy
	if entPermissionPolicy.IsInherited {
		// If inherited, check the policy type is set properly
		if entPermissionPolicy.Type != policyType {
			// If not, set it properly
			err = entPermissionPolicy.Update().SetType(policyType).Exec(ctx)
			if err != nil {
				return fmt.Errorf("failed to update existing inherited permission policy: %v", err)
			}
		}

		// Inherit this policy to all of this group's children

		// iterate over the childern of this group and call self on it
		entChildGroups, err := entGroup.QueryChildren().All(ctx)
		if err != nil {
			return fmt.Errorf("failed to query group children")
		}
		for _, entChild := range entChildGroups {
			err = inheritPermissionPolicy(ctx, txClient, policyType, entPermission, entChild)
			if err != nil {
				return fmt.Errorf("failed to set permissions on child: %v", err)
			}
		}
	}
	// If not inherited, leave it alone as it's overridden and don't propogate to children

	return nil
}

func (pe *PermissionEngine) RequestPermission(ctx context.Context, entUser *ent.User, permissionKey string) (bool, error) {
	// Check if user is a part of any groups which have (or inherit) an ALLOW permission policy for this permission
	isAllowed, err := entUser.QueryGroups().Where(
		group.HasPermissionPoliciesWith(
			permissionpolicy.And(
				permissionpolicy.HasPermissionWith(permission.KeyEQ(permissionKey)),
				permissionpolicy.TypeEQ(permissionpolicy.TypeALLOW),
			),
		),
	).Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to query for permission: %v", err)
	}
	return isAllowed, nil
}
