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
	"github.com/cble-platform/cble/backend/engine/models"
	"github.com/cble-platform/cble/backend/ent/blueprint"
	"github.com/cble-platform/cble/backend/ent/deployment"
	"github.com/cble-platform/cble/backend/ent/predicate"
	"github.com/cble-platform/cble/backend/ent/project"
	entprovider "github.com/cble-platform/cble/backend/ent/provider"
	"github.com/cble-platform/cble/backend/ent/resource"
	"github.com/google/uuid"
)

// BlueprintUpdate is the builder for updating Blueprint entities.
type BlueprintUpdate struct {
	config
	hooks    []Hook
	mutation *BlueprintMutation
}

// Where appends a list predicates to the BlueprintUpdate builder.
func (bu *BlueprintUpdate) Where(ps ...predicate.Blueprint) *BlueprintUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BlueprintUpdate) SetUpdatedAt(t time.Time) *BlueprintUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetName sets the "name" field.
func (bu *BlueprintUpdate) SetName(s string) *BlueprintUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bu *BlueprintUpdate) SetNillableName(s *string) *BlueprintUpdate {
	if s != nil {
		bu.SetName(*s)
	}
	return bu
}

// SetDescription sets the "description" field.
func (bu *BlueprintUpdate) SetDescription(s string) *BlueprintUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bu *BlueprintUpdate) SetNillableDescription(s *string) *BlueprintUpdate {
	if s != nil {
		bu.SetDescription(*s)
	}
	return bu
}

// SetBlueprintTemplate sets the "blueprint_template" field.
func (bu *BlueprintUpdate) SetBlueprintTemplate(b []byte) *BlueprintUpdate {
	bu.mutation.SetBlueprintTemplate(b)
	return bu
}

// SetVariableTypes sets the "variable_types" field.
func (bu *BlueprintUpdate) SetVariableTypes(mvt map[string]models.BlueprintVariableType) *BlueprintUpdate {
	bu.mutation.SetVariableTypes(mvt)
	return bu
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (bu *BlueprintUpdate) SetProviderID(id uuid.UUID) *BlueprintUpdate {
	bu.mutation.SetProviderID(id)
	return bu
}

// SetProvider sets the "provider" edge to the Provider entity.
func (bu *BlueprintUpdate) SetProvider(p *Provider) *BlueprintUpdate {
	return bu.SetProviderID(p.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (bu *BlueprintUpdate) SetProjectID(id uuid.UUID) *BlueprintUpdate {
	bu.mutation.SetProjectID(id)
	return bu
}

// SetProject sets the "project" edge to the Project entity.
func (bu *BlueprintUpdate) SetProject(p *Project) *BlueprintUpdate {
	return bu.SetProjectID(p.ID)
}

// AddResourceIDs adds the "resources" edge to the Resource entity by IDs.
func (bu *BlueprintUpdate) AddResourceIDs(ids ...uuid.UUID) *BlueprintUpdate {
	bu.mutation.AddResourceIDs(ids...)
	return bu
}

// AddResources adds the "resources" edges to the Resource entity.
func (bu *BlueprintUpdate) AddResources(r ...*Resource) *BlueprintUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return bu.AddResourceIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (bu *BlueprintUpdate) AddDeploymentIDs(ids ...uuid.UUID) *BlueprintUpdate {
	bu.mutation.AddDeploymentIDs(ids...)
	return bu
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (bu *BlueprintUpdate) AddDeployments(d ...*Deployment) *BlueprintUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return bu.AddDeploymentIDs(ids...)
}

// Mutation returns the BlueprintMutation object of the builder.
func (bu *BlueprintUpdate) Mutation() *BlueprintMutation {
	return bu.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (bu *BlueprintUpdate) ClearProvider() *BlueprintUpdate {
	bu.mutation.ClearProvider()
	return bu
}

// ClearProject clears the "project" edge to the Project entity.
func (bu *BlueprintUpdate) ClearProject() *BlueprintUpdate {
	bu.mutation.ClearProject()
	return bu
}

// ClearResources clears all "resources" edges to the Resource entity.
func (bu *BlueprintUpdate) ClearResources() *BlueprintUpdate {
	bu.mutation.ClearResources()
	return bu
}

// RemoveResourceIDs removes the "resources" edge to Resource entities by IDs.
func (bu *BlueprintUpdate) RemoveResourceIDs(ids ...uuid.UUID) *BlueprintUpdate {
	bu.mutation.RemoveResourceIDs(ids...)
	return bu
}

// RemoveResources removes "resources" edges to Resource entities.
func (bu *BlueprintUpdate) RemoveResources(r ...*Resource) *BlueprintUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return bu.RemoveResourceIDs(ids...)
}

// ClearDeployments clears all "deployments" edges to the Deployment entity.
func (bu *BlueprintUpdate) ClearDeployments() *BlueprintUpdate {
	bu.mutation.ClearDeployments()
	return bu
}

// RemoveDeploymentIDs removes the "deployments" edge to Deployment entities by IDs.
func (bu *BlueprintUpdate) RemoveDeploymentIDs(ids ...uuid.UUID) *BlueprintUpdate {
	bu.mutation.RemoveDeploymentIDs(ids...)
	return bu
}

// RemoveDeployments removes "deployments" edges to Deployment entities.
func (bu *BlueprintUpdate) RemoveDeployments(d ...*Deployment) *BlueprintUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return bu.RemoveDeploymentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlueprintUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlueprintUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlueprintUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlueprintUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BlueprintUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := blueprint.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlueprintUpdate) check() error {
	if _, ok := bu.mutation.ProviderID(); bu.mutation.ProviderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Blueprint.provider"`)
	}
	if _, ok := bu.mutation.ProjectID(); bu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Blueprint.project"`)
	}
	return nil
}

