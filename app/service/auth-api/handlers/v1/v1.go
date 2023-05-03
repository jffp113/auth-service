package v1

import (
	"com.cross-join.crossviewer.authservice/app/service/auth-api/handlers/v1/login"
	"com.cross-join.crossviewer.authservice/business/data"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
)

type Config struct {
	Log    *zap.SugaredLogger
	Db     data.Client
	Tracer trace.Tracer
}

// Mux registers all the api  routes and then custom.
// This bypassing the use of the
// DefaultServerMux. Using the DefaultServerMux would be a security risk since
// a dependency could inject a handler into our service without us knowing it.
func Mux(cfg Config) http.Handler {
	mux := http.NewServeMux()

	login.Register(mux, login.Config{
		Db:     cfg.Db,
		Log:    cfg.Log,
		Tracer: cfg.Tracer,
	})

	return mux
}
