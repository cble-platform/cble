// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble/backend/ent/blueprint"
	"github.com/cble-platform/cble/backend/ent/deployment"
	"github.com/cble-platform/cble/backend/ent/group"
	"github.com/cble-platform/cble/backend/ent/project"
	"github.com/cble-platform/cble/backend/ent/user"
	"github.com/google/uuid"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProjectCreate) SetCreatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProjectCreate) SetUpdatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUpdatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetQuotaCPU sets the "quota_cpu" field.
func (pc *ProjectCreate) SetQuotaCPU(i int) *ProjectCreate {
	pc.mutation.SetQuotaCPU(i)
	return pc
}

// SetUsageCPU sets the "usage_cpu" field.
func (pc *ProjectCreate) SetUsageCPU(i int) *ProjectCreate {
	pc.mutation.SetUsageCPU(i)
	return pc
}

// SetNillableUsageCPU sets the "usage_cpu" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUsageCPU(i *int) *ProjectCreate {
	if i != nil {
		pc.SetUsageCPU(*i)
	}
	return pc
}

// SetQuotaRAM sets the "quota_ram" field.
func (pc *ProjectCreate) SetQuotaRAM(i int) *ProjectCreate {
	pc.mutation.SetQuotaRAM(i)
	return pc
}

// SetUsageRAM sets the "usage_ram" field.
func (pc *ProjectCreate) SetUsageRAM(i int) *ProjectCreate {
	pc.mutation.SetUsageRAM(i)
	return pc
}

// SetNillableUsageRAM sets the "usage_ram" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUsageRAM(i *int) *ProjectCreate {
	if i != nil {
		pc.SetUsageRAM(*i)
	}
	return pc
}

// SetQuotaDisk sets the "quota_disk" field.
func (pc *ProjectCreate) SetQuotaDisk(i int) *ProjectCreate {
	pc.mutation.SetQuotaDisk(i)
	return pc
}

// SetUsageDisk sets the "usage_disk" field.
func (pc *ProjectCreate) SetUsageDisk(i int) *ProjectCreate {
	pc.mutation.SetUsageDisk(i)
	return pc
}

// SetNillableUsageDisk sets the "usage_disk" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUsageDisk(i *int) *ProjectCreate {
	if i != nil {
		pc.SetUsageDisk(*i)
	}
	return pc
}

// SetQuotaNetwork sets the "quota_network" field.
func (pc *ProjectCreate) SetQuotaNetwork(i int) *ProjectCreate {
	pc.mutation.SetQuotaNetwork(i)
	return pc
}

// SetUsageNetwork sets the "usage_network" field.
func (pc *ProjectCreate) SetUsageNetwork(i int) *ProjectCreate {
	pc.mutation.SetUsageNetwork(i)
	return pc
}

// SetNillableUsageNetwork sets the "usage_network" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUsageNetwork(i *int) *ProjectCreate {
	if i != nil {
		pc.SetUsageNetwork(*i)
	}
	return pc
}

// SetQuotaRouter sets the "quota_router" field.
func (pc *ProjectCreate) SetQuotaRouter(i int) *ProjectCreate {
	pc.mutation.SetQuotaRouter(i)
	return pc
}

// SetUsageRouter sets the "usage_router" field.
func (pc *ProjectCreate) SetUsageRouter(i int) *ProjectCreate {
	pc.mutation.SetUsageRouter(i)
	return pc
}

// SetNillableUsageRouter sets the "usage_router" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUsageRouter(i *int) *ProjectCreate {
	if i != nil {
		pc.SetUsageRouter(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProjectCreate) SetID(u uuid.UUID) *ProjectCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableID(u *uuid.UUID) *ProjectCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (pc *ProjectCreate) AddMemberIDs(ids ...uuid.UUID) *ProjectCreate {
	pc.mutation.AddMemberIDs(ids...)
	return pc
}

// AddMembers adds the "members" edges to the User entity.
func (pc *ProjectCreate) AddMembers(u ...*User) *ProjectCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddMemberIDs(ids...)
}

// AddGroupMemberIDs adds the "group_members" edge to the Group entity by IDs.
func (pc *ProjectCreate) AddGroupMemberIDs(ids ...uuid.UUID) *ProjectCreate {
	pc.mutation.AddGroupMemberIDs(ids...)
	return pc
}

// AddGroupMembers adds the "group_members" edges to the Group entity.
func (pc *ProjectCreate) AddGroupMembers(g ...*Group) *ProjectCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return pc.AddGroupMemberIDs(ids...)
}

// AddBlueprintIDs adds the "blueprints" edge to the Blueprint entity by IDs.
func (pc *ProjectCreate) AddBlueprintIDs(ids ...uuid.UUID) *ProjectCreate {
	pc.mutation.AddBlueprintIDs(ids...)
	return pc
}

