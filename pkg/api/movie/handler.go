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
		web.InternalServerError(w, web.ErrorResponse{Error: err.Error()})
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
	movie := &Movie{}
	if err := web.ReadBody(r.Body, movie); err != nil {
		web.InternalServerError(w, web.ErrorResponse{Error: err.Error()})
		return
	}
	_ = this.movieService.save(movie)
	web.Ok(w, nil)
}

func (this *handler) GetBasePath() string {
	return "/api/Movie"
}

func (this *handler) GetMiddlewares() routing.Middlewares {
	return routing.Middlewares{
		middleware.Jwt,
	}
}

func (this *handler) GetRoutes() []routing.ApiRoute {
	return []routing.ApiRoute{
		// swagger:route POST /api/movie saveMovie
		//
		// Guarda una pelicula
		//
		// Guarda una pelicula (no debe existir previamente)
		//
		//     Consumes:
		//	   - application/json
		//     Produces:
		//
		//     Schemes: https
		//
		//     Security:
		//
		//     Responses:
		//       default: Movie
		//       200: Movie
		//       500: ErrorResponse
		{
			Handler:  this.saveMovie,
			Method:   routing.POST,
			Endpoint: "/",
		},
		// swagger:route GET /api/movie listMovies
		//
		// Lista todas las peliculas
		//
		// Aca va la descripcion
		//
		//     Consumes:
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: https
		//
		//     Security:
		//
		//     Responses:
		//       default: []Movie
		//       200: []Movie
		//       500: ErrorResponse
		{
			Handler:  this.getMovies,
			Method:   routing.GET,
			Endpoint: "/",
		},
		// swagger:route GET /api/movie/{id} getMovieById
		//
		// Aca va el resumen
		//
		// Aca va la descripcion
		//
		//     Consumes:
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: https
		//
		//     Security:
		//
		//     Responses:
		//       default: Movie
		//       200: Movie
		//       500: ErrorResponse
		{
			Handler:  this.getById,
			Method:   routing.GET,
			Endpoint: "/{id}",
		},
	}

}
