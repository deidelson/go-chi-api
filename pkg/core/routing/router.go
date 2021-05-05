package routing

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

const (
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
	DELETE = "DELETE"
)

var (
	apiRouter *ApiRouter
)

type ApiRoute struct {
	Endpoint string
	Method   string
	//TODO ver que conviene
	Handler http.HandlerFunc
}

type ApiHandler interface {
	GetBasePath() string
	GetRoutes() []ApiRoute
	GetMiddlewares() []func(http.Handler) http.Handler
}

type ApiRouter struct {
	routerEngine *chi.Mux
}

func GetApiRouter() *ApiRouter {
	if apiRouter == nil {
		apiRouter = &ApiRouter{
			routerEngine: chi.NewRouter(),
		}
		apiRouter.routerEngine.Use(middleware.RequestID)
		apiRouter.routerEngine.Use(middleware.RealIP)
		apiRouter.routerEngine.Use(middleware.Logger)
		apiRouter.routerEngine.Use(middleware.Recoverer)
	}
	return apiRouter
}

func (this *ApiRouter) AddHandler(handler ApiHandler) {
	this.routerEngine.Route(handler.GetBasePath(), func(r chi.Router) {
		if handler.GetMiddlewares() != nil {
			r.Use(handler.GetMiddlewares()...)
		}

		for _, route := range handler.GetRoutes() {
			r.Method(route.Method, route.Endpoint, route.Handler)
		}
	})
}

func (this *ApiRouter) Start() {
	err := http.ListenAndServe(":8080", this.routerEngine)
	if err != nil {
		panic(err)
	}
}
