// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"com.cross-join.crossviewer.authService/business/data/ent/predicate"
	"com.cross-join.crossviewer.authService/business/data/ent/role"
	"com.cross-join.crossviewer.authService/business/data/ent/user"
	"com.cross-join.crossviewer.authService/business/data/ent/usersroles"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersRolesUpdate is the builder for updating UsersRoles entities.
type UsersRolesUpdate struct {
	config
	hooks    []Hook
	mutation *UsersRolesMutation
}

// Where appends a list predicates to the UsersRolesUpdate builder.
func (uru *UsersRolesUpdate) Where(ps ...predicate.UsersRoles) *UsersRolesUpdate {
	uru.mutation.Where(ps...)
	return uru
}

// SetUserID sets the "user_id" field.
func (uru *UsersRolesUpdate) SetUserID(i int) *UsersRolesUpdate {
	uru.mutation.SetUserID(i)
	return uru
}

// SetRoleID sets the "role_id" field.
func (uru *UsersRolesUpdate) SetRoleID(i int) *UsersRolesUpdate {
	uru.mutation.SetRoleID(i)
	return uru
}

// SetCreatedAt sets the "created_at" field.
func (uru *UsersRolesUpdate) SetCreatedAt(t time.Time) *UsersRolesUpdate {
	uru.mutation.SetCreatedAt(t)
	return uru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uru *UsersRolesUpdate) SetNillableCreatedAt(t *time.Time) *UsersRolesUpdate {
	if t != nil {
		uru.SetCreatedAt(*t)
	}
	return uru
}

// SetUpdatedAt sets the "updated_at" field.
func (uru *UsersRolesUpdate) SetUpdatedAt(t time.Time) *UsersRolesUpdate {
	uru.mutation.SetUpdatedAt(t)
	return uru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uru *UsersRolesUpdate) SetNillableUpdatedAt(t *time.Time) *UsersRolesUpdate {
	if t != nil {
		uru.SetUpdatedAt(*t)
	}
	return uru
}

// SetUser sets the "user" edge to the User entity.
func (uru *UsersRolesUpdate) SetUser(u *User) *UsersRolesUpdate {
	return uru.SetUserID(u.ID)
}

// SetRolesID sets the "roles" edge to the Role entity by ID.
func (uru *UsersRolesUpdate) SetRolesID(id int) *UsersRolesUpdate {
	uru.mutation.SetRolesID(id)
	return uru
}

// SetRoles sets the "roles" edge to the Role entity.
func (uru *UsersRolesUpdate) SetRoles(r *Role) *UsersRolesUpdate {
	return uru.SetRolesID(r.ID)
}

// Mutation returns the UsersRolesMutation object of the builder.
func (uru *UsersRolesUpdate) Mutation() *UsersRolesMutation {
	return uru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uru *UsersRolesUpdate) ClearUser() *UsersRolesUpdate {
	uru.mutation.ClearUser()
	return uru
}

// ClearRoles clears the "roles" edge to the Role entity.
func (uru *UsersRolesUpdate) ClearRoles() *UsersRolesUpdate {
	uru.mutation.ClearRoles()
	return uru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UsersRolesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UsersRolesMutation](ctx, uru.sqlSave, uru.mutation, uru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UsersRolesUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UsersRolesUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UsersRolesUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uru *UsersRolesUpdate) check() error {
	if _, ok := uru.mutation.UserID(); uru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersRoles.user"`)
	}
	if _, ok := uru.mutation.RolesID(); uru.mutation.RolesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersRoles.roles"`)
	}
	return nil
}

