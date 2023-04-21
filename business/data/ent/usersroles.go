// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"com.cross-join.crossviewer.authService/business/data/ent/role"
	"com.cross-join.crossviewer.authService/business/data/ent/user"
	"com.cross-join.crossviewer.authService/business/data/ent/usersroles"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UsersRoles is the model entity for the UsersRoles schema.
type UsersRoles struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID int `json:"role_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UsersRolesQuery when eager-loading is set.
	Edges        UsersRolesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UsersRolesEdges holds the relations/edges for other nodes in the graph.
type UsersRolesEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Roles holds the value of the roles edge.
	Roles *Role `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UsersRolesEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UsersRolesEdges) RolesOrErr() (*Role, error) {
	if e.loadedTypes[1] {
		if e.Roles == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UsersRoles) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersroles.FieldID, usersroles.FieldUserID, usersroles.FieldRoleID:
			values[i] = new(sql.NullInt64)
		case usersroles.FieldCreatedAt, usersroles.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UsersRoles fields.
func (ur *UsersRoles) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersroles.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ur.ID = int(value.Int64)
		case usersroles.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ur.UserID = int(value.Int64)
			}
		case usersroles.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				ur.RoleID = int(value.Int64)
			}
		case usersroles.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ur.CreatedAt = value.Time
			}
		case usersroles.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ur.UpdatedAt = value.Time
			}
		default:
			ur.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UsersRoles.
// This includes values selected through modifiers, order, etc.
func (ur *UsersRoles) Value(name string) (ent.Value, error) {
	return ur.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UsersRoles entity.
func (ur *UsersRoles) QueryUser() *UserQuery {
	return NewUsersRolesClient(ur.config).QueryUser(ur)
}

// QueryRoles queries the "roles" edge of the UsersRoles entity.
func (ur *UsersRoles) QueryRoles() *RoleQuery {
	return NewUsersRolesClient(ur.config).QueryRoles(ur)
}

// Update returns a builder for updating this UsersRoles.
// Note that you need to call UsersRoles.Unwrap() before calling this method if this UsersRoles
// was returned from a transaction, and the transaction was committed or rolled back.
func (ur *UsersRoles) Update() *UsersRolesUpdateOne {
	return NewUsersRolesClient(ur.config).UpdateOne(ur)
}

// Unwrap unwraps the UsersRoles entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ur *UsersRoles) Unwrap() *UsersRoles {
	_tx, ok := ur.config.driver.(*txDriver)
	if !ok {
		panic("ent: UsersRoles is not a transactional entity")
	}
	ur.config.driver = _tx.drv
	return ur
}

// String implements the fmt.Stringer.
func (ur *UsersRoles) String() string {
	var builder strings.Builder
	builder.WriteString("UsersRoles(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ur.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ur.UserID))
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", ur.RoleID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ur.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ur.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UsersRolesSlice is a parsable slice of UsersRoles.
type UsersRolesSlice []*UsersRoles