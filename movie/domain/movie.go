package domain

import (
	"errors"
	"time"
)

var (
	ErrAlreadyExists = errors.New("movie already exists")
	ErrNotFound      = errors.New("not found")
)

type Movie struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Duration    int        `json:"duration"`
	Desc        string     `json:"desc"`
	ComeOutDate *time.Time `json:"come_out_date,omitempty"` // use pointer since zero value for time.Time is struct, and omitempty not
}

type Repository interface {
	Store(*Movie) error
	Fetch(ID string) (*Movie, error)
	FetchAll() ([]*Movie, error)
}
