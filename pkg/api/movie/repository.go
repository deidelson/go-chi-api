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
	unexistentMovie    = errors.New("unexistent.movie")
	existentMovie      = errors.New("existent.movie")
)

type repositoryImpl struct {
	movies map[int]movie
}

func (repository *repositoryImpl) save(movie *movie) error {
	if _, exist := repository.movies[(*movie).Id]; exist {
		return existentMovie
	}
	repository.movies[movie.Id] = *movie
	return nil
}

func (repository *repositoryImpl) getById(id int) (*movie, error) {
	movie, exist := repository.movies[id]
	if exist {
		return &movie, nil
	}
	return nil, unexistentMovie
}

func (repository *repositoryImpl) findAll() []movie {
	movies := make([]movie, 0, 0)
	for _, value := range repository.movies {
		movies = append(movies, value)
	}
	return movies
}