func (uru *UsersRolesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersroles.Table, usersroles.Columns, sqlgraph.NewFieldSpec(usersroles.FieldID, field.TypeInt))
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uru.mutation.CreatedAt(); ok {
		_spec.SetField(usersroles.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uru.mutation.UpdatedAt(); ok {
		_spec.SetField(usersroles.FieldUpdatedAt, field.TypeTime, value)
	}
	if uru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.UserTable,
			Columns: []string{usersroles.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.UserTable,
			Columns: []string{usersroles.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uru.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.RolesTable,
			Columns: []string{usersroles.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.RolesTable,
			Columns: []string{usersroles.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersroles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uru.mutation.done = true
	return n, nil
}

// UsersRolesUpdateOne is the builder for updating a single UsersRoles entity.
type UsersRolesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersRolesMutation
}

// SetUserID sets the "user_id" field.
func (uruo *UsersRolesUpdateOne) SetUserID(i int) *UsersRolesUpdateOne {
	uruo.mutation.SetUserID(i)
	return uruo
}

// SetRoleID sets the "role_id" field.
func (uruo *UsersRolesUpdateOne) SetRoleID(i int) *UsersRolesUpdateOne {
	uruo.mutation.SetRoleID(i)
	return uruo
}

// SetCreatedAt sets the "created_at" field.
func (uruo *UsersRolesUpdateOne) SetCreatedAt(t time.Time) *UsersRolesUpdateOne {
	uruo.mutation.SetCreatedAt(t)
	return uruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uruo *UsersRolesUpdateOne) SetNillableCreatedAt(t *time.Time) *UsersRolesUpdateOne {
	if t != nil {
		uruo.SetCreatedAt(*t)
	}
	return uruo
}

// SetUpdatedAt sets the "updated_at" field.
func (uruo *UsersRolesUpdateOne) SetUpdatedAt(t time.Time) *UsersRolesUpdateOne {
	uruo.mutation.SetUpdatedAt(t)
	return uruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uruo *UsersRolesUpdateOne) SetNillableUpdatedAt(t *time.Time) *UsersRolesUpdateOne {
	if t != nil {
		uruo.SetUpdatedAt(*t)
	}
	return uruo
}

// SetUser sets the "user" edge to the User entity.
func (uruo *UsersRolesUpdateOne) SetUser(u *User) *UsersRolesUpdateOne {
	return uruo.SetUserID(u.ID)
}

// SetRolesID sets the "roles" edge to the Role entity by ID.
func (uruo *UsersRolesUpdateOne) SetRolesID(id int) *UsersRolesUpdateOne {
	uruo.mutation.SetRolesID(id)
	return uruo
}

// SetRoles sets the "roles" edge to the Role entity.
func (uruo *UsersRolesUpdateOne) SetRoles(r *Role) *UsersRolesUpdateOne {
	return uruo.SetRolesID(r.ID)
}

// Mutation returns the UsersRolesMutation object of the builder.
func (uruo *UsersRolesUpdateOne) Mutation() *UsersRolesMutation {
	return uruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uruo *UsersRolesUpdateOne) ClearUser() *UsersRolesUpdateOne {
	uruo.mutation.ClearUser()
	return uruo
}

// ClearRoles clears the "roles" edge to the Role entity.
func (uruo *UsersRolesUpdateOne) ClearRoles() *UsersRolesUpdateOne {
	uruo.mutation.ClearRoles()
	return uruo
}

// Where appends a list predicates to the UsersRolesUpdate builder.
func (uruo *UsersRolesUpdateOne) Where(ps ...predicate.UsersRoles) *UsersRolesUpdateOne {
	uruo.mutation.Where(ps...)
	return uruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uruo *UsersRolesUpdateOne) Select(field string, fields ...string) *UsersRolesUpdateOne {
	uruo.fields = append([]string{field}, fields...)
	return uruo
}

// Save executes the query and returns the updated UsersRoles entity.
func (uruo *UsersRolesUpdateOne) Save(ctx context.Context) (*UsersRoles, error) {
	return withHooks[*UsersRoles, UsersRolesMutation](ctx, uruo.sqlSave, uruo.mutation, uruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UsersRolesUpdateOne) SaveX(ctx context.Context) *UsersRoles {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UsersRolesUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UsersRolesUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uruo *UsersRolesUpdateOne) check() error {
	if _, ok := uruo.mutation.UserID(); uruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersRoles.user"`)
	}
	if _, ok := uruo.mutation.RolesID(); uruo.mutation.RolesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UsersRoles.roles"`)
	}
	return nil
}

func (uruo *UsersRolesUpdateOne) sqlSave(ctx context.Context) (_node *UsersRoles, err error) {
	if err := uruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersroles.Table, usersroles.Columns, sqlgraph.NewFieldSpec(usersroles.FieldID, field.TypeInt))
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UsersRoles.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersroles.FieldID)
		for _, f := range fields {
			if !usersroles.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usersroles.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uruo.mutation.CreatedAt(); ok {
		_spec.SetField(usersroles.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uruo.mutation.UpdatedAt(); ok {
		_spec.SetField(usersroles.FieldUpdatedAt, field.TypeTime, value)
	}
	if uruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.UserTable,
			Columns: []string{usersroles.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.UserTable,
			Columns: []string{usersroles.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uruo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.RolesTable,
			Columns: []string{usersroles.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersroles.RolesTable,
			Columns: []string{usersroles.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UsersRoles{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersroles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uruo.mutation.done = true
	return _node, nil
}
