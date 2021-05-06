package movie

import "github.com/deidelson/go-chi-api/pkg/core/routing"

func GetMovieHandlerInstance() routing.ApiHandler {
	if handlerInstance == nil {
		handlerInstance = &handler{
			movieService: getMovieServiceInstance(getMovieRepositoryInstance()),
		}
	}
	return handlerInstance
}

func getMovieServiceInstance(movieRepository repository) service {
	if serviceInstance == nil {
		serviceInstance = &serviceImpl{
			movieRepository: movieRepository,
		}
	}
	return serviceInstance
}

func getMovieRepositoryInstance() repository {
	if repositoryInstance == nil {
		repositoryInstance = &repositoryImpl{
			movies: make(map[int]movie),
		}
		repositoryInstance.save(&movie{
			Id:       0,
			Name:     "Django",
			Director: "Tarantino",
			Year:     2009,
		})
	}
	return repositoryInstance
}
