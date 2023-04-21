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

type UsersGroups struct {
	ent.Schema
}

// Fields of the User.
func (UsersGroups) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("group_id"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (UsersGroups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("roles", Group.Type).
			Required().
			Unique().
			Field("group_id"),
	}
}

func (UsersGroups) Indexed() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").
			Unique(),
		index.Fields("group_id").
			Unique(),
	}
}

// Annotations of the User.
func (UsersGroups) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "auth_users_groups"},
	}
}
