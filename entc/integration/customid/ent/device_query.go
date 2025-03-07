// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/device"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/entc/integration/customid/ent/session"
	"entgo.io/ent/schema/field"
)

// DeviceQuery is the builder for querying Device entities.
type DeviceQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.Device
	withActiveSession *SessionQuery
	withSessions      *SessionQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceQuery builder.
func (dq *DeviceQuery) Where(ps ...predicate.Device) *DeviceQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DeviceQuery) Limit(limit int) *DeviceQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DeviceQuery) Offset(offset int) *DeviceQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeviceQuery) Unique(unique bool) *DeviceQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DeviceQuery) Order(o ...OrderFunc) *DeviceQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryActiveSession chains the current query on the "active_session" edge.
func (dq *DeviceQuery) QueryActiveSession() *SessionQuery {
	query := &SessionQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, selector),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, device.ActiveSessionTable, device.ActiveSessionColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySessions chains the current query on the "sessions" edge.
func (dq *DeviceQuery) QuerySessions() *SessionQuery {
	query := &SessionQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(device.Table, device.FieldID, selector),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, device.SessionsTable, device.SessionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Device entity from the query.
// Returns a *NotFoundError when no Device was found.
func (dq *DeviceQuery) First(ctx context.Context) (*Device, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{device.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeviceQuery) FirstX(ctx context.Context) *Device {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Device ID from the query.
// Returns a *NotFoundError when no Device ID was found.
func (dq *DeviceQuery) FirstID(ctx context.Context) (id schema.ID, err error) {
	var ids []schema.ID
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{device.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeviceQuery) FirstIDX(ctx context.Context) schema.ID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Device entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Device entity is found.
// Returns a *NotFoundError when no Device entities are found.
func (dq *DeviceQuery) Only(ctx context.Context) (*Device, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{device.Label}
	default:
		return nil, &NotSingularError{device.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeviceQuery) OnlyX(ctx context.Context) *Device {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Device ID in the query.
// Returns a *NotSingularError when more than one Device ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeviceQuery) OnlyID(ctx context.Context) (id schema.ID, err error) {
	var ids []schema.ID
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{device.Label}
	default:
		err = &NotSingularError{device.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeviceQuery) OnlyIDX(ctx context.Context) schema.ID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Devices.
func (dq *DeviceQuery) All(ctx context.Context) ([]*Device, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeviceQuery) AllX(ctx context.Context) []*Device {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Device IDs.
func (dq *DeviceQuery) IDs(ctx context.Context) ([]schema.ID, error) {
	var ids []schema.ID
	if err := dq.Select(device.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeviceQuery) IDsX(ctx context.Context) []schema.ID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeviceQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeviceQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeviceQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeviceQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeviceQuery) Clone() *DeviceQuery {
	if dq == nil {
		return nil
	}
	return &DeviceQuery{
		config:            dq.config,
		limit:             dq.limit,
		offset:            dq.offset,
		order:             append([]OrderFunc{}, dq.order...),
		predicates:        append([]predicate.Device{}, dq.predicates...),
		withActiveSession: dq.withActiveSession.Clone(),
		withSessions:      dq.withSessions.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithActiveSession tells the query-builder to eager-load the nodes that are connected to
// the "active_session" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeviceQuery) WithActiveSession(opts ...func(*SessionQuery)) *DeviceQuery {
	query := &SessionQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withActiveSession = query
	return dq
}

// WithSessions tells the query-builder to eager-load the nodes that are connected to
// the "sessions" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeviceQuery) WithSessions(opts ...func(*SessionQuery)) *DeviceQuery {
	query := &SessionQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withSessions = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (dq *DeviceQuery) GroupBy(field string, fields ...string) *DeviceGroupBy {
	grbuild := &DeviceGroupBy{config: dq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = device.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (dq *DeviceQuery) Select(fields ...string) *DeviceSelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DeviceSelect{DeviceQuery: dq}
	selbuild.label = device.Label
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

func (dq *DeviceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !device.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeviceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Device, error) {
	var (
		nodes       = []*Device{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [2]bool{
			dq.withActiveSession != nil,
			dq.withSessions != nil,
		}
	)
	if dq.withActiveSession != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, device.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Device).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Device{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withActiveSession; query != nil {
		if err := dq.loadActiveSession(ctx, query, nodes, nil,
			func(n *Device, e *Session) { n.Edges.ActiveSession = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withSessions; query != nil {
		if err := dq.loadSessions(ctx, query, nodes,
			func(n *Device) { n.Edges.Sessions = []*Session{} },
			func(n *Device, e *Session) { n.Edges.Sessions = append(n.Edges.Sessions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DeviceQuery) loadActiveSession(ctx context.Context, query *SessionQuery, nodes []*Device, init func(*Device), assign func(*Device, *Session)) error {
	ids := make([]schema.ID, 0, len(nodes))
	nodeids := make(map[schema.ID][]*Device)
	for i := range nodes {
		if nodes[i].device_active_session == nil {
			continue
		}
		fk := *nodes[i].device_active_session
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(session.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_active_session" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DeviceQuery) loadSessions(ctx context.Context, query *SessionQuery, nodes []*Device, init func(*Device), assign func(*Device, *Session)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[schema.ID]*Device)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Session(func(s *sql.Selector) {
		s.Where(sql.InValues(device.SessionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.device_sessions
		if fk == nil {
			return fmt.Errorf(`foreign-key "device_sessions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_sessions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DeviceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeviceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DeviceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeBytes,
				Column: device.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for i := range fields {
			if fields[i] != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeviceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(device.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = device.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceGroupBy is the group-by builder for Device entities.
type DeviceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeviceGroupBy) Aggregate(fns ...AggregateFunc) *DeviceGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DeviceGroupBy) Scan(ctx context.Context, v any) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DeviceGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range dgb.fields {
		if !device.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DeviceGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DeviceSelect is the builder for selecting fields of Device entities.
type DeviceSelect struct {
	*DeviceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeviceSelect) Scan(ctx context.Context, v any) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DeviceQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DeviceSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
