package users

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"context"
	"fmt"
	"time"
)

type Storer interface {
	Create(ctx context.Context, prd NewUser) (User, error)
	//Update(ctx context.Context, prd Product) error
	//Delete(ctx context.Context, prd Product) error
	//Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Product, error)
	//Count(ctx context.Context, filter QueryFilter) (int, error)
	//QueryByID(ctx context.Context, productID uuid.UUID) (Product, error)
	//QueryByUserID(ctx context.Context, userID uuid.UUID) ([]Product, error)
}

type Store struct {
	cli data.Client
	//TODO add logger
}

func New(cli data.Client) Store {
	return Store{
		cli: cli,
	}
}

func (s *Store) Create(ctx context.Context, nu NewUser) (User, error) {
	now := time.Now()
	u := User{
		Id:          0,
		FullName:    nu.FullName,
		Username:    nu.Username,
		Hash:        nu.Hash,
		Email:       nu.Email,
		Preferences: nu.Preferences,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	resp := s.cli.Create(&u)

	if resp.Error != nil {
		return User{}, fmt.Errorf("adding user to database: %w", resp.Error)
	}

	return u, nil
}

//func GetById(ctx context.Context, cli data.Client, id int) (*ent.User, error) {
//	u, err := cli.User.Get(ctx, id)
//
//	if err != nil {
//		return nil, fmt.Errorf("querying user by id=%v: %w", id, err)
//	}
//
//	return u, nil
//}

//func GetAll(ctx context.Context, cli data.Client) ([]*ent.User, error) {
//	us, err := cli.User.Query().All(ctx)
//
//	if err != nil {
//		return nil, fmt.Errorf("quering all users: %w", err)
//	}
//
//	return us, nil
//}

//func Update(ctx context.Context, cli data.Client, id int, nu NewUser) (*ent.User, error) {
//	updateU := cli.User.UpdateOneID(id).
//		SetUpdatedAt(time.Now())
//
//	if nu.FullName != "" {
//		updateU.SetFullName(nu.FullName)
//	}
//	if nu.Username != "" {
//		updateU.SetUsername(nu.Username)
//	}
//	if nu.Hash != "" {
//		updateU.SetHash(nu.Hash)
//	}
//	if nu.Email != "" {
//		updateU.SetEmail(nu.Email)
//	}
//	if nu.Preferences != "" {
//		updateU.SetPreferences(nu.Preferences)
//	}
//
//	u, err := updateU.Save(ctx)
//
//	if err != nil {
//		return nil, fmt.Errorf("updating user to database: %w", err)
//	}
//
//	return u, nil
//}

//func FetchAll(ctx context.Context, cli data.Client) (*ent.User, error) {
//	cli.User.
//}
