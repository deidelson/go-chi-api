package main

import (
	"github.com/deidelson/go-chi-api/pkg/core/routing"
	"github.com/deidelson/go-chi-api/pkg/core/web"
	"net/http"
)

type handler struct {
	
}

func (this *handler) GetBasePath() string {
	return "/api"
}
func (this *handler) GetRoutes() []routing.ApiRoute {
	return []routing.ApiRoute {
		{
			Handler: func(w http.ResponseWriter, r *http.Request) {
				web.Ok(w, "Hello world desde chi")
			},
			Method: routing.GET,
			Endpoint: "/hello",
		},

	}
}
func (this *handler) GetMiddlewares() routing.Middlewares {
	return nil
}

func main() {

	router := routing.GetApiRouter()

	router.AddHandler(&handler{})

	router.Start()
}
