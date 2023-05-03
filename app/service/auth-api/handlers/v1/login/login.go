package login

import (
	"com.cross-join.crossviewer.authservice/business/data"
	"com.cross-join.crossviewer.authservice/business/data/users"
	"com.cross-join.crossviewer.authservice/foundation/context"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Config struct {
	Log    *zap.SugaredLogger
	Db     data.Client
	Tracer trace.Tracer
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HandlerWithError func(http.ResponseWriter, *http.Request) error
type Handler func(http.ResponseWriter, *http.Request)

func Login(c Config) Handler {
	return AppHandler(c, func(w http.ResponseWriter, r *http.Request) error {
		bs, err := io.ReadAll(r.Body)

		if err != nil {
			return err
		}

		var user User

		err = json.Unmarshal(bs, &user)

		if err != nil {
			return err
		}

		db := users.New(c.Db)

		u, err := db.QueryByUsername(r.Context(), user.Username)

		if err != nil {
			return err
		}

		if u.Hash != user.Password {
			return errors.New("wrong password")
		}

		return nil
	})
}

func AppHandler(c Config, handler HandlerWithError) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := c.Tracer.Start(r.Context(), "http handler")
		traceId := span.SpanContext().TraceID().String()
		ctx = context.Wrap(ctx, traceId, c.Tracer)
		r = r.WithContext(ctx)

		defer span.End()

		err := handler(w, r)
		if err != nil {
			c.Log.Errorw("executing handler", "error", err, "trace_id", traceId)
			w.WriteHeader(http.StatusInternalServerError)
			span.SetStatus(codes.Error, fmt.Sprintf("%s", err))
		} else {
			c.Log.Infow("executing handler", "trace_id", traceId)
		}
	}
}

func Register(mux *http.ServeMux, cfg Config) {
	mux.HandleFunc("/login", Login(cfg))
	//mux.Haxndle() refresh
}
