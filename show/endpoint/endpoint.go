package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/wagaru/ticket/pkg/common_error"
	"github.com/wagaru/ticket/show/domain"
	"github.com/wagaru/ticket/show/service"
)

type Error struct {
	Err error `json:"error,omitempty"`
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

type GetShowRequest struct {
	MovieID  uint `json:"movie_id"`
	CinemaID uint `json:"cinema_id"`
}

type GetShowResponse struct {
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

func MakeGetShowEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetShowRequest)
		if !ok {
			return GetShowResponse{Error: Error{common_error.ErrInvalidInput}}, nil
		}
		shows, err := svc.GetShow(req.MovieID, req.CinemaID)
		return GetShowResponse{shows, Error{err}}, nil
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
