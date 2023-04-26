package users

import (
	"com.cross-join.crossviewer.authservice/app/service/auth-api/graph/model"
	"com.cross-join.crossviewer.authservice/business/core/mappers"
	"com.cross-join.crossviewer.authservice/business/data/schema"
	"com.cross-join.crossviewer.authservice/business/data/users"
	"com.cross-join.crossviewer.authservice/foundation/hasher"
	"context"
	"fmt"
	"go.uber.org/zap"
)

//TODO, should we create a new model for the core, like the DB?

type Core struct {
	store users.Storer
	log   *zap.SugaredLogger
}

func NewCore(log *zap.SugaredLogger, store users.Storer) Core {
	return Core{
		store: store,
		log:   log,
	}
}

func (c *Core) CreateUser(ctx context.Context, nu model.UserInput) (*model.User, error) {

	hash, err := hasher.CreateHash(*nu.Password)

	if err != nil {
		return nil, fmt.Errorf("creating password hash: %w", err)
	}

	dbnu := schema.NewUser{
		FullName:    *nu.FullName,
		Username:    *nu.Username,
		Email:       *nu.Email,
		Preferences: *nu.Preferences,
		Hash:        hash,
	}

	u, err := c.store.Create(ctx, dbnu)

	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	return &model.User{
		ID:          u.Id,
		FullName:    u.FullName,
		Username:    u.Username,
		Email:       u.Email,
		Preferences: &u.Preferences,
	}, nil
}

func (c *Core) AllUsers(ctx context.Context) ([]*model.User, error) {

	us, err := c.store.Query(ctx)

	if err != nil {
		return nil, fmt.Errorf("querying all users: %w", err)
	}

	return mappers.MapUsers(us), nil
}

func (c *Core) UsersById(ctx context.Context, userId int) (*model.User, error) {
	us, err := c.store.QueryById(ctx, userId)

	if err != nil {
		return nil, fmt.Errorf("querying users by id: %w", err)
	}

	return mappers.MapUser(us), nil
}

func (c *Core) UserRoles(ctx context.Context, userId int) ([]*model.Role, error) {
	rs, err := c.store.QueryUserRoles(ctx, userId)

	if err != nil {
		return nil, fmt.Errorf("querying users roles: %w", err)
	}

	return mappers.MapRoles(rs), nil
}
