package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/blueprintengine/models"
	"github.com/google/uuid"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
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
		field.String("key").
			Comment("Store the resource key from the blueprint"),
		field.JSON("object", &models.Object{}).
			Comment("Store the resource object from the blueprint"),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blueprint", Blueprint.Type).
			Unique().
			Required().
			Comment("Blueprint containing this resource"),
		edge.To("required_by", Resource.Type).
			Comment("Stores all dependents of this resource").
			From("depends_on").
			Comment("Stores all dependencies of this resource"),
	}
}
