package graph

import (
	"context"
	"fmt"

	"github.com/cble-platform/cble/backend/auth"
	"github.com/cble-platform/cble/backend/ent"
	"github.com/cble-platform/cble/backend/ent/group"
	"github.com/cble-platform/cble/backend/ent/groupmembership"
	"github.com/cble-platform/cble/backend/ent/membership"
	"github.com/cble-platform/cble/backend/ent/predicate"
	"github.com/cble-platform/cble/backend/ent/project"
	"github.com/cble-platform/cble/backend/ent/user"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CurrentUserHasMinimumProjectRole(ctx context.Context, client *ent.Client, projectId uuid.UUID, role membership.Role) (bool, error) {
	// Get the current authenticated user
	currentUser, err := auth.ForContext(ctx)
	if err != nil {
		return false, gqlerror.Errorf("failed to get user from context: %v", err)
	}

	var groupRolePredicate predicate.GroupMembership
	if role == membership.RoleAdmin {
		// Admin or up
		groupRolePredicate = groupmembership.RoleIn(
			groupmembership.RoleAdmin,
		)
	} else if role == membership.RoleDeveloper {
		// Developer or up
		groupRolePredicate = groupmembership.RoleIn(
			groupmembership.RoleAdmin,
			groupmembership.RoleDeveloper,
		)
	} else if role == membership.RoleDeployer {
		// Deployer or up
		groupRolePredicate = groupmembership.RoleIn(
			groupmembership.RoleAdmin,
			groupmembership.RoleDeveloper,
			groupmembership.RoleDeployer,
		)
	} else if role == membership.RoleViewer {
		// Viewer or up
		groupRolePredicate = groupmembership.RoleIn(
			groupmembership.RoleAdmin,
			groupmembership.RoleDeveloper,
			groupmembership.RoleDeployer,
			groupmembership.RoleViewer,
		)
	}
	// Check if user inherits group roles
	hasGroupRole, err := client.GroupMembership.Query().Where(
		groupmembership.HasProjectWith(project.ID(projectId)),
		groupmembership.HasGroupWith(
			group.HasUsersWith(user.ID(currentUser.ID)),
		),
		groupRolePredicate,
	).Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to query group membership: %v", err)
	}
	// Has the group role, return true
	if hasGroupRole {
		return true, nil
	}

	var rolePredicate predicate.Membership
	if role == membership.RoleAdmin {
		// Admin or up
		rolePredicate = membership.RoleIn(
			membership.RoleAdmin,
		)
	} else if role == membership.RoleDeveloper {
		// Developer or up
		rolePredicate = membership.RoleIn(
			membership.RoleAdmin,
			membership.RoleDeveloper,
		)
	} else if role == membership.RoleDeployer {
		// Deployer or up
		rolePredicate = membership.RoleIn(
			membership.RoleAdmin,
			membership.RoleDeveloper,
			membership.RoleDeployer,
		)
	} else if role == membership.RoleViewer {
		// Viewer or up
		rolePredicate = membership.RoleIn(
			membership.RoleAdmin,
			membership.RoleDeveloper,
			membership.RoleDeployer,
			membership.RoleViewer,
		)
	}
	// Check if user has role directly
	hasRole, err := client.Membership.Query().Where(
		membership.HasProjectWith(project.ID(projectId)),
		membership.HasUserWith(
			user.ID(currentUser.ID),
		),
		rolePredicate,
	).Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to query membership: %v", err)
	}
	// Has the role, return true
	if hasRole {
		return true, nil
	}

	return false, nil
}
