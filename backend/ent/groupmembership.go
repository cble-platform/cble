// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cble-platform/cble/backend/ent/group"
	"github.com/cble-platform/cble/backend/ent/groupmembership"
	"github.com/cble-platform/cble/backend/ent/project"
	"github.com/google/uuid"
)

// GroupMembership is the model entity for the GroupMembership schema.
type GroupMembership struct {
	config `json:"-"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID uuid.UUID `json:"project_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID uuid.UUID `json:"group_id,omitempty"`
	// Role holds the value of the "role" field.
	Role groupmembership.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupMembershipQuery when eager-loading is set.
	Edges        GroupMembershipEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupMembershipEdges holds the relations/edges for other nodes in the graph.
type GroupMembershipEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMembershipEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupMembershipEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupMembership) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupmembership.FieldRole:
			values[i] = new(sql.NullString)
		case groupmembership.FieldProjectID, groupmembership.FieldGroupID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupMembership fields.
func (gm *GroupMembership) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupmembership.FieldProjectID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				gm.ProjectID = *value
			}
		case groupmembership.FieldGroupID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value != nil {
				gm.GroupID = *value
			}
		case groupmembership.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				gm.Role = groupmembership.Role(value.String)
			}
		default:
			gm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupMembership.
// This includes values selected through modifiers, order, etc.
func (gm *GroupMembership) Value(name string) (ent.Value, error) {
	return gm.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the GroupMembership entity.
func (gm *GroupMembership) QueryProject() *ProjectQuery {
	return NewGroupMembershipClient(gm.config).QueryProject(gm)
}

// QueryGroup queries the "group" edge of the GroupMembership entity.
func (gm *GroupMembership) QueryGroup() *GroupQuery {
	return NewGroupMembershipClient(gm.config).QueryGroup(gm)
}

// Update returns a builder for updating this GroupMembership.
// Note that you need to call GroupMembership.Unwrap() before calling this method if this GroupMembership
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GroupMembership) Update() *GroupMembershipUpdateOne {
	return NewGroupMembershipClient(gm.config).UpdateOne(gm)
}

// Unwrap unwraps the GroupMembership entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GroupMembership) Unwrap() *GroupMembership {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupMembership is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GroupMembership) String() string {
	var builder strings.Builder
	builder.WriteString("GroupMembership(")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.GroupID))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", gm.Role))
	builder.WriteByte(')')
	return builder.String()
}

// GroupMemberships is a parsable slice of GroupMembership.
type GroupMemberships []*GroupMembership
