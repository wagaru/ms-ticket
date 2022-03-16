package service

import (
	"github.com/wagaru/ticket/show/domain"
	"github.com/wagaru/ticket/show/repository"
)

type Service interface {

	// add new movie
	AddMovie(m *domain.Movie) error

	// get all movies played in cinema
	GetCinemaMovies(cinemaID uint) ([]*domain.Movie, error)

	// add new cinema
	AddCinema(cinema *domain.Cinema) error

	// get all cinemas
	GetCinemas() ([]*domain.Cinema, error)

	// add new show
	AddShow(show *domain.Show) error

	// get all shows playing movie with movieID and in cinema cinemaID
	GetShows(movieID uint, cinemaID uint) ([]*domain.Show, error)

	// get show seats
	GetShowSeats(showID uint) ([]*domain.CinemaSeat, error)
}

type service struct {
	repo repository.Repository
	//routing RoutingService
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
		//routing: routing,
	}
}

func (s *service) AddMovie(movie *domain.Movie) error {
	return s.repo.StoreMovie(movie)
}

func (s *service) GetCinemaMovies(cinemaID uint) ([]*domain.Movie, error) {
	return s.repo.FetchMovies(cinemaID)
}

func (s *service) AddCinema(cinema *domain.Cinema) error {
	return s.repo.StoreCinema(cinema)
}

func (s *service) GetCinemas() ([]*domain.Cinema, error) {
	return s.repo.FetchAllCinemas()
}

func (s *service) AddShow(show *domain.Show) error {
	return s.repo.StoreShow(show)
}

func (s *service) GetShows(movieID uint, cinemaID uint) ([]*domain.Show, error) {
	return s.repo.FetchShows(movieID, cinemaID)
}

func (s *service) GetShowSeats(showID uint) ([]*domain.CinemaSeat, error) {
	return s.repo.FetchShowSeats(showID)
}
