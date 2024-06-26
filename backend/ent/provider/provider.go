// Code generated by ent, DO NOT EDIT.

package entprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the provider type in the database.
	Label = "provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldProviderGitURL holds the string denoting the provider_git_url field in the database.
	FieldProviderGitURL = "provider_git_url"
	// FieldProviderVersion holds the string denoting the provider_version field in the database.
	FieldProviderVersion = "provider_version"
	// FieldConfigBytes holds the string denoting the config_bytes field in the database.
	FieldConfigBytes = "config_bytes"
	// FieldIsLoaded holds the string denoting the is_loaded field in the database.
	FieldIsLoaded = "is_loaded"
	// EdgeBlueprints holds the string denoting the blueprints edge name in mutations.
	EdgeBlueprints = "blueprints"
	// Table holds the table name of the provider in the database.
	Table = "providers"
	// BlueprintsTable is the table that holds the blueprints relation/edge.
	BlueprintsTable = "blueprints"
	// BlueprintsInverseTable is the table name for the Blueprint entity.
	// It exists in this package in order to avoid circular dependency with the "blueprint" package.
	BlueprintsInverseTable = "blueprints"
	// BlueprintsColumn is the table column denoting the blueprints relation/edge.
	BlueprintsColumn = "blueprint_provider"
)

// Columns holds all SQL columns for provider fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDisplayName,
	FieldProviderGitURL,
	FieldProviderVersion,
	FieldConfigBytes,
	FieldIsLoaded,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsLoaded holds the default value on creation for the "is_loaded" field.
	DefaultIsLoaded bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Provider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByProviderGitURL orders the results by the provider_git_url field.
func ByProviderGitURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderGitURL, opts...).ToFunc()
}

// ByProviderVersion orders the results by the provider_version field.
func ByProviderVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderVersion, opts...).ToFunc()
}

// ByIsLoaded orders the results by the is_loaded field.
func ByIsLoaded(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsLoaded, opts...).ToFunc()
}

// ByBlueprintsCount orders the results by blueprints count.
func ByBlueprintsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBlueprintsStep(), opts...)
	}
}

// ByBlueprints orders the results by blueprints terms.
func ByBlueprints(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlueprintsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBlueprintsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlueprintsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, BlueprintsTable, BlueprintsColumn),
	)
}
