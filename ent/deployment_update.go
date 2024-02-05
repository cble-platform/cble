// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/blueprint"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/ent/predicate"
	"github.com/cble-platform/cble-backend/ent/user"
	"github.com/google/uuid"
)

// DeploymentUpdate is the builder for updating Deployment entities.
type DeploymentUpdate struct {
	config
	hooks    []Hook
	mutation *DeploymentMutation
}

// Where appends a list predicates to the DeploymentUpdate builder.
func (du *DeploymentUpdate) Where(ps ...predicate.Deployment) *DeploymentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DeploymentUpdate) SetUpdatedAt(t time.Time) *DeploymentUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetName sets the "name" field.
func (du *DeploymentUpdate) SetName(s string) *DeploymentUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableName(s *string) *DeploymentUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetDescription sets the "description" field.
func (du *DeploymentUpdate) SetDescription(s string) *DeploymentUpdate {
	du.mutation.SetDescription(s)
	return du
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableDescription(s *string) *DeploymentUpdate {
	if s != nil {
		du.SetDescription(*s)
	}
	return du
}

// SetState sets the "state" field.
func (du *DeploymentUpdate) SetState(d deployment.State) *DeploymentUpdate {
	du.mutation.SetState(d)
	return du
}

// SetNillableState sets the "state" field if the given value is not nil.
func (du *DeploymentUpdate) SetNillableState(d *deployment.State) *DeploymentUpdate {
	if d != nil {
		du.SetState(*d)
	}
	return du
}

// SetTemplateVars sets the "template_vars" field.
func (du *DeploymentUpdate) SetTemplateVars(m map[string]string) *DeploymentUpdate {
	du.mutation.SetTemplateVars(m)
	return du
}

// SetBlueprintID sets the "blueprint" edge to the Blueprint entity by ID.
func (du *DeploymentUpdate) SetBlueprintID(id uuid.UUID) *DeploymentUpdate {
	du.mutation.SetBlueprintID(id)
	return du
}

// SetBlueprint sets the "blueprint" edge to the Blueprint entity.
func (du *DeploymentUpdate) SetBlueprint(b *Blueprint) *DeploymentUpdate {
	return du.SetBlueprintID(b.ID)
}

// AddDeploymentNodeIDs adds the "deployment_nodes" edge to the DeploymentNode entity by IDs.
func (du *DeploymentUpdate) AddDeploymentNodeIDs(ids ...uuid.UUID) *DeploymentUpdate {
	du.mutation.AddDeploymentNodeIDs(ids...)
	return du
}

// AddDeploymentNodes adds the "deployment_nodes" edges to the DeploymentNode entity.
func (du *DeploymentUpdate) AddDeploymentNodes(d ...*DeploymentNode) *DeploymentUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDeploymentNodeIDs(ids...)
}

// SetRequesterID sets the "requester" edge to the User entity by ID.
func (du *DeploymentUpdate) SetRequesterID(id uuid.UUID) *DeploymentUpdate {
	du.mutation.SetRequesterID(id)
	return du
}

// SetRequester sets the "requester" edge to the User entity.
func (du *DeploymentUpdate) SetRequester(u *User) *DeploymentUpdate {
	return du.SetRequesterID(u.ID)
}

// Mutation returns the DeploymentMutation object of the builder.
func (du *DeploymentUpdate) Mutation() *DeploymentMutation {
	return du.mutation
}

// ClearBlueprint clears the "blueprint" edge to the Blueprint entity.
func (du *DeploymentUpdate) ClearBlueprint() *DeploymentUpdate {
	du.mutation.ClearBlueprint()
	return du
}

// ClearDeploymentNodes clears all "deployment_nodes" edges to the DeploymentNode entity.
func (du *DeploymentUpdate) ClearDeploymentNodes() *DeploymentUpdate {
	du.mutation.ClearDeploymentNodes()
	return du
}

// RemoveDeploymentNodeIDs removes the "deployment_nodes" edge to DeploymentNode entities by IDs.
func (du *DeploymentUpdate) RemoveDeploymentNodeIDs(ids ...uuid.UUID) *DeploymentUpdate {
	du.mutation.RemoveDeploymentNodeIDs(ids...)
	return du
}

