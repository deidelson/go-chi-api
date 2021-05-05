package routing

import "net/http"

type ApiRoute struct {
	Endpoint string
	Method   string
	Handler http.HandlerFunc
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
