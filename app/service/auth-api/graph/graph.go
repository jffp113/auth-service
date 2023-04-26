package graph

import (
	usersCore "com.cross-join.crossviewer.authservice/business/core/users"
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/users"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.uber.org/zap"
	"net/http"
)

type MuxConfig struct {
	DbCli data.Client
	//UsersCore users.Core
	Log *zap.SugaredLogger
}

func Mux(config *MuxConfig) http.Handler {
	mux := http.NewServeMux()

	srv := buildGraphqlServer(config)

	mux.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	return srv
}

func buildGraphqlServer(config *MuxConfig) *handler.Server {
	resolver := Resolver{
		UsersCore: NewUserCore(config.DbCli, config.Log),
		Log:       config.Log,
	}

	srv := handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: &resolver,
		Directives: DirectiveRoot{
			Authenticated: Authenticated,
			HasRole:       HashRole,
		},
	}))

	return srv
}

func NewUserCore(db data.Client, log *zap.SugaredLogger) usersCore.Core {
	usersStore := users.New(db)
	return usersCore.NewCore(log, &usersStore)
}

// //Directives
func Authenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return next(ctx)
}

func HashRole(ctx context.Context, obj interface{}, next graphql.Resolver, roles *string) (res interface{}, err error) {
	return next(ctx)
}