func (bu *BlueprintUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(blueprint.Table, blueprint.Columns, sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(blueprint.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(blueprint.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.SetField(blueprint.FieldDescription, field.TypeString, value)
	}
	if value, ok := bu.mutation.BlueprintTemplate(); ok {
		_spec.SetField(blueprint.FieldBlueprintTemplate, field.TypeBytes, value)
	}
	if value, ok := bu.mutation.VariableTypes(); ok {
		_spec.SetField(blueprint.FieldVariableTypes, field.TypeJSON, value)
	}
	if bu.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProviderTable,
			Columns: []string{blueprint.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entprovider.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProviderTable,
			Columns: []string{blueprint.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entprovider.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProjectTable,
			Columns: []string{blueprint.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProjectTable,
			Columns: []string{blueprint.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.ResourcesTable,
			Columns: []string{blueprint.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !bu.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.ResourcesTable,
			Columns: []string{blueprint.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.ResourcesTable,
			Columns: []string{blueprint.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.DeploymentsTable,
			Columns: []string{blueprint.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedDeploymentsIDs(); len(nodes) > 0 && !bu.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.DeploymentsTable,
			Columns: []string{blueprint.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.DeploymentsTable,
			Columns: []string{blueprint.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blueprint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlueprintUpdateOne is the builder for updating a single Blueprint entity.
type BlueprintUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlueprintMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BlueprintUpdateOne) SetUpdatedAt(t time.Time) *BlueprintUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetName sets the "name" field.
func (buo *BlueprintUpdateOne) SetName(s string) *BlueprintUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (buo *BlueprintUpdateOne) SetNillableName(s *string) *BlueprintUpdateOne {
	if s != nil {
		buo.SetName(*s)
	}
	return buo
}

// SetDescription sets the "description" field.
func (buo *BlueprintUpdateOne) SetDescription(s string) *BlueprintUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (buo *BlueprintUpdateOne) SetNillableDescription(s *string) *BlueprintUpdateOne {
	if s != nil {
		buo.SetDescription(*s)
	}
	return buo
}

// SetBlueprintTemplate sets the "blueprint_template" field.
func (buo *BlueprintUpdateOne) SetBlueprintTemplate(b []byte) *BlueprintUpdateOne {
	buo.mutation.SetBlueprintTemplate(b)
	return buo
}

// SetVariableTypes sets the "variable_types" field.
func (buo *BlueprintUpdateOne) SetVariableTypes(mvt map[string]models.BlueprintVariableType) *BlueprintUpdateOne {
	buo.mutation.SetVariableTypes(mvt)
	return buo
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (buo *BlueprintUpdateOne) SetProviderID(id uuid.UUID) *BlueprintUpdateOne {
	buo.mutation.SetProviderID(id)
	return buo
}

// SetProvider sets the "provider" edge to the Provider entity.
func (buo *BlueprintUpdateOne) SetProvider(p *Provider) *BlueprintUpdateOne {
	return buo.SetProviderID(p.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (buo *BlueprintUpdateOne) SetProjectID(id uuid.UUID) *BlueprintUpdateOne {
	buo.mutation.SetProjectID(id)
	return buo
}

// SetProject sets the "project" edge to the Project entity.
func (buo *BlueprintUpdateOne) SetProject(p *Project) *BlueprintUpdateOne {
	return buo.SetProjectID(p.ID)
}

// AddResourceIDs adds the "resources" edge to the Resource entity by IDs.
func (buo *BlueprintUpdateOne) AddResourceIDs(ids ...uuid.UUID) *BlueprintUpdateOne {
	buo.mutation.AddResourceIDs(ids...)
	return buo
}

// AddResources adds the "resources" edges to the Resource entity.
func (buo *BlueprintUpdateOne) AddResources(r ...*Resource) *BlueprintUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return buo.AddResourceIDs(ids...)
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (buo *BlueprintUpdateOne) AddDeploymentIDs(ids ...uuid.UUID) *BlueprintUpdateOne {
	buo.mutation.AddDeploymentIDs(ids...)
	return buo
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (buo *BlueprintUpdateOne) AddDeployments(d ...*Deployment) *BlueprintUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return buo.AddDeploymentIDs(ids...)
}

// Mutation returns the BlueprintMutation object of the builder.
func (buo *BlueprintUpdateOne) Mutation() *BlueprintMutation {
	return buo.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (buo *BlueprintUpdateOne) ClearProvider() *BlueprintUpdateOne {
	buo.mutation.ClearProvider()
	return buo
}

// ClearProject clears the "project" edge to the Project entity.
func (buo *BlueprintUpdateOne) ClearProject() *BlueprintUpdateOne {
	buo.mutation.ClearProject()
	return buo
}

// ClearResources clears all "resources" edges to the Resource entity.
func (buo *BlueprintUpdateOne) ClearResources() *BlueprintUpdateOne {
	buo.mutation.ClearResources()
	return buo
}

// RemoveResourceIDs removes the "resources" edge to Resource entities by IDs.
func (buo *BlueprintUpdateOne) RemoveResourceIDs(ids ...uuid.UUID) *BlueprintUpdateOne {
	buo.mutation.RemoveResourceIDs(ids...)
	return buo
}

// RemoveResources removes "resources" edges to Resource entities.
func (buo *BlueprintUpdateOne) RemoveResources(r ...*Resource) *BlueprintUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return buo.RemoveResourceIDs(ids...)
}

// ClearDeployments clears all "deployments" edges to the Deployment entity.
func (buo *BlueprintUpdateOne) ClearDeployments() *BlueprintUpdateOne {
	buo.mutation.ClearDeployments()
	return buo
}

// RemoveDeploymentIDs removes the "deployments" edge to Deployment entities by IDs.
func (buo *BlueprintUpdateOne) RemoveDeploymentIDs(ids ...uuid.UUID) *BlueprintUpdateOne {
	buo.mutation.RemoveDeploymentIDs(ids...)
	return buo
}

// RemoveDeployments removes "deployments" edges to Deployment entities.
func (buo *BlueprintUpdateOne) RemoveDeployments(d ...*Deployment) *BlueprintUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return buo.RemoveDeploymentIDs(ids...)
}

// Where appends a list predicates to the BlueprintUpdate builder.
func (buo *BlueprintUpdateOne) Where(ps ...predicate.Blueprint) *BlueprintUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlueprintUpdateOne) Select(field string, fields ...string) *BlueprintUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blueprint entity.
func (buo *BlueprintUpdateOne) Save(ctx context.Context) (*Blueprint, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlueprintUpdateOne) SaveX(ctx context.Context) *Blueprint {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlueprintUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlueprintUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BlueprintUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := blueprint.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlueprintUpdateOne) check() error {
	if _, ok := buo.mutation.ProviderID(); buo.mutation.ProviderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Blueprint.provider"`)
	}
	if _, ok := buo.mutation.ProjectID(); buo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Blueprint.project"`)
	}
	return nil
}

func (buo *BlueprintUpdateOne) sqlSave(ctx context.Context) (_node *Blueprint, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(blueprint.Table, blueprint.Columns, sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blueprint.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blueprint.FieldID)
		for _, f := range fields {
			if !blueprint.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blueprint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(blueprint.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(blueprint.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.SetField(blueprint.FieldDescription, field.TypeString, value)
	}
	if value, ok := buo.mutation.BlueprintTemplate(); ok {
		_spec.SetField(blueprint.FieldBlueprintTemplate, field.TypeBytes, value)
	}
	if value, ok := buo.mutation.VariableTypes(); ok {
		_spec.SetField(blueprint.FieldVariableTypes, field.TypeJSON, value)
	}
	if buo.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProviderTable,
			Columns: []string{blueprint.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entprovider.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProviderTable,
			Columns: []string{blueprint.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entprovider.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProjectTable,
			Columns: []string{blueprint.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blueprint.ProjectTable,
			Columns: []string{blueprint.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.ResourcesTable,
			Columns: []string{blueprint.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedResourcesIDs(); len(nodes) > 0 && !buo.mutation.ResourcesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.ResourcesTable,
			Columns: []string{blueprint.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.ResourcesTable,
			Columns: []string{blueprint.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.DeploymentsTable,
			Columns: []string{blueprint.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedDeploymentsIDs(); len(nodes) > 0 && !buo.mutation.DeploymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.DeploymentsTable,
			Columns: []string{blueprint.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blueprint.DeploymentsTable,
			Columns: []string{blueprint.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Blueprint{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blueprint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
