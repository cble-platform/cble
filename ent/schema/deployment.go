package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Deployment holds the schema definition for the Deployment entity.
type Deployment struct {
	ent.Schema
}

// Fields of the Deployment.
func (Deployment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.JSON("template_vars", map[string]interface{}{}).
			Default(make(map[string]interface{})),
		field.JSON("deployment_vars", map[string]interface{}{}).
			Default(make(map[string]interface{})),
		field.JSON("deployment_state", map[string]string{}).
			Default(make(map[string]string)),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blueprint", Blueprint.Type).
			Unique().
			Required(),
		edge.To("requester", User.Type).
			Unique().
			Required(),
	}
}