// AddBlueprints adds the "blueprints" edges to the Blueprint entity.
func (pc *ProjectCreate) AddBlueprints(b ...*Blueprint) *ProjectCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBlueprintIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (pc *ProjectCreate) AddDeploymentIDs(ids ...uuid.UUID) *ProjectCreate {
	pc.mutation.AddDeploymentIDs(ids...)
	return pc
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (pc *ProjectCreate) AddDeployments(d ...*Deployment) *ProjectCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDeploymentIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := project.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := project.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.UsageCPU(); !ok {
		v := project.DefaultUsageCPU
		pc.mutation.SetUsageCPU(v)
	}
	if _, ok := pc.mutation.UsageRAM(); !ok {
		v := project.DefaultUsageRAM
		pc.mutation.SetUsageRAM(v)
	}
	if _, ok := pc.mutation.UsageDisk(); !ok {
		v := project.DefaultUsageDisk
		pc.mutation.SetUsageDisk(v)
	}
	if _, ok := pc.mutation.UsageNetwork(); !ok {
		v := project.DefaultUsageNetwork
		pc.mutation.SetUsageNetwork(v)
	}
	if _, ok := pc.mutation.UsageRouter(); !ok {
		v := project.DefaultUsageRouter
		pc.mutation.SetUsageRouter(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := project.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Project.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Project.updated_at"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Project.name"`)}
	}
	if _, ok := pc.mutation.QuotaCPU(); !ok {
		return &ValidationError{Name: "quota_cpu", err: errors.New(`ent: missing required field "Project.quota_cpu"`)}
	}
	if _, ok := pc.mutation.UsageCPU(); !ok {
		return &ValidationError{Name: "usage_cpu", err: errors.New(`ent: missing required field "Project.usage_cpu"`)}
	}
	if _, ok := pc.mutation.QuotaRAM(); !ok {
		return &ValidationError{Name: "quota_ram", err: errors.New(`ent: missing required field "Project.quota_ram"`)}
	}
	if _, ok := pc.mutation.UsageRAM(); !ok {
		return &ValidationError{Name: "usage_ram", err: errors.New(`ent: missing required field "Project.usage_ram"`)}
	}
	if _, ok := pc.mutation.QuotaDisk(); !ok {
		return &ValidationError{Name: "quota_disk", err: errors.New(`ent: missing required field "Project.quota_disk"`)}
	}
	if _, ok := pc.mutation.UsageDisk(); !ok {
		return &ValidationError{Name: "usage_disk", err: errors.New(`ent: missing required field "Project.usage_disk"`)}
	}
	if _, ok := pc.mutation.QuotaNetwork(); !ok {
		return &ValidationError{Name: "quota_network", err: errors.New(`ent: missing required field "Project.quota_network"`)}
	}
	if _, ok := pc.mutation.UsageNetwork(); !ok {
		return &ValidationError{Name: "usage_network", err: errors.New(`ent: missing required field "Project.usage_network"`)}
	}
	if _, ok := pc.mutation.QuotaRouter(); !ok {
		return &ValidationError{Name: "quota_router", err: errors.New(`ent: missing required field "Project.quota_router"`)}
	}
	if _, ok := pc.mutation.UsageRouter(); !ok {
		return &ValidationError{Name: "usage_router", err: errors.New(`ent: missing required field "Project.usage_router"`)}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(project.Table, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(project.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.QuotaCPU(); ok {
		_spec.SetField(project.FieldQuotaCPU, field.TypeInt, value)
		_node.QuotaCPU = value
	}
	if value, ok := pc.mutation.UsageCPU(); ok {
		_spec.SetField(project.FieldUsageCPU, field.TypeInt, value)
		_node.UsageCPU = value
	}
	if value, ok := pc.mutation.QuotaRAM(); ok {
		_spec.SetField(project.FieldQuotaRAM, field.TypeInt, value)
		_node.QuotaRAM = value
	}
	if value, ok := pc.mutation.UsageRAM(); ok {
		_spec.SetField(project.FieldUsageRAM, field.TypeInt, value)
		_node.UsageRAM = value
	}
	if value, ok := pc.mutation.QuotaDisk(); ok {
		_spec.SetField(project.FieldQuotaDisk, field.TypeInt, value)
		_node.QuotaDisk = value
	}
	if value, ok := pc.mutation.UsageDisk(); ok {
		_spec.SetField(project.FieldUsageDisk, field.TypeInt, value)
		_node.UsageDisk = value
	}
	if value, ok := pc.mutation.QuotaNetwork(); ok {
		_spec.SetField(project.FieldQuotaNetwork, field.TypeInt, value)
		_node.QuotaNetwork = value
	}
	if value, ok := pc.mutation.UsageNetwork(); ok {
		_spec.SetField(project.FieldUsageNetwork, field.TypeInt, value)
		_node.UsageNetwork = value
	}
	if value, ok := pc.mutation.QuotaRouter(); ok {
		_spec.SetField(project.FieldQuotaRouter, field.TypeInt, value)
		_node.QuotaRouter = value
	}
	if value, ok := pc.mutation.UsageRouter(); ok {
		_spec.SetField(project.FieldUsageRouter, field.TypeInt, value)
		_node.UsageRouter = value
	}
	if nodes := pc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.MembersTable,
			Columns: project.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &MembershipCreate{config: pc.config, mutation: newMembershipMutation(pc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.GroupMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.GroupMembersTable,
			Columns: project.GroupMembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &GroupMembershipCreate{config: pc.config, mutation: newGroupMembershipMutation(pc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BlueprintsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   project.BlueprintsTable,
			Columns: []string{project.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   project.DeploymentsTable,
			Columns: []string{project.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	err      error
	builders []*ProjectCreate
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
