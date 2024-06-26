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
	"github.com/cble-platform/cble/backend/ent/grantedpermission"
	"github.com/cble-platform/cble/backend/ent/group"
	"github.com/cble-platform/cble/backend/ent/predicate"
	"github.com/cble-platform/cble/backend/ent/user"
	"github.com/cble-platform/cble/backend/permission/actions"
	"github.com/google/uuid"
)

// GrantedPermissionUpdate is the builder for updating GrantedPermission entities.
type GrantedPermissionUpdate struct {
	config
	hooks    []Hook
	mutation *GrantedPermissionMutation
}

// Where appends a list predicates to the GrantedPermissionUpdate builder.
func (gpu *GrantedPermissionUpdate) Where(ps ...predicate.GrantedPermission) *GrantedPermissionUpdate {
	gpu.mutation.Where(ps...)
	return gpu
}

// SetUpdatedAt sets the "updated_at" field.
func (gpu *GrantedPermissionUpdate) SetUpdatedAt(t time.Time) *GrantedPermissionUpdate {
	gpu.mutation.SetUpdatedAt(t)
	return gpu
}

// SetSubjectType sets the "subject_type" field.
func (gpu *GrantedPermissionUpdate) SetSubjectType(gt grantedpermission.SubjectType) *GrantedPermissionUpdate {
	gpu.mutation.SetSubjectType(gt)
	return gpu
}

// SetNillableSubjectType sets the "subject_type" field if the given value is not nil.
func (gpu *GrantedPermissionUpdate) SetNillableSubjectType(gt *grantedpermission.SubjectType) *GrantedPermissionUpdate {
	if gt != nil {
		gpu.SetSubjectType(*gt)
	}
	return gpu
}

