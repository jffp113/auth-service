package main

import (
	"com.cross-join.crossviewer.authservice/app/service/auth-api/handlers/debug"
	v1 "com.cross-join.crossviewer.authservice/app/service/auth-api/handlers/v1"
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/foundation/logger"
	"context"
	"errors"
	"expvar"
	"fmt"
	conf "github.com/ardanlabs/conf/v3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
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
			User         string `conf:"default:postgres"`
			Password     string `conf:"default:postgres,mask"`
			Host         string `conf:"default:db"`
			Port         int    `conf:"default:5432"`
			Name         string `conf:"default:postgres"`
			MaxIdleConns int    `conf:"default:2"`    // TODO implement
			MaxOpenConns int    `conf:"default:0"`    // TODO implement
			DisableTLS   bool   `conf:"default:true"` // TODO implement
		}
		Otel struct {
			ReporterURI string  `conf:"default:https://otlp.nr-data.net:4317"`
			ServiceName string  `conf:"default:auth-api"`
			Probability float64 `conf:"default:1"`
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
		data.WithLogger(log),
	)

	if err != nil {
		return fmt.Errorf("initializing BD: %w", err)
	}

	// =========================================================================
	// Start Tracing Support

	log.Infow("startup", "status", "initializing OT/Otel tracing support")

	traceProvider, err := startTracing(
		cfg.Otel.ServiceName,
		cfg.Otel.ReporterURI,
		cfg.Otel.Probability,
		log.Desugar(),
	)
	if err != nil {
		return fmt.Errorf("starting tracing: %w", err)
	}
	defer traceProvider.Shutdown(context.Background())

	tracer := traceProvider.Tracer("service")

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

	//mux := graph.Mux(&graph.MuxConfig{
	//	DbCli: dbCli,
	//	Log:   log,
	//})
	//

	mux := v1.Mux(v1.Config{
		Log:    log,
		Db:     dbCli,
		Tracer: tracer,
	})

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      mux,
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

// startTracing configure open telemetry to be used with zipkin.
func startTracing(serviceName string, reporterURI string, probability float64, log *zap.Logger) (*trace.TracerProvider, error) {

	var headers = map[string]string{
		"api-key": "<your_license_key>",
	}

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithHeaders(headers),
		otlptracegrpc.WithEndpoint(reporterURI),
		otlptracegrpc.WithInsecure(),
	)

	exporter, err := otlptrace.New(context.Background(), client)

	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(probability)),
		trace.WithBatcher(exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
				attribute.String("exporter", "zipkin"),
			),
		),
	)

	// We must set this provider as the global provider for things to work,
	// but we pass this provider around the program where needed to collect
	// our traces.
	otel.SetTracerProvider(traceProvider)

	return traceProvider, nil
}
