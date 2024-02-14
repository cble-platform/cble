// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/grantedpermission"
	"github.com/cble-platform/cble-backend/ent/group"
	"github.com/cble-platform/cble-backend/ent/predicate"
	"github.com/cble-platform/cble-backend/ent/user"
	"github.com/google/uuid"
)

// GrantedPermissionQuery is the builder for querying GrantedPermission entities.
type GrantedPermissionQuery struct {
	config
	ctx        *QueryContext
	order      []grantedpermission.OrderOption
	inters     []Interceptor
	predicates []predicate.GrantedPermission
	withUser   *UserQuery
	withGroup  *GroupQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GrantedPermissionQuery builder.
func (gpq *GrantedPermissionQuery) Where(ps ...predicate.GrantedPermission) *GrantedPermissionQuery {
	gpq.predicates = append(gpq.predicates, ps...)
	return gpq
}

// Limit the number of records to be returned by this query.
func (gpq *GrantedPermissionQuery) Limit(limit int) *GrantedPermissionQuery {
	gpq.ctx.Limit = &limit
	return gpq
}

// Offset to start from.
func (gpq *GrantedPermissionQuery) Offset(offset int) *GrantedPermissionQuery {
	gpq.ctx.Offset = &offset
	return gpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gpq *GrantedPermissionQuery) Unique(unique bool) *GrantedPermissionQuery {
	gpq.ctx.Unique = &unique
	return gpq
}

// Order specifies how the records should be ordered.
func (gpq *GrantedPermissionQuery) Order(o ...grantedpermission.OrderOption) *GrantedPermissionQuery {
	gpq.order = append(gpq.order, o...)
	return gpq
}

// QueryUser chains the current query on the "user" edge.
func (gpq *GrantedPermissionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: gpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grantedpermission.Table, grantedpermission.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, grantedpermission.UserTable, grantedpermission.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(gpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (gpq *GrantedPermissionQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: gpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grantedpermission.Table, grantedpermission.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, grantedpermission.GroupTable, grantedpermission.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(gpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GrantedPermission entity from the query.
// Returns a *NotFoundError when no GrantedPermission was found.
func (gpq *GrantedPermissionQuery) First(ctx context.Context) (*GrantedPermission, error) {
	nodes, err := gpq.Limit(1).All(setContextOp(ctx, gpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grantedpermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) FirstX(ctx context.Context) *GrantedPermission {
	node, err := gpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GrantedPermission ID from the query.
// Returns a *NotFoundError when no GrantedPermission ID was found.
func (gpq *GrantedPermissionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(1).IDs(setContextOp(ctx, gpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grantedpermission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GrantedPermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GrantedPermission entity is found.
// Returns a *NotFoundError when no GrantedPermission entities are found.
func (gpq *GrantedPermissionQuery) Only(ctx context.Context) (*GrantedPermission, error) {
	nodes, err := gpq.Limit(2).All(setContextOp(ctx, gpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grantedpermission.Label}
	default:
		return nil, &NotSingularError{grantedpermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) OnlyX(ctx context.Context) *GrantedPermission {
	node, err := gpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GrantedPermission ID in the query.
// Returns a *NotSingularError when more than one GrantedPermission ID is found.
// Returns a *NotFoundError when no entities are found.
func (gpq *GrantedPermissionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(2).IDs(setContextOp(ctx, gpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grantedpermission.Label}
	default:
		err = &NotSingularError{grantedpermission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GrantedPermissions.
func (gpq *GrantedPermissionQuery) All(ctx context.Context) ([]*GrantedPermission, error) {
	ctx = setContextOp(ctx, gpq.ctx, "All")
	if err := gpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GrantedPermission, *GrantedPermissionQuery]()
	return withInterceptors[[]*GrantedPermission](ctx, gpq, qr, gpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) AllX(ctx context.Context) []*GrantedPermission {
	nodes, err := gpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GrantedPermission IDs.
func (gpq *GrantedPermissionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if gpq.ctx.Unique == nil && gpq.path != nil {
		gpq.Unique(true)
	}
	ctx = setContextOp(ctx, gpq.ctx, "IDs")
	if err = gpq.Select(grantedpermission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gpq *GrantedPermissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gpq.ctx, "Count")
	if err := gpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gpq, querierCount[*GrantedPermissionQuery](), gpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) CountX(ctx context.Context) int {
	count, err := gpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gpq *GrantedPermissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gpq.ctx, "Exist")
	switch _, err := gpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gpq *GrantedPermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := gpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GrantedPermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gpq *GrantedPermissionQuery) Clone() *GrantedPermissionQuery {
	if gpq == nil {
		return nil
	}
	return &GrantedPermissionQuery{
		config:     gpq.config,
		ctx:        gpq.ctx.Clone(),
		order:      append([]grantedpermission.OrderOption{}, gpq.order...),
		inters:     append([]Interceptor{}, gpq.inters...),
		predicates: append([]predicate.GrantedPermission{}, gpq.predicates...),
		withUser:   gpq.withUser.Clone(),
		withGroup:  gpq.withGroup.Clone(),
		// clone intermediate query.
		sql:  gpq.sql.Clone(),
		path: gpq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (gpq *GrantedPermissionQuery) WithUser(opts ...func(*UserQuery)) *GrantedPermissionQuery {
	query := (&UserClient{config: gpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gpq.withUser = query
	return gpq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (gpq *GrantedPermissionQuery) WithGroup(opts ...func(*GroupQuery)) *GrantedPermissionQuery {
	query := (&GroupClient{config: gpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gpq.withGroup = query
	return gpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GrantedPermission.Query().
//		GroupBy(grantedpermission.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gpq *GrantedPermissionQuery) GroupBy(field string, fields ...string) *GrantedPermissionGroupBy {
	gpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GrantedPermissionGroupBy{build: gpq}
	grbuild.flds = &gpq.ctx.Fields
	grbuild.label = grantedpermission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.GrantedPermission.Query().
//		Select(grantedpermission.FieldCreatedAt).
//		Scan(ctx, &v)
func (gpq *GrantedPermissionQuery) Select(fields ...string) *GrantedPermissionSelect {
	gpq.ctx.Fields = append(gpq.ctx.Fields, fields...)
	sbuild := &GrantedPermissionSelect{GrantedPermissionQuery: gpq}
	sbuild.label = grantedpermission.Label
	sbuild.flds, sbuild.scan = &gpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GrantedPermissionSelect configured with the given aggregations.
func (gpq *GrantedPermissionQuery) Aggregate(fns ...AggregateFunc) *GrantedPermissionSelect {
	return gpq.Select().Aggregate(fns...)
}

func (gpq *GrantedPermissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gpq); err != nil {
				return err
			}
		}
	}
	for _, f := range gpq.ctx.Fields {
		if !grantedpermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gpq.path != nil {
		prev, err := gpq.path(ctx)
		if err != nil {
			return err
		}
		gpq.sql = prev
	}
	return nil
}

func (gpq *GrantedPermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GrantedPermission, error) {
	var (
		nodes       = []*GrantedPermission{}
		withFKs     = gpq.withFKs
		_spec       = gpq.querySpec()
		loadedTypes = [2]bool{
			gpq.withUser != nil,
			gpq.withGroup != nil,
		}
	)
	if gpq.withUser != nil || gpq.withGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, grantedpermission.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GrantedPermission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GrantedPermission{config: gpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gpq.withUser; query != nil {
		if err := gpq.loadUser(ctx, query, nodes, nil,
			func(n *GrantedPermission, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := gpq.withGroup; query != nil {
		if err := gpq.loadGroup(ctx, query, nodes, nil,
			func(n *GrantedPermission, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gpq *GrantedPermissionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*GrantedPermission, init func(*GrantedPermission), assign func(*GrantedPermission, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*GrantedPermission)
	for i := range nodes {
		if nodes[i].granted_permission_user == nil {
			continue
		}
		fk := *nodes[i].granted_permission_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "granted_permission_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gpq *GrantedPermissionQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GrantedPermission, init func(*GrantedPermission), assign func(*GrantedPermission, *Group)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*GrantedPermission)
	for i := range nodes {
		if nodes[i].granted_permission_group == nil {
			continue
		}
		fk := *nodes[i].granted_permission_group
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "granted_permission_group" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gpq *GrantedPermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gpq.querySpec()
	_spec.Node.Columns = gpq.ctx.Fields
	if len(gpq.ctx.Fields) > 0 {
		_spec.Unique = gpq.ctx.Unique != nil && *gpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gpq.driver, _spec)
}

func (gpq *GrantedPermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(grantedpermission.Table, grantedpermission.Columns, sqlgraph.NewFieldSpec(grantedpermission.FieldID, field.TypeUUID))
	_spec.From = gpq.sql
	if unique := gpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gpq.path != nil {
		_spec.Unique = true
	}
	if fields := gpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grantedpermission.FieldID)
		for i := range fields {
			if fields[i] != grantedpermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gpq *GrantedPermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gpq.driver.Dialect())
	t1 := builder.Table(grantedpermission.Table)
	columns := gpq.ctx.Fields
	if len(columns) == 0 {
		columns = grantedpermission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gpq.sql != nil {
		selector = gpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gpq.ctx.Unique != nil && *gpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gpq.predicates {
		p(selector)
	}
	for _, p := range gpq.order {
		p(selector)
	}
	if offset := gpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GrantedPermissionGroupBy is the group-by builder for GrantedPermission entities.
type GrantedPermissionGroupBy struct {
	selector
	build *GrantedPermissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gpgb *GrantedPermissionGroupBy) Aggregate(fns ...AggregateFunc) *GrantedPermissionGroupBy {
	gpgb.fns = append(gpgb.fns, fns...)
	return gpgb
}

// Scan applies the selector query and scans the result into the given value.
func (gpgb *GrantedPermissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gpgb.build.ctx, "GroupBy")
	if err := gpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GrantedPermissionQuery, *GrantedPermissionGroupBy](ctx, gpgb.build, gpgb, gpgb.build.inters, v)
}

func (gpgb *GrantedPermissionGroupBy) sqlScan(ctx context.Context, root *GrantedPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gpgb.fns))
	for _, fn := range gpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gpgb.flds)+len(gpgb.fns))
		for _, f := range *gpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GrantedPermissionSelect is the builder for selecting fields of GrantedPermission entities.
type GrantedPermissionSelect struct {
	*GrantedPermissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gps *GrantedPermissionSelect) Aggregate(fns ...AggregateFunc) *GrantedPermissionSelect {
	gps.fns = append(gps.fns, fns...)
	return gps
}

// Scan applies the selector query and scans the result into the given value.
func (gps *GrantedPermissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gps.ctx, "Select")
	if err := gps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GrantedPermissionQuery, *GrantedPermissionSelect](ctx, gps.GrantedPermissionQuery, gps, gps.inters, v)
}

func (gps *GrantedPermissionSelect) sqlScan(ctx context.Context, root *GrantedPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gps.fns))
	for _, fn := range gps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
