package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VirtualizationProvider holds the schema definition for the VirtualizationProvider entity.
type VirtualizationProvider struct {
	ent.Schema
}

// Fields of the VirtualizationProvider.
func (VirtualizationProvider) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("display_name"),
		field.String("provider_git_url"),
		field.String("provider_version"),
		field.Bytes("config_bytes"),
		field.Bool("is_loaded").Default(false),
	}
}

// Edges of the VirtualizationProvider.
func (VirtualizationProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blueprints", Blueprint.Type).
			Ref("virtualization_provider"),
	}
}
