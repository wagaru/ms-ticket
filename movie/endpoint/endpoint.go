package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/wagaru/ticket/movie/service"
)

type Endpoints struct {
	PostMovieEndpoint endpoint.Endpoint
	GetMovieEndpoint  endpoint.Endpoint
	GetMoviesEndpoint endpoint.Endpoint
}

type PostMovieRequest struct {
	Movie service.Movie
}

type PostMovieResponse struct {
	Err error `json:"err,omitempty"`
}

func (p PostMovieResponse) GetError() error {
	return p.Err
}

type GetMovieRequest struct {
	ID string
}

type GetMovieResponse struct {
	Movie service.Movie
	Err   error `json:"err,omitempty"`
}

func (p GetMovieResponse) GetError() error {
	return p.Err
}

type GetMoviesRequest struct {
}

type GetMoviesResponse struct {
	Movies []service.Movie `json:"movies,omitempty"`
	Err    error           `json:"error,omitempty"`
}

func (p GetMoviesResponse) GetError() error {
	return p.Err
}

func MakeEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		PostMovieEndpoint: makePostMovieEndpoint(svc),
		GetMovieEndpoint:  makeGetMovieEndpoint(svc),
		GetMoviesEndpoint: makeGetMoviesEndpoint(svc),
	}
}

func makePostMovieEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostMovieRequest)
		err := svc.PostMovie(ctx, req.Movie)
		return PostMovieResponse{Err: err}, nil
	}
}

func makeGetMovieEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMovieRequest)
		movie, err := svc.GetMovie(ctx, req.ID)
		return GetMovieResponse{Movie: movie, Err: err}, nil
	}
}

func makeGetMoviesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		movies, err := svc.GetMovies(ctx)
		return GetMoviesResponse{Movies: movies, Err: err}, nil
	}
}

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
