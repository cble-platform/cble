// Code generated by ent, DO NOT EDIT.

package blueprint

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cble-platform/cble-backend/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldDescription, v))
}

// BlueprintTemplate applies equality check predicate on the "blueprint_template" field. It's identical to BlueprintTemplateEQ.
func BlueprintTemplate(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldBlueprintTemplate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldContainsFold(FieldDescription, v))
}

// BlueprintTemplateEQ applies the EQ predicate on the "blueprint_template" field.
func BlueprintTemplateEQ(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldEQ(FieldBlueprintTemplate, v))
}

// BlueprintTemplateNEQ applies the NEQ predicate on the "blueprint_template" field.
func BlueprintTemplateNEQ(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNEQ(FieldBlueprintTemplate, v))
}

// BlueprintTemplateIn applies the In predicate on the "blueprint_template" field.
func BlueprintTemplateIn(vs ...[]byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldIn(FieldBlueprintTemplate, vs...))
}

// BlueprintTemplateNotIn applies the NotIn predicate on the "blueprint_template" field.
func BlueprintTemplateNotIn(vs ...[]byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldNotIn(FieldBlueprintTemplate, vs...))
}

// BlueprintTemplateGT applies the GT predicate on the "blueprint_template" field.
func BlueprintTemplateGT(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGT(FieldBlueprintTemplate, v))
}

// BlueprintTemplateGTE applies the GTE predicate on the "blueprint_template" field.
func BlueprintTemplateGTE(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldGTE(FieldBlueprintTemplate, v))
}

// BlueprintTemplateLT applies the LT predicate on the "blueprint_template" field.
func BlueprintTemplateLT(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLT(FieldBlueprintTemplate, v))
}

// BlueprintTemplateLTE applies the LTE predicate on the "blueprint_template" field.
func BlueprintTemplateLTE(v []byte) predicate.Blueprint {
	return predicate.Blueprint(sql.FieldLTE(FieldBlueprintTemplate, v))
}

// HasProvider applies the HasEdge predicate on the "provider" edge.
func HasProvider() predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderWith applies the HasEdge predicate on the "provider" edge with a given conditions (other predicates).
func HasProviderWith(preds ...predicate.Provider) predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := newProviderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResources applies the HasEdge predicate on the "resources" edge.
func HasResources() predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ResourcesTable, ResourcesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourcesWith applies the HasEdge predicate on the "resources" edge with a given conditions (other predicates).
func HasResourcesWith(preds ...predicate.Resource) predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := newResourcesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeployments applies the HasEdge predicate on the "deployments" edge.
func HasDeployments() predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentsWith applies the HasEdge predicate on the "deployments" edge with a given conditions (other predicates).
func HasDeploymentsWith(preds ...predicate.Deployment) predicate.Blueprint {
	return predicate.Blueprint(func(s *sql.Selector) {
		step := newDeploymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blueprint) predicate.Blueprint {
	return predicate.Blueprint(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blueprint) predicate.Blueprint {
	return predicate.Blueprint(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blueprint) predicate.Blueprint {
	return predicate.Blueprint(sql.NotPredicates(p))
}
