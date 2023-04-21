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
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("fullName"),
		field.String("username").NotEmpty(),
		field.String("email"),
		field.String("hash"),
		field.String("preferences"), //TODO should be JSON
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).
			Through("auth_users_roles", UsersRoles.Type),
		edge.From("claims", Claim.Type).
			Ref("user"),
		edge.To("groups", Group.Type).
			Through("auth_users_groups", UsersGroups.Type),
	}
}

func (User) Indexed() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("username").Unique(),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "auth_users"},
	}
}
