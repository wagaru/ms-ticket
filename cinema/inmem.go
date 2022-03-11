package main

import "sync"

type Repository interface {
	Store(*Cinema) error
	Fetch(ID string) (*Cinema, error)
	FetchByCity(city City) ([]*Cinema, error)
	FetchAll() ([]*Cinema, error)
}

type inMemRepository struct {
	mux  sync.RWMutex
	data map[string]*Cinema
}

func NewInMemRepository() Repository {
	return &inMemRepository{
		data: make(map[string]*Cinema),
	}
}

func (in *inMemRepository) Store(cinema *Cinema) error {
	in.mux.Lock()
	defer in.mux.Unlock()
	if _, ok := in.data[cinema.ID]; ok {
		return ErrAlreadyExists
	}
	in.data[cinema.ID] = cinema
	return nil
}

func (in *inMemRepository) Fetch(ID string) (*Cinema, error) {
	in.mux.RLock()
	defer in.mux.RUnlock()
	cinema, ok := in.data[ID]
	if !ok {
		return nil, ErrNotFound
	}
	return cinema, nil
}

func (in *inMemRepository) FetchByCity(city City) ([]*Cinema, error) {
	in.mux.RLock()
	defer in.mux.RUnlock()
	res := make([]*Cinema, 0)
	for _, cinema := range in.data {
		if cinema.City == city {
			res = append(res, cinema)
		}
	}
	return res, nil
}

func (in *inMemRepository) FetchAll() ([]*Cinema, error) {
	in.mux.RLock()
	defer in.mux.RUnlock()
	res := make([]*Cinema, 0, len(in.data))
	for _, cinema := range in.data {
		res = append(res, cinema)
	}
	return res, nil
}
