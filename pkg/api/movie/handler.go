package movie

import (
	"github.com/deidelson/go-chi-api/pkg/core/convertion"
	"github.com/deidelson/go-chi-api/pkg/core/routing"
	"github.com/deidelson/go-chi-api/pkg/core/routing/middleware"
	"github.com/deidelson/go-chi-api/pkg/core/web"
	"net/http"
)

var (
	handlerInstance *handler
)

type handler struct {
	movieService service
}

func (this *handler) getMovies(w http.ResponseWriter, r *http.Request) {
	web.Ok(w, this.movieService.findAll())
}

func (this *handler) getById(w http.ResponseWriter, r *http.Request) {
	id, err := convertion.StringToInt(web.GetPathVariable(r, "id"))

	if err != nil {
		web.InternalServerError(w, err.Error())
		return
	}

	movie, err := this.movieService.getById(id)

	if err != nil {
		web.Conflict(w, err.Error())
		return
	}

	web.Ok(w, movie)

}

func (this *handler) saveMovie(w http.ResponseWriter, r *http.Request) {
	movie := &movie{}
	if err := web.ReadBody(r.Body, movie); err != nil {
		web.InternalServerError(w, err.Error())
		return
	}
	_ = this.movieService.save(movie)
	web.Ok(w, nil)
}

func (this *handler) GetBasePath() string {
	return "/api/movie"
}

func (this *handler) GetMiddlewares() routing.Middlewares {
	return routing.Middlewares{
		middleware.Jwt,
	}
}

func (this *handler) GetRoutes() []routing.ApiRoute {
	return []routing.ApiRoute{
		{
			Handler:  this.saveMovie,
			Method:   routing.POST,
			Endpoint: "/",
		},
		{
			Handler:  this.getMovies,
			Method:   routing.GET,
			Endpoint: "/",
		},
		{
			Handler:  this.getById,
			Method:   routing.GET,
			Endpoint: "/{id}",
		},
	}

}
