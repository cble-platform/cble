// Code generated by ent, DO NOT EDIT.

package deployment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deployment type in the database.
	Label = "deployment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTemplateVars holds the string denoting the template_vars field in the database.
	FieldTemplateVars = "template_vars"
	// FieldDeploymentVars holds the string denoting the deployment_vars field in the database.
	FieldDeploymentVars = "deployment_vars"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// EdgeBlueprint holds the string denoting the blueprint edge name in mutations.
	EdgeBlueprint = "blueprint"
	// EdgeRequester holds the string denoting the requester edge name in mutations.
	EdgeRequester = "requester"
	// Table holds the table name of the deployment in the database.
	Table = "deployments"
	// BlueprintTable is the table that holds the blueprint relation/edge.
	BlueprintTable = "deployments"
	// BlueprintInverseTable is the table name for the Blueprint entity.
	// It exists in this package in order to avoid circular dependency with the "blueprint" package.
	BlueprintInverseTable = "blueprints"
	// BlueprintColumn is the table column denoting the blueprint relation/edge.
	BlueprintColumn = "deployment_blueprint"
	// RequesterTable is the table that holds the requester relation/edge.
	RequesterTable = "deployments"
	// RequesterInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RequesterInverseTable = "users"
	// RequesterColumn is the table column denoting the requester relation/edge.
	RequesterColumn = "deployment_requester"
)

// Columns holds all SQL columns for deployment fields.
var Columns = []string{
	FieldID,
	FieldTemplateVars,
	FieldDeploymentVars,
	FieldIsActive,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "deployments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"deployment_blueprint",
	"deployment_requester",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTemplateVars holds the default value on creation for the "template_vars" field.
	DefaultTemplateVars map[string]interface{}
	// DefaultDeploymentVars holds the default value on creation for the "deployment_vars" field.
	DefaultDeploymentVars map[string]interface{}
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive map[string]int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Deployment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBlueprintField orders the results by blueprint field.
func ByBlueprintField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlueprintStep(), sql.OrderByField(field, opts...))
	}
}

// ByRequesterField orders the results by requester field.
func ByRequesterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRequesterStep(), sql.OrderByField(field, opts...))
	}
}
func newBlueprintStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlueprintInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BlueprintTable, BlueprintColumn),
	)
}
func newRequesterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RequesterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RequesterTable, RequesterColumn),
	)
}
