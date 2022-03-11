package main

type Service interface {
	PostCinema(cinema *Cinema) error
	GetCinemasByCity(city City) ([]*Cinema, error)
	GetCinemas() ([]*Cinema, error)
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
