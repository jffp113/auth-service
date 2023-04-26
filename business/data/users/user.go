package users

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/schema"
	"context"
	"fmt"
	"time"
)

type Storer interface {
	Create(ctx context.Context, prd schema.NewUser) (schema.User, error)
	Query(ctx context.Context) ([]schema.User, error)
	QueryById(ctx context.Context, userId int) (schema.User, error)
	QueryUserRoles(ctx context.Context, userId int) ([]schema.Role, error)
	//Update(ctx context.Context, prd Product) error
	//Delete(ctx context.Context, prd Product) error
	//Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Product, error)
	//Count(ctx context.Context, filter QueryFilter) (int, error)
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

func (s *Store) Create(ctx context.Context, nu schema.NewUser) (schema.User, error) {
	now := time.Now()
	u := schema.User{
		Id:          0,
		FullName:    nu.FullName,
		Username:    nu.Username,
		Hash:        nu.Hash,
		Email:       nu.Email,
		Preferences: nu.Preferences,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	resp := s.cli.WithContext(ctx).
		Create(&u)

	if resp.Error != nil {
		return schema.User{}, fmt.Errorf("adding user to database: %w", resp.Error)
	}

	return u, nil
}

func (s *Store) QueryById(ctx context.Context, id int) (schema.User, error) {
	var u schema.User

	resp := s.cli.WithContext(ctx).
		First(&u, id)

	if resp.Error != nil {
		return schema.User{}, fmt.Errorf("querying user by id=%v: %w", id, resp.Error)
	}

	return u, nil
}

func (s *Store) Query(ctx context.Context) ([]schema.User, error) {
	var us []schema.User

	result := s.cli.WithContext(ctx).
		Find(&us)

	if result.Error != nil {
		return nil, fmt.Errorf("querying users: %w", result.Error)
	}

	return us, nil
}

func (s *Store) QueryUserRoles(ctx context.Context, userId int) ([]schema.Role, error) {
	var u schema.User

	result := s.cli.WithContext(ctx).
		Preload("Roles").
		First(&u, userId)

	if result.Error != nil {
		return nil, fmt.Errorf("querying user(%v) roles: %w", userId, result.Error)
	}

	return u.Roles, nil
}

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
