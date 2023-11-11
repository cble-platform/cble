// Code generated by ent, DO NOT EDIT.

package blueprint

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the blueprint type in the database.
	Label = "blueprint"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBlueprintTemplate holds the string denoting the blueprint_template field in the database.
	FieldBlueprintTemplate = "blueprint_template"
	// EdgeParentGroup holds the string denoting the parent_group edge name in mutations.
	EdgeParentGroup = "parent_group"
	// EdgeProvider holds the string denoting the provider edge name in mutations.
	EdgeProvider = "provider"
	// EdgeDeployments holds the string denoting the deployments edge name in mutations.
	EdgeDeployments = "deployments"
	// Table holds the table name of the blueprint in the database.
	Table = "blueprints"
	// ParentGroupTable is the table that holds the parent_group relation/edge.
	ParentGroupTable = "blueprints"
	// ParentGroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	ParentGroupInverseTable = "groups"
	// ParentGroupColumn is the table column denoting the parent_group relation/edge.
	ParentGroupColumn = "blueprint_parent_group"
	// ProviderTable is the table that holds the provider relation/edge.
	ProviderTable = "blueprints"
	// ProviderInverseTable is the table name for the Provider entity.
	// It exists in this package in order to avoid circular dependency with the "provider" package.
	ProviderInverseTable = "providers"
	// ProviderColumn is the table column denoting the provider relation/edge.
	ProviderColumn = "blueprint_provider"
	// DeploymentsTable is the table that holds the deployments relation/edge.
	DeploymentsTable = "deployments"
	// DeploymentsInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentsInverseTable = "deployments"
	// DeploymentsColumn is the table column denoting the deployments relation/edge.
	DeploymentsColumn = "deployment_blueprint"
)

// Columns holds all SQL columns for blueprint fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBlueprintTemplate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "blueprints"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"blueprint_parent_group",
	"blueprint_provider",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Blueprint queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByParentGroupField orders the results by parent_group field.
func ByParentGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentGroupStep(), sql.OrderByField(field, opts...))
	}
}

// ByProviderField orders the results by provider field.
func ByProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProviderStep(), sql.OrderByField(field, opts...))
	}
}

// ByDeploymentsCount orders the results by deployments count.
func ByDeploymentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeploymentsStep(), opts...)
	}
}

// ByDeployments orders the results by deployments terms.
func ByDeployments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeploymentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newParentGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParentGroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ParentGroupTable, ParentGroupColumn),
	)
}
func newProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProviderTable, ProviderColumn),
	)
}
func newDeploymentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeploymentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, DeploymentsTable, DeploymentsColumn),
	)
}
