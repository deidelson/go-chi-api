package main

import (
	"github.com/deidelson/go-chi-api/pkg/api/movie"
	"github.com/deidelson/go-chi-api/pkg/api/public"
	"github.com/deidelson/go-chi-api/pkg/core/routing"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

func main() {

	router := routing.GetApiRouter()

	//TODO llevar middlewares al core
	router.AddGlobalMiddleware(middleware.RequestID)
	router.AddGlobalMiddleware(middleware.RealIP)
	router.AddGlobalMiddleware(middleware.Logger)
	router.AddGlobalMiddleware(middleware.Timeout(60 * time.Second))

	router.AddHandler(movie.GetMovieHandlerInstance())
	router.AddHandler(public.GetLoginHandlerInstance())

	router.Start()
}
