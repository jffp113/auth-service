package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// User holds the schema definition for the User entity.
type UsersRoles struct {
	ent.Schema
}

// Fields of the User.
func (UsersRoles) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("role_id"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (UsersRoles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("roles", Role.Type).
			Required().
			Unique().
			Field("role_id"),
	}
}

func (UsersRoles) Indexed() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").
			Unique(),
		index.Fields("role_id").
			Unique(),
	}
}

// Annotations of the User.
func (UsersRoles) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "auth_users_roles"},
	}
}
