package movie

type service interface {
	save(movie *movie) error
	getById(id int) (*movie, error)
	findAll() []movie
}

var (
	serviceInstance service
)

type serviceImpl struct {
	movieRepository repository
}

func (service *serviceImpl) save(movie *movie) error {
	return service.movieRepository.save(movie)
}

func (service *serviceImpl) getById(id int) (*movie, error) {
	return service.movieRepository.getById(id)
}

func (service *serviceImpl) findAll() []movie {
	return service.movieRepository.findAll()
}
