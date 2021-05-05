package movie

import (
	"errors"
)

type repository interface {
	save(movie *movie) error
	getById(id int) (*movie, error)
	findAll() []movie
}

var (
	repositoryInstance repository
)

type repositoryImpl struct {
	movies map[int]movie
}

func (repository *repositoryImpl) save(movie *movie) error {
	if _, exist := repository.movies[(*movie).Id]; exist {
		return errors.New("Ya existe la pelicula")
	}
	repository.movies[movie.Id] = *movie
	return nil
}

func (repository *repositoryImpl) getById(id int) (*movie, error) {
	movie, exist := repository.movies[id]
	if exist {
		return &movie, nil
	}
	return nil, errors.New("No existe la pelicula")
}

func (repository *repositoryImpl) findAll() []movie {
	movies := make([]movie, 0, 0)
	for _, value := range repository.movies {
		movies = append(movies, value)
	}
	return movies
}
