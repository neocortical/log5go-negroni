package log5gonegroni

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

type middleware struct {
	middlewareFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

func (mw *middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	mw.middlewareFunc(rw, r, next)
}

// NewMiddleware returns a new Negroni middleware based on the supplied closure
func NewMiddleware(middlewareFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) negroni.Handler {
	return &middleware{
		middlewareFunc: middlewareFunc,
	}
}
