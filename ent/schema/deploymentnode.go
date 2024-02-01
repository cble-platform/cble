package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DeploymentNode holds the schema definition for the DeploymentNode entity.
type DeploymentNode struct {
	ent.Schema
}

// Fields of the DeploymentNode.
func (DeploymentNode) Fields() []ent.Field {
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
		field.Enum("state").
			Values("to_deploy", "to_destroy", "to_redeploy", "parent_awaiting", "child_awaiting", "in_progress", "complete", "tainted", "failed", "destroyed").
			Comment("The state of the deployed resource (should only by updated by the deploy engine)"),
		field.JSON("vars", map[string]string{}).
			Default(make(map[string]string)).
			Comment("Stores metadata about the deployed resource for use with the provider"),
	}
}

// Edges of the DeploymentNode.
func (DeploymentNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deployment", Deployment.Type).
			Unique().
			Required().
			Comment("The deployment for this node"),
		edge.To("resource", Resource.Type).
			Unique().
			Required().
			Comment("The resource this node represents"),
		edge.To("next_nodes", DeploymentNode.Type).
			Comment("The next nodes in the dependency tree").
			From("prev_nodes").
			Comment("The previous nodes in the dependency tree"),
	}
}
