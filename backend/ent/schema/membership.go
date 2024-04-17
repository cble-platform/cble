package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cble-platform/cble/backend/ent/mixins"
	"github.com/google/uuid"
)

// Membership holds the schema definition for the Membership entity.
type Membership struct {
	ent.Schema
}

// Fields of the Membership.
func (Membership) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.UUID("project_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Enum("role").Values("viewer", "deployer", "developer", "admin").
			Default("deployer"),
	}
}

// Indexes of the Membership.
func (Membership) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("project_id", "user_id").
			Unique(),
	}
}

// Edges of the Membership.
func (Membership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("project", Project.Type).
			Required().
			Unique().
			Field("project_id"),
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
	}
}

// Mixins of the Membership.
func (Membership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Timestamps{},
	}
}
