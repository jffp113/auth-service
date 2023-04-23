package main

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/users"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	cli, err := data.New(ctx,
		data.WithDebug(true),
		data.WithCredentials("xviewer", "xviewer"),
		data.WithHostAndPort("localhost", 4438),
		data.WithDbName("xviewer_meta"),
	)

	if err != nil {
		panic(err)
	}

	storer := users.New(cli)

	u, err := storer.Create(ctx, users.NewUser{
		FullName:    "Jorge",
		Username:    "jff.pereira2",
		Email:       "jff.pereira@hotmail.com",
		Hash:        "qwerty",
		Preferences: "{}",
	})

	fmt.Println(u)

}
