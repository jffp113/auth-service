package main

import (
	"com.cross-join.crossviewer.authservice/app/service/auth-api/graph"
	"com.cross-join.crossviewer.authservice/app/service/auth-api/handlers/debug"
	usersCore "com.cross-join.crossviewer.authservice/business/core/users"
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/users"
	"com.cross-join.crossviewer.authservice/foundation/logger"
	"context"
	"errors"
	"expvar"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	conf "github.com/ardanlabs/conf/v3"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const build = "develop"

func main() {
	log, err := logger.New("AUTH-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Fatal(err)
	}
}

func run(log *zap.SugaredLogger) error {
	// =========================================================================
	// GOMAXPROCS
	opt := maxprocs.Logger(log.Infof)
	if _, err := maxprocs.Set(opt); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}

	// =========================================================================
	// Configuration
	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:3000"`
			DebugHost       string        `conf:"default:0.0.0.0:4000"`
		}
		Auth struct {
			KeysFolder string `conf:"default:zarf/keys/"`
			ActiveKID  string `conf:"default:54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"`
		}
		DB struct {
			User         string `conf:"default:xviewer"`
			Password     string `conf:"default:xviewer,mask"`
			Host         string `conf:"default:localhost"`
			Port         int    `conf:"default:4438"`
			Name         string `conf:"default:xviewer_meta"`
			MaxIdleConns int    `conf:"default:2"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = "AUTH"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}
	// =========================================================================
	// App Starting

	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown completed")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating configs for output: %w", err)
	}

	log.Infow("startup", "config", out)

	expvar.NewString("build").Set(build)

	// =========================================================================
	// Database Support
	log.Infow("startup", "status", "initializing database support", "host", cfg.DB.Host)

	ctx := context.Background()
	dbCli, err := data.New(ctx,
		data.WithCredentials(cfg.DB.User, cfg.DB.Password),
		data.WithHostAndPort(cfg.DB.Host, cfg.DB.Port),
		data.WithDbName(cfg.DB.Name),
	)

	if err != nil {
		return fmt.Errorf("initializing BD: %w", err)
	}

	//////////////////////////User core init////////////////////////////
	usersStore := users.New(dbCli)
	users := usersCore.NewCore(log, &usersStore)

	resolver := graph.Resolver{
		Users: users,
		Log:   log,
	}

	// =========================================================================
	// Start Debug Service

	log.Infow("startup", "status", "initializing debug endpoints", "host", cfg.Web.DebugHost)
	go func() {
		debugMux := debug.Mux(log)
		http.ListenAndServe(cfg.Web.DebugHost, debugMux)
	}()

	// =========================================================================
	// Start API Service
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver,
		Directives: struct {
			Authenticated func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
			HasRole       func(ctx context.Context, obj interface{}, next graphql.Resolver, roles *string) (res interface{}, err error)
		}{Authenticated: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			return next(ctx)
		}},
	}))

	var mux http.ServeMux

	mux.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      &mux,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Infow("startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case sig := <-shutdown:
		log.Infow("startup", "status", "shutdown started", "signal", sig)
		defer log.Infow("startup", "status", "shutdown completed", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully %w", err)
		}
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
