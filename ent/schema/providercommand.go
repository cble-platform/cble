package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Provider holds the schema definition for the Provider entity.
type ProviderCommand struct {
	ent.Schema
}

// Fields of the Provider.
func (ProviderCommand) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("command_type").Values("CONFIGURE", "DEPLOY", "DESTROY", "CONSOLE", "RESOURCES"),
		field.Enum("status").Values("QUEUED", "FAILED", "SUCCEEDED", "INPROGRESS").Default("QUEUED"),
		field.Strings("arguments").Optional(),
		field.Time("start_time").Optional(),
		field.Time("end_time").Optional(),
		field.String("output").Default(""),
		field.String("error").Default(""),
	}
}

// Edges of the Provider.
func (ProviderCommand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider", Provider.Type).
			Unique().
			Required(),
		edge.To("deployment", Deployment.Type).
			Unique(),
	}
}
