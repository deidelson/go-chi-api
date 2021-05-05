package main

import (
	"github.com/deidelson/go-chi-api/pkg/core/routing"
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
				w.Write([]byte("Hello world"))
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
