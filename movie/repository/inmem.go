package repository

import (
	"sync"

	"github.com/wagaru/ticket/movie/domain"
)

type inMemRepository struct {
	mux  sync.RWMutex
	data map[string]*domain.Movie
}

func NewInMemRepository() domain.Repository {
	return &inMemRepository{
		data: make(map[string]*domain.Movie),
	}
}

func (in *inMemRepository) Store(movie *domain.Movie) error {
	in.mux.Lock()
	defer in.mux.Unlock()
	if _, ok := in.data[movie.ID]; ok {
		return domain.ErrAlreadyExists
	}
	in.data[movie.ID] = movie
	return nil
}

func (in *inMemRepository) Fetch(ID string) (*domain.Movie, error) {
	in.mux.RLock()
	defer in.mux.RUnlock()
	movie, ok := in.data[ID]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return movie, nil
}

func (in *inMemRepository) FetchAll() ([]*domain.Movie, error) {
	in.mux.RLock()
	defer in.mux.RUnlock()
	data := make([]*domain.Movie, 0, len(in.data))
	for _, m := range in.data {
		data = append(data, m)
	}
	return data, nil
}