// RemoveDeploymentNodes removes "deployment_nodes" edges to DeploymentNode entities.
func (du *DeploymentUpdate) RemoveDeploymentNodes(d ...*DeploymentNode) *DeploymentUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDeploymentNodeIDs(ids...)
}

// ClearRequester clears the "requester" edge to the User entity.
func (du *DeploymentUpdate) ClearRequester() *DeploymentUpdate {
	du.mutation.ClearRequester()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeploymentUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeploymentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeploymentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeploymentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeploymentUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := deployment.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeploymentUpdate) check() error {
	if v, ok := du.mutation.State(); ok {
		if err := deployment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Deployment.state": %w`, err)}
		}
	}
	if _, ok := du.mutation.BlueprintID(); du.mutation.BlueprintCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deployment.blueprint"`)
	}
	if _, ok := du.mutation.RequesterID(); du.mutation.RequesterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deployment.requester"`)
	}
	return nil
}

func (du *DeploymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(deployment.Table, deployment.Columns, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(deployment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(deployment.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Description(); ok {
		_spec.SetField(deployment.FieldDescription, field.TypeString, value)
	}
	if value, ok := du.mutation.State(); ok {
		_spec.SetField(deployment.FieldState, field.TypeEnum, value)
	}
	if value, ok := du.mutation.TemplateVars(); ok {
		_spec.SetField(deployment.FieldTemplateVars, field.TypeJSON, value)
	}
	if du.mutation.BlueprintCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.BlueprintTable,
			Columns: []string{deployment.BlueprintColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.BlueprintIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.BlueprintTable,
			Columns: []string{deployment.BlueprintColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.DeploymentNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   deployment.DeploymentNodesTable,
			Columns: []string{deployment.DeploymentNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDeploymentNodesIDs(); len(nodes) > 0 && !du.mutation.DeploymentNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   deployment.DeploymentNodesTable,
			Columns: []string{deployment.DeploymentNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DeploymentNodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   deployment.DeploymentNodesTable,
			Columns: []string{deployment.DeploymentNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.RequesterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.RequesterTable,
			Columns: []string{deployment.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RequesterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.RequesterTable,
			Columns: []string{deployment.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeploymentUpdateOne is the builder for updating a single Deployment entity.
type DeploymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeploymentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DeploymentUpdateOne) SetUpdatedAt(t time.Time) *DeploymentUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetName sets the "name" field.
func (duo *DeploymentUpdateOne) SetName(s string) *DeploymentUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableName(s *string) *DeploymentUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetDescription sets the "description" field.
func (duo *DeploymentUpdateOne) SetDescription(s string) *DeploymentUpdateOne {
	duo.mutation.SetDescription(s)
	return duo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableDescription(s *string) *DeploymentUpdateOne {
	if s != nil {
		duo.SetDescription(*s)
	}
	return duo
}

// SetState sets the "state" field.
func (duo *DeploymentUpdateOne) SetState(d deployment.State) *DeploymentUpdateOne {
	duo.mutation.SetState(d)
	return duo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (duo *DeploymentUpdateOne) SetNillableState(d *deployment.State) *DeploymentUpdateOne {
	if d != nil {
		duo.SetState(*d)
	}
	return duo
}

// SetTemplateVars sets the "template_vars" field.
func (duo *DeploymentUpdateOne) SetTemplateVars(m map[string]string) *DeploymentUpdateOne {
	duo.mutation.SetTemplateVars(m)
	return duo
}

// SetBlueprintID sets the "blueprint" edge to the Blueprint entity by ID.
func (duo *DeploymentUpdateOne) SetBlueprintID(id uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.SetBlueprintID(id)
	return duo
}

// SetBlueprint sets the "blueprint" edge to the Blueprint entity.
func (duo *DeploymentUpdateOne) SetBlueprint(b *Blueprint) *DeploymentUpdateOne {
	return duo.SetBlueprintID(b.ID)
}

// AddDeploymentNodeIDs adds the "deployment_nodes" edge to the DeploymentNode entity by IDs.
func (duo *DeploymentUpdateOne) AddDeploymentNodeIDs(ids ...uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.AddDeploymentNodeIDs(ids...)
	return duo
}

// AddDeploymentNodes adds the "deployment_nodes" edges to the DeploymentNode entity.
func (duo *DeploymentUpdateOne) AddDeploymentNodes(d ...*DeploymentNode) *DeploymentUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDeploymentNodeIDs(ids...)
}

// SetRequesterID sets the "requester" edge to the User entity by ID.
func (duo *DeploymentUpdateOne) SetRequesterID(id uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.SetRequesterID(id)
	return duo
}

// SetRequester sets the "requester" edge to the User entity.
func (duo *DeploymentUpdateOne) SetRequester(u *User) *DeploymentUpdateOne {
	return duo.SetRequesterID(u.ID)
}

// Mutation returns the DeploymentMutation object of the builder.
func (duo *DeploymentUpdateOne) Mutation() *DeploymentMutation {
	return duo.mutation
}

// ClearBlueprint clears the "blueprint" edge to the Blueprint entity.
func (duo *DeploymentUpdateOne) ClearBlueprint() *DeploymentUpdateOne {
	duo.mutation.ClearBlueprint()
	return duo
}

// ClearDeploymentNodes clears all "deployment_nodes" edges to the DeploymentNode entity.
func (duo *DeploymentUpdateOne) ClearDeploymentNodes() *DeploymentUpdateOne {
	duo.mutation.ClearDeploymentNodes()
	return duo
}

// RemoveDeploymentNodeIDs removes the "deployment_nodes" edge to DeploymentNode entities by IDs.
func (duo *DeploymentUpdateOne) RemoveDeploymentNodeIDs(ids ...uuid.UUID) *DeploymentUpdateOne {
	duo.mutation.RemoveDeploymentNodeIDs(ids...)
	return duo
}

// RemoveDeploymentNodes removes "deployment_nodes" edges to DeploymentNode entities.
func (duo *DeploymentUpdateOne) RemoveDeploymentNodes(d ...*DeploymentNode) *DeploymentUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDeploymentNodeIDs(ids...)
}

// ClearRequester clears the "requester" edge to the User entity.
func (duo *DeploymentUpdateOne) ClearRequester() *DeploymentUpdateOne {
	duo.mutation.ClearRequester()
	return duo
}

// Where appends a list predicates to the DeploymentUpdate builder.
func (duo *DeploymentUpdateOne) Where(ps ...predicate.Deployment) *DeploymentUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeploymentUpdateOne) Select(field string, fields ...string) *DeploymentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Deployment entity.
func (duo *DeploymentUpdateOne) Save(ctx context.Context) (*Deployment, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeploymentUpdateOne) SaveX(ctx context.Context) *Deployment {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeploymentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeploymentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeploymentUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := deployment.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeploymentUpdateOne) check() error {
	if v, ok := duo.mutation.State(); ok {
		if err := deployment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Deployment.state": %w`, err)}
		}
	}
	if _, ok := duo.mutation.BlueprintID(); duo.mutation.BlueprintCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deployment.blueprint"`)
	}
	if _, ok := duo.mutation.RequesterID(); duo.mutation.RequesterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deployment.requester"`)
	}
	return nil
}

func (duo *DeploymentUpdateOne) sqlSave(ctx context.Context) (_node *Deployment, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(deployment.Table, deployment.Columns, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Deployment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deployment.FieldID)
		for _, f := range fields {
			if !deployment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deployment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(deployment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(deployment.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Description(); ok {
		_spec.SetField(deployment.FieldDescription, field.TypeString, value)
	}
	if value, ok := duo.mutation.State(); ok {
		_spec.SetField(deployment.FieldState, field.TypeEnum, value)
	}
	if value, ok := duo.mutation.TemplateVars(); ok {
		_spec.SetField(deployment.FieldTemplateVars, field.TypeJSON, value)
	}
	if duo.mutation.BlueprintCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.BlueprintTable,
			Columns: []string{deployment.BlueprintColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.BlueprintIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.BlueprintTable,
			Columns: []string{deployment.BlueprintColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.DeploymentNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   deployment.DeploymentNodesTable,
			Columns: []string{deployment.DeploymentNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDeploymentNodesIDs(); len(nodes) > 0 && !duo.mutation.DeploymentNodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   deployment.DeploymentNodesTable,
			Columns: []string{deployment.DeploymentNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DeploymentNodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   deployment.DeploymentNodesTable,
			Columns: []string{deployment.DeploymentNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.RequesterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.RequesterTable,
			Columns: []string{deployment.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RequesterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.RequesterTable,
			Columns: []string{deployment.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Deployment{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deployment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
