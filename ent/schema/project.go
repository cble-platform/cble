package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/mixins"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("name").
			Comment("The name of the project"),
		field.Int("quota_cpu").
			Comment("The quota for number of CPU cores"),
		field.Int("quota_ram").
			Comment("The quota for total RAM usage (MiB)"),
		field.Int("quota_disk").
			Comment("The quota for total disk usage (MiB)"),
		field.Int("quota_network").
			Comment("The quota for number of networks"),
		field.Int("quota_router").
			Comment("The quota for number of routers"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", User.Type).
			Comment("Users who have access to this project"),
		edge.To("group_members", Group.Type).
			Comment("Groups who have access to this project"),
		edge.From("blueprints", Blueprint.Type).
			Ref("project").
			Comment("Blueprints which belong to this project"),
		edge.From("deployments", Deployment.Type).
			Ref("project").
			Comment("Deployments which belong to this project"),
	}
}

// Mixins of the Project.
func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Timestamps{},
	}
}
