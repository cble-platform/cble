package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Provider holds the schema definition for the Provider entity.
type Provider struct {
	ent.Schema
}

// Fields of the Provider.
func (Provider) Fields() []ent.Field {
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

// Edges of the Provider.
func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blueprints", Blueprint.Type).
			Ref("provider"),
	}
}
