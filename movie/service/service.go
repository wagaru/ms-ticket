package service

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Movie struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Duration    int       `json:"duration"`
	Desc        string    `json:"desc"`
	ComeOutDate time.Time `json:"come_out_date,omitempty"`
}

var (
	ErrAlreadyExists = errors.New("movie already exists")
	ErrNotFound      = errors.New("not found")
)

type Service interface {
	PostMovie(ctx context.Context, m Movie) error
	GetMovie(ctx context.Context, ID string) (Movie, error)
	GetMovies(ctx context.Context) ([]Movie, error)
}

type inMemService struct {
	mutex sync.RWMutex
	data  map[string]Movie
}

func NewInMemService() Service {
	return &inMemService{
		data: make(map[string]Movie),
	}
}

func (m *inMemService) PostMovie(ctx context.Context, movie Movie) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.data[movie.ID]; ok {
		return ErrAlreadyExists
	}
	m.data[movie.ID] = movie
	return nil
}

func (m *inMemService) GetMovie(ctx context.Context, ID string) (Movie, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	movie, ok := m.data[ID]
	if !ok {
		return Movie{}, ErrNotFound
	}
	return movie, nil
}

func (m *inMemService) GetMovies(ctx context.Context) ([]Movie, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	res := make([]Movie, 0)
	for _, movie := range m.data {
		res = append(res, movie)
	}
	return res, nil
}
