package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cble-platform/cble-backend/ent/mixins"
	"github.com/google/uuid"
)

// GroupMembership holds the schema definition for the GroupMembership entity.
type GroupMembership struct {
	ent.Schema
}

// Fields of the GroupMembership.
func (GroupMembership) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}),
		field.UUID("group_id", uuid.UUID{}),
		field.Enum("role").Values("viewer", "deployer", "developer", "admin").
			Default("deployer"),
	}
}

// Indexes of the GroupMembership.
func (GroupMembership) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("project_id", "group_id").
			Unique(),
	}
}

// Edges of the GroupMembership.
func (GroupMembership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("project", Project.Type).
			Required().
			Unique().
			Field("project_id"),
		edge.To("group", Group.Type).
			Required().
			Unique().
			Field("group_id"),
	}
}

// Mixins of the GroupMembership.
func (GroupMembership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Timestamps{},
	}
}
