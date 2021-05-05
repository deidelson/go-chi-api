package routing

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

const (
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
	DELETE = "DELETE"
)

var (
	apiRouterInstance *ApiRouter
)

type ApiRouter struct {
	routerEngine *chi.Mux
}

func GetApiRouter() *ApiRouter {
	log.Println("Starting router")
	if apiRouterInstance == nil {
		apiRouterInstance = &ApiRouter{
			routerEngine: chi.NewRouter(),
		}
		log.Println("Using default middlewares")
		apiRouterInstance.routerEngine.Use(middleware.RequestID)
		apiRouterInstance.routerEngine.Use(middleware.RealIP)
		apiRouterInstance.routerEngine.Use(middleware.Logger)
		apiRouterInstance.routerEngine.Use(middleware.Recoverer)
	}
	return apiRouterInstance
}

//func (this *ApiRouter) AddGlobalMiddleware(handler ApiHandler)

func (this *ApiRouter) AddHandler(handler ApiHandler) {
	this.routerEngine.Route(handler.GetBasePath(), func(r chi.Router) {
		log.Println("Registering handler for route: ", handler.GetBasePath())
		if handler.GetMiddlewares() != nil && handler.GetMiddlewares().isNotEmpty() {
			r.Use(handler.GetMiddlewares()...)
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
