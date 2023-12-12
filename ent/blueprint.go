// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cble-platform/cble-backend/ent/blueprint"
	"github.com/cble-platform/cble-backend/ent/group"
	"github.com/cble-platform/cble-backend/ent/provider"
	"github.com/google/uuid"
)

// Blueprint is the model entity for the Blueprint schema.
type Blueprint struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// BlueprintTemplate holds the value of the "blueprint_template" field.
	BlueprintTemplate []byte `json:"blueprint_template,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlueprintQuery when eager-loading is set.
	Edges                  BlueprintEdges `json:"edges"`
	blueprint_parent_group *uuid.UUID
	blueprint_provider     *uuid.UUID
	selectValues           sql.SelectValues
}

// BlueprintEdges holds the relations/edges for other nodes in the graph.
type BlueprintEdges struct {
	// ParentGroup holds the value of the parent_group edge.
	ParentGroup *Group `json:"parent_group,omitempty"`
	// Provider holds the value of the provider edge.
	Provider *Provider `json:"provider,omitempty"`
	// Deployments holds the value of the deployments edge.
	Deployments []*Deployment `json:"deployments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentGroupOrErr returns the ParentGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlueprintEdges) ParentGroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.ParentGroup == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.ParentGroup, nil
	}
	return nil, &NotLoadedError{edge: "parent_group"}
}

// ProviderOrErr returns the Provider value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlueprintEdges) ProviderOrErr() (*Provider, error) {
	if e.loadedTypes[1] {
		if e.Provider == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: provider.Label}
		}
		return e.Provider, nil
	}
	return nil, &NotLoadedError{edge: "provider"}
}

// DeploymentsOrErr returns the Deployments value or an error if the edge
// was not loaded in eager-loading.
func (e BlueprintEdges) DeploymentsOrErr() ([]*Deployment, error) {
	if e.loadedTypes[2] {
		return e.Deployments, nil
	}
	return nil, &NotLoadedError{edge: "deployments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blueprint) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blueprint.FieldBlueprintTemplate:
			values[i] = new([]byte)
		case blueprint.FieldName, blueprint.FieldDescription:
			values[i] = new(sql.NullString)
		case blueprint.FieldCreatedAt, blueprint.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case blueprint.FieldID:
			values[i] = new(uuid.UUID)
		case blueprint.ForeignKeys[0]: // blueprint_parent_group
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case blueprint.ForeignKeys[1]: // blueprint_provider
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blueprint fields.
func (b *Blueprint) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blueprint.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case blueprint.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case blueprint.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case blueprint.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case blueprint.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case blueprint.FieldBlueprintTemplate:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field blueprint_template", values[i])
			} else if value != nil {
				b.BlueprintTemplate = *value
			}
		case blueprint.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field blueprint_parent_group", values[i])
			} else if value.Valid {
				b.blueprint_parent_group = new(uuid.UUID)
				*b.blueprint_parent_group = *value.S.(*uuid.UUID)
			}
		case blueprint.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field blueprint_provider", values[i])
			} else if value.Valid {
				b.blueprint_provider = new(uuid.UUID)
				*b.blueprint_provider = *value.S.(*uuid.UUID)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Blueprint.
// This includes values selected through modifiers, order, etc.
func (b *Blueprint) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryParentGroup queries the "parent_group" edge of the Blueprint entity.
func (b *Blueprint) QueryParentGroup() *GroupQuery {
	return NewBlueprintClient(b.config).QueryParentGroup(b)
}

// QueryProvider queries the "provider" edge of the Blueprint entity.
func (b *Blueprint) QueryProvider() *ProviderQuery {
	return NewBlueprintClient(b.config).QueryProvider(b)
}

// QueryDeployments queries the "deployments" edge of the Blueprint entity.
func (b *Blueprint) QueryDeployments() *DeploymentQuery {
	return NewBlueprintClient(b.config).QueryDeployments(b)
}

// Update returns a builder for updating this Blueprint.
// Note that you need to call Blueprint.Unwrap() before calling this method if this Blueprint
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blueprint) Update() *BlueprintUpdateOne {
	return NewBlueprintClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Blueprint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blueprint) Unwrap() *Blueprint {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blueprint is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blueprint) String() string {
	var builder strings.Builder
	builder.WriteString("Blueprint(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(b.Description)
	builder.WriteString(", ")
	builder.WriteString("blueprint_template=")
	builder.WriteString(fmt.Sprintf("%v", b.BlueprintTemplate))
	builder.WriteByte(')')
	return builder.String()
}

// Blueprints is a parsable slice of Blueprint.
type Blueprints []*Blueprint
