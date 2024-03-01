package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cble-platform/cble-backend/ent/mixins"
	"github.com/cble-platform/cble-backend/permission/actions"
	"github.com/google/uuid"
)

// GrantedPermission holds the schema definition for the GrantedPermission entity.
type GrantedPermission struct {
	ent.Schema
}

// Fields of the GrantedPermission.
func (GrantedPermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.Enum("subject_type").
			Values("user", "group").
			Comment("The type of subject this permission applies to"),
		field.UUID("subject_id", uuid.UUID{}).
			Comment("The ID of subject this permission applies to"),
		field.Enum("object_type").
			Values("blueprint", "deployment", "group", "permission", "project", "provider", "user").
			Comment("The type of object this permission applies to"),
		field.UUID("object_id", uuid.UUID{}).
			Comment("The ID of object this permission applies to (or `uuid.Nil` for wildcard)"),
		field.Enum("action").
			GoType(actions.PermissionAction("")).
			Comment("The action associated with the object"),
	}
}

func (GrantedPermission) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("subject_type", "subject_id", "object_type", "object_id", "action").
			Unique(),
	}
}

// Edges of the GrantedPermission.
func (GrantedPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Comment("The subject user (if of type user)"),
		edge.To("group", Group.Type).
			Unique().
			Comment("The subject group (if of type user)"),
	}
}

// Mixins of the GrantedPermission.
func (GrantedPermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Timestamps{},
	}
}
