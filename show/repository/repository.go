package repository

import (
	"github.com/wagaru/ticket/show/domain"
)

type Repository interface {
	StoreMovie(movie *domain.Movie) error
	FetchMovies(cinemaID uint) ([]*domain.Movie, error)
	StoreCinema(cinema *domain.Cinema) error
	FetchAllCinemas() ([]*domain.Cinema, error)
	StoreShow(*domain.Show) error
	FetchShows(movieID uint, cinemaID uint) ([]*domain.Show, error)
	FetchShowSeats(showID uint) ([]*domain.CinemaSeat, error)
}
