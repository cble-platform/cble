// Code generated by ent, DO NOT EDIT.

package entprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cble-platform/cble/backend/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldUpdatedAt, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldDisplayName, v))
}

// ProviderGitURL applies equality check predicate on the "provider_git_url" field. It's identical to ProviderGitURLEQ.
func ProviderGitURL(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldProviderGitURL, v))
}

// ProviderVersion applies equality check predicate on the "provider_version" field. It's identical to ProviderVersionEQ.
func ProviderVersion(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldProviderVersion, v))
}

// ConfigBytes applies equality check predicate on the "config_bytes" field. It's identical to ConfigBytesEQ.
func ConfigBytes(v []byte) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldConfigBytes, v))
}

// IsLoaded applies equality check predicate on the "is_loaded" field. It's identical to IsLoadedEQ.
func IsLoaded(v bool) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldIsLoaded, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldUpdatedAt, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContainsFold(FieldDisplayName, v))
}

// ProviderGitURLEQ applies the EQ predicate on the "provider_git_url" field.
func ProviderGitURLEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldProviderGitURL, v))
}

// ProviderGitURLNEQ applies the NEQ predicate on the "provider_git_url" field.
func ProviderGitURLNEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldProviderGitURL, v))
}

// ProviderGitURLIn applies the In predicate on the "provider_git_url" field.
func ProviderGitURLIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldProviderGitURL, vs...))
}

// ProviderGitURLNotIn applies the NotIn predicate on the "provider_git_url" field.
func ProviderGitURLNotIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldProviderGitURL, vs...))
}

// ProviderGitURLGT applies the GT predicate on the "provider_git_url" field.
func ProviderGitURLGT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldProviderGitURL, v))
}

// ProviderGitURLGTE applies the GTE predicate on the "provider_git_url" field.
func ProviderGitURLGTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldProviderGitURL, v))
}

// ProviderGitURLLT applies the LT predicate on the "provider_git_url" field.
func ProviderGitURLLT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldProviderGitURL, v))
}

// ProviderGitURLLTE applies the LTE predicate on the "provider_git_url" field.
func ProviderGitURLLTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldProviderGitURL, v))
}

// ProviderGitURLContains applies the Contains predicate on the "provider_git_url" field.
func ProviderGitURLContains(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContains(FieldProviderGitURL, v))
}

// ProviderGitURLHasPrefix applies the HasPrefix predicate on the "provider_git_url" field.
func ProviderGitURLHasPrefix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasPrefix(FieldProviderGitURL, v))
}

// ProviderGitURLHasSuffix applies the HasSuffix predicate on the "provider_git_url" field.
func ProviderGitURLHasSuffix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasSuffix(FieldProviderGitURL, v))
}

// ProviderGitURLEqualFold applies the EqualFold predicate on the "provider_git_url" field.
func ProviderGitURLEqualFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEqualFold(FieldProviderGitURL, v))
}

// ProviderGitURLContainsFold applies the ContainsFold predicate on the "provider_git_url" field.
func ProviderGitURLContainsFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContainsFold(FieldProviderGitURL, v))
}

// ProviderVersionEQ applies the EQ predicate on the "provider_version" field.
func ProviderVersionEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldProviderVersion, v))
}

// ProviderVersionNEQ applies the NEQ predicate on the "provider_version" field.
func ProviderVersionNEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldProviderVersion, v))
}

// ProviderVersionIn applies the In predicate on the "provider_version" field.
func ProviderVersionIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldProviderVersion, vs...))
}

// ProviderVersionNotIn applies the NotIn predicate on the "provider_version" field.
func ProviderVersionNotIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldProviderVersion, vs...))
}

// ProviderVersionGT applies the GT predicate on the "provider_version" field.
func ProviderVersionGT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldProviderVersion, v))
}

// ProviderVersionGTE applies the GTE predicate on the "provider_version" field.
func ProviderVersionGTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldProviderVersion, v))
}

// ProviderVersionLT applies the LT predicate on the "provider_version" field.
func ProviderVersionLT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldProviderVersion, v))
}

// ProviderVersionLTE applies the LTE predicate on the "provider_version" field.
func ProviderVersionLTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldProviderVersion, v))
}

// ProviderVersionContains applies the Contains predicate on the "provider_version" field.
func ProviderVersionContains(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContains(FieldProviderVersion, v))
}

// ProviderVersionHasPrefix applies the HasPrefix predicate on the "provider_version" field.
func ProviderVersionHasPrefix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasPrefix(FieldProviderVersion, v))
}

// ProviderVersionHasSuffix applies the HasSuffix predicate on the "provider_version" field.
func ProviderVersionHasSuffix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasSuffix(FieldProviderVersion, v))
}

// ProviderVersionEqualFold applies the EqualFold predicate on the "provider_version" field.
func ProviderVersionEqualFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEqualFold(FieldProviderVersion, v))
}

// ProviderVersionContainsFold applies the ContainsFold predicate on the "provider_version" field.
func ProviderVersionContainsFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContainsFold(FieldProviderVersion, v))
}

// ConfigBytesEQ applies the EQ predicate on the "config_bytes" field.
func ConfigBytesEQ(v []byte) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldConfigBytes, v))
}

// ConfigBytesNEQ applies the NEQ predicate on the "config_bytes" field.
func ConfigBytesNEQ(v []byte) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldConfigBytes, v))
}

// ConfigBytesIn applies the In predicate on the "config_bytes" field.
func ConfigBytesIn(vs ...[]byte) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldConfigBytes, vs...))
}

// ConfigBytesNotIn applies the NotIn predicate on the "config_bytes" field.
func ConfigBytesNotIn(vs ...[]byte) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldConfigBytes, vs...))
}

// ConfigBytesGT applies the GT predicate on the "config_bytes" field.
func ConfigBytesGT(v []byte) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldConfigBytes, v))
}

// ConfigBytesGTE applies the GTE predicate on the "config_bytes" field.
func ConfigBytesGTE(v []byte) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldConfigBytes, v))
}

// ConfigBytesLT applies the LT predicate on the "config_bytes" field.
func ConfigBytesLT(v []byte) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldConfigBytes, v))
}

// ConfigBytesLTE applies the LTE predicate on the "config_bytes" field.
func ConfigBytesLTE(v []byte) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldConfigBytes, v))
}

// IsLoadedEQ applies the EQ predicate on the "is_loaded" field.
func IsLoadedEQ(v bool) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldIsLoaded, v))
}

// IsLoadedNEQ applies the NEQ predicate on the "is_loaded" field.
func IsLoadedNEQ(v bool) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldIsLoaded, v))
}

// HasBlueprints applies the HasEdge predicate on the "blueprints" edge.
func HasBlueprints() predicate.Provider {
	return predicate.Provider(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BlueprintsTable, BlueprintsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlueprintsWith applies the HasEdge predicate on the "blueprints" edge with a given conditions (other predicates).
func HasBlueprintsWith(preds ...predicate.Blueprint) predicate.Provider {
	return predicate.Provider(func(s *sql.Selector) {
		step := newBlueprintsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Provider) predicate.Provider {
	return predicate.Provider(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Provider) predicate.Provider {
	return predicate.Provider(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Provider) predicate.Provider {
	return predicate.Provider(sql.NotPredicates(p))
}
