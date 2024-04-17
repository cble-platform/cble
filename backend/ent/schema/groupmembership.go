package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// GroupMembership holds the schema definition for the GroupMembership entity.
type GroupMembership struct {
	ent.Schema
}

func (GroupMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("project_id", "group_id"),
	}
}

// Fields of the GroupMembership.
func (GroupMembership) Fields() []ent.Field {
	return []ent.Field{
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
