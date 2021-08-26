package movie

type service interface {
	save(movie *Movie) error
	getById(id int) (*Movie, error)
	findAll() []Movie
}

var (
	serviceInstance service
)

type serviceImpl struct {
	movieRepository repository
}

func (service *serviceImpl) save(movie *Movie) error {
	return service.movieRepository.save(movie)
}

func (service *serviceImpl) getById(id int) (*Movie, error) {
	return service.movieRepository.getById(id)
}

func (service *serviceImpl) findAll() []Movie {
	return service.movieRepository.findAll()
}
