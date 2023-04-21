package main

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/user"
	"context"
	"fmt"
	"log"
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
		log.Fatalln(err)
	}

	user, err := user.GetById(ctx, cli, 1)

	if err != nil {
		log.Fatalln(err)
	}

	roles, err := user.QueryRoles().All(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	claims, err := user.QueryClaims().All(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	groups, err := user.QueryGroups().All(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)
	fmt.Println(roles)
	fmt.Println(claims)
	fmt.Println(groups)

}
