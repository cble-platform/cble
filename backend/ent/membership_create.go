// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble/backend/ent/membership"
	"github.com/cble-platform/cble/backend/ent/project"
	"github.com/cble-platform/cble/backend/ent/user"
	"github.com/google/uuid"
)

// MembershipCreate is the builder for creating a Membership entity.
type MembershipCreate struct {
	config
	mutation *MembershipMutation
	hooks    []Hook
}

// SetProjectID sets the "project_id" field.
func (mc *MembershipCreate) SetProjectID(u uuid.UUID) *MembershipCreate {
	mc.mutation.SetProjectID(u)
	return mc
}

// SetUserID sets the "user_id" field.
func (mc *MembershipCreate) SetUserID(u uuid.UUID) *MembershipCreate {
	mc.mutation.SetUserID(u)
	return mc
}

// SetRole sets the "role" field.
func (mc *MembershipCreate) SetRole(m membership.Role) *MembershipCreate {
	mc.mutation.SetRole(m)
	return mc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableRole(m *membership.Role) *MembershipCreate {
	if m != nil {
		mc.SetRole(*m)
	}
	return mc
}

// SetProject sets the "project" edge to the Project entity.
func (mc *MembershipCreate) SetProject(p *Project) *MembershipCreate {
	return mc.SetProjectID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (mc *MembershipCreate) SetUser(u *User) *MembershipCreate {
	return mc.SetUserID(u.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (mc *MembershipCreate) Mutation() *MembershipMutation {
	return mc.mutation
}

// Save creates the Membership in the database.
func (mc *MembershipCreate) Save(ctx context.Context) (*Membership, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MembershipCreate) SaveX(ctx context.Context) *Membership {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MembershipCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MembershipCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MembershipCreate) defaults() {
	if _, ok := mc.mutation.Role(); !ok {
		v := membership.DefaultRole
		mc.mutation.SetRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MembershipCreate) check() error {
	if _, ok := mc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project_id", err: errors.New(`ent: missing required field "Membership.project_id"`)}
	}
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Membership.user_id"`)}
	}
	if _, ok := mc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "Membership.role"`)}
	}
	if v, ok := mc.mutation.Role(); ok {
		if err := membership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Membership.role": %w`, err)}
		}
	}
	if _, ok := mc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New(`ent: missing required edge "Membership.project"`)}
	}
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Membership.user"`)}
	}
	return nil
}

func (mc *MembershipCreate) sqlSave(ctx context.Context) (*Membership, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (mc *MembershipCreate) createSpec() (*Membership, *sqlgraph.CreateSpec) {
	var (
		_node = &Membership{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(membership.Table, nil)
	)
	if value, ok := mc.mutation.Role(); ok {
		_spec.SetField(membership.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := mc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   membership.ProjectTable,
			Columns: []string{membership.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MembershipCreateBulk is the builder for creating many Membership entities in bulk.
type MembershipCreateBulk struct {
	config
	err      error
	builders []*MembershipCreate
}

// Save creates the Membership entities in the database.
func (mcb *MembershipCreateBulk) Save(ctx context.Context) ([]*Membership, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Membership, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MembershipMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MembershipCreateBulk) SaveX(ctx context.Context) []*Membership {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MembershipCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MembershipCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
