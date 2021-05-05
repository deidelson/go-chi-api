package movie

import (
	"github.com/deidelson/go-chi-api/pkg/core/convertion"
	"github.com/deidelson/go-chi-api/pkg/core/routing"
	"github.com/deidelson/go-chi-api/pkg/core/web"
	"github.com/go-chi/chi/v5"
	"net/http"
)

var (
	handlerInstance *handler
)

type handler struct {
	movieService service
}

func (handler *handler) getMovies(w http.ResponseWriter, r *http.Request) {
	web.Ok(w, handler.movieService.findAll())
}

func (handler *handler) getById(w http.ResponseWriter, r *http.Request) {
	//TODO move chi.URLParam to utility for technology abstraction
	id, err := convertion.StringToInt(chi.URLParam(r, "id"))

	if err != nil {
		web.InternalServerError(w, err.Error())
		return
	}

	movie, err := handler.movieService.getById(id)

	if err != nil {
		web.Confict(w, err.Error())
		return
	}

	web.Ok(w, movie)

}

func (handler *handler) saveMovie(w http.ResponseWriter, r *http.Request) {
	movie := &movie{}
	if err := web.ReadBody(r.Body, movie); err != nil {
		web.InternalServerError(w, err.Error())
		return
	}
	_ = handler.movieService.save(movie)
	web.Ok(w, nil)
}

func (this *handler) GetBasePath() string {
	return "/api/movie"
}

func (this *handler) GetMiddlewares() routing.Middlewares {
	return routing.Middlewares{}
}

func (this *handler) GetRoutes() []routing.ApiRoute {
	return []routing.ApiRoute{
		{
			Handler:  this.saveMovie,
			Method:   routing.POST,
			Endpoint: "/",
		},
		{
			Handler: this.getMovies,
			Method:  routing.GET,

			Endpoint: "/",
		},
		{
			Handler:  this.getById,
			Method:   routing.GET,
			Endpoint: "/{id}",
		},
	}

}
