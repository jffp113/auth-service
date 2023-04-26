package roles

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/schema"
	"context"
)

type RolesStorer interface {
	//Create(ctx context.Context, prd NewUser) (User, error)
	//Update(ctx context.Context, prd Product) error
	//Delete(ctx context.Context, prd Product) error
	//Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Product, error)
	//Count(ctx context.Context, filter QueryFilter) (int, error)
	//QueryByID(ctx context.Context, productID uuid.UUID) (Product, error)
	//QueryByUserID(ctx context.Context, userID uuid.UUID) ([]Product, error)

	QueryByUserId(ctx context.Context, userID int) ([]schema.Role, error)
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

//func (s *Store) QueryByUserId(ctx context.Context, userID int) ([]Role, error) {
//
//}
