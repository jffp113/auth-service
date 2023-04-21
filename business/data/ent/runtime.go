// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"com.cross-join.crossviewer.authservice/business/data/ent/claim"
	"com.cross-join.crossviewer.authservice/business/data/ent/group"
	"com.cross-join.crossviewer.authservice/business/data/ent/role"
	"com.cross-join.crossviewer.authservice/business/data/ent/schema"
	"com.cross-join.crossviewer.authservice/business/data/ent/user"
	"com.cross-join.crossviewer.authservice/business/data/ent/usersgroups"
	"com.cross-join.crossviewer.authservice/business/data/ent/usersroles"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	claimFields := schema.Claim{}.Fields()
	_ = claimFields
	// claimDescCreatedAt is the schema descriptor for created_at field.
	claimDescCreatedAt := claimFields[4].Descriptor()
	// claim.DefaultCreatedAt holds the default value on creation for the created_at field.
	claim.DefaultCreatedAt = claimDescCreatedAt.Default.(func() time.Time)
	// claimDescUpdatedAt is the schema descriptor for updated_at field.
	claimDescUpdatedAt := claimFields[5].Descriptor()
	// claim.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	claim.DefaultUpdatedAt = claimDescUpdatedAt.Default.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[4].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupFields[5].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[3].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[4].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	usersgroupsFields := schema.UsersGroups{}.Fields()
	_ = usersgroupsFields
	// usersgroupsDescCreatedAt is the schema descriptor for created_at field.
	usersgroupsDescCreatedAt := usersgroupsFields[2].Descriptor()
	// usersgroups.DefaultCreatedAt holds the default value on creation for the created_at field.
	usersgroups.DefaultCreatedAt = usersgroupsDescCreatedAt.Default.(func() time.Time)
	// usersgroupsDescUpdatedAt is the schema descriptor for updated_at field.
	usersgroupsDescUpdatedAt := usersgroupsFields[3].Descriptor()
	// usersgroups.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usersgroups.DefaultUpdatedAt = usersgroupsDescUpdatedAt.Default.(func() time.Time)
	usersrolesFields := schema.UsersRoles{}.Fields()
	_ = usersrolesFields
	// usersrolesDescCreatedAt is the schema descriptor for created_at field.
	usersrolesDescCreatedAt := usersrolesFields[2].Descriptor()
	// usersroles.DefaultCreatedAt holds the default value on creation for the created_at field.
	usersroles.DefaultCreatedAt = usersrolesDescCreatedAt.Default.(func() time.Time)
	// usersrolesDescUpdatedAt is the schema descriptor for updated_at field.
	usersrolesDescUpdatedAt := usersrolesFields[3].Descriptor()
	// usersroles.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usersroles.DefaultUpdatedAt = usersrolesDescUpdatedAt.Default.(func() time.Time)
}
