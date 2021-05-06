package routing

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

var (
	apiRouterInstance *ApiRouter
)

type ApiRouter struct {
	routerEngine RouterEngine
}

func GetApiRouter() *ApiRouter {
	log.Println("Starting router")
	if apiRouterInstance == nil {
		apiRouterInstance = &ApiRouter{
			routerEngine: chi.NewRouter(),
		}
		log.Println("Adding recovery middleware")
		apiRouterInstance.routerEngine.Use(middleware.Recoverer)
	}
	return apiRouterInstance
}

func (this *ApiRouter) AddGlobalMiddleware(mWare Middleware) {
	this.routerEngine.Use(mWare)
}

func (this *ApiRouter) AddHandler(handler ApiHandler) {
	this.routerEngine.Route(handler.GetBasePath(), func(r chi.Router) {
		log.Println("Registering handler for route: ", handler.GetBasePath())

		if handler.GetMiddlewares() != nil && handler.GetMiddlewares().isNotEmpty() {
			for _, mw := range handler.GetMiddlewares() {
				r.Use(mw)
			}
		}

		for _, route := range handler.GetRoutes() {
			r.Method(route.Method, route.Endpoint, route.Handler)
		}
	})
}

func (this *ApiRouter) Start() {
	log.Println("Starting server")
	err := http.ListenAndServe(":8080", this.routerEngine)
	if err != nil {
		panic(err)
	}
}
