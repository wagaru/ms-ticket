package service

import (
	"context"

	"github.com/wagaru/ticket/movie/domain"
)

type Service interface {
	// add new movie
	PostMovie(ctx context.Context, m *domain.Movie) error

	// get movie with movieID
	GetMovie(ctx context.Context, ID string) (*domain.Movie, error)

	// get all movies
	GetMovies(ctx context.Context) ([]*domain.Movie, error)
}

type service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &service{repo}
}

func (m *service) PostMovie(ctx context.Context, movie *domain.Movie) error {
	return m.repo.Store(movie)
}

func (m *service) GetMovie(ctx context.Context, ID string) (*domain.Movie, error) {
	return m.repo.Fetch(ID)
}

func (m *service) GetMovies(ctx context.Context) ([]*domain.Movie, error) {
	return m.repo.FetchAll()
}
