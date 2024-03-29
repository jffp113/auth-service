package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"errors"
	"fmt"

	"com.cross-join.crossviewer.authservice/app/service/auth-api/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	u, err := r.UsersCore.CreateUser(ctx, input)

	if err != nil {
		r.Log.Errorw("Creating User", "error", err, "input", input)
		return nil, errors.New("could not create user")
	}

	return u, err
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// AddUserToRoles is the resolver for the addUserToRoles field.
func (r *mutationResolver) AddUserToRoles(ctx context.Context, userID int, input *model.UserRolesInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: AddUserToRoles - addUserToRoles"))
}

// RemoveUserFromRoles is the resolver for the removeUserFromRoles field.
func (r *mutationResolver) RemoveUserFromRoles(ctx context.Context, userID int, input *model.UserRolesInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: RemoveUserFromRoles - removeUserFromRoles"))
}

// AddUserToGroups is the resolver for the addUserToGroups field.
func (r *mutationResolver) AddUserToGroups(ctx context.Context, userID int, input *model.UserGroupsInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: AddUserToGroups - addUserToGroups"))
}

// RemoveUserFromGroups is the resolver for the removeUserFromGroups field.
func (r *mutationResolver) RemoveUserFromGroups(ctx context.Context, userID int, input *model.UserGroupsInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: RemoveUserFromGroups - removeUserFromGroups"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	us, err := r.UsersCore.AllUsers(ctx)

	if err != nil {
		r.Log.Errorw("graphql: users", "error", err)
		return nil, errors.New("could not fetch users")
	}

	return us, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	us, err := r.UsersCore.UsersById(ctx, id)

	if err != nil {
		r.Log.Errorw("graphql: user", "error", err)
		return nil, errors.New("could not fetch user")
	}

	return us, nil
}

// Roles is the resolver for the roles field.
func (r *userResolver) Roles(ctx context.Context, obj *model.User) ([]*model.Role, error) {
	rs, err := r.UsersCore.UserRoles(ctx, obj.ID)

	if err != nil {
		r.Log.Errorw("graphql: user: roles", "error", err)
		return nil, errors.New("could not fetch user roles")
	}

	return rs, nil
}

// Groups is the resolver for the groups field.
func (r *userResolver) Groups(ctx context.Context, obj *model.User) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// Claims is the resolver for the claims field.
func (r *userResolver) Claims(ctx context.Context, obj *model.User) ([]*model.Claim, error) {
	panic(fmt.Errorf("not implemented: Claims - claims"))
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
