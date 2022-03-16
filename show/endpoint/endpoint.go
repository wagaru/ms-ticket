package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/wagaru/ticket/pkg/common_error"
	"github.com/wagaru/ticket/show/domain"
	"github.com/wagaru/ticket/show/service"
)

type Endpoints struct {
	AddMovieEndpoint        endpoint.Endpoint
	GetCinemaMoviesEndpoint endpoint.Endpoint
	AddCinemaEndpoint       endpoint.Endpoint
	GetCinemasEndpoint      endpoint.Endpoint
	AddShowEndpoint         endpoint.Endpoint
	GetShowsEndpoint        endpoint.Endpoint
	GetShowSeatsEndpoint    endpoint.Endpoint
}

func MakeEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		AddMovieEndpoint:        MakeAddMovieEndpoint(svc),
		GetCinemaMoviesEndpoint: MakeGetCinemaMoviesEndpoint(svc),
		AddCinemaEndpoint:       MakeAddCinemaEndpoint(svc),
		GetCinemasEndpoint:      MakeGetCinemasEndpoint(svc),
		AddShowEndpoint:         MakeAddShowEndpoint(svc),
		GetShowsEndpoint:        MakeGetShowsEndpoint(svc),
		GetShowSeatsEndpoint:    MakeGetShowSeatsEndpoint(svc),
	}
}

type Error struct {
	Err error `json:"error,omitempty"`
}

func (e *Error) GetError() error {
	return e.Err
}

type AddMovieRequest struct {
	Movie *domain.Movie
}

type AddMovieResponse struct {
	Error
}

type GetCinemaMoviesRequest struct {
	CinemaID uint `json:"cinema_id"`
}

type GetCinemaMoviesResponse struct {
	Movies []*domain.Movie `json:"movies"`
	Error
}

type AddCinemaRequest struct {
	Cinema *domain.Cinema
}

type AddCinemaResponse struct {
	Error
}

type GetCinemasRequest struct {
}

type GetCinemasResponse struct {
	Cinemas []*domain.Cinema `json:"cinemas"`
	Error
}

type AddShowRequest struct {
	Show *domain.Show
}

type AddShowResponse struct {
	Error
}

type GetShowsRequest struct {
	MovieID  uint `json:"movie_id"`
	CinemaID uint `json:"cinema_id"`
}

type GetShowsResponse struct {
	Shows []*domain.Show `json:"shows"`
	Error
}

type GetShowSeatsRequest struct {
	ShowID uint `json:"show_id"`
}

type GetShowSeatsResponse struct {
	Seats []*domain.CinemaSeat `json:"seats"`
	Error
}

func MakeAddMovieEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(AddMovieRequest)
		if !ok {
			return AddMovieResponse{Error: Error{common_error.ErrInvalidInput}}, nil
		}
		return AddMovieResponse{Error: Error{svc.AddMovie(req.Movie)}}, nil
	}
}

func MakeGetCinemaMoviesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetCinemaMoviesRequest)
		if !ok {
			return GetCinemaMoviesResponse{Error: Error{common_error.ErrInvalidInput}}, nil
		}
		movies, err := svc.GetCinemaMovies(req.CinemaID)
		return GetCinemaMoviesResponse{movies, Error{err}}, nil
	}
}

func MakeAddCinemaEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(AddCinemaRequest)
		if !ok {
			return AddCinemaResponse{Error: Error{common_error.ErrInvalidInput}}, nil
		}
		return AddCinemaResponse{Error: Error{svc.AddCinema(req.Cinema)}}, nil
	}
}

func MakeGetCinemasEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		cinemas, err := svc.GetCinemas()
		return GetCinemasResponse{cinemas, Error{err}}, nil
	}
}

func MakeAddShowEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(AddShowRequest)
		if !ok {
			return AddShowResponse{Error: Error{common_error.ErrInvalidInput}}, nil
		}
		return AddShowResponse{Error: Error{svc.AddShow(req.Show)}}, nil
	}
}

func MakeGetShowsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetShowsRequest)
		if !ok {
			return GetShowsResponse{Error: Error{common_error.ErrInvalidInput}}, nil
		}
		shows, err := svc.GetShows(req.MovieID, req.CinemaID)
		return GetShowsResponse{shows, Error{err}}, nil
	}
}

func MakeGetShowSeatsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetShowSeatsRequest)
		if !ok {
			return GetShowSeatsResponse{Error: Error{common_error.ErrInvalidInput}}, nil
		}
		seats, err := svc.GetShowSeats(req.ShowID)
		return GetShowSeatsResponse{seats, Error{err}}, nil
	}
}
