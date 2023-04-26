package dataloaders

import (
	"com.cross-join.crossviewer.authservice/app/service/auth-api/graph/model"
	"com.cross-join.crossviewer.authservice/business/core/users"
	"context"
	"errors"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"net/http"
	"reflect"
)

//https://github.com/zenyui/gqlgen-dataloader/blob/9e47f4af593b0dd4b771d2e5a61d03dd78c37445/graph/dataloader/dataloader.go

const dataloaderKey = "dataloader"

type Cores struct {
	users users.Core
}

type DataLoader struct {
	cores Cores

	userIdsToRolesIdsLoader *dataloader.Loader
	rolesLoader             *dataloader.Loader
}

func New(cores Cores) *DataLoader {
	return &DataLoader{}
}

func (dl *DataLoader) GetRoles(ctx context.Context, rolesIds []int) ([]*model.Role, error) {
	keys := toKeySlice(rolesIds)
	thuck := dl.rolesLoader.LoadMany(ctx, keys)

	rawResult, _ := thuck() //TODO, what to do with a slice of errors??

	result, err := convertSlice[any, *model.Role](rawResult)

	if err != nil {
		return nil, fmt.Errorf("while converting to roles: %w", err)
	}

	return result, nil
}

func (dl *DataLoader) GetRolesByUserId(ctx context.Context, userId int) ([]*model.Role, error) {
	thunk := dl.userIdsToRolesIdsLoader.Load(ctx, toKey(userId))
	result, err := thunk()

	if err != nil {
		return nil, fmt.Errorf("dataloading user ids to role ids: %w", err)
	}

	ids, err := anyToIntSlice(result)

	if err != nil {
		return nil, fmt.Errorf("converting dataloader ids result: %w", err)
	}

	return dl.GetRoles(ctx, ids)
}

func InjectDataloader(dl *DataLoader) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), dataloaderKey, dl)
			nr := r.WithContext(ctx)
			h.ServeHTTP(w, nr)
		})
	}
}

func For(ctx context.Context) *DataLoader {
	raw := ctx.Value(dataloaderKey)
	return raw.(*DataLoader)
}

func convertSlice[I any, R any](in []I) ([]R, error) {
	result := make([]R, 0, len(in))

	for i, _ := range in {
		r, ok := any(in[i]).(R)
		if !ok {
			return nil, fmt.Errorf("converting to type: %s != %s",
				reflect.TypeOf(in[i]), reflect.TypeOf(r))
		}
		result = append(result, r)
	}

	return result, nil
}

func toKeySlice(keys []int) dataloader.Keys {
	var result []dataloader.Key

	for _, k := range keys {
		result = append(result, toKey(k))
	}

	return result
}

func toKey(key int) dataloader.Key {
	strKey := fmt.Sprintf("%v", key)
	return dataloader.StringKey(strKey)
}

func anyToIntSlice(s any) ([]int, error) {
	is, ok := s.([]int)

	if !ok {
		return nil, errors.New("could not convert to int slice")
	}

	return is, nil
}
