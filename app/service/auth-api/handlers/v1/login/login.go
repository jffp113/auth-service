package login

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"go.uber.org/zap"
	"net/http"
)

type Config struct {
	log *zap.SugaredLogger
	db  data.Client
}

func Register(mux http.ServeMux, cfg Config) {
	//mux.Handle() login
	//mux.Handle() refresh
}
