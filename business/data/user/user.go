package user

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/ent"
	"context"
	"fmt"
	"time"
)

func Add(ctx context.Context, cli data.Client, nu NewUser) (*ent.User, error) {
	now := time.Now()

	u, err := cli.User.Create().
		SetFullName(nu.FullName).
		SetUsername(nu.Username).
		SetEmail(nu.Email).
		SetPreferences(nu.Preferences).
		SetHash(nu.Hash).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("adding user to database: %w", err)
	}

	return u, nil
}

func GetById(ctx context.Context, cli data.Client, id int) (*ent.User, error) {
	u, err := cli.User.Get(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("querying user by id=%v: %w", id, err)
	}

	return u, nil
}

func GetAll(ctx context.Context, cli data.Client) ([]*ent.User, error) {
	us, err := cli.User.Query().All(ctx)

	if err != nil {
		return nil, fmt.Errorf("quering all users: %w", err)
	}

	return us, nil
}

func Update(ctx context.Context, cli data.Client, id int, nu NewUser) (*ent.User, error) {
	updateU := cli.User.UpdateOneID(id).
		SetUpdatedAt(time.Now())

	if nu.FullName != "" {
		updateU.SetFullName(nu.FullName)
	}
	if nu.Username != "" {
		updateU.SetUsername(nu.Username)
	}
	if nu.Hash != "" {
		updateU.SetHash(nu.Hash)
	}
	if nu.Email != "" {
		updateU.SetEmail(nu.Email)
	}
	if nu.Preferences != "" {
		updateU.SetPreferences(nu.Preferences)
	}

	u, err := updateU.Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("updating user to database: %w", err)
	}

	return u, nil
}

//func FetchAll(ctx context.Context, cli data.Client) (*ent.User, error) {
//	cli.User.
//}
