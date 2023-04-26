package main

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
)

//userids 1,2,3,4,5,6
//5 -> [role. role, role]
//6 -> [role, role, role]

func main() {

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		for i, k := range keys {
			fmt.Println(k)
			results = append(results, &dataloader.Result{
				Data: []int{i},
			})
		}
		return results
	}

	loaderRolesIds := dataloader.NewBatchedLoader(batchFn)

	rolesFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		fmt.Println(keys)
		for _, k := range keys {
			results = append(results, &dataloader.Result{
				Data: "Role" + k.String(),
			})
		}
		return results
	}

	rolesLoader := dataloader.NewBatchedLoader(rolesFn)

	ctx := context.Background()
	thunk := loaderRolesIds.Load(ctx, dataloader.StringKey("key1"))
	result, err := thunk()
	AddToLoader(ctx, rolesLoader, result)

	thunk = loaderRolesIds.Load(ctx, dataloader.StringKey("key2"))
	result, err = thunk()
	AddToLoader(ctx, rolesLoader, result)

	thunk = loaderRolesIds.Load(ctx, dataloader.StringKey("key3"))
	result, err = thunk()
	fmt.Println("here")
	fmt.Println(result)
	thunkMany := AddToLoader(ctx, rolesLoader, result)
	r2, _ := thunkMany()

	thunk = loaderRolesIds.Load(ctx, dataloader.StringKey("key4"))
	result, err = thunk()
	thunkMany = AddToLoader(ctx, rolesLoader, result)

	r, _ := thunkMany()
	fmt.Println("result:")
	fmt.Println(result)
	fmt.Println(r2)
	fmt.Println(r)

	if err != nil {

	}
}

func AddToLoader(ctx context.Context, dl *dataloader.Loader, keys any) dataloader.ThunkMany {
	ks, ok := keys.([]int)

	if !ok {
		panic("error")
	}

	var result dataloader.Keys
	for k := range ks {
		result = append(result, dataloader.StringKey(fmt.Sprintf("%v", k)))
	}
	return dl.LoadMany(ctx, result)
}
