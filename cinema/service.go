package main

type Service interface {
	// add new cinema
	PostCinema(cinema *Cinema) error

	// get cinemas by city
	GetCinemasByCity(city City) ([]*Cinema, error)

	// get all cinemas
	GetCinemas() ([]*Cinema, error)

	// get single cinema by cinemaID
	GetCinema(ID string) (*Cinema, error)
}

type ServiceMiddleware func(Service) Service

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) PostCinema(cinema *Cinema) error {
	return s.repo.Store(cinema)
}

func (s *service) GetCinemasByCity(city City) ([]*Cinema, error) {
	return s.repo.FetchByCity(city)
}

func (s *service) GetCinemas() ([]*Cinema, error) {
	return s.repo.FetchAll()
}

func (s *service) GetCinema(ID string) (*Cinema, error) {
	return s.repo.Fetch(ID)
}
