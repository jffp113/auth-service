// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"com.cross-join.crossviewer.authService/business/data/ent/predicate"
	"com.cross-join.crossviewer.authService/business/data/ent/usersgroups"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersGroupsDelete is the builder for deleting a UsersGroups entity.
type UsersGroupsDelete struct {
	config
	hooks    []Hook
	mutation *UsersGroupsMutation
}

// Where appends a list predicates to the UsersGroupsDelete builder.
func (ugd *UsersGroupsDelete) Where(ps ...predicate.UsersGroups) *UsersGroupsDelete {
	ugd.mutation.Where(ps...)
	return ugd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ugd *UsersGroupsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UsersGroupsMutation](ctx, ugd.sqlExec, ugd.mutation, ugd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ugd *UsersGroupsDelete) ExecX(ctx context.Context) int {
	n, err := ugd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ugd *UsersGroupsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usersgroups.Table, sqlgraph.NewFieldSpec(usersgroups.FieldID, field.TypeInt))
	if ps := ugd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ugd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ugd.mutation.done = true
	return affected, err
}

// UsersGroupsDeleteOne is the builder for deleting a single UsersGroups entity.
type UsersGroupsDeleteOne struct {
	ugd *UsersGroupsDelete
}

// Where appends a list predicates to the UsersGroupsDelete builder.
func (ugdo *UsersGroupsDeleteOne) Where(ps ...predicate.UsersGroups) *UsersGroupsDeleteOne {
	ugdo.ugd.mutation.Where(ps...)
	return ugdo
}

// Exec executes the deletion query.
func (ugdo *UsersGroupsDeleteOne) Exec(ctx context.Context) error {
	n, err := ugdo.ugd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usersgroups.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ugdo *UsersGroupsDeleteOne) ExecX(ctx context.Context) {
	if err := ugdo.Exec(ctx); err != nil {
		panic(err)
	}
}