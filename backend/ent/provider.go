// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	entprovider "github.com/cble-platform/cble/backend/ent/provider"
	"github.com/google/uuid"
)

// Provider is the model entity for the Provider schema.
type Provider struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// ProviderGitURL holds the value of the "provider_git_url" field.
	ProviderGitURL string `json:"provider_git_url,omitempty"`
	// ProviderVersion holds the value of the "provider_version" field.
	ProviderVersion string `json:"provider_version,omitempty"`
	// ConfigBytes holds the value of the "config_bytes" field.
	ConfigBytes []byte `json:"config_bytes,omitempty"`
	// IsLoaded holds the value of the "is_loaded" field.
	IsLoaded bool `json:"is_loaded,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProviderQuery when eager-loading is set.
	Edges        ProviderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProviderEdges holds the relations/edges for other nodes in the graph.
type ProviderEdges struct {
	// Blueprints holds the value of the blueprints edge.
	Blueprints []*Blueprint `json:"blueprints,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BlueprintsOrErr returns the Blueprints value or an error if the edge
// was not loaded in eager-loading.
func (e ProviderEdges) BlueprintsOrErr() ([]*Blueprint, error) {
	if e.loadedTypes[0] {
		return e.Blueprints, nil
	}
	return nil, &NotLoadedError{edge: "blueprints"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Provider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entprovider.FieldConfigBytes:
			values[i] = new([]byte)
		case entprovider.FieldIsLoaded:
			values[i] = new(sql.NullBool)
		case entprovider.FieldDisplayName, entprovider.FieldProviderGitURL, entprovider.FieldProviderVersion:
			values[i] = new(sql.NullString)
		case entprovider.FieldCreatedAt, entprovider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case entprovider.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Provider fields.
func (pr *Provider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entprovider.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case entprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case entprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case entprovider.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				pr.DisplayName = value.String
			}
		case entprovider.FieldProviderGitURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_git_url", values[i])
			} else if value.Valid {
				pr.ProviderGitURL = value.String
			}
		case entprovider.FieldProviderVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_version", values[i])
			} else if value.Valid {
				pr.ProviderVersion = value.String
			}
		case entprovider.FieldConfigBytes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field config_bytes", values[i])
			} else if value != nil {
				pr.ConfigBytes = *value
			}
		case entprovider.FieldIsLoaded:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_loaded", values[i])
			} else if value.Valid {
				pr.IsLoaded = value.Bool
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Provider.
// This includes values selected through modifiers, order, etc.
func (pr *Provider) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryBlueprints queries the "blueprints" edge of the Provider entity.
func (pr *Provider) QueryBlueprints() *BlueprintQuery {
	return NewProviderClient(pr.config).QueryBlueprints(pr)
}

// Update returns a builder for updating this Provider.
// Note that you need to call Provider.Unwrap() before calling this method if this Provider
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Provider) Update() *ProviderUpdateOne {
	return NewProviderClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Provider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Provider) Unwrap() *Provider {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Provider is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Provider) String() string {
	var builder strings.Builder
	builder.WriteString("Provider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(pr.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("provider_git_url=")
	builder.WriteString(pr.ProviderGitURL)
	builder.WriteString(", ")
	builder.WriteString("provider_version=")
	builder.WriteString(pr.ProviderVersion)
	builder.WriteString(", ")
	builder.WriteString("config_bytes=")
	builder.WriteString(fmt.Sprintf("%v", pr.ConfigBytes))
	builder.WriteString(", ")
	builder.WriteString("is_loaded=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsLoaded))
	builder.WriteByte(')')
	return builder.String()
}

// Providers is a parsable slice of Provider.
type Providers []*Provider
