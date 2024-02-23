package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/engine/models"
	"github.com/cble-platform/cble-backend/ent/mixins"
	"github.com/google/uuid"
)

// Blueprint holds the schema definition for the Blueprint entity.
type Blueprint struct {
	ent.Schema
}

// Fields of the Blueprint.
func (Blueprint) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("name").
			Comment("Display name of the blueprint"),
		field.String("description").
			Comment("Display description of the blueprint (supports markdown)"),
		field.Bytes("blueprint_template").
			Comment("The blueprint file contents"),
		field.JSON("variable_types", map[string]models.BlueprintVariableType{}).
			Comment("Stores the names of variables and their data type"),
	}
}

// Edges of the Blueprint.
func (Blueprint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider", Provider.Type).
			Unique().
			Required().
			Comment("The provider to use for this blueprint"),
		edge.From("resources", Resource.Type).
			Ref("blueprint").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			Comment("The resources which are part of this blueprint"),
		edge.From("deployments", Deployment.Type).
			Ref("blueprint").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Restrict,
			}).
			Comment("All deployments of this blueprints"),
	}
}

// Mixins of the Blueprint.
func (Blueprint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Timestamps{},
	}
}
