package v1

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"go.uber.org/zap"
	"net/http"
)

type Config struct {
	log *zap.SugaredLogger
	db  data.Client
}

// Mux registers all the api  routes and then custom.
// This bypassing the use of the
// DefaultServerMux. Using the DefaultServerMux would be a security risk since
// a dependency could inject a handler into our service without us knowing it.
func Mux(cfg Config) http.Handler {
	mux := http.NewServeMux()

	return mux
}
