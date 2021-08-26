package movie

import (
	"errors"
)

type repository interface {
	save(movie *Movie) error
	getById(id int) (*Movie, error)
	findAll() []Movie
}

var (
	repositoryInstance repository
	unexistentMovie    = errors.New("unexistent.movie")
	existentMovie      = errors.New("existent.movie")
)

type repositoryImpl struct {
	movies map[int]Movie
}

func (repository *repositoryImpl) save(movie *Movie) error {
	if _, exist := repository.movies[(*movie).Id]; exist {
		return existentMovie
	}
	repository.movies[movie.Id] = *movie
	return nil
}

func (repository *repositoryImpl) getById(id int) (*Movie, error) {
	movie, exist := repository.movies[id]
	if exist {
		return &movie, nil
	}
	return nil, unexistentMovie
}

func (repository *repositoryImpl) findAll() []Movie {
	movies := make([]Movie, 0, 0)
	for _, value := range repository.movies {
		movies = append(movies, value)
	}
	return movies
}
