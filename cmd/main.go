package main

import (
	"github.com/deidelson/go-chi-api/pkg/api/movie"
	"github.com/deidelson/go-chi-api/pkg/core/routing"
)

func main() {

	router := routing.GetApiRouter()

	router.AddHandler(movie.GetMovieControllerInstance())

	router.Start()
}
