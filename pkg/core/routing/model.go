package routing

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type RouterEngine interface {
	Use(middlewares ...func(http.Handler) http.Handler)
	Method(method, pattern string, h http.Handler)
	Route(pattern string, fn func(r chi.Router)) chi.Router
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type ApiRoute struct {
	Endpoint string
	Method   string
	Handler  http.HandlerFunc
}

type Middleware func(http.Handler) http.Handler

type Middlewares []Middleware

func (this Middlewares) isNotEmpty() bool {
	return len(this) > 0
}

type ApiHandler interface {
	GetBasePath() string
	GetRoutes() []ApiRoute
	GetMiddlewares() Middlewares
}
