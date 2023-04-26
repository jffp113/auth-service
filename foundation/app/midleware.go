package app

import "net/http"

type Middleware func(h http.Handler) http.Handler

func WrapMiddleware(core http.Handler, handlers ...Middleware) http.Handler {
	var last = core

	for i, _ := range handlers {
		h := handlers[len(handlers)-1-i]
		handler := h(last)
		last = handler
	}

	return last
}

//core, a, b, c

//a(b(c(core)))