// SetSubjectID sets the "subject_id" field.
func (gpu *GrantedPermissionUpdate) SetSubjectID(u uuid.UUID) *GrantedPermissionUpdate {
	gpu.mutation.SetSubjectID(u)
	return gpu
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (gpu *GrantedPermissionUpdate) SetNillableSubjectID(u *uuid.UUID) *GrantedPermissionUpdate {
	if u != nil {
		gpu.SetSubjectID(*u)
	}
	return gpu
}

// SetObjectType sets the "object_type" field.
func (gpu *GrantedPermissionUpdate) SetObjectType(gt grantedpermission.ObjectType) *GrantedPermissionUpdate {
	gpu.mutation.SetObjectType(gt)
	return gpu
}

// SetNillableObjectType sets the "object_type" field if the given value is not nil.
func (gpu *GrantedPermissionUpdate) SetNillableObjectType(gt *grantedpermission.ObjectType) *GrantedPermissionUpdate {
	if gt != nil {
		gpu.SetObjectType(*gt)
	}
	return gpu
}

// SetObjectID sets the "object_id" field.
func (gpu *GrantedPermissionUpdate) SetObjectID(u uuid.UUID) *GrantedPermissionUpdate {
	gpu.mutation.SetObjectID(u)
	return gpu
}

// SetNillableObjectID sets the "object_id" field if the given value is not nil.
func (gpu *GrantedPermissionUpdate) SetNillableObjectID(u *uuid.UUID) *GrantedPermissionUpdate {
	if u != nil {
		gpu.SetObjectID(*u)
	}
	return gpu
}

// SetAction sets the "action" field.
func (gpu *GrantedPermissionUpdate) SetAction(aa actions.PermissionAction) *GrantedPermissionUpdate {
	gpu.mutation.SetAction(aa)
	return gpu
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (gpu *GrantedPermissionUpdate) SetNillableAction(aa *actions.PermissionAction) *GrantedPermissionUpdate {
	if aa != nil {
		gpu.SetAction(*aa)
	}
	return gpu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gpu *GrantedPermissionUpdate) SetUserID(id uuid.UUID) *GrantedPermissionUpdate {
	gpu.mutation.SetUserID(id)
	return gpu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (gpu *GrantedPermissionUpdate) SetNillableUserID(id *uuid.UUID) *GrantedPermissionUpdate {
	if id != nil {
		gpu = gpu.SetUserID(*id)
	}
	return gpu
}

// SetUser sets the "user" edge to the User entity.
func (gpu *GrantedPermissionUpdate) SetUser(u *User) *GrantedPermissionUpdate {
	return gpu.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gpu *GrantedPermissionUpdate) SetGroupID(id uuid.UUID) *GrantedPermissionUpdate {
	gpu.mutation.SetGroupID(id)
	return gpu
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gpu *GrantedPermissionUpdate) SetNillableGroupID(id *uuid.UUID) *GrantedPermissionUpdate {
	if id != nil {
		gpu = gpu.SetGroupID(*id)
	}
	return gpu
}

// SetGroup sets the "group" edge to the Group entity.
func (gpu *GrantedPermissionUpdate) SetGroup(g *Group) *GrantedPermissionUpdate {
	return gpu.SetGroupID(g.ID)
}

// Mutation returns the GrantedPermissionMutation object of the builder.
func (gpu *GrantedPermissionUpdate) Mutation() *GrantedPermissionMutation {
	return gpu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gpu *GrantedPermissionUpdate) ClearUser() *GrantedPermissionUpdate {
	gpu.mutation.ClearUser()
	return gpu
}

// ClearGroup clears the "group" edge to the Group entity.
func (gpu *GrantedPermissionUpdate) ClearGroup() *GrantedPermissionUpdate {
	gpu.mutation.ClearGroup()
	return gpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gpu *GrantedPermissionUpdate) Save(ctx context.Context) (int, error) {
	gpu.defaults()
	return withHooks(ctx, gpu.sqlSave, gpu.mutation, gpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gpu *GrantedPermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := gpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gpu *GrantedPermissionUpdate) Exec(ctx context.Context) error {
	_, err := gpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpu *GrantedPermissionUpdate) ExecX(ctx context.Context) {
	if err := gpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gpu *GrantedPermissionUpdate) defaults() {
	if _, ok := gpu.mutation.UpdatedAt(); !ok {
		v := grantedpermission.UpdateDefaultUpdatedAt()
		gpu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gpu *GrantedPermissionUpdate) check() error {
	if v, ok := gpu.mutation.SubjectType(); ok {
		if err := grantedpermission.SubjectTypeValidator(v); err != nil {
			return &ValidationError{Name: "subject_type", err: fmt.Errorf(`ent: validator failed for field "GrantedPermission.subject_type": %w`, err)}
		}
	}
	if v, ok := gpu.mutation.ObjectType(); ok {
		if err := grantedpermission.ObjectTypeValidator(v); err != nil {
			return &ValidationError{Name: "object_type", err: fmt.Errorf(`ent: validator failed for field "GrantedPermission.object_type": %w`, err)}
		}
	}
	if v, ok := gpu.mutation.Action(); ok {
		if err := grantedpermission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "GrantedPermission.action": %w`, err)}
		}
	}
	return nil
}

func (gpu *GrantedPermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(grantedpermission.Table, grantedpermission.Columns, sqlgraph.NewFieldSpec(grantedpermission.FieldID, field.TypeUUID))
	if ps := gpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gpu.mutation.UpdatedAt(); ok {
		_spec.SetField(grantedpermission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gpu.mutation.SubjectType(); ok {
		_spec.SetField(grantedpermission.FieldSubjectType, field.TypeEnum, value)
	}
	if value, ok := gpu.mutation.SubjectID(); ok {
		_spec.SetField(grantedpermission.FieldSubjectID, field.TypeUUID, value)
	}
	if value, ok := gpu.mutation.ObjectType(); ok {
		_spec.SetField(grantedpermission.FieldObjectType, field.TypeEnum, value)
	}
	if value, ok := gpu.mutation.ObjectID(); ok {
		_spec.SetField(grantedpermission.FieldObjectID, field.TypeUUID, value)
	}
	if value, ok := gpu.mutation.Action(); ok {
		_spec.SetField(grantedpermission.FieldAction, field.TypeEnum, value)
	}
	if gpu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.UserTable,
			Columns: []string{grantedpermission.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.UserTable,
			Columns: []string{grantedpermission.UserColumn},
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
	if gpu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.GroupTable,
			Columns: []string{grantedpermission.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.GroupTable,
			Columns: []string{grantedpermission.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grantedpermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gpu.mutation.done = true
	return n, nil
}

// GrantedPermissionUpdateOne is the builder for updating a single GrantedPermission entity.
type GrantedPermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GrantedPermissionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gpuo *GrantedPermissionUpdateOne) SetUpdatedAt(t time.Time) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetUpdatedAt(t)
	return gpuo
}

// SetSubjectType sets the "subject_type" field.
func (gpuo *GrantedPermissionUpdateOne) SetSubjectType(gt grantedpermission.SubjectType) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetSubjectType(gt)
	return gpuo
}

// SetNillableSubjectType sets the "subject_type" field if the given value is not nil.
func (gpuo *GrantedPermissionUpdateOne) SetNillableSubjectType(gt *grantedpermission.SubjectType) *GrantedPermissionUpdateOne {
	if gt != nil {
		gpuo.SetSubjectType(*gt)
	}
	return gpuo
}

// SetSubjectID sets the "subject_id" field.
func (gpuo *GrantedPermissionUpdateOne) SetSubjectID(u uuid.UUID) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetSubjectID(u)
	return gpuo
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (gpuo *GrantedPermissionUpdateOne) SetNillableSubjectID(u *uuid.UUID) *GrantedPermissionUpdateOne {
	if u != nil {
		gpuo.SetSubjectID(*u)
	}
	return gpuo
}

// SetObjectType sets the "object_type" field.
func (gpuo *GrantedPermissionUpdateOne) SetObjectType(gt grantedpermission.ObjectType) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetObjectType(gt)
	return gpuo
}

// SetNillableObjectType sets the "object_type" field if the given value is not nil.
func (gpuo *GrantedPermissionUpdateOne) SetNillableObjectType(gt *grantedpermission.ObjectType) *GrantedPermissionUpdateOne {
	if gt != nil {
		gpuo.SetObjectType(*gt)
	}
	return gpuo
}

// SetObjectID sets the "object_id" field.
func (gpuo *GrantedPermissionUpdateOne) SetObjectID(u uuid.UUID) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetObjectID(u)
	return gpuo
}

// SetNillableObjectID sets the "object_id" field if the given value is not nil.
func (gpuo *GrantedPermissionUpdateOne) SetNillableObjectID(u *uuid.UUID) *GrantedPermissionUpdateOne {
	if u != nil {
		gpuo.SetObjectID(*u)
	}
	return gpuo
}

// SetAction sets the "action" field.
func (gpuo *GrantedPermissionUpdateOne) SetAction(aa actions.PermissionAction) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetAction(aa)
	return gpuo
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (gpuo *GrantedPermissionUpdateOne) SetNillableAction(aa *actions.PermissionAction) *GrantedPermissionUpdateOne {
	if aa != nil {
		gpuo.SetAction(*aa)
	}
	return gpuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gpuo *GrantedPermissionUpdateOne) SetUserID(id uuid.UUID) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetUserID(id)
	return gpuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (gpuo *GrantedPermissionUpdateOne) SetNillableUserID(id *uuid.UUID) *GrantedPermissionUpdateOne {
	if id != nil {
		gpuo = gpuo.SetUserID(*id)
	}
	return gpuo
}

// SetUser sets the "user" edge to the User entity.
func (gpuo *GrantedPermissionUpdateOne) SetUser(u *User) *GrantedPermissionUpdateOne {
	return gpuo.SetUserID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gpuo *GrantedPermissionUpdateOne) SetGroupID(id uuid.UUID) *GrantedPermissionUpdateOne {
	gpuo.mutation.SetGroupID(id)
	return gpuo
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gpuo *GrantedPermissionUpdateOne) SetNillableGroupID(id *uuid.UUID) *GrantedPermissionUpdateOne {
	if id != nil {
		gpuo = gpuo.SetGroupID(*id)
	}
	return gpuo
}

// SetGroup sets the "group" edge to the Group entity.
func (gpuo *GrantedPermissionUpdateOne) SetGroup(g *Group) *GrantedPermissionUpdateOne {
	return gpuo.SetGroupID(g.ID)
}

// Mutation returns the GrantedPermissionMutation object of the builder.
func (gpuo *GrantedPermissionUpdateOne) Mutation() *GrantedPermissionMutation {
	return gpuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (gpuo *GrantedPermissionUpdateOne) ClearUser() *GrantedPermissionUpdateOne {
	gpuo.mutation.ClearUser()
	return gpuo
}

// ClearGroup clears the "group" edge to the Group entity.
func (gpuo *GrantedPermissionUpdateOne) ClearGroup() *GrantedPermissionUpdateOne {
	gpuo.mutation.ClearGroup()
	return gpuo
}

// Where appends a list predicates to the GrantedPermissionUpdate builder.
func (gpuo *GrantedPermissionUpdateOne) Where(ps ...predicate.GrantedPermission) *GrantedPermissionUpdateOne {
	gpuo.mutation.Where(ps...)
	return gpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gpuo *GrantedPermissionUpdateOne) Select(field string, fields ...string) *GrantedPermissionUpdateOne {
	gpuo.fields = append([]string{field}, fields...)
	return gpuo
}

// Save executes the query and returns the updated GrantedPermission entity.
func (gpuo *GrantedPermissionUpdateOne) Save(ctx context.Context) (*GrantedPermission, error) {
	gpuo.defaults()
	return withHooks(ctx, gpuo.sqlSave, gpuo.mutation, gpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gpuo *GrantedPermissionUpdateOne) SaveX(ctx context.Context) *GrantedPermission {
	node, err := gpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gpuo *GrantedPermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := gpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpuo *GrantedPermissionUpdateOne) ExecX(ctx context.Context) {
	if err := gpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gpuo *GrantedPermissionUpdateOne) defaults() {
	if _, ok := gpuo.mutation.UpdatedAt(); !ok {
		v := grantedpermission.UpdateDefaultUpdatedAt()
		gpuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gpuo *GrantedPermissionUpdateOne) check() error {
	if v, ok := gpuo.mutation.SubjectType(); ok {
		if err := grantedpermission.SubjectTypeValidator(v); err != nil {
			return &ValidationError{Name: "subject_type", err: fmt.Errorf(`ent: validator failed for field "GrantedPermission.subject_type": %w`, err)}
		}
	}
	if v, ok := gpuo.mutation.ObjectType(); ok {
		if err := grantedpermission.ObjectTypeValidator(v); err != nil {
			return &ValidationError{Name: "object_type", err: fmt.Errorf(`ent: validator failed for field "GrantedPermission.object_type": %w`, err)}
		}
	}
	if v, ok := gpuo.mutation.Action(); ok {
		if err := grantedpermission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "GrantedPermission.action": %w`, err)}
		}
	}
	return nil
}

func (gpuo *GrantedPermissionUpdateOne) sqlSave(ctx context.Context) (_node *GrantedPermission, err error) {
	if err := gpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(grantedpermission.Table, grantedpermission.Columns, sqlgraph.NewFieldSpec(grantedpermission.FieldID, field.TypeUUID))
	id, ok := gpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GrantedPermission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grantedpermission.FieldID)
		for _, f := range fields {
			if !grantedpermission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != grantedpermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(grantedpermission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gpuo.mutation.SubjectType(); ok {
		_spec.SetField(grantedpermission.FieldSubjectType, field.TypeEnum, value)
	}
	if value, ok := gpuo.mutation.SubjectID(); ok {
		_spec.SetField(grantedpermission.FieldSubjectID, field.TypeUUID, value)
	}
	if value, ok := gpuo.mutation.ObjectType(); ok {
		_spec.SetField(grantedpermission.FieldObjectType, field.TypeEnum, value)
	}
	if value, ok := gpuo.mutation.ObjectID(); ok {
		_spec.SetField(grantedpermission.FieldObjectID, field.TypeUUID, value)
	}
	if value, ok := gpuo.mutation.Action(); ok {
		_spec.SetField(grantedpermission.FieldAction, field.TypeEnum, value)
	}
	if gpuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.UserTable,
			Columns: []string{grantedpermission.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.UserTable,
			Columns: []string{grantedpermission.UserColumn},
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
	if gpuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.GroupTable,
			Columns: []string{grantedpermission.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grantedpermission.GroupTable,
			Columns: []string{grantedpermission.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GrantedPermission{config: gpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grantedpermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gpuo.mutation.done = true
	return _node, nil
}
