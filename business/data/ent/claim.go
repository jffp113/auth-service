// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"com.cross-join.crossviewer.authService/business/data/ent/claim"
	"com.cross-join.crossviewer.authService/business/data/ent/user"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Claim is the model entity for the Claim schema.
type Claim struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Service holds the value of the "service" field.
	Service string `json:"service,omitempty"`
	// Claims holds the value of the "claims" field.
	Claims string `json:"claims,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClaimQuery when eager-loading is set.
	Edges        ClaimEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClaimEdges holds the relations/edges for other nodes in the graph.
type ClaimEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClaimEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Claim) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case claim.FieldID, claim.FieldUserID:
			values[i] = new(sql.NullInt64)
		case claim.FieldService, claim.FieldClaims:
			values[i] = new(sql.NullString)
		case claim.FieldCreatedAt, claim.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Claim fields.
func (c *Claim) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case claim.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case claim.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.UserID = int(value.Int64)
			}
		case claim.FieldService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service", values[i])
			} else if value.Valid {
				c.Service = value.String
			}
		case claim.FieldClaims:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field claims", values[i])
			} else if value.Valid {
				c.Claims = value.String
			}
		case claim.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case claim.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Claim.
// This includes values selected through modifiers, order, etc.
func (c *Claim) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Claim entity.
func (c *Claim) QueryUser() *UserQuery {
	return NewClaimClient(c.config).QueryUser(c)
}

// Update returns a builder for updating this Claim.
// Note that you need to call Claim.Unwrap() before calling this method if this Claim
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Claim) Update() *ClaimUpdateOne {
	return NewClaimClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Claim entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Claim) Unwrap() *Claim {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Claim is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Claim) String() string {
	var builder strings.Builder
	builder.WriteString("Claim(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteString(", ")
	builder.WriteString("service=")
	builder.WriteString(c.Service)
	builder.WriteString(", ")
	builder.WriteString("claims=")
	builder.WriteString(c.Claims)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Claims is a parsable slice of Claim.
type Claims []*Claim