package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("last_accessed").
			Optional().
			Default(time.Now).
			Comment("The last time this deployment was accessed (used for auto-suspending deployments)"),
		field.String("name").
			Comment("Display name of the deployment (defaults to blueprint name)"),
		field.String("description").
			Comment("Display description of the deployment (supports markdown; defaults to blueprint description)"),
		field.Enum("state").
			Values("awaiting", "in_progress", "complete", "failed", "destroyed", "suspended").
			Comment("The overall state of the deployment (should only by updated by the deploy engine)"),
		field.JSON("template_vars", map[string]string{}).
			Default(make(map[string]string)).
			Comment("Stores the variable values to be injected into the blueprint template"),
		field.Time("expires_at").
			Comment("The time this deployment expires"),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blueprint", Blueprint.Type).
			Unique().
			Required().
			Comment("The blueprint for this deployment"),
		edge.From("deployment_nodes", DeploymentNode.Type).
			Ref("deployment").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			Comment("The deployment nodes belonging to this deployment"),
		edge.To("requester", User.Type).
			Unique().
			Required().
			Comment("The user who requested this deployment"),
	}
}
