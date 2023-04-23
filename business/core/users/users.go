package users

import (
	"com.cross-join.crossviewer.authservice/app/service/auth-api/graph/model"
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

	dbnu := users.NewUser{
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
