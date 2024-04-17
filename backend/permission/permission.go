package permission

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble/backend/auth"
	"github.com/cble-platform/cble/backend/ent"
	"github.com/cble-platform/cble/backend/ent/grantedpermission"
	"github.com/cble-platform/cble/backend/permission/actions"
	"github.com/google/uuid"
)

func GrantPermission(ctx context.Context, client *ent.Client, subjectType grantedpermission.SubjectType, subjectID uuid.UUID, objectType grantedpermission.ObjectType, objectID uuid.UUID, action actions.PermissionAction) (*ent.GrantedPermission, error) {
	// Check if permission is already granted
	entGrantedPermission, err := client.GrantedPermission.Query().Where(
		grantedpermission.SubjectTypeEQ(subjectType),
		grantedpermission.SubjectIDEQ(subjectID),
		grantedpermission.ObjectTypeEQ(objectType),
		grantedpermission.ObjectIDEQ(objectID),
		grantedpermission.ActionEQ(action),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("failed to query for granted permission %v", err)
	}

	// Create it if not exists
	if entGrantedPermission == nil {
		// Setup the creation query
		permissionCreate := client.GrantedPermission.Create().
			SetSubjectType(subjectType).
			SetSubjectID(subjectID).
			SetObjectType(objectType).
			SetObjectID(objectID).
			SetAction(action)

		// Query the subject based on type
		if subjectType == grantedpermission.SubjectTypeUser {
			entUser, err := client.User.Get(ctx, subjectID)
			if err != nil {
				return nil, fmt.Errorf("failed to query user subject with ID %s: %s", subjectID, err)
			}
			permissionCreate = permissionCreate.SetUser(entUser)
		} else if subjectType == grantedpermission.SubjectTypeGroup {
			entGroup, err := client.Group.Get(ctx, subjectID)
			if err != nil {
				return nil, fmt.Errorf("failed to query group subject with ID %s: %s", subjectID, err)
			}
			permissionCreate = permissionCreate.SetGroup(entGroup)
		}

		// Create the granted permission
		return permissionCreate.Save(ctx)
	}

	// Return nil if already exists
	return entGrantedPermission, nil
}

func RevokePermission(ctx context.Context, client *ent.Client, subjectType grantedpermission.SubjectType, subjectID uuid.UUID, objectType grantedpermission.ObjectType, objectID uuid.UUID, action actions.PermissionAction) error {
	// Check if permission is already granted
	entGrantedPermission, err := client.GrantedPermission.Query().Where(
		grantedpermission.SubjectTypeEQ(subjectType),
		grantedpermission.SubjectIDEQ(subjectID),
		grantedpermission.ObjectTypeEQ(objectType),
		grantedpermission.ObjectIDEQ(objectID),
		grantedpermission.ActionEQ(action),
	).Only(ctx)
	if err != nil {
		// If not found, return nil since didn't need revoking
		if ent.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("failed to query for granted permission %v", err)
	}

	// Delete it if exists
	return client.GrantedPermission.DeleteOne(entGrantedPermission).Exec(ctx)
}

// HasPermission reports whether a given user has a given permission for action on object. Use [github.com/google/uuid.Nil] to denote a wildcard object ID.
func HasPermission(ctx context.Context, client *ent.Client, entUser *ent.User, objectType grantedpermission.ObjectType, objectID uuid.UUID, action actions.PermissionAction) (bool, error) {
	// Get all group membership ID's
	entGroupIds, err := entUser.QueryGroups().IDs(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to query group id's from current user: %v", err)
	}

	return client.GrantedPermission.Query().Where(
		grantedpermission.And(
			// Subject Match
			grantedpermission.Or(
				// Permission by user
				grantedpermission.And(
					grantedpermission.SubjectTypeEQ(grantedpermission.SubjectTypeUser),
					grantedpermission.SubjectIDEQ(entUser.ID),
				),
				// Permission by groups
				grantedpermission.And(
					grantedpermission.SubjectTypeEQ(grantedpermission.SubjectTypeGroup),
					grantedpermission.SubjectIDIn(entGroupIds...),
				),
			),
			// Object Match
			grantedpermission.And(
				// Type and action
				grantedpermission.ObjectTypeEQ(objectType),
				grantedpermission.ActionEQ(action),
				// Either object ID or wildcard
				grantedpermission.Or(
					grantedpermission.ObjectIDEQ(uuid.Nil),
					grantedpermission.ObjectIDEQ(objectID),
				),
			),
		),
	).Exist(ctx)
}

// CurrentUserHasPermission reports whether the current user (pulled from context) has a given permission for action on object. Use [github.com/google/uuid.Nil] to denote a wildcard object ID.
func CurrentUserHasPermission(ctx context.Context, client *ent.Client, objectType grantedpermission.ObjectType, objectID uuid.UUID, action actions.PermissionAction) (bool, error) {
	// Get the current authenticated user
	currentUser, err := auth.ForContext(ctx)
	if err != nil {
		return false, auth.AUTH_REQUIRED_GQL_ERROR
	}

	return HasPermission(ctx, client, currentUser, objectType, objectID, action)
}

func DisplayString(subjectType grantedpermission.SubjectType, subjectID uuid.UUID, objectType grantedpermission.ObjectType, objectID uuid.UUID, action actions.PermissionAction) string {
	objectIdString := objectID.String()
	if objectID == uuid.Nil {
		objectIdString = "*"
	}
	return fmt.Sprintf("/%s/%s/%s/%s/%s", subjectType, subjectID, objectType, objectIdString, action.DisplayString())
}
